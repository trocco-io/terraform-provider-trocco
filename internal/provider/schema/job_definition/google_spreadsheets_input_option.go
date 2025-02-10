package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func GoogleSpreadsheetsInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes about source Google Spreadsheets",
		Attributes: map[string]schema.Attribute{
			"google_spreadsheets_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Id of Snowflake connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
			"input_option_columns": schema.ListNestedAttribute{
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
							MarkdownDescription: "Column type",
							Validators: []validator.String{
								stringvalidator.OneOf("boolean", "long", "timestamp", "double", "string", "json"),
							},
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Column format",
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
		},
	}
}
