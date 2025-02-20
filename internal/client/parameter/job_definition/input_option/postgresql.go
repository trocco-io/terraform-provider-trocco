package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type PostgreSQLInputOptionInput struct {
	Database                           string                                  `json:"database"`
	Schema                             string                                  `json:"schema"`
	Table                              *parameter.NullableString               `json:"table,omitempty"`
	Query                              *parameter.NullableString               `json:"query"`
	IncrementalColumns                 *parameter.NullableString               `json:"incremental_columns,omitempty"`
	LastRecord                         *parameter.NullableString               `json:"last_record,omitempty"`
	IncrementalLoadingEnabled          bool                                    `json:"incremental_loading_enabled"`
	FetchRows                          int64                                   `json:"fetch_rows"`
	ConnectTimeout                     int64                                   `json:"connect_timeout"`
	SocketTimeout                      int64                                   `json:"socket_timeout"`
	DefaultTimeZone                    string                                  `json:"default_time_zone,omitempty"`
	PostgreSQLConnectionID             int64                                   `json:"postgresql_connection_id"`
	PostgreSQLInputOptionColumnOptions *[]PostgreSQLInputOptionColumnOptions   `json:"postgresql_input_option_column_options"`
	CustomVariableSettings             *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdatePostgreSQLInputOptionInput struct {
	Database                           *string                                 `json:"database,omitempty"`
	Schema                             *string                                 `json:"schema,omitempty"`
	Table                              *parameter.NullableString               `json:"table,omitempty"`
	Query                              *parameter.NullableString               `json:"query,omitempty"`
	IncrementalColumns                 *parameter.NullableString               `json:"incremental_columns,omitempty"`
	LastRecord                         *parameter.NullableString               `json:"last_record,omitempty"`
	IncrementalLoadingEnabled          *bool                                   `json:"incremental_loading_enabled,omitempty"`
	FetchRows                          *int64                                  `json:"fetch_rows,omitempty"`
	ConnectTimeout                     *int64                                  `json:"connect_timeout,omitempty"`
	SocketTimeout                      *int64                                  `json:"socket_timeout,omitempty"`
	DefaultTimeZone                    *parameter.NullableString               `json:"default_time_zone,omitempty"`
	UseLegacyDatetimeCode              *bool                                   `json:"use_legacy_datetime_code,omitempty"`
	PostgreSQLConnectionID             *int64                                  `json:"postgresql_connection_id,omitempty"`
	PostgreSQLInputOptionColumnOptions *[]PostgreSQLInputOptionColumnOptions   `json:"postgresql_input_option_column_options,omitempty"`
	CustomVariableSettings             *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type PostgreSQLInputOptionColumnOptions struct {
	ColumnName      string `json:"cloumn_name,omitempty"`
	ColumnValueType string `json:"cloumn_value_type,omitempty"`
}
