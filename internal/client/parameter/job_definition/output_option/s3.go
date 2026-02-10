package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type S3OutputOptionInput struct {
	S3ConnectionID         int64                                   `json:"s3_connection_id"`
	Bucket                 string                                  `json:"bucket"`
	PathPrefix             string                                  `json:"path_prefix"`
	Region                 string                                  `json:"region"`
	FileExt                string                                  `json:"file_ext"`
	SequenceFormat         string                                  `json:"sequence_format,omitempty"`
	CannedAcl              string                                  `json:"canned_acl"`
	IsMinimumOutputTasks   bool                                    `json:"is_minimum_output_tasks"`
	MultipartUploadEnabled bool                                    `json:"multipart_upload_enabled"`
	FormatterType          string                                  `json:"formatter_type"`
	EncoderType            string                                  `json:"encoder_type"`
	CsvFormatter           *CsvFormatterInput                      `json:"csv_formatter,omitempty"`
	JsonlFormatter         *JsonlFormatterInput                    `json:"jsonl_formatter,omitempty"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateS3OutputOptionInput struct {
	S3ConnectionID         *int64                                  `json:"s3_connection_id,omitempty"`
	Bucket                 *string                                 `json:"bucket,omitempty"`
	PathPrefix             *string                                 `json:"path_prefix,omitempty"`
	Region                 *string                                 `json:"region,omitempty"`
	FileExt                *string                                 `json:"file_ext,omitempty"`
	SequenceFormat         *string                                 `json:"sequence_format,omitempty"`
	CannedAcl              *string                                 `json:"canned_acl,omitempty"`
	IsMinimumOutputTasks   *bool                                   `json:"is_minimum_output_tasks,omitempty"`
	MultipartUploadEnabled *bool                                   `json:"multipart_upload_enabled,omitempty"`
	FormatterType          *string                                 `json:"formatter_type,omitempty"`
	EncoderType            *string                                 `json:"encoder_type,omitempty"`
	CsvFormatter           *CsvFormatterInput                      `json:"csv_formatter,omitempty"`
	JsonlFormatter         *JsonlFormatterInput                    `json:"jsonl_formatter,omitempty"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type CsvFormatterInput struct {
	Delimiter                           string                           `json:"delimiter,omitempty"`
	Newline                             string                           `json:"newline,omitempty"`
	NewlineInField                      string                           `json:"newline_in_field,omitempty"`
	Charset                             string                           `json:"charset,omitempty"`
	QuotePolicy                         string                           `json:"quote_policy,omitempty"`
	Escape                              string                           `json:"escape,omitempty"`
	HeaderLine                          bool                             `json:"header_line,omitempty"`
	NullStringEnabled                   bool                             `json:"null_string_enabled,omitempty"`
	NullString                          *string                          `json:"null_string,omitempty"`
	DefaultTimeZone                     string                           `json:"default_time_zone,omitempty"`
	CsvFormatterColumnOptionsAttributes *[]CsvFormatterColumnOptionInput `json:"csv_formatter_column_options_attributes,omitempty"`
}

type CsvFormatterColumnOptionInput struct {
	Name     string  `json:"name"`
	Format   string  `json:"format,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

type JsonlFormatterInput struct {
	Encoding   string  `json:"encoding,omitempty"`
	Newline    string  `json:"newline,omitempty"`
	DateFormat *string `json:"date_format,omitempty"`
	Timezone   *string `json:"timezone,omitempty"`
}
