package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type YahooAdsApiYssInputOptionInput struct {
	AccountID               string                                  `json:"account_id"`
	BaseAccountID           string                                  `json:"base_account_id"`
	Service                 *parameter.NullableString               `json:"service"`
	ExcludeZeroImpressions  bool                                    `json:"exclude_zero_impressions"`
	ReportType              *parameter.NullableString               `json:"report_type,omitempty"`
	StartDate               *parameter.NullableString               `json:"start_date,omitempty"`
	EndDate                 *parameter.NullableString               `json:"end_date,omitempty"`
	YahooAdsApiConnectionID int64                                   `json:"yahoo_ads_api_connection_id"`
	InputOptionColumns      []YahooAdsApiYssInputOptionColumn       `json:"input_option_columns"`
	CustomVariableSettings  *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateYahooAdsApiYssInputOptionInput struct {
	AccountID               *parameter.NullableString               `json:"account_id,omitempty"`
	BaseAccountID           *parameter.NullableString               `json:"base_account_id,omitempty"`
	Service                 *parameter.NullableString               `json:"service,omitempty"`
	ExcludeZeroImpressions  *parameter.NullableBool                 `json:"exclude_zero_impressions,omitempty"`
	ReportType              *parameter.NullableString               `json:"report_type,omitempty"`
	StartDate               *parameter.NullableString               `json:"start_date,omitempty"`
	EndDate                 *parameter.NullableString               `json:"end_date,omitempty"`
	YahooAdsApiConnectionID *parameter.NullableInt64                `json:"yahoo_ads_api_connection_id,omitempty"`
	InputOptionColumns      []YahooAdsApiYssInputOptionColumn       `json:"input_option_columns,omitempty"`
	CustomVariableSettings  *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type YahooAdsApiYssInputOptionColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
