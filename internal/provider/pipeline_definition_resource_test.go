package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccPipelineDefinitionResource_SnowflakeDataCheck(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with Windows line endings
			{
				Config: providerConfig + `
resource "trocco_connection" "snowflake" {
  connection_type = "snowflake"
  name = "Snowflake Test"
  host = "example.snowflakecomputing.com"
  auth_method = "user_password"
  user_name = "test_user"
  password = "test_password"
}

resource "trocco_pipeline_definition" "test" {
  name = "test_snowflake_data_check"
  
  tasks = [
    {
      key  = "snowflake_data_check"
      type = "snowflake_data_check"

      snowflake_data_check_config = {
        name          = "Example"
        connection_id = trocco_connection.snowflake.id
        query         = "SELECT COUNT() AS count_of_time\r\nFROM sample.test.test\r\nWHERE time = '$day$'"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
        warehouse     = "EXAMPLE"
      }
    }
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "name", "test_snowflake_data_check"),
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "tasks.0.type", "snowflake_data_check"),
				),
			},
			// Import testing
			{
				ResourceName:            "trocco_pipeline_definition.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"tasks.0.key"}, // Skip key verification (will be fixed in a future PR)
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					pipelineDefinitionID := s.RootModule().Resources["trocco_pipeline_definition.test"].Primary.ID
					return pipelineDefinitionID, nil
				},
			},
			// Update with Unix line endings (should not cause a plan difference)
			{
				Config: providerConfig + `
resource "trocco_connection" "snowflake" {
  connection_type = "snowflake"
  name = "Snowflake Test"
  host = "example.snowflakecomputing.com"
  auth_method = "user_password"
  user_name = "test_user"
  password = "test_password"
}

resource "trocco_pipeline_definition" "test" {
  name = "test_snowflake_data_check"
  
  tasks = [
    {
      key  = "snowflake_data_check"
      type = "snowflake_data_check"

      snowflake_data_check_config = {
        name          = "Example"
        connection_id = trocco_connection.snowflake.id
        query         = "SELECT COUNT() AS count_of_time\nFROM sample.test.test\nWHERE time = '$day$'"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
        warehouse     = "EXAMPLE"
      }
    }
  ]
}
`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // Expect no changes
			},
			// Update with different whitespace (should not cause a plan difference)
			{
				Config: providerConfig + `
resource "trocco_connection" "snowflake" {
  connection_type = "snowflake"
  name = "Snowflake Test"
  host = "example.snowflakecomputing.com"
  auth_method = "user_password"
  user_name = "test_user"
  password = "test_password"
}

resource "trocco_pipeline_definition" "test" {
  name = "test_snowflake_data_check"
  
  tasks = [
    {
      key  = "snowflake_data_check"
      type = "snowflake_data_check"

      snowflake_data_check_config = {
        name          = "Example"
        connection_id = trocco_connection.snowflake.id
        query         = "SELECT COUNT() AS count_of_time      \nFROM sample.test.test        \nWHERE time = '$day$'"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
        warehouse     = "EXAMPLE"
      }
    }
  ]
}
`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // Expect no changes
			},
		},
	})
}

func TestAccPipelineDefinitionResource_BigqueryDataCheck(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with Windows line endings
			{
				Config: providerConfig + `
resource "trocco_connection" "bigquery" {
  connection_type = "bigquery"
  name = "BigQuery Test"
  project_id = "test-project"
  service_account_json_key = "{\"type\":\"service_account\",\"project_id\":\"\",\"private_key_id\":\"\",\"private_key\":\"\"}"
}

resource "trocco_pipeline_definition" "test" {
  name = "test_bigquery_data_check"
  
  tasks = [
    {
      key  = "bigquery_data_check"
      type = "bigquery_data_check"

      bigquery_data_check_config = {
        name          = "Example"
        connection_id = trocco_connection.bigquery.id
        query         = "SELECT COUNT(*) AS count\r\nFROM project_dataset_table\r\nWHERE date = '$day$'\r\n;"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
      }
    }
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "name", "test_bigquery_data_check"),
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "tasks.0.type", "bigquery_data_check"),
				),
			},
			// Import testing
			{
				ResourceName:            "trocco_pipeline_definition.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"tasks.0.key"}, // Skip key verification (will be fixed in a future PR)
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					pipelineDefinitionID := s.RootModule().Resources["trocco_pipeline_definition.test"].Primary.ID
					return pipelineDefinitionID, nil
				},
			},
			// Update with Unix line endings (should not cause a plan difference)
			{
				Config: providerConfig + `
resource "trocco_connection" "bigquery" {
  connection_type = "bigquery"
  name = "BigQuery Test"
  project_id = "test-project"
  service_account_json_key = "{\"type\":\"service_account\",\"project_id\":\"\",\"private_key_id\":\"\",\"private_key\":\"\"}"
}

resource "trocco_pipeline_definition" "test" {
  name = "test_bigquery_data_check"
  
  tasks = [
    {
      key  = "bigquery_data_check"
      type = "bigquery_data_check"

      bigquery_data_check_config = {
        name          = "Example"
        connection_id = trocco_connection.bigquery.id
        query         = "SELECT COUNT(*) AS count\nFROM project_dataset_table\nWHERE date = '$day$'\n;"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
      }
    }
  ]
}
`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // Expect no changes
			},
		},
	})
}

func TestAccPipelineDefinitionResource_RedshiftDataCheck(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with Windows line endings
			{
				Config: providerConfig + `
resource "trocco_pipeline_definition" "test" {
  name = "test_redshift_data_check"
  
  tasks = [
    {
      key  = "redshift_data_check"
      type = "redshift_data_check"

      redshift_data_check_config = {
        name          = "Example"
        connection_id = 256  // FYI: this is magic number. redshift connection creation via terraform has not been implemented yet
        query         = "SELECT COUNT(*) AS count\r\nFROM schema.table\r\nWHERE date = '$day$'"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
        database      = "test_db"
      }
    }
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "name", "test_redshift_data_check"),
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "tasks.0.type", "redshift_data_check"),
				),
			},
			// Import testing
			{
				ResourceName:            "trocco_pipeline_definition.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"tasks.0.key"}, // Skip key verification (will be fixed in a future PR)
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					pipelineDefinitionID := s.RootModule().Resources["trocco_pipeline_definition.test"].Primary.ID
					return pipelineDefinitionID, nil
				},
			},
			// Update with Unix line endings (should not cause a plan difference)
			{
				Config: providerConfig + `
resource "trocco_pipeline_definition" "test" {
  name = "test_redshift_data_check"
  
  tasks = [
    {
      key  = "redshift_data_check"
      type = "redshift_data_check"

      redshift_data_check_config = {
        name          = "Example"
        connection_id = 256  // FYI: this is magic number. redshift connection creation via terraform has not been implemented yet
        query         = "SELECT COUNT(*) AS count\nFROM schema.table\nWHERE date = '$day$'"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
        database      = "test_db"
      }
    }
  ]
}
`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // Expect no changes
			},
		},
	})
}

func TestAccPipelineDefinitionResource_DatamartWithCustomVariableLoop(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
// BigQuery connection definition
resource "trocco_connection" "bigquery" {
  connection_type = "bigquery"
  name = "BigQuery Test"
  project_id = "test-project"
  service_account_json_key = "{\"type\":\"service_account\",\"project_id\":\"\",\"private_key_id\":\"\",\"private_key\":\"\"}"
}

// Snowflake connection definition
resource "trocco_connection" "snowflake" {
  connection_type = "snowflake"
  name = "Snowflake Test"
  host = "example.snowflakecomputing.com"
  auth_method = "user_password"
  user_name = "test_user"
  password = "test_password"
}

// BigQuery datamart definition
resource "trocco_bigquery_datamart_definition" "bigquery_datamart" {
  name                     = "BQ datamart sample"
  is_runnable_concurrently = false
  bigquery_connection_id   = trocco_connection.bigquery.id
  query                    = "SELECT * FROM $table_name$"
  query_mode               = "query"
  location                 = "asia-northeast1"
  custom_variable_settings = [
    {
      name  = "$table_name$",
      type  = "string",
      value = "foo"
    }
  ]
}

// Pipeline definition
resource "trocco_pipeline_definition" "test" {
  name = "test_datamart_with_custom_variable_loop"
  
  tasks = [
    {
      key  = "bigquery_datamart_with_bigquery_loop"
      type = "trocco_bigquery_datamart"
      trocco_bigquery_datamart_config = {
        definition_id = trocco_bigquery_datamart_definition.bigquery_datamart.id
        custom_variable_loop = {
          type = "bigquery"
          bigquery_config = {
            connection_id = trocco_connection.bigquery.id
            query         = "SELECT DISTINCT table_name\r\nFROM project.INFORMATION_SCHEMA.TABLES\r\nWHERE table_schema = 'public'\r\nORDER BY table_name"
            variables     = ["$table_name$"]
          }
        }
      }
    },
    {
      key  = "bigquery_datamart_with_snowflake_loop"
      type = "trocco_bigquery_datamart"
      trocco_bigquery_datamart_config = {
        definition_id = trocco_bigquery_datamart_definition.bigquery_datamart.id
        custom_variable_loop = {
          type = "snowflake"
          snowflake_config = {
            connection_id = trocco_connection.snowflake.id
            query         = "SELECT DISTINCT table_name\r\nFROM information_schema.tables\r\nWHERE table_schema = 'public'\r\nORDER BY table_name"
            variables     = ["$table_name$"]
            warehouse     = "COMPUTE_WH"
          }
        }
      }
    },
    {
      key  = "bigquery_datamart_with_redshift_loop"
      type = "trocco_bigquery_datamart"
      trocco_bigquery_datamart_config = {
        definition_id = trocco_bigquery_datamart_definition.bigquery_datamart.id
        custom_variable_loop = {
          type = "redshift"
          redshift_config = {
            connection_id = 256  // FYI: this is magic number. redshift connection creation via terraform has not been implemented yet
            query         = "SELECT DISTINCT table_name\r\nFROM information_schema.tables\r\nWHERE table_schema = 'public'\r\nORDER BY table_name"
            variables     = ["$table_name$"]
            database      = "test_db"
          }
        }
      }
    }
  ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "name", "test_datamart_with_custom_variable_loop"),
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "tasks.0.type", "trocco_bigquery_datamart"),
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "tasks.0.trocco_bigquery_datamart_config.custom_variable_loop.bigquery_config.variables.0", "$table_name$"),
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "tasks.1.type", "trocco_bigquery_datamart"),
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "tasks.1.trocco_bigquery_datamart_config.custom_variable_loop.snowflake_config.variables.0", "$table_name$"),
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "tasks.2.type", "trocco_bigquery_datamart"),
					resource.TestCheckResourceAttr("trocco_pipeline_definition.test", "tasks.2.trocco_bigquery_datamart_config.custom_variable_loop.redshift_config.variables.0", "$table_name$"),
				),
			},
			// Unix line endings test
			{
				Config: providerConfig + `
// BigQuery connection definition
resource "trocco_connection" "bigquery" {
  connection_type = "bigquery"
  name = "BigQuery Test"
  project_id = "test-project"
  service_account_json_key = "{\"type\":\"service_account\",\"project_id\":\"\",\"private_key_id\":\"\",\"private_key\":\"\"}"
}

// Snowflake connection definition
resource "trocco_connection" "snowflake" {
  connection_type = "snowflake"
  name = "Snowflake Test"
  host = "example.snowflakecomputing.com"
  auth_method = "user_password"
  user_name = "test_user"
  password = "test_password"
}

// BigQuery datamart definition
resource "trocco_bigquery_datamart_definition" "bigquery_datamart" {
  name                     = "BQ datamart sample"
  is_runnable_concurrently = false
  bigquery_connection_id   = trocco_connection.bigquery.id
  query                    = "SELECT * FROM $table_name$"
  query_mode               = "query"
  location                 = "asia-northeast1"
    custom_variable_settings = [
    {
      name  = "$table_name$",
      type  = "string",
      value = "foo"
    }
  ]
}

// Pipeline definition
resource "trocco_pipeline_definition" "test" {
  name = "test_datamart_with_custom_variable_loop"
  
  tasks = [
    {
      key  = "bigquery_datamart_with_bigquery_loop"
      type = "trocco_bigquery_datamart"
      trocco_bigquery_datamart_config = {
        definition_id = trocco_bigquery_datamart_definition.bigquery_datamart.id
        custom_variable_loop = {
          type = "bigquery"
          bigquery_config = {
            connection_id = trocco_connection.bigquery.id
            query         = "SELECT DISTINCT table_name\nFROM project.INFORMATION_SCHEMA.TABLES\nWHERE table_schema = 'public'\nORDER BY table_name"
            variables     = ["$table_name$"]
          }
        }
      }
    },
    {
      key  = "bigquery_datamart_with_snowflake_loop"
      type = "trocco_bigquery_datamart"
      trocco_bigquery_datamart_config = {
        definition_id = trocco_bigquery_datamart_definition.bigquery_datamart.id
        custom_variable_loop = {
          type = "snowflake"
          snowflake_config = {
            connection_id = trocco_connection.snowflake.id
            query         = "SELECT DISTINCT table_name\nFROM information_schema.tables\nWHERE table_schema = 'public'\nORDER BY table_name"
            variables     = ["$table_name$"]
            warehouse     = "COMPUTE_WH"
          }
        }
      }
    },
    {
      key  = "bigquery_datamart_with_redshift_loop"
      type = "trocco_bigquery_datamart"
      trocco_bigquery_datamart_config = {
        definition_id = trocco_bigquery_datamart_definition.bigquery_datamart.id
        custom_variable_loop = {
          type = "redshift"
          redshift_config = {
            connection_id = 256  // FYI: this is magic number. redshift connection creation via terraform has not been implemented yet
            query         = "SELECT DISTINCT table_name\nFROM information_schema.tables\nWHERE table_schema = 'public'\nORDER BY table_name"
            variables     = ["$table_name$"]
            database      = "test_db"
          }
        }
      }
    }
  ]
}
`,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // Verify that there are no differences due to line ending variations
			},
		},
	})
}
