package pipeline_definition

import (
	parameter "terraform-provider-trocco/internal/client/parameter"
)

type RedshiftDataCheckTaskConfigInput struct {
	Name            string                   `json:"name,omitempty"`
	ConnectionID    int64                    `json:"connection_id,omitempty"`
	Query           string                   `json:"query,omitempty"`
	Operator        string                   `json:"operator,omitempty"`
	QueryResult     *parameter.NullableInt64 `json:"query_result,omitempty"`
	AcceptsNull     *parameter.NullableBool  `json:"accepts_null,omitempty"`
	Database        string                   `json:"database,omitempty"`
	CustomVariables []CustomVariable         `json:"custom_variables,omitempty"`
}
