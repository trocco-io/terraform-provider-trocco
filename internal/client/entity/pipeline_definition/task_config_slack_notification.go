package pipeline_definition

type SlackNotificationTaskConfig struct {
	Name         string `json:"name"`
	ConnectionID int64  `json:"connection_id"`
	Message      string `json:"message"`
	IgnoreError  bool   `json:"ignore_error"`
}
