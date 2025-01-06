package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"

	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
)

type TroccoAzureSynapseAnalyticsDatamartTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `tfsdk:"custom_variable_loop"`
}

func NewTroccoAzureSynapseAnalyticsDatamartTaskConfig(c *we.TroccoAzureSynapseAnalyticsDatamartTaskConfig) *TroccoAzureSynapseAnalyticsDatamartTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoAzureSynapseAnalyticsDatamartTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),

		CustomVariableLoop: NewCustomVariableLoop(c.CustomVariableLoop),
	}
}

func (c *TroccoAzureSynapseAnalyticsDatamartTaskConfig) ToInput() *wp.TroccoAzureSynapseAnalyticsDatamartTaskConfig {
	in := &wp.TroccoAzureSynapseAnalyticsDatamartTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}

	if c.CustomVariableLoop != nil {
		in.CustomVariableLoop = lo.ToPtr(c.CustomVariableLoop.ToInput())
	}

	return in
}
