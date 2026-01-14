package output_option

type PostgresqlOutputOption struct {
	Database               string   `json:"database"`
	Schema                 string   `json:"schema"`
	Table                  string   `json:"table"`
	Mode                   *string  `json:"mode"`
	DefaultTimeZone        *string  `json:"default_time_zone"`
	PostgresqlConnectionId int64    `json:"postgresql_connection_id"`
	MergeKeys              []string `json:"merge_keys"`
}
