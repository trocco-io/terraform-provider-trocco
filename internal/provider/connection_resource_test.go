package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccConnectionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_connection" "test" {
					  connection_type = "bigquery"
					  name = "test"
					  description = "The quick brown fox jumps over the lazy dog."
					  project_id = "test"

					  service_account_json_key = "{\"type\":\"service_account\",\"project_id\":\"\",\"private_key_id\":\"\",\"private_key\":\"\"}"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_connection.test", "connection_type", "bigquery"),
					resource.TestCheckResourceAttr("trocco_connection.test", "name", "test"),
					resource.TestCheckResourceAttr("trocco_connection.test", "description", "The quick brown fox jumps over the lazy dog."),
					resource.TestCheckResourceAttr("trocco_connection.test", "service_account_json_key", "{\"type\":\"service_account\",\"project_id\":\"\",\"private_key_id\":\"\",\"private_key\":\"\"}"),
					resource.TestCheckResourceAttrSet("trocco_connection.test", "id"),
				),
			},
			{
				ResourceName:            "trocco_connection.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_account_json_key"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					connectionID := s.RootModule().Resources["trocco_connection.test"].Primary.ID

					return fmt.Sprintf("bigquery,%s", connectionID), nil
				},
			},
			// Snowflake
			{
				Config: providerConfig + `
					resource "trocco_connection" "snowflake_test" {
					  connection_type = "snowflake"
					  auth_method = "user_password"

					  name = "snowflake test"
					  host = "example.snowflakecomputing.com"
					  user_name = "root"
					  password = "password"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_connection.snowflake_test", "connection_type", "snowflake"),
					resource.TestCheckResourceAttr("trocco_connection.snowflake_test", "name", "snowflake test"),
					resource.TestCheckResourceAttrSet("trocco_connection.snowflake_test", "id"),
				),
			},
			// MySQL
			{
				Config: providerConfig + `
					resource "trocco_connection" "mysql_test" {
					  connection_type = "mysql"

					  name = "mysql test"
					  host = "localhost"
					  user_name = "root"
					  password = "password"
					  port = 3306
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_connection.mysql_test", "connection_type", "mysql"),
					resource.TestCheckResourceAttr("trocco_connection.mysql_test", "name", "mysql test"),
					resource.TestCheckResourceAttrSet("trocco_connection.mysql_test", "id"),
				),
			},
			// PostgreSQL
			{
				Config: providerConfig + `
					resource "trocco_connection" "postgresql_test" {
					  connection_type = "postgresql"
					  name = "postgresql test"
					  host = "localhost"
					  user_name = "root"
					  password = "password"
					  port = 5432
					  driver = "postgresql_42_5_1"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_connection.postgresql_test", "connection_type", "postgresql"),
					resource.TestCheckResourceAttrSet("trocco_connection.postgresql_test", "id"),
				),
			},
			// Google Analytics4
			{
				Config: providerConfig + `
					resource "trocco_connection" "google_analytics4_test" {
					  connection_type = "google_analytics4"
					  name = "test"
					  description = "test"
						service_account_json_key = "{\"type\":\"service_account\",\"project_id\":\"\",\"private_key_id\":\"\",\"private_key\":\"\"}"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_connection.google_analytics4_test", "connection_type", "google_analytics4"),
					resource.TestCheckResourceAttr("trocco_connection.google_analytics4_test", "name", "test"),
					resource.TestCheckResourceAttr("trocco_connection.google_analytics4_test", "description", "test"),
					resource.TestCheckResourceAttr("trocco_connection.google_analytics4_test", "service_account_json_key", "{\"type\":\"service_account\",\"project_id\":\"\",\"private_key_id\":\"\",\"private_key\":\"\"}"),
					resource.TestCheckResourceAttrSet("trocco_connection.google_analytics4_test", "id"),
				),
			},
		},
	})
}

func TestInvalidDriver(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_connection" "invalid_driver_test" {
					  connection_type = "postgresql"
					  name = "invalid driver test"
					  host = "localhost"
					  user_name = "root"
					  password = "password"
					  port = 5432
					  driver = "invalid_driver"
					}
				`,
				ExpectError: regexp.MustCompile("driver: `invalid_driver` is invalid for PostgreSQL connection. "),
			},
			{
				Config: providerConfig + `
					resource "trocco_connection" "mismatch_driver_test_postgresql" {
					  connection_type = "postgresql"
					  name = "invalid driver test"
					  host = "localhost"
					  user_name = "root"
					  password = "password"
					  port = 5432
					  driver = "mysql_connector_java_5_1_49"
					}
				`,
				ExpectError: regexp.MustCompile("are: postgresql_42_5_1, postgresql_9_4_1205_jdbc41"),
			},
			{
				Config: providerConfig + `
					resource "trocco_connection" "mismatch_driver_test_mysql" {
					  connection_type = "mysql"
					  name = "invalid driver test"
					  host = "localhost"
					  user_name = "root"
					  password = "password"
					  port = 3306
					  driver = "snowflake_jdbc_3_14_2"
					}
				`,
				ExpectError: regexp.MustCompile("are: mysql_connector_java_5_1_49"),
			},
			{
				Config: providerConfig + `
					resource "trocco_connection" "mismatch_driver_test_snowflake" {
					  connection_type = "snowflake"
					  name = "invalid driver test"

					  auth_method = "user_password"
					  host = "example.snowflakecomputing.com"
					  user_name = "root"
					  password = "password"
					  driver = "mysql_connector_java_5_1_49"
					}
				`,
				ExpectError: regexp.MustCompile("are: snowflake_jdbc_3_14_2, snowflake_jdbc_3_17_0"),
			},
		},
	})
}
