package pipeline_definition

type TroccoTransferBulkTaskConfig struct {
	DefinitionID               int64  `json:"definition_id"`
	IsParallelExecutionAllowed *bool  `json:"is_parallel_execution_allowed"`
	IsStoppedOnErrors          *bool  `json:"is_stopped_on_errors"`
	MaxErrors                  *int64 `json:"max_errors"`
}
