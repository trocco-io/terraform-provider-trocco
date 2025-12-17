package input_option

import "terraform-provider-trocco/internal/client/entity"

type DatabricksInputOption struct {
	DatabricksConnectionID int64  `json:"databricks_connection_id"`
	CatalogName            string `json:"catalog_name"`
	SchemaName             string `json:"schema_name"`
	Query                  string `json:"query"`
	FetchRows              int64  `json:"fetch_rows"`

	InputOptionColumns     []DatabricksColumn              `json:"input_option_columns"`
	CustomVariableSettings *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type DatabricksColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
