package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccJobDefinitionResourceGoogleAdsToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.google_ads_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/google_ads_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Google Ads to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from Google Ads to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "google_ads"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					// Google Ads input option checks
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.customer_id", "1234567890"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.resource_type", "campaign"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.start_date", "2024-01-01"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.end_date", "2024-01-31"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.input_option_columns.#", "5"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.input_option_columns.0.name", "campaign.name"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.input_option_columns.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.input_option_columns.1.name", "campaign.id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.input_option_columns.1.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.input_option_columns.2.name", "campaign.create_time"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.input_option_columns.2.type", "timestamp"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.input_option_columns.2.format", "%Y-%m-%d %H:%M:%S"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.input_option_columns.3.name", "metrics.ctr"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.input_option_columns.3.type", "double"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.input_option_columns.4.name", "campaign.experiment_type"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.input_option_columns.4.type", "boolean"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.conditions.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.conditions.0", "campaign.status = 'ENABLED'"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_ads_input_option.conditions.1", "metrics.impressions > 0"),
					// Filter columns checks
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "campaign_name"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "campaign_name"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.name", "campaign_id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.2.type", "timestamp"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.3.type", "double"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.4.type", "boolean"),
					// Output option checks
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "google_ads_campaign_test"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "append"),
				),
			},
		},
	})
}
