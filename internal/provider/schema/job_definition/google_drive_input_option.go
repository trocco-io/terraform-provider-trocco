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

func GoogleDriveInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes about source Google Drive",
		Attributes: map[string]schema.Attribute{
			"google_drive_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "ID of Google Drive connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"folder_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Google Drive folder ID",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"file_match_pattern": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "File name match pattern",
			},
			"is_skip_header_line": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Flag whether or not to skip the header line",
			},
			"stop_when_file_not_found": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Flag whether the transfer should continue if the file does not exist in the specified path",
			},
			"decompression_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Decompression type. Omit if the file is not compressed.",
				Validators: []validator.String{
					stringvalidator.OneOf("bzip2", "gzip", "targz", "zip"),
				},
			},
			"jsonpath_parser":          parser.JsonpathParserSchema(),
			"xml_parser":               parser.XmlParserSchema(),
			"excel_parser":             parser.ExcelParserSchema(),
			"ltsv_parser":              parser.LtsvParserSchema(),
			"jsonl_parser":             parser.JsonlParserSchema(),
			"csv_parser":               parser.CsvParserSchema(),
			"decoder":                  DecoderSchema(),
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&planModifier.FileParserPlanModifier{},
		},
	}
}
