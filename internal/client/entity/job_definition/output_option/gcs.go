package output_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type GcsOutputOption struct {
	GcsConnectionID        *int64                          `json:"gcs_connection_id"`
	Bucket                 string                          `json:"bucket"`
	PathPrefix             string                          `json:"path_prefix"`
	FileExt                string                          `json:"file_ext"`
	SequenceFormat         *string                         `json:"sequence_format"`
	IsMinimumOutputTasks   bool                            `json:"is_minimum_output_tasks"`
	EncoderType            string                          `json:"encoder_type"`
	Formatter              *GcsOutputOptionFormatter       `json:"formatter"`
	CustomVariableSettings *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type GcsOutputOptionFormatter struct {
	Type           string          `json:"type"`
	CsvFormatter   *CsvFormatter   `json:"csv_formatter,omitempty"`
	JsonlFormatter *JsonlFormatter `json:"jsonl_formatter,omitempty"`
}
