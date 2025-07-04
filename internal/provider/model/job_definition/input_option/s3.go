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

type S3InputOption struct {
	S3ConnectionID            types.Int64            `tfsdk:"s3_connection_id"`
	Bucket                    types.String           `tfsdk:"bucket"`
	PathPrefix                types.String           `tfsdk:"path_prefix"`
	PathMatchPattern          types.String           `tfsdk:"path_match_pattern"`
	Region                    types.String           `tfsdk:"region"`
	IncrementalLoadingEnabled types.Bool             `tfsdk:"incremental_loading_enabled"`
	IsSkipHeaderLine          types.Bool             `tfsdk:"is_skip_header_line"`
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

func NewS3InputOption(ctx context.Context, s3InputOption *input_options.S3InputOption) *S3InputOption {
	if s3InputOption == nil {
		return nil
	}

	result := &S3InputOption{
		S3ConnectionID:            types.Int64Value(s3InputOption.S3ConnectionID),
		Bucket:                    types.StringValue(s3InputOption.Bucket),
		PathPrefix:                types.StringValue(s3InputOption.PathPrefix),
		PathMatchPattern:          types.StringValue(s3InputOption.PathMatchPattern),
		Region:                    types.StringValue(s3InputOption.Region),
		IncrementalLoadingEnabled: types.BoolValue(s3InputOption.IncrementalLoadingEnabled),
		IsSkipHeaderLine:          types.BoolValue(s3InputOption.IsSkipHeaderLine),
		StopWhenFileNotFound:      types.BoolValue(s3InputOption.StopWhenFileNotFound),
		DecompressionType:         types.StringValue(s3InputOption.DecompressionType),
		CsvParser:                 parser.NewCsvParser(ctx, s3InputOption.CsvParser),
		JsonlParser:               parser.NewJsonlParser(ctx, s3InputOption.JsonlParser),
		JsonpathParser:            parser.NewJsonPathParser(ctx, s3InputOption.JsonpathParser),
		LtsvParser:                parser.NewLtsvParser(ctx, s3InputOption.LtsvParser),
		ExcelParser:               parser.NewExcelParser(ctx, s3InputOption.ExcelParser),
		XmlParser:                 parser.NewXmlParser(ctx, s3InputOption.XmlParser),
		ParquetParser:             parser.NewParquetParser(ctx, s3InputOption.ParquetParser),
		Decoder:                   NewDecoder(s3InputOption.Decoder),
	}

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, s3InputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func (s3InputOption *S3InputOption) ToInput(ctx context.Context) *input_options2.S3InputOptionInput {
	if s3InputOption == nil {
		return nil
	}
	customVarSettings := common.ExtractCustomVariableSettings(ctx, s3InputOption.CustomVariableSettings)

	return &input_options2.S3InputOptionInput{
		S3ConnectionID:            s3InputOption.S3ConnectionID.ValueInt64(),
		Bucket:                    s3InputOption.Bucket.ValueString(),
		PathPrefix:                model.NewNullableString(s3InputOption.PathPrefix),
		PathMatchPattern:          model.NewNullableString(s3InputOption.PathMatchPattern),
		Region:                    model.NewNullableString(s3InputOption.Region),
		IncrementalLoadingEnabled: model.NewNullableBool(s3InputOption.IncrementalLoadingEnabled),
		IsSkipHeaderLine:          model.NewNullableBool(s3InputOption.IsSkipHeaderLine),
		StopWhenFileNotFound:      model.NewNullableBool(s3InputOption.StopWhenFileNotFound),
		DecompressionType:         model.NewNullableString(s3InputOption.DecompressionType),
		CsvParser:                 s3InputOption.CsvParser.ToCsvParserInput(ctx),
		JsonlParser:               s3InputOption.JsonlParser.ToJsonlParserInput(ctx),
		JsonpathParser:            s3InputOption.JsonpathParser.ToJsonpathParserInput(ctx),
		LtsvParser:                s3InputOption.LtsvParser.ToLtsvParserInput(ctx),
		ExcelParser:               s3InputOption.ExcelParser.ToExcelParserInput(ctx),
		XmlParser:                 s3InputOption.XmlParser.ToXmlParserInput(ctx),
		ParquetParser:             s3InputOption.ParquetParser.ToParquetParserInput(ctx),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
		Decoder:                   s3InputOption.Decoder.ToDecoderInput(),
	}
}

func (s3InputOption *S3InputOption) ToUpdateInput(ctx context.Context) *input_options2.UpdateS3InputOptionInput {
	if s3InputOption == nil {
		return nil
	}
	customVarSettings := common.ExtractCustomVariableSettings(ctx, s3InputOption.CustomVariableSettings)

	return &input_options2.UpdateS3InputOptionInput{
		S3ConnectionID:            s3InputOption.S3ConnectionID.ValueInt64Pointer(),
		Bucket:                    s3InputOption.Bucket.ValueStringPointer(),
		PathPrefix:                model.NewNullableString(s3InputOption.PathPrefix),
		PathMatchPattern:          model.NewNullableString(s3InputOption.PathMatchPattern),
		Region:                    model.NewNullableString(s3InputOption.Region),
		IncrementalLoadingEnabled: model.NewNullableBool(s3InputOption.IncrementalLoadingEnabled),
		IsSkipHeaderLine:          model.NewNullableBool(s3InputOption.IsSkipHeaderLine),
		StopWhenFileNotFound:      model.NewNullableBool(s3InputOption.StopWhenFileNotFound),
		DecompressionType:         model.NewNullableString(s3InputOption.DecompressionType),
		CsvParser:                 s3InputOption.CsvParser.ToCsvParserInput(ctx),
		JsonlParser:               s3InputOption.JsonlParser.ToJsonlParserInput(ctx),
		JsonpathParser:            s3InputOption.JsonpathParser.ToJsonpathParserInput(ctx),
		LtsvParser:                s3InputOption.LtsvParser.ToLtsvParserInput(ctx),
		ExcelParser:               s3InputOption.ExcelParser.ToExcelParserInput(ctx),
		XmlParser:                 s3InputOption.XmlParser.ToXmlParserInput(ctx),
		ParquetParser:             s3InputOption.ParquetParser.ToParquetParserInput(ctx),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
		Decoder:                   s3InputOption.Decoder.ToDecoderInput(),
	}
}
