package workflow

type StringCustomVariableLoopConfig struct {
	Variables []StringCustomVariableLoopVariable `json:"variables,omitempty"`
}

type StringCustomVariableLoopVariable struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}
