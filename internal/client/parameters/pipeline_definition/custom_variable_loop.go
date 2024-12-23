package pipeline_definition

import (
	p "terraform-provider-trocco/internal/client/parameters"
)

//
// CustomVariableLoop
//

type CustomVariableLoop struct {
	Type                       string           `json:"type,omitempty"`
	IsParallelExecutionAllowed *p.NullableBool  `json:"is_parallel_execution_allowed,omitempty"`
	IsStoppedOnErrors          *p.NullableBool  `json:"is_stopped_on_errors,omitempty"`
	MaxErrors                  *p.NullableInt64 `json:"max_errors,omitempty"`

	StringConfig    *StringCustomVariableLoopConfig    `json:"string_config,omitempty"`
	PeriodConfig    *PeriodCustomVariableLoopConfig    `json:"period_config,omitempty"`
	BigqueryConfig  *BigqueryCustomVariableLoopConfig  `json:"bigquery_config,omitempty"`
	SnowflakeConfig *SnowflakeCustomVariableLoopConfig `json:"snowflake_config,omitempty"`
	RedshiftConfig  *RedshiftCustomVariableLoopConfig  `json:"redshift_config,omitempty"`
}

//
// StringCustomVariableLoopConfig
//

type StringCustomVariableLoopConfig struct {
	Variables []StringCustomVariableLoopVariable `json:"variables,omitempty"`
}

type StringCustomVariableLoopVariable struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}

//
// PeriodCustomVariableLoopConfig
//

type PeriodCustomVariableLoopConfig struct {
	Interval  string                             `json:"interval,omitempty"`
	TimeZone  string                             `json:"time_zone,omitempty"`
	From      PeriodCustomVariableLoopFrom       `json:"from,omitempty"`
	To        PeriodCustomVariableLoopTo         `json:"to,omitempty"`
	Variables []PeriodCustomVariableLoopVariable `json:"variables,omitempty"`
}

type PeriodCustomVariableLoopFrom struct {
	Value *int64 `json:"value,omitempty"`
	Unit  string `json:"unit,omitempty"`
}

type PeriodCustomVariableLoopTo struct {
	Value *int64 `json:"value,omitempty"`
	Unit  string `json:"unit,omitempty"`
}

type PeriodCustomVariableLoopVariable struct {
	Name   string                                 `json:"name,omitempty"`
	Offset PeriodCustomVariableLoopVariableOffset `json:"offset,omitempty"`
}

type PeriodCustomVariableLoopVariableOffset struct {
	Value *int64 `json:"value,omitempty"`
	Unit  string `json:"unit,omitempty"`
}

//
// StringCustomVariableLoopConfig
//

type BigqueryCustomVariableLoopConfig struct {
	ConnectionID int64    `json:"connection_id,omitempty"`
	Query        string   `json:"query,omitempty"`
	Variables    []string `json:"variables,omitempty"`
}

//
// RedshiftCustomVariableLoopConfig
//

type RedshiftCustomVariableLoopConfig struct {
	ConnectionID int64    `json:"connection_id,omitempty"`
	Query        string   `json:"query,omitempty"`
	Database     string   `json:"database,omitempty"`
	Variables    []string `json:"variables,omitempty"`
}

//
// SnowflakeCustomVariableLoopConfig
//

type SnowflakeCustomVariableLoopConfig struct {
	ConnectionID int64    `json:"connection_id,omitempty"`
	Query        string   `json:"query,omitempty"`
	Warehouse    string   `json:"warehouse,omitempty"`
	Variables    []string `json:"variables,omitempty"`
}
