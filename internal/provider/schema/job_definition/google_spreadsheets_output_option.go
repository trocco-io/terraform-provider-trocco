package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func GoogleSpreadsheetsOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination Snowflake settings",
		Attributes: map[string]schema.Attribute{
			"google_spreadsheets_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Snowflake connection ID",
			},
			"spreadsheets_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Spreadsheet ID",
			},
			"worksheet_title": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Worksheet title",
			},
			"timezone": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "Time zone",
				Default:             stringdefault.StaticString("Asia/Tokyo"),
			},
			"value_input_option": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "Value input option",
				Default:             stringdefault.StaticString("RAW"),
				Validators: []validator.String{
					stringvalidator.OneOf("RAW", "USER_ENTERED"),
				},
			},
			"mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("append"),
				Validators: []validator.String{
					stringvalidator.OneOf("append", "replace", "truncate_insert"),
				},
				MarkdownDescription: "Transfer mode",
			},
			"google_spreadsheets_output_option_sorts": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"column": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column name",
						},
						"order": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Data type",
							Validators: []validator.String{
								stringvalidator.OneOf("ascending", "descending"),
							},
						},
					},
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
		// PlanModifiers: []planmodifier.Object{
		// 	&planmodifier2.GoogleSpreadsheetsOutputOptionPlanModifier{},
		// },
	}
}
