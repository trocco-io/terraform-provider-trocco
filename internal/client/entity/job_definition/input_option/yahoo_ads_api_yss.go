package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type YahooAdsApiYssInputOption struct {
	AccountID               string                            `json:"account_id"`
	BaseAccountID           string                            `json:"base_account_id"`
	Service                 string                            `json:"service"`
	ExcludeZeroImpressions  bool                              `json:"exclude_zero_impressions"`
	ReportType              *string                           `json:"report_type"`
	StartDate               *string                           `json:"start_date"`
	EndDate                 *string                           `json:"end_date"`
	YahooAdsApiConnectionID int64                             `json:"yahoo_ads_api_connection_id"`
	InputOptionColumns      []YahooAdsApiYssInputOptionColumn `json:"input_option_columns"`
	CustomVariableSettings  *[]entity.CustomVariableSetting   `json:"custom_variable_settings"`
}

type YahooAdsApiYssInputOptionColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
