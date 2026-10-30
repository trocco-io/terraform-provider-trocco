package job_definition

import (
	troccoPlanModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func CustomConnectorOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of a destination that uses a custom connector definition (`trocco_custom_connector_output`)",
		PlanModifiers: []planmodifier.Object{
			&troccoPlanModifier.CustomConnectorOutputOptionPlanModifier{},
		},
		Attributes: map[string]schema.Attribute{
			"custom_connector_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of the custom connector connection (`trocco_connection` with `connection_type = \"custom_connector\"`). It must belong to the same custom connector as the endpoints below.",
			},
			"mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("insert"),
				Validators: []validator.String{
					stringvalidator.OneOf("insert", "upsert"),
				},
				MarkdownDescription: "Transfer mode. `insert` sends every record to the create endpoint; `upsert` additionally requires `update_key` and `update_custom_connector_output_endpoint_id`. Default is `insert`.",
			},
			"update_key": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Column used to decide whether a record already exists. Required when `mode` is `upsert`, and must not be specified when `mode` is `insert`.",
			},
			"create_custom_connector_output_endpoint_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of the custom connector endpoint (`trocco_custom_connector_output.<name>.endpoints[N].id`) used to send new records. Its `operation` must be `create`.",
			},
			"update_custom_connector_output_endpoint_id": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of the custom connector endpoint (`trocco_custom_connector_output.<name>.endpoints[N].id`) used to update existing records. Its `operation` must be `update`, and it must belong to the same custom connector as the create endpoint. Required when `mode` is `upsert`, and must not be specified when `mode` is `insert`.",
			},
			"create_endpoint_settings": customConnectorEndpointSettingsSchema(
				"Values submitted to the create endpoint.",
			),
			"update_endpoint_settings": customConnectorEndpointSettingsSchema(
				"Values submitted to the update endpoint. Only used when `mode` is `upsert`.",
			),
			"custom_variable_settings": CustomVariableSettingsSchema(),
			"url":                      customConnectorOutputComputedString("URL (snapshot from the custom connector definition)"),
			"auth_type":                customConnectorOutputComputedString("Authentication method (snapshot from the custom connector definition)"),
			"auth_header_name":         customConnectorOutputComputedString("Authentication header name (snapshot from the custom connector definition)"),
			"auth_header_scheme":       customConnectorOutputComputedString("Authentication header scheme (snapshot from the custom connector definition)"),
			"endpoints":                customConnectorOutputEndpointsSchema(),
		},
	}
}

// customConnectorOutputComputedString builds a read-only (server-derived) leaf
// attribute. See CustomConnectorOutputSnapshotStringPlanModifier for why the
// framework's built-in UseStateForUnknown is not used.
func customConnectorOutputComputedString(description string) schema.Attribute {
	return schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: description,
		PlanModifiers: []planmodifier.String{
			troccoPlanModifier.CustomConnectorOutputSnapshotStringPlanModifier{},
		},
	}
}

// customConnectorEndpointSettingsSchema builds create_endpoint_settings /
// update_endpoint_settings. These are configuration-only: the API accepts them
// but reports the resulting values through `endpoints` instead, merged with the
// definition's non-editable defaults. They are therefore not Computed, and the
// provider carries the configured value through create/update/read unchanged.
func customConnectorEndpointSettingsSchema(description string) schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		MarkdownDescription: description +
			" Only headers and query parameters marked `is_editable = true` on the referenced endpoint can be specified, " +
			"and a value must be supplied for every one of them marked `is_required = true`. " +
			"Path parameters may be left unset. " +
			"Omitting the whole attribute, or one of its collections, keeps the values stored on the server.\n\n" +
			"This attribute is configuration-only: the API accepts it but never reports it back, and the resulting " +
			"values are exposed through `endpoints` instead. Two consequences follow. It cannot be recovered on " +
			"`terraform import`, so an imported job definition shows it as unset until the configuration supplies it " +
			"again. And removing it from an existing configuration only clears it from the Terraform state: the values " +
			"already stored on the server are kept, since omitting the attribute is how \"keep the current values\" is " +
			"expressed. To actually clear a collection, set it to `[]` rather than removing it.",
		Attributes: map[string]schema.Attribute{
			"headers": customConnectorOutputNamedValueListSchema(
				"Request header values. Omitting this attribute keeps the existing values; specifying `[]` clears them; specifying values fully replaces them.",
			),
			"path_parameters": customConnectorOutputNamedValueListSchema(
				"Path parameter values, substituted into the placeholders of the referenced endpoint's path. Omitting this attribute keeps the existing values; specifying `[]` clears them; specifying values fully replaces them.",
			),
			"query_parameters": customConnectorOutputNamedValueListSchema(
				"Query parameter values. Omitting this attribute keeps the existing values; specifying `[]` clears them; specifying values fully replaces them.",
			),
		},
	}
}

func customConnectorOutputNamedValueListSchema(description string) schema.Attribute {
	return schema.ListNestedAttribute{
		Optional:            true,
		MarkdownDescription: description,
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

func customConnectorOutputEndpointsSchema() schema.Attribute {
	return schema.ListNestedAttribute{
		Computed: true,
		MarkdownDescription: "Snapshots of the endpoints this job definition uses, taken from the custom connector definition (in creation order). " +
			"Each entry is distinguished by `operation`; the `update` entry is absent while `mode` is `insert`.",
		PlanModifiers: []planmodifier.List{
			troccoPlanModifier.CustomConnectorOutputSnapshotListPlanModifier{},
		},
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"operation": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "Whether this endpoint sends new records (`create`) or updates existing ones (`update`)",
				},
				"request_type": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "Whether records are sent one per request (`single`) or in batches (`multiple`)",
				},
				"batch_size": schema.Int64Attribute{
					Computed:            true,
					MarkdownDescription: "Number of records sent per request when `request_type` is `multiple`",
				},
				"method": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "HTTP method",
				},
				"path": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "Endpoint path",
				},
				"payload_type": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "Request payload format",
				},
				"success_codes": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "Status codes treated as success",
				},
				"not_retryable_codes": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "Status codes excluded from retries",
				},
				"request_timeout_sec": schema.Int64Attribute{
					Computed:            true,
					MarkdownDescription: "Seconds to wait for a response before timing out",
				},
				"template": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "Request body template",
				},
				"headers":          customConnectorOutputComputedNamedValueListSchema("Request header values"),
				"path_parameters":  customConnectorOutputComputedNamedValueListSchema("Path parameter values"),
				"query_parameters": customConnectorOutputComputedNamedValueListSchema("Query parameter values"),
			},
		},
	}
}

func customConnectorOutputComputedNamedValueListSchema(description string) schema.Attribute {
	return schema.ListNestedAttribute{
		Computed:            true,
		MarkdownDescription: description,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "Name",
				},
				"value": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "Value",
				},
			},
		},
	}
}
