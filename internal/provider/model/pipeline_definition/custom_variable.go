package pipeline_definition

import (
	"context"
	"fmt"
	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameter"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomVariable struct {
	Name      types.String `tfsdk:"name"`
	Type      types.String `tfsdk:"type"`
	Value     types.String `tfsdk:"value"`
	Quantity  types.Int64  `tfsdk:"quantity"`
	Unit      types.String `tfsdk:"unit"`
	Direction types.String `tfsdk:"direction"`
	Format    types.String `tfsdk:"format"`
	TimeZone  types.String `tfsdk:"time_zone"`
}

func NewCustomVariables(ens []we.CustomVariable) (types.Set, error) {
	if len(ens) == 0 {
		return types.SetNull(types.ObjectType{AttrTypes: CustomVariableAttrTypes()}), nil
	}

	var mds []CustomVariable
	for _, en := range ens {
		mds = append(mds, NewCustomVariable(en))
	}

	ctx := context.Background()

	customVariables, diags := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: CustomVariableAttrTypes()}, mds)
	if diags.HasError() {
		return types.SetNull(types.ObjectType{AttrTypes: CustomVariableAttrTypes()}), fmt.Errorf("failed to convert customVariables to SetValue: %v", diags)
	}
	return customVariables, nil
}

func NewCustomVariable(en we.CustomVariable) CustomVariable {
	return CustomVariable{
		Name:      types.StringPointerValue(en.Name),
		Type:      types.StringPointerValue(en.Type),
		Value:     types.StringPointerValue(en.Value),
		Quantity:  types.Int64PointerValue(en.Quantity),
		Unit:      types.StringPointerValue(en.Unit),
		Direction: types.StringPointerValue(en.Direction),
		Format:    types.StringPointerValue(en.Format),
		TimeZone:  types.StringPointerValue(en.TimeZone),
	}
}

func (v *CustomVariable) ToInput() wp.CustomVariable {
	return wp.CustomVariable{
		Name:      v.Name.ValueStringPointer(),
		Type:      v.Type.ValueStringPointer(),
		Value:     v.Value.ValueStringPointer(),
		Quantity:  &p.NullableInt64{Valid: !v.Quantity.IsNull(), Value: v.Quantity.ValueInt64()},
		Unit:      v.Unit.ValueStringPointer(),
		Direction: v.Direction.ValueStringPointer(),
		Format:    v.Format.ValueStringPointer(),
		TimeZone:  v.TimeZone.ValueStringPointer(),
	}
}

func CustomVariableAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":      types.StringType,
		"type":      types.StringType,
		"value":     types.StringType,
		"quantity":  types.Int64Type,
		"unit":      types.StringType,
		"direction": types.StringType,
		"format":    types.StringType,
		"time_zone": types.StringType,
	}
}

func CustomVariablesToInput(set types.Set) []wp.CustomVariable {
	if set.IsNull() {
		return nil
	}

	var tfVars []CustomVariable
	diags := set.ElementsAs(context.Background(), &tfVars, false)
	if diags.HasError() {
		return nil
	}

	inputs := make([]wp.CustomVariable, 0, len(tfVars))
	for _, v := range tfVars {
		inputs = append(inputs, v.ToInput())
	}
	return inputs
}
