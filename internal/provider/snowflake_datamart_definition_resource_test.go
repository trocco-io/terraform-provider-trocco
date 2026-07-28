package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSnowflakeDatamartDefinitionResourceInsertMode(t *testing.T) {
	resourceName := "trocco_snowflake_datamart_definition.test_snowflake_datamart_insert"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/insert/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_snowflake_datamart_insert"),
					resource.TestCheckResourceAttr(resourceName, "query_mode", "insert"),
					resource.TestCheckResourceAttr(resourceName, "warehouse", "EXAMPLE_WH"),
					resource.TestCheckResourceAttr(resourceName, "destination_database", "DEST_DATABASE"),
					resource.TestCheckResourceAttr(resourceName, "destination_schema", "DEST_SCHEMA"),
					resource.TestCheckResourceAttr(resourceName, "destination_table", "DEST_TABLE"),
					resource.TestCheckResourceAttr(resourceName, "write_disposition", "truncate"),
				),
			},
		},
	})
}

func TestAccSnowflakeDatamartDefinitionResourceQueryMode(t *testing.T) {
	resourceName := "trocco_snowflake_datamart_definition.test_snowflake_datamart_query"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/query/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_snowflake_datamart_query"),
					resource.TestCheckResourceAttr(resourceName, "query_mode", "query"),
					resource.TestCheckResourceAttr(resourceName, "warehouse", "EXAMPLE_WH"),
					resource.TestCheckResourceAttr(resourceName, "statement_timeout", "3600"),
				),
			},
		},
	})
}

func TestAccSnowflakeDatamartDefinitionResourceMissingRequiredInsertFields(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/missing_required_insert_fields.tf"),
				ExpectError: regexp.MustCompile("destination_database is required for insert query mode"),
			},
		},
	})
}

func TestAccSnowflakeDatamartDefinitionResourceIncremental(t *testing.T) {
	resourceName := "trocco_snowflake_datamart_definition.test_snowflake_datamart_incremental"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/incremental.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_snowflake_datamart_incremental"),
					resource.TestCheckResourceAttr(resourceName, "write_disposition", "incremental"),
					resource.TestCheckResourceAttr(resourceName, "merge_keys.0", "id"),
					resource.TestCheckResourceAttr(resourceName, "on_matched_action", "upsert"),
					resource.TestCheckResourceAttr(resourceName, "schema_evolution_mode", "detect_only"),
					resource.TestCheckResourceAttr(resourceName, "lookback_period_column", "updated_at"),
					resource.TestCheckResourceAttr(resourceName, "lookback_period_column_type", "TIMESTAMP_NTZ"),
					resource.TestCheckResourceAttr(resourceName, "lookback_period_timezone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "lookback_period_from", "3"),
					resource.TestCheckResourceAttr(resourceName, "lookback_period_to", "0"),
					resource.TestCheckResourceAttr(resourceName, "lookback_period_unit", "days"),
				),
			},
		},
	})
}

func TestAccSnowflakeDatamartDefinitionResourceSCDType2(t *testing.T) {
	resourceName := "trocco_snowflake_datamart_definition.test_snowflake_datamart_scd_type_2"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/scd_type_2.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_snowflake_datamart_scd_type_2"),
					resource.TestCheckResourceAttr(resourceName, "write_disposition", "scd_type_2"),
					resource.TestCheckResourceAttr(resourceName, "merge_keys.0", "id"),
					resource.TestCheckResourceAttr(resourceName, "incremental_column", "updated_at"),
					resource.TestCheckResourceAttr(resourceName, "schema_evolution_mode", "detect_only"),
					resource.TestCheckResourceAttr(resourceName, "valid_from_column", "trocco_valid_from"),
					resource.TestCheckResourceAttr(resourceName, "valid_to_column", "trocco_valid_to"),
					resource.TestCheckResourceAttr(resourceName, "is_current_column", "trocco_is_current"),
				),
			},
		},
	})
}

func TestAccSnowflakeDatamartDefinitionResourceMissingIncrementalFields(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/missing_incremental_fields.tf"),
				ExpectError: regexp.MustCompile("merge_keys is required when write_disposition is incremental"),
			},
		},
	})
}

func TestAccSnowflakeDatamartDefinitionResourceMissingQualityCheckFields(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/missing_quality_check_fields.tf"),
				ExpectError: regexp.MustCompile("quality_check_on_violation is required when quality_check_enabled is true"),
			},
		},
	})
}
