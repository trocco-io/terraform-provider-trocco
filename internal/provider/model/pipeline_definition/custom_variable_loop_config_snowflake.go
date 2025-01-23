package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type SnowflakeCustomVariableLoopConfig struct {
	ConnectionID types.Int64    `tfsdk:"connection_id"`
	Query        types.String   `tfsdk:"query"`
	Warehouse    types.String   `tfsdk:"warehouse"`
	Variables    []types.String `tfsdk:"variables"`
}

func NewSnowflakeCustomVariableLoopConfig(en *we.SnowflakeCustomVariableLoopConfig) *SnowflakeCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	vs := []types.String{}
	for _, v := range en.Variables {
		vs = append(vs, types.StringValue(v))
	}

	return &SnowflakeCustomVariableLoopConfig{
		ConnectionID: types.Int64Value(en.ConnectionID),
		Query:        types.StringValue(en.Query),
		Warehouse:    types.StringValue(en.Warehouse),
		Variables:    vs,
	}
}

func (c *SnowflakeCustomVariableLoopConfig) ToInput() wp.SnowflakeCustomVariableLoopConfig {
	vs := []string{}
	for _, v := range c.Variables {
		vs = append(vs, v.ValueString())
	}

	return wp.SnowflakeCustomVariableLoopConfig{
		ConnectionID: c.ConnectionID.ValueInt64(),
		Query:        c.Query.ValueString(),
		Warehouse:    c.Warehouse.ValueString(),
		Variables:    vs,
	}
}
