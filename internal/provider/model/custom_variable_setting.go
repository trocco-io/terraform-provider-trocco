package model

import (
	"terraform-provider-trocco/internal/client/entity"
	"terraform-provider-trocco/internal/client/parameter"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomVariableSetting struct {
	Name      types.String `tfsdk:"name"`
	Type      types.String `tfsdk:"type"`
	Value     types.String `tfsdk:"value"`
	Quantity  types.Int64  `tfsdk:"quantity"`
	Unit      types.String `tfsdk:"unit"`
	Direction types.String `tfsdk:"direction"`
	Format    types.String `tfsdk:"format"`
	TimeZone  types.String `tfsdk:"time_zone"`
}

func NewCustomVariableSetting(customVariableSetting *entity.CustomVariableSetting) *CustomVariableSetting {
	if customVariableSetting == nil {
		return nil
	}
	return &CustomVariableSetting{
		Name:      types.StringValue(customVariableSetting.Name),
		Type:      types.StringValue(customVariableSetting.Type),
		Value:     types.StringPointerValue(customVariableSetting.Value),
		Quantity:  types.Int64PointerValue(customVariableSetting.Quantity),
		Unit:      types.StringPointerValue(customVariableSetting.Unit),
		Direction: types.StringPointerValue(customVariableSetting.Direction),
		Format:    types.StringPointerValue(customVariableSetting.Format),
		TimeZone:  types.StringPointerValue(customVariableSetting.TimeZone),
	}
}

func NewCustomVariableSettings(customVariableSettings *[]entity.CustomVariableSetting) *[]CustomVariableSetting {
	if customVariableSettings == nil {
		return nil
	}
	settings := make([]CustomVariableSetting, 0, len(*customVariableSettings))
	for _, setting := range *customVariableSettings {
		settings = append(settings, *NewCustomVariableSetting(&setting))
	}
	return &settings
}

func ToCustomVariableSettingInputs(settings *[]CustomVariableSetting) *[]parameter.CustomVariableSettingInput {
	if settings == nil {
		return nil
	}
	inputs := make([]parameter.CustomVariableSettingInput, 0, len(*settings))
	for _, setting := range *settings {
		inputs = append(inputs, parameter.CustomVariableSettingInput{
			Name:      setting.Name.ValueString(),
			Type:      setting.Type.ValueString(),
			Value:     setting.Value.ValueStringPointer(),
			Quantity:  setting.Quantity.ValueInt64Pointer(),
			Unit:      setting.Unit.ValueStringPointer(),
			Direction: setting.Direction.ValueStringPointer(),
			Format:    setting.Format.ValueStringPointer(),
			TimeZone:  setting.TimeZone.ValueStringPointer(),
		})
	}
	return &inputs
}
