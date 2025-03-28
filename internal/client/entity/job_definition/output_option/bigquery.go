package output_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type BigQueryOutputOption struct {
	CustomVariableSettings               *[]entity.CustomVariableSetting     `json:"custom_variable_settings"`
	Dataset                              string                              `json:"dataset"`
	Table                                string                              `json:"table"`
	AutoCreateDataset                    bool                                `json:"auto_create_dataset"`
	OpenTimeoutSec                       int64                               `json:"open_timeout_sec"`
	TimeoutSec                           int64                               `json:"timeout_sec"`
	SendTimeoutSec                       int64                               `json:"send_timeout_sec"`
	ReadTimeoutSec                       int64                               `json:"read_timeout_sec"`
	Retries                              int64                               `json:"retries"`
	Mode                                 string                              `json:"mode"`
	PartitioningType                     *string                             `json:"partitioning_type"`
	TimePartitioningType                 *string                             `json:"time_partitioning_type"`
	TimePartitioningField                *string                             `json:"time_partitioning_field"`
	TimePartitioningExpirationMs         *int64                              `json:"time_partitioning_expiration_ms"`
	Location                             *string                             `json:"location"`
	TemplateTable                        *string                             `json:"template_table"`
	BigQueryConnectionID                 int64                               `json:"bigquery_connection_id"`
	BigQueryOutputOptionColumnOptions    *[]BigQueryOutputOptionColumnOption `json:"bigquery_output_option_column_options"`
	BigQueryOutputOptionClusteringFields *[]string                           `json:"bigquery_output_option_clustering_fields"`
	BigQueryOutputOptionMergeKeys        *[]string                           `json:"bigquery_output_option_merge_keys"`
}

type BigQueryOutputOptionColumnOption struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Mode            string  `json:"mode"`
	TimestampFormat *string `json:"timestamp_format"`
	Timezone        *string `json:"timezone"`
	Description     *string `json:"description"`
}
