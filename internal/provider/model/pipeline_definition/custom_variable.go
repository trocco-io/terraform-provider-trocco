package pipeline_definition

import (
	"context"
	pipelineDefinitionEntities "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	parameter "terraform-provider-trocco/internal/client/parameter"
	pipelineDefinitionParameters "terraform-provider-trocco/internal/client/parameter/pipeline_definition"

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

func NewCustomVariables(ctx context.Context, ens []pipelineDefinitionEntities.CustomVariable) types.Set {
	objectType := types.ObjectType{
		AttrTypes: CustomVariableAttrTypes(),
	}

	if len(ens) == 0 {
		// If no custom variables are present, the API returns an empty array but the provider should set `null`.
		return types.SetNull(objectType)
	}

	var mds []CustomVariable
	for _, en := range ens {
		mds = append(mds, NewCustomVariable(en))
	}

	setValue, diags := types.SetValueFrom(ctx, objectType, mds)
	if diags.HasError() {
		return types.SetNull(objectType)
	}

	return setValue
}

func NewCustomVariable(en pipelineDefinitionEntities.CustomVariable) CustomVariable {
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

func (v *CustomVariable) ToInput() pipelineDefinitionParameters.CustomVariable {
	return pipelineDefinitionParameters.CustomVariable{
		Name:      v.Name.ValueStringPointer(),
		Type:      v.Type.ValueStringPointer(),
		Value:     v.Value.ValueStringPointer(),
		Quantity:  &parameter.NullableInt64{Valid: !v.Quantity.IsNull(), Value: v.Quantity.ValueInt64()},
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
