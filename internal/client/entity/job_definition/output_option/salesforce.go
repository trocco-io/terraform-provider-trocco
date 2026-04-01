package output_option

type SalesforceOutputOption struct {
	Object                 string  `json:"object"`
	ActionType             string  `json:"action_type"`
	ApiVersion             string  `json:"api_version"`
	UpsertKey              *string `json:"upsert_key"`
	UpdateKey              *string `json:"update_key"`
	IgnoreNulls            bool    `json:"ignore_nulls"`
	ThrowIfFailed          bool    `json:"throw_if_failed"`
	BatchSize              int64   `json:"batch_size"`
	SalesforceConnectionId int64   `json:"salesforce_connection_id"`
}
