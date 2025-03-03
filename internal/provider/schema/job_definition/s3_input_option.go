package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	planmodifier2 "terraform-provider-trocco/internal/provider/planmodifier"
	"terraform-provider-trocco/internal/provider/schema/job_definition/parser"
)

func S3InputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes about source S3",
		Attributes: map[string]schema.Attribute{
			"bucket": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Bucket name",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"path_prefix": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "Path prefix. If not entered, all files under the bucket will be targeted.",
			},
			"path_match_pattern": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "Path regular expression. If not entered, all files matching the path prefix will be included.",
			},
			"region": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("ap-northeast-1"),
				MarkdownDescription: "Region",
			},
			"incremental_loading_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "If it is true, to be incremental loading. If it is false, to be all record loading",
			},
			"is_skip_header_line": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Flag whether or not to skip header columns For CSV/TSV files that do not contain header columns, a temporary header name generated on the TROCCO side is assigned.",
			},
			"stop_when_file_not_found": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Flag whether the transfer should continue if the file does not exist in the specified path",
			},
			"s3_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Id of S3 connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"decompression_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("default"),
				MarkdownDescription: "Decompression type",
				Validators: []validator.String{
					stringvalidator.OneOf("default", "zip", "targz"),
				},
			},
			"parquet_parser":           parser.ParquetParserSchema(),
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
			&planmodifier2.FileParserPlanModifier{},
		},
	}
}
