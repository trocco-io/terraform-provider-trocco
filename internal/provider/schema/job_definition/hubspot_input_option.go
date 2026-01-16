package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func HubspotInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "attributes of source HubSpot",
		Attributes: map[string]schema.Attribute{
			"hubspot_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "id of HubSpot connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"target": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "type of data to retrieve from HubSpot",
				Validators: []validator.String{
					stringvalidator.OneOf("object", "association", "engagement", "engagement_association", "email_event", "pipeline", "pipeline_stage", "owner"),
				},
			},
			"from_object_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "source object type (required when target is association)",
			},
			"to_object_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "destination object type (required when target is association, engagement_association, pipeline, or pipeline_stage)",
			},
			"object_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "object type (required when target is object, pipeline, or pipeline_stage)",
			},
			"incremental_loading_enabled": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "enable incremental loading (only valid when target is object)",
			},
			"last_record_time": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "last record time (used when incremental loading is enabled)",
			},
			"email_event_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "email event type (required when target is email_event)",
			},
			"start_timestamp": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "start timestamp (used when target is email_event)",
			},
			"end_timestamp": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "end timestamp (used when target is email_event)",
			},
			"input_option_columns": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "list of columns to be retrieved and their types",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
							MarkdownDescription: "column name",
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "column type",
							Validators: []validator.String{
								stringvalidator.OneOf("boolean", "long", "timestamp", "double", "string", "json"),
							},
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "column format",
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
