package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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

	variablesList, diags := types.ListValueFrom(
		context.Background(),
		types.StringType,
		vs,
	)
	if diags.HasError() {
		return nil
	}

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
			for _, val := range stringValues {
				vs = append(vs, val.ValueString())
			}
		}
	}

	return wp.BigqueryCustomVariableLoopConfig{
		ConnectionID: c.ConnectionID.ValueInt64(),
		Query:        c.Query.ValueString(),
		Variables:    vs,
	}
}

func BigqueryCustomVariableLoopConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"connection_id": types.Int64Type,
		"query":         types.StringType,
		"variables":     types.ListType{ElemType: types.StringType},
	}
}
