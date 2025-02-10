package output_options

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	parameter "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GoogleSpreadsheetsOutputOption struct {
	GoogleSpreadsheetsConnectionId      types.Int64                            `tfsdk:"google_spreadsheets_connection_id"`
	SpreadsheetsID                      types.String                           `tfsdk:"spreadsheets_id"`
	WorksheetTitle                      types.String                           `tfsdk:"worksheet_title"`
	Timezone                            types.String                           `tfsdk:"timezone"`
	ValueInputOption                    types.String                           `tfsdk:"value_input_option"`
	Mode                                types.String                           `tfsdk:"mode"`
	GoogleSpreadsheetsOutputOptionSorts *[]googleSpreadsheetsOutputOptionSorts `tfsdk:"google_spreadsheets_output_option_sorts"`
	CustomVariableSettings              *[]model.CustomVariableSetting         `tfsdk:"custom_variable_settings"`
}

type googleSpreadsheetsOutputOptionSorts struct {
	Column types.String `tfsdk:"column"`
	Order  types.String `tfsdk:"order"`
}

func NewGoogleSpreadsheetsOutputOption(googleSpreadsheetsOutputOption *output_option.GoogleSpreadsheetsOutputOption) *GoogleSpreadsheetsOutputOption {
	if googleSpreadsheetsOutputOption == nil {
		return nil
	}

	return &GoogleSpreadsheetsOutputOption{
		GoogleSpreadsheetsConnectionId:      types.Int64Value(googleSpreadsheetsOutputOption.GoogleSpreadsheetsConnectionId),
		GoogleSpreadsheetsOutputOptionSorts: newGoogleSpreadsheetsOutputOptionSorts(googleSpreadsheetsOutputOption.GoogleSpreadsheetsOutputOptionSorts),
		CustomVariableSettings:              model.NewCustomVariableSettings(googleSpreadsheetsOutputOption.CustomVariableSettings),
	}
}

func newGoogleSpreadsheetsOutputOptionSorts(sorts *[]output_option.GoogleSpreadsheetsOutputOptionSorts) *[]googleSpreadsheetsOutputOptionSorts {
	if sorts == nil {
		return nil
	}

	outputs := make([]googleSpreadsheetsOutputOptionSorts, 0, len(*sorts))
	for _, input := range *sorts {
		newSort := googleSpreadsheetsOutputOptionSorts{
			Column: types.StringValue(input.Column),
			Order:  types.StringValue(input.Order),
		}
		outputs = append(outputs, newSort)
	}
	return &outputs
}

func (outputOption *GoogleSpreadsheetsOutputOption) ToInput() *parameter.GoogleSpreadsheetsOutputOptionInput {
	if outputOption == nil {
		return nil
	}

	return &parameter.GoogleSpreadsheetsOutputOptionInput{
		GoogleSpreadsheetsConnectionId:      outputOption.GoogleSpreadsheetsConnectionId.ValueInt64(),
		SpreadsheetsID:                      outputOption.SpreadsheetsID.ValueString(),
		WorksheetTitle:                      outputOption.WorksheetTitle.ValueString(),
		Timezone:                            outputOption.Timezone.ValueString(),
		ValueInputOption:                      outputOption.ValueInputOption.ValueString(),
		Mode:                                outputOption.Mode.ValueString(),
		GoogleSpreadsheetsOutputOptionSorts: model.WrapObjectList(toGoogleSpreadsheetsOutputOptionSorts(outputOption)),
		CustomVariableSettings:              model.ToCustomVariableSettingInputs(outputOption.CustomVariableSettings),
	}
}

func (outputOption *GoogleSpreadsheetsOutputOption) ToUpdateInput() *parameter.UpdateGoogleSpreadsheetsOutputOptionInput {
	if outputOption == nil {
		return nil
	}

	return &parameter.UpdateGoogleSpreadsheetsOutputOptionInput{
		GoogleSpreadsheetsConnectionId:      outputOption.GoogleSpreadsheetsConnectionId.ValueInt64Pointer(),
		SpreadsheetsID:                      outputOption.SpreadsheetsID.ValueStringPointer(),
		WorksheetTitle:                      outputOption.WorksheetTitle.ValueStringPointer(),
		Timezone:                            outputOption.Timezone.ValueStringPointer(),
		ValueInputOption:                      outputOption.ValueInputOption.ValueStringPointer(),
		Mode:                                outputOption.Mode.ValueStringPointer(),
		GoogleSpreadsheetsOutputOptionSorts: model.WrapObjectList(toGoogleSpreadsheetsOutputOptionSorts(outputOption)),
		CustomVariableSettings:              model.ToCustomVariableSettingInputs(outputOption.CustomVariableSettings),
	}
}

func toGoogleSpreadsheetsOutputOptionSorts(outputOption *GoogleSpreadsheetsOutputOption) *[]parameter.GoogleSpreadsheetsOutputOptionSortsInput {
	var sorts *[]parameter.GoogleSpreadsheetsOutputOptionSortsInput
	if outputOption.GoogleSpreadsheetsOutputOptionSorts != nil {
		s := make([]parameter.GoogleSpreadsheetsOutputOptionSortsInput, 0, len(*outputOption.GoogleSpreadsheetsOutputOptionSorts))
		for _, input := range *outputOption.GoogleSpreadsheetsOutputOptionSorts {
			s = append(s, parameter.GoogleSpreadsheetsOutputOptionSortsInput{
				Column: input.Column.ValueString(),
				Order:  input.Order.ValueString(),
			})
		}
		sorts = &s
	}
	return sorts
}
