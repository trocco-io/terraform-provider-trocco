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
