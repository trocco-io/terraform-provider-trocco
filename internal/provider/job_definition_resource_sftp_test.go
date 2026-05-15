package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceSftpToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.sftp_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with SFTP CSV parser
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/sftp_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test job_definition"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "sftp"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),

					// Check SFTP input option fields
					resource.TestCheckResourceAttrSet(resourceName, "input_option.sftp_input_option.sftp_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sftp_input_option.path_prefix", "/data/files/"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sftp_input_option.path_match_pattern", ".*\\.csv$"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sftp_input_option.incremental_loading_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sftp_input_option.stop_when_file_not_found", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sftp_input_option.decompression_type", "guess"),

					// Check CSV parser
					resource.TestCheckResourceAttr(resourceName, "input_option.sftp_input_option.csv_parser.delimiter", ","),
					resource.TestCheckResourceAttr(resourceName, "input_option.sftp_input_option.csv_parser.escape", "\\"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sftp_input_option.csv_parser.quote", "\""),
					resource.TestCheckResourceAttr(resourceName, "input_option.sftp_input_option.csv_parser.columns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sftp_input_option.csv_parser.columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.sftp_input_option.csv_parser.columns.0.type", "long"),
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

// TestAccJobDefinitionResourceBigQueryToSftpCSV tests SFTP output option with CSV formatter.

func TestAccJobDefinitionResourceBigQueryToSftpCSV(t *testing.T) {
	resourceName := "trocco_job_definition.bigquery_to_sftp_csv"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with CSV formatter
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/bigquery_to_sftp/create_csv.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "BigQuery to SFTP CSV Export"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "sftp"),

					// Check SFTP output option fields
					resource.TestCheckResourceAttrSet(resourceName, "output_option.sftp_output_option.sftp_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.path_prefix", "/exports/users/users_$export_date$"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.file_ext", ".csv"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.is_minimum_output_tasks", "false"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.encoder_type", "gzip"),

					// Check CSV formatter
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.csv_formatter.delimiter", ","),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.csv_formatter.newline", "CRLF"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.csv_formatter.charset", "UTF-8"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.csv_formatter.header_line", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.csv_formatter.null_string_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.csv_formatter.null_string", "NULL"),

					// Check CSV column options
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.csv_formatter.csv_formatter_column_options_attributes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.csv_formatter.csv_formatter_column_options_attributes.0.name", "created_at"),

					// Check custom variables
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.custom_variable_settings.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.custom_variable_settings.0.name", "$export_date$"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.custom_variable_settings.0.type", "timestamp"),
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

// TestAccJobDefinitionResourceBigQueryToSftpJSONL tests SFTP output option with JSONL formatter.

func TestAccJobDefinitionResourceBigQueryToSftpJSONL(t *testing.T) {
	resourceName := "trocco_job_definition.bigquery_to_sftp_jsonl"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Step 1: Create with JSONL formatter
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/bigquery_to_sftp/create_jsonl.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "BigQuery to SFTP JSONL Export"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "sftp"),

					// Check SFTP output option fields
					resource.TestCheckResourceAttrSet(resourceName, "output_option.sftp_output_option.sftp_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.path_prefix", "/analytics/events/$date$/events"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.file_ext", ".jsonl"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.is_minimum_output_tasks", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.encoder_type", "gzip"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.sequence_format", "%03d.%02d"),

					// Check JSONL formatter
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.jsonl_formatter.encoding", "UTF-8"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.jsonl_formatter.newline", "LF"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.jsonl_formatter.date_format", "%Y-%m-%d %H:%M:%S"),
					resource.TestCheckResourceAttr(resourceName, "output_option.sftp_output_option.jsonl_formatter.timezone", "UTC"),
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
