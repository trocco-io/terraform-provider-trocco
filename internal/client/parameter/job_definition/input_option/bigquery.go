package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
	jobDefinitionParameters "terraform-provider-trocco/internal/client/parameter/job_definition"
)

type BigqueryInputOptionInput struct {
	BigqueryConnectionID   int64                                   `json:"bigquery_connection_id"`
	GcsUri                 string                                  `json:"gcs_uri"`
	GcsUriFormat           *parameter.NullableString               `json:"gcs_uri_format,omitempty"`
	Query                  string                                  `json:"query"`
	TempDataset            string                                  `json:"temp_dataset"`
	IsStandardSQL          *bool                                   `json:"is_standard_sql,omitempty"`
	CleanupGcsFiles        *bool                                   `json:"cleanup_gcs_files,omitempty"`
	FileFormat             *parameter.NullableString               `json:"file_format,omitempty"`
	Location               *parameter.NullableString               `json:"location,omitempty"`
	Cache                  *bool                                   `json:"cache,omitempty"`
	BigqueryJobWaitSecond  *int64                                  `json:"bigquery_job_wait_second,omitempty"`
	Columns                []BigqueryColumn                        `json:"columns,omitempty"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	Decoder                *jobDefinitionParameters.DecoderInput   `json:"decoder,omitempty"`
}

type UpdateBigqueryInputOptionInput struct {
	BigqueryConnectionID   *int64                                  `json:"bigquery_connection_id,omitempty"`
	GcsUri                 *parameter.NullableString               `json:"gcs_uri,omitempty"`
	GcsUriFormat           *parameter.NullableString               `json:"gcs_uri_format,omitempty"`
	Query                  *parameter.NullableString               `json:"query,omitempty"`
	TempDataset            *parameter.NullableString               `json:"temp_dataset,omitempty"`
	IsStandardSQL          *bool                                   `json:"is_standard_sql,omitempty"`
	CleanupGcsFiles        *bool                                   `json:"cleanup_gcs_files,omitempty"`
	FileFormat             *parameter.NullableString               `json:"file_format,omitempty"`
	Location               *parameter.NullableString               `json:"location,omitempty"`
	Cache                  *bool                                   `json:"cache,omitempty"`
	BigqueryJobWaitSecond  *int64                                  `json:"bigquery_job_wait_second,omitempty"`
	Columns                []BigqueryColumn                        `json:"columns,omitempty"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	Decoder                *jobDefinitionParameters.DecoderInput   `json:"decoder,omitempty"`
}

type BigqueryColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
