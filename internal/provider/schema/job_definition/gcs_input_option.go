package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	planModifier "terraform-provider-trocco/internal/provider/planmodifier"
	"terraform-provider-trocco/internal/provider/schema/job_definition/parser"
)

func GcsInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes about source GCS",
		Attributes: map[string]schema.Attribute{
			"bucket": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Bucket name",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"path_prefix": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Path prefix",
			},
			"incremental_loading_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "If it is true, to be incremental loading. If it is false, to be all record loading",
			},
			"last_path": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Last path transferred. It is only enabled when incremental loading is true. When updating differences, data behind in lexicographic order from the path specified here is transferred. If the form is blank, the data is transferred from the beginning. Do not change this value unless there is a special reason. Duplicate data may occur.",
			},
			"stop_when_file_not_found": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Flag whether the transfer should continue if the file does not exist in the specified path",
			},
			"gcs_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Id of GCS connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"decompression_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Decompression type",
				Validators: []validator.String{
					stringvalidator.OneOf("gzip", "bzip2", "zip", "targz"),
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
			&planModifier.FileParserPlanModifier{},
			&planModifier.GcsInputOptionPlanModifier{},
		},
	}
}
