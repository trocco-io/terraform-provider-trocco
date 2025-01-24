package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
					resource "trocco_resource_group" "test" {
					  name = "test"
					  description = "test"
					  teams = [
						{
						  team_id = 1
						  role = "operator"
						}
					  ]
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_resource_group.test", "name", "test"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "description", "test"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "teams.#", "1"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "teams.0.team_id", "1"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "teams.0.role", "operator"),
				),
			},
			// ImportState testing
			{
				ResourceName:            "trocco_resource_group.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"id"},
			},
			// Update testing
			{
				Config: providerConfig + `
					resource "trocco_resource_group" "test" {
					  name = "updated"
					  description = "updated"
					  teams = [
						{
						  team_id = 1
						  role = "administrator"
						},
						{
						  team_id = 2
						  role = "operator"
						}
					  ]
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_resource_group.test", "name", "updated"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "description", "updated"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "teams.#", "2"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "teams.0.team_id", "1"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "teams.0.role", "administrator"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "teams.1.team_id", "2"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "teams.1.role", "operator"),
				),
			},
		},
	})
}

func TestAccResourceGroupNoTeams(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_resource_group" "test" {
					  name = "test"
					  description = "test"
					  teams = []
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_resource_group.test", "name", "test"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "description", "test"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "teams.#", "0"),
				),
			},
		},
	})
}

func TestAccResourceGroupDuplicateRoles(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_resource_group" "test" {
					  name = "test"
					  description = "test"
					  teams = [
					    { team_id = 1, role = "administrator" },
					    { team_id = 1, role = "operator" },
					  ]
					}
				`,
				ExpectError: regexp.MustCompile(`Team ID "1" is duplicated in the list.`),
			},
		},
	})

}

