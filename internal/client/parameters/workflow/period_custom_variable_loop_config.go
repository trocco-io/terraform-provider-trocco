package workflow

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
