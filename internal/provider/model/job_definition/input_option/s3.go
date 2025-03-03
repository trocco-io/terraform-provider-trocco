package input_options

import (
	input_options "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	input_options2 "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/input_option/parser"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type S3InputOption struct {
	S3ConnectionID            types.Int64                    `tfsdk:"s3_connection_id"`
	Bucket                    types.String                   `tfsdk:"bucket"`
	PathPrefix                types.String                   `tfsdk:"path_prefix"`
	PathMatchPattern          types.String                   `tfsdk:"path_match_pattern"`
	Region                    types.String                   `tfsdk:"region"`
	IncrementalLoadingEnabled types.Bool                     `tfsdk:"incremental_loading_enabled"`
	IsSkipHeaderLine          types.Bool                     `tfsdk:"is_skip_header_line"`
	StopWhenFileNotFound      types.Bool                     `tfsdk:"stop_when_file_not_found"`
	DecompressionType         types.String                   `tfsdk:"decompression_type"`
	CsvParser                 *parser.CsvParser              `tfsdk:"csv_parser"`
	JsonlParser               *parser.JsonlParser            `tfsdk:"jsonl_parser"`
	JsonpathParser            *parser.JsonpathParser         `tfsdk:"jsonpath_parser"`
	LtsvParser                *parser.LtsvParser             `tfsdk:"ltsv_parser"`
	ExcelParser               *parser.ExcelParser            `tfsdk:"excel_parser"`
	XmlParser                 *parser.XmlParser              `tfsdk:"xml_parser"`
	ParquetParser             *parser.ParquetParser          `tfsdk:"parquet_parser"`
	CustomVariableSettings    *[]model.CustomVariableSetting `tfsdk:"custom_variable_settings"`
	Decoder                   *Decoder                       `tfsdk:"decoder"`
}

func NewS3InputOption(s3InputOption *input_options.S3InputOption) *S3InputOption {
	if s3InputOption == nil {
		return nil
	}
	return &S3InputOption{
		S3ConnectionID:            types.Int64Value(s3InputOption.S3ConnectionID),
		Bucket:                    types.StringValue(s3InputOption.Bucket),
		PathPrefix:                types.StringValue(s3InputOption.PathPrefix),
		PathMatchPattern:          types.StringValue(s3InputOption.PathMatchPattern),
		Region:                    types.StringValue(s3InputOption.Region),
		IncrementalLoadingEnabled: types.BoolValue(s3InputOption.IncrementalLoadingEnabled),
		IsSkipHeaderLine:          types.BoolValue(s3InputOption.IsSkipHeaderLine),
		StopWhenFileNotFound:      types.BoolValue(s3InputOption.StopWhenFileNotFound),
		DecompressionType:         types.StringValue(s3InputOption.DecompressionType),
		CsvParser:                 parser.NewCsvParser(s3InputOption.CsvParser),
		JsonlParser:               parser.NewJsonlParser(s3InputOption.JsonlParser),
		JsonpathParser:            parser.NewJsonPathParser(s3InputOption.JsonpathParser),
		LtsvParser:                parser.NewLtsvParser(s3InputOption.LtsvParser),
		ExcelParser:               parser.NewExcelParser(s3InputOption.ExcelParser),
		XmlParser:                 parser.NewXmlParser(s3InputOption.XmlParser),
		ParquetParser:             parser.NewParquetParser(s3InputOption.ParquetParser),
		CustomVariableSettings:    model.NewCustomVariableSettings(s3InputOption.CustomVariableSettings),
		Decoder:                   NewDecoder(s3InputOption.Decoder),
	}
}

func (s3InputOption *S3InputOption) ToInput() *input_options2.S3InputOptionInput {
	if s3InputOption == nil {
		return nil
	}

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
		CsvParser:                 s3InputOption.CsvParser.ToCsvParserInput(),
		JsonlParser:               s3InputOption.JsonlParser.ToJsonlParserInput(),
		JsonpathParser:            s3InputOption.JsonpathParser.ToJsonpathParserInput(),
		LtsvParser:                s3InputOption.LtsvParser.ToLtsvParserInput(),
		ExcelParser:               s3InputOption.ExcelParser.ToExcelParserInput(),
		XmlParser:                 s3InputOption.XmlParser.ToXmlParserInput(),
		ParquetParser:             s3InputOption.ParquetParser.ToParquetParserInput(),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(s3InputOption.CustomVariableSettings),
		Decoder:                   s3InputOption.Decoder.ToDecoderInput(),
	}
}

func (s3InputOption *S3InputOption) ToUpdateInput() *input_options2.UpdateS3InputOptionInput {
	if s3InputOption == nil {
		return nil
	}

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
		CsvParser:                 s3InputOption.CsvParser.ToCsvParserInput(),
		JsonlParser:               s3InputOption.JsonlParser.ToJsonlParserInput(),
		JsonpathParser:            s3InputOption.JsonpathParser.ToJsonpathParserInput(),
		LtsvParser:                s3InputOption.LtsvParser.ToLtsvParserInput(),
		ExcelParser:               s3InputOption.ExcelParser.ToExcelParserInput(),
		XmlParser:                 s3InputOption.XmlParser.ToXmlParserInput(),
		ParquetParser:             s3InputOption.ParquetParser.ToParquetParserInput(),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(s3InputOption.CustomVariableSettings),
		Decoder:                   s3InputOption.Decoder.ToDecoderInput(),
	}
}
