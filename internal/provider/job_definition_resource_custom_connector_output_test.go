package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceMysqlToCustomConnector(t *testing.T) {
	resourceName := "trocco_job_definition.mysql_to_custom_connector"
	outputOption := "output_option.custom_connector_output_option"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_custom_connector/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "custom_connector"),
					resource.TestCheckResourceAttrSet(resourceName, outputOption+".custom_connector_connection_id"),
					resource.TestCheckResourceAttrSet(resourceName, outputOption+".create_custom_connector_output_endpoint_id"),
					// mode defaults to insert, and the update-only attributes stay unset.
					resource.TestCheckResourceAttr(resourceName, outputOption+".mode", "insert"),
					resource.TestCheckNoResourceAttr(resourceName, outputOption+".update_key"),
					resource.TestCheckNoResourceAttr(resourceName, outputOption+".update_custom_connector_output_endpoint_id"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".create_endpoint_settings.query_parameters.0.name", "dry_run"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".create_endpoint_settings.query_parameters.0.value", "true"),
					// Snapshot of the referenced custom connector definition.
					resource.TestCheckResourceAttrSet(resourceName, outputOption+".url"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".auth_type", "api_key"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.0.operation", "create"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.0.method", "POST"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.0.path", "/users"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.0.request_timeout_sec", "45"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.0.query_parameters.0.name", "dry_run"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.0.query_parameters.0.value", "true"),
					// X-Api-Version is not editable, so the server fills in the
					// definition's default even though it is never configured here.
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.0.headers.0.name", "X-Api-Version"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.0.headers.0.value", "2"),
				),
			},
			{
				// Switch to upsert: update_key and the update endpoint become
				// mandatory, and a second endpoint snapshot appears.
				Config: providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_custom_connector/update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "MySQL to Custom Connector Test (Updated)"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".mode", "upsert"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".update_key", "id"),
					resource.TestCheckResourceAttrSet(resourceName, outputOption+".update_custom_connector_output_endpoint_id"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.#", "2"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.0.operation", "create"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.0.query_parameters.0.value", "false"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.1.operation", "update"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.1.method", "PATCH"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.1.path_parameters.0.name", "user_id"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.1.path_parameters.0.value", "id"),
				),
			},
			{
				// Switch back to insert: the server clears update_key and drops
				// the update endpoint snapshot. query_parameters = [] clears the
				// previously configured value of the create endpoint.
				Config: providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_custom_connector/revert_to_insert.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, outputOption+".mode", "insert"),
					resource.TestCheckNoResourceAttr(resourceName, outputOption+".update_key"),
					resource.TestCheckNoResourceAttr(resourceName, outputOption+".update_custom_connector_output_endpoint_id"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.#", "1"),
					resource.TestCheckResourceAttr(resourceName, outputOption+".endpoints.0.query_parameters.#", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// create_endpoint_settings/update_endpoint_settings are
				// configuration-only: the API reports the resulting values
				// through `endpoints` instead, so they cannot be recovered on
				// import.
				ImportStateVerifyIgnore: []string{
					outputOption + ".create_endpoint_settings",
					outputOption + ".update_endpoint_settings",
				},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources[resourceName].Primary.ID
					return jobDefinitionId, nil
				},
			},
		},
	})
}
