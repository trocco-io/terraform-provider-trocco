package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceNotifications(t *testing.T) {
	resourceName := "trocco_job_definition.notifications_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/notifications/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "notifications_test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition with notifications"),
					// Slack message
					resource.TestCheckResourceAttr(resourceName, "notifications.0.message", "This is a multi-line message\nwith several lines\n  and some indentation\n    to test TrimmedStringType\n"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "slack"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.notification_type", "job"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.notify_when", "finished"),
					resource.TestCheckResourceAttrSet(resourceName, "notifications.0.id"),
					// Email message
					resource.TestCheckResourceAttr(resourceName, "notifications.1.message", "  This is another multi-line message\nwith leading and trailing whitespace\n  \n  to test TrimmedStringType\n  \n"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "email"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.notification_type", "job"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.notify_when", "finished"),
					resource.TestCheckResourceAttrSet(resourceName, "notifications.1.id"),
				),
			},
			// Import testing
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					// The message attributes are trimmed and set in state, so different from the resource config.
					"notifications.0.message",
					"notifications.1.message",
				},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources[resourceName].Primary.ID
					return jobDefinitionId, nil
				},
			},
			// Reordering the notifications in config should be reflected in state
			// without leaving a perpetual diff after refresh.
			{
				Config: providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/notifications/reorder.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "email"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "slack"),
				),
			},
			// Multiple notifications sharing the same (type, destination_type)
			// must keep their distinct destination IDs without being swapped.
			{
				Config: providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/notifications/multiple_slack.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
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
