package output_option

type PostgresqlOutputOption struct {
	Database               string   `json:"database"`
	Schema                 string   `json:"schema"`
	Table                  string   `json:"table"`
	Mode                   *string  `json:"mode"`
	DefaultTimeZone        *string  `json:"default_time_zone"`
	PostgresqlConnectionId int64    `json:"postgresql_connection_id"`
	MergeKeys              []string `json:"merge_keys"`
	RetryLimit             *int64   `json:"retry_limit"`
	RetryWait              *int64   `json:"retry_wait"`
	MaxRetryWait           *int64   `json:"max_retry_wait"`
	BeforeLoad             *string  `json:"before_load"`
	AfterLoad              *string  `json:"after_load"`
}
