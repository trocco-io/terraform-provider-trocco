package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type StringCustomVariableLoopConfig struct {
	Variables []StringCustomVariableLoopVariable `tfsdk:"variables"`
}

func NewStringCustomVariableLoopConfig(en *we.StringCustomVariableLoopConfig) *StringCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	variables := []StringCustomVariableLoopVariable{}
	for _, variable := range en.Variables {
		variables = append(variables, NewStringCustomVariableLoopVariable(variable))
	}

	return &StringCustomVariableLoopConfig{
		Variables: variables,
	}
}

func (c *StringCustomVariableLoopConfig) ToInput() wp.StringCustomVariableLoopConfig {
	vs := []wp.StringCustomVariableLoopVariable{}
	for _, v := range c.Variables {
		vs = append(vs, v.ToInput())
	}

	return wp.StringCustomVariableLoopConfig{
		Variables: vs,
	}
}

type StringCustomVariableLoopVariable struct {
	Name   types.String   `tfsdk:"name"`
	Values []types.String `tfsdk:"values"`
}

func NewStringCustomVariableLoopVariable(en we.StringCustomVariableLoopVariable) StringCustomVariableLoopVariable {
	values := []types.String{}
	for _, val := range en.Values {
		values = append(values, types.StringValue(val))
	}

	return StringCustomVariableLoopVariable{
		Name:   types.StringValue(en.Name),
		Values: values,
	}
}

func (v *StringCustomVariableLoopVariable) ToInput() wp.StringCustomVariableLoopVariable {
	values := []string{}
	for _, val := range v.Values {
		values = append(values, val.ValueString())
	}

	return wp.StringCustomVariableLoopVariable{
		Name:   v.Name.ValueString(),
		Values: values,
	}
}

func StringCustomVariableLoopConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"variables": types.ListType{
			ElemType: types.ObjectType{AttrTypes: StringCustomVariableLoopVariableAttrTypes()},
		},
	}
}

func StringCustomVariableLoopVariableAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   types.StringType,
		"values": types.ListType{ElemType: types.StringType},
	}
}
