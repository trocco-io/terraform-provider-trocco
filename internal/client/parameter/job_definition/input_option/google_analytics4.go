package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type GoogleAnalytics4InputOptionInput struct {
	GoogleAnalytics4ConnectionID          int64                                   `json:"google_analytics4_connection_id"`
	PropertyID                            string                                  `json:"property_id"`
	TimeSeries                            string                                  `json:"time_series"`
	StartDate                             *string                                 `json:"start_date,omitempty"`
	EndDate                               *string                                 `json:"end_date,omitempty"`
	IncrementalLoadingEnabled             *bool                                   `json:"incremental_loading_enabled,omitempty"`
	RetryLimit                            *int64                                  `json:"retry_limit,omitempty"`
	RetrySleep                            *int64                                  `json:"retry_sleep,omitempty"`
	RaiseOnOtherRow                       *bool                                   `json:"raise_on_other_row,omitempty"`
	LimitOfRows                           *int64                                  `json:"limit_of_rows,omitempty"`
	GoogleAnalytics4InputOptionDimensions []GoogleAnalytics4Dimension             `json:"google_analytics4_input_option_dimensions"`
	GoogleAnalytics4InputOptionMetrics    []GoogleAnalytics4Metric                `json:"google_analytics4_input_option_metrics"`
	InputOptionColumns                    []GoogleAnalytics4Column                `json:"input_option_columns"`
	CustomVariableSettings                *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateGoogleAnalytics4InputOptionInput struct {
	GoogleAnalytics4ConnectionID          *int64                                                   `json:"google_analytics4_connection_id,omitempty"`
	PropertyID                            *string                                                  `json:"property_id,omitempty"`
	TimeSeries                            *string                                                  `json:"time_series,omitempty"`
	StartDate                             *string                                                  `json:"start_date,omitempty"`
	EndDate                               *string                                                  `json:"end_date,omitempty"`
	IncrementalLoadingEnabled             *bool                                                    `json:"incremental_loading_enabled,omitempty"`
	RetryLimit                            *int64                                                   `json:"retry_limit,omitempty"`
	RetrySleep                            *int64                                                   `json:"retry_sleep,omitempty"`
	RaiseOnOtherRow                       *bool                                                    `json:"raise_on_other_row,omitempty"`
	LimitOfRows                           *int64                                                   `json:"limit_of_rows,omitempty"`
	GoogleAnalytics4InputOptionDimensions *parameter.NullableObjectList[GoogleAnalytics4Dimension] `json:"google_analytics4_input_option_dimensions,omitempty"`
	GoogleAnalytics4InputOptionMetrics    *[]GoogleAnalytics4Metric                                `json:"google_analytics4_input_option_metrics,omitempty"`
	InputOptionColumns                    *[]GoogleAnalytics4Column                                `json:"input_option_columns,omitempty"`
	CustomVariableSettings                *[]parameter.CustomVariableSettingInput                  `json:"custom_variable_settings,omitempty"`
}

type GoogleAnalytics4Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type GoogleAnalytics4Dimension struct {
	Name       string  `json:"name"`
	Expression *string `json:"expression,omitempty"`
}
type GoogleAnalytics4Metric struct {
	Name       string  `json:"name"`
	Expression *string `json:"expression,omitempty"`
}
