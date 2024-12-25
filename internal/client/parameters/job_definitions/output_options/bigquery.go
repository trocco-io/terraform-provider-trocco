package output_options

import (
	"terraform-provider-trocco/internal/client/parameters"
)

type BigQueryOutputOptionInput struct {
	CustomVariableSettings                 *[]parameters.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	Dataset                                string                                   `json:"dataset"`
	Table                                  string                                   `json:"table"`
	AutoCreateDataset                      bool                                     `json:"auto_create_dataset"`
	AutoCreateTable                        bool                                     `json:"auto_create_table"`
	OpenTimeoutSec                         int64                                    `json:"open_timeout_sec"`
	TimeoutSec                             int64                                    `json:"timeout_sec"`
	SendTimeoutSec                         int64                                    `json:"send_timeout_sec"`
	ReadTimeoutSec                         int64                                    `json:"read_timeout_sec"`
	Retries                                int64                                    `json:"retries"`
	Mode                                   string                                   `json:"mode"`
	PartitioningType                       *string                                  `json:"partitioning_type,omitempty"`
	TimePartitioningType                   *string                                  `json:"time_partitioning_type,omitempty"`
	TimePartitioningField                  *string                                  `json:"time_partitioning_field,omitempty"`
	TimePartitioningExpirationMs           *int64                                   `json:"time_partitioning_expiration_ms,omitempty"`
	TimePartitioningRequirePartitionFilter *bool                                    `json:"time_partitioning_require_partition_filter,omitempty"`
	Location                               *string                                  `json:"location,omitempty"`
	TemplateTable                          *string                                  `json:"template_table,omitempty"`
	BigQueryConnectionID                   int64                                    `json:"bigquery_connection_id"`
	BeforeLoad                             string                                   `json:"before_load"`
	BigQueryOutputOptionColumnOptions      *[]BigQueryOutputOptionColumnOptionInput `json:"bigquery_output_option_column_options,omitempty"`
	BigQueryOutputOptionClusteringFields   *[]string                                `json:"bigquery_output_option_clustering_fields,omitempty"`
	BigQueryOutputOptionMergeKeys          *[]string                                `json:"bigquery_output_option_merge_keys,omitempty"`
}

type UpdateBigQueryOutputOptionInput struct {
	CustomVariableSettings                 *[]parameters.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	Dataset                                *string                                  `json:"dataset,omitempty"`
	Table                                  *string                                  `json:"table,omitempty"`
	AutoCreateDataset                      *bool                                    `json:"auto_create_dataset,omitempty"`
	AutoCreateTable                        *bool                                    `json:"auto_create_table,omitempty"`
	OpenTimeoutSec                         *int64                                   `json:"open_timeout_sec,omitempty"`
	TimeoutSec                             *int64                                   `json:"timeout_sec,omitempty"`
	SendTimeoutSec                         *int64                                   `json:"send_timeout_sec,omitempty"`
	ReadTimeoutSec                         *int64                                   `json:"read_timeout_sec,omitempty"`
	Retries                                *int64                                   `json:"retries,omitempty"`
	Mode                                   *string                                  `json:"mode,omitempty"`
	PartitioningType                       *string                                  `json:"partitioning_type,omitempty"`
	TimePartitioningType                   *string                                  `json:"time_partitioning_type,omitempty"`
	TimePartitioningField                  *string                                  `json:"time_partitioning_field,omitempty"`
	TimePartitioningExpirationMs           *int64                                   `json:"time_partitioning_expiration_ms,omitempty"`
	TimePartitioningRequirePartitionFilter *bool                                    `json:"time_partitioning_require_partition_filter,omitempty"`
	Location                               *string                                  `json:"location,omitempty"`
	TemplateTable                          *string                                  `json:"template_table,omitempty"`
	BigQueryConnectionID                   *int64                                   `json:"bigquery_connection_id,omitempty"`
	BeforeLoad                             *string                                  `json:"before_load,omitempty"`
	BigQueryOutputOptionColumnOptions      *[]BigQueryOutputOptionColumnOptionInput `json:"bigquery_output_option_column_options,omitempty"`
	BigQueryOutputOptionClusteringFields   *[]string                                `json:"bigquery_output_option_clustering_fields,omitempty"`
	BigQueryOutputOptionMergeKeys          *[]string                                `json:"bigquery_output_option_merge_keys,omitempty"`
}

type BigQueryOutputOptionColumnOptionInput struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Mode            string  `json:"mode"`
	TimestampFormat *string `json:"timestamp_format,omitempty"`
	Timezone        *string `json:"timezone,omitempty"`
	Description     *string `json:"description,omitempty"`
}
