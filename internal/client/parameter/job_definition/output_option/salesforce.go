package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type SalesforceOutputOptionInput struct {
	Object                 string                    `json:"object"`
	ActionType             *parameter.NullableString `json:"action_type,omitempty"`
	ApiVersion             *parameter.NullableString `json:"api_version,omitempty"`
	UpsertKey              *parameter.NullableString `json:"upsert_key,omitempty"`
	IgnoreNulls            *parameter.NullableBool   `json:"ignore_nulls,omitempty"`
	ThrowIfFailed          *parameter.NullableBool   `json:"throw_if_failed,omitempty"`
	SalesforceConnectionId int64                     `json:"salesforce_connection_id"`
}

type UpdateSalesforceOutputOptionInput struct {
	Object                 *string                   `json:"object"`
	ActionType             *parameter.NullableString `json:"action_type,omitempty"`
	ApiVersion             *parameter.NullableString `json:"api_version,omitempty"`
	UpsertKey              *parameter.NullableString `json:"upsert_key,omitempty"`
	IgnoreNulls            *parameter.NullableBool   `json:"ignore_nulls,omitempty"`
	ThrowIfFailed          *parameter.NullableBool   `json:"throw_if_failed,omitempty"`
	SalesforceConnectionId *int64                    `json:"salesforce_connection_id"`
}
