package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
	job_definitions "terraform-provider-trocco/internal/client/entity/job_definition"
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
	CsvParser                 *job_definitions.CsvParser      `json:"csv_parser"`
	JsonlParser               *job_definitions.JsonlParser    `json:"jsonl_parser"`
	JsonpathParser            *job_definitions.JsonpathParser `json:"jsonpath_parser"`
	LtsvParser                *job_definitions.LtsvParser     `json:"ltsv_parser"`
	ExcelParser               *job_definitions.ExcelParser    `json:"excel_parser"`
	XmlParser                 *job_definitions.XmlParser      `json:"xml_parser"`
	ParquetParser             *job_definitions.ParquetParser  `json:"parquet_parser"`
	CustomVariableSettings    *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
	Decoder                   *job_definitions.Decoder        `json:"decoder"`
}
