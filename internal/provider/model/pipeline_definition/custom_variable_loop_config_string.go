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

func NewStringCustomVariableLoopConfig(ctx context.Context, en *we.StringCustomVariableLoopConfig) *StringCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	variables := []StringCustomVariableLoopVariable{}
	for _, variable := range en.Variables {
		variables = append(variables, NewStringCustomVariableLoopVariable(ctx, variable))
	}

	variablesList, diags := types.ListValueFrom(
		ctx,
		types.ObjectType{AttrTypes: StringCustomVariableLoopVariableAttrTypes()},
		variables,
	)
	if diags.HasError() {
		return nil
	}

	return &StringCustomVariableLoopConfig{
		Variables: variablesList,
	}
}

func (c *StringCustomVariableLoopConfig) ToInput(ctx context.Context) wp.StringCustomVariableLoopConfig {
	vs := []wp.StringCustomVariableLoopVariable{}

	var variables []StringCustomVariableLoopVariable
	diags := c.Variables.ElementsAs(ctx, &variables, false)
	if diags.HasError() {
		return wp.StringCustomVariableLoopConfig{
			Variables: []wp.StringCustomVariableLoopVariable{},
		}
	}

	for _, v := range variables {
		vs = append(vs, v.ToInput(ctx))
	}

	return wp.StringCustomVariableLoopConfig{
		Variables: vs,
	}
}

type StringCustomVariableLoopVariable struct {
	Name   types.String `tfsdk:"name"`
	Values types.List   `tfsdk:"values"`
}

func NewStringCustomVariableLoopVariable(ctx context.Context, en we.StringCustomVariableLoopVariable) StringCustomVariableLoopVariable {
	values := []attr.Value{}
	for _, val := range en.Values {
		values = append(values, types.StringValue(val))
	}

	valuesList, diags := types.ListValueFrom(
		ctx,
		types.StringType,
		values,
	)
	if diags.HasError() {
		return StringCustomVariableLoopVariable{
			Name:   types.StringValue(en.Name),
			Values: types.ListNull(types.StringType),
		}
	}

	return StringCustomVariableLoopVariable{
		Name:   types.StringValue(en.Name),
		Values: valuesList,
	}
}

func (v *StringCustomVariableLoopVariable) ToInput(ctx context.Context) wp.StringCustomVariableLoopVariable {
	values := []string{}

	var stringValues []types.String
	diags := v.Values.ElementsAs(ctx, &stringValues, false)
	if diags.HasError() {
		// In a real application, you might want to handle this error differently
		// For now, we'll just return an empty list if there's an error
		return wp.StringCustomVariableLoopVariable{
			Name:   v.Name.ValueString(),
			Values: []string{},
		}
	}

	for _, val := range stringValues {
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
