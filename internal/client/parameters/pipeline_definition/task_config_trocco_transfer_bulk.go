package pipeline_definition

import (
	p "terraform-provider-trocco/internal/client/parameters"
)

type TroccoTransferBulkTaskConfig struct {
	DefinitionID      int64            `json:"definition_id,omitempty"`
	IsStoppedOnErrors *p.NullableBool  `json:"is_stopped_on_errors,omitempty"`
	MaxErrors         *p.NullableInt64 `json:"max_errors,omitempty"`
}
