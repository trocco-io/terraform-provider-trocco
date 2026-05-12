package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceS3ToSnowflake(t *testing.T) {
	resourceName := "trocco_job_definition.s3_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/s3_to_snowflake/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "s3 to snowflake"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "s3"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "snowflake"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "custom_spec"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources["trocco_job_definition.s3_test"].Primary.ID
					return jobDefinitionId, nil
				},
			},
		},
	})
}

func TestAccJobDefinitionResourceS3ToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.s3_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/s3_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "S3 to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from S3 to BigQuery with filter_columns"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "custom_spec"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "s3"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.s3_input_option.bucket", "test_bucket"),
					resource.TestCheckResourceAttr(resourceName, "input_option.s3_input_option.path_prefix", "data/users.csv"),
					resource.TestCheckResourceAttr(resourceName, "input_option.s3_input_option.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "user_id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.name", "user_name"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.src", "name"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.1.default", "Unknown"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.3.name", "registration_date"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.3.format", "%Y-%m-%d"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "s3_to_bigquery_test_table"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "append"),
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

func TestAccJobDefinitionResourceMysqlToS3CSV(t *testing.T) {
	resourceName := "trocco_job_definition.mysql_to_s3_csv"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_s3/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "MySQL to S3 CSV Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from MySQL to S3 with CSV formatter"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "s3"),
					// MySQL input option attributes
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.table", "test_table"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.connect_timeout", "300"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.socket_timeout", "1800"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.default_time_zone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.fetch_rows", "1000"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.incremental_loading_enabled", "false"),
					// S3 output option attributes
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.bucket", "test-bucket"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.path_prefix", "output/data"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.region", "ap-northeast-1"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.file_ext", "csv.gz"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.sequence_format", ".%03d"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.canned_acl", "Private"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.is_minimum_output_tasks", "false"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.multipart_upload_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.formatter_type", "csv"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.encoder_type", "gzip"),
					// CSV formatter attributes
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.delimiter", ","),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.escape", "\\"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.header_line", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.charset", "UTF-8"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.quote_policy", "ALL"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.newline", "LF"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.newline_in_field", "LF"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.null_string_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.null_string", "NULL"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.default_time_zone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.csv_formatter_column_options_attributes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.csv_formatter_column_options_attributes.0.name", "created_at"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.csv_formatter_column_options_attributes.0.format", "%Y-%m-%d %H:%M:%S"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.csv_formatter.csv_formatter_column_options_attributes.0.timezone", "Asia/Tokyo"),
					// Filter columns
					resource.TestCheckResourceAttr(resourceName, "filter_columns.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
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

func TestAccJobDefinitionResourceMysqlToS3JSONL(t *testing.T) {
	resourceName := "trocco_job_definition.mysql_to_s3_jsonl"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_s3/create_jsonl.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "MySQL to S3 JSONL Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from MySQL to S3 with JSONL formatter"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "s3"),
					// MySQL input option attributes
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.table", "test_table"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.connect_timeout", "300"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.socket_timeout", "1800"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.default_time_zone", "Asia/Tokyo"),
					// S3 output option attributes
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.bucket", "test-bucket"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.path_prefix", "output/json"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.region", "us-west-2"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.file_ext", "jsonl.gz"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.sequence_format", ".%03d"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.is_minimum_output_tasks", "false"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.multipart_upload_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.formatter_type", "jsonl"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.encoder_type", "gzip"),
					// JSONL formatter attributes
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.jsonl_formatter.encoding", "UTF-8"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.jsonl_formatter.newline", "LF"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.jsonl_formatter.date_format", "%Y-%m-%d"),
					resource.TestCheckResourceAttr(resourceName, "output_option.s3_output_option.jsonl_formatter.timezone", "UTC"),
					// Filter columns
					resource.TestCheckResourceAttr(resourceName, "filter_columns.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
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
