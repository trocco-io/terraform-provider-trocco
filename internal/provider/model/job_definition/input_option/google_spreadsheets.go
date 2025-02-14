package input_options

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	param "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GoogleSpreadsheetsInputOption struct {
	GoogleSpreadsheetsConnectionID types.Int64                           `tfsdk:"google_spreadsheets_connection_id"`
	SpreadsheetsID                 types.String                          `tfsdk:"spreadsheets_id"`
	WorksheetTitle                 types.String                          `tfsdk:"worksheet_title"`
	StartRow                       types.Int64                           `tfsdk:"start_row"`
	StartColumn                    types.String                          `tfsdk:"start_column"`
	DefaultTimeZone                types.String                          `tfsdk:"default_time_zone"`
	NullString                     types.String                          `tfsdk:"null_string"`
	InputOptionColumns             []GoogleSpreadsheetsInputOptionColumn `tfsdk:"input_option_columns"`
	CustomVariableSettings         *[]model.CustomVariableSetting        `tfsdk:"custom_variable_settings"`
}

func (inputOption *GoogleSpreadsheetsInputOption) SpreadsheetsURL() types.String {
	if inputOption.SpreadsheetsID.IsNull() || inputOption.SpreadsheetsID.ValueString() == "" {
		return types.StringNull()
	}
	// ex) "https://docs.google.com/spreadsheets/d/MY_SHEETS_ID/edit#gid=0"
	url := "https://docs.google.com/spreadsheets/d/" + inputOption.SpreadsheetsID.ValueString() + "/edit#gid=0"
	return types.StringValue(url)
}

type GoogleSpreadsheetsInputOptionColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func NewGoogleSpreadsheetsInputOption(inputOption *input_option.GoogleSpreadsheetsInputOption) *GoogleSpreadsheetsInputOption {
	if inputOption == nil {
		return nil
	}

	return &GoogleSpreadsheetsInputOption{
		SpreadsheetsID:                 types.StringValue(inputOption.SpreadsheetsID()),
		WorksheetTitle:                 types.StringValue(inputOption.WorksheetTitle),
		StartRow:                       types.Int64Value(inputOption.StartRow),
		StartColumn:                    types.StringValue(inputOption.StartColumn),
		DefaultTimeZone:                types.StringValue(inputOption.DefaultTimeZone),
		NullString:                     types.StringValue(inputOption.NullString),
		GoogleSpreadsheetsConnectionID: types.Int64Value(inputOption.GoogleSpreadsheetsConnectionID),
		InputOptionColumns:             newGoogleSpreadsheetsInputOptionColumns(inputOption.InputOptionColumns),
		CustomVariableSettings:         model.NewCustomVariableSettings(inputOption.CustomVariableSettings),
	}
}

func newGoogleSpreadsheetsInputOptionColumns(inputOptionColumns []input_option.GoogleSpreadsheetsInputOptionColumn) []GoogleSpreadsheetsInputOptionColumn {
	if inputOptionColumns == nil {
		return nil
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
	return columns
}

func (inputOption *GoogleSpreadsheetsInputOption) ToInput() *param.GoogleSpreadsheetsInputOptionInput {
	if inputOption == nil {
		return nil
	}

	return &param.GoogleSpreadsheetsInputOptionInput{
		SpreadsheetsURL:                inputOption.SpreadsheetsURL().ValueString(),
		WorksheetTitle:                 inputOption.WorksheetTitle.ValueString(),
		StartRow:                       inputOption.StartRow.ValueInt64(),
		StartColumn:                    inputOption.StartColumn.ValueString(),
		DefaultTimeZone:                inputOption.DefaultTimeZone.ValueString(),
		NullString:                     inputOption.NullString.ValueString(),
		GoogleSpreadsheetsConnectionID: inputOption.GoogleSpreadsheetsConnectionID.ValueInt64(),
		InputOptionColumns:             toGoogleSpreadsheetsInputOptionColumnsInput(inputOption.InputOptionColumns),
		CustomVariableSettings:         model.ToCustomVariableSettingInputs(inputOption.CustomVariableSettings),
	}
}

func (inputOption *GoogleSpreadsheetsInputOption) ToUpdateInput() *param.UpdateGoogleSpreadsheetsInputOptionInput {
	if inputOption == nil {
		return nil
	}

	inputOptionColumns := toGoogleSpreadsheetsInputOptionColumnsInput(inputOption.InputOptionColumns)

	return &param.UpdateGoogleSpreadsheetsInputOptionInput{
		SpreadsheetsURL:                model.NewNullableString(inputOption.SpreadsheetsURL()),
		WorksheetTitle:                 model.NewNullableString(inputOption.WorksheetTitle),
		StartRow:                       model.NewNullableInt64(inputOption.StartRow),
		StartColumn:                    model.NewNullableString(inputOption.StartColumn),
		DefaultTimeZone:                model.NewNullableString(inputOption.DefaultTimeZone),
		NullString:                     model.NewNullableString(inputOption.NullString),
		GoogleSpreadsheetsConnectionID: model.NewNullableInt64(inputOption.GoogleSpreadsheetsConnectionID),
		InputOptionColumns:             inputOptionColumns,
		CustomVariableSettings:         model.ToCustomVariableSettingInputs(inputOption.CustomVariableSettings),
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
