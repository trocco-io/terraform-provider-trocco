package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type SnowflakeCustomVariableLoopConfig struct {
	ConnectionID types.Int64  `tfsdk:"connection_id"`
	Query        types.String `tfsdk:"query"`
	Warehouse    types.String `tfsdk:"warehouse"`
	Variables    types.Set    `tfsdk:"variables"`
}

func NewSnowflakeCustomVariableLoopConfig(en *we.SnowflakeCustomVariableLoopConfig, ctx context.Context) *SnowflakeCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	elements := []attr.Value{}
	for _, v := range en.Variables {
		elements = append(elements, types.StringValue(v))
	}

	var variables types.Set
	if len(elements) == 0 {
		variables = types.SetNull(types.StringType)
	} else {
		set, _ := types.SetValue(types.StringType, elements)
		variables = set
	}

	return &SnowflakeCustomVariableLoopConfig{
		ConnectionID: types.Int64Value(en.ConnectionID),
		Query:        types.StringValue(en.Query),
		Warehouse:    types.StringValue(en.Warehouse),
		Variables:    variables,
	}
}

func (c *SnowflakeCustomVariableLoopConfig) ToInput(ctx context.Context) wp.SnowflakeCustomVariableLoopConfig {
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

	return wp.SnowflakeCustomVariableLoopConfig{
		ConnectionID: c.ConnectionID.ValueInt64(),
		Query:        c.Query.ValueString(),
		Warehouse:    c.Warehouse.ValueString(),
		Variables:    vs,
	}
}

func SnowflakeCustomVariableLoopConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"connection_id": types.Int64Type,
		"query":         types.StringType,
		"warehouse":     types.StringType,
		"variables":     types.SetType{ElemType: types.StringType},
	}
}
