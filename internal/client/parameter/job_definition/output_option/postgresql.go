package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type PostgresqlOutputOptionInput struct {
	Database               string                                `json:"database"`
	Schema                 string                                `json:"schema"`
	Table                  string                                `json:"table"`
	Mode                   string                                `json:"mode"`
	DefaultTimeZone        string                                `json:"default_time_zone"`
	PostgresqlConnectionId int64                                 `json:"postgresql_connection_id"`
	MergeKeys              *parameter.NullableObjectList[string] `json:"merge_keys,omitempty"`
	RetryLimit             *int64                                `json:"retry_limit,omitempty"`
	RetryWait              *int64                                `json:"retry_wait,omitempty"`
	MaxRetryWait           *int64                                `json:"max_retry_wait,omitempty"`
	BeforeLoad             *string                               `json:"before_load,omitempty"`
	AfterLoad              *string                               `json:"after_load,omitempty"`
}

type UpdatePostgresqlOutputOptionInput struct {
	Database               *string                               `json:"database,omitempty"`
	Schema                 *string                               `json:"schema,omitempty"`
	Table                  *string                               `json:"table,omitempty"`
	Mode                   *string                               `json:"mode,omitempty"`
	DefaultTimeZone        *string                               `json:"default_time_zone,omitempty"`
	PostgresqlConnectionId *int64                                `json:"postgresql_connection_id,omitempty"`
	MergeKeys              *parameter.NullableObjectList[string] `json:"merge_keys,omitempty"`
	RetryLimit             *int64                                `json:"retry_limit,omitempty"`
	RetryWait              *int64                                `json:"retry_wait,omitempty"`
	MaxRetryWait           *int64                                `json:"max_retry_wait,omitempty"`
	BeforeLoad             *string                               `json:"before_load,omitempty"`
	AfterLoad              *string                               `json:"after_load,omitempty"`
}
