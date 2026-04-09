package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type RedshiftOutputOptionInput struct {
	RedshiftConnectionID              int64                                                                `json:"redshift_connection_id"`
	Database                          string                                                               `json:"database"`
	Schema                            string                                                               `json:"schema"`
	Table                             string                                                               `json:"table"`
	CreateTableConstraint             *string                                                              `json:"create_table_constraint,omitempty"`
	CreateTableOption                 *string                                                              `json:"create_table_option,omitempty"`
	S3Bucket                          string                                                               `json:"s3_bucket"`
	S3KeyPrefix                       string                                                               `json:"s3_key_prefix"`
	DeleteS3TempFile                  bool                                                                 `json:"delete_s3_temp_file"`
	CopyIAMRoleName                   *string                                                              `json:"copy_iam_role_name,omitempty"`
	RetryLimit                        int64                                                                `json:"retry_limit"`
	RetryWait                         int64                                                                `json:"retry_wait"`
	MaxRetryWait                      int64                                                                `json:"max_retry_wait"`
	Mode                              string                                                               `json:"mode"`
	DefaultTimeZone                   string                                                               `json:"default_time_zone"`
	BeforeLoad                        *string                                                              `json:"before_load,omitempty"`
	AfterLoad                         *string                                                              `json:"after_load,omitempty"`
	BatchSize                         int64                                                                `json:"batch_size"`
	RedshiftOutputOptionColumnOptions *parameter.NullableObjectList[RedshiftOutputOptionColumnOptionInput] `json:"redshift_output_option_column_options,omitempty"`
	RedshiftOutputOptionMergeKeys     *parameter.NullableObjectList[string]                                `json:"redshift_output_option_merge_keys,omitempty"`
	CustomVariableSettings            *[]parameter.CustomVariableSettingInput                              `json:"custom_variable_settings,omitempty"`
}

type UpdateRedshiftOutputOptionInput struct {
	RedshiftConnectionID              *int64                                                               `json:"redshift_connection_id,omitempty"`
	Database                          *string                                                              `json:"database,omitempty"`
	Schema                            *string                                                              `json:"schema,omitempty"`
	Table                             *string                                                              `json:"table,omitempty"`
	CreateTableConstraint             *string                                                              `json:"create_table_constraint,omitempty"`
	CreateTableOption                 *string                                                              `json:"create_table_option,omitempty"`
	S3Bucket                          *string                                                              `json:"s3_bucket,omitempty"`
	S3KeyPrefix                       *string                                                              `json:"s3_key_prefix,omitempty"`
	DeleteS3TempFile                  *bool                                                                `json:"delete_s3_temp_file,omitempty"`
	CopyIAMRoleName                   *string                                                              `json:"copy_iam_role_name,omitempty"`
	RetryLimit                        *int64                                                               `json:"retry_limit,omitempty"`
	RetryWait                         *int64                                                               `json:"retry_wait,omitempty"`
	MaxRetryWait                      *int64                                                               `json:"max_retry_wait,omitempty"`
	Mode                              *string                                                              `json:"mode,omitempty"`
	DefaultTimeZone                   *string                                                              `json:"default_time_zone,omitempty"`
	BeforeLoad                        *string                                                              `json:"before_load,omitempty"`
	AfterLoad                         *string                                                              `json:"after_load,omitempty"`
	BatchSize                         *int64                                                               `json:"batch_size,omitempty"`
	RedshiftOutputOptionColumnOptions *parameter.NullableObjectList[RedshiftOutputOptionColumnOptionInput] `json:"redshift_output_option_column_options,omitempty"`
	RedshiftOutputOptionMergeKeys     *parameter.NullableObjectList[string]                                `json:"redshift_output_option_merge_keys,omitempty"`
	CustomVariableSettings            *[]parameter.CustomVariableSettingInput                              `json:"custom_variable_settings,omitempty"`
}

type RedshiftOutputOptionColumnOptionInput struct {
	Name            string  `json:"name"`
	Type            *string `json:"type,omitempty"`
	ValueType       *string `json:"value_type,omitempty"`
	TimestampFormat *string `json:"timestamp_format,omitempty"`
	Timezone        *string `json:"timezone,omitempty"`
}
