package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourcePostgresqlToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.postgresql_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/postgresql_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "PostgreSQL to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from PostgreSQL to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "3"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "postgresql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.postgresql_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.postgresql_input_option.schema", "public"),
					resource.TestCheckResourceAttr(resourceName, "input_option.postgresql_input_option.fetch_rows", "1000"),
					resource.TestCheckResourceAttr(resourceName, "input_option.postgresql_input_option.default_time_zone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "postgresql_to_bigquery_test_table"),
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

func TestAccJobDefinitionResourceMysqlToPostgresql(t *testing.T) {
	resourceName := "trocco_job_definition.mysql_to_postgresql"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_postgresql/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test job_definition"),
					resource.TestCheckResourceAttr(resourceName, "description", "test description"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "postgresql"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.table", "test_table"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.connect_timeout", "300"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.socket_timeout", "1801"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.default_time_zone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "output_option.postgresql_output_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "output_option.postgresql_output_option.schema", "public"),
					resource.TestCheckResourceAttr(resourceName, "output_option.postgresql_output_option.table", "test_table"),
					resource.TestCheckResourceAttr(resourceName, "output_option.postgresql_output_option.mode", "merge"),
					resource.TestCheckResourceAttr(resourceName, "output_option.postgresql_output_option.default_time_zone", "UTC"),
					resource.TestCheckResourceAttrSet(resourceName, "output_option.postgresql_output_option.postgresql_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.postgresql_output_option.merge_keys.#", "1"),
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
