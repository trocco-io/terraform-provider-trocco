package workflow

//
// CustomVariableLoop
//

type CustomVariableLoop struct {
	Type string `json:"type"`

	StringConfig    *StringCustomVariableLoopConfig    `json:"string_config"`
	PeriodConfig    *PeriodCustomVariableLoopConfig    `json:"period_config"`
	BigqueryConfig  *BigqueryCustomVariableLoopConfig  `json:"bigquery_config"`
	SnowflakeConfig *SnowflakeCustomVariableLoopConfig `json:"snowflake_config"`
	RedshiftConfig  *RedshiftCustomVariableLoopConfig  `json:"redshift_config"`
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
	Interval  string                             `json:"interval"`
	TimeZone  string                             `json:"time_zone"`
	From      PeriodCustomVariableLoopFrom       `json:"from"`
	To        PeriodCustomVariableLoopTo         `json:"to"`
	Variables []PeriodCustomVariableLoopVariable `json:"variables"`
}

type PeriodCustomVariableLoopFrom struct {
	Value int64  `json:"value"`
	Unit  string `json:"unit"`
}

type PeriodCustomVariableLoopTo struct {
	Value int64  `json:"value"`
	Unit  string `json:"unit"`
}

type PeriodCustomVariableLoopVariable struct {
	Name   string                                 `json:"name"`
	Offset PeriodCustomVariableLoopVariableOffset `json:"offset"`
}

type PeriodCustomVariableLoopVariableOffset struct {
	Value int64  `json:"value"`
	Unit  string `json:"unit"`
}

//
// BigqueryCustomVariableLoopConfig
//

type BigqueryCustomVariableLoopConfig struct {
	ConnectionID int64    `json:"connection_id"`
	Query        string   `json:"query"`
	Variables    []string `json:"variables"`
}

//
// RedshiftCustomVariableLoopConfig
//

type RedshiftCustomVariableLoopConfig struct {
	ConnectionID int64    `json:"connection_id"`
	Query        string   `json:"query"`
	Database     string   `json:"database"`
	Variables    []string `json:"variables"`
}

//
// SnowflakeCustomVariableLoopConfig
//

type SnowflakeCustomVariableLoopConfig struct {
	ConnectionID int64    `json:"connection_id"`
	Query        string   `json:"query"`
	Warehouse    string   `json:"warehouse"`
	Variables    []string `json:"variables"`
}
