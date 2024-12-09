package workflow

type StringCustomVariableLoopConfig struct {
	Variables []StringCustomVariableLoopVariable `json:"variables"`
}

type StringCustomVariableLoopVariable struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}
