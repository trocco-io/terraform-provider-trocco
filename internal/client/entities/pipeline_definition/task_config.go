package workflow

type TroccoBigQueryDatamartTaskConfig struct {
	DefinitionID int64 `json:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop"`
}

type TroccoTransferTaskConfig struct {
	DefinitionID int64 `json:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop"`
}

type TroccoTransferBulkTaskConfig struct {
	DefinitionID int64 `json:"definition_id"`
}

type DBTTaskConfig struct {
	DefinitionID int64 `json:"definition_id"`
}

type TroccoAgentTaskConfig struct {
	DefinitionID int64 `json:"definition_id"`
}

type TroccoRedshiftDatamartTaskConfig struct {
	DefinitionID int64 `json:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop"`
}

type TroccoSnowflakeDatamartTaskConfig struct {
	DefinitionID int64 `json:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop"`
}

type WorkflowTaskConfig struct {
	DefinitionID int64 `json:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop"`
}

type SlackNotificationTaskConfig struct {
	Name         string `json:"name"`
	ConnectionID int64  `json:"connection_id"`
	Message      string `json:"message"`
}

type TableauDataExtractionTaskConfig struct {
	Name         string `json:"name"`
	ConnectionID int64  `json:"connection_id"`
	TaskID       string `json:"task_id"`
}