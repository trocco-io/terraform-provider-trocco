package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTeamResource(t *testing.T) {
	resourceName := "trocco_team.test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + LoadTextFile("testdata/team/basic_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test"),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "members.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "members.0.role", "team_admin"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"id"},
			},
			// Update testing
			{
				Config: providerConfig + LoadTextFile("testdata/team/basic_update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "updated"),
					resource.TestCheckResourceAttr(resourceName, "description", "updated"),
					resource.TestCheckResourceAttr(resourceName, "members.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "members.0.role", "team_admin"),
					resource.TestCheckResourceAttr(resourceName, "members.1.role", "team_member"),
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
				Config:      providerConfig + LoadTextFile("testdata/team/no_members.tf"),
				ExpectError: regexp.MustCompile(`Empty Team Members`),
			},
		},
	})
}

func TestAccTeamInvalidRole(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/team/no_admin.tf"),
				ExpectError: regexp.MustCompile(`Missing Team Admin`),
			},
		},
	})
}

func TestAccTeamWithUnknownValues(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test with unknown values from another resource (simulating plan phase)
			{
				Config: providerConfig + `
					# Create a team that will be referenced
					resource "trocco_team" "source" {
					  name = "source-team"
					  description = "source team"
					  members = [
						{
						  user_id = 10626
						  role = "team_admin"
						}
					  ]
					}

					# Create another team with unknown values during plan
					resource "trocco_team" "test" {
					  name = trocco_team.source.name  # This will be unknown during initial plan
					  description = "test team with unknown values"
					  members = [
						{
						  user_id = 10626
						  role = "team_admin"
						}
					  ]
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_team.test", "name", "source-team"),
					resource.TestCheckResourceAttr("trocco_team.test", "description", "test team with unknown values"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.#", "1"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.0.role", "team_admin"),
				),
			},
		},
	})
}

func TestAccTeamValidationErrors(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test validation error with multiple team_members but no team_admin
			{
				Config: providerConfig + `
					resource "trocco_team" "test" {
					  name = "test-no-admin"
					  description = "team without admin"
					  members = [
						{
						  user_id = 10626
						  role = "team_member"
						},
						{
						  user_id = 10652
						  role = "team_member"
						},
						{
						  user_id = 10653
						  role = "team_member"
						}
					  ]
					}
				`,
				ExpectError: regexp.MustCompile(`Missing Team Admin`),
			},
			// Test validation error with single team_member
			{
				Config: providerConfig + `
					resource "trocco_team" "test" {
					  name = "test-single-member"
					  description = "team with single non-admin member"
					  members = [
						{
						  user_id = 10652
						  role = "team_member"
						}
					  ]
					}
				`,
				ExpectError: regexp.MustCompile(`Missing Team Admin`),
			},
			// Test validation passes when team_admin is added
			{
				Config: providerConfig + `
					resource "trocco_team" "test" {
					  name = "test-with-admin"
					  description = "team with admin and members"
					  members = [
						{
						  user_id = 10626
						  role = "team_admin"
						},
						{
						  user_id = 10652
						  role = "team_member"
						}
					  ]
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_team.test", "name", "test-with-admin"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.#", "2"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.0.role", "team_admin"),
					resource.TestCheckResourceAttr("trocco_team.test", "members.1.role", "team_member"),
				),
			},
		},
	})
}
