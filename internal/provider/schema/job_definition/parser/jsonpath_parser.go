package parser

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func JsonpathParserSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "For files in jsonpath format, this parameter is required.",
		Attributes: map[string]schema.Attribute{
			"root": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "JSONPath",
			},
			"default_time_zone": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Default time zone",
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
						"time_zone": schema.StringAttribute{
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
