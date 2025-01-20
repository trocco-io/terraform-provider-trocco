package input_options

import (
	"terraform-provider-trocco/internal/client/parameters"
	"terraform-provider-trocco/internal/client/parameters/job_definitions"
)

type GcsInputOptionInput struct {
	GcsConnectionID           int64                                    `json:"gcs_connection_id"`
	Bucket                    string                                   `json:"bucket"`
	PathPrefix                string                                   `json:"path_prefix"`
	IncrementalLoadingEnabled bool                                     `json:"incremental_loading_enabled"`
	LastPath                  *parameters.NullableString               `json:"last_path,omitempty"`
	StopWhenFileNotFound      bool                                     `json:"stop_when_file_not_found"`
	DecompressionType         *parameters.NullableString               `json:"decompression_type,omitempty"`
	CsvParser                 *job_definitions.CsvParserInput          `json:"csv_parser,omitempty"`
	JsonlParser               *job_definitions.JsonlParserInput        `json:"jsonl_parser,omitempty"`
	JsonpathParser            *job_definitions.JsonpathParserInput     `json:"jsonpath_parser,omitempty"`
	LtsvParser                *job_definitions.LtsvParserInput         `json:"ltsv_parser,omitempty"`
	ExcelParser               *job_definitions.ExcelParserInput        `json:"excel_parser,omitempty"`
	XmlParser                 *job_definitions.XmlParserInput          `json:"xml_parser,omitempty"`
	ParquetParser             *job_definitions.ParquetParserInput      `json:"parquet_parser,omitempty"`
	CustomVariableSettings    *[]parameters.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	Decoder                   *job_definitions.DecoderInput            `json:"decoder,omitempty"`
}

type UpdateGcsInputOptionInput struct {
	GcsConnectionID           *int64                                   `json:"gcs_connection_id,omitempty"`
	Bucket                    *string                                  `json:"bucket,omitempty"`
	PathPrefix                *string                                  `json:"path_prefix,omitempty"`
	IncrementalLoadingEnabled *bool                                    `json:"incremental_loading_enabled,omitempty"`
	LastPath                  *parameters.NullableString               `json:"last_path,omitempty"`
	StopWhenFileNotFound      *bool                                    `json:"stop_when_file_not_found,omitempty"`
	DecompressionType         *parameters.NullableString               `json:"decompression_type,omitempty"`
	CsvParser                 *job_definitions.CsvParserInput          `json:"csv_parser,omitempty"`
	JsonlParser               *job_definitions.JsonlParserInput        `json:"jsonl_parser,omitempty"`
	JsonpathParser            *job_definitions.JsonpathParserInput     `json:"jsonpath_parser,omitempty"`
	LtsvParser                *job_definitions.LtsvParserInput         `json:"ltsv_parser,omitempty"`
	ExcelParser               *job_definitions.ExcelParserInput        `json:"excel_parser,omitempty"`
	XmlParser                 *job_definitions.XmlParserInput          `json:"xml_parser,omitempty"`
	ParquetParser             *job_definitions.ParquetParserInput      `json:"parquet_parser,omitempty"`
	CustomVariableSettings    *[]parameters.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	Decoder                   *job_definitions.DecoderInput            `json:"decoder,omitempty"`
}
