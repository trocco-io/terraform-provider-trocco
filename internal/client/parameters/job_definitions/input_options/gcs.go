package input_options

import (
	"terraform-provider-trocco/internal/client/parameters"
	"terraform-provider-trocco/internal/client/parameters/job_definitions"
)

type GcsInputOptionInput struct {
	GcsConnectionID           int64                                    `json:"gcs_connection_id"`
	Bucket                    string                                   `json:"bucket"`
	PathPrefix                *string                                  `json:"path_prefix,omitempty"`
	IncrementalLoadingEnabled bool                                     `json:"incremental_loading_enabled"`
	LastPath                  *string                                  `json:"last_path,omitempty"`
	StopWhenFileNotFound      bool                                     `json:"stop_when_file_not_found"`
	DecompressionType         *string                                  `json:"decompression_type,omitempty"`
	CsvParsers                *job_definitions.CsvParserInput          `json:"csv_parsers,omitempty"`
	JsonlParsers              *job_definitions.JsonlParserInput        `json:"jsonl_parsers,omitempty"`
	JsonpathParsers           *job_definitions.JsonpathParserInput     `json:"jsonpath_parsers,omitempty"`
	LtsvParsers               *job_definitions.LtsvParserInput         `json:"ltsv_parsers,omitempty"`
	ExcelParsers              *job_definitions.ExcelParserInput        `json:"excel_parsers,omitempty"`
	XmlParsers                *job_definitions.XmlParserInput          `json:"xml_parsers,omitempty"`
	ParquetParsers            *job_definitions.ParquetParserInput      `json:"parquet_parsers,omitempty"`
	CustomVariableSettings    *[]parameters.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	Decoder                   *job_definitions.DecoderInput            `json:"decoder,omitempty"`
}

type UpdateGcsInputOptionInput struct {
	GcsConnectionID           *int64                                   `json:"gcs_connection_id,omitempty"`
	Bucket                    *string                                  `json:"bucket,omitempty"`
	PathPrefix                *string                                  `json:"path_prefix,omitempty"`
	IncrementalLoadingEnabled *bool                                    `json:"incremental_loading_enabled,omitempty"`
	LastPath                  *string                                  `json:"last_path,omitempty"`
	StopWhenFileNotFound      *bool                                    `json:"stop_when_file_not_found,omitempty"`
	DecompressionType         *string                                  `json:"decompression_type,omitempty"`
	CsvParsers                *job_definitions.CsvParserInput          `json:"csv_parsers,omitempty"`
	JsonlParsers              *job_definitions.JsonlParserInput        `json:"jsonl_parsers,omitempty"`
	JsonpathParsers           *job_definitions.JsonpathParserInput     `json:"jsonpath_parsers,omitempty"`
	LtsvParsers               *job_definitions.LtsvParserInput         `json:"ltsv_parsers,omitempty"`
	ExcelParsers              *job_definitions.ExcelParserInput        `json:"excel_parsers,omitempty"`
	XmlParsers                *job_definitions.XmlParserInput          `json:"xml_parsers,omitempty"`
	ParquetParsers            *job_definitions.ParquetParserInput      `json:"parquet_parsers,omitempty"`
	CustomVariableSettings    *[]parameters.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	Decoder                   *job_definitions.DecoderInput            `json:"decoder,omitempty"`
}
