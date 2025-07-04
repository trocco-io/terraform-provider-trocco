package input_options

import (
	"context"
	input_options "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	input_options2 "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"
	"terraform-provider-trocco/internal/provider/model/job_definition/input_option/parser"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GcsInputOption struct {
	GcsConnectionID           types.Int64            `tfsdk:"gcs_connection_id"`
	Bucket                    types.String           `tfsdk:"bucket"`
	PathPrefix                types.String           `tfsdk:"path_prefix"`
	IncrementalLoadingEnabled types.Bool             `tfsdk:"incremental_loading_enabled"`
	LastPath                  types.String           `tfsdk:"last_path"`
	StopWhenFileNotFound      types.Bool             `tfsdk:"stop_when_file_not_found"`
	DecompressionType         types.String           `tfsdk:"decompression_type"`
	CsvParser                 *parser.CsvParser      `tfsdk:"csv_parser"`
	JsonlParser               *parser.JsonlParser    `tfsdk:"jsonl_parser"`
	JsonpathParser            *parser.JsonpathParser `tfsdk:"jsonpath_parser"`
	LtsvParser                *parser.LtsvParser     `tfsdk:"ltsv_parser"`
	ExcelParser               *parser.ExcelParser    `tfsdk:"excel_parser"`
	XmlParser                 *parser.XmlParser      `tfsdk:"xml_parser"`
	ParquetParser             *parser.ParquetParser  `tfsdk:"parquet_parser"`
	CustomVariableSettings    types.List             `tfsdk:"custom_variable_settings"`
	Decoder                   *Decoder               `tfsdk:"decoder"`
}

func NewGcsInputOption(ctx context.Context, gcsInputOption *input_options.GcsInputOption) *GcsInputOption {
	if gcsInputOption == nil {
		return nil
	}

	result := &GcsInputOption{
		GcsConnectionID:           types.Int64Value(gcsInputOption.GcsConnectionID),
		Bucket:                    types.StringValue(gcsInputOption.Bucket),
		PathPrefix:                types.StringValue(gcsInputOption.PathPrefix),
		IncrementalLoadingEnabled: types.BoolValue(gcsInputOption.IncrementalLoadingEnabled),
		LastPath:                  types.StringPointerValue(gcsInputOption.LastPath),
		StopWhenFileNotFound:      types.BoolValue(gcsInputOption.StopWhenFileNotFound),
		DecompressionType:         types.StringPointerValue(gcsInputOption.DecompressionType),
		CsvParser:                 parser.NewCsvParser(ctx, gcsInputOption.CsvParser),
		JsonlParser:               parser.NewJsonlParser(ctx, gcsInputOption.JsonlParser),
		JsonpathParser:            parser.NewJsonPathParser(ctx, gcsInputOption.JsonpathParser),
		LtsvParser:                parser.NewLtsvParser(ctx, gcsInputOption.LtsvParser),
		ExcelParser:               parser.NewExcelParser(ctx, gcsInputOption.ExcelParser),
		XmlParser:                 parser.NewXmlParser(ctx, gcsInputOption.XmlParser),
		ParquetParser:             parser.NewParquetParser(ctx, gcsInputOption.ParquetParser),
		Decoder:                   NewDecoder(gcsInputOption.Decoder),
	}

	CustomVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, gcsInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = CustomVariableSettings

	return result
}

func (gcsInputOption *GcsInputOption) ToInput(ctx context.Context) *input_options2.GcsInputOptionInput {
	if gcsInputOption == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, gcsInputOption.CustomVariableSettings)

	return &input_options2.GcsInputOptionInput{
		GcsConnectionID:           gcsInputOption.GcsConnectionID.ValueInt64(),
		Bucket:                    gcsInputOption.Bucket.ValueString(),
		PathPrefix:                gcsInputOption.PathPrefix.ValueString(),
		IncrementalLoadingEnabled: gcsInputOption.IncrementalLoadingEnabled.ValueBool(),
		LastPath:                  model.NewNullableString(gcsInputOption.LastPath),
		StopWhenFileNotFound:      gcsInputOption.StopWhenFileNotFound.ValueBool(),
		DecompressionType:         model.NewNullableString(gcsInputOption.DecompressionType),
		CsvParser:                 gcsInputOption.CsvParser.ToCsvParserInput(ctx),
		JsonlParser:               gcsInputOption.JsonlParser.ToJsonlParserInput(ctx),
		JsonpathParser:            gcsInputOption.JsonpathParser.ToJsonpathParserInput(ctx),
		LtsvParser:                gcsInputOption.LtsvParser.ToLtsvParserInput(ctx),
		ExcelParser:               gcsInputOption.ExcelParser.ToExcelParserInput(ctx),
		XmlParser:                 gcsInputOption.XmlParser.ToXmlParserInput(ctx),
		ParquetParser:             gcsInputOption.ParquetParser.ToParquetParserInput(ctx),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
		Decoder:                   gcsInputOption.Decoder.ToDecoderInput(),
	}
}

func (gcsInputOption *GcsInputOption) ToUpdateInput(ctx context.Context) *input_options2.UpdateGcsInputOptionInput {
	if gcsInputOption == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, gcsInputOption.CustomVariableSettings)

	return &input_options2.UpdateGcsInputOptionInput{
		GcsConnectionID:           gcsInputOption.GcsConnectionID.ValueInt64Pointer(),
		Bucket:                    gcsInputOption.Bucket.ValueStringPointer(),
		PathPrefix:                gcsInputOption.PathPrefix.ValueStringPointer(),
		IncrementalLoadingEnabled: gcsInputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		LastPath:                  model.NewNullableString(gcsInputOption.LastPath),
		StopWhenFileNotFound:      gcsInputOption.StopWhenFileNotFound.ValueBoolPointer(),
		DecompressionType:         model.NewNullableString(gcsInputOption.DecompressionType),
		CsvParser:                 gcsInputOption.CsvParser.ToCsvParserInput(ctx),
		JsonlParser:               gcsInputOption.JsonlParser.ToJsonlParserInput(ctx),
		JsonpathParser:            gcsInputOption.JsonpathParser.ToJsonpathParserInput(ctx),
		LtsvParser:                gcsInputOption.LtsvParser.ToLtsvParserInput(ctx),
		ExcelParser:               gcsInputOption.ExcelParser.ToExcelParserInput(ctx),
		XmlParser:                 gcsInputOption.XmlParser.ToXmlParserInput(ctx),
		ParquetParser:             gcsInputOption.ParquetParser.ToParquetParserInput(ctx),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
		Decoder:                   gcsInputOption.Decoder.ToDecoderInput(),
	}
}
