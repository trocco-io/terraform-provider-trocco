package input_options

import "terraform-provider-trocco/internal/client/parameter"

type OracleInputOptionInput struct {
	OracleConnectionID         int64                                   `json:"oracle_connection_id"`
	Database                   *parameter.NullableString               `json:"database,omitempty"`
	ConnectionType             *parameter.NullableString               `json:"connection_type,omitempty"`
	NetServiceName             *parameter.NullableString               `json:"net_service_name,omitempty"`
	Schema                     *parameter.NullableString               `json:"schema,omitempty"`
	Query                      *string                                 `json:"query,omitempty"`
	IncrementalLoadingEnabled  bool                                    `json:"incremental_loading_enabled"`
	Table                      *string                                 `json:"table,omitempty"`
	IncrementalColumns         *parameter.NullableString               `json:"incremental_columns,omitempty"`
	DefaultTimeZone            *string                                 `json:"default_time_zone,omitempty"`
	InputOptionColumns         *[]OracleInputOptionColumnInput         `json:"input_option_columns,omitempty"`
	InputOptionColumnOptions   *[]OracleInputOptionColumnOptionInput   `json:"input_option_column_options,omitempty"`
	CustomVariableSettings     *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type OracleInputOptionColumnInput struct {
	Name     string                    `json:"name"`
	Type     string                    `json:"type"`
	Format   *parameter.NullableString `json:"format,omitempty"`
	Timezone *parameter.NullableString `json:"timezone,omitempty"`
}

type OracleInputOptionColumnOptionInput struct {
	ColumnName      string `json:"column_name"`
	ColumnValueType string `json:"column_value_type"`
}

type UpdateOracleInputOptionInput struct {
	OracleConnectionID         *int64                                  `json:"oracle_connection_id,omitempty"`
	Database                   *parameter.NullableString               `json:"database,omitempty"`
	ConnectionType             *parameter.NullableString               `json:"connection_type,omitempty"`
	NetServiceName             *parameter.NullableString               `json:"net_service_name,omitempty"`
	Schema                     *parameter.NullableString               `json:"schema,omitempty"`
	Query                      *string                                 `json:"query,omitempty"`
	IncrementalLoadingEnabled  *bool                                   `json:"incremental_loading_enabled,omitempty"`
	Table                      *string                                 `json:"table,omitempty"`
	IncrementalColumns         *parameter.NullableString               `json:"incremental_columns,omitempty"`
	DefaultTimeZone            *string                                 `json:"default_time_zone,omitempty"`
	InputOptionColumns         *[]OracleInputOptionColumnInput         `json:"input_option_columns,omitempty"`
	InputOptionColumnOptions   *[]OracleInputOptionColumnOptionInput   `json:"input_option_column_options,omitempty"`
	CustomVariableSettings     *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}
