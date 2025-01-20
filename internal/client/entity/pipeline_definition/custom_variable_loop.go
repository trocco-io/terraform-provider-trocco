package pipeline_definition

type CustomVariableLoop struct {
	Type                       string `json:"type"`
	IsParallelExecutionAllowed *bool  `json:"is_parallel_execution_allowed"`
	IsStoppedOnErrors          *bool  `json:"is_stopped_on_errors"`
	MaxErrors                  *int64 `json:"max_errors"`

	StringConfig    *StringCustomVariableLoopConfig    `json:"string_config"`
	PeriodConfig    *PeriodCustomVariableLoopConfig    `json:"period_config"`
	BigqueryConfig  *BigqueryCustomVariableLoopConfig  `json:"bigquery_config"`
	SnowflakeConfig *SnowflakeCustomVariableLoopConfig `json:"snowflake_config"`
	RedshiftConfig  *RedshiftCustomVariableLoopConfig  `json:"redshift_config"`
}
