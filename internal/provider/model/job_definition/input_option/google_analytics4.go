package input_options

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	param "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GoogleAnalytics4InputOption struct {
	GoogleAnalytics4ConnectionID types.Int64                    `tfsdk:"google_analytics4_connection_id"`
	PropertyID                   types.String                   `tfsdk:"property_id"`
	TimeSeries                   types.String                   `tfsdk:"time_series"`
	StartDate                    types.String                   `tfsdk:"start_date"`
	EndDate                      types.String                   `tfsdk:"end_date"`
	IncrementalLoadingEnabled    types.Bool                     `tfsdk:"incremental_loading_enabled"`
	RetryLimit                   types.Int64                    `tfsdk:"retry_limit"`
	RetrySleep                   types.Int64                    `tfsdk:"retry_sleep"`
	RaiseOnOtherRow              types.Bool                     `tfsdk:"raise_on_other_row"`
	LimitOfRows                  types.Int64                    `tfsdk:"limit_of_rows"`
	GoogleAnalytics4Dimensions   []GoogleAnalytics4Dimension    `tfsdk:"google_analytics4_input_option_dimensions"`
	GoogleAnalytics4Metrics      []GoogleAnalytics4Metric       `tfsdk:"google_analytics4_input_option_metrics"`
	InputOptionColumns           []GoogleAnalytics4Column       `tfsdk:"input_option_columns"`
	CustomVariableSettings       *[]model.CustomVariableSetting `tfsdk:"custom_variable_settings"`
}

type GoogleAnalytics4Dimension struct {
	Name       types.String `tfsdk:"name"`
	Expression types.String `tfsdk:"expression"`
}

type GoogleAnalytics4Metric struct {
	Name       types.String `tfsdk:"name"`
	Expression types.String `tfsdk:"expression"`
}

type GoogleAnalytics4Column struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

func NewGoogleAnalytics4InputOption(inputOption *input_option.GoogleAnalytics4InputOption) *GoogleAnalytics4InputOption {
	if inputOption == nil {
		return nil
	}

	return &GoogleAnalytics4InputOption{
		GoogleAnalytics4ConnectionID: types.Int64Value(inputOption.GoogleAnalytics4ConnectionID),
		PropertyID:                   types.StringValue(inputOption.PropertyID),
		TimeSeries:                   types.StringValue(inputOption.TimeSeries),
		StartDate:                    types.StringPointerValue(inputOption.StartDate),
		EndDate:                      types.StringPointerValue(inputOption.EndDate),
		IncrementalLoadingEnabled:    types.BoolPointerValue(inputOption.IncrementalLoadingEnabled),
		RetryLimit:                   types.Int64PointerValue(inputOption.RetryLimit),
		RetrySleep:                   types.Int64PointerValue(inputOption.RetrySleep),
		RaiseOnOtherRow:              types.BoolPointerValue(inputOption.RaiseOnOtherRow),
		LimitOfRows:                  types.Int64PointerValue(inputOption.LimitOfRows),
		GoogleAnalytics4Dimensions:   newGoogleAnalytics4Dimensions(inputOption.GoogleAnalytics4Dimensions),
		GoogleAnalytics4Metrics:      newGoogleAnalytics4Metrics(inputOption.GoogleAnalytics4Metrics),
		InputOptionColumns:           newGoogleAnalytics4InputOptionColumns(inputOption.InputOptionColumns),
		CustomVariableSettings:       model.NewCustomVariableSettings(inputOption.CustomVariableSettings),
	}
}
func newGoogleAnalytics4Dimensions(inputOptionDimensions []input_option.GoogleAnalytics4Dimension) []GoogleAnalytics4Dimension {
	if inputOptionDimensions == nil {
		return nil
	}
	if len(inputOptionDimensions) == 0 {
		return nil
	}
	dimensions := make([]GoogleAnalytics4Dimension, 0, len(inputOptionDimensions))
	for _, input := range inputOptionDimensions {
		column := GoogleAnalytics4Dimension{
			Name:       types.StringValue(input.Name),
			Expression: types.StringValue(input.Expression),
		}
		dimensions = append(dimensions, column)
	}
	return dimensions
}

func newGoogleAnalytics4Metrics(inputOptionMetrics []input_option.GoogleAnalytics4Metric) []GoogleAnalytics4Metric {
	if inputOptionMetrics == nil {
		return nil
	}
	metrics := make([]GoogleAnalytics4Metric, 0, len(inputOptionMetrics))
	for _, input := range inputOptionMetrics {
		column := GoogleAnalytics4Metric{
			Name:       types.StringValue(input.Name),
			Expression: types.StringValue(input.Expression),
		}
		metrics = append(metrics, column)
	}
	return metrics
}

func newGoogleAnalytics4InputOptionColumns(inputOptionColumns []input_option.GoogleAnalytics4Column) []GoogleAnalytics4Column {
	if inputOptionColumns == nil {
		return nil
	}
	columns := make([]GoogleAnalytics4Column, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := GoogleAnalytics4Column{
			Name: types.StringValue(input.Name),
			Type: types.StringValue(input.Type),
		}
		columns = append(columns, column)
	}
	return columns
}

func (inputOption *GoogleAnalytics4InputOption) ToInput() *param.GoogleAnalytics4InputOptionInput {
	if inputOption == nil {
		return nil
	}

	return &param.GoogleAnalytics4InputOptionInput{
		GoogleAnalytics4ConnectionID:          inputOption.GoogleAnalytics4ConnectionID.ValueInt64(),
		PropertyID:                            inputOption.PropertyID.ValueString(),
		TimeSeries:                            inputOption.TimeSeries.ValueString(),
		StartDate:                             inputOption.StartDate.ValueStringPointer(),
		EndDate:                               inputOption.EndDate.ValueStringPointer(),
		IncrementalLoadingEnabled:             inputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		RetryLimit:                            inputOption.RetryLimit.ValueInt64Pointer(),
		RetrySleep:                            inputOption.RetrySleep.ValueInt64Pointer(),
		RaiseOnOtherRow:                       inputOption.RaiseOnOtherRow.ValueBoolPointer(),
		LimitOfRows:                           inputOption.LimitOfRows.ValueInt64Pointer(),
		GoogleAnalytics4InputOptionDimensions: toGoogleAnalytics4DimensionsInput(inputOption.GoogleAnalytics4Dimensions),
		GoogleAnalytics4InputOptionMetrics:    toGoogleAnalytics4MetricsInput(inputOption.GoogleAnalytics4Metrics),
		InputOptionColumns:                    toGoogleAnalytics4ColumnsInput(inputOption.InputOptionColumns),
		CustomVariableSettings:                model.ToCustomVariableSettingInputs(inputOption.CustomVariableSettings),
	}
}

func (inputOption *GoogleAnalytics4InputOption) ToUpdateInput() *param.UpdateGoogleAnalytics4InputOptionInput {
	if inputOption == nil {
		return nil
	}

	dimensions := toGoogleAnalytics4DimensionsInput(inputOption.GoogleAnalytics4Dimensions)
	metrics := toGoogleAnalytics4MetricsInput(inputOption.GoogleAnalytics4Metrics)
	columns := toGoogleAnalytics4ColumnsInput(inputOption.InputOptionColumns)

	return &param.UpdateGoogleAnalytics4InputOptionInput{
		GoogleAnalytics4ConnectionID:          inputOption.GoogleAnalytics4ConnectionID.ValueInt64Pointer(),
		PropertyID:                            inputOption.PropertyID.ValueStringPointer(),
		TimeSeries:                            inputOption.TimeSeries.ValueStringPointer(),
		StartDate:                             inputOption.StartDate.ValueStringPointer(),
		EndDate:                               inputOption.EndDate.ValueStringPointer(),
		IncrementalLoadingEnabled:             inputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		RetryLimit:                            inputOption.RetryLimit.ValueInt64Pointer(),
		RetrySleep:                            inputOption.RetrySleep.ValueInt64Pointer(),
		RaiseOnOtherRow:                       inputOption.RaiseOnOtherRow.ValueBoolPointer(),
		LimitOfRows:                           inputOption.LimitOfRows.ValueInt64Pointer(),
		GoogleAnalytics4InputOptionDimensions: model.WrapObjectList(&dimensions),
		GoogleAnalytics4InputOptionMetrics:    &metrics,
		InputOptionColumns:                    &columns,
		CustomVariableSettings:                model.ToCustomVariableSettingInputs(inputOption.CustomVariableSettings),
	}
}

func toGoogleAnalytics4ColumnsInput(columns []GoogleAnalytics4Column) []param.GoogleAnalytics4Column {
	if columns == nil {
		return nil
	}

	inputs := make([]param.GoogleAnalytics4Column, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, param.GoogleAnalytics4Column{
			Name: column.Name.ValueString(),
			Type: column.Type.ValueString(),
		})
	}
	return inputs
}

func toGoogleAnalytics4DimensionsInput(dimensions []GoogleAnalytics4Dimension) []param.GoogleAnalytics4Dimension {
	if dimensions == nil {
		return nil
	}

	inputs := make([]param.GoogleAnalytics4Dimension, 0, len(dimensions))
	for _, dimension := range dimensions {
		inputs = append(inputs, param.GoogleAnalytics4Dimension{
			Name:       dimension.Name.ValueString(),
			Expression: dimension.Expression.ValueStringPointer(),
		})
	}
	return inputs
}

func toGoogleAnalytics4MetricsInput(metrics []GoogleAnalytics4Metric) []param.GoogleAnalytics4Metric {
	if metrics == nil {
		return nil
	}

	inputs := make([]param.GoogleAnalytics4Metric, 0, len(metrics))
	for _, metric := range metrics {
		inputs = append(inputs, param.GoogleAnalytics4Metric{
			Name:       metric.Name.ValueString(),
			Expression: metric.Expression.ValueStringPointer(),
		})
	}
	return inputs
}
