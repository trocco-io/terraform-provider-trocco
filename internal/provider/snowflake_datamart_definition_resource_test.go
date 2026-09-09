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
			// Import testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// The query attribute is trimmed and set in state, so different from the resource config.
					"query",
				},
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
			// Import testing: SCD Type 2 specific attributes must be restored on import.
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// The query attribute is trimmed and set in state, so different from the resource config.
					"query",
				},
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

func TestAccSnowflakeDatamartDefinitionResourceNotifications(t *testing.T) {
	resourceName := "trocco_snowflake_datamart_definition.test_snowflake_datamart_notifications"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/notifications/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_snowflake_datamart_notifications"),
					// Verify that the slack notification message is correctly set
					resource.TestCheckResourceAttr(resourceName, "notifications.0.message", "This is a multi-line message\nwith several lines\n  and some indentation\n    to test TrimmedStringType\n"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "slack"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.notification_type", "job"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.notify_when", "finished"),
					resource.TestCheckResourceAttrSet(resourceName, "notifications.0.id"),
					// Verify that the email notification message is correctly set
					resource.TestCheckResourceAttr(resourceName, "notifications.1.message", "  This is another multi-line message\nwith leading and trailing whitespace\n\n  to test TrimmedStringType\n\n"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "email"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.notification_type", "record"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.record_count", "100"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.record_operator", "above"),
					resource.TestCheckResourceAttrSet(resourceName, "notifications.1.id"),
				),
			},
			// Reordering the notifications in config should be reflected in state
			// without leaving a perpetual diff after refresh.
			{
				Config: providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/notifications/reorder.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "email"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "slack"),
				),
			},
			// Multiple notifications sharing the same (type, destination_type)
			// must keep their distinct destination IDs without being swapped.
			{
				Config: providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/notifications/multiple_slack.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "notifications.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "slack"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.notification_type", "job"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.notify_when", "finished"),
					resource.TestCheckResourceAttrPair(resourceName, "notifications.0.slack_channel_id", "trocco_notification_destination.slack", "id"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "slack"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.notification_type", "job"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.notify_when", "failed"),
					resource.TestCheckResourceAttrPair(resourceName, "notifications.1.slack_channel_id", "trocco_notification_destination.slack_b", "id"),
				),
			},
		},
	})
}

func TestAccSnowflakeDatamartDefinitionResourceWriteDispositionTransition(t *testing.T) {
	resourceName := "trocco_snowflake_datamart_definition.test_snowflake_datamart_transition"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/transition/append.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "write_disposition", "append"),
					resource.TestCheckNoResourceAttr(resourceName, "schema_evolution_mode"),
					resource.TestCheckResourceAttr(resourceName, "notifications.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "slack"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "email"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "3"),
				),
			},
			{
				Config: providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/transition/replace.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "write_disposition", "replace"),
				),
			},
			{
				Config: providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/transition/incremental_skip.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "write_disposition", "incremental"),
					resource.TestCheckResourceAttr(resourceName, "merge_keys.0", "id"),
					resource.TestCheckResourceAttr(resourceName, "on_matched_action", "skip"),
					resource.TestCheckResourceAttr(resourceName, "schema_evolution_mode", "auto_add_column"),
					resource.TestCheckNoResourceAttr(resourceName, "valid_from_column"),
					// Notifications must keep their configured order across updates.
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "slack"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "email"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "3"),
				),
			},
			{
				Config: providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/transition/scd_type_2.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "write_disposition", "scd_type_2"),
					resource.TestCheckResourceAttr(resourceName, "merge_keys.0", "id"),
					resource.TestCheckResourceAttr(resourceName, "incremental_column", "updated_at"),
					resource.TestCheckResourceAttr(resourceName, "schema_evolution_mode", "auto_add_column"),
					resource.TestCheckResourceAttr(resourceName, "valid_from_column", "trocco_valid_from"),
					resource.TestCheckResourceAttr(resourceName, "valid_to_column", "trocco_valid_to"),
					resource.TestCheckResourceAttr(resourceName, "is_current_column", "trocco_is_current"),
					resource.TestCheckNoResourceAttr(resourceName, "on_matched_action"),
				),
			},
			// Switching back to truncate must clear every incremental / SCD Type 2
			// attribute without leaving an inconsistent state.
			{
				Config: providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/transition/truncate.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "write_disposition", "truncate"),
					resource.TestCheckNoResourceAttr(resourceName, "merge_keys.#"),
					resource.TestCheckNoResourceAttr(resourceName, "on_matched_action"),
					resource.TestCheckNoResourceAttr(resourceName, "incremental_column"),
					resource.TestCheckNoResourceAttr(resourceName, "schema_evolution_mode"),
					resource.TestCheckNoResourceAttr(resourceName, "valid_from_column"),
					resource.TestCheckNoResourceAttr(resourceName, "valid_to_column"),
					resource.TestCheckNoResourceAttr(resourceName, "is_current_column"),
					// Notifications and schedules must survive the transitions unchanged.
					resource.TestCheckResourceAttr(resourceName, "notifications.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "slack"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.notification_type", "job"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "email"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.notification_type", "record"),
					resource.TestCheckResourceAttr(resourceName, "schedules.#", "3"),
				),
			},
		},
	})
}

func TestAccSnowflakeDatamartDefinitionResourceQualityChecks(t *testing.T) {
	resourceName := "trocco_snowflake_datamart_definition.test_snowflake_datamart_quality_checks"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/quality_checks/create.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "quality_check_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "quality_check_on_violation", "fail"),
					resource.TestCheckResourceAttr(resourceName, "quality_check_lookback_period_column", "updated_at"),
					resource.TestCheckResourceAttr(resourceName, "quality_check_lookback_period_column_type", "TIMESTAMP_NTZ"),
					resource.TestCheckResourceAttr(resourceName, "quality_check_lookback_period_timezone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "quality_check_lookback_period_from", "3"),
					resource.TestCheckResourceAttr(resourceName, "quality_check_lookback_period_to", "0"),
					resource.TestCheckResourceAttr(resourceName, "quality_check_lookback_period_unit", "days"),
					resource.TestCheckResourceAttr(resourceName, "quality_checks.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "quality_checks.0.check_type", "not_null"),
					resource.TestCheckResourceAttr(resourceName, "quality_checks.0.column_names.0", "id"),
					resource.TestCheckResourceAttr(resourceName, "quality_checks.1.check_type", "composite_unique"),
					resource.TestCheckResourceAttr(resourceName, "quality_checks.1.column_names.#", "2"),
				),
			},
			// Disabling quality checks must clear every quality check attribute.
			{
				Config: providerConfig + LoadTextFile("testdata/snowflake_datamart_definition/quality_checks/disable.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "quality_check_enabled", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "quality_check_on_violation"),
					resource.TestCheckNoResourceAttr(resourceName, "quality_check_lookback_period_column"),
					resource.TestCheckNoResourceAttr(resourceName, "quality_checks.#"),
				),
			},
		},
	})
}
