package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccNotificationDestinationResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
					resource "trocco_notification_destination" "email_test" {
						type = "email"
						email_config = {
							email = "test@example.com"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_notification_destination.email_test", "type", "email"),
					resource.TestCheckResourceAttr("trocco_notification_destination.email_test", "email_config.email", "test@example.com"),
					resource.TestCheckResourceAttrSet("trocco_notification_destination.email_test", "id"),
				),
			},
			{
				ResourceName:      "trocco_notification_destination.email_test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					id := s.RootModule().Resources["trocco_notification_destination.email_test"].Primary.ID
					return fmt.Sprintf("email,%s", id), nil
				},
			},
			// Slack_Channel
			{
				Config: providerConfig + `
					resource "trocco_notification_destination" "slack_channel_test" {
						type = "slack_channel"
						slack_channel_config = {
							channel = "trocco-log2"
							webhook_url = "https://hooks.slack.com/services/test"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_notification_destination.slack_channel_test", "type", "slack_channel"),
					resource.TestCheckResourceAttr("trocco_notification_destination.slack_channel_test", "slack_channel_config.channel", "trocco-log2"),
					resource.TestCheckResourceAttr("trocco_notification_destination.slack_channel_test", "slack_channel_config.webhook_url", "https://hooks.slack.com/services/test"),
					resource.TestCheckResourceAttrSet("trocco_notification_destination.slack_channel_test", "id"),
				),
			},
		},
	})
}

func TestInvalidNotificationDestinationType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Invalid type
			{
				Config: providerConfig + `
					resource "trocco_notification_destination" "invalid_type_test" {
					  type = "invalid_type"
					  email_config = {
						email = "test@example.com"
					  }
					}
				`,
				ExpectError: regexp.MustCompile(`"type" must be either "email" or "slack_channel".`),
			},
			// Valid type but missing email_config for email type
			{
				Config: providerConfig + `
					resource "trocco_notification_destination" "missing_email_config" {
					  type = "email"
					}
				`,
				ExpectError: regexp.MustCompile("`email_config.email` is required when type is 'email'."),
			},
			// Valid type but missing slack_channel_config for slack_channel type
			{
				Config: providerConfig + `
					resource "trocco_notification_destination" "missing_slack_channel_config" {
					  type = "slack_channel"
					}
				`,
				ExpectError: regexp.MustCompile("`slack_channel_config` is required when type is 'slack_channel'"),
			},
		},
	})
}
