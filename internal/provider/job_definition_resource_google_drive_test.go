package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceGoogleDriveToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.google_drive_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/google_drive_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Google Drive to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from Google Drive to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "0"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "google_drive"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttrSet(resourceName, "input_option.google_drive_input_option.google_drive_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_drive_input_option.folder_id", "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_drive_input_option.file_match_pattern", ""),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_drive_input_option.is_skip_header_line", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_drive_input_option.stop_when_file_not_found", "false"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "google_drive_to_bigquery_test_table"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "append"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.location", "US"),
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

func TestAccJobDefinitionResourceBigQueryToGoogleDrive(t *testing.T) {
	resourceName := "trocco_job_definition.bigquery_to_google_drive"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/bigquery_to_google_drive/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "BigQuery to Google Drive CSV Export"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "google_drive"),

					// Check Google Drive output option fields
					resource.TestCheckResourceAttrSet(resourceName, "output_option.google_drive_output_option.google_drive_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_drive_output_option.main_folder_id", "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_drive_output_option.file_name", "users_export.csv"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_drive_output_option.formatter_type", "csv"),

					// Check CSV formatter
					resource.TestCheckResourceAttr(resourceName, "output_option.google_drive_output_option.csv_formatter.delimiter", ","),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_drive_output_option.csv_formatter.newline", "CRLF"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_drive_output_option.csv_formatter.charset", "UTF-8"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_drive_output_option.csv_formatter.header_line", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_drive_output_option.csv_formatter.null_string_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_drive_output_option.csv_formatter.null_string", "NULL"),

					// Check CSV column options
					resource.TestCheckResourceAttr(resourceName, "output_option.google_drive_output_option.csv_formatter.csv_formatter_column_options_attributes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "output_option.google_drive_output_option.csv_formatter.csv_formatter_column_options_attributes.0.name", "created_at"),
				),
			},
			// Step 2: Import state
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources[resourceName].Primary.ID
					return jobDefinitionId, nil
				},
			},
		},
	})
}
