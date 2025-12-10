package output_options

type DatabricksOutputOptionInput struct {
	DatabricksConnectionID int64  `json:"databricks_connection_id"`
	CatalogName            string `json:"catalog_name"`
	SchemaName             string `json:"schema_name"`
	Table                  string `json:"table"`
	BatchSize              int64  `json:"batch_size"`
	Mode                   string `json:"mode"`
	DefaultTimeZone        string `json:"default_time_zone"`
}

type UpdateDatabricksOutputOptionInput struct {
	DatabricksConnectionID *int64  `json:"databricks_connection_id,omitempty"`
	CatalogName            *string `json:"catalog_name,omitempty"`
	SchemaName             *string `json:"schema_name,omitempty"`
	Table                  *string `json:"table,omitempty"`
	BatchSize              *int64  `json:"batch_size,omitempty"`
	Mode                   *string `json:"mode,omitempty"`
	DefaultTimeZone        *string `json:"default_time_zone,omitempty"`
}
