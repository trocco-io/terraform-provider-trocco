package provider

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"

	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/client/parameter"
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
	t.Run("sqlserver", func(t *testing.T) {
		testAccConnectionResourceSQLServer(t)
	})
	t.Run("sqlserver_default_ssl", func(t *testing.T) {
		testAccConnectionResourceSQLServerDefaultSSL(t)
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
	t.Run("kintone_client_certificate", func(t *testing.T) {
		testAccConnectionResourceKintoneClientCertificate(t)
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
	t.Run("oracle_update", func(t *testing.T) {
		testAccConnectionResourceOracleUpdate(t)
	})
	t.Run("oracle_gateway_tns", func(t *testing.T) {
		testAccConnectionResourceOracleGatewayTns(t)
	})
	t.Run("oracle_privatelink", func(t *testing.T) {
		testAccConnectionResourceOraclePrivatelink(t)
	})
	t.Run("oracle_user_name_required", func(t *testing.T) {
		testAccConnectionResourceOracleUserNameRequired(t)
	})
	t.Run("oracle_driver_required", func(t *testing.T) {
		testAccConnectionResourceOracleDriverRequired(t)
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

func testAccConnectionResourceSQLServer(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.sqlserver_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/sqlserver_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "sqlserver"),
					resource.TestCheckResourceAttr(resourceName, "name", "sqlserver test"),
					resource.TestCheckResourceAttr(resourceName, "driver", "ms_sqlserver_jdbc_driver_13_4"),
					resource.TestCheckResourceAttr(resourceName, "ssl_enabled", "true"),
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
					return fmt.Sprintf("sqlserver,%s", connectionID), nil
				},
			},
		},
	})
}

// ssl_enabled is omitted, so it must default to false without leaving a diff after apply or refresh.
func testAccConnectionResourceSQLServerDefaultSSL(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.sqlserver_default_ssl_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/sqlserver_default_ssl.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ssl_enabled", "false"),
				),
			},
			{
				RefreshState: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ssl_enabled", "false"),
				),
			},
		},
	})
}

func TestInvalidSQLServerConnection(t *testing.T) {
	testCases := []struct {
		name        string
		configFile  string
		expectError string
	}{
		{
			name:        "missing_driver",
			configFile:  "testdata/connection/missing_driver_sqlserver.tf",
			expectError: `driver\s+is\s+required\s+for\s+SQL\s+Server\s+connection`,
		},
		{
			name:        "ssl_block",
			configFile:  "testdata/connection/ssl_block_sqlserver.tf",
			expectError: `ssl\s+is\s+not\s+supported\s+for\s+SQL\s+Server\s+connection`,
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
					// client_certificate_enabled is planned as null for non-kintone connection types.
					resource.TestCheckNoResourceAttr(resourceName, "client_certificate_enabled"),
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
					resource.TestCheckResourceAttr(resourceName, "client_certificate_enabled", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccConnectionResourceKintoneClientCertificate(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.kintone_client_certificate_test"

	// Self-signed test PFX files (base64) injected into the configurations so the same strings
	// can be sent through the API client to simulate a change outside Terraform.
	pfxWithPassword := strings.TrimSpace(LoadTextFile("testdata/connection/kintone_client_certificate_with_password.pfx.b64"))
	pfxWithoutPassword := strings.TrimSpace(LoadTextFile("testdata/connection/kintone_client_certificate_without_password.pfx.b64"))
	createConfig := fmt.Sprintf(LoadTextFile("testdata/connection/kintone_client_certificate_create.tf"), pfxWithPassword)
	updateDescriptionConfig := fmt.Sprintf(LoadTextFile("testdata/connection/kintone_client_certificate_update_description.tf"), pfxWithPassword)
	passwordDroppedConfig := fmt.Sprintf(LoadTextFile("testdata/connection/kintone_client_certificate_password_dropped.tf"), pfxWithPassword)
	updateConfig := fmt.Sprintf(LoadTextFile("testdata/connection/kintone_client_certificate_update.tf"), pfxWithoutPassword)
	removeConfig := LoadTextFile("testdata/connection/kintone_client_certificate_remove.tf")

	var connectionID int64
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + createConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "connection_type", "kintone"),
					resource.TestCheckResourceAttr(resourceName, "domain", "test_domain.s.cybozu.com"),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "client_certificate"),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_password", "test_pfx_password"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				// The API never returns the secrets, so they are not populated on import.
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password", "client_certificate", "client_certificate_password"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					return fmt.Sprintf("kintone,%s", s.RootModule().Resources[resourceName].Primary.ID), nil
				},
			},
			{
				Config: providerConfig + updateDescriptionConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", "This is a Kintone connection example (updated)"),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "client_certificate"),
					resource.TestCheckResourceAttr(resourceName, "client_certificate_password", "test_pfx_password"),
				),
			},
			{
				// The API rejects the re-sent certificate without its password; the state stays as in the previous step.
				Config:      providerConfig + passwordDroppedConfig,
				ExpectError: regexp.MustCompile(`Unable to update connection`),
			},
			{
				Config: providerConfig + updateConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "client_certificate_enabled", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "client_certificate"),
					resource.TestCheckNoResourceAttr(resourceName, "client_certificate_password"),
				),
			},
			{
				Config: providerConfig + removeConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "client_certificate_enabled", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "client_certificate"),
					resource.TestCheckNoResourceAttr(resourceName, "client_certificate_password"),
					func(s *terraform.State) error {
						id, err := strconv.ParseInt(s.RootModule().Resources[resourceName].Primary.ID, 10, 64)
						if err != nil {
							return err
						}
						connectionID = id
						return nil
					},
				),
			},
			{
				// A certificate enabled outside Terraform is read back as client_certificate_enabled = true,
				// so the plan shows true -> false while the configuration has no client_certificate.
				PreConfig: func() {
					enabled := true
					_, err := testAccTroccoClient().UpdateConnection("kintone", connectionID, &client.UpdateConnectionInput{
						ClientCertificateEnabled:  &enabled,
						ClientCertificate:         &parameter.NullableString{Value: pfxWithPassword, Valid: true},
						ClientCertificatePassword: &parameter.NullableString{Value: "test_pfx_password", Valid: true},
					})
					if err != nil {
						t.Fatalf("failed to enable the client certificate through the api: %v", err)
					}
				},
				Config:             providerConfig + removeConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
				// A PlanOnly step skips the pre-apply plan, so PreApply checks never run. The change
				// made in PreConfig is only visible after a refresh, hence PostApplyPostRefresh.
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(resourceName, plancheck.ResourceActionUpdate),
						plancheck.ExpectKnownValue(resourceName, tfjsonpath.New("client_certificate_enabled"), knownvalue.Bool(false)),
					},
				},
			},
			{
				Config: providerConfig + removeConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "client_certificate_enabled", "false"),
				),
			},
		},
	})
}

// testAccTroccoClient returns an API client for the environment under test, for changes made outside Terraform.
func testAccTroccoClient() *client.TroccoClient {
	return client.NewDevTroccoClient(os.Getenv("TROCCO_API_KEY"), os.Getenv("TROCCO_TEST_URL"))
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
		{
			name:        "mismatch_driver_oracle",
			configFile:  "testdata/connection/mismatch_driver_oracle.tf",
			expectError: "are: 12c-ojdbc7.jar, 19c-ojdbc8.jar, 26ai-ojdbc8.jar",
		},
		{
			name:        "mismatch_driver_sqlserver",
			configFile:  "testdata/connection/mismatch_driver_sqlserver.tf",
			expectError: "`mysql_connector_java_5_1_49` is invalid for SQL Server connection",
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

func TestInvalidKintoneClientCertificate(t *testing.T) {
	testCases := []struct {
		name        string
		configFile  string
		expectError string
	}{
		{
			name:        "password_without_cert",
			configFile:  "testdata/connection/kintone_client_certificate_password_without_cert.tf",
			expectError: `client_certificate is required for Kintone connection when\s+client_certificate_password is set`,
		},
		{
			name:        "enabled_is_read_only",
			configFile:  "testdata/connection/kintone_client_certificate_enabled_without_cert.tf",
			expectError: "read-only",
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
					resource.TestCheckResourceAttr(resourceName, "has_wallet_file", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// ssl_ca is write-only, same as password.
				ImportStateVerifyIgnore: []string{"password", "ssl_ca"},
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
				// Specifying both host/port and tns_admin_ora must fail
				// at plan time in ValidateConfig.
				Config:      providerConfig + LoadTextFile("testdata/connection/oracle/conflict_host_tns.tf"),
				ExpectError: regexp.MustCompile(`cannot be specified at the same time`),
			},
			{
				// A host-only (no port) combined with tns_admin_ora must also
				// fail at plan time: the conflict check must not require both
				// host and port to be set before detecting the conflict.
				Config:      providerConfig + LoadTextFile("testdata/connection/oracle/conflict_host_only_tns.tf"),
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
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// gateway is not restored on import: Read() falls back to state.Gateway
				// wholesale, so immediately after import (no prior state) the entire
				// gateway block is lost, not just its secret fields. This is a known,
				// pre-existing limitation shared by every connector that uses gateway
				// (not specific to Oracle), so only the gateway attributes are ignored
				// here; all other fields are still verified to round-trip correctly.
				ImportStateVerifyIgnore: []string{
					"password",
					"gateway",
					"gateway.host",
					"gateway.port",
					"gateway.user_name",
					"gateway.password",
					"gateway.key",
					"gateway.key_passphrase",
				},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					connectionID := s.RootModule().Resources[resourceName].Primary.ID
					return fmt.Sprintf("oracle,%s", connectionID), nil
				},
			},
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
			{
				// Updates wallet_file to a new value; has_wallet_file should stay true.
				Config: providerConfig + LoadTextFile("testdata/connection/oracle/wallet_update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "has_wallet_file", "true"),
				),
			},
			{
				// Removes wallet_file from config entirely; the API should clear it
				// server-side (has_wallet_file becomes false), not merely leave the
				// old value in place.
				Config: providerConfig + LoadTextFile("testdata/connection/oracle/wallet_delete.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "has_wallet_file", "false"),
				),
			},
		},
	})
}

func testAccConnectionResourceOracleUpdate(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.oracle_update_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/connection/oracle/update_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "host", "oracle-before-update.example.com"),
					resource.TestCheckResourceAttr(resourceName, "driver", "19c-ojdbc8.jar"),
					resource.TestCheckNoResourceAttr(resourceName, "tns_admin_ora"),
				),
			},
			{
				// Switches from host/port to TNS naming, and changes driver to a
				// non-default value, in the same update.
				Config: providerConfig + LoadTextFile("testdata/connection/oracle/update_to_tns.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "driver", "26ai-ojdbc8.jar"),
					resource.TestCheckResourceAttrSet(resourceName, "tns_admin_ora"),
					resource.TestCheckNoResourceAttr(resourceName, "host"),
					resource.TestCheckNoResourceAttr(resourceName, "port"),
				),
			},
		},
	})
}

func testAccConnectionResourceOracleGatewayTns(t *testing.T) {
	t.Helper()
	resourceName := "trocco_connection.oracle_gateway_tns_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// gateway combined with tns_admin_ora: needed for existing
				// AWS PrivateLink-based Autonomous Database connections (see
				// Design Decisions in the PR description).
				Config: providerConfig + LoadTextFile("testdata/connection/oracle/gateway_tns_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "tns_admin_ora"),
					resource.TestCheckResourceAttr(resourceName, "gateway.host", "bastion.example.com"),
					resource.TestCheckResourceAttr(resourceName, "gateway.port", "22"),
					resource.TestCheckResourceAttr(resourceName, "gateway.user_name", "ec2-user"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccConnectionResourceOraclePrivatelink(t *testing.T) {
	t.Helper()
	// SSH tunnels can only be created by an admin and must belong to the same
	// account as the API key, so the ID differs per environment.
	sshTunnelID := os.Getenv("TROCCO_TEST_SSH_TUNNEL_ID")
	if sshTunnelID == "" {
		t.Skip("TROCCO_TEST_SSH_TUNNEL_ID is not set")
	}
	if _, err := strconv.ParseInt(sshTunnelID, 10, 64); err != nil {
		t.Fatalf("TROCCO_TEST_SSH_TUNNEL_ID must be an integer: %s", sshTunnelID)
	}
	resourceName := "trocco_connection.oracle_privatelink_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + fmt.Sprintf(LoadTextFile("testdata/connection/oracle/privatelink_create.tf"), sshTunnelID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "aws_privatelink_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "ssh_tunnel_id", sshTunnelID),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
		},
	})
}

func testAccConnectionResourceOracleUserNameRequired(t *testing.T) {
	t.Helper()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Enforced by the provider's own ValidateConfig (plan-time), not just
				// the API's server-side presence validation.
				Config:      providerConfig + LoadTextFile("testdata/connection/oracle/missing_user_name.tf"),
				ExpectError: regexp.MustCompile(`user_name is required for Oracle connection`),
			},
		},
	})
}

func testAccConnectionResourceOracleDriverRequired(t *testing.T) {
	t.Helper()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Unlike MySQL/Snowflake, the Oracle backend fills in a default
				// driver when none is sent, which would otherwise produce an
				// inconsistent result after apply. driver must be required
				// at plan time for Oracle, same as PostgreSQL.
				Config:      providerConfig + LoadTextFile("testdata/connection/oracle/missing_driver.tf"),
				ExpectError: regexp.MustCompile(`driver is required for Oracle connection`),
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
