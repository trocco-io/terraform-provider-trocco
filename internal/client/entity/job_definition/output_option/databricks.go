package output_option

type DatabricksOutputOption struct {
	DatabricksConnectionID              int64                                `json:"databricks_connection_id"`
	CatalogName                         string                               `json:"catalog_name"`
	SchemaName                          string                               `json:"schema_name"`
	Table                               string                               `json:"table"`
	BatchSize                           int64                                `json:"batch_size"`
	Mode                                string                               `json:"mode"`
	DefaultTimeZone                     string                               `json:"default_time_zone"`
	DatabricksOutputOptionColumnOptions []DatabricksOutputOptionColumnOption `json:"databricks_output_option_column_options"`
	DatabricksOutputOptionMergeKeys     []string                             `json:"databricks_output_option_merge_keys"`
}

type DatabricksOutputOptionColumnOption struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	ValueType       *string `json:"value_type"`
	TimestampFormat *string `json:"timestamp_format"`
	Timezone        *string `json:"timezone"`
}
