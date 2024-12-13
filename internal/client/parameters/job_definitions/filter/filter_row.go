package filter

type FilterRowsInput struct {
	Condition           string                    `json:"condition"`
	FilterRowConditions []FilterRowConditionInput `json:"filter_row_conditions"`
}

type FilterRowConditionInput struct {
	Column   string `json:"column"`
	Operator string `json:"operator"`
	Argument string `json:"argument"`
}
