package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type DatabricksInputOptionInput struct {
	DatabricksConnectionID int64                    `json:"databricks_connection_id"`
	CatalogName            string                   `json:"catalog_name"`
	SchemaName             string                   `json:"schema_name"`
	Query                  string                   `json:"query"`
	FetchRows              *parameter.NullableInt64 `json:"fetch_rows,omitempty"`

	InputOptionColumns     []DatabricksInputOptionColumn           `json:"input_option_columns"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateDatabricksInputOptionInput struct {
	DatabricksConnectionID int64                    `json:"databricks_connection_id"`
	CatalogName            string                   `json:"catalog_name"`
	SchemaName             string                   `json:"schema_name"`
	Query                  string                   `json:"query"`
	FetchRows              *parameter.NullableInt64 `json:"fetch_rows,omitempty"`

	InputOptionColumns     []DatabricksInputOptionColumn           `json:"input_option_columns,omitempty"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type DatabricksInputOptionColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
