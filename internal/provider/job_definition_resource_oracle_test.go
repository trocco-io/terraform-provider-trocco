package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceOracleIncrementalToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.oracle_incremental_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_incremental_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "oracle"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.schema", "test_schema"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.table", "test_table"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.incremental_loading_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.incremental_columns", "updated_at"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.last_record", "2026-01-01 00:00:00"),
					resource.TestCheckNoResourceAttr(resourceName, "input_option.oracle_input_option.query"),
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

func TestAccJobDefinitionResourceOracleTnsToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.oracle_tns_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_tns_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "oracle"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.net_service_name", "orcl_high"),
					resource.TestCheckNoResourceAttr(resourceName, "input_option.oracle_input_option.database"),
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

func TestAccJobDefinitionResourceOracleDefaultsToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.oracle_defaults_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_defaults_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// connection_type / incremental_loading_enabled を config で省略した場合、
					// デフォルト値（sid / false）が適用され、Provider produced inconsistent result に
					// ならないことを確認する。source_time_zone はデフォルト値を持たないため、
					// 省略時は未設定のまま（API側ではnull、実行時は転送元マシンのタイムゾーンとして解釈）
					// であることを確認する。
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.connection_type", "sid"),
					resource.TestCheckNoResourceAttr(resourceName, "input_option.oracle_input_option.source_time_zone"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.incremental_loading_enabled", "false"),
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

func TestAccJobDefinitionResourceOracleInputOptionConflict(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// input_option_type = "oracle" なのに mysql_input_option も
				// 同時に設定した場合、InputOptionPlanModifier によって
				// プラン時にエラーになることを確認する。
				Config:      providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_input_option_conflict/create.tf"),
				ExpectError: regexp.MustCompile(`(?s)invalid attribute.*mysql_input_option`),
			},
		},
	})
}

func TestAccJobDefinitionResourceOracleToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.oracle_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Oracle to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from Oracle to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "3"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "oracle"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.connection_type", "sid"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.source_time_zone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.input_option_columns.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.input_option_columns.3.timezone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "oracle_to_bigquery_test_table"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "append"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.location", "US"),
				),
			},
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_to_bigquery/update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Oracle to BigQuery Test (Updated)"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "5"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.database", "test_database_updated"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.source_time_zone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.input_option_columns.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.input_option_columns.2.timezone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "oracle_to_bigquery_test_table_updated"),
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
