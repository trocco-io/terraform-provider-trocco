package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
)

type TroccoAgentTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`
}

func NewTroccoAgentTaskConfig(c *we.TroccoAgentTaskConfig) *TroccoAgentTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoAgentTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),
	}
}

func (c *TroccoAgentTaskConfig) ToInput() *wp.TroccoAgentTaskConfig {
	return &wp.TroccoAgentTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}
}
