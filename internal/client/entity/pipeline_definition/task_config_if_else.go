package pipeline_definition

type IfElseTaskConfig struct {
	Name            string           `json:"name"`
	ConditionGroups *ConditionGroups `json:"condition_groups"`
	Destinations    *Destinations    `json:"destinations"`
}

type ConditionGroups struct {
	SetType    string       `json:"set_type"`
	Conditions []*Condition `json:"conditions"`
}

type Condition struct {
	Identifier string `json:"identifier"`
	Variable   string `json:"variable"`
	Operator   string `json:"operator"`
	Value      string `json:"value"`
}

type Destinations struct {
	If   []string `json:"if"`
	Else []string `json:"else"`
}
