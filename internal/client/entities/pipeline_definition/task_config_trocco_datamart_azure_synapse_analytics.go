package pipeline_definition

type TroccoAzureSynapseAnalyticsDatamartTaskConfig struct {
	DefinitionID int64 `json:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `json:"custom_variable_loop"`
}
