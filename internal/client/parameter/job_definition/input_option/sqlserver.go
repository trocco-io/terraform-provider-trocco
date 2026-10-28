package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type SQLServerInputOptionInput struct {
	Database                  string                                                  `json:"database"`
	Schema                    *parameter.NullableString                               `json:"schema,omitempty"`
	Table                     *parameter.NullableString                               `json:"table,omitempty"`
	Query                     *parameter.NullableString                               `json:"query,omitempty"`
	IncrementalColumns        *parameter.NullableString                               `json:"incremental_columns,omitempty"`
	LastRecord                *parameter.NullableString                               `json:"last_record,omitempty"`
	IncrementalLoadingEnabled *parameter.NullableBool                                 `json:"incremental_loading_enabled,omitempty"`
	FetchRows                 *parameter.NullableInt64                                `json:"fetch_rows,omitempty"`
	ConnectTimeout            *parameter.NullableInt64                                `json:"connect_timeout,omitempty"`
	SocketTimeout             *parameter.NullableInt64                                `json:"socket_timeout,omitempty"`
	DefaultTimeZone           *parameter.NullableString                               `json:"default_time_zone,omitempty"`
	SQLServerConnectionID     int64                                                   `json:"sqlserver_connection_id"`
	InputOptionColumns        []SQLServerInputOptionColumn                            `json:"input_option_columns"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput                 `json:"custom_variable_settings,omitempty"`
	InputOptionColumnOptions  *parameter.NullableObjectList[InputOptionColumnOptions] `json:"input_option_column_options,omitempty"`
}

type UpdateSQLServerInputOptionInput struct {
	Database                  *string                                                 `json:"database,omitempty"`
	Schema                    *parameter.NullableString                               `json:"schema,omitempty"`
	Table                     *parameter.NullableString                               `json:"table,omitempty"`
	Query                     *parameter.NullableString                               `json:"query,omitempty"`
	IncrementalColumns        *parameter.NullableString                               `json:"incremental_columns,omitempty"`
	LastRecord                *parameter.NullableString                               `json:"last_record,omitempty"`
	IncrementalLoadingEnabled *bool                                                   `json:"incremental_loading_enabled,omitempty"`
	FetchRows                 *parameter.NullableInt64                                `json:"fetch_rows,omitempty"`
	ConnectTimeout            *parameter.NullableInt64                                `json:"connect_timeout,omitempty"`
	SocketTimeout             *parameter.NullableInt64                                `json:"socket_timeout,omitempty"`
	DefaultTimeZone           *parameter.NullableString                               `json:"default_time_zone,omitempty"`
	SQLServerConnectionID     *int64                                                  `json:"sqlserver_connection_id,omitempty"`
	InputOptionColumns        []SQLServerInputOptionColumn                            `json:"input_option_columns,omitempty"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput                 `json:"custom_variable_settings,omitempty"`
	InputOptionColumnOptions  *parameter.NullableObjectList[InputOptionColumnOptions] `json:"input_option_column_options,omitempty"`
}

type SQLServerInputOptionColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
