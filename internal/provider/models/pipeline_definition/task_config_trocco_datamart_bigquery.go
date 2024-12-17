package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"

	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
)

//
// TroccoBigQueryDatamartTaskConfig
//

type TroccoBigqueryDatamartTaskConfig struct {
	DefinitionID types.Int64 `tfsdk:"definition_id"`

	CustomVariableLoop *CustomVariableLoop `tfsdk:"custom_variable_loop"`
}

func NewTroccoBigQueryDatamartTaskConfig(c *we.TroccoBigQueryDatamartTaskConfig) *TroccoBigqueryDatamartTaskConfig {
	if c == nil {
		return nil
	}

	return &TroccoBigqueryDatamartTaskConfig{
		DefinitionID: types.Int64Value(c.DefinitionID),

		CustomVariableLoop: NewCustomVariableLoop(c.CustomVariableLoop),
	}
}

func (c *TroccoBigqueryDatamartTaskConfig) ToInput() *wp.TroccoBigQueryDatamartTaskConfig {
	in := &wp.TroccoBigQueryDatamartTaskConfig{
		DefinitionID: c.DefinitionID.ValueInt64(),
	}

	if c.CustomVariableLoop != nil {
		in.CustomVariableLoop = lo.ToPtr(c.CustomVariableLoop.ToInput())
	}

	return in
}
