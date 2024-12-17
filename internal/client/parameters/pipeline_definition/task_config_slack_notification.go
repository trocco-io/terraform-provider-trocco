package workflow

type SlackNotificationTaskConfig struct {
	Name         string `json:"name,omitempty"`
	ConnectionID int64  `json:"connection_id,omitempty"`
	Message      string `json:"message,omitempty"`
}
