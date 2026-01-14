package job_definition

import (
	planModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func SftpOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "attributes of destination SFTP settings",
		Attributes: map[string]schema.Attribute{
			"sftp_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "id of SFTP connection",
			},
			"path_prefix": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "path prefix for output files. Can contain custom variables (e.g., $start_time$)",
			},
			"file_ext": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "file extension",
			},
			"is_minimum_output_tasks": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "output file mode. true = output file number suppression mode, false = parallel transfer",
			},
			"encoder_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "gzip", "bzip2", "zip"),
				},
				MarkdownDescription: "encoder type. Valid values: `` (no compression), `gzip`, `bzip2`, `zip`",
			},
			"csv_formatter": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "CSV formatter configuration. Required when formatter_type is `csv`",
				Attributes: map[string]schema.Attribute{
					"delimiter": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString(","),
						MarkdownDescription: "delimiter character",
					},
					"newline": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString("CRLF"),
						Validators: []validator.String{
							stringvalidator.OneOf("CRLF", "LF", "CR"),
						},
						MarkdownDescription: "newline character. Valid values: `CRLF`, `LF`, `CR`",
					},
					"newline_in_field": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString("LF"),
						Validators: []validator.String{
							stringvalidator.OneOf("CRLF", "LF", "CR"),
						},
						MarkdownDescription: "newline character in field. Valid values: `CRLF`, `LF`, `CR`",
					},
					"charset": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("UTF-8"),
						MarkdownDescription: "character encoding",
					},
					"quote_policy": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString("MINIMAL"),
						Validators: []validator.String{
							stringvalidator.OneOf("ALL", "MINIMAL", "NONE"),
						},
						MarkdownDescription: "quote policy. Valid values: `ALL`, `MINIMAL`, `NONE`",
					},
					"escape": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("\\"),
						MarkdownDescription: "escape character",
					},
					"header_line": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						Default:             booldefault.StaticBool(true),
						MarkdownDescription: "whether to include header line",
					},
					"null_string_enabled": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						Default:             booldefault.StaticBool(false),
						MarkdownDescription: "whether to enable null string representation",
					},
					"null_string": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "null string representation",
					},
					"default_time_zone": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("UTC"),
						MarkdownDescription: "default timezone",
					},
					"csv_formatter_column_options_attributes": schema.ListNestedAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "column-specific options",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "column name",
								},
								"format": schema.StringAttribute{
									Optional:            true,
									Computed:            true,
									Default:             stringdefault.StaticString("%Y-%m-%d %H:%M:%S.%6N %z"),
									MarkdownDescription: "date format",
								},
								"timezone": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "timezone",
								},
							},
						},
						PlanModifiers: []planmodifier.List{
							planModifier.EmptyListForNull(),
						},
					},
				},
			},
			"jsonl_formatter": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "JSONL formatter configuration. Required when formatter_type is `jsonl`",
				Attributes: map[string]schema.Attribute{
					"encoding": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString("UTF-8"),
						Validators: []validator.String{
							stringvalidator.OneOf("UTF-8", "UTF-16LE", "UTF-32BE", "UTF-32LE"),
						},
						MarkdownDescription: "character encoding. Valid values: `UTF-8`, `UTF-16LE`, `UTF-32BE`, `UTF-32LE`",
					},
					"newline": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString("LF"),
						Validators: []validator.String{
							stringvalidator.OneOf("CRLF", "LF", "CR", "NUL", "NO"),
						},
						MarkdownDescription: "newline character. Valid values: `CRLF`, `LF`, `CR`, `NUL`, `NO`",
					},
					"date_format": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "date format",
					},
					"timezone": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "timezone",
					},
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
