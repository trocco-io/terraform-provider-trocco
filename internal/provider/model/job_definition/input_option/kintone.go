package input_options

import (
	"context"
	"fmt"
	inputOptionEntities "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	inputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type KintoneInputOption struct {
	AppID                  types.String `tfsdk:"app_id"`
	GuestSpaceID           types.String `tfsdk:"guest_space_id"`
	ExpandSubtable         types.Bool   `tfsdk:"expand_subtable"`
	Query                  types.String `tfsdk:"query"`
	KintoneConnectionID    types.Int64  `tfsdk:"kintone_connection_id"`
	InputOptionColumns     types.List   `tfsdk:"input_option_columns"`
	CustomVariableSettings types.List   `tfsdk:"custom_variable_settings"`
}

type KintoneInputOptionColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func (KintoneInputOptionColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   types.StringType,
		"type":   types.StringType,
		"format": types.StringType,
	}
}

func NewKintoneInputOption(ctx context.Context, inputOption *inputOptionEntities.KintoneInputOption) *KintoneInputOption {
	if inputOption == nil {
		return nil
	}

	result := &KintoneInputOption{
		AppID:               types.StringValue(inputOption.AppID),
		GuestSpaceID:        types.StringPointerValue(inputOption.GuestSpaceID),
		ExpandSubtable:      types.BoolValue(inputOption.ExpandSubtable),
		Query:               types.StringPointerValue(inputOption.Query),
		KintoneConnectionID: types.Int64Value(inputOption.KintoneConnectionID),
	}

	columns, err := newKintoneInputOptionColumns(ctx, inputOption.InputOptionColumns)
	if err != nil {
		return nil
	}
	result.InputOptionColumns = columns

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, inputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func newKintoneInputOptionColumns(
	ctx context.Context,
	inputOptionColumns []inputOptionEntities.KintoneInputOptionColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: KintoneInputOptionColumn{}.attrTypes(),
	}

	if inputOptionColumns == nil {
		return types.ListNull(objectType), nil
	}

	columns := make([]KintoneInputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := KintoneInputOptionColumn{
			Name:   types.StringValue(input.Name),
			Type:   types.StringValue(input.Type),
			Format: types.StringPointerValue(input.Format),
		}
		columns = append(columns, column)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columns)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert input option columns to ListValue: %v", diags)
	}
	return listValue, nil
}

func (inputOption *KintoneInputOption) ToInput(ctx context.Context) *inputOptionParameters.KintoneInputOptionInput {
	if inputOption == nil {
		return nil
	}

	var columnValues []KintoneInputOptionColumn
	if !inputOption.InputOptionColumns.IsNull() && !inputOption.InputOptionColumns.IsUnknown() {
		diags := inputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	return &inputOptionParameters.KintoneInputOptionInput{
		AppID:                  inputOption.AppID.ValueString(),
		GuestSpaceID:           model.NewNullableString(inputOption.GuestSpaceID),
		ExpandSubtable:         model.NewNullableBool(inputOption.ExpandSubtable),
		Query:                  model.NewNullableString(inputOption.Query),
		KintoneConnectionID:    inputOption.KintoneConnectionID.ValueInt64(),
		InputOptionColumns:     toKintoneInputOptionColumnsInput(columnValues),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (inputOption *KintoneInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateKintoneInputOptionInput {
	if inputOption == nil {
		return nil
	}

	var columnValues []KintoneInputOptionColumn
	if !inputOption.InputOptionColumns.IsNull() {
		if !inputOption.InputOptionColumns.IsUnknown() {
			diags := inputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnValues = []KintoneInputOptionColumn{}
		}
	} else {
		columnValues = []KintoneInputOptionColumn{}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	return &inputOptionParameters.UpdateKintoneInputOptionInput{
		AppID:                  model.NewNullableString(inputOption.AppID),
		GuestSpaceID:           model.NewNullableString(inputOption.GuestSpaceID),
		ExpandSubtable:         model.NewNullableBool(inputOption.ExpandSubtable),
		Query:                  model.NewNullableString(inputOption.Query),
		KintoneConnectionID:    model.NewNullableInt64(inputOption.KintoneConnectionID),
		InputOptionColumns:     toKintoneInputOptionColumnsInput(columnValues),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func toKintoneInputOptionColumnsInput(columns []KintoneInputOptionColumn) []inputOptionParameters.KintoneInputOptionColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.KintoneInputOptionColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, inputOptionParameters.KintoneInputOptionColumn{
			Name:   column.Name.ValueString(),
			Type:   column.Type.ValueString(),
			Format: column.Format.ValueStringPointer(),
		})
	}
	return inputs
}
