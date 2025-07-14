package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
	jobDefinitionParameters "terraform-provider-trocco/internal/client/parameter/job_definition"
)

type GcsInputOptionInput struct {
	GcsConnectionID           int64                                        `json:"gcs_connection_id"`
	Bucket                    string                                       `json:"bucket"`
	PathPrefix                string                                       `json:"path_prefix"`
	IncrementalLoadingEnabled bool                                         `json:"incremental_loading_enabled"`
	LastPath                  *parameter.NullableString                    `json:"last_path,omitempty"`
	StopWhenFileNotFound      bool                                         `json:"stop_when_file_not_found"`
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

type UpdateGcsInputOptionInput struct {
	GcsConnectionID           *int64                                       `json:"gcs_connection_id,omitempty"`
	Bucket                    *string                                      `json:"bucket,omitempty"`
	PathPrefix                *string                                      `json:"path_prefix,omitempty"`
	IncrementalLoadingEnabled *bool                                        `json:"incremental_loading_enabled,omitempty"`
	LastPath                  *parameter.NullableString                    `json:"last_path,omitempty"`
	StopWhenFileNotFound      *bool                                        `json:"stop_when_file_not_found,omitempty"`
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
