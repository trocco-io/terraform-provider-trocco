package output_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type SftpOutputOption struct {
	CustomVariableSettings *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
	SftpConnectionID       *int64                          `json:"sftp_connection_id"`
	PathPrefix             string                          `json:"path_prefix"`
	FileExt                string                          `json:"file_ext"`
	IsMinimumOutputTasks   bool                            `json:"is_minimum_output_tasks"`
	FormatterType          string                          `json:"formatter_type"`
	EncoderType            string                          `json:"encoder_type"`
	CsvFormatter           *SftpOutputOptionCsvFormatter   `json:"csv_formatter"`
	JsonlFormatter         *SftpOutputOptionJsonlFormatter `json:"jsonl_formatter"`
	Formatter              *SftpOutputOptionFormatter      `json:"formatter"`
}

type SftpOutputOptionCsvFormatter struct {
	Delimiter                           string                                      `json:"delimiter"`
	Newline                             string                                      `json:"newline"`
	NewlineInField                      string                                      `json:"newline_in_field"`
	Charset                             string                                      `json:"charset"`
	QuotePolicy                         string                                      `json:"quote_policy"`
	Escape                              string                                      `json:"escape"`
	HeaderLine                          bool                                        `json:"header_line"`
	NullStringEnabled                   bool                                        `json:"null_string_enabled"`
	NullString                          *string                                     `json:"null_string"`
	DefaultTimeZone                     string                                      `json:"default_time_zone"`
	CsvFormatterColumnOptionsAttributes *[]SftpOutputOptionCsvFormatterColumnOption `json:"csv_formatter_column_options_attributes"`
}

type SftpOutputOptionCsvFormatterColumnOption struct {
	Name     string  `json:"name"`
	Format   string  `json:"format"`
	Timezone *string `json:"timezone"`
}

type SftpOutputOptionJsonlFormatter struct {
	Encoding   string  `json:"encoding"`
	Newline    string  `json:"newline"`
	DateFormat *string `json:"date_format"`
	Timezone   *string `json:"timezone"`
}

type SftpOutputOptionFormatter struct {
	Type           string                          `json:"type"`
	CsvFormatter   *SftpOutputOptionCsvFormatter   `json:"csv_formatter,omitempty"`
	JsonlFormatter *SftpOutputOptionJsonlFormatter `json:"jsonl_formatter,omitempty"`

	// Legacy fields for CSV formatter (kept for backward compatibility)
	Delimiter         *string                                  `json:"delimiter,omitempty"`
	Newline           *string                                  `json:"newline,omitempty"`
	NewlineInField    *string                                  `json:"newline_in_field,omitempty"`
	Charset           *string                                  `json:"charset,omitempty"`
	QuotePolicy       *string                                  `json:"quote_policy,omitempty"`
	Escape            *string                                  `json:"escape,omitempty"`
	HeaderLine        *bool                                    `json:"header_line,omitempty"`
	NullStringEnabled *bool                                    `json:"null_string_enabled,omitempty"`
	NullString        *string                                  `json:"null_string,omitempty"`
	DefaultTimeZone   *string                                  `json:"default_time_zone,omitempty"`
	ColumnOptions     *[]SftpOutputOptionFormatterColumnOption `json:"column_options,omitempty"`

	// Legacy fields for JSONL formatter (kept for backward compatibility)
	Encoding   *string `json:"encoding,omitempty"`
	DateFormat *string `json:"date_format,omitempty"`
	Timezone   *string `json:"timezone,omitempty"`
}

type SftpOutputOptionFormatterColumnOption struct {
	Name     *string `json:"name,omitempty"`
	Format   *string `json:"format,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}
