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
				Config:      providerConfig + LoadTextFile("../../examples/testdata/bigquery_datamart_definition/create.tf"),
				ExpectError: nil,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "test_bigquery_datamart"),
					resource.TestCheckResourceAttr(resourceName, "query", "    SELECT * FROM examples\n"),
				),
			},
		},
	})
}
