package output_option

type DatabricksOutputOption struct {
	DatabricksConnectionID int64  `json:"databricks_connection_id"`
	CatalogName            string `json:"catalog_name"`
	SchemaName             string `json:"schema_name"`
	Table                  string `json:"table"`
	BatchSize              int64  `json:"batch_size"`
}
