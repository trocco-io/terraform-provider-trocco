package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccNotificationDestination_Email(t *testing.T) {
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
					resource.TestCheckNoResourceAttr("trocco_notification_destination.email_test", "slack_channel_config"),
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
		},
	})
}
