package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceFacebookAdsInsightsToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.facebook_ads_insights_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/facebook_ads_insights_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Facebook Ads Insights to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from Facebook Ads Insights to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "facebook_ads_insights"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.facebook_ads_insights_connection_id", "922"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.ad_account_id", "act_123456789"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.level", "campaign"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.time_range_since", "2024-01-01"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.time_range_until", "2024-01-31"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.use_unified_attribution_setting", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.fields.#", "5"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.fields.0.name", "campaign_id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.fields.1.name", "campaign_name"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.fields.2.name", "impressions"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.fields.3.name", "clicks"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.fields.4.name", "spend"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.breakdowns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.breakdowns.0.name", "country"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.breakdowns.1.name", "device_platform"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.action_attribution_windows.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.action_attribution_windows.0.name", "1d_click"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.action_attribution_windows.1.name", "7d_click"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.action_breakdowns.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.action_breakdowns.0.name", "action_type"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.custom_variable_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.custom_variable_settings.0.name", "$window_start$"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.custom_variable_settings.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "input_option.facebook_ads_insights_input_option.custom_variable_settings.0.value", "2024-01-01"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "campaign_id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "campaign_id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.name", "impressions"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.src", "impressions"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "facebook_ads_insights_test"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "append"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.location", "US"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources[resourceName].Primary.ID
					return jobDefinitionId, nil
				},
			},
		},
	})
}
