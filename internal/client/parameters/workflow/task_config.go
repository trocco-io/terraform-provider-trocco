package workflow

type TroccoTransferTaskConfig struct {
	DefinitionID int64 `json:"definition_id,omitempty"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop,omitempty"`
}

type SlackNotificationTaskConfig struct {
	Name         string `json:"name,omitempty"`
	ConnectionID int64  `json:"connection_id,omitempty"`
	Message      string `json:"message,omitempty"`
}
