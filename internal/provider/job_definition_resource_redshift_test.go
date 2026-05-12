package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceRedshiftToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.redshift_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/redshift_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Redshift to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from Redshift to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "redshift"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.redshift_input_option.database", "analytics"),
					resource.TestCheckResourceAttr(resourceName, "input_option.redshift_input_option.query", "SELECT * FROM test_table WHERE status = 'active'"),
					resource.TestCheckResourceAttr(resourceName, "input_option.redshift_input_option.schema", "public"),
					resource.TestCheckResourceAttr(resourceName, "input_option.redshift_input_option.fetch_rows", "1000"),
					resource.TestCheckResourceAttr(resourceName, "input_option.redshift_input_option.connect_timeout", "30"),
					resource.TestCheckResourceAttr(resourceName, "input_option.redshift_input_option.socket_timeout", "60"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.name", "created_at"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.type", "timestamp"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "redshift_test_table"),
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

func TestAccJobDefinitionResourceMysqlToRedshift(t *testing.T) {
	resourceName := "trocco_job_definition.mysql_to_redshift"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_redshift/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "MySQL to Redshift Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from MySQL to Redshift"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "redshift"),

					// Check MySQL input option
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.table", "test_table"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.connect_timeout", "300"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.socket_timeout", "1800"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.incremental_loading_enabled", "false"),

					// Check Redshift output option
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.database", "analytics"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.schema", "$schema$"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.table", "users"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.mode", "merge"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.s3_bucket", "my-redshift-bucket"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.s3_key_prefix", "/redshift-temp"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.delete_s3_temp_file", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.retry_limit", "12"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.retry_wait", "1000"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.max_retry_wait", "1800000"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.default_time_zone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.batch_size", "1024"),

					// Check custom variables and merge settings
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.custom_variable_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.custom_variable_settings.0.name", "$schema$"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.custom_variable_settings.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.custom_variable_settings.0.value", "public"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.redshift_output_option_merge_keys.#", "2"),

					// Check column options
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.redshift_output_option_column_options.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.redshift_output_option_column_options.0.name", "user_id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.redshift_output_option_column_options.0.type", "BIGINT"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.redshift_output_option_column_options.0.value_type", "long"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.redshift_output_option_column_options.1.name", "user_name"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.redshift_output_option_column_options.1.type", "VARCHAR"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.redshift_output_option_column_options.1.value_type", "string"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.redshift_output_option_column_options.2.name", "created_timestamp"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.redshift_output_option_column_options.2.type", "TIMESTAMP"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.redshift_output_option_column_options.2.value_type", "timestamp"),
					resource.TestCheckResourceAttr(resourceName, "output_option.redshift_output_option.redshift_output_option_column_options.2.timezone", "Asia/Tokyo"),

					// Check filter columns
					resource.TestCheckResourceAttr(resourceName, "filter_columns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.name", "type"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.src", "type"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.type", "string"),
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
