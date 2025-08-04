package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	pipelineDefinitionEntities "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	pipelineDefinitionParameters "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type RedshiftCustomVariableLoopConfig struct {
	ConnectionID types.Int64  `tfsdk:"connection_id"`
	Query        types.String `tfsdk:"query"`
	Database     types.String `tfsdk:"database"`
	Variables    types.Set    `tfsdk:"variables"`
}

func NewRedshiftCustomVariableLoopConfig(ctx context.Context, en *pipelineDefinitionEntities.RedshiftCustomVariableLoopConfig) *RedshiftCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	variables, diags := types.SetValueFrom(ctx, types.StringType, en.Variables)
	if diags.HasError() {
		variables = types.SetNull(types.StringType)
	}

	return &RedshiftCustomVariableLoopConfig{
		ConnectionID: types.Int64Value(en.ConnectionID),
		Query:        types.StringValue(en.Query),
		Database:     types.StringValue(en.Database),
		Variables:    variables,
	}
}

func (c *RedshiftCustomVariableLoopConfig) ToInput(ctx context.Context) pipelineDefinitionParameters.RedshiftCustomVariableLoopConfig {
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

	return pipelineDefinitionParameters.RedshiftCustomVariableLoopConfig{
		ConnectionID: c.ConnectionID.ValueInt64(),
		Query:        c.Query.ValueString(),
		Database:     c.Database.ValueString(),
		Variables:    vs,
	}
}

func RedshiftCustomVariableLoopConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"connection_id": types.Int64Type,
		"query":         types.StringType,
		"database":      types.StringType,
		"variables":     types.SetType{ElemType: types.StringType},
	}
}
