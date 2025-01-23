package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type MySQLInputOptionInput struct {
	Database                  string                                  `json:"database"`
	Table                     *parameter.NullableString               `json:"table,omitempty"`
	Query                     *parameter.NullableString               `json:"query"`
	IncrementalColumns        *parameter.NullableString               `json:"incremental_columns,omitempty"`
	LastRecord                *parameter.NullableString               `json:"last_record,omitempty"`
	IncrementalLoadingEnabled bool                                    `json:"incremental_loading_enabled"`
	FetchRows                 int64                                   `json:"fetch_rows"`
	ConnectTimeout            int64                                   `json:"connect_timeout"`
	SocketTimeout             int64                                   `json:"socket_timeout"`
	DefaultTimeZone           *parameter.NullableString               `json:"default_time_zone,omitempty"`
	UseLegacyDatetimeCode     bool                                    `json:"use_legacy_datetime_code,omitempty"`
	MySQLConnectionID         int64                                   `json:"mysql_connection_id"`
	InputOptionColumns        []InputOptionColumn                     `json:"input_option_columns"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateMySQLInputOptionInput struct {
	Database                  *string                                 `json:"database,omitempty"`
	Table                     *parameter.NullableString               `json:"table,omitempty"`
	Query                     *parameter.NullableString               `json:"query,omitempty"`
	IncrementalColumns        *parameter.NullableString               `json:"incremental_columns,omitempty"`
	LastRecord                *parameter.NullableString               `json:"last_record,omitempty"`
	IncrementalLoadingEnabled *bool                                   `json:"incremental_loading_enabled,omitempty"`
	FetchRows                 *int64                                  `json:"fetch_rows,omitempty"`
	ConnectTimeout            *int64                                  `json:"connect_timeout,omitempty"`
	SocketTimeout             *int64                                  `json:"socket_timeout,omitempty"`
	DefaultTimeZone           *parameter.NullableString               `json:"default_time_zone,omitempty"`
	UseLegacyDatetimeCode     *bool                                   `json:"use_legacy_datetime_code,omitempty"`
	MySQLConnectionID         *int64                                  `json:"mysql_connection_id,omitempty"`
	InputOptionColumns        *[]InputOptionColumn                    `json:"input_option_columns,omitempty"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type InputOptionColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
