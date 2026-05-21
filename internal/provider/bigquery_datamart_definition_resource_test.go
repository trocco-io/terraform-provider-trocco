package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDatamartDefinitionResourceForBigquery(t *testing.T) {
	resourceName := "trocco_bigquery_datamart_definition.test_bigquery_datamart"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/bigquery_datamart_definition/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_bigquery_datamart"),
					resource.TestCheckResourceAttr(resourceName, "query", "    SELECT * FROM examples\n"),
					resource.TestCheckResourceAttr(resourceName, "before_load", "    DELETE FROM examples\n    WHERE created_at < '2024-01-01'\n"),
				),
			},
		},
	})
}

func TestAccDatamartDefinitionResourceForBigqueryNotifications(t *testing.T) {
	resourceName := "trocco_bigquery_datamart_definition.test_bigquery_datamart_notifications"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/bigquery_datamart_definition/notifications/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_bigquery_datamart_notifications"),
					resource.TestCheckResourceAttr(resourceName, "query", "    SELECT * FROM examples\n"),
					// Verify that the slack notification message is correctly set
					resource.TestCheckResourceAttr(resourceName, "notifications.0.message", "This is a multi-line message\nwith several lines\n  and some indentation\n    to test TrimmedStringType\n"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "slack"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.notification_type", "job"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.notify_when", "finished"),
					// Verify that the email notification message is correctly set
					resource.TestCheckResourceAttr(resourceName, "notifications.1.message", "  This is another multi-line message\nwith leading and trailing whitespace\n  \n  to test TrimmedStringType\n  \n"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "email"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.notification_type", "record"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.record_count", "100"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.record_operator", "above"),
				),
			},
			// Reordering the notifications in config should be reflected in state
			// without leaving a perpetual diff after refresh.
			{
				Config: providerConfig + LoadTextFile("testdata/bigquery_datamart_definition/notifications/reorder.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "email"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "slack"),
				),
			},
			// Multiple notifications sharing the same (type, destination_type)
			// must keep their distinct destination IDs without being swapped.
			{
				Config: providerConfig + LoadTextFile("testdata/bigquery_datamart_definition/notifications/multiple_slack.tf"),
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

func TestAccDatamartDefinitionResourceForBigqueryIncremental(t *testing.T) {
	resourceName := "trocco_bigquery_datamart_definition.test_incremental"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/bigquery_datamart_definition/incremental.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_incremental"),
					resource.TestCheckResourceAttr(resourceName, "write_disposition", "incremental"),
					resource.TestCheckResourceAttr(resourceName, "merge_keys.0", "id"),
					resource.TestCheckResourceAttr(resourceName, "on_matched_action", "upsert"),
					resource.TestCheckResourceAttr(resourceName, "schema_evolution_mode", "detect_only"),
					resource.TestCheckResourceAttr(resourceName, "lookback_period_column", "updated_at"),
					resource.TestCheckResourceAttr(resourceName, "lookback_period_column_type", "TIMESTAMP"),
					resource.TestCheckResourceAttr(resourceName, "lookback_period_timezone", "Asia/Tokyo"),
					resource.TestCheckResourceAttr(resourceName, "lookback_period_from", "3"),
					resource.TestCheckResourceAttr(resourceName, "lookback_period_to", "0"),
					resource.TestCheckResourceAttr(resourceName, "lookback_period_unit", "days"),
				),
			},
		},
	})
}

func TestAccDatamartDefinitionResourceForBigquerySCDType2(t *testing.T) {
	resourceName := "trocco_bigquery_datamart_definition.test_scd_type_2"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      providerConfig + LoadTextFile("testdata/bigquery_datamart_definition/scd_type_2.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_scd_type_2"),
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

func TestAccDatamartDefinitionResourceForBigqueryTruncateWriteDisposition(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Test case: write_disposition is "truncate" and before_load is set (should fail)
				Config:      providerConfig + LoadTextFile("testdata/bigquery_datamart_definition/truncate_with_before_load.tf"),
				ExpectError: regexp.MustCompile("before_load is only available in insert query mode"),
			},
			{
				// Test case: write_disposition is "truncate" and before_load is not set (should pass)
				Config:      providerConfig + LoadTextFile("testdata/bigquery_datamart_definition/truncate_without_before_load.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_bigquery_datamart_definition.test_truncate_without_before_load", "name", "test_truncate_without_before_load"),
					resource.TestCheckResourceAttr("trocco_bigquery_datamart_definition.test_truncate_without_before_load", "write_disposition", "truncate"),
				),
			},
		},
	})
}
