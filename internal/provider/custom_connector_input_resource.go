package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/client/parameter"
	"terraform-provider-trocco/internal/provider/model"
	troccoPlanModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                   = &customConnectorInputResource{}
	_ resource.ResourceWithConfigure      = &customConnectorInputResource{}
	_ resource.ResourceWithImportState    = &customConnectorInputResource{}
	_ resource.ResourceWithValidateConfig = &customConnectorInputResource{}
)

var urlSchemeRegex = regexp.MustCompile(`^https?://`)
var commaSeparatedCodesRegex = regexp.MustCompile(`^\d+(,\d+)*$`)

// Defaults mirroring the server-side column defaults. The server assigns every
// endpoint attribute unconditionally, so the request must always carry a value:
// a missing key is written as NULL and rejected by the presence validations
// instead of falling back to the column default. Pinning a static default here
// keeps these attributes optional in configuration while guaranteeing that.
// `name` and `path` have no column default and stay required.
const (
	customConnectorInputDefaultMethod            = "GET"
	customConnectorInputDefaultJsonpathRoot      = "$.*"
	customConnectorInputDefaultSuccessCodes      = "200"
	customConnectorInputDefaultNotRetryableCodes = "400,401,403,404"
	customConnectorInputDefaultRequestTimeoutSec = 30
)

func NewCustomConnectorInputResource() resource.Resource {
	return &customConnectorInputResource{}
}

type customConnectorInputResource struct {
	client *client.TroccoClient
}

func (r *customConnectorInputResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_connector_input"
}

func (r *customConnectorInputResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.TroccoClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.TroccoClient, got: %T", req.ProviderData),
		)
		return
	}
	r.client = c
}

func customConnectorFieldOptionAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "Field name.",
		},
		"display_name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "Display name shown in the transfer setting UI.",
		},
		"default_value": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "Default value. Required when `is_editable` is `false`.",
		},
		"is_editable": schema.BoolAttribute{
			Required:            true,
			MarkdownDescription: "Whether the transfer setting can override this value.",
		},
		"is_required": schema.BoolAttribute{
			Required:            true,
			MarkdownDescription: "Whether the transfer setting must supply this value.",
		},
	}
}

func customConnectorPaginatorAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"inject_into": schema.StringAttribute{
			Optional: true,
			Computed: true,
			Validators: []validator.String{
				stringvalidator.OneOf("query", "request_body"),
			},
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
			MarkdownDescription: "Where to inject pagination parameters. `request_body` requires the endpoint's `method` to be `POST` and `request_body` to be set. When omitted, the server applies its default.",
		},
		"page_increment_strategy": schema.SingleNestedAttribute{
			Optional: true,
			Validators: []validator.Object{
				objectvalidator.ExactlyOneOf(
					path.MatchRelative().AtParent().AtName("offset_increment_strategy"),
					path.MatchRelative().AtParent().AtName("cursor_based_strategy"),
				),
			},
			Attributes: map[string]schema.Attribute{
				"page_size": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Number of records per page.",
				},
				"last_page_size": schema.StringAttribute{
					Optional: true,
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("total_pages"),
							path.MatchRelative().AtParent().AtName("stop_on_page"),
						),
						stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("max_request_count")),
					},
					MarkdownDescription: "Size of the last page. A number or a JSONPath expression (e.g. `$.meta.last_page_size`).",
				},
				"start_from_page": schema.Int64Attribute{
					Required:            true,
					MarkdownDescription: "Page number to start from.",
				},
				"stop_on_page": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Page number to stop on.",
				},
				"total_pages": schema.StringAttribute{
					Optional:            true,
					MarkdownDescription: "Total number of pages. A number or a JSONPath expression (e.g. `$.meta.total_pages`).",
				},
				"max_request_count": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Maximum number of requests to issue.",
				},
			},
			MarkdownDescription: "Page-increment pagination strategy. Exactly one of `page_increment_strategy` / `offset_increment_strategy` / `cursor_based_strategy` must be set.",
		},
		"offset_increment_strategy": schema.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]schema.Attribute{
				"page_size": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Number of records per page.",
				},
				"last_page_size": schema.StringAttribute{
					Optional: true,
					Validators: []validator.String{
						stringvalidator.ExactlyOneOf(
							path.MatchRelative().AtParent().AtName("total_records"),
							path.MatchRelative().AtParent().AtName("stop_on_offset"),
						),
						stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("max_request_count")),
					},
					MarkdownDescription: "Size of the last page. A number or a JSONPath expression (e.g. `$.meta.last_page_size`).",
				},
				"start_from_offset": schema.Int64Attribute{
					Required:            true,
					MarkdownDescription: "Offset to start from.",
				},
				"stop_on_offset": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Offset to stop on.",
				},
				"total_records": schema.StringAttribute{
					Optional:            true,
					MarkdownDescription: "Total number of records. A number or a JSONPath expression (e.g. `$.meta.total_records`).",
				},
				"max_request_count": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Maximum number of requests to issue.",
				},
			},
			MarkdownDescription: "Offset-increment pagination strategy.",
		},
		"cursor_based_strategy": schema.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]schema.Attribute{
				"page_size": schema.Int64Attribute{
					Optional:            true,
					MarkdownDescription: "Number of records per page.",
				},
				"cursor_value": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Cursor value.",
				},
			},
			MarkdownDescription: "Cursor-based pagination strategy.",
		},
		"page_size_option": schema.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]schema.Attribute{
				"field_name": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Field name used to pass the page size.",
				},
			},
			MarkdownDescription: "Where to pass the page size.",
		},
		"page_token_option": schema.SingleNestedAttribute{
			Optional: true,
			Attributes: map[string]schema.Attribute{
				"field_name": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Field name used to pass the page token (page number/offset/cursor).",
				},
			},
			MarkdownDescription: "Where to pass the page token (page number/offset/cursor).",
		},
	}
}

func (r *customConnectorInputResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO custom connector input (transfer source) definition resource.\n\n" +
			"`grant_type = \"authorization_code\"` can be declared here, but the resulting connection cannot be fully " +
			"authorized via the API/Terraform: TROCCO requires an interactive OAuth2 authorization flow performed in " +
			"the browser (login, consent, and the `code` callback). Prefer `grant_type = \"client_credentials\"` for " +
			"connections that need to work end-to-end via Terraform.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The ID of the custom connector.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the custom connector.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "A memo describing the custom connector.",
			},
			"url": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(urlSchemeRegex, "must start with http:// or https://"),
				},
				MarkdownDescription: "Base URL. Must start with `http://` or `https://`; private IPs are rejected by the server.",
			},
			"auth_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("api_key", "oauth2"),
				},
				MarkdownDescription: "Authentication method. Determines which other attributes are required (e.g. `access_token_uri` when `oauth2`).",
			},
			"auth_header_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Authentication header name. When omitted, the server applies its default.",
			},
			"auth_header_scheme": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Authentication header scheme. When omitted, the server applies its default.",
			},
			"grant_type": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("authorization_code", "client_credentials"),
				},
				MarkdownDescription: "OAuth2 grant type. Only meaningful when `auth_type` is `oauth2`; the server discards this (along with `auth_uri`/`access_token_uri`) when `auth_type` is changed away from `oauth2`.",
			},
			"auth_uri": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(urlSchemeRegex, "must start with http:// or https://"),
				},
				MarkdownDescription: "OAuth2 authorization endpoint URI. Required when `grant_type` is `authorization_code`.",
			},
			"access_token_uri": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(urlSchemeRegex, "must start with http:// or https://"),
				},
				MarkdownDescription: "OAuth2 token endpoint URI. Required when `auth_type` is `oauth2`.",
			},
			"created_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The time the custom connector was created.",
			},
			"updated_at": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The time the custom connector was last updated.",
			},
			"endpoints": schema.ListNestedAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{
					troccoPlanModifier.EmptyListForNull(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "Server-assigned endpoint ID.",
						},
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Endpoint name. Must be unique within the connector; used to match endpoints across applies (renaming an endpoint is treated as delete+create).",
						},
						"path": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Endpoint path.",
						},
						"method": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString(customConnectorInputDefaultMethod),
							Validators: []validator.String{
								stringvalidator.OneOf("GET", "POST"),
							},
							MarkdownDescription: "HTTP method. Defaults to `GET`.",
						},
						"request_body": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtMost(16383),
							},
							MarkdownDescription: "Request body (up to 16383 characters).",
						},
						"jsonpath_root": schema.StringAttribute{
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString(customConnectorInputDefaultJsonpathRoot),
							MarkdownDescription: "JSONPath used to extract the record array from the response. Defaults to `$.*`.",
						},
						"success_codes": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString(customConnectorInputDefaultSuccessCodes),
							Validators: []validator.String{
								stringvalidator.RegexMatches(commaSeparatedCodesRegex, "must be a comma-separated list of numbers"),
							},
							MarkdownDescription: "Comma-separated HTTP status codes treated as success. Defaults to `200`.",
						},
						"not_retryable_codes": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString(customConnectorInputDefaultNotRetryableCodes),
							Validators: []validator.String{
								stringvalidator.RegexMatches(commaSeparatedCodesRegex, "must be a comma-separated list of numbers"),
							},
							MarkdownDescription: "Comma-separated HTTP status codes that must not be retried. Defaults to `400,401,403,404`.",
						},
						"request_timeout_sec": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							Default:  int64default.StaticInt64(customConnectorInputDefaultRequestTimeoutSec),
							Validators: []validator.Int64{
								int64validator.Between(1, 1800),
							},
							MarkdownDescription: "Seconds to wait for a response before timing out (1-1800). Defaults to 30.",
						},
						"query_parameters": schema.ListNestedAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.List{
								troccoPlanModifier.EmptyListForNull(),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: customConnectorFieldOptionAttributes(),
								PlanModifiers: []planmodifier.Object{
									&troccoPlanModifier.CustomConnectorFieldOptionPlanModifier{},
								},
							},
							MarkdownDescription: "Query parameter definitions. Set to `[]` to clear.",
						},
						"headers": schema.ListNestedAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.List{
								troccoPlanModifier.EmptyListForNull(),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: customConnectorFieldOptionAttributes(),
								PlanModifiers: []planmodifier.Object{
									&troccoPlanModifier.CustomConnectorFieldOptionPlanModifier{},
								},
							},
							MarkdownDescription: "Header definitions. Set to `[]` to clear.",
						},
						"path_parameters": schema.ListNestedAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.List{
								troccoPlanModifier.EmptyListForNull(),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "Path parameter name.",
									},
									"value": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "Value. Must appear in the endpoint's `path` as `{value}`, and is replaced there by the value a transfer setting supplies for this parameter.",
									},
								},
							},
							MarkdownDescription: "Path parameter definitions. Set to `[]` to clear.",
						},
						"request_body_parameters": schema.ListNestedAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.List{
								troccoPlanModifier.EmptyListForNull(),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: customConnectorFieldOptionAttributes(),
								PlanModifiers: []planmodifier.Object{
									&troccoPlanModifier.CustomConnectorFieldOptionPlanModifier{},
								},
							},
							MarkdownDescription: "Request body parameter definitions. Each `name` can be referenced as a `{name}` placeholder inside `request_body`, and a transfer setting supplies its value. Set to `[]` to clear.",
						},
						"paginator": schema.SingleNestedAttribute{
							Optional: true,
							PlanModifiers: []planmodifier.Object{
								&troccoPlanModifier.CustomConnectorPaginatorPlanModifier{},
							},
							Attributes:          customConnectorPaginatorAttributes(),
							MarkdownDescription: "Pagination settings.",
						},
					},
				},
				MarkdownDescription: "Ordered list of endpoint definitions. Set to `[]` to clear.",
			},
		},
	}
}

func (r *customConnectorInputResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var plan model.CustomConnectorInputModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.AuthType.IsNull() && !plan.AuthType.IsUnknown() && plan.AuthType.ValueString() == "oauth2" {
		if plan.AccessTokenURI.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("access_token_uri"),
				"Missing Access Token URI",
				"`access_token_uri` is required when `auth_type` is \"oauth2\".",
			)
		}
	}
	if !plan.GrantType.IsNull() && !plan.GrantType.IsUnknown() && plan.GrantType.ValueString() == "authorization_code" {
		if plan.AuthURI.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("auth_uri"),
				"Missing Auth URI",
				"`auth_uri` is required when `grant_type` is \"authorization_code\".",
			)
		}
	}

	endpoints, diags := extractCustomConnectorEndpointModels(ctx, plan.Endpoints)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	seen := make(map[string]bool, len(endpoints))
	for i, e := range endpoints {
		endpointPath := path.Root("endpoints").AtListIndex(i)

		if !e.Name.IsUnknown() {
			name := e.Name.ValueString()
			if seen[name] {
				resp.Diagnostics.AddAttributeError(
					endpointPath.AtName("name"),
					"Duplicate Endpoint Name",
					fmt.Sprintf("Endpoint name %q is duplicated. Endpoint names must be unique within a connector.", name),
				)
			}
			seen[name] = true
		}

		// The endpoint path is a plain string here: path parameters are
		// substituted by a literal `{name}` replacement, so no placeholder
		// syntax guidance is needed.
		resp.Diagnostics.Append(validateCustomConnectorPathParameters(ctx, endpointPath, e.Path, e.PathParameters, "")...)
	}
}

// validateCustomConnectorPathParameters mirrors the server-side validation
// shared by both custom connector definitions, which requires each path
// parameter's `value` to appear in the endpoint's `path` as `{value}`. Without
// it the server answers with an opaque 400 that names no attribute, since the
// failure surfaces through validates_associated.
//
// pathSyntaxHint is appended to the error detail by callers whose `path` is a
// template rather than a literal string, and therefore needs the placeholder
// spelled a particular way for the substring match to succeed.
func validateCustomConnectorPathParameters(
	ctx context.Context,
	endpointPath path.Path,
	endpointPathValue types.String,
	pathParameters types.List,
	pathSyntaxHint string,
) diag.Diagnostics {
	var diags diag.Diagnostics
	if endpointPathValue.IsNull() || endpointPathValue.IsUnknown() {
		return diags
	}
	params, d := extractCustomConnectorPathParameterModels(ctx, pathParameters)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}

	pathValue := endpointPathValue.ValueString()
	for i, p := range params {
		if p.Value.IsNull() || p.Value.IsUnknown() {
			continue
		}
		value := p.Value.ValueString()
		if strings.Contains(pathValue, "{"+value+"}") {
			continue
		}
		detail := fmt.Sprintf("Path parameter value %q must appear in the endpoint path as {%s}.", value, value)
		if pathSyntaxHint != "" {
			detail += " " + pathSyntaxHint
		}
		diags.AddAttributeError(
			endpointPath.AtName("path_parameters").AtListIndex(i).AtName("value"),
			"Path Parameter Not In Path",
			detail,
		)
	}
	return diags
}

// extractCustomConnectorPathParameterModels converts a `path_parameters` list
// into a Go slice, keeping Null/Unknown elements intact so callers can skip
// values that are not yet known at validation time.
func extractCustomConnectorPathParameterModels(ctx context.Context, list types.List) ([]model.CustomConnectorPathParameterModel, diag.Diagnostics) {
	if list.IsNull() || list.IsUnknown() {
		return []model.CustomConnectorPathParameterModel{}, nil
	}
	var params []model.CustomConnectorPathParameterModel
	diags := list.ElementsAs(ctx, &params, false)
	if diags.HasError() {
		return nil, diags
	}
	return params, nil
}

func (r *customConnectorInputResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan model.CustomConnectorInputModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoints, diags := extractCustomConnectorEndpointModels(ctx, plan.Endpoints)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	endpointInputs, diags := buildCustomConnectorEndpointInputs(ctx, endpoints, nil)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := parameter.CreateCustomConnectorInputInput{
		Name:        plan.Name.ValueString(),
		Description: model.NewNullableString(plan.Description),
		URL:         plan.URL.ValueString(),
		AuthType:    plan.AuthType.ValueString(),
		Endpoints:   endpointInputs,
	}
	if !plan.AuthHeaderName.IsNull() && !plan.AuthHeaderName.IsUnknown() {
		input.SetAuthHeaderName(plan.AuthHeaderName.ValueString())
	}
	if !plan.AuthHeaderScheme.IsNull() && !plan.AuthHeaderScheme.IsUnknown() {
		input.SetAuthHeaderScheme(plan.AuthHeaderScheme.ValueString())
	}
	if !plan.GrantType.IsNull() && !plan.GrantType.IsUnknown() {
		input.SetGrantType(plan.GrantType.ValueString())
	}
	if !plan.AuthURI.IsNull() && !plan.AuthURI.IsUnknown() {
		input.SetAuthURI(plan.AuthURI.ValueString())
	}
	if !plan.AccessTokenURI.IsNull() && !plan.AccessTokenURI.IsUnknown() {
		input.SetAccessTokenURI(plan.AccessTokenURI.ValueString())
	}

	def, err := r.client.CreateCustomConnectorInput(&input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating custom connector input",
			fmt.Sprintf("Unable to create custom connector input, got error: %s", err),
		)
		return
	}

	state, err := model.NewCustomConnectorInputModel(ctx, def, endpoints)
	if err != nil {
		resp.Diagnostics.AddError("Converting custom connector input", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *customConnectorInputResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var prior model.CustomConnectorInputModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	if resp.Diagnostics.HasError() {
		return
	}

	def, err := r.client.GetCustomConnectorInput(prior.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading custom connector input",
			fmt.Sprintf("Unable to read custom connector input, got error: %s", err),
		)
		return
	}

	refEndpoints, diags := extractCustomConnectorEndpointModels(ctx, prior.Endpoints)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state, err := model.NewCustomConnectorInputModel(ctx, def, refEndpoints)
	if err != nil {
		resp.Diagnostics.AddError("Converting custom connector input", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *customConnectorInputResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, prior model.CustomConnectorInputModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	if resp.Diagnostics.HasError() {
		return
	}

	planEndpoints, diags := extractCustomConnectorEndpointModels(ctx, plan.Endpoints)
	resp.Diagnostics.Append(diags...)
	priorEndpoints, diags := extractCustomConnectorEndpointModels(ctx, prior.Endpoints)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	existingIDByName := make(map[string]int64, len(priorEndpoints))
	for _, e := range priorEndpoints {
		existingIDByName[e.Name.ValueString()] = e.ID.ValueInt64()
	}
	endpointInputs, diags := buildCustomConnectorEndpointInputs(ctx, planEndpoints, existingIDByName)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := parameter.UpdateCustomConnectorInputInput{
		Description: model.NewNullableString(plan.Description),
		Endpoints:   endpointInputs,
	}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		input.SetName(plan.Name.ValueString())
	}
	if !plan.URL.IsNull() && !plan.URL.IsUnknown() {
		input.SetURL(plan.URL.ValueString())
	}
	if !plan.AuthType.IsNull() && !plan.AuthType.IsUnknown() {
		input.SetAuthType(plan.AuthType.ValueString())
	}
	if !plan.AuthHeaderName.IsNull() && !plan.AuthHeaderName.IsUnknown() {
		input.SetAuthHeaderName(plan.AuthHeaderName.ValueString())
	}
	if !plan.AuthHeaderScheme.IsNull() && !plan.AuthHeaderScheme.IsUnknown() {
		input.SetAuthHeaderScheme(plan.AuthHeaderScheme.ValueString())
	}
	if !plan.GrantType.IsNull() && !plan.GrantType.IsUnknown() {
		input.SetGrantType(plan.GrantType.ValueString())
	}
	if !plan.AuthURI.IsNull() && !plan.AuthURI.IsUnknown() {
		input.SetAuthURI(plan.AuthURI.ValueString())
	}
	if !plan.AccessTokenURI.IsNull() && !plan.AccessTokenURI.IsUnknown() {
		input.SetAccessTokenURI(plan.AccessTokenURI.ValueString())
	}

	def, err := r.client.UpdateCustomConnectorInput(prior.ID.ValueInt64(), &input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating custom connector input",
			fmt.Sprintf("Unable to update custom connector input, got error: %s", err),
		)
		return
	}

	state, err := model.NewCustomConnectorInputModel(ctx, def, planEndpoints)
	if err != nil {
		resp.Diagnostics.AddError("Converting custom connector input", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *customConnectorInputResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state model.CustomConnectorInputModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteCustomConnectorInput(state.ID.ValueInt64()); err != nil {
		resp.Diagnostics.AddError(
			"Deleting custom connector input",
			fmt.Sprintf("Unable to delete custom connector input, got error: %s", err),
		)
		return
	}
}

func (r *customConnectorInputResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing custom connector input",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

// extractCustomConnectorEndpointModels converts the plan/state's types.List into
// a Go slice. Null/Unknown (e.g. `endpoints` omitted) is mapped to an empty
// slice so the request payload always carries `endpoints: []`.
func extractCustomConnectorEndpointModels(ctx context.Context, list types.List) ([]model.CustomConnectorEndpointModel, diag.Diagnostics) {
	if list.IsNull() || list.IsUnknown() {
		return []model.CustomConnectorEndpointModel{}, nil
	}
	var endpoints []model.CustomConnectorEndpointModel
	diags := list.ElementsAs(ctx, &endpoints, false)
	if diags.HasError() {
		return nil, diags
	}
	return endpoints, nil
}

// buildCustomConnectorEndpointInputs builds the request payload for `endpoints`.
// existingIDByName maps endpoint name -> server-assigned id, resolved from
// prior state, so that Update preserves identity for endpoints whose name is
// unchanged (see the resource documentation for the naming/matching contract).
func buildCustomConnectorEndpointInputs(ctx context.Context, endpoints []model.CustomConnectorEndpointModel, existingIDByName map[string]int64) ([]parameter.CustomConnectorEndpointInput, diag.Diagnostics) {
	var diags diag.Diagnostics
	inputs := make([]parameter.CustomConnectorEndpointInput, 0, len(endpoints))
	for _, e := range endpoints {
		input := parameter.CustomConnectorEndpointInput{
			Name:              e.Name.ValueString(),
			Path:              e.Path.ValueString(),
			Method:            e.Method.ValueString(),
			JsonpathRoot:      e.JsonpathRoot.ValueString(),
			SuccessCodes:      e.SuccessCodes.ValueString(),
			NotRetryableCodes: e.NotRetryableCodes.ValueString(),
		}
		if id, ok := existingIDByName[e.Name.ValueString()]; ok {
			input.SetID(id)
		}
		if !e.RequestBody.IsNull() && !e.RequestBody.IsUnknown() {
			input.SetRequestBody(e.RequestBody.ValueString())
		}
		if !e.RequestTimeoutSec.IsNull() && !e.RequestTimeoutSec.IsUnknown() {
			input.SetRequestTimeoutSec(e.RequestTimeoutSec.ValueInt64())
		}

		queryParams, qDiags := extractCustomConnectorFieldOptionInputs(ctx, e.QueryParameters)
		diags.Append(qDiags...)
		input.QueryParameters = queryParams

		headers, hDiags := extractCustomConnectorFieldOptionInputs(ctx, e.Headers)
		diags.Append(hDiags...)
		input.Headers = headers

		pathParams, ppDiags := extractCustomConnectorPathParameterInputs(ctx, e.PathParameters)
		diags.Append(ppDiags...)
		input.PathParameters = pathParams

		requestBodyParams, rbDiags := extractCustomConnectorFieldOptionInputs(ctx, e.RequestBodyParameters)
		diags.Append(rbDiags...)
		input.RequestBodyParameters = requestBodyParams

		if e.Paginator != nil {
			input.SetPaginator(buildCustomConnectorPaginatorInput(e.Paginator))
		}

		inputs = append(inputs, input)
	}
	return inputs, diags
}

func extractCustomConnectorFieldOptionInputs(ctx context.Context, list types.List) ([]parameter.CustomConnectorFieldOptionInput, diag.Diagnostics) {
	if list.IsNull() || list.IsUnknown() {
		return []parameter.CustomConnectorFieldOptionInput{}, nil
	}
	var options []model.CustomConnectorFieldOptionModel
	diags := list.ElementsAs(ctx, &options, false)
	if diags.HasError() {
		return nil, diags
	}
	out := make([]parameter.CustomConnectorFieldOptionInput, 0, len(options))
	for _, o := range options {
		input := parameter.CustomConnectorFieldOptionInput{
			Name:        o.Name.ValueString(),
			DisplayName: o.DisplayName.ValueString(),
			IsEditable:  o.IsEditable.ValueBool(),
			IsRequired:  o.IsRequired.ValueBool(),
		}
		if !o.DefaultValue.IsNull() && !o.DefaultValue.IsUnknown() {
			v := o.DefaultValue.ValueString()
			input.DefaultValue = &v
		}
		out = append(out, input)
	}
	return out, nil
}

func extractCustomConnectorPathParameterInputs(ctx context.Context, list types.List) ([]parameter.CustomConnectorPathParameterInput, diag.Diagnostics) {
	if list.IsNull() || list.IsUnknown() {
		return []parameter.CustomConnectorPathParameterInput{}, nil
	}
	var params []model.CustomConnectorPathParameterModel
	diags := list.ElementsAs(ctx, &params, false)
	if diags.HasError() {
		return nil, diags
	}
	out := make([]parameter.CustomConnectorPathParameterInput, 0, len(params))
	for _, p := range params {
		out = append(out, parameter.CustomConnectorPathParameterInput{
			Name:  p.Name.ValueString(),
			Value: p.Value.ValueString(),
		})
	}
	return out, nil
}

func buildCustomConnectorPaginatorInput(p *model.CustomConnectorPaginatorModel) parameter.CustomConnectorPaginatorInput {
	input := parameter.CustomConnectorPaginatorInput{}
	if !p.InjectInto.IsNull() && !p.InjectInto.IsUnknown() {
		v := p.InjectInto.ValueString()
		input.InjectInto = &v
	}

	switch {
	case p.PageIncrementStrategy != nil:
		s := p.PageIncrementStrategy
		strategy := parameter.CustomConnectorPaginationStrategyInput{Type: "page_increment"}
		if !s.PageSize.IsNull() && !s.PageSize.IsUnknown() {
			v := s.PageSize.ValueInt64()
			strategy.PageSize = &v
		}
		if !s.LastPageSize.IsNull() && !s.LastPageSize.IsUnknown() {
			v := s.LastPageSize.ValueString()
			strategy.LastPageSize = &v
		}
		if !s.StartFromPage.IsNull() && !s.StartFromPage.IsUnknown() {
			v := s.StartFromPage.ValueInt64()
			strategy.StartFromPage = &v
		}
		if !s.StopOnPage.IsNull() && !s.StopOnPage.IsUnknown() {
			v := s.StopOnPage.ValueInt64()
			strategy.StopOnPage = &v
		}
		if !s.TotalPages.IsNull() && !s.TotalPages.IsUnknown() {
			v := s.TotalPages.ValueString()
			strategy.TotalPages = &v
		}
		if !s.MaxRequestCount.IsNull() && !s.MaxRequestCount.IsUnknown() {
			v := s.MaxRequestCount.ValueInt64()
			strategy.MaxRequestCount = &v
		}
		input.PaginationStrategy = strategy
	case p.OffsetIncrementStrategy != nil:
		s := p.OffsetIncrementStrategy
		strategy := parameter.CustomConnectorPaginationStrategyInput{Type: "offset_increment"}
		if !s.PageSize.IsNull() && !s.PageSize.IsUnknown() {
			v := s.PageSize.ValueInt64()
			strategy.PageSize = &v
		}
		if !s.LastPageSize.IsNull() && !s.LastPageSize.IsUnknown() {
			v := s.LastPageSize.ValueString()
			strategy.LastPageSize = &v
		}
		if !s.StartFromOffset.IsNull() && !s.StartFromOffset.IsUnknown() {
			v := s.StartFromOffset.ValueInt64()
			strategy.StartFromOffset = &v
		}
		if !s.StopOnOffset.IsNull() && !s.StopOnOffset.IsUnknown() {
			v := s.StopOnOffset.ValueInt64()
			strategy.StopOnOffset = &v
		}
		if !s.TotalRecords.IsNull() && !s.TotalRecords.IsUnknown() {
			v := s.TotalRecords.ValueString()
			strategy.TotalRecords = &v
		}
		if !s.MaxRequestCount.IsNull() && !s.MaxRequestCount.IsUnknown() {
			v := s.MaxRequestCount.ValueInt64()
			strategy.MaxRequestCount = &v
		}
		input.PaginationStrategy = strategy
	case p.CursorBasedStrategy != nil:
		s := p.CursorBasedStrategy
		strategy := parameter.CustomConnectorPaginationStrategyInput{Type: "cursor_based"}
		if !s.PageSize.IsNull() && !s.PageSize.IsUnknown() {
			v := s.PageSize.ValueInt64()
			strategy.PageSize = &v
		}
		if !s.CursorValue.IsNull() && !s.CursorValue.IsUnknown() {
			v := s.CursorValue.ValueString()
			strategy.CursorValue = &v
		}
		input.PaginationStrategy = strategy
	}

	if p.PageSizeOption != nil {
		input.PageSizeOption = &parameter.CustomConnectorFieldNameOptionInput{FieldName: p.PageSizeOption.FieldName.ValueString()}
	}
	if p.PageTokenOption != nil {
		input.PageTokenOption = &parameter.CustomConnectorFieldNameOptionInput{FieldName: p.PageTokenOption.FieldName.ValueString()}
	}

	return input
}
