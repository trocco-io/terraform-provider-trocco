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
