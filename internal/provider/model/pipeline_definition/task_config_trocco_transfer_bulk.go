package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	pipelineDefinitionEntities "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	parameter "terraform-provider-trocco/internal/client/parameter"
	pipelineDefinitionParameters "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type TroccoTransferBulkTaskConfig struct {
	DefinitionID               types.Int64 `tfsdk:"definition_id"`
	IsParallelExecutionAllowed types.Bool  `tfsdk:"is_parallel_execution_allowed"`
	IsStoppedOnErrors          types.Bool  `tfsdk:"is_stopped_on_errors"`
	MaxErrors                  types.Int64 `tfsdk:"max_errors"`
}

func NewTroccoTransferBulkTaskConfig(en *pipelineDefinitionEntities.TroccoTransferBulkTaskConfig) *TroccoTransferBulkTaskConfig {
	if en == nil {
		return nil
	}

	return &TroccoTransferBulkTaskConfig{
		DefinitionID:               types.Int64Value(en.DefinitionID),
		IsParallelExecutionAllowed: types.BoolPointerValue(en.IsParallelExecutionAllowed),
		IsStoppedOnErrors:          types.BoolPointerValue(en.IsStoppedOnErrors),
		MaxErrors:                  types.Int64PointerValue(en.MaxErrors),
	}
}

func (c *TroccoTransferBulkTaskConfig) ToInput() *pipelineDefinitionParameters.TroccoTransferBulkTaskConfig {
	return &pipelineDefinitionParameters.TroccoTransferBulkTaskConfig{
		DefinitionID:               c.DefinitionID.ValueInt64(),
		IsParallelExecutionAllowed: &parameter.NullableBool{Valid: !c.IsParallelExecutionAllowed.IsNull(), Value: c.IsParallelExecutionAllowed.ValueBool()},
		IsStoppedOnErrors:          &parameter.NullableBool{Valid: !c.IsStoppedOnErrors.IsNull(), Value: c.IsStoppedOnErrors.ValueBool()},
		MaxErrors:                  &parameter.NullableInt64{Valid: !c.MaxErrors.IsNull(), Value: c.MaxErrors.ValueInt64()},
	}
}

func TroccoTransferBulkTaskConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"definition_id":                 types.Int64Type,
		"is_parallel_execution_allowed": types.BoolType,
		"is_stopped_on_errors":          types.BoolType,
		"max_errors":                    types.Int64Type,
	}
}
