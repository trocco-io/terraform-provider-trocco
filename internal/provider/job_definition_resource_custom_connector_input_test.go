package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceCustomConnectorToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.custom_connector_to_bigquery"
	inputOption := "input_option.custom_connector_input_option"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/custom_connector_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "custom_connector"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttrSet(resourceName, inputOption+".custom_connector_endpoint_id"),
					resource.TestCheckResourceAttrSet(resourceName, inputOption+".custom_connector_connection_id"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".query_parameters.0.name", "limit"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".query_parameters.0.value", "100"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".request_body_parameters.0.name", "status"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".request_body_parameters.0.value", "pending"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".endpoint_path", "/users"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".endpoint_method", "GET"),
					resource.TestCheckResourceAttrSet(resourceName, inputOption+".url"),
					resource.TestCheckResourceAttrSet(resourceName, inputOption+".auth_type"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".jsonpath_parser.root", "$.data"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".jsonpath_parser.columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".paginator.strategy_type", "page_increment"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".paginator.page_increment_strategy.page_size", "100"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".request_timeout_sec", "45"),
				),
			},
			{
				// query_parameters and request_body_parameters are omitted in
				// update.tf: verify omitting the key keeps the previously
				// configured values.
				Config: providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/custom_connector_to_bigquery/update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Custom Connector to BigQuery Test (Updated)"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".query_parameters.0.name", "limit"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".query_parameters.0.value", "100"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".request_body_parameters.0.value", "pending"),
				),
			},
			{
				// The job definition is pointed at another endpoint of the
				// same custom connector definition: the API rebuilds the
				// snapshot from the newly referenced endpoint while updating
				// the job definition, so the plan must not pin the snapshot
				// attributes to their prior values.
				//
				// The definition itself is deliberately identical in every
				// step of this test. TROCCO rebuilds the snapshot of every job
				// definition referencing a custom connector asynchronously
				// whenever that definition changes, which races with a job
				// definition update sent in the same apply and intermittently
				// fails it. Switching endpoints covers the same provider code
				// path - ModifyPlan keys off "an update is planned", not off
				// what triggered the new snapshot - without that race.
				Config: providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/custom_connector_to_bigquery/switch_endpoint.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Custom Connector to BigQuery Test (Switched)"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".endpoint_path", "/v2/users"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".success_codes", "200,201"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".request_timeout_sec", "60"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".paginator.page_increment_strategy.stop_on_page", "20"),
					// Omitted named value lists survive the rebuilt snapshot.
					resource.TestCheckResourceAttr(resourceName, inputOption+".query_parameters.0.value", "100"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".request_body_parameters.0.value", "pending"),
				),
			},
			{
				// query_parameters = [] and request_body_parameters = [] in
				// clear.tf: verify an explicit empty list clears the previously
				// configured values.
				Config: providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/custom_connector_to_bigquery/clear.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, inputOption+".query_parameters.#", "0"),
					resource.TestCheckResourceAttr(resourceName, inputOption+".request_body_parameters.#", "0"),
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
