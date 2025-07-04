package output_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type SnowflakeOutputOption struct {
	Warehouse                          string                              `json:"warehouse"`
	Database                           string                              `json:"database"`
	Schema                             string                              `json:"schema"`
	Table                              string                              `json:"table"`
	Mode                               *string                             `json:"mode"`
	EmptyFieldAsNull                   *bool                               `json:"empty_field_as_null"`
	DeleteStageOnError                 *bool                               `json:"delete_stage_on_error"`
	BatchSize                          *int64                              `json:"batch_size"`
	RetryLimit                         *int64                              `json:"retry_limit"`
	RetryWait                          *int64                              `json:"retry_wait"`
	MaxRetryWait                       *int64                              `json:"max_retry_wait"`
	DefaultTimeZone                    *string                             `json:"default_time_zone"`
	SnowflakeConnectionID              int64                               `json:"snowflake_connection_id"`
	SnowflakeOutputOptionColumnOptions []SnowflakeOutputOptionColumnOption `json:"snowflake_output_option_column_options"`
	SnowflakeOutputOptionMergeKeys     []string                            `json:"snowflake_output_option_merge_keys"`
	CustomVariableSettings             *[]entity.CustomVariableSetting     `json:"custom_variable_settings"`
}

type SnowflakeOutputOptionColumnOption struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	ValueType       *string `json:"value_type"`
	TimestampFormat *string `json:"timestamp_format"`
	Timezone        *string `json:"timezone"`
}
