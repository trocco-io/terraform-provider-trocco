package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type SnowflakeInputOption struct {
	Warehouse              string                          `json:"warehouse"`
	Database               string                          `json:"database"`
	Schema                 string                          `json:"schema"`
	Query                  string                          `json:"query"`
	FetchRows              *int64                          `json:"fetch_rows"`
	ConnectTimeout         *int64                          `json:"connect_timeout"`
	SocketTimeout          *int64                          `json:"socket_timeout"`
	SnowflakeConnectionID  int64                           `json:"snowflake_connection_id"`
	InputOptionColumns     []SnowflakeInputOptionColumn    `json:"input_option_columns"`
	CustomVariableSettings *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type SnowflakeInputOptionColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
