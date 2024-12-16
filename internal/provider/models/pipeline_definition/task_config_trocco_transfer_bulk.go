package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
)

type TroccoTransferBulkTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`
}

func NewTroccoTransferBulkTaskConfig(c *we.TroccoTransferBulkTaskConfig) *TroccoTransferBulkTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoTransferBulkTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),
	}
}

func (c *TroccoTransferBulkTaskConfig) ToInput() *wp.TroccoTransferBulkTaskConfig {
	return &wp.TroccoTransferBulkTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}
}
