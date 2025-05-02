package input_options

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	param "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type KintoneInputOption struct {
	AppID                  types.String                   `tfsdk:"app_id"`
	GuestSpaceID           types.String                   `tfsdk:"guest_space_id"`
	ExpandSubtable         types.Bool                     `tfsdk:"expand_subtable"`
	Query                  types.String                   `tfsdk:"query"`
	KintoneConnectionID    types.Int64                    `tfsdk:"kintone_connection_id"`
	InputOptionColumns     []KintoneInputOptionColumn     `tfsdk:"input_option_columns"`
	CustomVariableSettings *[]model.CustomVariableSetting `tfsdk:"custom_variable_settings"`
}

type KintoneInputOptionColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func NewKintoneInputOption(inputOption *input_option.KintoneInputOption) *KintoneInputOption {
	if inputOption == nil {
		return nil
	}

	return &KintoneInputOption{
		AppID:                  types.StringValue(inputOption.AppID),
		GuestSpaceID:           types.StringPointerValue(inputOption.GuestSpaceID),
		ExpandSubtable:         types.BoolValue(inputOption.ExpandSubtable),
		Query:                  types.StringPointerValue(inputOption.Query),
		KintoneConnectionID:    types.Int64Value(inputOption.KintoneConnectionID),
		InputOptionColumns:     newKintoneInputOptionColumns(inputOption.InputOptionColumns),
		CustomVariableSettings: model.NewCustomVariableSettings(inputOption.CustomVariableSettings),
	}
}

func newKintoneInputOptionColumns(inputOptionColumns []input_option.KintoneInputOptionColumn) []KintoneInputOptionColumn {
	if inputOptionColumns == nil {
		return nil
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
	return columns
}

func (inputOption *KintoneInputOption) ToInput() *param.KintoneInputOptionInput {
	if inputOption == nil {
		return nil
	}

	return &param.KintoneInputOptionInput{
		AppID:                  inputOption.AppID.ValueString(),
		GuestSpaceID:           model.NewNullableString(inputOption.GuestSpaceID),
		ExpandSubtable:         model.NewNullableBool(inputOption.ExpandSubtable),
		Query:                  model.NewNullableString(inputOption.Query),
		KintoneConnectionID:    inputOption.KintoneConnectionID.ValueInt64(),
		InputOptionColumns:     toKintoneInputOptionColumnsInput(inputOption.InputOptionColumns),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(inputOption.CustomVariableSettings),
	}
}

func (inputOption *KintoneInputOption) ToUpdateInput() *param.UpdateKintoneInputOptionInput {
	if inputOption == nil {
		return nil
	}

	return &param.UpdateKintoneInputOptionInput{
		AppID:                  model.NewNullableString(inputOption.AppID),
		GuestSpaceID:           model.NewNullableString(inputOption.GuestSpaceID),
		ExpandSubtable:         model.NewNullableBool(inputOption.ExpandSubtable),
		Query:                  model.NewNullableString(inputOption.Query),
		KintoneConnectionID:    model.NewNullableInt64(inputOption.KintoneConnectionID),
		InputOptionColumns:     toKintoneInputOptionColumnsInput(inputOption.InputOptionColumns),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(inputOption.CustomVariableSettings),
	}
}

func toKintoneInputOptionColumnsInput(columns []KintoneInputOptionColumn) []param.KintoneInputOptionColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]param.KintoneInputOptionColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, param.KintoneInputOptionColumn{
			Name:   column.Name.ValueString(),
			Type:   column.Type.ValueString(),
			Format: column.Format.ValueStringPointer(),
		})
	}
	return inputs
}
