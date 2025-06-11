package provider

import (
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

					resource.TestCheckResourceAttr(resourceName, "tasks.0.redshift_data_check_config.custom_variables.0.name", "$string$"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.redshift_data_check_config.custom_variables.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.redshift_data_check_config.custom_variables.0.value", "foo"),
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
					resource.TestCheckResourceAttr(resourceName, "notifications.0.type", "job_execution"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "slack"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.slack_config.message", "This is a multi-line message\nwith several lines\n  and some indentation\n    to test TrimmedStringType\n"),

					resource.TestCheckResourceAttr(resourceName, "notifications.1.type", "job_time_alert"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "email"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.email_config.message", "  This is another multi-line message\nwith leading and trailing whitespace\n  \n  to test TrimmedStringType\n  \n"),
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
					"notifications.0.slack_config.message",
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
