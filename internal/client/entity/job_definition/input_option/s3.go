package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
	jobDefinitions "terraform-provider-trocco/internal/client/entity/job_definition"
)

type S3InputOption struct {
	S3ConnectionID            int64                           `json:"s3_connection_id"`
	Bucket                    string                          `json:"bucket"`
	PathPrefix                string                          `json:"path_prefix"`
	PathMatchPattern          string                          `json:"path_match_pattern"`
	Region                    string                          `json:"region"`
	IncrementalLoadingEnabled bool                            `json:"incremental_loading_enabled"`
	IsSkipHeaderLine          bool                            `json:"is_skip_header_line"`
	StopWhenFileNotFound      bool                            `json:"stop_when_file_not_found"`
	DecompressionType         string                          `json:"decompression_type"`
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
