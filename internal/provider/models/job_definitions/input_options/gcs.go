package input_options

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions/input_options"
	input_options2 "terraform-provider-trocco/internal/client/parameters/job_definitions/input_options"
	"terraform-provider-trocco/internal/provider/models"
	"terraform-provider-trocco/internal/provider/models/job_definitions/input_options/parser"
)

type GcsInputOption struct {
	GcsConnectionID           types.Int64                     `tfsdk:"gcs_connection_id"`
	Bucket                    types.String                    `tfsdk:"bucket"`
	PathPrefix                types.String                    `tfsdk:"path_prefix"`
	IncrementalLoadingEnabled types.Bool                      `tfsdk:"incremental_loading_enabled"`
	LastPath                  types.String                    `tfsdk:"last_path"`
	StopWhenFileNotFound      types.Bool                      `tfsdk:"stop_when_file_not_found"`
	DecompressionType         types.String                    `tfsdk:"decompression_type"`
	CsvParsers                *parser.CsvParser               `tfsdk:"csv_parsers"`
	JsonlParsers              *parser.JsonlParser             `tfsdk:"jsonl_parsers"`
	JsonpathParsers           *parser.JsonpathParser          `tfsdk:"jsonpath_parsers"`
	LtsvParsers               *parser.LtsvParser              `tfsdk:"ltsv_parsers"`
	ExcelParsers              *parser.ExcelParser             `tfsdk:"excel_parsers"`
	XmlParsers                *parser.XmlParser               `tfsdk:"xml_parsers"`
	ParquetParsers            *parser.ParquetParser           `tfsdk:"parquet_parsers"`
	CustomVariableSettings    *[]models.CustomVariableSetting `tfsdk:"custom_variable_settings"`
	Decoder                   *Decoder                        `tfsdk:"decoder"`
}

func NewGcsInputOption(gcsInputOption *input_options.GcsInputOption) *GcsInputOption {
	if gcsInputOption == nil {
		return nil
	}
	return &GcsInputOption{
		GcsConnectionID:           types.Int64Value(gcsInputOption.GcsConnectionID),
		Bucket:                    types.StringValue(gcsInputOption.Bucket),
		PathPrefix:                types.StringPointerValue(gcsInputOption.PathPrefix),
		IncrementalLoadingEnabled: types.BoolValue(gcsInputOption.IncrementalLoadingEnabled),
		LastPath:                  types.StringPointerValue(gcsInputOption.LastPath),
		StopWhenFileNotFound:      types.BoolValue(gcsInputOption.StopWhenFileNotFound),
		DecompressionType:         types.StringPointerValue(gcsInputOption.DecompressionType),
		CsvParsers:                parser.NewCsvParser(gcsInputOption.CsvParsers),
		JsonlParsers:              parser.NewJsonlParser(gcsInputOption.JsonlParsers),
		JsonpathParsers:           parser.NewJsonPathParser(gcsInputOption.JsonpathParsers),
		LtsvParsers:               parser.NewLtsvParser(gcsInputOption.LtsvParsers),
		ExcelParsers:              parser.NewExcelParser(gcsInputOption.ExcelParsers),
		XmlParsers:                parser.NewXmlParser(gcsInputOption.XmlParsers),
		ParquetParsers:            parser.NewParquetParser(gcsInputOption.ParquetParsers),
		CustomVariableSettings:    models.NewCustomVariableSettings(gcsInputOption.CustomVariableSettings),
		Decoder:                   NewDecoder(gcsInputOption.Decoder),
	}
}

func (gcsInputOption *GcsInputOption) ToInput() *input_options2.GcsInputOptionInput {
	if gcsInputOption == nil {
		return nil
	}

	return &input_options2.GcsInputOptionInput{
		GcsConnectionID:           gcsInputOption.GcsConnectionID.ValueInt64(),
		Bucket:                    gcsInputOption.Bucket.String(),
		PathPrefix:                gcsInputOption.PathPrefix.ValueStringPointer(),
		IncrementalLoadingEnabled: gcsInputOption.IncrementalLoadingEnabled.ValueBool(),
		LastPath:                  gcsInputOption.LastPath.ValueStringPointer(),
		StopWhenFileNotFound:      gcsInputOption.StopWhenFileNotFound.ValueBool(),
		DecompressionType:         gcsInputOption.DecompressionType.ValueStringPointer(),
		CsvParsers:                gcsInputOption.CsvParsers.ToCsvParserInput(),
		JsonlParsers:              gcsInputOption.JsonlParsers.ToJsonlParserInput(),
		JsonpathParsers:           gcsInputOption.JsonpathParsers.ToJsonpathParserInput(),
		LtsvParsers:               gcsInputOption.LtsvParsers.ToLtsvParserInput(),
		ExcelParsers:              gcsInputOption.ExcelParsers.ToExcelParserInput(),
		XmlParsers:                gcsInputOption.XmlParsers.ToXmlParserInput(),
		ParquetParsers:            gcsInputOption.ParquetParsers.ToParquetParserInput(),
		CustomVariableSettings:    models.ToCustomVariableSettingInputs(gcsInputOption.CustomVariableSettings),
		Decoder:                   gcsInputOption.Decoder.ToDecoderInput(),
	}
}
