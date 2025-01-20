package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type TroccoSnowflakeDatamartTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `tfsdk:"custom_variable_loop"`
}

func NewTroccoSnowflakeDatamartTaskConfig(c *we.TroccoSnowflakeDatamartTaskConfig) *TroccoSnowflakeDatamartTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoSnowflakeDatamartTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),

		CustomVariableLoop: NewCustomVariableLoop(c.CustomVariableLoop),
	}
}

func (c *TroccoSnowflakeDatamartTaskConfig) ToInput() *wp.TroccoSnowflakeDatamartTaskConfig {
	in := &wp.TroccoSnowflakeDatamartTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}

	if c.CustomVariableLoop != nil {
		in.CustomVariableLoop = lo.ToPtr(c.CustomVariableLoop.ToInput())
	}

	return in
}
