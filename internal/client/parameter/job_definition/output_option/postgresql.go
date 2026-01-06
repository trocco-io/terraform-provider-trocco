package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type PostgresqlOutputOptionInput struct {
	Database                        string                                `json:"database"`
	Schema                          string                                `json:"schema"`
	Table                           string                                `json:"table"`
	Mode                            *parameter.NullableString             `json:"mode,omitempty"`
	DefaultTimeZone                 *parameter.NullableString             `json:"default_time_zone,omitempty"`
	PostgresqlConnectionId          int64                                 `json:"postgresql_connection_id"`
	PostgresqlOutputOptionMergeKeys *parameter.NullableObjectList[string] `json:"postgresql_output_option_merge_keys,omitempty"`
}

type UpdatePostgresqlOutputOptionInput struct {
	Database                        *string                               `json:"database,omitempty"`
	Schema                          *string                               `json:"schema,omitempty"`
	Table                           *string                               `json:"table,omitempty"`
	Mode                            *parameter.NullableString             `json:"mode,omitempty"`
	DefaultTimeZone                 *parameter.NullableString             `json:"default_time_zone,omitempty"`
	PostgresqlConnectionId          *int64                                `json:"postgresql_connection_id,omitempty"`
	PostgresqlOutputOptionMergeKeys *parameter.NullableObjectList[string] `json:"postgresql_output_option_merge_keys,omitempty"`
}
