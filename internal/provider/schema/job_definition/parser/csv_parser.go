package parser

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func CsvParserSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "For files in CSV format, this parameter is required",
		Attributes: map[string]schema.Attribute{
			"delimiter": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Delimiter",
			},
			"quote": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Quote character",
			},
			"escape": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Escape character",
			},
			"skip_header_lines": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Number of header lines to skip",
			},
			"null_string_enabled": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "Flag whether or not to set the string to be replaced by NULL",
			},
			"null_string": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Replacement source string to be converted to NULL",
			},
			"trim_if_not_quoted": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "Flag whether or not to remove spaces from the value if it is not quoted",
			},
			"quotes_in_quoted_fields": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Processing method for irregular quarts",
				Validators:          []validator.String{stringvalidator.OneOf("ACCEPT_ONLY_RFC4180_ESCAPED", "ACCEPT_STRAY_QUOTES_ASSUMING_NO_DELIMITERS_IN_FIELDS")},
			},
			"comment_line_marker": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Comment line marker. Skip if this character is at the beginning of a line",
			},
			"allow_optional_columns": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "If true, NULL-complete the missing columns. If false, treat as invalid record.",
			},
			"allow_extra_columns": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "If true, ignore the column. If false, treat as invalid record.",
			},
			"max_quoted_size_limit": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Maximum amount of data that can be enclosed in quotation marks.",
			},
			"stop_on_invalid_record": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "Flag whether or not to abort the transfer if an invalid record is found.",
			},
			"default_time_zone": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Default time zone",
			},
			"default_date": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Default date",
			},
			"newline": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.
						OneOf(
							"CRLF",
							"LF",
							"CR",
						)},
				MarkdownDescription: "Newline character",
			},
			"charset": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Character set",
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
							Validators:          []validator.String{stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean")},
							MarkdownDescription: "Column type",
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Format of the column",
						},
						"date": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Date",
						},
					},
				},
			},
		},
	}
}