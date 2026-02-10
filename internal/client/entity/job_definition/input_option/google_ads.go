package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type GoogleAdsInputOption struct {
	CustomerID             string                          `json:"customer_id"`
	ResourceType           string                          `json:"resource_type"`
	StartDate              *string                         `json:"start_date"`
	EndDate                *string                         `json:"end_date"`
	GoogleAdsConnectionID  int64                           `json:"google_ads_connection_id"`
	InputOptionColumns     []GoogleAdsColumn               `json:"input_option_columns"`
	Conditions             []string                        `json:"conditions"`
	CustomVariableSettings *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type GoogleAdsColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
