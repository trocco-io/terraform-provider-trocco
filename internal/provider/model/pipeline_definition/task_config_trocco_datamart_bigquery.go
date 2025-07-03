package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

//
// TroccoBigqueryDatamartTaskConfig
//

type TroccoBigqueryDatamartTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `tfsdk:"custom_variable_loop"`
}

func NewTroccoBigqueryDatamartTaskConfig(ctx context.Context, c *we.TroccoBigqueryDatamartTaskConfig) *TroccoBigqueryDatamartTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoBigqueryDatamartTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),

		CustomVariableLoop: NewCustomVariableLoop(ctx, c.CustomVariableLoop),
	}
}

func (c *TroccoBigqueryDatamartTaskConfig) ToInput(ctx context.Context) *wp.TroccoBigqueryDatamartTaskConfig {
	in := &wp.TroccoBigqueryDatamartTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}

	if c.CustomVariableLoop != nil {
		in.CustomVariableLoop = lo.ToPtr(c.CustomVariableLoop.ToInput(ctx))
	}

	return in
}

func TroccoBigqueryDatamartTaskConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"definition_id": types.Int64Type,
		"custom_variable_loop": types.ObjectType{
			AttrTypes: CustomVariableLoopAttrTypes(),
		},
	}
}
