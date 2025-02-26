package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type PostgreSQLInputOptionInput struct {
	Database                  string                                  `json:"database"`
	Schema                    string                                  `json:"schema"`
	Table                     *parameter.NullableString               `json:"table,omitempty"`
	Query                     *parameter.NullableString               `json:"query,omitempty"`
	IncrementalColumns        *parameter.NullableString               `json:"incremental_columns,omitempty"`
	LastRecord                *parameter.NullableString               `json:"last_record,omitempty"`
	IncrementalLoadingEnabled bool                                    `json:"incremental_loading_enabled"`
	FetchRows                 int64                                   `json:"fetch_rows"`
	ConnectTimeout            int64                                   `json:"connect_timeout"`
	SocketTimeout             int64                                   `json:"socket_timeout"`
	DefaultTimeZone           string                                  `json:"default_time_zone"`
	PostgreSQLConnectionID    int64                                   `json:"postgresql_connection_id"`
	InputOptionColumns        []PostgreSQLInputOptionColumn           `json:"input_option_columns"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	InputOptionColumnOptions  *[]InputOptionColumnOptions             `json:"input_option_column_options,omitempty"`
}

type InputOptionColumnOptions struct {
	ColumnName      string `json:"column_name,omitempty"`
	ColumnValueType string `json:"column_value_type,omitempty"`
}

type UpdatePostgreSQLInputOptionInput struct {
	Database                  *string                                 `json:"database,omitempty"`
	Schema                    *string                                 `json:"schema,omitempty"`
	Table                     *parameter.NullableString               `json:"table,omitempty"`
	Query                     *parameter.NullableString               `json:"query,omitempty"`
	IncrementalColumns        *parameter.NullableString               `json:"incremental_columns,omitempty"`
	LastRecord                *parameter.NullableString               `json:"last_record,omitempty"`
	IncrementalLoadingEnabled *bool                                   `json:"incremental_loading_enabled,omitempty"`
	FetchRows                 *int64                                  `json:"fetch_rows,omitempty"`
	ConnectTimeout            *int64                                  `json:"connect_timeout,omitempty"`
	SocketTimeout             *int64                                  `json:"socket_timeout,omitempty"`
	DefaultTimeZone           *string                                 `json:"default_time_zone,omitempty"`
	PostgreSQLConnectionID    *int64                                  `json:"postgresql_connection_id,omitempty"`
	InputOptionColumns        []PostgreSQLInputOptionColumn           `json:"input_option_columns"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	InputOptionColumnOptions  *[]InputOptionColumnOptions             `json:"input_option_column_options,omitempty"`
}

type PostgreSQLInputOptionColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
