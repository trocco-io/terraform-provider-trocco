package pipeline_definition

type CustomVariableLoop struct {
	Type string `json:"type"`

	StringConfig    *StringCustomVariableLoopConfig    `json:"string_config"`
	PeriodConfig    *PeriodCustomVariableLoopConfig    `json:"period_config"`
	BigqueryConfig  *BigqueryCustomVariableLoopConfig  `json:"bigquery_config"`
	SnowflakeConfig *SnowflakeCustomVariableLoopConfig `json:"snowflake_config"`
	RedshiftConfig  *RedshiftCustomVariableLoopConfig  `json:"redshift_config"`
}
