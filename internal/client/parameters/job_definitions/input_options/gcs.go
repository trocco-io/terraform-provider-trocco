package input_options

import (
	"terraform-provider-trocco/internal/client/parameters"
	"terraform-provider-trocco/internal/client/parameters/job_definitions"
)

type GcsInputOption struct {
	GcsConnectionID           int64                               `json:"gcs_connection_id"`
	Bucket                    string                              `json:"bucket"`
	PathPrefix                *string                             `json:"path_prefix"`
	IncrementalLoadingEnabled bool                                `json:"incremental_loading_enabled"`
	LastPath                  *string                             `json:"last_path"`
	StopWhenFileNotFound      bool                                `json:"stop_when_file_not_found"`
	DecompressionType         *string                             `json:"decompression_type"`
	CsvParsers                *job_definitions.CsvParser          `json:"csv_parsers"`
	JsonlParsers              *job_definitions.JsonlParser        `json:"jsonl_parsers"`
	JsonpathParsers           *job_definitions.JsonpathParser     `json:"jsonpath_parsers"`
	LtsvParsers               *job_definitions.LtsvParser         `json:"ltsv_parsers"`
	ExcelParsers              *job_definitions.ExcelParser        `json:"excel_parsers"`
	XmlParsers                *job_definitions.XmlParser          `json:"xml_parsers"`
	ParquetParsers            *job_definitions.ParquetParser      `json:"parquet_parsers"`
	CustomVariableSettings    *[]parameters.CustomVariableSetting `json:"custom_variable_settings"`
	Decoder                   *job_definitions.Decoder            `json:"decoder"`
}

type GcsInputOptionInput struct {
	GcsConnectionID           int64                                `json:"gcs_connection_id"`
	Bucket                    string                               `json:"bucket"`
	PathPrefix                *string                              `json:"path_prefix,omitempty"`
	IncrementalLoadingEnabled bool                                 `json:"incremental_loading_enabled"`
	LastPath                  *string                              `json:"last_path,omitempty"`
	StopWhenFileNotFound      bool                                 `json:"stop_when_file_not_found"`
	DecompressionType         *string                              `json:"decompression_type,omitempty"`
	CsvParsers                *job_definitions.CsvParser           `json:"csv_parsers,omitempty"`
	JsonlParsers              *job_definitions.JsonlParserInput    `json:"jsonl_parsers,omitempty"`
	JsonpathParsers           *job_definitions.JsonpathParserInput `json:"jsonpath_parsers,omitempty"`
	LtsvParsers               *job_definitions.LtsvParserInput     `json:"ltsv_parsers,omitempty"`
	ExcelParsers              *job_definitions.ExcelParserInput    `json:"excel_parsers,omitempty"`
	XmlParsers                *job_definitions.XmlParserInput      `json:"xml_parsers,omitempty"`
	ParquetParsers            *job_definitions.ParquetParserInput  `json:"parquet_parsers,omitempty"`
	CustomVariableSettings    *[]parameters.CustomVariableSetting  `json:"custom_variable_settings,omitempty"`
	Decoder                   *job_definitions.Decoder             `json:"decoder,omitempty"`
}

type UpdateGcsInputOptionInput struct {
	GcsConnectionID           *int64                               `json:"gcs_connection_id,omitempty"`
	Bucket                    *string                              `json:"bucket,omitempty"`
	PathPrefix                *string                              `json:"path_prefix,omitempty"`
	IncrementalLoadingEnabled *bool                                `json:"incremental_loading_enabled,omitempty"`
	LastPath                  *string                              `json:"last_path,omitempty"`
	StopWhenFileNotFound      *bool                                `json:"stop_when_file_not_found,omitempty"`
	DecompressionType         *string                              `json:"decompression_type,omitempty"`
	CsvParsers                *job_definitions.CsvParserInput      `json:"csv_parsers,omitempty"`
	JsonlParsers              *job_definitions.JsonlParserInput    `json:"jsonl_parsers,omitempty"`
	JsonpathParsers           *job_definitions.JsonpathParserInput `json:"jsonpath_parsers,omitempty"`
	LtsvParsers               *job_definitions.LtsvParserInput     `json:"ltsv_parsers,omitempty"`
	ExcelParsers              *job_definitions.ExcelParserInput    `json:"excel_parsers,omitempty"`
	XmlParsers                *job_definitions.XmlParserInput      `json:"xml_parsers,omitempty"`
	ParquetParsers            *job_definitions.ParquetParserInput  `json:"parquet_parsers,omitempty"`
	CustomVariableSettings    *[]parameters.CustomVariableSetting  `json:"custom_variable_settings,omitempty"`
	Decoder                   *job_definitions.Decoder             `json:"decoder,omitempty"`
}
