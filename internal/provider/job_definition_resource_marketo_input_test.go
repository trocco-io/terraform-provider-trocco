package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceMarketoToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.marketo_lead_with_date_filter"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/marketo_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Marketo to BigQuery - Lead with Date Filter"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "marketo"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.target", "lead"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.from_date", "2025-03-01"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.end_date", "2025-03-18"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.use_updated_at", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.polling_interval_second", "60"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.bulk_job_timeout_second", "3600"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.input_option_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.input_option_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.input_option_columns.1.name", "email"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.input_option_columns.1.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "marketing_data"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "marketo_leads"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "merge"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.location", "US"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.bigquery_output_option_merge_keys.0", "id"),
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

func TestAccJobDefinitionResourceMarketoActivityToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.marketo_activity_with_type_filter"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/marketo_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Marketo to BigQuery - Activity with Type Filter"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "marketo"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.target", "activity"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.from_date", "2025-03-10"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.end_date", "2025-03-18"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.activity_type_ids.0", "1"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.activity_type_ids.1", "6"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.activity_type_ids.2", "12"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.polling_interval_second", "120"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.bulk_job_timeout_second", "7200"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.input_option_columns.0.name", "lead_id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.input_option_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.input_option_columns.1.name", "activity_type"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.input_option_columns.1.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "marketing_data"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "marketo_activities"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "merge"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.bigquery_output_option_merge_keys.0", "lead_id"),
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

func TestAccJobDefinitionResourceMarketoCustomObjectToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.marketo_custom_object_with_filter"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/marketo_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Marketo to BigQuery - Custom Object with Filter"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "marketo"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.target", "custom_object"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.custom_object_api_name", "company"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.custom_object_filter_type", "id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.custom_object_filter_from_value", "1000"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.custom_object_filter_to_value", "2000"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.custom_object_fields.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.custom_object_fields.1.name", "name"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.input_option_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.input_option_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "marketing_data"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "marketo_custom_objects"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "merge"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.bigquery_output_option_merge_keys.0", "id"),
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

func TestAccJobDefinitionResourceMarketoFolderToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.marketo_folder_transfer"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/marketo_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Marketo to BigQuery - Folder Transfer"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "marketo"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.target", "folder"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.root_type", "program"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.root_id", "456"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.max_depth", "3"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.workspace", "Marketing"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.input_option_columns.0.name", "id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.input_option_columns.0.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "marketing_data"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "marketo_folders"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "merge"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.bigquery_output_option_merge_keys.0", "id"),
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

func TestAccJobDefinitionResourceMarketoWithCustomVariablesToBigQuery(t *testing.T) {
	resourceName := "trocco_job_definition.marketo_with_custom_variables_to_bigquery"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextFile("testdata/fixtures/bigquery_connection.tf") + LoadTextFile("testdata/job_definition/marketo_to_bigquery/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "Marketo to BigQuery - Dynamic Configuration"),
					resource.TestCheckResourceAttr(resourceName, "input_option_type", "marketo"),
					resource.TestCheckResourceAttr(resourceName, "output_option_type", "bigquery"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.target", "lead"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.from_date", "$start_date$"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.end_date", "$end_date$"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.custom_variable_settings.0.name", "$start_date$"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.custom_variable_settings.0.type", "timestamp_runtime"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.custom_variable_settings.0.quantity", "7"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.custom_variable_settings.0.unit", "date"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.custom_variable_settings.0.direction", "ago"),
					resource.TestCheckResourceAttr(resourceName, "input_option.marketo_input_option.custom_variable_settings.1.name", "$end_date$"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.dataset", "marketing_data"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.table", "marketo_leads_dynamic"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.mode", "merge"),
					resource.TestCheckResourceAttr(resourceName, "output_option.bigquery_output_option.bigquery_output_option_merge_keys.0", "id"),
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

// TestAccJobDefinitionResourcePagerdutyToBigQuery is auto-generated by tool. DO NOT EDIT.
