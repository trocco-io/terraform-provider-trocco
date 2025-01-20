package parser

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func XmlParserSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "For files in xml format, this parameter is required.",
		Attributes: map[string]schema.Attribute{
			"root": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Root element",
			},
			"columns": schema.ListNestedAttribute{
				Required: true,
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
						"path": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "XPath",
						},
						"timezone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "time zone",
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
