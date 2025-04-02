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
					resource "trocco_notification_destination" "invalid_type" {
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
					resource "trocco_notification_destination" "email" {
					  type = "email"
					}
				`,
				ExpectError: regexp.MustCompile("`email_config.email` is required when type is 'email'."),
			},
			// missing email
			{
				Config: providerConfig + `
					resource "trocco_notification_destination" "email" {
						type = "email"
						email_config = {
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Incorrect attribute value type`),
			},
			// Valid type but conflicting slack_channel_config for email type
			{
				Config: providerConfig + `
					resource "trocco_notification_destination" "email" {
						type = "email"
						email_config = {
							email = "test@example.com"
						}
						slack_channel_config = {
							channel     = "trocco-log2"
							webhook_url = "https://hooks.slack.com/services/test"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Invalid Attribute Combination`),
			},
			// Valid type but missing slack_channel_config for slack_channel type
			{
				Config: providerConfig + `
					resource "trocco_notification_destination" "slack" {
					  type = "slack_channel"
					}
				`,
				ExpectError: regexp.MustCompile("`slack_channel_config` is required when type is 'slack_channel'."),
			},
			// Valid type but conflicting email_config for slack_channel type
			{
				Config: providerConfig + `
					resource "trocco_notification_destination" "slack" {
						type = "slack_channel"
						email_config = {
							email = "test@example.com"
						}
						slack_channel_config = {
							channel     = "trocco-log2"
							webhook_url = "https://hooks.slack.com/services/test"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Invalid Attribute Combination`),
			},
			// missing channel and webhook_url
			{
				Config: providerConfig + `
					resource "trocco_notification_destination" "slack" {
						type = "slack_channel"
						slack_channel_config = {
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Incorrect attribute value type`),
			},
		},
	})
}

func TestInvalidEmailValidation(t *testing.T) {
	invalidEmails := []struct {
		email         string
		expectedError string
	}{
		// Invalid email address: Missing dot in domain
		{"notify@examplecom", `invalid email address`},

		// Invalid email address: Domain part incomplete
		{"notify@.com", `invalid email address`},

		// Invalid email address: Missing username
		{"@example.com", `invalid email address`},

		// Invalid email address: Extra dot in domain
		{"notify@com.", `invalid email address`},

		// Invalid email address: Incomplete domain
		{"notify@example", `invalid email address`},

		// Invalid email address: Space in domain
		{"notify@exa mple.com", `invalid email address`},

		// Invalid email address: Multiple '@' symbols
		{"notify@com@domain.com", `invalid email address`},
	}
	for _, testCase := range invalidEmails {
		t.Run(testCase.email, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: providerConfig + `
							resource "trocco_notification_destination" "email" {
								type = "email"
								email_config = {
									email = "` + testCase.email + `"
								}
							}
						`,
						ExpectError: regexp.MustCompile(testCase.expectedError),
					},
				},
			})
		})
	}
}
