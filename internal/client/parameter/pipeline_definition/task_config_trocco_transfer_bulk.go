package pipeline_definition

import (
	parameter "terraform-provider-trocco/internal/client/parameter"
)

type TroccoTransferBulkTaskConfig struct {
	DefinitionID               int64                    `json:"definition_id,omitempty"`
	IsParallelExecutionAllowed *parameter.NullableBool  `json:"is_parallel_execution_allowed,omitempty"`
	IsStoppedOnErrors          *parameter.NullableBool  `json:"is_stopped_on_errors,omitempty"`
	MaxErrors                  *parameter.NullableInt64 `json:"max_errors,omitempty"`
}
