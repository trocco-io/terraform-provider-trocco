package job_definition

import (
	"regexp"
	troccoPlanModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func YahooAdsApiYssInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of source yahoo_ads_api_yss",
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "account id",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"base_account_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "base account id",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"service": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("report_definition_service"),
				MarkdownDescription: "service",
				Validators: []validator.String{
					stringvalidator.OneOf("report_definition_service", "campaign_export_service"),
				},
			},
			"report_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "report_type",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ACCOUNT",
						"CAMPAIGN",
						"ADGROUP",
						"AD",
						"KEYWORDS",
						"SEARCH_QUERY",
						"GEO",
						"GEO_TARGET",
						"SCHEDULE_TARGET",
						"BID_STRATEGY",
						"ADGROUP_TARGET_LIST",
						"CAMPAIGN_TARGET_LIST",
						"LANDING_PAGE_URL",
						"KEYWORDLESS_QUERY",
						"WEBPAGE_CRITERION",
						"BID_MODIFIER",
						"CAMPAIGN_ASSET",
						"ADGROUP_ASSET",
						"RESPONSIVE_ADS_FOR_SEARCH_ASSET",
						"ASSET_COMBINATIONS",
					),
				},
			},
			"start_date": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "start_date",
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^(.*\$.*|%[A-Za-z0-9%/:.\-]+|\d{8,20}|\d{4}[-/]\d{2}[-/]\d{2})$`),
						"must be a valid date or template string (e.g. 20240501, ${start_date}, $startdate$01, %Y/%m/%d)",
					),
				},
			},
			"end_date": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "end_date",
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^(.*\$.*|%[A-Za-z0-9%/:.\-]+|\d{8,20}|\d{4}[-/]\d{2}[-/]\d{2})$`),
						"must be a valid date or template string (e.g. 20240501, ${start_date}, $startdate$01, %Y/%m/%d)",
					),
				},
			},
			"exclude_zero_impressions": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "exclude_zero_impressions",
			},
			"yahoo_ads_api_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of yahoo_ads_api connection",
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
			"input_option_columns": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "List of columns to be retrieved and their types",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
							MarkdownDescription: "Column name",
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column type",
							Validators: []validator.String{
								stringvalidator.OneOf("boolean", "long", "timestamp", "double", "string", "json"),
							},
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Column format",
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
					},
					PlanModifiers: []planmodifier.Object{
						&troccoPlanModifier.YahooAdsApiYssInputOptionColumnPlanModifier{},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
		},
		PlanModifiers: []planmodifier.Object{
			&troccoPlanModifier.YahooAdsApiYssInputOptionPlanModifier{},
		},
	}
}
