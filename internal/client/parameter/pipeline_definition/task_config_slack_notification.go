package pipeline_definition

import (
	p "terraform-provider-trocco/internal/client/parameter"
)

type SlackNotificationTaskConfig struct {
	Name         string          `json:"name,omitempty"`
	ConnectionID int64           `json:"connection_id,omitempty"`
	Message      string          `json:"message,omitempty"`
	IgnoreError  *p.NullableBool `json:"ignore_error,omitempty"`
}
