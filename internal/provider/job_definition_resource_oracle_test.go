package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceOracleIncrementalToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.oracle_incremental_to_bigquery"
	resource.ParallelTest(t, resource.TestCase{
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
				// Switches from table/incremental mode to query mode on update.
				// table/incremental_columns/last_record should be cleared (not left
				// at their prior values) once incremental_loading_enabled becomes false.
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_incremental_to_bigquery/switch_to_query_mode.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.incremental_loading_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "input_option.oracle_input_option.query"),
					resource.TestCheckNoResourceAttr(resourceName, "input_option.oracle_input_option.table"),
					resource.TestCheckNoResourceAttr(resourceName, "input_option.oracle_input_option.incremental_columns"),
					resource.TestCheckNoResourceAttr(resourceName, "input_option.oracle_input_option.last_record"),
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

func TestAccJobDefinitionResourceOracleMultiColumnIncrementalToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.oracle_multi_incremental_to_bigquery"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_multi_incremental_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.incremental_columns", "id,updated_at"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.last_record", "100,2026-01-01 00:00:00"),
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
	resource.ParallelTest(t, resource.TestCase{
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
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_defaults_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// When connection_type / incremental_loading_enabled are omitted from the
					// config, the defaults (sid / false) must be applied without causing
					// "Provider produced inconsistent result". source_time_zone has no default,
					// so it must stay unset when omitted (null on the API side; at runtime the
					// timezone of the machine running the transfer is used).
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

func TestAccJobDefinitionResourceOracleServiceNameToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.oracle_service_name_to_bigquery"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_service_name_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.connection_type", "service_name"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.custom_variable_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.custom_variable_settings.0.name", "$category$"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.custom_variable_settings.0.value", "electronics"),
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

func TestAccJobDefinitionResourceOracleUpdateConnectionMethod(t *testing.T) {
	resourceName := "trocco_job_definition.oracle_update_connection_method"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_update_connection_method/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.connection_type", "sid"),
					resource.TestCheckNoResourceAttr(resourceName, "input_option.oracle_input_option.net_service_name"),
				),
			},
			{
				// Switches the input_option from database/connection_type to
				// net_service_name on update. database should be cleared (not left
				// at its prior value); connection_type keeps its Computed default
				// ("sid") regardless, since it always has a server-side default.
				Config: providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_update_connection_method/update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.net_service_name", "orcl"),
					resource.TestCheckNoResourceAttr(resourceName, "input_option.oracle_input_option.database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.oracle_input_option.connection_type", "sid"),
				),
			},
		},
	})
}

func TestAccJobDefinitionResourceOracleMissingInputOptionColumns(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// input_option_columns is Required in the schema, so omitting it
				// fails at plan/validate time, not at apply time.
				Config:      providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_missing_input_option_columns/create.tf"),
				ExpectError: regexp.MustCompile(`(?s)input_option_columns.*required`),
			},
		},
	})
}

func TestAccJobDefinitionResourceOracleInputOptionConflict(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Setting mysql_input_option alongside input_option_type = "oracle"
				// must fail at plan time via InputOptionPlanModifier.
				Config:      providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_input_option_conflict/create.tf"),
				ExpectError: regexp.MustCompile(`(?s)invalid attribute.*mysql_input_option`),
			},
		},
	})
}

func TestAccJobDefinitionResourceOracleToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.oracle_to_bigquery"
	resource.ParallelTest(t, resource.TestCase{
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
				// Removes source_time_zone from config entirely after it was set to
				// "UTC" above. It has no server-side default (unlike connection_type),
				// so the API must be told explicitly to clear it; omitting the key from
				// the update request would leave the prior value in place and desync
				// state from config on every subsequent plan.
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/oracle_to_bigquery/update_clear_time_zone.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckNoResourceAttr(resourceName, "input_option.oracle_input_option.source_time_zone"),
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
