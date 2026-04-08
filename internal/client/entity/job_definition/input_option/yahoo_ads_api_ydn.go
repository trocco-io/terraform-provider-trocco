package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type YahooAdsApiYdnInputOption struct {
	YahooAdsApiConnectionID          int64                             `json:"yahoo_ads_api_connection_id"`
	Target                           string                            `json:"target"`
	AccountID                        string                            `json:"account_id"`
	BaseAccountID                    *string                           `json:"base_account_id"`
	ReportType                       *string                           `json:"report_type"`
	StartDate                        string                            `json:"start_date"`
	EndDate                          string                            `json:"end_date"`
	IncludeDeleted                   bool                              `json:"include_deleted"`
	YahooAdsApiYdnInputOptionColumns []YahooAdsApiYdnInputOptionColumn `json:"input_option_columns"`
	CustomVariableSettings           *[]entity.CustomVariableSetting   `json:"custom_variable_settings"`
}

type YahooAdsApiYdnInputOptionColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
