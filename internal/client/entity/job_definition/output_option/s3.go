package output_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type S3OutputOption struct {
	S3ConnectionID         int64                           `json:"s3_connection_id"`
	Bucket                 string                          `json:"bucket"`
	PathPrefix             string                          `json:"path_prefix"`
	Region                 string                          `json:"region"`
	FileExt                string                          `json:"file_ext"`
	SequenceFormat         string                          `json:"sequence_format"`
	CannedAcl              string                          `json:"canned_acl"`
	IsMinimumOutputTasks   bool                            `json:"is_minimum_output_tasks"`
	MultipartUploadEnabled bool                            `json:"multipart_upload_enabled"`
	EncoderType            *string                         `json:"encoder_type"`
	Formatter              *S3Formatter                    `json:"formatter"`
	CustomVariableSettings *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type S3Formatter struct {
	Type           string          `json:"type"`
	CsvFormatter   *CsvFormatter   `json:"csv_formatter"`
	JsonlFormatter *JsonlFormatter `json:"jsonl_formatter"`
}

type CsvFormatter struct {
	Delimiter                           string                      `json:"delimiter"`
	Escape                              string                      `json:"escape"`
	HeaderLine                          bool                        `json:"header_line"`
	Charset                             string                      `json:"charset"`
	QuotePolicy                         string                      `json:"quote_policy"`
	Newline                             string                      `json:"newline"`
	NewlineInField                      string                      `json:"newline_in_field"`
	NullStringEnabled                   bool                        `json:"null_string_enabled"`
	NullString                          *string                     `json:"null_string"`
	DefaultTimeZone                     string                      `json:"default_time_zone"`
	CsvFormatterColumnOptionsAttributes *[]CsvFormatterColumnOption `json:"csv_formatter_column_options_attributes"`
}

type CsvFormatterColumnOption struct {
	Name     string  `json:"name"`
	Format   string  `json:"format"`
	Timezone *string `json:"timezone"`
}

type JsonlFormatter struct {
	Encoding   string  `json:"encoding"`
	Newline    string  `json:"newline"`
	DateFormat *string `json:"date_format"`
	Timezone   *string `json:"timezone"`
}
