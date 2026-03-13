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

func GoogleDriveOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination Google Drive settings",
		Attributes: map[string]schema.Attribute{
			"google_drive_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of Google Drive connection",
			},
			"main_folder_id": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Google Drive folder ID",
			},
			"child_folder_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Child folder name",
			},
			"file_name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Output file name",
			},
			"formatter_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf("csv"),
				},
				MarkdownDescription: "Formatter type",
			},
			"csv_formatter": schema.SingleNestedAttribute{
				Optional:            true,
				MarkdownDescription: "CSV formatter settings",
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
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
