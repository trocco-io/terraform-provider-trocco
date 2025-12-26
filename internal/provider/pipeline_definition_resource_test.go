package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccPipelineDefinitionResourceForDataCheckBigquery(t *testing.T) {
	resourceName := "trocco_pipeline_definition.bigquery_data_check_query_check"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/pipeline_definition/bigquery_data_check/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "bigquery_data_check"),
					resource.TestCheckResourceAttr(resourceName, "description", "    This is a pipeline definition for BigQuery data check.\n    It checks if the count of rows in the 'examples' table equals 1.\n"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.bigquery_data_check_config.query", "          SELECT COUNT(*) FROM examples\n"),
				),
				ImportStateVerifyIgnore: []string{
					// The `key` attribute does not exist in the TROCCO API,
					// therefore there is no value for it during import.
					"tasks.0.key",
					// INFO: The `query` attribute is trimmed and set in state, so different from the resource config.
					"tasks.0.bigquery_data_check_config.query",
				},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					pipelineDefinitionID := s.RootModule().Resources[resourceName].Primary.ID
					return pipelineDefinitionID, nil
				},
			},
		},
	})
}

func TestAccPipelineDefinitionResourceForDataCheckSnowflake(t *testing.T) {
	resourceName := "trocco_pipeline_definition.snowflake_data_check_query_check"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/pipeline_definition/snowflake_data_check/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "snowflake_data_check"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.snowflake_data_check_config.query", "          SELECT COUNT(*) FROM examples\n"),
				),
			},
			// Import testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// The `key` attribute does not exist in the TROCCO API,
					// therefore there is no value for it during import.
					"tasks.0.key",
					// INFO: The `query` attribute is trimmed and set in state, so different from the resource config.
					"tasks.0.snowflake_data_check_config.query",
				},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					pipelineDefinitionID := s.RootModule().Resources[resourceName].Primary.ID
					return pipelineDefinitionID, nil
				},
			},
		},
	})
}

func TestAccPipelineDefinitionResourceForDataCheckRedshift(t *testing.T) {
	resourceName := "trocco_pipeline_definition.redshift_data_check_query_check"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/pipeline_definition/redshift_data_check/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "redshift_data_check"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.redshift_data_check_config.query", "          SELECT COUNT(*) FROM examples\n"),
				),
			},
			// Import testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// The `key` attribute does not exist in the TROCCO API,
					// therefore there is no value for it during import.
					"tasks.0.key",
					// INFO: The `query` attribute is trimmed and set in state, so different from the resource config.
					"tasks.0.redshift_data_check_config.query",
				},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					pipelineDefinitionID := s.RootModule().Resources[resourceName].Primary.ID
					return pipelineDefinitionID, nil
				},
			},
		},
	})
}

func TestAccPipelineDefinitionResourceForNotifications(t *testing.T) {
	resourceName := "trocco_pipeline_definition.notifications_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/pipeline_definition/notifications/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "notifications_test"),

					// Check slack notification
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "notifications.*", map[string]string{
						"type":                 "job_execution",
						"destination_type":     "slack",
						"slack_config.message": "This is a multi-line message\nwith several lines\n  and some indentation\n    to test TrimmedStringType\n",
					}),

					// Check email notification
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "notifications.*", map[string]string{
						"type":                 "job_time_alert",
						"destination_type":     "email",
						"email_config.message": "  This is another multi-line message\nwith leading and trailing whitespace\n  \n  to test TrimmedStringType\n  \n",
					}),
				),
			},
			// Import testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// The `key` attribute does not exist in the TROCCO API,
					// therefore there is no value for it during import.
					"tasks.0.key",
					// INFO: The message attributes are trimmed and set in state, so different from the resource config.
					// Explicitly specify all possible indices for both message types
					"notifications.0.slack_config.message",
					"notifications.1.slack_config.message",
					"notifications.0.email_config.message",
					"notifications.1.email_config.message",
					// INFO: The `query` attribute is trimmed and set in state, so different from the resource config.
					"tasks.0.bigquery_data_check_config.query",
				},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					pipelineDefinitionID := s.RootModule().Resources[resourceName].Primary.ID
					return pipelineDefinitionID, nil
				},
			},
		},
	})
}

func TestAccPipelineDefinitionResourceForCustomVariableLoopInvalidType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test invalid type
			{
				Config:      providerConfig + LoadTextFile("testdata/pipeline_definition/custom_variable_loop/invalid_type.tf"),
				ExpectError: regexp.MustCompile(`must be one of: \["string" "period" "bigquery" "snowflake" "redshift"\]`),
			},
		},
	})
}

func TestAccPipelineDefinitionResourceForCustomVariableLoopMissingConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test missing config
			{
				Config:      providerConfig + LoadTextFile("testdata/pipeline_definition/custom_variable_loop/missing_config.tf"),
				ExpectError: regexp.MustCompile(`Missing Required Configuration`),
			},
		},
	})
}

func TestAccPipelineDefinitionResourceForBigQueryDatamartWithStringLoop(t *testing.T) {
	resourceName := "trocco_pipeline_definition.trocco_bigquery_datamart"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/pipeline_definition/custom_variable_loop/valid_string_config.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "trocco_bigquery_datamart"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.key", "trocco_bigquery_datamart"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.string_config.variables.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.string_config.variables.0.name", "$foo$"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.string_config.variables.0.values.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.string_config.variables.1.name", "$bar$"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.string_config.variables.1.values.#", "2"),
				),
				ImportStateVerifyIgnore: []string{
					"tasks.0.key",
					"tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.string_config.variables.1.values",
				},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return s.RootModule().Resources[resourceName].Primary.ID, nil
				},
			},
		},
	})
}

func TestAccPipelineDefinitionResourceForBigQueryDatamartWithBigQueryLoop(t *testing.T) {
	resourceName := "trocco_pipeline_definition.trocco_bigquery_datamart"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/pipeline_definition/custom_variable_loop/valid_bigquery_config.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "trocco_bigquery_datamart"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.key", "trocco_bigquery_datamart"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.bigquery_config.query", "SELECT foo, bar FROM sample"),
					resource.TestCheckTypeSetElemAttr(resourceName,
						"tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.bigquery_config.variables.*",
						"$foo$",
					),
					resource.TestCheckTypeSetElemAttr(resourceName,
						"tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.bigquery_config.variables.*",
						"$bar$",
					),
				),
				ImportStateVerifyIgnore: []string{
					"tasks.0.key",
					"tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.bigquery_config.query",
					"tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.bigquery_config.variables",
				},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return s.RootModule().Resources[resourceName].Primary.ID, nil
				},
			},
		},
	})
}

func TestAccPipelineDefinitionResourceForIfElse(t *testing.T) {
	resourceName := "trocco_pipeline_definition.if_else_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/pipeline_definition/if_else/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "if_else_test"),
					resource.TestCheckResourceAttr(resourceName, "tasks.#", "4"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "tasks.*", map[string]string{
						"key":                 "if_else",
						"type":                "if_else",
						"if_else_config.name": "Check transfer status",
						"if_else_config.condition_groups.set_type": "and",
					}),
				),
			},
			// Import testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// The `key` attribute does not exist in the TROCCO API,
					// therefore there is no value for it during import.
					"tasks.0.key",
					"tasks.1.key",
					"tasks.2.key",
					"tasks.3.key",
					// task_dependencies use task keys which are not available during import
					"task_dependencies.0.source",
					"task_dependencies.0.destination",
					"task_dependencies.1.source",
					"task_dependencies.1.destination",
					"task_dependencies.2.source",
					"task_dependencies.2.destination",
					// if_else_config references task keys which are not available during import
					"tasks.0.if_else_config.condition_groups.conditions.0.task_key",
					"tasks.0.if_else_config.destinations.if.0",
					"tasks.0.if_else_config.destinations.else.0",
				},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					pipelineDefinitionID := s.RootModule().Resources[resourceName].Primary.ID
					return pipelineDefinitionID, nil
				},
			},
		},
	})
}
