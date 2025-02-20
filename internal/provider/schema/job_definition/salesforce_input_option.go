package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	planmodifier2 "terraform-provider-trocco/internal/provider/planmodifier"
)

func SalesforceInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes about source Salesforce",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Object name",
			},
			"api_version": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Api version",
				Default:             stringdefault.StaticString("54.0"),
			},
			"soql": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "SOQL. If object_acquisition_method is 'soql', this field is required.",
			},
			"object_acquisition_method": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf("all_columns", "soql"),
				},
				MarkdownDescription: "Object Acquisition Method. If 'all_columns' is specified, soql is automatically completed.",
				Default:             stringdefault.StaticString("all_columns"),
			},
			"is_convert_type_custom_columns": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "false: transfer based on custom columns. true: Change custom columns to STRING type. If you select change custom columns to STRING type, all custom columns except BOOLEAN type will be changed to STRING type.",
			},
			"include_deleted_or_archived_records": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Extraction of deleted and archived records",
			},
			"salesforce_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Id of Salesforce connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"columns": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "List of columns to be retrieved and their types",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
							MarkdownDescription: "Column name",
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column type.",
							Validators: []validator.String{
								stringvalidator.OneOf("boolean", "long", "timestamp", "double", "string", "json"),
							},
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "format",
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&planmodifier2.SalesforceInputOptionPlanModifier{},
		},
	}
}
