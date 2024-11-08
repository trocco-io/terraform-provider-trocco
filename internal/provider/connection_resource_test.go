package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccConnectionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_connection" "test" {
					  connection_type = "bigquery"

					  name = "test"
					  description = "The quick brown fox jumps over the lazy dog."

					  service_account_json_key = "{\"type\":\"service_account\",\"project_id\":\"\",\"private_key_id\":\"\",\"private_key\":\"\"}"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_connection.test", "connection_type", "bigquery"),
					resource.TestCheckResourceAttr("trocco_connection.test", "name", "test"),
					resource.TestCheckResourceAttr("trocco_connection.test", "description", "The quick brown fox jumps over the lazy dog."),
					resource.TestCheckResourceAttr("trocco_connection.test", "service_account_json_key", "{\"type\":\"service_account\",\"project_id\":\"\",\"private_key_id\":\"\",\"private_key\":\"\"}"),
					resource.TestCheckResourceAttrSet("trocco_connection.test", "id"),
				),
			},
			{
				ResourceName:            "trocco_connection.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_account_json_key"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					connectionID := s.RootModule().Resources["trocco_connection.test"].Primary.ID

					return fmt.Sprintf("bigquery,%s", connectionID), nil
				},
			},
		},
	})
}
