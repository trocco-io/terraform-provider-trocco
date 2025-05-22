package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceMysqlToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.mysql_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/mysql_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test job_definition"),
					resource.TestCheckResourceAttr(resourceName, "description", "test description"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "large"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "1"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources["trocco_job_definition.mysql_to_bigquery"].Primary.ID

					return jobDefinitionId, nil
				},
			},
		},
	})
}

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

func TestAccJobDefinitionResourceGoogleAnalytics4ToSnowflake(t *testing.T) {
	resourceName := "trocco_job_definition.ga4_to_snowflake"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/google_analytics4_to_snowflake/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "GA4 to Snowflake"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.time_series", "dateHour"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.start_date", "2daysAgo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.end_date", "1daysAgo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.google_analytics4_input_option_dimensions.0.name", "yyyymm"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.google_analytics4_input_option_dimensions.0.expression", "{\n  \"concatenate\": {\n    \"dimensionNames\": [\"year\",\"month\"],\n    \"delimiter\": \"-\"\n  }\n}\n"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.google_analytics4_input_option_metrics.0.name", "totalUsers"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.google_analytics4_input_option_metrics.0.expression", ""),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.incremental_loading_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.retry_limit", "5"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.retry_sleep", "2"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.raise_on_other_row", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.limit_of_rows", "10000"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.0.name", "date_hour"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.0.type", "timestamp"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.1.name", "yyyymm"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.1.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.2.name", "total_users"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.2.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.3.name", "property_id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.3.type", "string"),
				),
			},
			// ImportState testing
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
			// Update testing with null dimension
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/google_analytics4_to_snowflake/update_dimension_null.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "GA4 to Snowflake"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.time_series", "dateHour"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.start_date", "2daysAgo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.google_analytics4_input_option_dimensions.#", "0"),
				),
			},
		},
	})
}

func TestAccJobDefinitionResourceGoogleAnalytics4ToSnowflakeInvalid(t *testing.T) {
	resourceName := "trocco_job_definition.ga4_to_snowflake"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/google_analytics4_to_snowflake/update_dimension_too_many.tf"),
				ExpectError:  regexp.MustCompile(`list must contain at most 8 elements, got: 9`),
			},
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/google_analytics4_to_snowflake/update_metrics_required.tf"),
				ExpectError:  regexp.MustCompile(`"google_analytics4_input_option_metrics" is required.`),
			},
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/google_analytics4_to_snowflake/update_metrics_too_many.tf"),
				ExpectError:  regexp.MustCompile(`list must contain at least 1 elements and at most 10 elements, got: 11`),
			},
		},
	})
}

func TestAccJobDefinitionResourceKintoneToSnowflake(t *testing.T) {
	resourceName := "trocco_job_definition.kintone_to_snowflake"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/kintone_to_snowflake/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "kintone to snowflake"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "0"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "output_option.snowflake_output_option.batch_size", "50"),
					resource.TestCheckResourceAttr(resourceName, "output_option.snowflake_output_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.app_id", "123"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.expand_subtable", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "input_option.kintone_input_option.guest_space_id"),
					resource.TestCheckNoResourceAttr(resourceName, "input_option.kintone_input_option.query"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.0.name", "duration"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.1.name", "date"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.1.type", "timestamp"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.1.format", "%Y%m%d"),
				),
			},
			// ImportState testing
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

func TestAccJobDefinitionResourceKintoneToSnowflakeInvalid(t *testing.T) {
	resourceName := "trocco_job_definition.kintone_to_snowflake"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/kintone_to_snowflake/update_app_id_required.tf"),
				ExpectError:  regexp.MustCompile(`Missing Configuration for Required Attribute`),
			},
		},
	})
}

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

func TestAccJobDefinitionResourceGcsToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.gcs_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/gcs_to_bigquery/create.tf"),
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

func TestAccJobDefinitionResourcePostgresqlToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.postgresql_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/postgresql_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "PostgreSQL to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from PostgreSQL to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "3"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "postgresql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.postgresql_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.postgresql_input_option.schema", "public"),
					resource.TestCheckResourceAttr(resourceName, "input_option.postgresql_input_option.fetch_rows", "1000"),
					resource.TestCheckResourceAttr(resourceName, "input_option.postgresql_input_option.default_time_zone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "postgresql_to_bigquery_test_table"),
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

func TestAccJobDefinitionResourceS3ToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.s3_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/s3_to_bigquery/create.tf"),
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

func TestAccJobDefinitionResourceSalesforceToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.salesforce_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/salesforce_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Salesforce to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from Salesforce to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "salesforce"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.salesforce_input_option.object", "Contact"),
					resource.TestCheckResourceAttr(resourceName, "input_option.salesforce_input_option.object_acquisition_method", "soql"),
					resource.TestCheckResourceAttr(resourceName, "input_option.salesforce_input_option.soql", "SELECT Id, Name, Email, CreatedDate FROM Contact"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "contact_id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "Id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "salesforce_to_bigquery_test_table"),
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

func TestAccJobDefinitionResourceBigQueryToSnowflake(t *testing.T) {
	resourceName := "trocco_job_definition.bigquery_to_snowflake"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/bigquery_to_snowflake/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "BigQuery to Snowflake Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from BigQuery to Snowflake"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "snowflake"),
					resource.TestCheckResourceAttr(resourceName, "input_option.bigquery_input_option.query", "SELECT id, name, email, created_at FROM `test_dataset.test_table`"),
					resource.TestCheckResourceAttr(resourceName, "input_option.bigquery_input_option.location", "US"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "user_id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "output_option.snowflake_output_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "output_option.snowflake_output_option.schema", "PUBLIC"),
					resource.TestCheckResourceAttr(resourceName, "output_option.snowflake_output_option.table", "bigquery_to_snowflake_test_table"),
					resource.TestCheckResourceAttr(resourceName, "output_option.snowflake_output_option.mode", "insert"),
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

func TestAccJobDefinitionResourceSnowflakeToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.snowflake_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/snowflake_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Snowflake to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from Snowflake to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "snowflake"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.snowflake_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.snowflake_input_option.schema", "PUBLIC"),
					resource.TestCheckResourceAttr(resourceName, "input_option.snowflake_input_option.query", "SELECT id, name, email, created_at FROM test_table"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "user_id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "snowflake_to_bigquery_test_table"),
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

func TestAccJobDefinitionResourceNotifications(t *testing.T) {
	resourceName := "trocco_job_definition.notifications_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/job_definition/notifications/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "notifications_test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition with notifications"),
					// Emailの通知メッセージが正しく設定されていることを確認
					resource.TestCheckResourceAttr(resourceName, "notifications.0.message", "  This is another multi-line message\nwith leading and trailing whitespace\n  \n  to test TrimmedStringType\n  \n"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "email"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.notification_type", "exec_time"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.minutes", "30"),
					// Slackの通知メッセージが正しく設定されていることを確認
					resource.TestCheckResourceAttr(resourceName, "notifications.1.message", "This is a multi-line message\nwith several lines\n  and some indentation\n    to test TrimmedStringType\n"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "slack"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.notification_type", "job"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.notify_when", "finished"),
				),
			},
			// Import testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// The message attributes are trimmed and set in state, so different from the resource config.
					"notifications.0.message",
					"notifications.1.message",
				},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources[resourceName].Primary.ID
					return jobDefinitionId, nil
				},
			},
		},
	})
}
