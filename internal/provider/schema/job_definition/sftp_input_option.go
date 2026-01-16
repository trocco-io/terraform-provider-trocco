package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	planModifier "terraform-provider-trocco/internal/provider/planmodifier"
	"terraform-provider-trocco/internal/provider/schema/job_definition/parser"
)

func SftpInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes about source SFTP",
		Attributes: map[string]schema.Attribute{
			"sftp_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "ID of SFTP connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"path_prefix": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Path to file/folder on SFTP server. Supports custom variables (replaced at runtime).",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"path_match_pattern": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "Pattern to match files (regex). Supports custom variables.",
			},
			"incremental_loading_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "If true, enables incremental loading. If false, performs all record loading.",
			},
			"last_path": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Path of last transferred file (for incremental loading). Only used when incremental_loading_enabled is true.",
			},
			"stop_when_file_not_found": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "If true, stop job with error when file not found. If false, continue job (skip file).",
			},
			"decompression_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("guess"),
				MarkdownDescription: "Compression type of file. Valid values: guess (auto-detect), zip, targz.",
				Validators: []validator.String{
					stringvalidator.OneOf("guess", "zip", "targz"),
				},
			},
			"decoder":                  DecoderSchema(),
			"custom_variable_settings": CustomVariableSettingsSchema(),
			"csv_parser":               parser.CsvParserSchema(),
			"jsonl_parser":             parser.JsonlParserSchema(),
			"jsonpath_parser":          parser.JsonpathParserSchema(),
			"ltsv_parser":              parser.LtsvParserSchema(),
			"excel_parser":             parser.ExcelParserSchema(),
			"xml_parser":               parser.XmlParserSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&planModifier.FileParserPlanModifier{},
		},
	}
}
