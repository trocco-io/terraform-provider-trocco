package pipeline_definition

import (
	p "terraform-provider-trocco/internal/client/parameters"
)

type SnowflakeDataCheckTaskConfigInput struct {
	Name            string           `json:"name,omitempty"`
	ConnectionID    int64            `json:"connection_id,omitempty"`
	Query           string           `json:"query,omitempty"`
	Operator        string           `json:"operator,omitempty"`
	QueryResult     *p.NullableInt64 `json:"query_result,omitempty"`
	AcceptsNull     *p.NullableBool  `json:"accepts_null,omitempty"`
	Warehouse       string           `json:"warehouse,omitempty"`
	CustomVariables []CustomVariable `json:"custom_variables,omitempty"`
}
