package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
	jobDefinitionEntities "terraform-provider-trocco/internal/client/entity/job_definition"
)

type GoogleDriveInputOption struct {
	GoogleDriveConnectionID int64                                 `json:"google_drive_connection_id"`
	FolderID                string                                `json:"folder_id"`
	FileMatchPattern        string                                `json:"file_match_pattern"`
	IsSkipHeaderLine        bool                                  `json:"is_skip_header_line"`
	StopWhenFileNotFound    bool                                  `json:"stop_when_file_not_found"`
	DecompressionType       *string                               `json:"decompression_type"`
	CsvParser               *jobDefinitionEntities.CsvParser      `json:"csv_parser"`
	JsonlParser             *jobDefinitionEntities.JsonlParser    `json:"jsonl_parser"`
	JsonpathParser          *jobDefinitionEntities.JsonpathParser `json:"jsonpath_parser"`
	LtsvParser              *jobDefinitionEntities.LtsvParser     `json:"ltsv_parser"`
	ExcelParser             *jobDefinitionEntities.ExcelParser    `json:"excel_parser"`
	XmlParser               *jobDefinitionEntities.XmlParser      `json:"xml_parser"`
	CustomVariableSettings  *[]entity.CustomVariableSetting       `json:"custom_variable_settings"`
	Decoder                 *jobDefinitionEntities.Decoder        `json:"decoder"`
}
