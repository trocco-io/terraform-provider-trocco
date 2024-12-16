package input_options

import "terraform-provider-trocco/internal/client/entities"

type MySQLInputOption struct {
	Database                  string                            `json:"database"`
	Table                     *string                           `json:"table"`
	Query                     string                            `json:"query"`
	IncrementalColumns        *string                           `json:"incremental_columns"`
	LastRecord                *string                           `json:"last_record"`
	IncrementalLoadingEnabled bool                              `json:"incremental_loading_enabled"`
	FetchRows                 int64                             `json:"fetch_rows"`
	ConnectTimeout            int64                             `json:"connect_timeout"`
	SocketTimeout             int64                             `json:"socket_timeout"`
	DefaultTimeZone           *string                           `json:"default_time_zone"`
	UseLegacyDatetimeCode     *bool                             `json:"use_legacy_datetime_code"`
	MySQLConnectionID         int64                             `json:"mysql_connection_id"`
	InputOptionColumns        []InputOptionColumn               `json:"input_option_columns"`
	CustomVariableSettings    *[]entities.CustomVariableSetting `json:"custom_variable_settings"`
}

type InputOptionColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
