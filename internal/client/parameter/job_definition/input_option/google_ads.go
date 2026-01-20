package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type GoogleAdsInputOptionInput struct {
	CustomerID             string                                  `json:"customer_id"`
	ResourceType           string                                  `json:"resource_type"`
	StartDate              *parameter.NullableString               `json:"start_date,omitempty"`
	EndDate                *parameter.NullableString               `json:"end_date,omitempty"`
	GoogleAdsConnectionID  int64                                   `json:"google_ads_connection_id"`
	InputOptionColumns     []GoogleAdsColumn                       `json:"input_option_columns"`
	Conditions             []string                                `json:"conditions,omitempty"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateGoogleAdsInputOptionInput struct {
	CustomerID             *parameter.NullableString               `json:"customer_id,omitempty"`
	ResourceType           *parameter.NullableString               `json:"resource_type,omitempty"`
	StartDate              *parameter.NullableString               `json:"start_date,omitempty"`
	EndDate                *parameter.NullableString               `json:"end_date,omitempty"`
	GoogleAdsConnectionID  *parameter.NullableInt64                `json:"google_ads_connection_id,omitempty"`
	InputOptionColumns     *[]GoogleAdsColumn                      `json:"input_option_columns,omitempty"`
	Conditions             *[]string                               `json:"conditions,omitempty"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type GoogleAdsColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format,omitempty"`
}
