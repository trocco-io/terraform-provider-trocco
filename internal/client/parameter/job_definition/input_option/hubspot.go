package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type HubspotInputOptionInput struct {
	HubspotConnectionID       int64                                   `json:"hubspot_connection_id"`
	Target                    string                                  `json:"target"`
	FromObjectType            *string                                 `json:"from_object_type,omitempty"`
	ToObjectType              *string                                 `json:"to_object_type,omitempty"`
	ObjectType                *string                                 `json:"object_type,omitempty"`
	IncrementalLoadingEnabled *bool                                   `json:"incremental_loading_enabled,omitempty"`
	EmailEventType            *string                                 `json:"email_event_type,omitempty"`
	StartTimestamp            *string                                 `json:"start_timestamp,omitempty"`
	EndTimestamp              *string                                 `json:"end_timestamp,omitempty"`
	InputOptionColumns        []HubspotColumn                         `json:"input_option_columns"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateHubspotInputOptionInput struct {
	HubspotConnectionID       *int64                                  `json:"hubspot_connection_id,omitempty"`
	Target                    *string                                 `json:"target,omitempty"`
	FromObjectType            *string                                 `json:"from_object_type,omitempty"`
	ToObjectType              *string                                 `json:"to_object_type,omitempty"`
	ObjectType                *string                                 `json:"object_type,omitempty"`
	IncrementalLoadingEnabled *bool                                   `json:"incremental_loading_enabled,omitempty"`
	EmailEventType            *string                                 `json:"email_event_type,omitempty"`
	StartTimestamp            *string                                 `json:"start_timestamp,omitempty"`
	EndTimestamp              *string                                 `json:"end_timestamp,omitempty"`
	InputOptionColumns        []HubspotColumn                         `json:"input_option_columns,omitempty"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type HubspotColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
