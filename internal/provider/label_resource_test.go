package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLabelResource(t *testing.T) {
	resourceName := "trocco_label.test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + LoadTextFile("testdata/label/basic_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Test Label"),
					resource.TestCheckResourceAttr(resourceName, "description", "This is a test label"),
					resource.TestCheckResourceAttr(resourceName, "color", "#FFFFFF"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr("trocco_label.test_omitted_description", "description", ""),
					resource.TestCheckResourceAttr("trocco_label.test_empty_description", "description", ""),
				),
			},
			// ImportState testing
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"description"},
			},
			// Update and Read testing
			{
				Config: providerConfig + LoadTextFile("testdata/label/basic_update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Updated Label"),
					resource.TestCheckResourceAttr(resourceName, "description", "This is an updated test label"),
					resource.TestCheckResourceAttr(resourceName, "color", "#000000"),
				),
			},
			{
				Config: providerConfig + `
					resource "trocco_label" "test_empty_description" {
					    name = "Updated Label, Second Time"
					    color = "#000000"
					}
				`,
				Check: resource.TestCheckResourceAttr("trocco_label.test_empty_description", "description", ""),
			},
			{
				Config: providerConfig + `
					resource "trocco_label" "test_empty_description_test2" {
					    name = "Updated Label, Third Time"
					    color = "#000000"
					    description = ""
					}
				`,
				Check: resource.TestCheckResourceAttr("trocco_label.test_empty_description_test2", "description", ""),
			},
		},
	})
}

func TestAccLabelResourceInvalidColor(t *testing.T) {
	testCases := []struct {
		name        string
		configFile  string
		expectError string
	}{
		{
			name:        "invalid_color",
			configFile:  "testdata/label/invalid_color.tf",
			expectError: "must be in format #RRGGBB or #RGB",
		},
		{
			name:        "empty_color",
			configFile:  "testdata/label/empty_color.tf",
			expectError: "must be in format #RRGGBB or #RGB",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config:      providerConfig + LoadTextFile(tc.configFile),
						ExpectError: regexp.MustCompile(tc.expectError),
					},
				},
			})
		})
	}
}

func TestAccLabelResourceInvalidName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/label/invalid_name.tf"),
				ExpectError: regexp.MustCompile(`UTF-8 character count must be at least 1, got: 0`),
			},
		},
	})
}
