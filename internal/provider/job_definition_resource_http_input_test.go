package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceHttpToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.http_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/http_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "HTTP to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from HTTP to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "http"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.http_input_option.url", "https://example.com"),
					resource.TestCheckResourceAttr(resourceName, "input_option.http_input_option.method", "GET"),
					resource.TestCheckResourceAttr(resourceName, "input_option.http_input_option.request_params.0.key", "foo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.http_input_option.request_params.0.value", "bar"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "http_to_bigquery_table"),
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

func TestAccJobDefinitionResourceHttpToBigQueryInvalid(t *testing.T) {
	resourceName := "trocco_job_definition.http_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/http_to_bigquery/conflict_body_and_param.tf"),
				PlanOnly:     true,
				ExpectError:  regexp.MustCompile(`Error: request_body conflicts with request_params`),
			},
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/http_to_bigquery/pager_offset_missing_param.tf"),
				PlanOnly:     true,
				ExpectError:  regexp.MustCompile(`Error: pager_from_param is required when pager_type is offset`),
			},
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/http_to_bigquery/pager_cursor_missing_param.tf"),
				PlanOnly:     true,
				ExpectError:  regexp.MustCompile(`Error: cursor_request_parameter_cursor_name is required when pager_type is cursor`),
			},
		},
	})
}
