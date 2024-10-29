package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUserResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
					resource "trocco_user" "test" {
					  email = "test@example.com"
					  password = "3XRambMkp-Hw"
					  role = "admin"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_user.test", "email", "test@example.com"),
					resource.TestCheckResourceAttr("trocco_user.test", "role", "admin"),
					resource.TestCheckResourceAttr("trocco_user.test", "can_use_audit_log", "false"),
					resource.TestCheckResourceAttr("trocco_user.test", "is_restricted_connection_modify", "false"),
					resource.TestCheckResourceAttrSet("trocco_user.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:            "trocco_user.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			// Update and Read testing
			{
				Config: providerConfig + `
					resource "trocco_user" "test" {
					  email = "test@example.com"
					  role = "member"
					  can_use_audit_log = true
					  is_restricted_connection_modify = true
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_user.test", "email", "test@example.com"),
					resource.TestCheckResourceAttr("trocco_user.test", "role", "member"),
					resource.TestCheckResourceAttr("trocco_user.test", "can_use_audit_log", "true"),
					resource.TestCheckResourceAttr("trocco_user.test", "is_restricted_connection_modify", "true"),
				),
			},
		},
	})
}

func TestAccUserResourceInvalidEmail(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_user" "test" {
					  email = "test"
					  password = "3XRambMkp-Hw"
					  role = "admin"
					}
				`,
				ExpectError: regexp.MustCompile(`invalid email address`),
			},
		},
	})
}

func TestAccUserResourceInvalidPassword(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_user" "test" {
					  email = "test@example.com"
					  password = "abc123"
					  role = "admin"
					}
				`,
				ExpectError: regexp.MustCompile(`password string length`),
			},
			{
				Config: providerConfig + `
					resource "trocco_user" "test" {
					  email = "test@example.com"
					  password = "1111111111111"
					  role = "admin"
					}
				`,
				ExpectError: regexp.MustCompile(`must contain at least one letter`),
			},
			{
				Config: providerConfig + `
					resource "trocco_user" "test" {
					  email = "test@example.com"
					  password = "aaaaaaaaaaaaa"
					  role = "admin"
					}
				`,
				ExpectError: regexp.MustCompile(`must contain at least one number`),
			},
			{
				Config: providerConfig + `
					resource "trocco_user" "test" {
					  email = "test@example.com"
					  role = "admin"
					}
				`,
				ExpectError: regexp.MustCompile(`Missing Required Attribute`),
			},
		},
	})
}

func TestAccUserResourceInvalidRole(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_user" "test" {
					  email = "test@example.com"
					  password = "abc123"
					  role = "invalid role"
					}
				`,
				ExpectError: regexp.MustCompile(`role value must be one of`),
			},
		},
	})
}
