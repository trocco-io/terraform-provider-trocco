package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type BigqueryCustomVariableLoopConfig struct {
	ConnectionID types.Int64    `tfsdk:"connection_id"`
	Query        types.String   `tfsdk:"query"`
	Variables    []types.String `tfsdk:"variables"`
}

func NewBigqueryCustomVariableLoopConfig(en *we.BigqueryCustomVariableLoopConfig) *BigqueryCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	vs := []types.String{}
	for _, v := range en.Variables {
		vs = append(vs, types.StringValue(v))
	}

	return &BigqueryCustomVariableLoopConfig{
		ConnectionID: types.Int64Value(en.ConnectionID),
		Query:        types.StringValue(en.Query),
		Variables:    vs,
	}
}

func (c *BigqueryCustomVariableLoopConfig) ToInput() wp.BigqueryCustomVariableLoopConfig {
	vs := []string{}
	for _, v := range c.Variables {
		vs = append(vs, v.ValueString())
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
