package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	pipelineDefinitionEntities "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	pipelineDefinitionParameters "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type TroccoDBTTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`
}

func NewTroccoDBTTaskConfig(c *pipelineDefinitionEntities.TroccoDBTTaskConfig) *TroccoDBTTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoDBTTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),
	}
}

func (c *TroccoDBTTaskConfig) ToInput() *pipelineDefinitionParameters.TroccoDBTTaskConfig {
	return &pipelineDefinitionParameters.TroccoDBTTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}
}

func TroccoDBTTaskConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"definition_id": types.Int64Type,
	}
}
