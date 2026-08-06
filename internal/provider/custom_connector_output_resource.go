package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/client/parameter"
	"terraform-provider-trocco/internal/provider/model"
	troccoPlanModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ resource.Resource                   = &customConnectorOutputResource{}
	_ resource.ResourceWithConfigure      = &customConnectorOutputResource{}
	_ resource.ResourceWithImportState    = &customConnectorOutputResource{}
	_ resource.ResourceWithValidateConfig = &customConnectorOutputResource{}
)

// Defaults mirroring the server-side column defaults. The server assigns every
// endpoint attribute unconditionally, so the request must always carry a value:
// a missing key is written as NULL and rejected by the presence validations
// instead of falling back to the column default.
const (
	customConnectorOutputDefaultRequestType       = "single"
	customConnectorOutputDefaultBatchSize         = 1
	customConnectorOutputDefaultMethod            = "POST"
	customConnectorOutputDefaultOperation         = "create"
	customConnectorOutputDefaultPayloadType       = "json"
	customConnectorOutputDefaultSuccessCodes      = "200,201,204"
	customConnectorOutputDefaultNotRetryableCodes = "400,401,403,404"
	customConnectorOutputDefaultRequestTimeoutSec = 30
)

// customConnectorOutputPathSyntaxHint supplements the shared path parameter
// validation: an output endpoint's `path` is rendered as a Liquid template, and
// the server matches the placeholder as a plain substring, so the Liquid
// braces must not contain surrounding whitespace.
const customConnectorOutputPathSyntaxHint = "The endpoint path is a Liquid template, so write the placeholder without surrounding whitespace, " +
	"e.g. path = \"/users/{{user_id}}\" rather than \"/users/{{ user_id }}\"."

func NewCustomConnectorOutputResource() resource.Resource {
	return &customConnectorOutputResource{}
}

type customConnectorOutputResource struct {
	client *client.TroccoClient
}

func (r *customConnectorOutputResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_connector_output"
}

func (r *customConnectorOutputResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *customConnectorOutputResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO custom connector output (transfer destination) definition resource.\n\n" +
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
						// The server has no presence validation on an output
						// endpoint's `name` (unlike the transfer source), so an
						// empty name would be persisted silently. The remaining
						// non-empty checks below guard against the server's
						// presence errors surfacing as an opaque 400 ("Custom
						// connector output endpoints is invalid") via
						// validates_associated, which names no attribute.
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
							MarkdownDescription: "Endpoint name. Must be unique within the connector; used to match endpoints across applies (renaming an endpoint is treated as delete+create).",
						},
						"request_type": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString(customConnectorOutputDefaultRequestType),
							Validators: []validator.String{
								stringvalidator.OneOf("single", "multiple"),
							},
							MarkdownDescription: "Whether records are sent one at a time (`single`) or in batches (`multiple`). Defaults to `single`.",
						},
						"batch_size": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							Default:  int64default.StaticInt64(customConnectorOutputDefaultBatchSize),
							Validators: []validator.Int64{
								int64validator.Between(1, 100),
							},
							MarkdownDescription: "Number of records per request (1-100). Must be `1` when `request_type` is `single`. Defaults to `1`.",
						},
						"method": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString(customConnectorOutputDefaultMethod),
							Validators: []validator.String{
								stringvalidator.OneOf("GET", "POST", "PUT", "PATCH"),
							},
							MarkdownDescription: "HTTP method. Defaults to `POST`.",
						},
						"operation": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString(customConnectorOutputDefaultOperation),
							Validators: []validator.String{
								stringvalidator.OneOf("create", "update"),
							},
							MarkdownDescription: "Whether the endpoint creates (`create`) or updates (`update`) records. A job definition references one `create` endpoint and, in upsert mode, one `update` endpoint. Defaults to `create`.",
						},
						"path": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
							MarkdownDescription: "Endpoint path. Rendered as a Liquid template when `request_type` is `single`.",
						},
						"payload_type": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString(customConnectorOutputDefaultPayloadType),
							Validators: []validator.String{
								stringvalidator.OneOf("json"),
							},
							MarkdownDescription: "Request payload format. Only `json` is supported. Defaults to `json`.",
						},
						"success_codes": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString(customConnectorOutputDefaultSuccessCodes),
							Validators: []validator.String{
								stringvalidator.RegexMatches(commaSeparatedCodesRegex, "must be a comma-separated list of numbers"),
							},
							MarkdownDescription: "Comma-separated HTTP status codes treated as success. Defaults to `200,201,204`.",
						},
						"not_retryable_codes": schema.StringAttribute{
							Optional: true,
							Computed: true,
							Default:  stringdefault.StaticString(customConnectorOutputDefaultNotRetryableCodes),
							Validators: []validator.String{
								stringvalidator.RegexMatches(commaSeparatedCodesRegex, "must be a comma-separated list of numbers"),
							},
							MarkdownDescription: "Comma-separated HTTP status codes that must not be retried. Defaults to `400,401,403,404`.",
						},
						"request_timeout_sec": schema.Int64Attribute{
							Optional: true,
							Computed: true,
							Default:  int64default.StaticInt64(customConnectorOutputDefaultRequestTimeoutSec),
							Validators: []validator.Int64{
								int64validator.Between(1, 60),
							},
							MarkdownDescription: "Seconds to wait for a response before timing out (1-60). Defaults to 30.",
						},
						"template": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.LengthAtLeast(1),
							},
							MarkdownDescription: "Liquid template used to build the request body. Must be valid Liquid syntax; the server is the authority.",
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
										MarkdownDescription: "Value. Must appear as `{value}` in the endpoint's `path`, so write the Liquid placeholder without surrounding whitespace (`/users/{{user_id}}`, not `/users/{{ user_id }}`).",
									},
								},
							},
							MarkdownDescription: "Path parameter definitions. Set to `[]` to clear.",
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
					},
				},
				MarkdownDescription: "Ordered list of endpoint definitions. Set to `[]` to clear. " +
					"Multiple endpoints may share the same `operation`; the server does not enforce uniqueness, " +
					"and a job definition selects the endpoints it uses by ID.",
			},
		},
	}
}

func (r *customConnectorOutputResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var plan model.CustomConnectorOutputModel
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

	endpoints, diags := extractCustomConnectorOutputEndpointModels(ctx, plan.Endpoints)
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

		// `request_type` defaults to `single`, for which the server requires
		// `batch_size` to be exactly 1.
		singleRequest := e.RequestType.IsNull() || (!e.RequestType.IsUnknown() && e.RequestType.ValueString() == customConnectorOutputDefaultRequestType)
		if singleRequest && !e.BatchSize.IsNull() && !e.BatchSize.IsUnknown() && e.BatchSize.ValueInt64() != customConnectorOutputDefaultBatchSize {
			resp.Diagnostics.AddAttributeError(
				endpointPath.AtName("batch_size"),
				"Invalid Batch Size",
				"`batch_size` must be 1 when `request_type` is \"single\".",
			)
		}

		resp.Diagnostics.Append(validateCustomConnectorPathParameters(ctx, endpointPath, e.Path, e.PathParameters, customConnectorOutputPathSyntaxHint)...)
	}
}

func (r *customConnectorOutputResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan model.CustomConnectorOutputModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoints, diags := extractCustomConnectorOutputEndpointModels(ctx, plan.Endpoints)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	endpointInputs, diags := buildCustomConnectorOutputEndpointInputs(ctx, endpoints, nil)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := parameter.CreateCustomConnectorOutputInput{
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

	def, err := r.client.CreateCustomConnectorOutput(&input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating custom connector output",
			fmt.Sprintf("Unable to create custom connector output, got error: %s", err),
		)
		return
	}

	state, err := model.NewCustomConnectorOutputModel(ctx, def, endpoints)
	if err != nil {
		resp.Diagnostics.AddError("Converting custom connector output", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *customConnectorOutputResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var prior model.CustomConnectorOutputModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	if resp.Diagnostics.HasError() {
		return
	}

	def, err := r.client.GetCustomConnectorOutput(prior.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading custom connector output",
			fmt.Sprintf("Unable to read custom connector output, got error: %s", err),
		)
		return
	}

	refEndpoints, diags := extractCustomConnectorOutputEndpointModels(ctx, prior.Endpoints)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state, err := model.NewCustomConnectorOutputModel(ctx, def, refEndpoints)
	if err != nil {
		resp.Diagnostics.AddError("Converting custom connector output", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *customConnectorOutputResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, prior model.CustomConnectorOutputModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	if resp.Diagnostics.HasError() {
		return
	}

	planEndpoints, diags := extractCustomConnectorOutputEndpointModels(ctx, plan.Endpoints)
	resp.Diagnostics.Append(diags...)
	priorEndpoints, diags := extractCustomConnectorOutputEndpointModels(ctx, prior.Endpoints)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	existingIDByName := make(map[string]int64, len(priorEndpoints))
	for _, e := range priorEndpoints {
		existingIDByName[e.Name.ValueString()] = e.ID.ValueInt64()
	}
	endpointInputs, diags := buildCustomConnectorOutputEndpointInputs(ctx, planEndpoints, existingIDByName)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := parameter.UpdateCustomConnectorOutputInput{
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

	def, err := r.client.UpdateCustomConnectorOutput(prior.ID.ValueInt64(), &input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating custom connector output",
			fmt.Sprintf("Unable to update custom connector output, got error: %s", err),
		)
		return
	}

	state, err := model.NewCustomConnectorOutputModel(ctx, def, planEndpoints)
	if err != nil {
		resp.Diagnostics.AddError("Converting custom connector output", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *customConnectorOutputResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state model.CustomConnectorOutputModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteCustomConnectorOutput(state.ID.ValueInt64()); err != nil {
		resp.Diagnostics.AddError(
			"Deleting custom connector output",
			fmt.Sprintf("Unable to delete custom connector output, got error: %s", err),
		)
		return
	}
}

func (r *customConnectorOutputResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing custom connector output",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

// extractCustomConnectorOutputEndpointModels converts the plan/state's
// types.List into a Go slice. Null/Unknown (e.g. `endpoints` omitted) is mapped
// to an empty slice so the request payload always carries `endpoints: []`.
func extractCustomConnectorOutputEndpointModels(ctx context.Context, list types.List) ([]model.CustomConnectorOutputEndpointModel, diag.Diagnostics) {
	if list.IsNull() || list.IsUnknown() {
		return []model.CustomConnectorOutputEndpointModel{}, nil
	}
	var endpoints []model.CustomConnectorOutputEndpointModel
	diags := list.ElementsAs(ctx, &endpoints, false)
	if diags.HasError() {
		return nil, diags
	}
	return endpoints, nil
}

// buildCustomConnectorOutputEndpointInputs builds the request payload for
// `endpoints`. existingIDByName maps endpoint name -> server-assigned id,
// resolved from prior state, so that Update preserves identity for endpoints
// whose name is unchanged (see the resource documentation for the
// naming/matching contract).
func buildCustomConnectorOutputEndpointInputs(ctx context.Context, endpoints []model.CustomConnectorOutputEndpointModel, existingIDByName map[string]int64) ([]parameter.CustomConnectorOutputEndpointInput, diag.Diagnostics) {
	var diags diag.Diagnostics
	inputs := make([]parameter.CustomConnectorOutputEndpointInput, 0, len(endpoints))
	for _, e := range endpoints {
		input := parameter.CustomConnectorOutputEndpointInput{
			Name:              e.Name.ValueString(),
			RequestType:       e.RequestType.ValueString(),
			BatchSize:         e.BatchSize.ValueInt64(),
			Method:            e.Method.ValueString(),
			Operation:         e.Operation.ValueString(),
			Path:              e.Path.ValueString(),
			PayloadType:       e.PayloadType.ValueString(),
			SuccessCodes:      e.SuccessCodes.ValueString(),
			NotRetryableCodes: e.NotRetryableCodes.ValueString(),
			RequestTimeoutSec: e.RequestTimeoutSec.ValueInt64(),
			Template:          e.Template.ValueString(),
		}
		if id, ok := existingIDByName[e.Name.ValueString()]; ok {
			input.SetID(id)
		}

		headers, hDiags := extractCustomConnectorFieldOptionInputs(ctx, e.Headers)
		diags.Append(hDiags...)
		input.Headers = headers

		pathParams, ppDiags := extractCustomConnectorPathParameterInputs(ctx, e.PathParameters)
		diags.Append(ppDiags...)
		input.PathParameters = pathParams

		queryParams, qDiags := extractCustomConnectorFieldOptionInputs(ctx, e.QueryParameters)
		diags.Append(qDiags...)
		input.QueryParameters = queryParams

		inputs = append(inputs, input)
	}
	return inputs, diags
}
