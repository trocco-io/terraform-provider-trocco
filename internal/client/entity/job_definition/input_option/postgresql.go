package input_option

import "terraform-provider-trocco/internal/client/entity"

type PostgreSQLInputOption struct {
	PostgreSQLConnectionID             int64                                 `json:"postgresql_connection_id"`
	Database                           string                                `json:"database"`
	Schema                             string                                `json:"schema"`
	Query                              *string                               `json:"query"`
	IncrementalLoadingEnabled          bool                                  `json:"incremental_loading_enabled"`
	Table                              *string                               `json:"table"`
	IncrementalColumns                 *string                               `json:"incremental_columns"`
	LastRecord                         *string                               `json:"last_record"`
	FetchRows                          int64                                 `json:"fetch_rows"`
	ConnectTimeout                     int64                                 `json:"connect_timeout"`
	SocketTimeout                      int64                                 `json:"socket_timeout"`
	DefaultTimeZone                    string                                `json:"default_time_zone"`
	PostgreSQLInputOptionColumnOptions *[]PostgreSQLInputOptionColumnOptions `json:"postgresql_input_option_column_options"`
	CustomVariableSettings             *[]entity.CustomVariableSetting       `json:"custom_variable_settings"`
}

type PostgreSQLInputOptionColumnOptions struct {
	ColumnName      string `json:"column_name"`
	ColumnValueType string `json:"column_value_type"`
}
