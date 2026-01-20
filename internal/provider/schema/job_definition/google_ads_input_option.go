package job_definition

import (
	planModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func GoogleAdsInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes about source Google Ads",
		Attributes: map[string]schema.Attribute{
			"customer_id": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Google Ads Customer ID (10-digit number)",
			},
			"resource_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"account_budget", "account_budget_proposal", "account_link", "ad_group", "ad_group_ad",
						"ad_group_ad_asset_view", "ad_group_ad_label", "ad_group_asset", "ad_group_audience_view",
						"ad_group_bid_modifier", "ad_group_criterion", "ad_group_criterion_label",
						"ad_group_criterion_simulation", "ad_group_label", "ad_group_simulation", "ad_parameter",
						"ad_schedule_view", "age_range_view", "asset", "asset_group_asset", "batch_job",
						"bidding_strategy", "billing_setup", "campaign", "campaign_asset", "campaign_audience_view",
						"campaign_bid_modifier", "campaign_budget", "campaign_criterion", "campaign_draft",
						"campaign_label", "campaign_shared_set", "carrier_constant", "change_event", "change_status",
						"click_view", "conversion_action", "currency_constant", "custom_interest", "customer",
						"customer_asset", "customer_client", "customer_client_link", "customer_label",
						"customer_manager_link", "customer_negative_criterion", "detail_placement_view",
						"display_keyword_view", "distance_view", "domain_category",
						"dynamic_search_ads_search_term_view", "expanded_landing_page_view", "gender_view",
						"geo_target_constant", "geographic_view", "group_placement_view", "hotel_group_view",
						"hotel_performance_view", "income_range_view", "keyword_plan", "keyword_plan_ad_group",
						"keyword_plan_ad_group_keyword", "keyword_plan_campaign", "keyword_plan_campaign_keyword",
						"keyword_view", "label", "landing_page_view", "language_constant", "location_view",
						"managed_placement_view", "media_file", "mobile_app_category_constant",
						"mobile_device_constant", "offline_user_data_job", "operating_system_version_constant",
						"paid_organic_search_term_view", "parental_status_view", "product_group_view",
						"recommendation", "remarketing_action", "search_term_view", "shared_criterion", "shared_set",
						"shopping_performance_view", "third_party_app_analytics_link", "topic_constant", "topic_view",
						"user_interest", "user_list", "user_location_view", "video",
					),
				},
				MarkdownDescription: "Resource type to retrieve (e.g., customer, campaign, ad_group, ad_group_ad)",
			},
			"start_date": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Data retrieval start date (YYYY-MM-DD format). May not be specified for some resource types.",
			},
			"end_date": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Data retrieval end date (YYYY-MM-DD format). May not be specified for some resource types.",
			},
			"google_ads_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Id of Google Ads connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"input_option_columns": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "List of fields to be retrieved",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
							MarkdownDescription: "Field name",
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Data type",
							Validators: []validator.String{
								stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean", "json"),
							},
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Format for timestamp type",
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"conditions": schema.ListAttribute{
				Optional:            true,
				Computed:            true,
				ElementType:         schema.StringAttribute{}.GetType(),
				MarkdownDescription: "List of WHERE clause conditions",
				PlanModifiers: []planmodifier.List{
					planModifier.EmptyListForNull(),
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
