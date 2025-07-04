package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type RedshiftCustomVariableLoopConfig struct {
	ConnectionID types.Int64    `tfsdk:"connection_id"`
	Query        types.String   `tfsdk:"query"`
	Database     types.String   `tfsdk:"database"`
	Variables    []types.String `tfsdk:"variables"`
}

func NewRedshiftCustomVariableLoopConfig(en *we.RedshiftCustomVariableLoopConfig) *RedshiftCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	vs := []types.String{}
	for _, v := range en.Variables {
		vs = append(vs, types.StringValue(v))
	}

	return &RedshiftCustomVariableLoopConfig{
		ConnectionID: types.Int64Value(en.ConnectionID),
		Query:        types.StringValue(en.Query),
		Database:     types.StringValue(en.Database),
		Variables:    vs,
	}
}

func (c *RedshiftCustomVariableLoopConfig) ToInput() wp.RedshiftCustomVariableLoopConfig {
	vs := []string{}
	for _, v := range c.Variables {
		vs = append(vs, v.ValueString())
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
