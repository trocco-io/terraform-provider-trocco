package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceMongoDBToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.mongodb_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/mongodb_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "MongoDB to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from MongoDB to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mongodb"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.collection", "test_collection"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.query", "{\"status\": \"active\"}"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.incremental_loading_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.incremental_columns", "created_at"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.last_record", "{\"created_at\":\"2024-01-01 00:00:00\"}"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.input_option_columns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.input_option_columns.0.name", "_id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.input_option_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.input_option_columns.1.name", "created_at"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.input_option_columns.1.type", "timestamp"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.input_option_columns.1.format", "%Y-%m-%d %H:%M:%S"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mongodb_input_option.input_option_columns.1.timezone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "mongodb_test_table"),
				),
			},
			// ImportState testing
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
