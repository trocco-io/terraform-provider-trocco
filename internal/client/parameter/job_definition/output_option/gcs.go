package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type GcsOutputOptionInput struct {
	GcsConnectionID        int64                                   `json:"gcs_connection_id"`
	Bucket                 string                                  `json:"bucket"`
	PathPrefix             string                                  `json:"path_prefix"`
	FileExt                string                                  `json:"file_ext"`
	SequenceFormat         *parameter.NullableString               `json:"sequence_format,omitempty"`
	IsMinimumOutputTasks   bool                                    `json:"is_minimum_output_tasks"`
	FormatterType          string                                  `json:"formatter_type"`
	EncoderType            string                                  `json:"encoder_type"`
	CsvFormatter           *GcsOutputOptionCsvFormatterInput       `json:"csv_formatter,omitempty"`
	JsonlFormatter         *GcsOutputOptionJsonlFormatterInput     `json:"jsonl_formatter,omitempty"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateGcsOutputOptionInput struct {
	GcsConnectionID        *int64                                  `json:"gcs_connection_id,omitempty"`
	Bucket                 *string                                 `json:"bucket,omitempty"`
	PathPrefix             *string                                 `json:"path_prefix,omitempty"`
	FileExt                *string                                 `json:"file_ext,omitempty"`
	SequenceFormat         *parameter.NullableString               `json:"sequence_format,omitempty"`
	IsMinimumOutputTasks   *bool                                   `json:"is_minimum_output_tasks,omitempty"`
	FormatterType          *string                                 `json:"formatter_type,omitempty"`
	EncoderType            *string                                 `json:"encoder_type,omitempty"`
	CsvFormatter           *GcsOutputOptionCsvFormatterInput       `json:"csv_formatter,omitempty"`
	JsonlFormatter         *GcsOutputOptionJsonlFormatterInput     `json:"jsonl_formatter,omitempty"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type GcsOutputOptionCsvFormatterInput struct {
	Delimiter                           string                                          `json:"delimiter,omitempty"`
	Newline                             string                                          `json:"newline,omitempty"`
	NewlineInField                      string                                          `json:"newline_in_field,omitempty"`
	Charset                             string                                          `json:"charset,omitempty"`
	QuotePolicy                         string                                          `json:"quote_policy,omitempty"`
	Escape                              string                                          `json:"escape,omitempty"`
	HeaderLine                          bool                                            `json:"header_line,omitempty"`
	NullStringEnabled                   bool                                            `json:"null_string_enabled,omitempty"`
	NullString                          *string                                         `json:"null_string,omitempty"`
	DefaultTimeZone                     string                                          `json:"default_time_zone,omitempty"`
	CsvFormatterColumnOptionsAttributes *[]GcsOutputOptionCsvFormatterColumnOptionInput `json:"csv_formatter_column_options_attributes,omitempty"`
}

type GcsOutputOptionCsvFormatterColumnOptionInput struct {
	Name     string  `json:"name"`
	Format   string  `json:"format,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

type GcsOutputOptionJsonlFormatterInput struct {
	Encoding   string  `json:"encoding,omitempty"`
	Newline    string  `json:"newline,omitempty"`
	DateFormat *string `json:"date_format,omitempty"`
	Timezone   *string `json:"timezone,omitempty"`
}
