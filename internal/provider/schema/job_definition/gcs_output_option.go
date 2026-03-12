package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	planModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func GcsOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination Google Cloud Storage settings",
		Attributes: map[string]schema.Attribute{
			"gcs_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of GCS connection",
			},
			"bucket": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "GCS bucket name",
			},
			"path_prefix": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "File path prefix. Can contain custom variables (e.g., $start_time$)",
			},
			"file_ext": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "File extension",
			},
			"sequence_format": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(".%03d.%02d"),
				MarkdownDescription: "Sequence format for output files",
			},
			"is_minimum_output_tasks": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "Output file mode. true = output file number suppression mode, false = parallel transfer",
			},
			"formatter_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("csv", "jsonl"),
				},
				MarkdownDescription: "Formatter type. Valid values: `csv`, `jsonl`",
			},
			"encoder_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "gzip", "bzip2", "zip"),
				},
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "Encoder type. Valid values: `` (no compression), `gzip`, `bzip2`, `zip`",
			},
			"csv_formatter": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "CSV formatter configuration. Required when formatter_type is `csv`",
				Attributes: map[string]schema.Attribute{
					"delimiter": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString(","),
						MarkdownDescription: "Delimiter character",
					},
					"newline": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString("CRLF"),
						Validators: []validator.String{
							stringvalidator.OneOf("CRLF", "LF", "CR"),
						},
						MarkdownDescription: "Newline character. Valid values: `CRLF`, `LF`, `CR`",
					},
					"newline_in_field": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString("LF"),
						Validators: []validator.String{
							stringvalidator.OneOf("CRLF", "LF", "CR"),
						},
						MarkdownDescription: "Newline character in field. Valid values: `CRLF`, `LF`, `CR`",
					},
					"charset": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("UTF-8"),
						MarkdownDescription: "Character encoding",
					},
					"quote_policy": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString("MINIMAL"),
						Validators: []validator.String{
							stringvalidator.OneOf("ALL", "MINIMAL", "NONE"),
						},
						MarkdownDescription: "Quote policy. Valid values: `ALL`, `MINIMAL`, `NONE`",
					},
					"escape": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("\\"),
						MarkdownDescription: "Escape character",
					},
					"header_line": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						Default:             booldefault.StaticBool(true),
						MarkdownDescription: "Whether to include header line",
					},
					"null_string_enabled": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						Default:             booldefault.StaticBool(false),
						MarkdownDescription: "Whether to enable null string representation",
					},
					"null_string": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Null string representation",
					},
					"default_time_zone": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("UTC"),
						MarkdownDescription: "Default timezone",
					},
					"csv_formatter_column_options_attributes": schema.ListNestedAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Column-specific options",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "Column name",
								},
								"format": schema.StringAttribute{
									Optional:            true,
									Computed:            true,
									Default:             stringdefault.StaticString("%Y-%m-%d %H:%M:%S.%6N %z"),
									MarkdownDescription: "Date format",
								},
								"timezone": schema.StringAttribute{
									Optional:            true,
									MarkdownDescription: "Timezone",
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
						MarkdownDescription: "Character encoding. Valid values: `UTF-8`, `UTF-16LE`, `UTF-32BE`, `UTF-32LE`",
					},
					"newline": schema.StringAttribute{
						Optional: true,
						Computed: true,
						Default:  stringdefault.StaticString("LF"),
						Validators: []validator.String{
							stringvalidator.OneOf("CRLF", "LF", "CR", "NUL", "NO"),
						},
						MarkdownDescription: "Newline character. Valid values: `CRLF`, `LF`, `CR`, `NUL`, `NO`",
					},
					"date_format": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Date format",
					},
					"timezone": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Timezone",
					},
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
