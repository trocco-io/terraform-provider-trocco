package input_option

import "terraform-provider-trocco/internal/client/entity"

type OracleInputOption struct {
	OracleConnectionID        int64                            `json:"oracle_connection_id"`
	Database                  *string                          `json:"database"`
	ConnectionType            string                           `json:"connection_type"`
	NetServiceName            *string                          `json:"net_service_name"`
	Schema                    *string                          `json:"schema"`
	Query                     *string                          `json:"query"`
	IncrementalLoadingEnabled bool                             `json:"incremental_loading_enabled"`
	Table                     *string                          `json:"table"`
	IncrementalColumns        *string                          `json:"incremental_columns"`
	SourceTimeZone            *string                          `json:"source_time_zone"`
	InputOptionColumns        *[]OracleInputOptionColumn       `json:"input_option_columns"`
	InputOptionColumnOptions  *[]OracleInputOptionColumnOption `json:"input_option_column_options"`
	CustomVariableSettings    *[]entity.CustomVariableSetting  `json:"custom_variable_settings"`
}

type OracleInputOptionColumn struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Format   *string `json:"format"`
	Timezone *string `json:"timezone"`
}

type OracleInputOptionColumnOption struct {
	ColumnName      string `json:"column_name"`
	ColumnValueType string `json:"column_value_type"`
}
