package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type BigqueryInputOption struct {
	BigqueryConnectionID  int64   `json:"bigquery_connection_id"`
	GcsUri                string  `json:"gcs_uri"`
	GcsUriFormat          *string `json:"gcs_uri_format"`
	Query                 string  `json:"query"`
	TempDataset           string  `json:"temp_dataset"`
	IsStandardSQL         *bool   `json:"is_standard_sql"`
	CleanupGcsFiles       *bool   `json:"cleanup_gcs_files"`
	FileFormat            *string `json:"file_format"`
	Location              *string `json:"location"`
	Cache                 *bool   `json:"cache"`
	BigqueryJobWaitSecond *int64  `json:"bigquery_job_wait_second"`

	Columns                []BigqueryColumn                `json:"columns"`
	CustomVariableSettings *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type BigqueryColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
