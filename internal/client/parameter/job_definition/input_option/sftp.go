package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
	jobDefinitionParameters "terraform-provider-trocco/internal/client/parameter/job_definition"
)

type SftpInputOptionInput struct {
	SftpConnectionID          int64                                        `json:"sftp_connection_id"`
	PathPrefix                string                                       `json:"path_prefix"`
	PathMatchPattern          *parameter.NullableString                    `json:"path_match_pattern,omitempty"`
	IncrementalLoadingEnabled bool                                         `json:"incremental_loading_enabled"`
	LastPath                  *parameter.NullableString                    `json:"last_path,omitempty"`
	StopWhenFileNotFound      bool                                         `json:"stop_when_file_not_found"`
	DecompressionType         string                                       `json:"decompression_type"`
	Decoder                   *jobDefinitionParameters.DecoderInput        `json:"decoder,omitempty"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput      `json:"custom_variable_settings,omitempty"`
	CsvParser                 *jobDefinitionParameters.CsvParserInput      `json:"csv_parser,omitempty"`
	JsonlParser               *jobDefinitionParameters.JsonlParserInput    `json:"jsonl_parser,omitempty"`
	JsonpathParser            *jobDefinitionParameters.JsonpathParserInput `json:"jsonpath_parser,omitempty"`
	LtsvParser                *jobDefinitionParameters.LtsvParserInput     `json:"ltsv_parser,omitempty"`
	ExcelParser               *jobDefinitionParameters.ExcelParserInput    `json:"excel_parser,omitempty"`
	XmlParser                 *jobDefinitionParameters.XmlParserInput      `json:"xml_parser,omitempty"`
}

type UpdateSftpInputOptionInput struct {
	SftpConnectionID          *int64                                       `json:"sftp_connection_id,omitempty"`
	PathPrefix                *string                                      `json:"path_prefix,omitempty"`
	PathMatchPattern          *parameter.NullableString                    `json:"path_match_pattern,omitempty"`
	IncrementalLoadingEnabled *bool                                        `json:"incremental_loading_enabled,omitempty"`
	LastPath                  *parameter.NullableString                    `json:"last_path,omitempty"`
	StopWhenFileNotFound      *bool                                        `json:"stop_when_file_not_found,omitempty"`
	DecompressionType         *string                                      `json:"decompression_type,omitempty"`
	Decoder                   *jobDefinitionParameters.DecoderInput        `json:"decoder,omitempty"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput      `json:"custom_variable_settings,omitempty"`
	CsvParser                 *jobDefinitionParameters.CsvParserInput      `json:"csv_parser,omitempty"`
	JsonlParser               *jobDefinitionParameters.JsonlParserInput    `json:"jsonl_parser,omitempty"`
	JsonpathParser            *jobDefinitionParameters.JsonpathParserInput `json:"jsonpath_parser,omitempty"`
	LtsvParser                *jobDefinitionParameters.LtsvParserInput     `json:"ltsv_parser,omitempty"`
	ExcelParser               *jobDefinitionParameters.ExcelParserInput    `json:"excel_parser,omitempty"`
	XmlParser                 *jobDefinitionParameters.XmlParserInput      `json:"xml_parser,omitempty"`
}
