package workflow

type TroccoTransferTaskConfig struct {
	DefinitionID int64 `json:"definition_id,omitempty"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop,omitempty"`
}

type TroccoTransferBulkTaskConfig struct {
	DefinitionID int64 `json:"definition_id,omitempty"`
}

type DBTTaskConfig struct {
	DefinitionID int64 `json:"definition_id,omitempty"`
}

type TroccoAgentTaskConfig struct {
	DefinitionID int64 `json:"definition_id,omitempty"`
}

type TroccoBigQueryDatamartTaskConfig struct {
	DefinitionID int64 `json:"definition_id,omitempty"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop,omitempty"`
}

type TroccoRedshiftDatamartTaskConfig struct {
	DefinitionID int64 `json:"definition_id,omitempty"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop,omitempty"`
}

type TroccoSnowflakeDatamartTaskConfig struct {
	DefinitionID int64 `json:"definition_id,omitempty"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop,omitempty"`
}

type WorkflowTaskConfig struct {
	DefinitionID int64 `json:"definition_id,omitempty"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop,omitempty"`
}

type SlackNotificationTaskConfig struct {
	Name         string `json:"name,omitempty"`
	ConnectionID int64  `json:"connection_id,omitempty"`
	Message      string `json:"message,omitempty"`
}

// type WorkflowHTTPRequestTaskConfig struct {
// 	Name              string                               `json:"name"`
// 	ConnectionID      *int64                               `json:"connection_id"`
// 	HTTPMethod        string                               `json:"http_method"`
// 	URL               string                               `json:"url"`
// 	RequestBody       *string                              `json:"request_body"`
// 	RequestHeaders    []WorkflowTaskRequestHeaderConfig    `json:"request_headers"`
// 	RequestParameters []WorkflowTaskRequestParameterConfig `json:"request_parameters"`
// 	CustomVariables   []WorkflowTaskCustomVariableConfig   `json:"custom_variables"`
// }

// type WorkflowTaskRequestHeaderConfig struct {
// 	Key     string `json:"key"`
// 	Value   string `json:"value"`
// 	Masking bool   `json:"masking"`
// }

// type WorkflowTaskRequestParameterConfig struct {
// 	Key     string `json:"key"`
// 	Value   string `json:"value"`
// 	Masking bool   `json:"masking"`
// }
