package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceDatabricksToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.databricks_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/databricks_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test databricks_to_bigquery job"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for Databricks to BigQuery transfer"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "databricks"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					// Databricks input option attributes
					resource.TestCheckResourceAttr(resourceName, "input_option.databricks_input_option.catalog_name", "test_catalog"),
					resource.TestCheckResourceAttr(resourceName, "input_option.databricks_input_option.schema_name", "test_schema"),
					resource.TestCheckResourceAttrSet(resourceName, "input_option.databricks_input_option.databricks_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.databricks_input_option.input_option_columns.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "input_option.databricks_input_option.input_option_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.databricks_input_option.input_option_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "input_option.databricks_input_option.input_option_columns.1.name", "user_name"),
					resource.TestCheckResourceAttr(resourceName, "input_option.databricks_input_option.input_option_columns.1.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "input_option.databricks_input_option.input_option_columns.2.name", "email"),
					resource.TestCheckResourceAttr(resourceName, "input_option.databricks_input_option.input_option_columns.2.type", "string"),
					// BigQuery output option attributes
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "databricks_users"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "replace"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.location", "US"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.auto_create_dataset", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "output_option.bigquery_output_option.bigquery_connection_id"),
					// Filter columns
					resource.TestCheckResourceAttr(resourceName, "filter_columns.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.name", "user_name"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.2.name", "email"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.2.type", "string"),
				),
			},
			// Import testing
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

func TestAccJobDefinitionResourceMysqlToDatabricks(t *testing.T) {
	resourceName := "trocco_job_definition.mysql_to_databricks"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_databricks/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "MySQL to Databricks Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from MySQL to Databricks"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "databricks"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.table", "test_table"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.connect_timeout", "300"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.socket_timeout", "1800"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.default_time_zone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.catalog_name", "test_catalog"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.schema_name", "test_schema"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.table", "test_table"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.batch_size", "40000"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.mode", "merge"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.default_time_zone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.databricks_output_option_column_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.databricks_output_option_column_options.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.databricks_output_option_column_options.0.type", "TIMESTAMP"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.databricks_output_option_column_options.0.value_type", "timestamp"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.databricks_output_option_column_options.0.timestamp_format", "%Y-%m-%d %H:%M:%S"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.databricks_output_option_column_options.0.timezone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "output_option.databricks_output_option.databricks_output_option_merge_keys.#", "1"),
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
