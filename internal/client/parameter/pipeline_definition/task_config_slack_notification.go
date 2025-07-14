package pipeline_definition

import (
	parameter "terraform-provider-trocco/internal/client/parameter"
)

type SlackNotificationTaskConfig struct {
	Name         string                  `json:"name,omitempty"`
	ConnectionID int64                   `json:"connection_id,omitempty"`
	Message      string                  `json:"message,omitempty"`
	IgnoreError  *parameter.NullableBool `json:"ignore_error,omitempty"`
}
