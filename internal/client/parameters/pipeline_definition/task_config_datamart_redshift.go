package workflow

type TroccoRedshiftDatamartTaskConfig struct {
	DefinitionID int64 `json:"definition_id,omitempty"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop,omitempty"`
}