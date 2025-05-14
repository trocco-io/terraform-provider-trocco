package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPipelineDefinitionResourceForDatacheckBigquery(t *testing.T) {
	resourceName := "trocco_pipeline_definition.bigquery_data_check_query_check"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("../../examples/testdata/pipeline_definition/bigquery_data_check/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "bigquery_data_check"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.bigquery_data_check_config.query", "          SELECT COUNT(*) FROM examples\n"),
				),
			},
		},
	})
}

func TestAccPipelineDefinitionResourceForDatacheckSnowflake(t *testing.T) {
	resourceName := "trocco_pipeline_definition.snowflake_data_check_query_check"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("../../examples/testdata/pipeline_definition/snowflake_data_check/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "snowflake_data_check"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.snowflake_data_check_config.query", "          SELECT COUNT(*) FROM examples\n"),
				),
			},
		},
	})
}

func TestAccPipelineDefinitionResourceForDatacheckRedshift(t *testing.T) {
	resourceName := "trocco_pipeline_definition.redshift_data_check_query_check"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("../../examples/testdata/pipeline_definition/redshift_data_check/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "redshift_data_check"),
					resource.TestCheckResourceAttr(resourceName, "tasks.0.redshift_data_check_config.query", "          SELECT COUNT(*) FROM examples\n"),
				),
			},
		},
	})
}
