package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type DatabricksOutputOptionInput struct {
	DatabricksConnectionID              int64                                                                  `json:"databricks_connection_id"`
	CatalogName                         string                                                                 `json:"catalog_name"`
	SchemaName                          string                                                                 `json:"schema_name"`
	Table                               string                                                                 `json:"table"`
	BatchSize                           *int64                                                                 `json:"batch_size,omitempty"`
	Mode                                string                                                                 `json:"mode"`
	DefaultTimeZone                     *string                                                                `json:"default_time_zone,omitempty"`
	DatabricksOutputOptionColumnOptions *parameter.NullableObjectList[DatabricksOutputOptionColumnOptionInput] `json:"databricks_output_option_column_options,omitempty"`
	DatabricksOutputOptionMergeKeys     *parameter.NullableObjectList[string]                                  `json:"databricks_output_option_merge_keys,omitempty"`
}

type UpdateDatabricksOutputOptionInput struct {
	DatabricksConnectionID              *int64                                                                 `json:"databricks_connection_id,omitempty"`
	CatalogName                         *string                                                                `json:"catalog_name,omitempty"`
	SchemaName                          *string                                                                `json:"schema_name,omitempty"`
	Table                               *string                                                                `json:"table,omitempty"`
	BatchSize                           *int64                                                                 `json:"batch_size,omitempty"`
	Mode                                *string                                                                `json:"mode,omitempty"`
	DefaultTimeZone                     *string                                                                `json:"default_time_zone,omitempty"`
	DatabricksOutputOptionColumnOptions *parameter.NullableObjectList[DatabricksOutputOptionColumnOptionInput] `json:"databricks_output_option_column_options,omitempty"`
	DatabricksOutputOptionMergeKeys     *parameter.NullableObjectList[string]                                  `json:"databricks_output_option_merge_keys,omitempty"`
}

type DatabricksOutputOptionColumnOptionInput struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	ValueType       *string `json:"value_type,omitempty"`
	TimestampFormat *string `json:"timestamp_format,omitempty"`
	Timezone        *string `json:"timezone,omitempty"`
}
