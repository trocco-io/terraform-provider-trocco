package pipeline_definition

import (
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

func NewTroccoBigqueryDatamartTaskConfig(c *we.TroccoBigqueryDatamartTaskConfig) *TroccoBigqueryDatamartTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoBigqueryDatamartTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),

		CustomVariableLoop: NewCustomVariableLoop(c.CustomVariableLoop),
	}
}

func (c *TroccoBigqueryDatamartTaskConfig) ToInput() *wp.TroccoBigqueryDatamartTaskConfig {
	in := &wp.TroccoBigqueryDatamartTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}

	if c.CustomVariableLoop != nil {
		in.CustomVariableLoop = lo.ToPtr(c.CustomVariableLoop.ToInput())
	}

	return in
}
