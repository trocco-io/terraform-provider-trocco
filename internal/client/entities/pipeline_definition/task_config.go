package pipeline_definition

type TroccoBigqueryDatamartTaskConfig struct {
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

type BigqueryDataCheckTaskConfig struct {
	Name            string           `json:"name"`
	ConnectionID    int64            `json:"connection_id"`
	Query           string           `json:"query"`
	Operator        string           `json:"operator"`
	QueryResult     int64            `json:"query_result"`
	AcceptsNull     bool             `json:"accepts_null"`
	CustomVariables []CustomVariable `json:"custom_variables"`
}

type SnowflakeDataCheckTaskConfig struct {
	Name            string           `json:"name"`
	ConnectionID    int64            `json:"connection_id"`
	Query           string           `json:"query"`
	Operator        string           `json:"operator"`
	QueryResult     int64            `json:"query_result"`
	AcceptsNull     bool             `json:"accepts_null"`
	Warehouse       string           `json:"warehouse"`
	CustomVariables []CustomVariable `json:"custom_variables"`
}

type RedshiftDataCheckTaskConfig struct {
	Name            string           `json:"name"`
	ConnectionID    int64            `json:"connection_id"`
	Query           string           `json:"query"`
	Operator        string           `json:"operator"`
	QueryResult     int64            `json:"query_result"`
	AcceptsNull     bool             `json:"accepts_null"`
	Database        string           `json:"database"`
	CustomVariables []CustomVariable `json:"custom_variables"`
}

type HTTPRequestTaskConfig struct {
	Name              string             `json:"name"`
	ConnectionID      *int64             `json:"connection_id"`
	HTTPMethod        string             `json:"http_method"`
	URL               string             `json:"url"`
	RequestBody       *string            `json:"request_body"`
	RequestHeaders    []RequestHeader    `json:"request_headers"`
	RequestParameters []RequestParameter `json:"request_parameters"`
	CustomVariables   []CustomVariable   `json:"custom_variables"`
}

type RequestHeader struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Masking bool   `json:"masking"`
}

type RequestParameter struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Masking bool   `json:"masking"`
}
