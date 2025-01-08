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
	CsvParser                 *job_definitions.CsvParser        `json:"csv_parser"`
	JsonlParser               *job_definitions.JsonlParser      `json:"jsonl_parser"`
	JsonpathParser            *job_definitions.JsonpathParser   `json:"jsonpath_parser"`
	LtsvParser                *job_definitions.LtsvParser       `json:"ltsv_parser"`
	ExcelParser               *job_definitions.ExcelParser      `json:"excel_parser"`
	XmlParser                 *job_definitions.XmlParser        `json:"xml_parser"`
	ParquetParser             *job_definitions.ParquetParser    `json:"parquet_parser"`
	CustomVariableSettings    *[]entities.CustomVariableSetting `json:"custom_variable_settings"`
	Decoder                   *job_definitions.Decoder          `json:"decoder"`
}
