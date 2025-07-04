package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type SnowflakeOutputOptionInput struct {
	Warehouse                          string                                                                `json:"warehouse"`
	Database                           string                                                                `json:"database"`
	Schema                             string                                                                `json:"schema"`
	Table                              string                                                                `json:"table"`
	Mode                               *parameter.NullableString                                             `json:"mode,omitempty"`
	EmptyFieldAsNull                   *parameter.NullableBool                                               `json:"empty_field_as_null,omitempty"`
	DeleteStageOnError                 *parameter.NullableBool                                               `json:"delete_stage_on_error,omitempty"`
	BatchSize                          *parameter.NullableInt64                                              `json:"batch_size,omitempty"`
	RetryLimit                         *parameter.NullableInt64                                              `json:"retry_limit,omitempty"`
	RetryWait                          *parameter.NullableInt64                                              `json:"retry_wait,omitempty"`
	MaxRetryWait                       *parameter.NullableInt64                                              `json:"max_retry_wait,omitempty"`
	DefaultTimeZone                    *parameter.NullableString                                             `json:"default_time_zone,omitempty"`
	SnowflakeConnectionID              int64                                                                 `json:"snowflake_connection_id"`
	SnowflakeOutputOptionColumnOptions *parameter.NullableObjectList[SnowflakeOutputOptionColumnOptionInput] `json:"snowflake_output_option_column_options,omitempty"`
	SnowflakeOutputOptionMergeKeys     *parameter.NullableObjectList[string]                                 `json:"snowflake_output_option_merge_keys,omitempty"`
	CustomVariableSettings             *[]parameter.CustomVariableSettingInput                               `json:"custom_variable_settings,omitempty"`
}

type UpdateSnowflakeOutputOptionInput struct {
	Warehouse                          *string                                                               `json:"warehouse,omitempty"`
	Database                           *string                                                               `json:"database,omitempty"`
	Schema                             *string                                                               `json:"schema,omitempty"`
	Table                              *string                                                               `json:"table,omitempty"`
	Mode                               *parameter.NullableString                                             `json:"mode,omitempty"`
	EmptyFieldAsNull                   *parameter.NullableBool                                               `json:"empty_field_as_null,omitempty"`
	DeleteStageOnError                 *parameter.NullableBool                                               `json:"delete_stage_on_error,omitempty"`
	BatchSize                          *parameter.NullableInt64                                              `json:"batch_size,omitempty"`
	RetryLimit                         *parameter.NullableInt64                                              `json:"retry_limit,omitempty"`
	RetryWait                          *parameter.NullableInt64                                              `json:"retry_wait,omitempty"`
	MaxRetryWait                       *parameter.NullableInt64                                              `json:"max_retry_wait,omitempty"`
	DefaultTimeZone                    *parameter.NullableString                                             `json:"default_time_zone,omitempty"`
	SnowflakeConnectionID              *int64                                                                `json:"snowflake_connection_id,omitempty"`
	SnowflakeOutputOptionColumnOptions *parameter.NullableObjectList[SnowflakeOutputOptionColumnOptionInput] `json:"snowflake_output_option_column_options,omitempty"`
	SnowflakeOutputOptionMergeKeys     *parameter.NullableObjectList[string]                                 `json:"snowflake_output_option_merge_keys,omitempty"`
	CustomVariableSettings             *[]parameter.CustomVariableSettingInput                               `json:"custom_variable_settings,omitempty"`
}

type SnowflakeOutputOptionColumnOptionInput struct {
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	ValueType       *string `json:"value_type,omitempty"`
	TimestampFormat *string `json:"timestamp_format,omitempty"`
	Timezone        *string `json:"timezone,omitempty"`
}
