package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
	jobDefinitionEntities "terraform-provider-trocco/internal/client/entity/job_definition"
)

type SftpInputOption struct {
	SftpConnectionID          int64                                 `json:"sftp_connection_id"`
	PathPrefix                string                                `json:"path_prefix"`
	PathMatchPattern          *string                               `json:"path_match_pattern"`
	IncrementalLoadingEnabled bool                                  `json:"incremental_loading_enabled"`
	LastPath                  *string                               `json:"last_path"`
	StopWhenFileNotFound      bool                                  `json:"stop_when_file_not_found"`
	DecompressionType         string                                `json:"decompression_type"`
	Decoder                   *jobDefinitionEntities.Decoder        `json:"decoder"`
	CustomVariableSettings    *[]entity.CustomVariableSetting       `json:"custom_variable_settings"`
	CsvParser                 *jobDefinitionEntities.CsvParser      `json:"csv_parser"`
	JsonlParser               *jobDefinitionEntities.JsonlParser    `json:"jsonl_parser"`
	JsonpathParser            *jobDefinitionEntities.JsonpathParser `json:"jsonpath_parser"`
	LtsvParser                *jobDefinitionEntities.LtsvParser     `json:"ltsv_parser"`
	ExcelParser               *jobDefinitionEntities.ExcelParser    `json:"excel_parser"`
	XmlParser                 *jobDefinitionEntities.XmlParser      `json:"xml_parser"`
}
