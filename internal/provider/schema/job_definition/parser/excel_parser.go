package parser

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func ExcelParserSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "For files in excel format, this parameter is required.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"default_time_zone": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Default time zone",
			},
			"sheet_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Sheet name",
			},
			"skip_header_lines": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Number of header lines to skip",
			},
			"columns": schema.ListNestedAttribute{
				MarkdownDescription: "List of columns to be retrieved and their types",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column name",
						},
						"type": schema.StringAttribute{
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean", "json")},
							MarkdownDescription: "Column type",
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Format of the column.",
						},
						"formula_handling": schema.StringAttribute{
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("cashed_value", "evaluate")},
							MarkdownDescription: "Formula handling",
						},
					},
				},
			},
		},
	}
}
