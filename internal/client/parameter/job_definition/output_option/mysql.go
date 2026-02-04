package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type MysqlOutputOptionInput struct {
	MysqlConnectionID              int64                                                             `json:"mysql_connection_id"`
	Database                       string                                                            `json:"database"`
	Table                          string                                                            `json:"table"`
	Mode                           string                                                            `json:"mode"`
	RetryLimit                     int64                                                             `json:"retry_limit"`
	RetryWait                      int64                                                             `json:"retry_wait"`
	MaxRetryWait                   int64                                                             `json:"max_retry_wait"`
	DefaultTimeZone                string                                                            `json:"default_time_zone"`
	BeforeLoad                     *string                                                           `json:"before_load,omitempty"`
	AfterLoad                      *string                                                           `json:"after_load,omitempty"`
	MysqlOutputOptionColumnOptions *parameter.NullableObjectList[MysqlOutputOptionColumnOptionInput] `json:"mysql_output_option_column_options,omitempty"`
	CustomVariableSettings         *[]parameter.CustomVariableSettingInput                           `json:"custom_variable_settings,omitempty"`
}

type UpdateMysqlOutputOptionInput struct {
	MysqlConnectionID              *int64                                                            `json:"mysql_connection_id,omitempty"`
	Database                       *string                                                           `json:"database,omitempty"`
	Table                          *string                                                           `json:"table,omitempty"`
	Mode                           *string                                                           `json:"mode,omitempty"`
	RetryLimit                     *int64                                                            `json:"retry_limit,omitempty"`
	RetryWait                      *int64                                                            `json:"retry_wait,omitempty"`
	MaxRetryWait                   *int64                                                            `json:"max_retry_wait,omitempty"`
	DefaultTimeZone                *string                                                           `json:"default_time_zone,omitempty"`
	BeforeLoad                     *string                                                           `json:"before_load,omitempty"`
	AfterLoad                      *string                                                           `json:"after_load,omitempty"`
	MysqlOutputOptionColumnOptions *parameter.NullableObjectList[MysqlOutputOptionColumnOptionInput] `json:"mysql_output_option_column_options,omitempty"`
	CustomVariableSettings         *[]parameter.CustomVariableSettingInput                           `json:"custom_variable_settings,omitempty"`
}

type MysqlOutputOptionColumnOptionInput struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Scale     *int64 `json:"scale,omitempty"`
	Precision *int64 `json:"precision,omitempty"`
}
