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

type GoogleAdsInputOption struct {
	CustomerID             types.String `tfsdk:"customer_id"`
	ResourceType           types.String `tfsdk:"resource_type"`
	StartDate              types.String `tfsdk:"start_date"`
	EndDate                types.String `tfsdk:"end_date"`
	GoogleAdsConnectionID  types.Int64  `tfsdk:"google_ads_connection_id"`
	InputOptionColumns     types.List   `tfsdk:"input_option_columns"`
	Conditions             types.List   `tfsdk:"conditions"`
	CustomVariableSettings types.List   `tfsdk:"custom_variable_settings"`
}

type GoogleAdsColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func (GoogleAdsColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   types.StringType,
		"type":   types.StringType,
		"format": types.StringType,
	}
}

func NewGoogleAdsInputOption(ctx context.Context, inputOption *inputOptionEntities.GoogleAdsInputOption) *GoogleAdsInputOption {
	if inputOption == nil {
		return nil
	}

	result := &GoogleAdsInputOption{
		CustomerID:            types.StringValue(inputOption.CustomerID),
		ResourceType:          types.StringValue(inputOption.ResourceType),
		StartDate:             types.StringPointerValue(inputOption.StartDate),
		EndDate:               types.StringPointerValue(inputOption.EndDate),
		GoogleAdsConnectionID: types.Int64Value(inputOption.GoogleAdsConnectionID),
	}

	columns, err := newGoogleAdsColumns(ctx, inputOption.InputOptionColumns)
	if err != nil {
		return nil
	}
	result.InputOptionColumns = columns

	conditions, err := newGoogleAdsConditions(inputOption.Conditions)
	if err != nil {
		return nil
	}
	result.Conditions = conditions

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, inputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func newGoogleAdsColumns(
	ctx context.Context,
	googleAdsColumns []inputOptionEntities.GoogleAdsColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: GoogleAdsColumn{}.attrTypes(),
	}

	if googleAdsColumns == nil {
		return types.ListNull(objectType), nil
	}

	columns := make([]GoogleAdsColumn, 0, len(googleAdsColumns))
	for _, input := range googleAdsColumns {
		column := GoogleAdsColumn{
			Name:   types.StringValue(input.Name),
			Type:   types.StringValue(input.Type),
			Format: types.StringPointerValue(input.Format),
		}
		columns = append(columns, column)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columns)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert google ads columns to ListValue: %v", diags)
	}
	return listValue, nil
}

func newGoogleAdsConditions(
	conditions []string,
) (types.List, error) {
	if conditions == nil {
		return types.ListNull(types.StringType), nil
	}

	conditionValues := make([]attr.Value, 0, len(conditions))
	for _, condition := range conditions {
		conditionValues = append(conditionValues, types.StringValue(condition))
	}

	listValue, diags := types.ListValue(types.StringType, conditionValues)
	if diags.HasError() {
		return types.ListNull(types.StringType), fmt.Errorf("failed to convert google ads conditions to ListValue: %v", diags)
	}
	return listValue, nil
}

func (inputOption *GoogleAdsInputOption) ToInput(ctx context.Context) *inputOptionParameters.GoogleAdsInputOptionInput {
	if inputOption == nil {
		return nil
	}

	var columnValues []GoogleAdsColumn
	if !inputOption.InputOptionColumns.IsNull() && !inputOption.InputOptionColumns.IsUnknown() {
		diags := inputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
		if diags.HasError() {
			return nil
		}
	}

	var conditionValues []string
	if !inputOption.Conditions.IsNull() && !inputOption.Conditions.IsUnknown() {
		diags := inputOption.Conditions.ElementsAs(ctx, &conditionValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	return &inputOptionParameters.GoogleAdsInputOptionInput{
		CustomerID:             inputOption.CustomerID.ValueString(),
		ResourceType:           inputOption.ResourceType.ValueString(),
		StartDate:              model.NewNullableString(inputOption.StartDate),
		EndDate:                model.NewNullableString(inputOption.EndDate),
		GoogleAdsConnectionID:  inputOption.GoogleAdsConnectionID.ValueInt64(),
		InputOptionColumns:     toGoogleAdsColumnsInput(columnValues),
		Conditions:             conditionValues,
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (inputOption *GoogleAdsInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateGoogleAdsInputOptionInput {
	if inputOption == nil {
		return nil
	}

	var columnValues []GoogleAdsColumn
	if !inputOption.InputOptionColumns.IsNull() {
		if !inputOption.InputOptionColumns.IsUnknown() {
			diags := inputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnValues = []GoogleAdsColumn{}
		}
	} else {
		columnValues = []GoogleAdsColumn{}
	}

	var conditionValues []string
	if !inputOption.Conditions.IsNull() {
		if !inputOption.Conditions.IsUnknown() {
			diags := inputOption.Conditions.ElementsAs(ctx, &conditionValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			conditionValues = []string{}
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	result := &inputOptionParameters.UpdateGoogleAdsInputOptionInput{
		CustomerID:             model.NewNullableString(inputOption.CustomerID),
		ResourceType:           model.NewNullableString(inputOption.ResourceType),
		StartDate:              model.NewNullableString(inputOption.StartDate),
		EndDate:                model.NewNullableString(inputOption.EndDate),
		GoogleAdsConnectionID:  model.NewNullableInt64(inputOption.GoogleAdsConnectionID),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
	}

	if columnValues != nil {
		columns := toGoogleAdsColumnsInput(columnValues)
		result.InputOptionColumns = &columns
	}

	if conditionValues != nil {
		result.Conditions = &conditionValues
	}

	return result
}

func toGoogleAdsColumnsInput(columns []GoogleAdsColumn) []inputOptionParameters.GoogleAdsColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.GoogleAdsColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, inputOptionParameters.GoogleAdsColumn{
			Name:   column.Name.ValueString(),
			Type:   column.Type.ValueString(),
			Format: column.Format.ValueStringPointer(),
		})
	}
	return inputs
}
