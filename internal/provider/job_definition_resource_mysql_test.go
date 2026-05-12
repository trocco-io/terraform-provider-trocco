package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceMysqlToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.mysql_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test job_definition"),
					resource.TestCheckResourceAttr(resourceName, "description", "test description"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "large"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources["trocco_job_definition.mysql_to_bigquery"].Primary.ID

					return jobDefinitionId, nil
				},
			},
		},
	})
}

func TestAccJobDefinitionResourceKintoneToMysql(t *testing.T) {
	resourceName := "trocco_job_definition.kintone_to_mysql"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/kintone_to_mysql/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Kintone to Mysql Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for Kintone to Mysql transfer"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "0"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "false"),

					// Input option checks
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "kintone"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.app_id", "403"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.expand_subtable", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.#", "3"),

					// Output option checks - Mysql specific
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.database", "$db_name$"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.table", "$table_name$"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.mode", "insert"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.retry_limit", "12"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.retry_wait", "1000"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.max_retry_wait", "1800000"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.default_time_zone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.before_load", "DELETE FROM test_table WHERE status = 'pending';"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.after_load", "UPDATE test_table SET updated_at = NOW();"),

					// Column options checks
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.mysql_output_option_column_options.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.mysql_output_option_column_options.0.name", "description"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.mysql_output_option_column_options.0.type", "TEXT"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.mysql_output_option_column_options.1.name", "price"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.mysql_output_option_column_options.1.type", "DECIMAL"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.mysql_output_option_column_options.1.scale", "2"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.mysql_output_option_column_options.1.precision", "10"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.mysql_output_option_column_options.2.name", "notes"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.mysql_output_option_column_options.2.type", "LONGTEXT"),

					// Custom variable settings checks
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.custom_variable_settings.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.custom_variable_settings.0.name", "$db_name$"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.custom_variable_settings.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.custom_variable_settings.0.value", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.custom_variable_settings.1.name", "$table_name$"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.custom_variable_settings.1.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "output_option.mysql_output_option.custom_variable_settings.1.value", "test_table"),
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
