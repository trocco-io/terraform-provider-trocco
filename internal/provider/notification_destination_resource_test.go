package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccNotificationDestinationResource(t *testing.T) {
	t.Run("email", func(t *testing.T) {
		testAccNotificationDestinationResourceEmail(t)
	})
	t.Run("slack_channel", func(t *testing.T) {
		testAccNotificationDestinationResourceSlackChannel(t)
	})
	t.Run("http", func(t *testing.T) {
		testAccNotificationDestinationResourceHTTP(t)
	})
}

func testAccNotificationDestinationResourceHTTP(t *testing.T) {
	t.Helper()
	resourceName := "trocco_notification_destination.http_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/notification_destination/http_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "type", "http"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "http_config.name", "terraform-acc-http"),
					resource.TestCheckResourceAttr(resourceName, "http_config.url", "https://example.com/webhook"),
					resource.TestCheckResourceAttr(resourceName, "http_config.description", "created by acceptance test"),
					resource.TestCheckResourceAttr(resourceName, "http_config.headers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "http_config.headers.0.key", "Authorization"),
					resource.TestCheckResourceAttr(resourceName, "http_config.headers.0.value", "Bearer acc-secret"),
					resource.TestCheckResourceAttr(resourceName, "http_config.headers.0.masking", "true"),
					resource.TestCheckResourceAttr(resourceName, "http_config.headers.1.key", "Content-Type"),
					resource.TestCheckResourceAttr(resourceName, "http_config.headers.1.value", "application/json"),
					resource.TestCheckResourceAttr(resourceName, "http_config.headers.1.masking", "false"),
					resource.TestCheckResourceAttr(resourceName, "http_config.query_params.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "http_config.query_params.0.key", "token"),
					resource.TestCheckResourceAttr(resourceName, "http_config.query_params.0.value", "acc-token"),
					resource.TestCheckResourceAttr(resourceName, "http_config.query_params.0.masking", "true"),
				),
			},
			// Refresh: masked values are not returned by the API and must be restored from the state.
			{
				Config:   providerConfig + LoadTextFile("testdata/notification_destination/http_create.tf"),
				PlanOnly: true,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					id := s.RootModule().Resources[resourceName].Primary.ID
					return fmt.Sprintf("http,%s", id), nil
				},
				// masked values cannot be read back from the API
				ImportStateVerifyIgnore: []string{"http_config.headers.0.value", "http_config.query_params.0.value"},
			},
			{
				Config: providerConfig + LoadTextFile("testdata/notification_destination/http_update.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "http_config.name", "terraform-acc-http-updated"),
					resource.TestCheckResourceAttr(resourceName, "http_config.url", "https://example.com/webhook/v2"),
					resource.TestCheckNoResourceAttr(resourceName, "http_config.description"),
					resource.TestCheckResourceAttr(resourceName, "http_config.headers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "http_config.headers.0.key", "X-Api-Key"),
					resource.TestCheckResourceAttr(resourceName, "http_config.headers.0.value", "rotated-secret"),
					resource.TestCheckResourceAttr(resourceName, "http_config.headers.0.masking", "true"),
					resource.TestCheckNoResourceAttr(resourceName, "http_config.query_params.#"),
				),
			},
			{
				Config:   providerConfig + LoadTextFile("testdata/notification_destination/http_update.tf"),
				PlanOnly: true,
			},
		},
	})
}

func testAccNotificationDestinationResourceEmail(t *testing.T) {
	t.Helper()
	resourceName := "trocco_notification_destination.email_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/notification_destination/email_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "type", "email"),
					resource.TestCheckResourceAttr(resourceName, "email_config.email", "test@example.com"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					id := s.RootModule().Resources[resourceName].Primary.ID
					return fmt.Sprintf("email,%s", id), nil
				},
			},
		},
	})
}

func testAccNotificationDestinationResourceSlackChannel(t *testing.T) {
	t.Helper()
	resourceName := "trocco_notification_destination.slack_channel_test"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/notification_destination/slack_channel_create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "type", "slack_channel"),
					resource.TestCheckResourceAttr(resourceName, "slack_channel_config.channel", "trocco-log2"),
					resource.TestCheckResourceAttr(resourceName, "slack_channel_config.webhook_url", "https://hooks.slack.com/services/test"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
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
				ExpectError: regexp.MustCompile(`"type" must be one of "email", "slack_channel" or "http".`),
			},
			// Valid type but missing http_config for http type
			{
				Config: providerConfig + `
					resource "trocco_notification_destination" "http" {
					  type = "http"
					}
				`,
				ExpectError: regexp.MustCompile("`http_config` is required when type is 'http'."),
			},
			// Valid type but conflicting email_config for http type
			{
				Config: providerConfig + `
					resource "trocco_notification_destination" "http" {
						type = "http"
						http_config = {
							name = "conflict"
							url  = "https://example.com/webhook"
						}
						email_config = {
							email = "test@example.com"
						}
					}
				`,
				ExpectError: regexp.MustCompile(`Invalid Attribute Combination`),
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
