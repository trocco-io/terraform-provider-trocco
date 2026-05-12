package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceHubspotToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.hubspot_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/hubspot_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Hubspot to BigQuery Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from Hubspot to BigQuery"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "true"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "hubspot"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.hubspot_input_option.hubspot_connection_id", "388"),
					resource.TestCheckResourceAttr(resourceName, "input_option.hubspot_input_option.target", "object"),
					resource.TestCheckResourceAttr(resourceName, "input_option.hubspot_input_option.object_type", "contact"),
					resource.TestCheckResourceAttr(resourceName, "input_option.hubspot_input_option.incremental_loading_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.hubspot_input_option.input_option_columns.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "input_option.hubspot_input_option.input_option_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.hubspot_input_option.input_option_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "input_option.hubspot_input_option.input_option_columns.1.name", "email"),
					resource.TestCheckResourceAttr(resourceName, "input_option.hubspot_input_option.input_option_columns.1.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "input_option.hubspot_input_option.input_option_columns.2.name", "created_at"),
					resource.TestCheckResourceAttr(resourceName, "input_option.hubspot_input_option.input_option_columns.2.type", "timestamp"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.name", "contact_id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.src", "id"),
					resource.TestCheckResourceAttr(resourceName, "filter_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "test_dataset"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "hubspot_to_bigquery_test_table"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "append"),
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

func TestAccJobDefinitionResourceMysqlToHubspot(t *testing.T) {
	resourceName := "trocco_job_definition.mysql_to_hubspot"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/mysql_connection.tf") + LoadTextFile("testdata/job_definition/mysql_to_hubspot/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "MySQL to HubSpot Test"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test job definition for transferring data from MySQL to HubSpot"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "2"),
					resource.TestCheckResourceAttr(resourceName, "is_runnable_concurrently", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "mysql"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "hubspot"),

					// Check MySQL input option
					resource.TestCheckResourceAttrSet(resourceName, "input_option.mysql_input_option.mysql_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.table", "test_table"),
					resource.TestCheckResourceAttr(resourceName, "input_option.mysql_input_option.incremental_loading_enabled", "false"),

					// Check HubSpot output option
					resource.TestCheckResourceAttrSet(resourceName, "output_option.hubspot_output_option.hubspot_connection_id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.hubspot_output_option.object_type", "task"),
					resource.TestCheckResourceAttr(resourceName, "output_option.hubspot_output_option.mode", "merge"),
					resource.TestCheckResourceAttr(resourceName, "output_option.hubspot_output_option.upsert_key", "id"),
					resource.TestCheckResourceAttr(resourceName, "output_option.hubspot_output_option.number_of_parallels", "2"),

					// Check associations
					resource.TestCheckResourceAttr(resourceName, "output_option.hubspot_output_option.associations.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "output_option.hubspot_output_option.associations.0.to_object_type", "contact"),
					resource.TestCheckResourceAttr(resourceName, "output_option.hubspot_output_option.associations.0.from_object_key", "contact_email"),
					resource.TestCheckResourceAttr(resourceName, "output_option.hubspot_output_option.associations.0.to_object_key", "email"),
					resource.TestCheckResourceAttr(resourceName, "output_option.hubspot_output_option.associations.1.to_object_type", "deal"),
					resource.TestCheckResourceAttr(resourceName, "output_option.hubspot_output_option.associations.1.from_object_key", "deal_task"),
					resource.TestCheckResourceAttr(resourceName, "output_option.hubspot_output_option.associations.1.to_object_key", "task"),
				),
			},
			// ImportState testing
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
