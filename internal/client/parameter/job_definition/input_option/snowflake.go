package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type SnowflakeInputOptionInput struct {
	Warehouse              string                                  `json:"warehouse"`
	Database               string                                  `json:"database"`
	Schema                 string                                  `json:"schema"`
	Query                  string                                  `json:"query"`
	FetchRows              *parameter.NullableInt64                `json:"fetch_rows,omitempty"`
	ConnectTimeout         *parameter.NullableInt64                `json:"connect_timeout,omitempty"`
	SocketTimeout          *parameter.NullableInt64                `json:"socket_timeout,omitempty"`
	SnowflakeConnectionID  int64                                   `json:"snowflake_connection_id"`
	InputOptionColumns     []SnowflakeInputOptionColumn            `json:"input_option_columns"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateSnowflakeInputOptionInput struct {
	Warehouse              *string                                 `json:"warehouse,omitempty"`
	Database               *string                                 `json:"database,omitempty"`
	Schema                 *string                                 `json:"schema,omitempty"`
	Query                  *string                                 `json:"query,omitempty"`
	FetchRows              *parameter.NullableInt64                `json:"fetch_rows,omitempty"`
	ConnectTimeout         *parameter.NullableInt64                `json:"connect_timeout,omitempty"`
	SocketTimeout          *parameter.NullableInt64                `json:"socket_timeout,omitempty"`
	SnowflakeConnectionID  *int64                                  `json:"snowflake_connection_id,omitempty"`
	InputOptionColumns     []SnowflakeInputOptionColumn            `json:"input_option_columns,omitempty"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type SnowflakeInputOptionColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
