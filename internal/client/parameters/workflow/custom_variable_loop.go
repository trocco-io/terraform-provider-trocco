package workflow

type CustomVariableLoop struct {
	Type string `json:"type,omitempty"`

	StringConfig    *StringCustomVariableLoopConfig    `json:"string_config,omitempty"`
	PeriodConfig    *PeriodCustomVariableLoopConfig    `json:"period_config,omitempty"`
	BigqueryConfig  *BigqueryCustomVariableLoopConfig  `json:"bigquery_config,omitempty"`
	SnowflakeConfig *SnowflakeCustomVariableLoopConfig `json:"snowflake_config,omitempty"`
	RedshiftConfig  *RedshiftCustomVariableLoopConfig  `json:"redshift_config,omitempty"`
}
