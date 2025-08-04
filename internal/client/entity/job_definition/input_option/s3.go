package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
	jobDefinitionEntities "terraform-provider-trocco/internal/client/entity/job_definition"
)

type S3InputOption struct {
	S3ConnectionID            int64                                 `json:"s3_connection_id"`
	Bucket                    string                                `json:"bucket"`
	PathPrefix                string                                `json:"path_prefix"`
	PathMatchPattern          string                                `json:"path_match_pattern"`
	Region                    string                                `json:"region"`
	IncrementalLoadingEnabled bool                                  `json:"incremental_loading_enabled"`
	IsSkipHeaderLine          bool                                  `json:"is_skip_header_line"`
	StopWhenFileNotFound      bool                                  `json:"stop_when_file_not_found"`
	DecompressionType         string                                `json:"decompression_type"`
	CsvParser                 *jobDefinitionEntities.CsvParser      `json:"csv_parser"`
	JsonlParser               *jobDefinitionEntities.JsonlParser    `json:"jsonl_parser"`
	JsonpathParser            *jobDefinitionEntities.JsonpathParser `json:"jsonpath_parser"`
	LtsvParser                *jobDefinitionEntities.LtsvParser     `json:"ltsv_parser"`
	ExcelParser               *jobDefinitionEntities.ExcelParser    `json:"excel_parser"`
	XmlParser                 *jobDefinitionEntities.XmlParser      `json:"xml_parser"`
	ParquetParser             *jobDefinitionEntities.ParquetParser  `json:"parquet_parser"`
	CustomVariableSettings    *[]entity.CustomVariableSetting       `json:"custom_variable_settings"`
	Decoder                   *jobDefinitionEntities.Decoder        `json:"decoder"`
}
