package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceYahooAdsApiYdnReportToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.yahoo_ads_report_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/yahoo_ads_api_ydn_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Yahoo Ads Report to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "description", "Daily report data from Yahoo Ads YDN synced to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "yahoo_ads_api_ydn"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttrSet(resourceName, "input_option.yahoo_ads_api_ydn_input_option.yahoo_ads_api_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.target", "report"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.base_account_id", "1234567890"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.account_id", "1234567890"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.start_date", "2024-01-01"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.end_date", "2024-01-31"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.include_deleted", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.input_option_columns.0.name", "DATE"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.input_option_columns.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.input_option_columns.1.name", "CAMPAIGN_ID"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.input_option_columns.1.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "yahoo_ads"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "report"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "merge"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.location", "US"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.timeout_sec", "300"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.retries", "5"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "DATE"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "DATE"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "string"),
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

func TestAccJobDefinitionResourceYahooAdsApiYdnStatsToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.yahoo_ads_stats_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/yahoo_ads_api_ydn_stats_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Yahoo Ads Campaign Stats to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "description", "Campaign stats data from Yahoo Ads YDN synced to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "3"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "yahoo_ads_api_ydn"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttrSet(resourceName, "input_option.yahoo_ads_api_ydn_input_option.yahoo_ads_api_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.target", "stats"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.report_type", "CAMPAIGN"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.base_account_id", "1234567890"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.account_id", "1234567890"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.start_date", "2024-01-01"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.end_date", "2024-01-31"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.include_deleted", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.input_option_columns.0.name", "CAMPAIGN_NAME"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.input_option_columns.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.input_option_columns.1.name", "CAMPAIGN_ID"),
					resource.TestCheckResourceAttr(resourceName, "input_option.yahoo_ads_api_ydn_input_option.input_option_columns.1.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "yahoo_ads"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "campaign_stats"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "replace"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.location", "US"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.auto_create_dataset", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.timeout_sec", "600"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.retries", "3"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "CAMPAIGN_NAME"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "CAMPAIGN_NAME"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "string"),
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
