package input_options

import (
	"terraform-provider-trocco/internal/client/entities"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
)

type GcsInputOption struct {
	GcsConnectionID           int64                             `json:"gcs_connection_id"`
	Bucket                    string                            `json:"bucket"`
	PathPrefix                *string                           `json:"path_prefix"`
	IncrementalLoadingEnabled bool                              `json:"incremental_loading_enabled"`
	LastPath                  *string                           `json:"last_path"`
	StopWhenFileNotFound      bool                              `json:"stop_when_file_not_found"`
	DecompressionType         *string                           `json:"decompression_type"`
	CsvParsers                *job_definitions.CsvParser        `json:"csv_parsers"`
	JsonlParsers              *job_definitions.JsonlParser      `json:"jsonl_parsers"`
	JsonpathParsers           *job_definitions.JsonpathParser   `json:"jsonpath_parsers"`
	LtsvParsers               *job_definitions.LtsvParser       `json:"ltsv_parsers"`
	ExcelParsers              *job_definitions.ExcelParser      `json:"excel_parsers"`
	XmlParsers                *job_definitions.XmlParser        `json:"xml_parsers"`
	ParquetParsers            *job_definitions.ParquetParser    `json:"parquet_parsers"`
	CustomVariableSettings    *[]entities.CustomVariableSetting `json:"custom_variable_settings"`
	Decoder                   *job_definitions.Decoder          `json:"decoder"`
}
