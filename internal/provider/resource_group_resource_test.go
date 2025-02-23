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
				  resource "trocco_team" "team1" {
					  name = "test"
					  members = [
					    {
					      user_id = 10626
					      role = "team_admin"
					    }
					  ]
					}
					resource "trocco_resource_group" "test" {
					  name = "test"
					  description = "test"
					  teams = [
						{
						  team_id = trocco_team.team1.id
						  role = "operator"
						}
					  ]
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_resource_group.test", "name", "test"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "description", "test"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "teams.#", "1"),
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
				        resource "trocco_team" "team1" {
					  name = "test1"
					  members = [
						{
						  user_id = 10626
						  role = "team_admin"
			                        }
					  ]
			                }
					resource "trocco_resource_group" "test" {
					  name = "updated"
					  description = "updated"
					  teams = [
						{
						  team_id = trocco_team.team1.id
						  role = "administrator"
						}
					  ]
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_resource_group.test", "name", "updated"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "description", "updated"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "teams.#", "1"),
					resource.TestCheckResourceAttr("trocco_resource_group.test", "teams.0.role", "administrator"),
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
				        resource "trocco_team" "team1" {
					  name = "test"
					  description = "test"
					  members = [
						{
						  user_id = 10626
						  role = "team_admin"
			                        }
					  ]
			                }
					resource "trocco_resource_group" "test" {
					  name = "test"
					  description = "test"
					  teams = [
					    { team_id = trocco_team.team1.id, role = "administrator" },
					    { team_id = trocco_team.team1.id, role = "operator" },
					  ]
					}
				`,
				ExpectError: regexp.MustCompile(`Team ID "<unknown>" is duplicated in the list.`),
			},
		},
	})

}
