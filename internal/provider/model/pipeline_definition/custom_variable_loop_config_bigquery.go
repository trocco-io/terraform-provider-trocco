package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	pipelineDefinitionEntities "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	pipelineDefinitionParameters "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type BigqueryCustomVariableLoopConfig struct {
	ConnectionID types.Int64  `tfsdk:"connection_id"`
	Query        types.String `tfsdk:"query"`
	Variables    types.Set    `tfsdk:"variables"`
}

func NewBigqueryCustomVariableLoopConfig(ctx context.Context, en *pipelineDefinitionEntities.BigqueryCustomVariableLoopConfig) *BigqueryCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	variables, diags := types.SetValueFrom(ctx, types.StringType, en.Variables)
	if diags.HasError() {
		variables = types.SetNull(types.StringType)
	}

	return &BigqueryCustomVariableLoopConfig{
		ConnectionID: types.Int64Value(en.ConnectionID),
		Query:        types.StringValue(en.Query),
		Variables:    variables,
	}
}

func (c *BigqueryCustomVariableLoopConfig) ToInput(ctx context.Context) pipelineDefinitionParameters.BigqueryCustomVariableLoopConfig {
	vs := []string{}
	if !c.Variables.IsNull() && !c.Variables.IsUnknown() {
		var variableValues []types.String
		diags := c.Variables.ElementsAs(ctx, &variableValues, false)
		if !diags.HasError() {
			for _, v := range variableValues {
				vs = append(vs, v.ValueString())
			}
		}
	}

	return pipelineDefinitionParameters.BigqueryCustomVariableLoopConfig{
		ConnectionID: c.ConnectionID.ValueInt64(),
		Query:        c.Query.ValueString(),
		Variables:    vs,
	}
}

func BigqueryCustomVariableLoopConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"connection_id": types.Int64Type,
		"query":         types.StringType,
		"variables":     types.SetType{ElemType: types.StringType},
	}
}
