package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceGcsToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.gcs_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/gcs_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "GCS to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from GCS to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "custom_spec"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "gcs"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.gcs_input_option.bucket", "example_bucket"),
					resource.TestCheckResourceAttr(resourceName, "input_option.gcs_input_option.path_prefix", "path/to/your/csv_file"),
					resource.TestCheckResourceAttr(resourceName, "input_option.gcs_input_option.stop_when_file_not_found", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.gcs_input_option.incremental_loading_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "gcs_to_bigquery_test_table"),
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

func TestAccJobDefinitionResourceMysqlToGcsCSV(t *testing.T) {
	resourceName := "trocco_job_definition.mysql_to_gcs"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_gcs/create_csv.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "MySQL to GCS Test (csv)"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from MySQL to GCS"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "gcs"),
					// MySQL input option attributes
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.table", "test_table"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.connect_timeout", "300"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.socket_timeout", "1800"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.default_time_zone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.incremental_loading_enabled", "false"),
					// GCS output option attributes
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.bucket", "my-test-bucket"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.path_prefix", "output/test/"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.file_ext", ".csv"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.sequence_format", ".%03d.%02d"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.is_minimum_output_tasks", "false"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.formatter_type", "csv"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.encoder_type", "gzip"),
					// CSV formatter attributes
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.delimiter", ","),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.newline", "CRLF"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.newline_in_field", "LF"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.charset", "UTF-8"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.escape", "\\"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.quote_policy", "MINIMAL"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.header_line", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.null_string_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.default_time_zone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.csv_formatter_column_options_attributes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.csv_formatter_column_options_attributes.0.name", "created_at"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.csv_formatter_column_options_attributes.0.format", "%Y-%m-%d %H:%M:%S"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.csv_formatter.csv_formatter_column_options_attributes.0.timezone", "Asia/Tokyo"),
					// Filter columns
					resource.TestCheckResourceAttr(resourceName, "filter_columns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.name", "type"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.src", "type"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.type", "string"),
				),
			},
			// Import testing
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

func TestAccJobDefinitionResourceMysqlToGcsJSONL(t *testing.T) {
	resourceName := "trocco_job_definition.mysql_to_gcs"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_gcs/create_jsonl.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "MySQL to GCS Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from MySQL to GCS"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "gcs"),
					// MySQL input option attributes
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.table", "test_table"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.connect_timeout", "300"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.socket_timeout", "1800"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.default_time_zone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.incremental_loading_enabled", "false"),
					// GCS output option attributes
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.bucket", "my-test-bucket"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.path_prefix", "output/test/"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.file_ext", ".jsonl"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.sequence_format", ".%03d.%02d"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.is_minimum_output_tasks", "false"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.formatter_type", "jsonl"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.encoder_type", "gzip"),
					// JSONL formatter attributes
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.jsonl_formatter.encoding", "UTF-8"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.jsonl_formatter.newline", "LF"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.jsonl_formatter.date_format", "%Y-%m-%d %H:%M:%S"),
					resource.TestCheckResourceAttr(resourceName, "output_option.gcs_output_option.jsonl_formatter.timezone", "UTC"),
					// Filter columns
					resource.TestCheckResourceAttr(resourceName, "filter_columns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.name", "type"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.src", "type"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.type", "string"),
				),
			},
			// Import testing
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
