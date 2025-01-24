package input_options

import (
	"terraform-provider-trocco/internal/client/entities/job_definitions/input_options"
	input_options2 "terraform-provider-trocco/internal/client/parameter/job_definitions/input_options"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/models"
	"terraform-provider-trocco/internal/provider/models/job_definitions/input_options/parser"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GcsInputOption struct {
	GcsConnectionID           types.Int64                     `tfsdk:"gcs_connection_id"`
	Bucket                    types.String                    `tfsdk:"bucket"`
	PathPrefix                types.String                    `tfsdk:"path_prefix"`
	IncrementalLoadingEnabled types.Bool                      `tfsdk:"incremental_loading_enabled"`
	LastPath                  types.String                    `tfsdk:"last_path"`
	StopWhenFileNotFound      types.Bool                      `tfsdk:"stop_when_file_not_found"`
	DecompressionType         types.String                    `tfsdk:"decompression_type"`
	CsvParser                 *parser.CsvParser               `tfsdk:"csv_parser"`
	JsonlParser               *parser.JsonlParser             `tfsdk:"jsonl_parser"`
	JsonpathParser            *parser.JsonpathParser          `tfsdk:"jsonpath_parser"`
	LtsvParser                *parser.LtsvParser              `tfsdk:"ltsv_parser"`
	ExcelParser               *parser.ExcelParser             `tfsdk:"excel_parser"`
	XmlParser                 *parser.XmlParser               `tfsdk:"xml_parser"`
	ParquetParser             *parser.ParquetParser           `tfsdk:"parquet_parser"`
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
		PathPrefix:                types.StringValue(gcsInputOption.PathPrefix),
		IncrementalLoadingEnabled: types.BoolValue(gcsInputOption.IncrementalLoadingEnabled),
		LastPath:                  types.StringPointerValue(gcsInputOption.LastPath),
		StopWhenFileNotFound:      types.BoolValue(gcsInputOption.StopWhenFileNotFound),
		DecompressionType:         types.StringPointerValue(gcsInputOption.DecompressionType),
		CsvParser:                 parser.NewCsvParser(gcsInputOption.CsvParser),
		JsonlParser:               parser.NewJsonlParser(gcsInputOption.JsonlParser),
		JsonpathParser:            parser.NewJsonPathParser(gcsInputOption.JsonpathParser),
		LtsvParser:                parser.NewLtsvParser(gcsInputOption.LtsvParser),
		ExcelParser:               parser.NewExcelParser(gcsInputOption.ExcelParser),
		XmlParser:                 parser.NewXmlParser(gcsInputOption.XmlParser),
		ParquetParser:             parser.NewParquetParser(gcsInputOption.ParquetParser),
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
		Bucket:                    gcsInputOption.Bucket.ValueString(),
		PathPrefix:                gcsInputOption.PathPrefix.ValueString(),
		IncrementalLoadingEnabled: gcsInputOption.IncrementalLoadingEnabled.ValueBool(),
		LastPath:                  model.NewNullableString(gcsInputOption.LastPath),
		StopWhenFileNotFound:      gcsInputOption.StopWhenFileNotFound.ValueBool(),
		DecompressionType:         model.NewNullableString(gcsInputOption.DecompressionType),
		CsvParser:                 gcsInputOption.CsvParser.ToCsvParserInput(),
		JsonlParser:               gcsInputOption.JsonlParser.ToJsonlParserInput(),
		JsonpathParser:            gcsInputOption.JsonpathParser.ToJsonpathParserInput(),
		LtsvParser:                gcsInputOption.LtsvParser.ToLtsvParserInput(),
		ExcelParser:               gcsInputOption.ExcelParser.ToExcelParserInput(),
		XmlParser:                 gcsInputOption.XmlParser.ToXmlParserInput(),
		ParquetParser:             gcsInputOption.ParquetParser.ToParquetParserInput(),
		CustomVariableSettings:    models.ToCustomVariableSettingInputs(gcsInputOption.CustomVariableSettings),
		Decoder:                   gcsInputOption.Decoder.ToDecoderInput(),
	}
}

func (gcsInputOption *GcsInputOption) ToUpdateInput() *input_options2.UpdateGcsInputOptionInput {
	if gcsInputOption == nil {
		return nil
	}

	return &input_options2.UpdateGcsInputOptionInput{
		GcsConnectionID:           gcsInputOption.GcsConnectionID.ValueInt64Pointer(),
		Bucket:                    gcsInputOption.Bucket.ValueStringPointer(),
		PathPrefix:                gcsInputOption.PathPrefix.ValueStringPointer(),
		IncrementalLoadingEnabled: gcsInputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		LastPath:                  model.NewNullableString(gcsInputOption.LastPath),
		StopWhenFileNotFound:      gcsInputOption.StopWhenFileNotFound.ValueBoolPointer(),
		DecompressionType:         model.NewNullableString(gcsInputOption.DecompressionType),
		CsvParser:                 gcsInputOption.CsvParser.ToCsvParserInput(),
		JsonlParser:               gcsInputOption.JsonlParser.ToJsonlParserInput(),
		JsonpathParser:            gcsInputOption.JsonpathParser.ToJsonpathParserInput(),
		LtsvParser:                gcsInputOption.LtsvParser.ToLtsvParserInput(),
		ExcelParser:               gcsInputOption.ExcelParser.ToExcelParserInput(),
		XmlParser:                 gcsInputOption.XmlParser.ToXmlParserInput(),
		ParquetParser:             gcsInputOption.ParquetParser.ToParquetParserInput(),
		CustomVariableSettings:    models.ToCustomVariableSettingInputs(gcsInputOption.CustomVariableSettings),
		Decoder:                   gcsInputOption.Decoder.ToDecoderInput(),
	}
}
