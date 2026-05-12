package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceSnowflakeToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.snowflake_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/snowflake_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Snowflake to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from Snowflake to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "snowflake"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.snowflake_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.snowflake_input_option.schema", "PUBLIC"),
					resource.TestCheckResourceAttr(resourceName, "input_option.snowflake_input_option.query", "SELECT id, name, email, created_at FROM test_table"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "user_id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "snowflake_to_bigquery_test_table"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "append"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.bigquery_output_option_merge_keys.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.bigquery_output_option_clustering_fields.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.bigquery_output_option_column_options.#", "0"),
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

func TestAccJobDefinitionResourceBigQueryToSnowflake(t *testing.T) {
	resourceName := "trocco_job_definition.bigquery_to_snowflake"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/bigquery_to_snowflake/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "BigQuery to Snowflake Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from BigQuery to Snowflake"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "snowflake"),
					resource.TestCheckResourceAttr(resourceName, "input_option.bigquery_input_option.query", "SELECT id, name, email, created_at FROM `test_dataset.test_table`"),
					resource.TestCheckResourceAttr(resourceName, "input_option.bigquery_input_option.location", "US"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "user_id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "output_option.snowflake_output_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "output_option.snowflake_output_option.schema", "PUBLIC"),
					resource.TestCheckResourceAttr(resourceName, "output_option.snowflake_output_option.table", "bigquery_to_snowflake_test_table"),
					resource.TestCheckResourceAttr(resourceName, "output_option.snowflake_output_option.mode", "insert"),
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
