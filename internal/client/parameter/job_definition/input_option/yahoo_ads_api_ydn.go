package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type YahooAdsApiYdnInputOptionInput struct {
	YahooAdsApiConnectionID          int64                                   `json:"yahoo_ads_api_connection_id"`
	Target                           string                                  `json:"target"`
	AccountID                        string                                  `json:"account_id"`
	BaseAccountID                    *parameter.NullableString               `json:"base_account_id,omitempty"`
	ReportType                       *parameter.NullableString               `json:"report_type,omitempty"`
	StartDate                        string                                  `json:"start_date"`
	EndDate                          string                                  `json:"end_date"`
	IncludeDeleted                   *parameter.NullableBool                 `json:"include_deleted,omitempty"`
	YahooAdsApiYdnInputOptionColumns []YahooAdsApiYdnInputOptionColumn       `json:"input_option_columns"`
	CustomVariableSettings           *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateYahooAdsApiYdnInputOptionInput struct {
	YahooAdsApiConnectionID          *parameter.NullableInt64                `json:"yahoo_ads_api_connection_id,omitempty"`
	Target                           *parameter.NullableString               `json:"target,omitempty"`
	AccountID                        *parameter.NullableString               `json:"account_id,omitempty"`
	BaseAccountID                    *parameter.NullableString               `json:"base_account_id,omitempty"`
	ReportType                       *parameter.NullableString               `json:"report_type,omitempty"`
	StartDate                        *parameter.NullableString               `json:"start_date,omitempty"`
	EndDate                          *parameter.NullableString               `json:"end_date,omitempty"`
	IncludeDeleted                   *parameter.NullableBool                 `json:"include_deleted,omitempty"`
	YahooAdsApiYdnInputOptionColumns []YahooAdsApiYdnInputOptionColumn       `json:"input_option_columns,omitempty"`
	CustomVariableSettings           *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type YahooAdsApiYdnInputOptionColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
