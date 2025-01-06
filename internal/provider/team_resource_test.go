package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTeamResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
					resource "trocco_team" "test" {
					  name = "test"
					  description = "test"
					  members = [
						{
						  user_id = 1
						  role = "team_admin"
						}
					  ]
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_team.test", "name", "test"),
					resource.TestCheckResourceAttr("trocco_team.test", "description", "test"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.#", "1"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.0.user_id", "1"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.0.role", "team_admin"),
				),
			},
			// ImportState testing
			{
				ResourceName:            "trocco_team.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"id"},
			},
			// Update testing
			{
				Config: providerConfig + `
					resource "trocco_team" "test" {
					  name = "updated"
					  description = "updated"
					  members = [
						{
						  user_id = 1
						  role = "team_admin"
						},
						{
						  user_id = 2
						  role = "team_member"
						}
					  ]
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_team.test", "name", "updated"),
					resource.TestCheckResourceAttr("trocco_team.test", "description", "updated"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.#", "2"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.0.user_id", "1"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.0.role", "team_admin"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.1.user_id", "2"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.1.role", "team_member"),
				),
			},
		},
	})
}

func TestAccTeamNoMembers(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_team" "test" {
					  name = "test"
					  description = "test"
					  members = []
					}
				`,
				ExpectError: regexp.MustCompile(`Missing Team Admin`),
			},
		},
	})
}

func TestAccTeamInvalidRole(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_team" "test" {
					  name = "test"
					  description = "test"
					  members = [
						{
						  user_id = 1
						  role = "team_member"
						}
					  ]
					}
				`,
				ExpectError: regexp.MustCompile(`Missing Team Admin`),
			},
		},
	})

}
