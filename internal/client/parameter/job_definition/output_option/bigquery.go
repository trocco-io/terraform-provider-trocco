package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type BigQueryOutputOptionInput struct {
	CustomVariableSettings               *[]parameter.CustomVariableSettingInput  `json:"custom_variable_settings,omitempty"`
	Dataset                              string                                   `json:"dataset"`
	Table                                string                                   `json:"table"`
	AutoCreateDataset                    bool                                     `json:"auto_create_dataset"`
	OpenTimeoutSec                       int64                                    `json:"open_timeout_sec"`
	TimeoutSec                           int64                                    `json:"timeout_sec"`
	SendTimeoutSec                       int64                                    `json:"send_timeout_sec"`
	ReadTimeoutSec                       int64                                    `json:"read_timeout_sec"`
	Retries                              int64                                    `json:"retries"`
	Mode                                 string                                   `json:"mode"`
	PartitioningType                     *parameter.NullableString                `json:"partitioning_type,omitempty"`
	TimePartitioningType                 *parameter.NullableString                `json:"time_partitioning_type,omitempty"`
	TimePartitioningField                *parameter.NullableString                `json:"time_partitioning_field,omitempty"`
	TimePartitioningExpirationMs         *parameter.NullableInt64                 `json:"time_partitioning_expiration_ms,omitempty"`
	Location                             string                                   `json:"location,omitempty"`
	TemplateTable                        *parameter.NullableString                `json:"template_table,omitempty"`
	BigQueryConnectionID                 int64                                    `json:"bigquery_connection_id"`
	BigQueryOutputOptionColumnOptions    *[]BigQueryOutputOptionColumnOptionInput `json:"bigquery_output_option_column_options,omitempty"`
	BigQueryOutputOptionClusteringFields []string                                 `json:"bigquery_output_option_clustering_fields,omitempty"`
	BigQueryOutputOptionMergeKeys        []string                                 `json:"bigquery_output_option_merge_keys,omitempty"`
}

type UpdateBigQueryOutputOptionInput struct {
	CustomVariableSettings               *[]parameter.CustomVariableSettingInput  `json:"custom_variable_settings,omitempty"`
	Dataset                              *string                                  `json:"dataset,omitempty"`
	Table                                *string                                  `json:"table,omitempty"`
	AutoCreateDataset                    *bool                                    `json:"auto_create_dataset,omitempty"`
	OpenTimeoutSec                       *int64                                   `json:"open_timeout_sec,omitempty"`
	TimeoutSec                           *int64                                   `json:"timeout_sec,omitempty"`
	SendTimeoutSec                       *int64                                   `json:"send_timeout_sec,omitempty"`
	ReadTimeoutSec                       *int64                                   `json:"read_timeout_sec,omitempty"`
	Retries                              *int64                                   `json:"retries,omitempty"`
	Mode                                 *string                                  `json:"mode,omitempty"`
	PartitioningType                     *parameter.NullableString                `json:"partitioning_type,omitempty"`
	TimePartitioningType                 *parameter.NullableString                `json:"time_partitioning_type,omitempty"`
	TimePartitioningField                *parameter.NullableString                `json:"time_partitioning_field,omitempty"`
	TimePartitioningExpirationMs         *parameter.NullableInt64                 `json:"time_partitioning_expiration_ms,omitempty"`
	Location                             *string                                  `json:"location,omitempty"`
	TemplateTable                        *parameter.NullableString                `json:"template_table,omitempty"`
	BigQueryConnectionID                 *int64                                   `json:"bigquery_connection_id,omitempty"`
	BigQueryOutputOptionColumnOptions    *[]BigQueryOutputOptionColumnOptionInput `json:"bigquery_output_option_column_options,omitempty"`
	BigQueryOutputOptionClusteringFields *[]string                                `json:"bigquery_output_option_clustering_fields,omitempty"`
	BigQueryOutputOptionMergeKeys        *[]string                                `json:"bigquery_output_option_merge_keys,omitempty"`
}

type BigQueryOutputOptionColumnOptionInput struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Mode            string  `json:"mode"`
	TimestampFormat *string `json:"timestamp_format,omitempty"`
	Timezone        *string `json:"timezone,omitempty"`
	Description     *string `json:"description,omitempty"`
}
