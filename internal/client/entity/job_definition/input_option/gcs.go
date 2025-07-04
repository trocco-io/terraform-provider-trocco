package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
	jobDefinitions "terraform-provider-trocco/internal/client/entity/job_definition"
)

type GcsInputOption struct {
	GcsConnectionID           int64                           `json:"gcs_connection_id"`
	Bucket                    string                          `json:"bucket"`
	PathPrefix                string                          `json:"path_prefix"`
	IncrementalLoadingEnabled bool                            `json:"incremental_loading_enabled"`
	LastPath                  *string                         `json:"last_path"`
	StopWhenFileNotFound      bool                            `json:"stop_when_file_not_found"`
	DecompressionType         *string                         `json:"decompression_type"`
	CsvParser                 *jobDefinitions.CsvParser       `json:"csv_parser"`
	JsonlParser               *jobDefinitions.JsonlParser     `json:"jsonl_parser"`
	JsonpathParser            *jobDefinitions.JsonpathParser  `json:"jsonpath_parser"`
	LtsvParser                *jobDefinitions.LtsvParser      `json:"ltsv_parser"`
	ExcelParser               *jobDefinitions.ExcelParser     `json:"excel_parser"`
	XmlParser                 *jobDefinitions.XmlParser       `json:"xml_parser"`
	ParquetParser             *jobDefinitions.ParquetParser   `json:"parquet_parser"`
	CustomVariableSettings    *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
	Decoder                   *jobDefinitions.Decoder         `json:"decoder"`
}
