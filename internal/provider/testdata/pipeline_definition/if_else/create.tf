resource "trocco_connection" "bigquery_conn" {
  connection_type = "bigquery"

  name        = "BigQuery for if-else test"
  description = "Connection for if-else task test"

  project_id               = "example"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}

resource "trocco_bigquery_datamart_definition" "datamart_def" {
  name                     = "datamart_for_if_else_test"
  is_runnable_concurrently = false
  bigquery_connection_id   = trocco_connection.bigquery_conn.id
  query                    = "SELECT 1"
  query_mode               = "insert"
  destination_dataset      = "test_dataset"
  destination_table        = "test_table"
  write_disposition        = "append"
}

resource "trocco_pipeline_definition" "if_else_test" {
  name        = "if_else_test"
  description = "Test pipeline for if-else task"

  tasks = [
    {
      key  = "datamart"
      type = "trocco_bigquery_datamart"
      trocco_bigquery_datamart_config = {
        definition_id = trocco_bigquery_datamart_definition.datamart_def.id
      }
    },
    {
      key  = "success_datamart"
      type = "trocco_bigquery_datamart"
      trocco_bigquery_datamart_config = {
        definition_id = trocco_bigquery_datamart_definition.datamart_def.id
      }
    },
    {
      key  = "failure_datamart"
      type = "trocco_bigquery_datamart"
      trocco_bigquery_datamart_config = {
        definition_id = trocco_bigquery_datamart_definition.datamart_def.id
      }
    },
    {
      key  = "if_else"
      type = "if_else"
      if_else_config = {
        name = "Check datamart status"
        condition_groups = {
          set_type = "and"
          conditions = [
            {
              variable = "status"
              task_key = "datamart"
              operator = "equal"
              value    = "succeeded"
            }
          ]
        }
        destinations = {
          if   = ["success_datamart"]
          else = ["failure_datamart"]
        }
      }
    }
  ]

  task_dependencies = [
    {
      source      = "datamart"
      destination = "if_else"
    },
    {
      source      = "if_else"
      destination = "success_datamart"
    },
    {
      source      = "if_else"
      destination = "failure_datamart"
    }
  ]
}
