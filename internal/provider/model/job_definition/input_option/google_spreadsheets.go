package input_options

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	param "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GoogleSpreadsheetsInputOption struct {
	SpreadsheetsURL                types.String `tfsdk:"spreadsheets_url"`
	WorksheetTitle                 types.String `tfsdk:"worksheet_title"`
	StartRow                       types.Int64  `tfsdk:"start_row"`
	StartColumn                    types.String `tfsdk:"start_column"`
	DefaultTimeZone                types.String `tfsdk:"default_time_zone"`
	NullString                     types.String `tfsdk:"null_string"`
	GoogleSpreadsheetsConnectionID types.Int64  `tfsdk:"google_spreadsheets_connection_id"`
	InputOptionColumns             types.List   `tfsdk:"input_option_columns"`
	CustomVariableSettings         types.List   `tfsdk:"custom_variable_settings"`
}

type GoogleSpreadsheetsInputOptionColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func (GoogleSpreadsheetsInputOptionColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   types.StringType,
		"type":   types.StringType,
		"format": types.StringType,
	}
}

func NewGoogleSpreadsheetsInputOption(inputOption *input_option.GoogleSpreadsheetsInputOption) *GoogleSpreadsheetsInputOption {
	if inputOption == nil {
		return nil
	}

	ctx := context.Background()

	result := &GoogleSpreadsheetsInputOption{
		SpreadsheetsURL:                types.StringValue(inputOption.SpreadsheetsURL),
		WorksheetTitle:                 types.StringValue(inputOption.WorksheetTitle),
		StartRow:                       types.Int64Value(inputOption.StartRow),
		StartColumn:                    types.StringValue(inputOption.StartColumn),
		DefaultTimeZone:                types.StringValue(inputOption.DefaultTimeZone),
		NullString:                     types.StringValue(inputOption.NullString),
		GoogleSpreadsheetsConnectionID: types.Int64Value(inputOption.GoogleSpreadsheetsConnectionID),
	}

	columns, err := newGoogleSpreadsheetsInputOptionColumns(ctx, inputOption.InputOptionColumns)
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

func newGoogleSpreadsheetsInputOptionColumns(
	ctx context.Context,
	inputOptionColumns []input_option.GoogleSpreadsheetsInputOptionColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: GoogleSpreadsheetsInputOptionColumn{}.attrTypes(),
	}

	if inputOptionColumns == nil {
		return types.ListNull(objectType), nil
	}

	columns := make([]GoogleSpreadsheetsInputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := GoogleSpreadsheetsInputOptionColumn{
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

func (inputOption *GoogleSpreadsheetsInputOption) ToInput() *param.GoogleSpreadsheetsInputOptionInput {
	if inputOption == nil {
		return nil
	}

	ctx := context.Background()

	var columnValues []GoogleSpreadsheetsInputOptionColumn
	if !inputOption.InputOptionColumns.IsNull() && !inputOption.InputOptionColumns.IsUnknown() {
		diags := inputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	return &param.GoogleSpreadsheetsInputOptionInput{
		SpreadsheetsURL:                inputOption.SpreadsheetsURL.ValueString(),
		WorksheetTitle:                 inputOption.WorksheetTitle.ValueString(),
		StartRow:                       inputOption.StartRow.ValueInt64(),
		StartColumn:                    inputOption.StartColumn.ValueString(),
		DefaultTimeZone:                inputOption.DefaultTimeZone.ValueString(),
		NullString:                     inputOption.NullString.ValueString(),
		GoogleSpreadsheetsConnectionID: inputOption.GoogleSpreadsheetsConnectionID.ValueInt64(),
		InputOptionColumns:             toGoogleSpreadsheetsInputOptionColumnsInput(columnValues),
		CustomVariableSettings:         model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (inputOption *GoogleSpreadsheetsInputOption) ToUpdateInput() *param.UpdateGoogleSpreadsheetsInputOptionInput {
	if inputOption == nil {
		return nil
	}

	ctx := context.Background()

	var columnValues []GoogleSpreadsheetsInputOptionColumn
	if !inputOption.InputOptionColumns.IsNull() {
		if !inputOption.InputOptionColumns.IsUnknown() {
			diags := inputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnValues = []GoogleSpreadsheetsInputOptionColumn{}
		}
	} else {
		columnValues = []GoogleSpreadsheetsInputOptionColumn{}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	return &param.UpdateGoogleSpreadsheetsInputOptionInput{
		SpreadsheetsURL:                inputOption.SpreadsheetsURL.ValueStringPointer(),
		WorksheetTitle:                 inputOption.WorksheetTitle.ValueStringPointer(),
		StartRow:                       inputOption.StartRow.ValueInt64Pointer(),
		StartColumn:                    inputOption.StartColumn.ValueStringPointer(),
		DefaultTimeZone:                inputOption.DefaultTimeZone.ValueStringPointer(),
		NullString:                     inputOption.NullString.ValueStringPointer(),
		GoogleSpreadsheetsConnectionID: inputOption.GoogleSpreadsheetsConnectionID.ValueInt64Pointer(),
		InputOptionColumns:             toGoogleSpreadsheetsInputOptionColumnsInput(columnValues),
		CustomVariableSettings:         model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func toGoogleSpreadsheetsInputOptionColumnsInput(columns []GoogleSpreadsheetsInputOptionColumn) []param.GoogleSpreadsheetsInputOptionColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]param.GoogleSpreadsheetsInputOptionColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, param.GoogleSpreadsheetsInputOptionColumn{
			Name:   column.Name.ValueString(),
			Type:   column.Type.ValueString(),
			Format: column.Format.ValueStringPointer(),
		})
	}
	return inputs
}
