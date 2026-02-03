package job_definition

import (
	planModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func S3OutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination Amazon S3 settings",
		Attributes: map[string]schema.Attribute{
			"s3_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Id of S3 connection",
			},
			"bucket": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "S3 bucket name",
			},
			"path_prefix": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "File path prefix",
			},
			"region": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "AWS region",
			},
			"file_ext": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "File extension",
			},
			"sequence_format": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Sequence format",
			},
			"canned_acl": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "S3 object ACL",
			},
			"is_minimum_output_tasks": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Minimum output tasks setting",
			},
			"multipart_upload_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Enable multipart upload",
			},
			"formatter_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf("csv", "jsonl"),
				},
				MarkdownDescription: "Formatter type",
			},
			"encoder_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf("", "gzip", "bzip2", "zip"),
				},
				MarkdownDescription: "Encoder type",
			},
			"csv_formatter": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "CSV formatter settings. Required when formatter_type is csv",
				Attributes: map[string]schema.Attribute{
					"delimiter": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Delimiter character. Use \\t for tab",
					},
					"escape": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Escape character",
					},
					"header_line": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Output header line",
					},
					"charset": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Character encoding (e.g., UTF-8, Shift_JIS)",
					},
					"quote_policy": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Quote policy",
					},
					"newline": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Newline code (e.g., LF, CRLF)",
					},
					"newline_in_field": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Newline processing in field",
					},
					"null_string_enabled": schema.BoolAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Enable null string",
					},
					"null_string": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "String to represent null values",
					},
					"default_time_zone": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Default time zone",
					},
					"csv_formatter_column_options_attributes": schema.ListNestedAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Column options for formatting",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required:            true,
									MarkdownDescription: "Column name",
								},
								"format": schema.StringAttribute{
									Optional:            true,
									Computed:            true,
									MarkdownDescription: "Format specification (e.g., date/time format)",
								},
								"timezone": schema.StringAttribute{
									Optional:            true,
									Computed:            true,
									MarkdownDescription: "Time zone for this column",
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
				MarkdownDescription: "JSONL formatter settings. Required when formatter_type is jsonl",
				Attributes: map[string]schema.Attribute{
					"encoding": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Character encoding (e.g., UTF-8)",
					},
					"newline": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Newline code (e.g., LF, CRLF)",
					},
					"date_format": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Date format",
					},
					"timezone": schema.StringAttribute{
						Optional:            true,
						Computed:            true,
						MarkdownDescription: "Time zone",
					},
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&planModifier.S3OutputOptionPlanModifier{},
		},
	}
}
