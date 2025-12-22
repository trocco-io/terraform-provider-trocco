package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type SftpOutputOptionInput struct {
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	SftpConnectionID       int64                                   `json:"sftp_connection_id"`
	PathPrefix             string                                  `json:"path_prefix"`
	FileExt                string                                  `json:"file_ext"`
	IsMinimumOutputTasks   bool                                    `json:"is_minimum_output_tasks"`
	FormatterType          string                                  `json:"formatter_type"`
	EncoderType            string                                  `json:"encoder_type"`
	CsvFormatter           *SftpOutputOptionCsvFormatterInput      `json:"csv_formatter,omitempty"`
	JsonlFormatter         *SftpOutputOptionJsonlFormatterInput    `json:"jsonl_formatter,omitempty"`
}

type UpdateSftpOutputOptionInput struct {
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	SftpConnectionID       *int64                                  `json:"sftp_connection_id,omitempty"`
	PathPrefix             *string                                 `json:"path_prefix,omitempty"`
	FileExt                *string                                 `json:"file_ext,omitempty"`
	IsMinimumOutputTasks   *bool                                   `json:"is_minimum_output_tasks,omitempty"`
	FormatterType          *string                                 `json:"formatter_type,omitempty"`
	EncoderType            *string                                 `json:"encoder_type,omitempty"`
	CsvFormatter           *SftpOutputOptionCsvFormatterInput      `json:"csv_formatter,omitempty"`
	JsonlFormatter         *SftpOutputOptionJsonlFormatterInput    `json:"jsonl_formatter,omitempty"`
}

type SftpOutputOptionCsvFormatterInput struct {
	Delimiter                           string                                           `json:"delimiter,omitempty"`
	Newline                             string                                           `json:"newline,omitempty"`
	NewlineInField                      string                                           `json:"newline_in_field,omitempty"`
	Charset                             string                                           `json:"charset,omitempty"`
	QuotePolicy                         string                                           `json:"quote_policy,omitempty"`
	Escape                              string                                           `json:"escape,omitempty"`
	HeaderLine                          bool                                             `json:"header_line,omitempty"`
	NullStringEnabled                   bool                                             `json:"null_string_enabled,omitempty"`
	NullString                          *parameter.NullableString                        `json:"null_string,omitempty"`
	DefaultTimeZone                     string                                           `json:"default_time_zone,omitempty"`
	CsvFormatterColumnOptionsAttributes *[]SftpOutputOptionCsvFormatterColumnOptionInput `json:"csv_formatter_column_options_attributes,omitempty"`
}

type SftpOutputOptionCsvFormatterColumnOptionInput struct {
	Name     string                    `json:"name"`
	Format   string                    `json:"format,omitempty"`
	Timezone *parameter.NullableString `json:"timezone,omitempty"`
}

type SftpOutputOptionJsonlFormatterInput struct {
	Encoding   string                    `json:"encoding,omitempty"`
	Newline    string                    `json:"newline,omitempty"`
	DateFormat *parameter.NullableString `json:"date_format,omitempty"`
	Timezone   *parameter.NullableString `json:"timezone,omitempty"`
}
