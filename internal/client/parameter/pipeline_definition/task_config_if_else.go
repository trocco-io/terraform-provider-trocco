package pipeline_definition

type IfElseTaskConfig struct {
	Name            string           `json:"name"`
	ConditionGroups *ConditionGroups `json:"condition_groups,omitempty"`
	Destinations    *Destinations    `json:"destinations"`
}

type ConditionGroups struct {
	SetType    string       `json:"set_type"`
	Conditions []*Condition `json:"conditions"`
}

type Condition struct {
	Variable string  `json:"variable"`
	TaskKey  string `json:"task_key"`
	Operator string  `json:"operator"`
	Value    string  `json:"value"`
}

type Destinations struct {
	If   []string `json:"if,omitempty"`
	Else []string `json:"else,omitempty"`
}
