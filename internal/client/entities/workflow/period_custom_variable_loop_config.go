package workflow

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
