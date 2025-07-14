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

type GoogleAnalytics4InputOption struct {
	GoogleAnalytics4ConnectionID types.Int64  `tfsdk:"google_analytics4_connection_id"`
	PropertyID                   types.String `tfsdk:"property_id"`
	TimeSeries                   types.String `tfsdk:"time_series"`
	StartDate                    types.String `tfsdk:"start_date"`
	EndDate                      types.String `tfsdk:"end_date"`
	IncrementalLoadingEnabled    types.Bool   `tfsdk:"incremental_loading_enabled"`
	RetryLimit                   types.Int64  `tfsdk:"retry_limit"`
	RetrySleep                   types.Int64  `tfsdk:"retry_sleep"`
	RaiseOnOtherRow              types.Bool   `tfsdk:"raise_on_other_row"`
	LimitOfRows                  types.Int64  `tfsdk:"limit_of_rows"`
	GoogleAnalytics4Dimensions   types.List   `tfsdk:"google_analytics4_input_option_dimensions"`
	GoogleAnalytics4Metrics      types.List   `tfsdk:"google_analytics4_input_option_metrics"`
	InputOptionColumns           types.List   `tfsdk:"input_option_columns"`
	CustomVariableSettings       types.List   `tfsdk:"custom_variable_settings"`
}

type GoogleAnalytics4Dimension struct {
	Name       types.String `tfsdk:"name"`
	Expression types.String `tfsdk:"expression"`
}

func (GoogleAnalytics4Dimension) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":       types.StringType,
		"expression": types.StringType,
	}
}

type GoogleAnalytics4Metric struct {
	Name       types.String `tfsdk:"name"`
	Expression types.String `tfsdk:"expression"`
}

func (GoogleAnalytics4Metric) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":       types.StringType,
		"expression": types.StringType,
	}
}

type GoogleAnalytics4Column struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

func (GoogleAnalytics4Column) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
		"type": types.StringType,
	}
}

func NewGoogleAnalytics4InputOption(ctx context.Context, inputOption *inputOptionEntities.GoogleAnalytics4InputOption) *GoogleAnalytics4InputOption {
	if inputOption == nil {
		return nil
	}

	result := &GoogleAnalytics4InputOption{
		GoogleAnalytics4ConnectionID: types.Int64Value(inputOption.GoogleAnalytics4ConnectionID),
		PropertyID:                   types.StringValue(inputOption.PropertyID),
		TimeSeries:                   types.StringValue(inputOption.TimeSeries),
		StartDate:                    types.StringPointerValue(inputOption.StartDate),
		EndDate:                      types.StringPointerValue(inputOption.EndDate),
		IncrementalLoadingEnabled:    types.BoolValue(inputOption.IncrementalLoadingEnabled),
		RetryLimit:                   types.Int64PointerValue(inputOption.RetryLimit),
		RetrySleep:                   types.Int64PointerValue(inputOption.RetrySleep),
		RaiseOnOtherRow:              types.BoolPointerValue(inputOption.RaiseOnOtherRow),
		LimitOfRows:                  types.Int64PointerValue(inputOption.LimitOfRows),
	}

	dimensions, err := newGoogleAnalytics4Dimensions(ctx, inputOption.GoogleAnalytics4Dimensions)
	if err != nil {
		return nil
	}
	result.GoogleAnalytics4Dimensions = dimensions

	metrics, err := newGoogleAnalytics4Metrics(ctx, inputOption.GoogleAnalytics4Metrics)
	if err != nil {
		return nil
	}
	result.GoogleAnalytics4Metrics = metrics

	columns, err := newGoogleAnalytics4InputOptionColumns(ctx, inputOption.InputOptionColumns)
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

func newGoogleAnalytics4Dimensions(
	ctx context.Context,
	inputOptionDimensions []inputOptionEntities.GoogleAnalytics4Dimension,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: GoogleAnalytics4Dimension{}.attrTypes(),
	}

	if inputOptionDimensions == nil {
		return types.ListNull(objectType), nil
	}

	dimensions := make([]GoogleAnalytics4Dimension, 0, len(inputOptionDimensions))
	for _, input := range inputOptionDimensions {
		dimension := GoogleAnalytics4Dimension{
			Name:       types.StringValue(input.Name),
			Expression: types.StringValue(input.Expression),
		}
		dimensions = append(dimensions, dimension)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, dimensions)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert google analytics 4 dimensions to ListValue: %v", diags)
	}
	return listValue, nil
}

func newGoogleAnalytics4Metrics(
	ctx context.Context,
	inputOptionMetrics []inputOptionEntities.GoogleAnalytics4Metric,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: GoogleAnalytics4Metric{}.attrTypes(),
	}

	if inputOptionMetrics == nil {
		return types.ListNull(objectType), nil
	}

	metrics := make([]GoogleAnalytics4Metric, 0, len(inputOptionMetrics))
	for _, input := range inputOptionMetrics {
		metric := GoogleAnalytics4Metric{
			Name:       types.StringValue(input.Name),
			Expression: types.StringValue(input.Expression),
		}
		metrics = append(metrics, metric)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, metrics)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert google analytics 4 metrics to ListValue: %v", diags)
	}
	return listValue, nil
}

func newGoogleAnalytics4InputOptionColumns(
	ctx context.Context,
	inputOptionColumns []inputOptionEntities.GoogleAnalytics4Column,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: GoogleAnalytics4Column{}.attrTypes(),
	}

	if inputOptionColumns == nil {
		return types.ListNull(objectType), nil
	}

	columns := make([]GoogleAnalytics4Column, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := GoogleAnalytics4Column{
			Name: types.StringValue(input.Name),
			Type: types.StringValue(input.Type),
		}
		columns = append(columns, column)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columns)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert google analytics 4 columns to ListValue: %v", diags)
	}
	return listValue, nil
}

func (inputOption *GoogleAnalytics4InputOption) ToInput(ctx context.Context) *inputOptionParameters.GoogleAnalytics4InputOptionInput {
	if inputOption == nil {
		return nil
	}

	var dimensionValues []GoogleAnalytics4Dimension
	if !inputOption.GoogleAnalytics4Dimensions.IsNull() && !inputOption.GoogleAnalytics4Dimensions.IsUnknown() {
		diags := inputOption.GoogleAnalytics4Dimensions.ElementsAs(ctx, &dimensionValues, false)
		if diags.HasError() {
			return nil
		}
	}

	var metricValues []GoogleAnalytics4Metric
	if !inputOption.GoogleAnalytics4Metrics.IsNull() && !inputOption.GoogleAnalytics4Metrics.IsUnknown() {
		diags := inputOption.GoogleAnalytics4Metrics.ElementsAs(ctx, &metricValues, false)
		if diags.HasError() {
			return nil
		}
	}

	var columnValues []GoogleAnalytics4Column
	if !inputOption.InputOptionColumns.IsNull() && !inputOption.InputOptionColumns.IsUnknown() {
		diags := inputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	return &inputOptionParameters.GoogleAnalytics4InputOptionInput{
		GoogleAnalytics4ConnectionID:          inputOption.GoogleAnalytics4ConnectionID.ValueInt64(),
		PropertyID:                            inputOption.PropertyID.ValueString(),
		TimeSeries:                            inputOption.TimeSeries.ValueString(),
		StartDate:                             model.NewNullableString(inputOption.StartDate),
		EndDate:                               model.NewNullableString(inputOption.EndDate),
		IncrementalLoadingEnabled:             inputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		RetryLimit:                            model.NewNullableInt64(inputOption.RetryLimit),
		RetrySleep:                            model.NewNullableInt64(inputOption.RetrySleep),
		RaiseOnOtherRow:                       inputOption.RaiseOnOtherRow.ValueBoolPointer(),
		LimitOfRows:                           model.NewNullableInt64(inputOption.LimitOfRows),
		GoogleAnalytics4InputOptionDimensions: toGoogleAnalytics4DimensionsInput(dimensionValues),
		GoogleAnalytics4InputOptionMetrics:    toGoogleAnalytics4MetricsInput(metricValues),
		InputOptionColumns:                    toGoogleAnalytics4ColumnsInput(columnValues),
		CustomVariableSettings:                model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (inputOption *GoogleAnalytics4InputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateGoogleAnalytics4InputOptionInput {
	if inputOption == nil {
		return nil
	}

	var dimensionValues []GoogleAnalytics4Dimension
	if !inputOption.GoogleAnalytics4Dimensions.IsNull() {
		if !inputOption.GoogleAnalytics4Dimensions.IsUnknown() {
			diags := inputOption.GoogleAnalytics4Dimensions.ElementsAs(ctx, &dimensionValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			dimensionValues = []GoogleAnalytics4Dimension{}
		}
	} else {
		dimensionValues = []GoogleAnalytics4Dimension{}
	}

	var metricValues []GoogleAnalytics4Metric
	if !inputOption.GoogleAnalytics4Metrics.IsNull() {
		if !inputOption.GoogleAnalytics4Metrics.IsUnknown() {
			diags := inputOption.GoogleAnalytics4Metrics.ElementsAs(ctx, &metricValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			metricValues = []GoogleAnalytics4Metric{}
		}
	} else {
		metricValues = []GoogleAnalytics4Metric{}
	}

	var columnValues []GoogleAnalytics4Column
	if !inputOption.InputOptionColumns.IsNull() {
		if !inputOption.InputOptionColumns.IsUnknown() {
			diags := inputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnValues = []GoogleAnalytics4Column{}
		}
	} else {
		columnValues = []GoogleAnalytics4Column{}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	dimensions := toGoogleAnalytics4DimensionsInput(dimensionValues)
	metrics := toGoogleAnalytics4MetricsInput(metricValues)
	columns := toGoogleAnalytics4ColumnsInput(columnValues)

	return &inputOptionParameters.UpdateGoogleAnalytics4InputOptionInput{
		GoogleAnalytics4ConnectionID:          model.NewNullableInt64(inputOption.GoogleAnalytics4ConnectionID),
		PropertyID:                            model.NewNullableString(inputOption.PropertyID),
		TimeSeries:                            model.NewNullableString(inputOption.TimeSeries),
		StartDate:                             model.NewNullableString(inputOption.StartDate),
		EndDate:                               model.NewNullableString(inputOption.EndDate),
		IncrementalLoadingEnabled:             inputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		RetryLimit:                            model.NewNullableInt64(inputOption.RetryLimit),
		RetrySleep:                            model.NewNullableInt64(inputOption.RetrySleep),
		RaiseOnOtherRow:                       inputOption.RaiseOnOtherRow.ValueBoolPointer(),
		LimitOfRows:                           model.NewNullableInt64(inputOption.LimitOfRows),
		GoogleAnalytics4InputOptionDimensions: model.WrapObjectList(&dimensions),
		GoogleAnalytics4InputOptionMetrics:    &metrics,
		InputOptionColumns:                    &columns,
		CustomVariableSettings:                model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func toGoogleAnalytics4ColumnsInput(columns []GoogleAnalytics4Column) []inputOptionParameters.GoogleAnalytics4Column {
	if columns == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.GoogleAnalytics4Column, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, inputOptionParameters.GoogleAnalytics4Column{
			Name: column.Name.ValueString(),
			Type: column.Type.ValueString(),
		})
	}
	return inputs
}

func toGoogleAnalytics4DimensionsInput(dimensions []GoogleAnalytics4Dimension) []inputOptionParameters.GoogleAnalytics4Dimension {
	if dimensions == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.GoogleAnalytics4Dimension, 0, len(dimensions))
	for _, dimension := range dimensions {
		inputs = append(inputs, inputOptionParameters.GoogleAnalytics4Dimension{
			Name:       dimension.Name.ValueString(),
			Expression: model.NewNullableString(dimension.Expression),
		})
	}
	return inputs
}

func toGoogleAnalytics4MetricsInput(metrics []GoogleAnalytics4Metric) []inputOptionParameters.GoogleAnalytics4Metric {
	if metrics == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.GoogleAnalytics4Metric, 0, len(metrics))
	for _, metric := range metrics {
		inputs = append(inputs, inputOptionParameters.GoogleAnalytics4Metric{
			Name:       metric.Name.ValueString(),
			Expression: model.NewNullableString(metric.Expression),
		})
	}
	return inputs
}
