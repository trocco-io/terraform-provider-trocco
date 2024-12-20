package pipeline_definition

type TroccoTransferBulkTaskConfig struct {
	DefinitionID      int64  `json:"definition_id"`
	IsStoppedOnErrors *bool  `json:"is_stopped_on_errors"`
	MaxErrors         *int64 `json:"max_errors"`
}
