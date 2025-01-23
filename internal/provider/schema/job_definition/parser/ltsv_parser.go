package parser

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func LtsvParserSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "For files in LTSV format, this parameter is required.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"newline": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("CRLF"),
				MarkdownDescription: "Newline character",
			},
			"charset": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Character set",
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
							Required:            true,
							Validators:          []validator.String{stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean")},
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
