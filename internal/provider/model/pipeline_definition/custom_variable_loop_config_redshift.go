package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type RedshiftCustomVariableLoopConfig struct {
	ConnectionID types.Int64  `tfsdk:"connection_id"`
	Query        types.String `tfsdk:"query"`
	Database     types.String `tfsdk:"database"`
	Variables    types.List   `tfsdk:"variables"`
}

func NewRedshiftCustomVariableLoopConfig(en *we.RedshiftCustomVariableLoopConfig) *RedshiftCustomVariableLoopConfig {
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

	return &RedshiftCustomVariableLoopConfig{
		ConnectionID: types.Int64Value(en.ConnectionID),
		Query:        types.StringValue(en.Query),
		Database:     types.StringValue(en.Database),
		Variables:    variablesList,
	}
}

func (c *RedshiftCustomVariableLoopConfig) ToInput() wp.RedshiftCustomVariableLoopConfig {
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

	return wp.RedshiftCustomVariableLoopConfig{
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
		"variables":     types.ListType{ElemType: types.StringType},
	}
}
