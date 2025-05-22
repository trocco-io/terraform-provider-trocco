package provider

import (
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
					// Verify that the email notification message is correctly set
					resource.TestCheckResourceAttr(resourceName, "notifications.0.message", "  This is another multi-line message\nwith leading and trailing whitespace\n  \n  to test TrimmedStringType\n  \n"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.destination_type", "email"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.notification_type", "record"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.record_count", "100"),
					resource.TestCheckResourceAttr(resourceName, "notifications.0.record_operator", "above"),
					// Verify that the slack notification message is correctly set
					resource.TestCheckResourceAttr(resourceName, "notifications.1.message", "This is a multi-line message\nwith several lines\n  and some indentation\n    to test TrimmedStringType\n"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.destination_type", "slack"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.notification_type", "job"),
					resource.TestCheckResourceAttr(resourceName, "notifications.1.notify_when", "finished"),
				),
			},
		},
	})
}
