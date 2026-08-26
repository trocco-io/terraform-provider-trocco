package job_definition

import (
	"sort"
	troccoPlanModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// customConnectorComputedString/Int64/Bool build read-only (server-derived)
// leaf attributes. They use CustomConnectorSnapshot*PlanModifier (not the
// framework's built-in UseStateForUnknown) because these snapshot fields can
// legitimately be null (e.g. an endpoint with no request body), and
// UseStateForUnknown leaves the plan Unknown whenever the prior state value
// is null - which would make the attribute show as "(known after apply)" on
// every plan even though nothing changed.
func customConnectorComputedString(description string) schema.Attribute {
	return schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: description,
		PlanModifiers: []planmodifier.String{
			troccoPlanModifier.CustomConnectorSnapshotStringPlanModifier{},
		},
	}
}

func customConnectorComputedInt64(description string) schema.Attribute {
	return schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: description,
		PlanModifiers: []planmodifier.Int64{
			troccoPlanModifier.CustomConnectorSnapshotInt64PlanModifier{},
		},
	}
}

func customConnectorComputedBool(description string) schema.Attribute {
	return schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: description,
		PlanModifiers: []planmodifier.Bool{
			troccoPlanModifier.CustomConnectorSnapshotBoolPlanModifier{},
		},
	}
}

func CustomConnectorInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of a source that uses a custom connector definition (`trocco_custom_connector_input`). " +
			"The attributes documented as a snapshot are read-only copies of the referenced endpoint, and TROCCO rebuilds all of them " +
			"every time this job definition is updated. Any update therefore plans them as `(known after apply)`, and changing the " +
			"endpoint definition - whether through `trocco_custom_connector_input` or the web UI - is reflected here on the next update " +
			"of this job definition, not when the definition itself changes.",
		Attributes: map[string]schema.Attribute{
			"custom_connector_endpoint_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of the custom connector endpoint (`trocco_custom_connector_input.<name>.endpoints[N].id`) to use as the transfer source",
			},
			"custom_connector_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of the custom connector connection (`trocco_connection` with `connection_type = \"custom_connector\"`)",
			},
			"query_parameters": customConnectorNamedValueListSchema(
				"Query parameter values. Only parameters marked `is_editable = true` on the referenced endpoint can be specified, " +
					"and a value must be supplied for every parameter marked `is_required = true`. " +
					"Omitting this attribute keeps the existing values; specifying `[]` clears them; specifying values fully replaces them.",
			),
			"headers": customConnectorNamedValueListSchema(
				"Request header values. Only headers marked `is_editable = true` on the referenced endpoint can be specified, " +
					"and a value must be supplied for every header marked `is_required = true`. " +
					"Omitting this attribute keeps the existing values; specifying `[]` clears them; specifying values fully replaces them.",
			),
			"path_parameters": customConnectorNamedValueListSchema(
				"Path parameter values. A value must be specified for every placeholder in the referenced endpoint's path. " +
					"Omitting this attribute keeps the existing values; specifying `[]` clears them; specifying values fully replaces them.",
			),
			"request_body_parameters": customConnectorNamedValueListSchema(
				"Request body parameter values, substituted into the `{name}` placeholders of the referenced endpoint's request body. " +
					"Only parameters marked `is_editable = true` on the referenced endpoint can be specified, " +
					"and a value must be supplied for every parameter marked `is_required = true`. " +
					"Omitting this attribute keeps the existing values; specifying `[]` clears them; specifying values fully replaces them.",
			),
			"jsonpath_parser":          customConnectorJsonpathParserSchema(),
			"custom_variable_settings": CustomVariableSettingsSchema(),
			"url":                      customConnectorComputedString("URL (snapshot from the custom connector definition)"),
			"auth_type":                customConnectorComputedString("Authentication method (snapshot from the custom connector definition)"),
			"auth_header_name":         customConnectorComputedString("Authentication header name (snapshot from the custom connector definition)"),
			"auth_header_scheme":       customConnectorComputedString("Authentication header scheme (snapshot from the custom connector definition)"),
			"endpoint_path":            customConnectorComputedString("Endpoint path (snapshot from the custom connector definition)"),
			"endpoint_method":          customConnectorComputedString("HTTP method (snapshot from the custom connector definition)"),
			"endpoint_request_body":    customConnectorComputedString("Request body (snapshot from the custom connector definition)"),
			"success_codes":            customConnectorComputedString("Status codes treated as success (snapshot from the custom connector definition)"),
			"not_retryable_codes":      customConnectorComputedString("Status codes excluded from retries (snapshot from the custom connector definition)"),
			"request_timeout_sec":      customConnectorComputedInt64("Seconds to wait for a response before timing out (snapshot from the custom connector definition). `null` for job definitions created before this field existed; treated as 30 in that case."),
			"paginator":                customConnectorPaginatorSchema(),
		},
	}
}

// CustomConnectorSnapshotAttribute identifies an attribute of
// custom_connector_input_option whose value the API derives from the custom
// connector definition rather than from the configuration.
type CustomConnectorSnapshotAttribute struct {
	Path path.Path
	Type attr.Type
}

// CustomConnectorSnapshotAttributes reports which attributes of the
// custom_connector_input_option object rooted at base are server-derived for
// the given configuration, so a plan can mark them unknown before an update.
//
// The set is derived from the schema rather than hard-coded: the API rebuilds
// the whole snapshot on every upsert, so any attribute the configuration does
// not own is subject to change, and a new snapshot attribute is picked up here
// without further changes. An attribute qualifies when it is Computed, carries
// no default (a default is a provider-side value, not a server-derived one)
// and is absent from the configuration. The last condition matters for the
// Optional + Computed named value lists: the API only reconciles them against
// the endpoint definition while they are omitted, and once they are configured
// Terraform requires the planned value to match the configuration anyway.
func CustomConnectorSnapshotAttributes(base path.Path, config types.Object) []CustomConnectorSnapshotAttribute {
	root, ok := CustomConnectorInputOptionSchema().(schema.SingleNestedAttribute)
	if !ok {
		return nil
	}

	attributes := customConnectorSnapshotAttributes(base, root.Attributes, config)
	sort.Slice(attributes, func(i, j int) bool {
		return attributes[i].Path.String() < attributes[j].Path.String()
	})
	return attributes
}

func customConnectorSnapshotAttributes(
	base path.Path,
	attributes map[string]schema.Attribute,
	config types.Object,
) []CustomConnectorSnapshotAttribute {
	if config.IsNull() || config.IsUnknown() {
		return nil
	}

	configAttributes := config.Attributes()

	var snapshotAttributes []CustomConnectorSnapshotAttribute
	for name, attribute := range attributes {
		attributePath := base.AtName(name)
		configValue := configAttributes[name]

		if attribute.IsComputed() && !customConnectorAttributeHasDefault(attribute) {
			if configValue == nil || configValue.IsNull() {
				snapshotAttributes = append(snapshotAttributes, CustomConnectorSnapshotAttribute{
					Path: attributePath,
					Type: attribute.GetType(),
				})
			}
			continue
		}

		// Descend into configurable nested objects (jsonpath_parser) to reach
		// the server-derived attributes they contain.
		nested, ok := attribute.(schema.SingleNestedAttribute)
		if !ok {
			continue
		}
		nestedConfig, ok := configValue.(types.Object)
		if !ok {
			continue
		}
		snapshotAttributes = append(
			snapshotAttributes,
			customConnectorSnapshotAttributes(attributePath, nested.Attributes, nestedConfig)...,
		)
	}
	return snapshotAttributes
}

// customConnectorAttributeHasDefault covers the attribute types used by
// CustomConnectorInputOptionSchema. An unlisted type simply reports no
// default, which at worst plans an attribute as unknown unnecessarily.
func customConnectorAttributeHasDefault(attribute schema.Attribute) bool {
	switch a := attribute.(type) {
	case schema.StringAttribute:
		return a.Default != nil
	case schema.Int64Attribute:
		return a.Default != nil
	case schema.BoolAttribute:
		return a.Default != nil
	case schema.ListAttribute:
		return a.Default != nil
	case schema.ListNestedAttribute:
		return a.Default != nil
	case schema.SingleNestedAttribute:
		return a.Default != nil
	default:
		return false
	}
}

func customConnectorNamedValueListSchema(description string) schema.Attribute {
	return schema.ListNestedAttribute{
		Optional: true,
		Computed: true,
		MarkdownDescription: description +
			" When this attribute is omitted and the snapshot is rebuilt, values whose name no longer exists on the referenced " +
			"endpoint are dropped, so the result of an update is only known after apply.",
		PlanModifiers: []planmodifier.List{
			troccoPlanModifier.CustomConnectorSnapshotListPlanModifier{},
		},
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Name",
					Validators: []validator.String{
						stringvalidator.UTF8LengthAtLeast(1),
					},
				},
				"value": schema.StringAttribute{
					Required:            true,
					MarkdownDescription: "Value",
				},
			},
		},
	}
}

func customConnectorJsonpathParserSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Required:            true,
		MarkdownDescription: "JSONPath parser settings",
		Attributes: map[string]schema.Attribute{
			"root": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "JSONPath root. Always derived from the referenced endpoint; cannot be set in configuration.",
				PlanModifiers: []planmodifier.String{
					troccoPlanModifier.CustomConnectorSnapshotStringPlanModifier{},
				},
			},
			"default_time_zone": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("UTC"),
				MarkdownDescription: "Default time zone",
			},
			"columns": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "List of columns to be retrieved and their types",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column name",
						},
						"type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean", "json"),
							},
							MarkdownDescription: "Column type",
						},
						"time_zone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Time zone",
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Format of the column",
						},
					},
				},
			},
		},
	}
}

func customConnectorFieldNameOptionSchema(description string) schema.Attribute {
	return schema.SingleNestedAttribute{
		Computed:            true,
		MarkdownDescription: description,
		PlanModifiers: []planmodifier.Object{
			troccoPlanModifier.CustomConnectorSnapshotObjectPlanModifier{},
		},
		Attributes: map[string]schema.Attribute{
			"field_name": customConnectorComputedString("Field name"),
		},
	}
}

// customConnectorPaginatorSchema mirrors the referenced endpoint's pagination
// settings. It intentionally omits the `page_end_option` the API also returns;
// see entity.CustomConnectorPaginator for why.
func customConnectorPaginatorSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Computed:            true,
		MarkdownDescription: "Pagination settings (snapshot from the custom connector definition). `null` if the endpoint has no pagination configured.",
		PlanModifiers: []planmodifier.Object{
			troccoPlanModifier.CustomConnectorSnapshotObjectPlanModifier{},
		},
		Attributes: map[string]schema.Attribute{
			"strategy_type":     customConnectorComputedString("Pagination strategy (`page_increment`, `offset_increment`, or `cursor_based`)"),
			"inject_into":       customConnectorComputedString("Where the paging parameter is injected (`query` or `request_body`)"),
			"page_size_option":  customConnectorFieldNameOptionSchema("Field used to pass the page size"),
			"page_token_option": customConnectorFieldNameOptionSchema("Field used to pass the page token (page number/offset/cursor)"),
			"page_increment_strategy": schema.SingleNestedAttribute{
				Computed:            true,
				MarkdownDescription: "Settings used when `strategy_type` is `page_increment`",
				PlanModifiers: []planmodifier.Object{
					troccoPlanModifier.CustomConnectorSnapshotObjectPlanModifier{},
				},
				Attributes: map[string]schema.Attribute{
					"page_size":               customConnectorComputedInt64("Page size"),
					"start_from_page":         customConnectorComputedInt64("Starting page number"),
					"first_page":              customConnectorComputedInt64("First page number"),
					"inject_on_first_request": customConnectorComputedBool("Whether to inject the paging parameter on the first request"),
					"last_page_size":          customConnectorComputedString("Size of the last page (a literal number or a JSONPath expression)"),
					"total_pages":             customConnectorComputedString("Total number of pages (a literal number or a JSONPath expression)"),
					"stop_on_page":            customConnectorComputedInt64("Page number to stop at"),
					"max_request_count":       customConnectorComputedInt64("Maximum number of requests (used with `last_page_size` to avoid infinite loops)"),
				},
			},
			"offset_increment_strategy": schema.SingleNestedAttribute{
				Computed:            true,
				MarkdownDescription: "Settings used when `strategy_type` is `offset_increment`",
				PlanModifiers: []planmodifier.Object{
					troccoPlanModifier.CustomConnectorSnapshotObjectPlanModifier{},
				},
				Attributes: map[string]schema.Attribute{
					"page_size":               customConnectorComputedInt64("Page size"),
					"start_from_offset":       customConnectorComputedInt64("Starting offset"),
					"first_offset":            customConnectorComputedInt64("First offset"),
					"inject_on_first_request": customConnectorComputedBool("Whether to inject the paging parameter on the first request"),
					"last_page_size":          customConnectorComputedString("Size of the last page (a literal number or a JSONPath expression)"),
					"total_records":           customConnectorComputedString("Total number of records (a literal number or a JSONPath expression)"),
					"stop_on_offset":          customConnectorComputedInt64("Offset to stop at"),
					"max_request_count":       customConnectorComputedInt64("Maximum number of requests (used with `last_page_size` to avoid infinite loops)"),
				},
			},
			"cursor_based_strategy": schema.SingleNestedAttribute{
				Computed:            true,
				MarkdownDescription: "Settings used when `strategy_type` is `cursor_based`",
				PlanModifiers: []planmodifier.Object{
					troccoPlanModifier.CustomConnectorSnapshotObjectPlanModifier{},
				},
				Attributes: map[string]schema.Attribute{
					"page_size":      customConnectorComputedInt64("Page size"),
					"cursor_value":   customConnectorComputedString("Cursor value (a literal value or a JSONPath expression)"),
					"last_page_size": customConnectorComputedString("Size of the last page (a literal number or a JSONPath expression)"),
					"stop_on_blank":  customConnectorComputedString("Whether to stop when the cursor is blank"),
				},
			},
		},
	}
}
