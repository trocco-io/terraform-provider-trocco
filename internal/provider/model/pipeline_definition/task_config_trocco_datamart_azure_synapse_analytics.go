package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type TroccoAzureSynapseAnalyticsDatamartTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `tfsdk:"custom_variable_loop"`
}

func NewTroccoAzureSynapseAnalyticsDatamartTaskConfig(c *we.TroccoAzureSynapseAnalyticsDatamartTaskConfig, ctx context.Context) *TroccoAzureSynapseAnalyticsDatamartTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoAzureSynapseAnalyticsDatamartTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),

		CustomVariableLoop: NewCustomVariableLoop(c.CustomVariableLoop, ctx),
	}
}

func (c *TroccoAzureSynapseAnalyticsDatamartTaskConfig) ToInput(ctx context.Context) *wp.TroccoAzureSynapseAnalyticsDatamartTaskConfig {
	in := &wp.TroccoAzureSynapseAnalyticsDatamartTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}

	if c.CustomVariableLoop != nil {
		in.CustomVariableLoop = lo.ToPtr(c.CustomVariableLoop.ToInput(ctx))
	}

	return in
}

func TroccoAzureSynapseAnalyticsDatamartTaskConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"definition_id": types.Int64Type,
		"custom_variable_loop": types.ObjectType{
			AttrTypes: CustomVariableLoopAttrTypes(),
		},
	}
}
