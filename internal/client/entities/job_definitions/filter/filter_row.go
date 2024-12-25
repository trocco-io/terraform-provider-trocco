package filter

type FilterRows struct {
	Condition           string               `json:"condition"`
	FilterRowConditions []FilterRowCondition `json:"filter_row_conditions"`
}

type FilterRowCondition struct {
	Column   string `json:"column"`
	Operator string `json:"operator"`
	Argument string `json:"argument"`
}
