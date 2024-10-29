package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccUserResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
					resource "trocco_user" "test" {
					  email = "test@example.com"
					  password = "3XRambMkp-Hw"
					  role = "admin"
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("trocco_user.test", "email", "test@example.com"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "trocco_user.test",
				ImportState:       true,
				ImportStateVerify: true,
				//ImportStateVerifyIgnore: []string{"last_updated"},
			},
			// Update and Read testing
		},
	})
}
