package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type FacebookAdsInsightsInputOption struct {
	FacebookAdsInsightsConnectionID int64                                `json:"facebook_ads_insights_connection_id"`
	AdAccountID                     string                               `json:"ad_account_id"`
	Level                           string                               `json:"level"`
	TimeRangeSince                  string                               `json:"time_range_since"`
	TimeRangeUntil                  string                               `json:"time_range_until"`
	UseUnifiedAttributionSetting    bool                                 `json:"use_unified_attribution_setting"`
	Fields                          []FacebookAdsInsightsField           `json:"fields"`
	Breakdowns                      []FacebookAdsInsightsBreakdown       `json:"breakdowns"`
	ActionAttributionWindows        []FacebookAdsInsightsAttrWindow      `json:"action_attribution_windows"`
	ActionBreakdowns                []FacebookAdsInsightsActionBreakdown `json:"action_breakdowns"`
	CustomVariableSettings          *[]entity.CustomVariableSetting      `json:"custom_variable_settings"`
}

type FacebookAdsInsightsField struct {
	Name string `json:"name"`
}

type FacebookAdsInsightsBreakdown struct {
	Name string `json:"name"`
}

type FacebookAdsInsightsAttrWindow struct {
	Name string `json:"name"`
}

type FacebookAdsInsightsActionBreakdown struct {
	Name string `json:"name"`
}
