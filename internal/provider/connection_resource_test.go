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
	t.Run("google_drive", func(t *testing.T) {
		testAccConnectionResourceGoogleDrive(t)
	})
	t.Run("kintone", func(t *testing.T) {
		testAccConnectionResourceKintone(t)
	})
	t.Run("marketo", func(t *testing.T) {
		testAccConnectionResourceMarketo(t)
	})
	// START [GENERATOR:CONNECTION_RESOURCE_TEST]
	t.Run("pagerduty", func(t *testing.T) {
		testAccConnectionResourcePagerduty(t)
	})
	t.Run("oracle", func(t *testing.T) {
		testAccConnectionResourceOracle(t)
	})
	t.Run("oracle_tns", func(t *testing.T) {
		testAccConnectionResourceOracleTns(t)
	})
	t.Run("oracle_host_tns_conflict", func(t *testing.T) {
		testAccConnectionResourceOracleHostTnsConflict(t)
	})
	t.Run("oracle_gateway", func(t *testing.T) {
		testAccConnectionResourceOracleGateway(t)
	})
	t.Run("oracle_wallet", func(t *testing.T) {
		testAccConnectionResourceOracleWallet(t)
	})
	// END [GENERATOR:CONNECTION_RESOURCE_TEST]
	t.Run("custom_connector", func(t *testing.T) {
		testAccConnectionResourceCustomConnector(t)
	})
	t.Run("custom_connector_oauth2", func(t *testing.T) {
		testAccConnectionResourceCustomConnectorOAuth2(t)
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

func testAccConnectionResourceGoogleDrive(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.google_drive_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/google_drive_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "google_drive"),
					resource.TestCheckResourceAttr(resourceName, "name", "google drive test"),
					resource.TestCheckResourceAttr(resourceName, "description", "test"),
					resource.TestCheckResourceAttr(resourceName, "service_account_json_key",
						"{\"type\":\"service_account\",\"project_id\":\"create_project_id\",\"private_key_id\":\"create_private_key_id\",\"private_key\":\"create_private_key\",\"client_email\":\"create_client_email\",\"client_id\":\"create_client_id\"}"),
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
					return fmt.Sprintf("google_drive,%s", connectionID), nil
				},
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

func TestInvalidCustomConnectorMissingID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/connection/invalid_custom_connector_missing_id.tf"),
				ExpectError: regexp.MustCompile("custom_connector_id is required for Custom Connector connection."),
			},
		},
	})
}

func TestInvalidCustomConnectorIDZero(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/connection/invalid_custom_connector_id_zero.tf"),
				ExpectError: regexp.MustCompile("value must be at least 1"),
			},
		},
	})
}

func testAccConnectionResourceMarketo(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.marketo"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/marketo/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "marketo"),
					resource.TestCheckResourceAttr(resourceName, "name", "Test Marketo Connection"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test Marketo connection for API operations"),
					resource.TestCheckResourceAttr(resourceName, "account_id", "123-ABC-456"),
					resource.TestCheckResourceAttr(resourceName, "client_id", "client_test_123"),
					resource.TestCheckResourceAttr(resourceName, "api_max_call_count", "5000"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_secret"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					connectionId := s.RootModule().Resources[resourceName].Primary.ID
					return fmt.Sprintf("marketo,%s", connectionId), nil
				},
			},
		},
	})
}

// testAccConnectionResourcePagerduty is auto-generated by tool. DO NOT EDIT.
func testAccConnectionResourcePagerduty(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.pagerduty_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/pagerduty/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "pagerduty"),
					resource.TestCheckResourceAttr(resourceName, "name", "Test Pagerduty Connection"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"api_key"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					connectionID := s.RootModule().Resources[resourceName].Primary.ID
					return fmt.Sprintf("pagerduty,%s", connectionID), nil
				},
			},
		},
	})
}

func testAccConnectionResourceOracle(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.oracle_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/oracle/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "oracle"),
					resource.TestCheckResourceAttr(resourceName, "name", "Test Oracle Connection"),
					resource.TestCheckResourceAttr(resourceName, "host", "oracle.example.com"),
					resource.TestCheckResourceAttr(resourceName, "port", "1521"),
					resource.TestCheckResourceAttr(resourceName, "user_name", "test_user"),
					resource.TestCheckResourceAttr(resourceName, "ssl_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// ssl_enabled is not restored on import: Read() only re-fetches it from
				// the API for the redshift connector and otherwise falls back to the
				// existing state, which is empty right after import. This is a known,
				// provider-wide limitation unrelated to Oracle (see ssl_enabled/gateway
				// import-restoration notes).
				ImportStateVerifyIgnore: []string{"password", "ssl_enabled"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					connectionID := s.RootModule().Resources[resourceName].Primary.ID
					return fmt.Sprintf("oracle,%s", connectionID), nil
				},
			},
		},
	})
}

func testAccConnectionResourceOracleTns(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.oracle_tns_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/oracle/tns_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "oracle"),
					resource.TestCheckResourceAttr(resourceName, "name", "Test Oracle TNS Connection"),
					resource.TestCheckResourceAttrSet(resourceName, "tns_admin_ora"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					connectionID := s.RootModule().Resources[resourceName].Primary.ID
					return fmt.Sprintf("oracle,%s", connectionID), nil
				},
			},
		},
	})
}

func testAccConnectionResourceOracleHostTnsConflict(t *testing.T) {
	t.Helper()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// host/port と tns_admin_ora を両方指定した場合は
				// ValidateConfig でプラン時にエラーになることを確認する。
				Config:      providerConfig + LoadTextFile("testdata/connection/oracle/conflict_host_tns.tf"),
				ExpectError: regexp.MustCompile(`cannot be specified at the same time`),
			},
		},
	})
}

func testAccConnectionResourceOracleGateway(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.oracle_gateway_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/oracle/gateway_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "oracle"),
					resource.TestCheckResourceAttr(resourceName, "gateway.host", "bastion.example.com"),
					resource.TestCheckResourceAttr(resourceName, "gateway.port", "22"),
					resource.TestCheckResourceAttr(resourceName, "gateway.user_name", "ec2-user"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// NOTE: import は検証しない。gateway は Read() が state.Gateway に
			// フォールバックする実装のため、import 直後（事前 state が無い状態）では
			// gateway ブロック全体が復元されない。これは gateway を使う全コネクタに
			// 共通する既存の実装上の制約であり、Oracle 固有の問題ではないため
			// 本 PR のスコープでは対応しない。
		},
	})
}

func testAccConnectionResourceOracleWallet(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.oracle_wallet_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/oracle/wallet_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "oracle"),
					resource.TestCheckResourceAttr(resourceName, "has_wallet_file", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password", "wallet_file"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					connectionID := s.RootModule().Resources[resourceName].Primary.ID
					return fmt.Sprintf("oracle,%s", connectionID), nil
				},
			},
		},
	})
}

func testAccConnectionResourceCustomConnector(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.custom_connector_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/custom_connector/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "custom_connector"),
					resource.TestCheckResourceAttr(resourceName, "name", "Test Custom Connector Connection"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test custom connector connection for acceptance testing"),
					resource.TestCheckResourceAttr(resourceName, "auth_type", "api_key"),
					resource.TestCheckNoResourceAttr(resourceName, "authorized"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "custom_connector_id", "trocco_custom_connector_input.test", "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"api_key"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					connectionID := s.RootModule().Resources[resourceName].Primary.ID
					return fmt.Sprintf("custom_connector,%s", connectionID), nil
				},
			},
			{
				Config: providerConfig + LoadTextFile("testdata/connection/custom_connector/update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
					resource.TestCheckResourceAttr(resourceName, "auth_type", "api_key"),
				),
			},
		},
	})
}

// testAccConnectionResourceCustomConnectorOAuth2 covers the oauth2 side of a
// custom connector connection, where the server serializes `scopes` and
// `authorized` (both are absent for an api_key definition). The first step
// deliberately omits `scopes`: the server answers with an empty list, which
// only round-trips because `scopes` is Optional + Computed.
func testAccConnectionResourceCustomConnectorOAuth2(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.custom_connector_oauth2_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/custom_connector/oauth2_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "custom_connector"),
					resource.TestCheckResourceAttr(resourceName, "auth_type", "oauth2"),
					resource.TestCheckResourceAttr(resourceName, "scopes.#", "0"),
					// A connection created through the API is never authorized,
					// regardless of grant_type.
					resource.TestCheckResourceAttr(resourceName, "authorized", "false"),
					resource.TestCheckResourceAttrPair(resourceName, "custom_connector_id", "trocco_custom_connector_input.oauth2_test", "id"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"oauth2_client_secret"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					connectionID := s.RootModule().Resources[resourceName].Primary.ID
					return fmt.Sprintf("custom_connector,%s", connectionID), nil
				},
			},
			{
				Config: providerConfig + LoadTextFile("testdata/connection/custom_connector/oauth2_update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "Updated description"),
					resource.TestCheckResourceAttr(resourceName, "scopes.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "scopes.0", "read"),
					resource.TestCheckResourceAttr(resourceName, "scopes.1", "write"),
					// The secret changed in this step, so `authorized` is planned
					// as unknown and re-derived; the server keeps it false.
					resource.TestCheckResourceAttr(resourceName, "authorized", "false"),
				),
			},
		},
	})
}
