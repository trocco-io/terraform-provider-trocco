package output_options

import (
	"context"
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

	if googleSpreadsheetsOutputOption.GoogleSpreadsheetsOutputOptionSorts != nil {
		sorts := make([]googleSpreadsheetsOutputOptionSorts, 0, len(*googleSpreadsheetsOutputOption.GoogleSpreadsheetsOutputOptionSorts))
		for _, input := range *googleSpreadsheetsOutputOption.GoogleSpreadsheetsOutputOptionSorts {
			newSort := googleSpreadsheetsOutputOptionSorts{
				Column: types.StringValue(input.Column),
				Order:  types.StringValue(input.Order),
			}
			sorts = append(sorts, newSort)
		}

		objectType := types.ObjectType{
			AttrTypes: googleSpreadsheetsOutputOptionSorts{}.attrTypes(),
		}

		listValue, _ := types.ListValueFrom(ctx, objectType, sorts)
		result.GoogleSpreadsheetsOutputOptionSorts = listValue
	} else {
		result.GoogleSpreadsheetsOutputOptionSorts = types.ListNull(types.ObjectType{
			AttrTypes: googleSpreadsheetsOutputOptionSorts{}.attrTypes(),
		})
	}

	result.CustomVariableSettings = ConvertCustomVariableSettingsToList(ctx, googleSpreadsheetsOutputOption.CustomVariableSettings)

	return result
}

func (outputOption *GoogleSpreadsheetsOutputOption) ToInput() *parameter.GoogleSpreadsheetsOutputOptionInput {
	if outputOption == nil {
		return nil
	}

	ctx := context.Background()

	var sorts *[]parameter.GoogleSpreadsheetsOutputOptionSortsInput
	if !outputOption.GoogleSpreadsheetsOutputOptionSorts.IsNull() && !outputOption.GoogleSpreadsheetsOutputOptionSorts.IsUnknown() {
		var sortValues []googleSpreadsheetsOutputOptionSorts
		outputOption.GoogleSpreadsheetsOutputOptionSorts.ElementsAs(ctx, &sortValues, false)

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
	if !outputOption.GoogleSpreadsheetsOutputOptionSorts.IsNull() && !outputOption.GoogleSpreadsheetsOutputOptionSorts.IsUnknown() {
		var sortValues []googleSpreadsheetsOutputOptionSorts
		outputOption.GoogleSpreadsheetsOutputOptionSorts.ElementsAs(ctx, &sortValues, false)

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
