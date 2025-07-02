package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type TroccoTransferTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `tfsdk:"custom_variable_loop"`
}

func NewTroccoTransferTaskConfig(c *we.TroccoTransferTaskConfig, ctx context.Context) *TroccoTransferTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoTransferTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),

		CustomVariableLoop: NewCustomVariableLoop(c.CustomVariableLoop, ctx),
	}
}

func (c *TroccoTransferTaskConfig) ToInput(ctx context.Context) *wp.TroccoTransferTaskConfig {
	in := &wp.TroccoTransferTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}

	if c.CustomVariableLoop != nil {
		in.CustomVariableLoop = lo.ToPtr(c.CustomVariableLoop.ToInput(ctx))
	}

	return in
}

func TroccoTransferTaskConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"definition_id": types.Int64Type,
		"custom_variable_loop": types.ObjectType{
			AttrTypes: CustomVariableLoopAttrTypes(),
		},
	}
}
