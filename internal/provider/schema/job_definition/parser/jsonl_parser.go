package parser

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func JsonlParserSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "For files in JSONL format, this parameter is required",
		Attributes: map[string]schema.Attribute{
			"stop_on_invalid_record": schema.BoolAttribute{
				Computed:            true,
				Optional:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Flag whether the transfer should stop if an invalid record is found",
			},
			"default_time_zone": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString("UTC"),
				MarkdownDescription: "Default time zone",
			},
			"newline": schema.StringAttribute{
				Computed:            true,
				Optional:            true,
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
							MarkdownDescription: "Column type",
							Validators:          []validator.String{stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean", "json")},
						},
						"time_zone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "time zone",
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Format of the column",
						},
					},
				},
			},
		},
	}
}
