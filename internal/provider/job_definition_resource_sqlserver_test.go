package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceSqlserverToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.sqlserver_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/sqlserver_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "SQL Server to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from SQL Server to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "3"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "sqlserver"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sqlserver_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sqlserver_input_option.schema", "dbo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sqlserver_input_option.incremental_loading_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sqlserver_input_option.fetch_rows", "1000"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sqlserver_input_option.connect_timeout", "300"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sqlserver_input_option.socket_timeout", "1801"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sqlserver_input_option.default_time_zone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sqlserver_input_option.input_option_columns.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sqlserver_input_option.input_option_column_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sqlserver_input_option.input_option_column_options.0.column_name", "created_at"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sqlserver_input_option.input_option_column_options.0.column_value_type", "string"),
					resource.TestCheckResourceAttrSet(resourceName, "input_option.sqlserver_input_option.sqlserver_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "sqlserver_to_bigquery_test_table"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "append"),
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

func TestJobDefinitionResourceSqlserverEmptyColumnOptions(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/job_definition/sqlserver_to_bigquery/empty_column_options.tf"),
				ExpectError: regexp.MustCompile(`input_option_column_options\s+list\s+must\s+contain\s+at\s+least\s+1\s+elements`),
			},
		},
	})
}
