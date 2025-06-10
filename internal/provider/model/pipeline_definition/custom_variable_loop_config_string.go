package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type StringCustomVariableLoopConfig struct {
	Variables types.List `tfsdk:"variables"`
}

func NewStringCustomVariableLoopConfig(en *we.StringCustomVariableLoopConfig) *StringCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	variables := []StringCustomVariableLoopVariable{}
	for _, variable := range en.Variables {
		variables = append(variables, NewStringCustomVariableLoopVariable(variable))
	}

	listType := types.ObjectType{
		AttrTypes: StringCustomVariableLoopVariableAttrTypes(),
	}
	variablesList, diags := types.ListValueFrom(context.Background(), listType, variables)
	if diags.HasError() {
		return nil
	}

	return &StringCustomVariableLoopConfig{
		Variables: variablesList,
	}
}

func (c *StringCustomVariableLoopConfig) ToInput() wp.StringCustomVariableLoopConfig {
	vs := []wp.StringCustomVariableLoopVariable{}

	if !c.Variables.IsNull() && !c.Variables.IsUnknown() {
		var variables []StringCustomVariableLoopVariable
		diags := c.Variables.ElementsAs(context.Background(), &variables, false)
		if !diags.HasError() {
			for _, v := range variables {
				vs = append(vs, v.ToInput())
			}
		}
	}

	return wp.StringCustomVariableLoopConfig{
		Variables: vs,
	}
}

type StringCustomVariableLoopVariable struct {
	Name   types.String `tfsdk:"name"`
	Values types.List   `tfsdk:"values"`
}

func NewStringCustomVariableLoopVariable(en we.StringCustomVariableLoopVariable) StringCustomVariableLoopVariable {
	values := []types.String{}
	for _, val := range en.Values {
		values = append(values, types.StringValue(val))
	}

	valuesList, _ := types.ListValueFrom(
		context.Background(),
		types.StringType,
		values,
	)

	return StringCustomVariableLoopVariable{
		Name:   types.StringValue(en.Name),
		Values: valuesList,
	}
}

func (v *StringCustomVariableLoopVariable) ToInput() wp.StringCustomVariableLoopVariable {
	values := []string{}

	if !v.Values.IsNull() && !v.Values.IsUnknown() {
		var stringValues []types.String
		diags := v.Values.ElementsAs(context.Background(), &stringValues, false)
		if !diags.HasError() {
			for _, val := range stringValues {
				values = append(values, val.ValueString())
			}
		}
	}

	return wp.StringCustomVariableLoopVariable{
		Name:   v.Name.ValueString(),
		Values: values,
	}
}

func StringCustomVariableLoopVariableAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   types.StringType,
		"values": types.ListType{ElemType: types.StringType},
	}
}

func StringCustomVariableLoopConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"variables": types.ListType{
			ElemType: types.ObjectType{AttrTypes: StringCustomVariableLoopVariableAttrTypes()},
		},
	}
}
