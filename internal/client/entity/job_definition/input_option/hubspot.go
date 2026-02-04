package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type HubspotInputOption struct {
	HubspotConnectionID       int64                           `json:"hubspot_connection_id"`
	Target                    string                          `json:"target"`
	FromObjectType            *string                         `json:"from_object_type"`
	ToObjectType              *string                         `json:"to_object_type"`
	ObjectType                *string                         `json:"object_type"`
	IncrementalLoadingEnabled *bool                           `json:"incremental_loading_enabled"`
	LastRecordTime            *string                         `json:"last_record_time"`
	EmailEventType            *string                         `json:"email_event_type"`
	StartTimestamp            *string                         `json:"start_timestamp"`
	EndTimestamp              *string                         `json:"end_timestamp"`
	InputOptionColumns        []HubspotColumn                 `json:"input_option_columns"`
	CustomVariableSettings    *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type HubspotColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
