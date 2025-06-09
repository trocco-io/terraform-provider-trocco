package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type BigqueryCustomVariableLoopConfig struct {
	ConnectionID types.Int64  `tfsdk:"connection_id"`
	Query        types.String `tfsdk:"query"`
	Variables    types.List   `tfsdk:"variables"`
}

func NewBigqueryCustomVariableLoopConfig(en *we.BigqueryCustomVariableLoopConfig) *BigqueryCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	vs := []types.String{}
	for _, v := range en.Variables {
		vs = append(vs, types.StringValue(v))
	}

	variablesList, _ := types.ListValueFrom(
		context.Background(),
		types.StringType,
		vs,
	)

	return &BigqueryCustomVariableLoopConfig{
		ConnectionID: types.Int64Value(en.ConnectionID),
		Query:        types.StringValue(en.Query),
		Variables:    variablesList,
	}
}

func (c *BigqueryCustomVariableLoopConfig) ToInput() wp.BigqueryCustomVariableLoopConfig {
	vs := []string{}

	if !c.Variables.IsNull() && !c.Variables.IsUnknown() {
		var stringValues []types.String
		diags := c.Variables.ElementsAs(context.Background(), &stringValues, false)
		if !diags.HasError() {
			for _, v := range stringValues {
				vs = append(vs, v.ValueString())
			}
		}
	}

	return wp.BigqueryCustomVariableLoopConfig{
		ConnectionID: c.ConnectionID.ValueInt64(),
		Query:        c.Query.ValueString(),
		Variables:    vs,
	}
}
