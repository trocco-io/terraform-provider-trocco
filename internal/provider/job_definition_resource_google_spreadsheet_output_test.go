package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceGoogleSpreadsheetToGoogleSpreadsheet(t *testing.T) {
	resourceName := "trocco_job_definition.sheets_to_sheets"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/google_spreadsheet_to_google_spreadsheet/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Google Spreadsheets to Google Spreadsheets"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from Google Spreadsheets to Google Spreadsheets"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "0"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "google_spreadsheets"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "google_spreadsheets"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_spreadsheets_input_option.worksheet_title", "input-data"),

					resource.TestCheckResourceAttr(resourceName, "output_option.google_spreadsheets_output_option.spreadsheets_id", "TEST_SHEETS_ID"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_spreadsheets_output_option.worksheet_title", "output-data"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_spreadsheets_output_option.timezone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_spreadsheets_output_option.value_input_option", "USER_ENTERED"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_spreadsheets_output_option.mode", "replace"),
					// google_spreadsheets_output_option_sorts
					resource.TestCheckResourceAttr(resourceName, "output_option.google_spreadsheets_output_option.google_spreadsheets_output_option_sorts.0.column", "created_at"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_spreadsheets_output_option.google_spreadsheets_output_option_sorts.0.order", "ascending"),
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
