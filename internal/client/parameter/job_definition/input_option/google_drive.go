package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
	jobDefinitionParameters "terraform-provider-trocco/internal/client/parameter/job_definition"
)

type GoogleDriveInputOptionInput struct {
	GoogleDriveConnectionID int64                                        `json:"google_drive_connection_id"`
	FolderID                string                                       `json:"folder_id"`
	FileMatchPattern        string                                       `json:"file_match_pattern"`
	IsSkipHeaderLine        bool                                         `json:"is_skip_header_line"`
	StopWhenFileNotFound    bool                                         `json:"stop_when_file_not_found"`
	DecompressionType       *parameter.NullableString                    `json:"decompression_type,omitempty"`
	CsvParser               *jobDefinitionParameters.CsvParserInput      `json:"csv_parser,omitempty"`
	JsonlParser             *jobDefinitionParameters.JsonlParserInput    `json:"jsonl_parser,omitempty"`
	JsonpathParser          *jobDefinitionParameters.JsonpathParserInput `json:"jsonpath_parser,omitempty"`
	LtsvParser              *jobDefinitionParameters.LtsvParserInput     `json:"ltsv_parser,omitempty"`
	ExcelParser             *jobDefinitionParameters.ExcelParserInput    `json:"excel_parser,omitempty"`
	XmlParser               *jobDefinitionParameters.XmlParserInput      `json:"xml_parser,omitempty"`
	ParquetParser           *jobDefinitionParameters.ParquetParserInput  `json:"parquet_parser,omitempty"`
	CustomVariableSettings  *[]parameter.CustomVariableSettingInput      `json:"custom_variable_settings,omitempty"`
	Decoder                 *jobDefinitionParameters.DecoderInput        `json:"decoder,omitempty"`
}

type UpdateGoogleDriveInputOptionInput struct {
	GoogleDriveConnectionID *int64                                       `json:"google_drive_connection_id,omitempty"`
	FolderID                *string                                      `json:"folder_id,omitempty"`
	FileMatchPattern        *string                                      `json:"file_match_pattern,omitempty"`
	IsSkipHeaderLine        *bool                                        `json:"is_skip_header_line,omitempty"`
	StopWhenFileNotFound    *bool                                        `json:"stop_when_file_not_found,omitempty"`
	DecompressionType       *parameter.NullableString                    `json:"decompression_type,omitempty"`
	CsvParser               *jobDefinitionParameters.CsvParserInput      `json:"csv_parser,omitempty"`
	JsonlParser             *jobDefinitionParameters.JsonlParserInput    `json:"jsonl_parser,omitempty"`
	JsonpathParser          *jobDefinitionParameters.JsonpathParserInput `json:"jsonpath_parser,omitempty"`
	LtsvParser              *jobDefinitionParameters.LtsvParserInput     `json:"ltsv_parser,omitempty"`
	ExcelParser             *jobDefinitionParameters.ExcelParserInput    `json:"excel_parser,omitempty"`
	XmlParser               *jobDefinitionParameters.XmlParserInput      `json:"xml_parser,omitempty"`
	ParquetParser           *jobDefinitionParameters.ParquetParserInput  `json:"parquet_parser,omitempty"`
	CustomVariableSettings  *[]parameter.CustomVariableSettingInput      `json:"custom_variable_settings,omitempty"`
	Decoder                 *jobDefinitionParameters.DecoderInput        `json:"decoder,omitempty"`
}
