package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type PeriodCustomVariableLoopConfig struct {
	Interval  types.String                 `tfsdk:"interval"`
	TimeZone  types.String                 `tfsdk:"time_zone"`
	From      PeriodCustomVariableLoopFrom `tfsdk:"from"`
	To        PeriodCustomVariableLoopTo   `tfsdk:"to"`
	Variables types.List                   `tfsdk:"variables"`
}

func NewPeriodCustomVariableLoopConfig(ctx context.Context, en *we.PeriodCustomVariableLoopConfig) *PeriodCustomVariableLoopConfig {
	if en == nil {
		return nil
	}

	variables := []PeriodCustomVariableLoopVariable{}
	for _, variable := range en.Variables {
		variables = append(variables, NewPeriodCustomVariableLoopVariable(variable))
	}

	variablesList, diags := types.ListValueFrom(
		ctx,
		types.ObjectType{AttrTypes: PeriodCustomVariableLoopVariableAttrTypes()},
		variables,
	)
	if diags.HasError() {
		return nil
	}

	return &PeriodCustomVariableLoopConfig{
		Interval:  types.StringValue(en.Interval),
		TimeZone:  types.StringValue(en.TimeZone),
		From:      NewPeriodCustomVariableLoopFrom(en.From),
		To:        NewPeriodCustomVariableLoopTo(en.To),
		Variables: variablesList,
	}
}

func (c *PeriodCustomVariableLoopConfig) ToInput(ctx context.Context) wp.PeriodCustomVariableLoopConfig {
	vars := []wp.PeriodCustomVariableLoopVariable{}

	var variables []PeriodCustomVariableLoopVariable
	diags := c.Variables.ElementsAs(ctx, &variables, false)
	if diags.HasError() {
		return wp.PeriodCustomVariableLoopConfig{}
	}

	for _, v := range variables {
		vars = append(vars, v.ToInput())
	}

	return wp.PeriodCustomVariableLoopConfig{
		Interval:  c.Interval.ValueString(),
		TimeZone:  c.TimeZone.ValueString(),
		From:      c.From.ToInput(),
		To:        c.To.ToInput(),
		Variables: vars,
	}
}

type PeriodCustomVariableLoopFrom struct {
	Value types.Int64  `tfsdk:"value"`
	Unit  types.String `tfsdk:"unit"`
}

func NewPeriodCustomVariableLoopFrom(en we.PeriodCustomVariableLoopFrom) PeriodCustomVariableLoopFrom {
	return PeriodCustomVariableLoopFrom{
		Value: types.Int64Value(en.Value),
		Unit:  types.StringValue(en.Unit),
	}
}

func (f *PeriodCustomVariableLoopFrom) ToInput() wp.PeriodCustomVariableLoopFrom {
	return wp.PeriodCustomVariableLoopFrom{
		Value: f.Value.ValueInt64Pointer(),
		Unit:  f.Unit.ValueString(),
	}
}

type PeriodCustomVariableLoopTo struct {
	Value types.Int64  `tfsdk:"value"`
	Unit  types.String `tfsdk:"unit"`
}

func NewPeriodCustomVariableLoopTo(en we.PeriodCustomVariableLoopTo) PeriodCustomVariableLoopTo {
	return PeriodCustomVariableLoopTo{
		Value: types.Int64Value(en.Value),
		Unit:  types.StringValue(en.Unit),
	}
}

func (t *PeriodCustomVariableLoopTo) ToInput() wp.PeriodCustomVariableLoopTo {
	return wp.PeriodCustomVariableLoopTo{
		Value: t.Value.ValueInt64Pointer(),
		Unit:  t.Unit.ValueString(),
	}
}

type PeriodCustomVariableLoopVariable struct {
	Name   types.String                           `tfsdk:"name"`
	Offset PeriodCustomVariableLoopVariableOffset `tfsdk:"offset"`
}

func NewPeriodCustomVariableLoopVariable(en we.PeriodCustomVariableLoopVariable) PeriodCustomVariableLoopVariable {
	return PeriodCustomVariableLoopVariable{
		Name:   types.StringValue(en.Name),
		Offset: NewStringCustomVariableLoopVariableOffset(en.Offset),
	}
}

func (v *PeriodCustomVariableLoopVariable) ToInput() wp.PeriodCustomVariableLoopVariable {
	return wp.PeriodCustomVariableLoopVariable{
		Name:   v.Name.ValueString(),
		Offset: v.Offset.ToInput(),
	}
}

type PeriodCustomVariableLoopVariableOffset struct {
	Value types.Int64  `tfsdk:"value"`
	Unit  types.String `tfsdk:"unit"`
}

func NewStringCustomVariableLoopVariableOffset(en we.PeriodCustomVariableLoopVariableOffset) PeriodCustomVariableLoopVariableOffset {
	return PeriodCustomVariableLoopVariableOffset{
		Value: types.Int64Value(en.Value),
		Unit:  types.StringValue(en.Unit),
	}
}

func (o *PeriodCustomVariableLoopVariableOffset) ToInput() wp.PeriodCustomVariableLoopVariableOffset {
	return wp.PeriodCustomVariableLoopVariableOffset{
		Value: o.Value.ValueInt64Pointer(),
		Unit:  o.Unit.ValueString(),
	}
}
func PeriodCustomVariableLoopConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"interval":  types.StringType,
		"time_zone": types.StringType,
		"from": types.ObjectType{
			AttrTypes: PeriodCustomVariableLoopFromAttrTypes(),
		},
		"to": types.ObjectType{
			AttrTypes: PeriodCustomVariableLoopToAttrTypes(),
		},
		"variables": types.ListType{
			ElemType: types.ObjectType{
				AttrTypes: PeriodCustomVariableLoopVariableAttrTypes(),
			},
		},
	}
}

func PeriodCustomVariableLoopVariableAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
		"offset": types.ObjectType{
			AttrTypes: PeriodCustomVariableLoopVariableOffsetAttrTypes(),
		},
	}
}

func PeriodCustomVariableLoopVariableOffsetAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"value": types.Int64Type,
		"unit":  types.StringType,
	}
}

func PeriodCustomVariableLoopFromAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"value": types.Int64Type,
		"unit":  types.StringType,
	}
}

func PeriodCustomVariableLoopToAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"value": types.Int64Type,
		"unit":  types.StringType,
	}
}
