package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccConnectionResource(t *testing.T) {
	t.Run("bigquery", func(t *testing.T) {
		testAccConnectionResourceBigQuery(t)
	})
	t.Run("snowflake", func(t *testing.T) {
		testAccConnectionResourceSnowflake(t)
	})
	t.Run("mysql", func(t *testing.T) {
		testAccConnectionResourceMySQL(t)
	})
	t.Run("postgresql", func(t *testing.T) {
		testAccConnectionResourcePostgreSQL(t)
	})
	t.Run("google_analytics4", func(t *testing.T) {
		testAccConnectionResourceGoogleAnalytics4(t)
	})
	t.Run("kintone", func(t *testing.T) {
		testAccConnectionResourceKintone(t)
	})
}

func testAccConnectionResourceBigQuery(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/bigquery_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "name", "test"),
					resource.TestCheckResourceAttr(resourceName, "description", "The quick brown fox jumps over the lazy dog."),
					resource.TestCheckResourceAttr(resourceName, "service_account_json_key", "{\"type\":\"service_account\",\"project_id\":\"\",\"private_key_id\":\"\",\"private_key\":\"\"}"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"service_account_json_key"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					connectionID := s.RootModule().Resources[resourceName].Primary.ID
					return fmt.Sprintf("bigquery,%s", connectionID), nil
				},
			},
		},
	})
}

func testAccConnectionResourceSnowflake(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.snowflake_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/snowflake_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "snowflake"),
					resource.TestCheckResourceAttr(resourceName, "name", "snowflake test"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccConnectionResourceMySQL(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.mysql_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/mysql_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "name", "mysql test"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccConnectionResourcePostgreSQL(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.postgresql_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/postgresql_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "postgresql"),
					resource.TestCheckResourceAttr(resourceName, "name", "postgresql test"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccConnectionResourceGoogleAnalytics4(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.google_analytics4_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/google_analytics4_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "google_analytics4"),
					resource.TestCheckResourceAttr(resourceName, "name", "test"),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "service_account_json_key",
						"{\"type\":\"service_account\",\"project_id\":\"create_project_id\",\"private_key_id\":\"create_private_key_id\",\"private_key\":\"create_private_key\",\"client_email\":\"create_client_email\",\"client_id\":\"create_client_id\"}"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccConnectionResourceKintone(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.kintone_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/kintone_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "kintone"),
					resource.TestCheckResourceAttr(resourceName, "name", "Kintone Test"),
					resource.TestCheckResourceAttr(resourceName, "domain", "test_domain"),
					resource.TestCheckResourceAttr(resourceName, "login_method", "username_and_password"),
					resource.TestCheckResourceAttr(resourceName, "password", "test_password"),
					resource.TestCheckResourceAttr(resourceName, "username", "test_username"),
					resource.TestCheckResourceAttr(resourceName, "basic_auth_username", "test_basic_auth_username"),
					resource.TestCheckResourceAttr(resourceName, "basic_auth_password", "test_basic_auth_password"),
					resource.TestCheckNoResourceAttr(resourceName, "token"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func TestInvalidDriver(t *testing.T) {
	testCases := []struct {
		name        string
		configFile  string
		expectError string
	}{
		{
			name:        "invalid_driver",
			configFile:  "testdata/connection/invalid_driver.tf",
			expectError: "driver: `invalid_driver` is invalid for PostgreSQL connection. ",
		},
		{
			name:        "mismatch_driver_postgresql",
			configFile:  "testdata/connection/mismatch_driver_postgresql.tf",
			expectError: "are: postgresql_42_5_1, postgresql_9_4_1205_jdbc41",
		},
		{
			name:        "mismatch_driver_mysql",
			configFile:  "testdata/connection/mismatch_driver_mysql.tf",
			expectError: "are: mysql_connector_java_5_1_49",
		},
		{
			name:        "mismatch_driver_snowflake",
			configFile:  "testdata/connection/mismatch_driver_snowflake.tf",
			expectError: "are: snowflake_jdbc_3_14_2, snowflake_jdbc_3_17_0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config:      providerConfig + LoadTextFile(tc.configFile),
						ExpectError: regexp.MustCompile(tc.expectError),
					},
				},
			})
		})
	}
}

func TestInvalidLoginMethod(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/connection/invalid_login_method.tf"),
				ExpectError: regexp.MustCompile("login_method: `invalid_login_method` is invalid for Kintone connection."),
			},
		},
	})
}
