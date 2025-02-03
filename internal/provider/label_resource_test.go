package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLabelResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
					resource "trocco_label" "test" {
						name = "Test Label"
						color = "#FFFFFF"
						description = "This is a test label"
					}

					resource "trocco_label" "test2" {
					    name = "Test Label 2"
					    color = "#FFFFFF"
					    description = "This is a test label"
					}

					resource "trocco_label" "test_omitted_description" {
					    name = "Test Label Using Omitted Description"
					    color = "#FFFFFF"
					}

					resource "trocco_label" "test_empty_description" {
					    name = "Test Label Using Empty Description"
					    color = "#FFFFFF"
					    description = ""
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_label.test", "name", "Test Label"),
					resource.TestCheckResourceAttr("trocco_label.test", "description", "This is a test label"),
					resource.TestCheckResourceAttr("trocco_label.test", "color", "#FFFFFF"),
					resource.TestCheckResourceAttrSet("trocco_label.test", "id"),

					resource.TestCheckResourceAttr("trocco_label.test_omitted_description", "description", ""),

					resource.TestCheckResourceAttr("trocco_label.test_empty_description", "description", ""),
				),
			},
			// ImportState testing
			{
				ResourceName:            "trocco_label.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"description"},
			},
			// Update and Read testing
			{
				Config: providerConfig + `
					resource "trocco_label" "test" {
						name = "Updated Label"
						description = "This is an updated test label"
						color = "#000000"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_label.test", "name", "Updated Label"),
					resource.TestCheckResourceAttr("trocco_label.test", "description", "This is an updated test label"),
					resource.TestCheckResourceAttr("trocco_label.test", "color", "#000000"),
				),
			},
			{
				Config: providerConfig + `
					resource "trocco_label" "test" {
					    name = "Updated Label, Second Time"
					    color = "#000000"
					}
				`,
				Check: resource.TestCheckResourceAttr("trocco_label.test_empty_description", "description", ""),
			},
			{
				Config: providerConfig + `
					resource "trocco_label" "test2" {
					    name = "Updated Label, Third Time"
					    color = "#000000"
					    description = ""
					}
				`,
				Check: resource.TestCheckResourceAttr("trocco_label.test_empty_description", "description", ""),
			},
		},
	})
}

func TestAccLabelResourceInvalidColor(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_label" "test" {
						name = "Test Label"
						description = "This is a test label"
						color = "invalid_color"
					}
				`,
				ExpectError: regexp.MustCompile(`must be in format #RRGGBB or #RGB`),
			},
			{
				Config: providerConfig + `
					resource "trocco_label" "test" {
					    name = "Test Label"
					    description = "This is a test label"
					    color = ""
					}
				`,
				ExpectError: regexp.MustCompile(`must be in format #RRGGBB or #RGB`),
			},
		},
	})
}

func TestAccLabelResourceInvalidName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_label" "test" {
						name = ""
						description = "This is a test label"
						color = "#FFFFFF"
					}
				`,
				ExpectError: regexp.MustCompile(`UTF-8 character count must be at least 1, got: 0`),
			},
		},
	})
}
