package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type GoogleAnalytics4InputOption struct {
	GoogleAnalytics4ConnectionID int64                           `json:"google_analytics4_connection_id"`
	PropertyID                   string                          `json:"property_id"`
	TimeSeries                   string                          `json:"time_series"`
	StartDate                    *string                         `json:"start_date"`
	EndDate                      *string                         `json:"end_date"`
	IncrementalLoadingEnabled    *bool                           `json:"incremental_loading_enabled"`
	RetryLimit                   *int64                          `json:"retry_limit"`
	RetrySleep                   *int64                          `json:"retry_sleep"`
	RaiseOnOtherRow              *bool                           `json:"raise_on_other_row"`
	LimitOfRows                  *int64                          `json:"limit_of_rows"`
	GoogleAnalytics4Dimensions   []GoogleAnalytics4Dimension     `json:"google_analytics4_input_option_dimensions"`
	GoogleAnalytics4Metrics      []GoogleAnalytics4Metric        `json:"google_analytics4_input_option_metrics"`
	InputOptionColumns           []GoogleAnalytics4Column        `json:"input_option_columns"`
	CustomVariableSettings       *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type GoogleAnalytics4Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type GoogleAnalytics4Dimension struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}
type GoogleAnalytics4Metric struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}
