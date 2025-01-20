package parser

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func ParquetParserSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "For files in parquet format, this parameter is required.",
		Attributes: map[string]schema.Attribute{
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
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean", "json")},
							MarkdownDescription: "Column type",
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Format of the column.",
						},
					},
				},
			},
		},
	}
}
