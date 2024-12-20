package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities"
)

type CustomVariableSetting struct {
	Name      types.String `tfsdk:"name"`
	Type      types.String `tfsdk:"type"`
	Value     types.String `tfsdk:"value"`
	Quantity  types.Int32  `tfsdk:"quantity"`
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
		Quantity:  types.Int32PointerValue(customVariableSetting.Quantity),
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
