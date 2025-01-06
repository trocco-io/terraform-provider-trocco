package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameters"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
)

type TroccoTransferBulkTaskConfig struct {
	DefinitionID               types.Int64 `tfsdk:"definition_id"`
	IsParallelExecutionAllowed types.Bool  `tfsdk:"is_parallel_execution_allowed"`
	IsStoppedOnErrors          types.Bool  `tfsdk:"is_stopped_on_errors"`
	MaxErrors                  types.Int64 `tfsdk:"max_errors"`
}

func NewTroccoTransferBulkTaskConfig(en *we.TroccoTransferBulkTaskConfig) *TroccoTransferBulkTaskConfig {
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

func (c *TroccoTransferBulkTaskConfig) ToInput() *wp.TroccoTransferBulkTaskConfig {
	return &wp.TroccoTransferBulkTaskConfig{
		DefinitionID:               c.DefinitionID.ValueInt64(),
		IsParallelExecutionAllowed: &p.NullableBool{Valid: !c.IsParallelExecutionAllowed.IsNull(), Value: c.IsParallelExecutionAllowed.ValueBool()},
		IsStoppedOnErrors:          &p.NullableBool{Valid: !c.IsStoppedOnErrors.IsNull(), Value: c.IsStoppedOnErrors.ValueBool()},
		MaxErrors:                  &p.NullableInt64{Valid: !c.MaxErrors.IsNull(), Value: c.MaxErrors.ValueInt64()},
	}
}
