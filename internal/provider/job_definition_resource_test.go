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
				Config:       providerConfig + LoadTextFile("../../examples/testdata/job_definition/mysql_to_bigquery/create.tf"),
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
				Config:       providerConfig + LoadTextFile("../../examples/testdata/job_definition/s3_to_snowflake/create.tf"),
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
				Config:       providerConfig + LoadTextFile("../../examples/testdata/job_definition/google_analytics4_to_snowflake/create.tf"),
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
				Config:       providerConfig + LoadTextFile("../../examples/testdata/job_definition/google_analytics4_to_snowflake/update_dimension_null.tf"),
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
				Config:       providerConfig + LoadTextFile("../../examples/testdata/job_definition/google_analytics4_to_snowflake/update_dimension_too_many.tf"),
				ExpectError:  regexp.MustCompile(`list must contain at most 8 elements, got: 9`),
			},
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("../../examples/testdata/job_definition/google_analytics4_to_snowflake/update_metrics_required.tf"),
				ExpectError:  regexp.MustCompile(`"google_analytics4_input_option_metrics" is required.`),
			},
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("../../examples/testdata/job_definition/google_analytics4_to_snowflake/update_metrics_too_many.tf"),
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
				Config:       providerConfig + LoadTextFile("../../examples/testdata/job_definition/kintone_to_snowflake/create.tf"),
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
				Config:       providerConfig + LoadTextFile("../../examples/testdata/job_definition/kintone_to_snowflake/update_app_id_required.tf"),
				ExpectError:  regexp.MustCompile(`Missing Configuration for Required Attribute`),
			},
		},
	})
}
