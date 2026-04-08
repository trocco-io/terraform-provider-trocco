package output_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type RedshiftOutputOption struct {
	RedshiftConnectionID              int64                              `json:"redshift_connection_id"`
	Database                          string                             `json:"database"`
	Schema                            string                             `json:"schema"`
	Table                             string                             `json:"table"`
	CreateTableConstraint             *string                            `json:"create_table_constraint"`
	CreateTableOption                 *string                            `json:"create_table_option"`
	S3Bucket                          string                             `json:"s3_bucket"`
	S3KeyPrefix                       string                             `json:"s3_key_prefix"`
	DeleteS3TempFile                  bool                               `json:"delete_s3_temp_file"`
	CopyIAMRoleName                   *string                            `json:"copy_iam_role_name"`
	RetryLimit                        int64                              `json:"retry_limit"`
	RetryWait                         int64                              `json:"retry_wait"`
	MaxRetryWait                      int64                              `json:"max_retry_wait"`
	Mode                              string                             `json:"mode"`
	DefaultTimeZone                   string                             `json:"default_time_zone"`
	BeforeLoad                        *string                            `json:"before_load"`
	AfterLoad                         *string                            `json:"after_load"`
	BatchSize                         int64                              `json:"batch_size"`
	RedshiftOutputOptionColumnOptions []RedshiftOutputOptionColumnOption `json:"redshift_output_option_column_options"`
	RedshiftOutputOptionMergeKeys     []string                           `json:"redshift_output_option_merge_keys"`
	CustomVariableSettings            *[]entity.CustomVariableSetting    `json:"custom_variable_settings"`
}

type RedshiftOutputOptionColumnOption struct {
	Name            string  `json:"name"`
	Type            *string `json:"type"`
	ValueType       *string `json:"value_type"`
	TimestampFormat *string `json:"timestamp_format"`
	Timezone        *string `json:"timezone"`
}
