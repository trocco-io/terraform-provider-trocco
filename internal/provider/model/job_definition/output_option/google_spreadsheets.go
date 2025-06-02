package output_options

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	parameter "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GoogleSpreadsheetsOutputOption struct {
	GoogleSpreadsheetsConnectionId      types.Int64  `tfsdk:"google_spreadsheets_connection_id"`
	SpreadsheetsID                      types.String `tfsdk:"spreadsheets_id"`
	WorksheetTitle                      types.String `tfsdk:"worksheet_title"`
	Timezone                            types.String `tfsdk:"timezone"`
	ValueInputOption                    types.String `tfsdk:"value_input_option"`
	Mode                                types.String `tfsdk:"mode"`
	GoogleSpreadsheetsOutputOptionSorts types.List   `tfsdk:"google_spreadsheets_output_option_sorts"`
	CustomVariableSettings              types.List   `tfsdk:"custom_variable_settings"`
}

type googleSpreadsheetsOutputOptionSorts struct {
	Column types.String `tfsdk:"column"`
	Order  types.String `tfsdk:"order"`
}

func (s googleSpreadsheetsOutputOptionSorts) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"column": types.StringType,
		"order":  types.StringType,
	}
}

func NewGoogleSpreadsheetsOutputOption(googleSpreadsheetsOutputOption *output_option.GoogleSpreadsheetsOutputOption) *GoogleSpreadsheetsOutputOption {
	if googleSpreadsheetsOutputOption == nil {
		return nil
	}

	ctx := context.Background()
	result := &GoogleSpreadsheetsOutputOption{
		GoogleSpreadsheetsConnectionId: types.Int64Value(googleSpreadsheetsOutputOption.GoogleSpreadsheetsConnectionId),
		SpreadsheetsID:                 types.StringValue(googleSpreadsheetsOutputOption.SpreadsheetsID),
		WorksheetTitle:                 types.StringValue(googleSpreadsheetsOutputOption.WorksheetTitle),
		Timezone:                       types.StringValue(googleSpreadsheetsOutputOption.Timezone),
		ValueInputOption:               types.StringValue(googleSpreadsheetsOutputOption.ValueInputOption),
		Mode:                           types.StringValue(googleSpreadsheetsOutputOption.Mode),
	}

	GoogleSpreadsheetsOutputOptionSorts, err := newGoogleSpreadsheetsOutputOptionSorts(ctx, googleSpreadsheetsOutputOption.GoogleSpreadsheetsOutputOptionSorts)
	if err != nil {
		return nil
	}
	result.GoogleSpreadsheetsOutputOptionSorts = GoogleSpreadsheetsOutputOptionSorts

	CustomVariableSettings, err := ConvertCustomVariableSettingsToList(ctx, googleSpreadsheetsOutputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = CustomVariableSettings
	return result
}

func newGoogleSpreadsheetsOutputOptionSorts(ctx context.Context, sorts *[]output_option.GoogleSpreadsheetsOutputOptionSorts) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: googleSpreadsheetsOutputOptionSorts{}.attrTypes(),
	}

	if sorts == nil {
		return types.ListNull(objectType), nil
	}

	converted := make([]googleSpreadsheetsOutputOptionSorts, 0, len(*sorts))
	for _, input := range *sorts {
		newSort := googleSpreadsheetsOutputOptionSorts{
			Column: types.StringValue(input.Column),
			Order:  types.StringValue(input.Order),
		}
		converted = append(converted, newSort)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, converted)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert to ListValue: %v", diags)
	}

	return listValue, nil
}

func (outputOption *GoogleSpreadsheetsOutputOption) ToInput() *parameter.GoogleSpreadsheetsOutputOptionInput {
	if outputOption == nil {
		return nil
	}

	ctx := context.Background()

	var sorts *[]parameter.GoogleSpreadsheetsOutputOptionSortsInput
	if !outputOption.GoogleSpreadsheetsOutputOptionSorts.IsNull() {
		var sortValues []googleSpreadsheetsOutputOptionSorts
		diags := outputOption.GoogleSpreadsheetsOutputOptionSorts.ElementsAs(ctx, &sortValues, false)
		if diags.HasError() {
			return nil
		}

		s := make([]parameter.GoogleSpreadsheetsOutputOptionSortsInput, 0, len(sortValues))
		for _, input := range sortValues {
			s = append(s, parameter.GoogleSpreadsheetsOutputOptionSortsInput{
				Column: input.Column.ValueString(),
				Order:  input.Order.ValueString(),
			})
		}
		sorts = &s
	}

	customVarSettings := ExtractCustomVariableSettings(ctx, outputOption.CustomVariableSettings)

	return &parameter.GoogleSpreadsheetsOutputOptionInput{
		GoogleSpreadsheetsConnectionId:      outputOption.GoogleSpreadsheetsConnectionId.ValueInt64(),
		SpreadsheetsID:                      outputOption.SpreadsheetsID.ValueString(),
		WorksheetTitle:                      outputOption.WorksheetTitle.ValueString(),
		Timezone:                            outputOption.Timezone.ValueString(),
		ValueInputOption:                    outputOption.ValueInputOption.ValueString(),
		Mode:                                outputOption.Mode.ValueString(),
		GoogleSpreadsheetsOutputOptionSorts: model.WrapObjectList(sorts),
		CustomVariableSettings:              model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (outputOption *GoogleSpreadsheetsOutputOption) ToUpdateInput() *parameter.UpdateGoogleSpreadsheetsOutputOptionInput {
	if outputOption == nil {
		return nil
	}

	ctx := context.Background()

	var sorts *[]parameter.GoogleSpreadsheetsOutputOptionSortsInput
	if !outputOption.GoogleSpreadsheetsOutputOptionSorts.IsNull() {
		var sortValues []googleSpreadsheetsOutputOptionSorts
		diags := outputOption.GoogleSpreadsheetsOutputOptionSorts.ElementsAs(ctx, &sortValues, false)
		if diags.HasError() {
			return nil
		}

		s := make([]parameter.GoogleSpreadsheetsOutputOptionSortsInput, 0, len(sortValues))
		for _, input := range sortValues {
			s = append(s, parameter.GoogleSpreadsheetsOutputOptionSortsInput{
				Column: input.Column.ValueString(),
				Order:  input.Order.ValueString(),
			})
		}
		sorts = &s
	}

	customVarSettings := ExtractCustomVariableSettings(ctx, outputOption.CustomVariableSettings)

	return &parameter.UpdateGoogleSpreadsheetsOutputOptionInput{
		GoogleSpreadsheetsConnectionId:      outputOption.GoogleSpreadsheetsConnectionId.ValueInt64Pointer(),
		SpreadsheetsID:                      outputOption.SpreadsheetsID.ValueStringPointer(),
		WorksheetTitle:                      outputOption.WorksheetTitle.ValueStringPointer(),
		Timezone:                            outputOption.Timezone.ValueStringPointer(),
		ValueInputOption:                    outputOption.ValueInputOption.ValueStringPointer(),
		Mode:                                outputOption.Mode.ValueStringPointer(),
		GoogleSpreadsheetsOutputOptionSorts: model.WrapObjectList(sorts),
		CustomVariableSettings:              model.ToCustomVariableSettingInputs(customVarSettings),
	}
}
