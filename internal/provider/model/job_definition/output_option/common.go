package output_options

import (
	"context"
	"fmt"
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

func ConvertCustomVariableSettingsToList(ctx context.Context, customVariableSettings *[]entity.CustomVariableSetting) (types.List, error) {
	if customVariableSettings == nil {
		return types.ListNull(types.ObjectType{
			AttrTypes: GetCustomVariableSettingAttrTypes(),
		}), nil
	}

	settings := model.NewCustomVariableSettings(customVariableSettings)
	if settings == nil {
		return types.ListNull(types.ObjectType{
			AttrTypes: GetCustomVariableSettingAttrTypes(),
		}), nil
	}

	objectType := types.ObjectType{
		AttrTypes: GetCustomVariableSettingAttrTypes(),
	}
	listValue, diags := types.ListValueFrom(ctx, objectType, *settings)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert to ListValue: %v", diags)
	}
	return listValue, nil
}

func ExtractCustomVariableSettings(ctx context.Context, list types.List) *[]model.CustomVariableSetting {
	if list.IsNull() && list.IsUnknown() {
		return nil
	}

	var settings []model.CustomVariableSetting
	diags := list.ElementsAs(ctx, &settings, false)
	if diags.HasError() {
		return nil
	}
	return &settings
}
