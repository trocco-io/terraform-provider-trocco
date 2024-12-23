package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities"
	"terraform-provider-trocco/internal/client/parameters"
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

func NewCustomVariableSetting(customVariableSetting *entities.CustomVariableSetting) *CustomVariableSetting {
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

func NewCustomVariableSettings(customVariableSettings *[]entities.CustomVariableSetting) *[]CustomVariableSetting {
	if customVariableSettings == nil {
		return nil
	}
	settings := make([]CustomVariableSetting, 0, len(*customVariableSettings))
	for _, setting := range *customVariableSettings {
		settings = append(settings, *NewCustomVariableSetting(&setting))
	}
	return &settings
}

func ToCustomVariableSettingInputs(settings *[]CustomVariableSetting) *[]parameters.CustomVariableSettingInput {
	if settings == nil {
		return nil
	}
	inputs := make([]parameters.CustomVariableSettingInput, 0, len(*settings))
	for _, setting := range *settings {
		inputs = append(inputs, parameters.CustomVariableSettingInput{
			Name:      setting.Name.ValueString(),
			Type:      setting.Type.ValueString(),
			Value:     setting.Value.ValueStringPointer(),
			Quantity:  &parameters.NullableInt64{Valid: !setting.Quantity.IsNull(), Value: setting.Quantity.ValueInt64()},
			Unit:      setting.Unit.ValueStringPointer(),
			Direction: setting.Direction.ValueStringPointer(),
			Format:    setting.Format.ValueStringPointer(),
			TimeZone:  setting.TimeZone.ValueStringPointer(),
		})
	}
	return &inputs
}

func CustomVariableEntitiesToModels(customVariables *[]entities.CustomVariableSetting) *[]CustomVariableSetting {
	if customVariables == nil {
		return nil
	}
	outputs := make([]CustomVariableSetting, 0, len(*customVariables))
	for _, setting := range *customVariables {
		outputs = append(outputs, CustomVariableSetting{
			Name:      types.StringValue(setting.Name),
			Type:      types.StringValue(setting.Type),
			Value:     types.StringPointerValue(setting.Value),
			Quantity:  types.Int64PointerValue(setting.Quantity),
			Unit:      types.StringPointerValue(setting.Unit),
			Direction: types.StringPointerValue(setting.Direction),
			Format:    types.StringPointerValue(setting.Format),
			TimeZone:  types.StringPointerValue(setting.TimeZone),
		})
	}
	return &outputs
}
