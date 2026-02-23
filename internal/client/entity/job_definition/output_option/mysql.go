package output_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type MysqlOutputOption struct {
	MysqlConnectionID              int64                            `json:"mysql_connection_id"`
	Database                       string                           `json:"database"`
	Table                          string                           `json:"table"`
	Mode                           string                           `json:"mode"`
	RetryLimit                     int64                            `json:"retry_limit"`
	RetryWait                      int64                            `json:"retry_wait"`
	MaxRetryWait                   int64                            `json:"max_retry_wait"`
	DefaultTimeZone                string                           `json:"default_time_zone"`
	BeforeLoad                     *string                          `json:"before_load"`
	AfterLoad                      *string                          `json:"after_load"`
	MysqlOutputOptionColumnOptions *[]MysqlOutputOptionColumnOption `json:"mysql_output_option_column_options"`
	CustomVariableSettings         *[]entity.CustomVariableSetting  `json:"custom_variable_settings"`
}

type MysqlOutputOptionColumnOption struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Scale     *int64 `json:"scale"`
	Precision *int64 `json:"precision"`
}
