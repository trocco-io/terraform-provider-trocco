package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourcePagerdutyToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.pagerduty_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/pagerduty_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "pagerduty"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.pagerduty_input_option.path", "escalation_policies"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources[resourceName].Primary.ID
					return jobDefinitionId, nil
				},
			},
		},
	})
}
