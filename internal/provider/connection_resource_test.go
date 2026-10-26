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
