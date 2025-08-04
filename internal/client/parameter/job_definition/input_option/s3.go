package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
	jobDefinitionParameters "terraform-provider-trocco/internal/client/parameter/job_definition"
)

type S3InputOptionInput struct {
	S3ConnectionID            int64                                        `json:"s3_connection_id"`
	Bucket                    string                                       `json:"bucket"`
	PathPrefix                *parameter.NullableString                    `json:"path_prefix,omitempty"`
	PathMatchPattern          *parameter.NullableString                    `json:"path_match_pattern,omitempty"`
	Region                    *parameter.NullableString                    `json:"region,omitempty"`
	IncrementalLoadingEnabled *parameter.NullableBool                      `json:"incremental_loading_enabled,omitempty"`
	IsSkipHeaderLine          *parameter.NullableBool                      `json:"is_skip_header_line,omitempty"`
	StopWhenFileNotFound      *parameter.NullableBool                      `json:"stop_when_file_not_found,omitempty"`
	DecompressionType         *parameter.NullableString                    `json:"decompression_type,omitempty"`
	CsvParser                 *jobDefinitionParameters.CsvParserInput      `json:"csv_parser,omitempty"`
	JsonlParser               *jobDefinitionParameters.JsonlParserInput    `json:"jsonl_parser,omitempty"`
	JsonpathParser            *jobDefinitionParameters.JsonpathParserInput `json:"jsonpath_parser,omitempty"`
	LtsvParser                *jobDefinitionParameters.LtsvParserInput     `json:"ltsv_parser,omitempty"`
	ExcelParser               *jobDefinitionParameters.ExcelParserInput    `json:"excel_parser,omitempty"`
	XmlParser                 *jobDefinitionParameters.XmlParserInput      `json:"xml_parser,omitempty"`
	ParquetParser             *jobDefinitionParameters.ParquetParserInput  `json:"parquet_parser,omitempty"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput      `json:"custom_variable_settings,omitempty"`
	Decoder                   *jobDefinitionParameters.DecoderInput        `json:"decoder,omitempty"`
}

type UpdateS3InputOptionInput struct {
	S3ConnectionID            *int64                                       `json:"s3_connection_id,omitempty"`
	Bucket                    *string                                      `json:"bucket,omitempty"`
	PathPrefix                *parameter.NullableString                    `json:"path_prefix,omitempty"`
	PathMatchPattern          *parameter.NullableString                    `json:"path_match_pattern,omitempty"`
	Region                    *parameter.NullableString                    `json:"region,omitempty"`
	IncrementalLoadingEnabled *parameter.NullableBool                      `json:"incremental_loading_enabled,omitempty"`
	IsSkipHeaderLine          *parameter.NullableBool                      `json:"is_skip_header_line,omitempty"`
	StopWhenFileNotFound      *parameter.NullableBool                      `json:"stop_when_file_not_found,omitempty"`
	DecompressionType         *parameter.NullableString                    `json:"decompression_type,omitempty"`
	CsvParser                 *jobDefinitionParameters.CsvParserInput      `json:"csv_parser,omitempty"`
	JsonlParser               *jobDefinitionParameters.JsonlParserInput    `json:"jsonl_parser,omitempty"`
	JsonpathParser            *jobDefinitionParameters.JsonpathParserInput `json:"jsonpath_parser,omitempty"`
	LtsvParser                *jobDefinitionParameters.LtsvParserInput     `json:"ltsv_parser,omitempty"`
	ExcelParser               *jobDefinitionParameters.ExcelParserInput    `json:"excel_parser,omitempty"`
	XmlParser                 *jobDefinitionParameters.XmlParserInput      `json:"xml_parser,omitempty"`
	ParquetParser             *jobDefinitionParameters.ParquetParserInput  `json:"parquet_parser,omitempty"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput      `json:"custom_variable_settings,omitempty"`
	Decoder                   *jobDefinitionParameters.DecoderInput        `json:"decoder,omitempty"`
}
