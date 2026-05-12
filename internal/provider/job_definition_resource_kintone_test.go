package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

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

func TestAccJobDefinitionResourceMysqlToKintone(t *testing.T) {
	resourceName := "trocco_job_definition.mysql_to_kintone"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_kintone/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "MySQL to Kintone Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from MySQL to Kintone"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "kintone"),
					// MySQL input option attributes
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.table", "test_table"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.connect_timeout", "300"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.socket_timeout", "1800"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.default_time_zone", "Asia/Tokyo"),
					// Kintone output option attributes
					resource.TestCheckResourceAttrSet(resourceName, "output_option.kintone_output_option.kintone_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.app_id", "123"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.guest_space_id", "1"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.mode", "upsert"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.update_key", "id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.ignore_nulls", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.reduce_key", "email"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.chunk_size", "150"),
					// Kintone output option column options
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.#", "4"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.0.field_code", "record_id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.0.type", "NUMBER"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.1.name", "created_date"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.1.field_code", "created_at"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.1.type", "DATE"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.1.timezone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.2.name", "updated_time"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.2.field_code", "updated_at"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.2.type", "TIME"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.2.timezone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.3.name", "sub_items"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.3.field_code", "items_table"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.3.type", "SUBTABLE"),
					resource.TestCheckResourceAttr(resourceName, "output_option.kintone_output_option.kintone_output_option_column_options.3.sort_column", "item_order"),
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

// TestAccJobDefinitionResourceSftpToBigQuery tests SFTP input option.
