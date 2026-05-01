package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func FacebookAdsInsightsInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes about source Facebook Ads Insights",
		Attributes: map[string]schema.Attribute{
			"facebook_ads_insights_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "ID of Facebook Ads Insights connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"ad_account_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Facebook Ad Account ID (format: act_XXXXXXXXX)",
			},
			"level": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Data retrieval level. Supported values: `account`, `campaign`, `adset`, `ad`",
				Validators: []validator.String{
					stringvalidator.OneOf("account", "campaign", "adset", "ad"),
				},
			},
			"time_range_since": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Start date for data retrieval (ISO8601 format: YYYY-MM-DD)",
			},
			"time_range_until": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "End date for data retrieval (ISO8601 format: YYYY-MM-DD)",
			},
			"use_unified_attribution_setting": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Use unified attribution setting. Default: true",
			},
			"fields": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "List of metrics/fields to be retrieved. At least one field is required.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Field name defined by Meta Ads Insights API",
							Validators: []validator.String{
								stringvalidator.OneOf(
									"account_currency", "account_id", "account_name", "action_values", "actions",
									"ad_click_actions", "ad_id", "ad_impression_actions", "ad_name",
									"adset_end", "adset_id", "adset_name", "adset_start",
									"age_targeting", "auction_bid", "auction_competitiveness", "auction_max_competitor_bid",
									"buying_type", "campaign_id", "campaign_name",
									"canvas_avg_view_percent", "canvas_avg_view_time", "clicks",
									"conversion_rate_ranking", "conversion_values", "conversions",
									"cost_per_15_sec_video_view", "cost_per_2_sec_continuous_video_view",
									"cost_per_action_type", "cost_per_ad_click", "cost_per_conversion",
									"cost_per_dda_countby_convs", "cost_per_estimated_ad_recallers",
									"cost_per_inline_link_click", "cost_per_inline_post_engagement",
									"cost_per_one_thousand_ad_impression", "cost_per_outbound_click",
									"cost_per_thruplay", "cost_per_unique_action_type", "cost_per_unique_click",
									"cost_per_unique_conversion", "cost_per_unique_inline_link_click",
									"cost_per_unique_outbound_click", "cpc", "cpm", "cpp",
									"created_time", "ctr", "date_start", "date_stop",
									"dda_countby_convs", "engagement_rate_ranking",
									"estimated_ad_recall_rate", "estimated_ad_recall_rate_lower_bound",
									"estimated_ad_recall_rate_upper_bound", "estimated_ad_recallers",
									"estimated_ad_recallers_lower_bound", "estimated_ad_recallers_upper_bound",
									"frequency", "full_view_impressions", "full_view_reach", "gender_targeting",
									"impressions", "inline_link_click_ctr", "inline_link_clicks", "inline_post_engagement",
									"instant_experience_clicks_to_open", "instant_experience_clicks_to_start",
									"instant_experience_outbound_clicks", "labels", "location",
									"mobile_app_purchase_roas", "objective", "outbound_clicks", "outbound_clicks_ctr",
									"place_page_name", "purchase_roas", "quality_ranking",
									"quality_score_ectr", "quality_score_ecvr", "quality_score_organic",
									"reach", "social_spend", "spend",
									"unique_actions", "unique_clicks", "unique_conversions", "unique_ctr",
									"unique_inline_link_click_ctr", "unique_inline_link_clicks",
									"unique_link_clicks_ctr", "unique_outbound_clicks", "unique_outbound_clicks_ctr",
									"unique_video_continuous_2_sec_watched_actions", "unique_video_view_15_sec",
									"updated_time", "video_15_sec_watched_actions", "video_30_sec_watched_actions",
									"video_avg_time_watched_actions", "video_continuous_2_sec_watched_actions",
									"video_p100_watched_actions", "video_p25_watched_actions",
									"video_p50_watched_actions", "video_p75_watched_actions",
									"video_p95_watched_actions", "video_play_actions",
									"video_thruplay_watched_actions", "video_time_watched_actions",
									"website_ctr", "website_purchase_roas", "wish_bid",
								),
							},
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"breakdowns": schema.ListNestedAttribute{
				Optional:            true,
				Computed:            true,
				Default:             listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType}}, []attr.Value{})),
				MarkdownDescription: "List of breakdowns for data segmentation",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Breakdown name defined by Meta Ads Insights API",
							Validators: []validator.String{
								stringvalidator.OneOf(
									"ad_format_asset", "age", "body_asset", "call_to_action_asset",
									"country", "description_asset", "device_platform", "dma",
									"frequency_value", "gender",
									"hourly_stats_aggregated_by_advertiser_time_zone",
									"hourly_stats_aggregated_by_audience_time_zone",
									"image_asset", "impression_device", "link_url_asset",
									"place_page_id", "platform_position", "product_id",
									"publisher_platform", "region", "title_asset", "video_asset",
								),
							},
						},
					},
				},
			},
			"action_attribution_windows": schema.ListNestedAttribute{
				Optional:            true,
				Computed:            true,
				Default:             listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType}}, []attr.Value{})),
				MarkdownDescription: "List of action attribution windows",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Attribution window value defined by Meta Ads API",
							Validators: []validator.String{
								stringvalidator.OneOf(
									"1d_view", "7d_view", "28d_view",
									"1d_click", "7d_click", "28d_click",
								),
							},
						},
					},
				},
			},
			"action_breakdowns": schema.ListNestedAttribute{
				Optional:            true,
				Computed:            true,
				Default:             listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{"name": types.StringType}}, []attr.Value{})),
				MarkdownDescription: "List of action breakdowns",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Action breakdown value defined by Meta Ads API",
							Validators: []validator.String{
								stringvalidator.OneOf(
									"action_device", "action_canvas_component_name",
									"action_carousel_card_id", "action_carousel_card_name",
									"action_destination", "action_reaction", "action_target_id",
									"action_type", "action_video_sound", "action_video_type",
								),
							},
						},
					},
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
