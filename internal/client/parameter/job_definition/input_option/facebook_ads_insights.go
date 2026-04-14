package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type FacebookAdsInsightsInputOptionInput struct {
	FacebookAdsInsightsConnectionID int64                                   `json:"facebook_ads_insights_connection_id"`
	AdAccountID                     string                                  `json:"ad_account_id"`
	Level                           string                                  `json:"level"`
	TimeRangeSince                  string                                  `json:"time_range_since"`
	TimeRangeUntil                  string                                  `json:"time_range_until"`
	UseUnifiedAttributionSetting    bool                                    `json:"use_unified_attribution_setting"`
	Fields                          []FacebookAdsInsightsField              `json:"fields"`
	Breakdowns                      []FacebookAdsInsightsBreakdown          `json:"breakdowns,omitempty"`
	ActionAttributionWindows        []FacebookAdsInsightsAttrWindow         `json:"action_attribution_windows,omitempty"`
	ActionBreakdowns                []FacebookAdsInsightsActionBreakdown    `json:"action_breakdowns,omitempty"`
	CustomVariableSettings          *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateFacebookAdsInsightsInputOptionInput struct {
	FacebookAdsInsightsConnectionID *parameter.NullableInt64                `json:"facebook_ads_insights_connection_id,omitempty"`
	AdAccountID                     *parameter.NullableString               `json:"ad_account_id,omitempty"`
	Level                           *parameter.NullableString               `json:"level,omitempty"`
	TimeRangeSince                  *parameter.NullableString               `json:"time_range_since,omitempty"`
	TimeRangeUntil                  *parameter.NullableString               `json:"time_range_until,omitempty"`
	UseUnifiedAttributionSetting    *bool                                   `json:"use_unified_attribution_setting,omitempty"`
	Fields                          *[]FacebookAdsInsightsField             `json:"fields,omitempty"`
	Breakdowns                      *[]FacebookAdsInsightsBreakdown         `json:"breakdowns,omitempty"`
	ActionAttributionWindows        *[]FacebookAdsInsightsAttrWindow        `json:"action_attribution_windows,omitempty"`
	ActionBreakdowns                *[]FacebookAdsInsightsActionBreakdown   `json:"action_breakdowns,omitempty"`
	CustomVariableSettings          *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
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
