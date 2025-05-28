package output_options

import (
	"context"
	"terraform-provider-trocco/internal/client/entity"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func GetCustomVariableSettingAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":      types.StringType,
		"type":      types.StringType,
		"value":     types.StringType,
		"quantity":  types.Int32Type,
		"unit":      types.StringType,
		"direction": types.StringType,
		"format":    types.StringType,
		"time_zone": types.StringType,
	}
}

func ConvertCustomVariableSettingsToList(ctx context.Context, customVariableSettings *[]entity.CustomVariableSetting) types.List {
	if customVariableSettings == nil {
		return types.ListNull(types.ObjectType{
			AttrTypes: GetCustomVariableSettingAttrTypes(),
		})
	}

	settings := model.NewCustomVariableSettings(customVariableSettings)
	if settings == nil {
		return types.ListNull(types.ObjectType{
			AttrTypes: GetCustomVariableSettingAttrTypes(),
		})
	}

	objectType := types.ObjectType{
		AttrTypes: GetCustomVariableSettingAttrTypes(),
	}
	listValue, _ := types.ListValueFrom(ctx, objectType, *settings)
	return listValue
}

func ExtractCustomVariableSettings(ctx context.Context, list types.List) *[]model.CustomVariableSetting {
	if list.IsNull() && list.IsUnknown() {
		return nil
	}

	var settings []model.CustomVariableSetting
	list.ElementsAs(ctx, &settings, false)
	return &settings
}
