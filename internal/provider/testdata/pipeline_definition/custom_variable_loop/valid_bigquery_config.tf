resource "trocco_connection" "my_conn" {
  connection_type = "bigquery"

  name        = "BigQuery Example"
  description = "This is a BigQuery connection example"

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

resource "trocco_bigquery_datamart_definition" "test_bigquery_datamart_notifications" {
  name                     = "test_bigquery_datamart_notifications"
  is_runnable_concurrently = false
  bigquery_connection_id   = trocco_connection.my_conn.id
  query                    = "SELECT * FROM examples"
  query_mode               = "insert"
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "append"

  custom_variable_settings = [
    {
      name  = "$foo$"
      type  = "string"
      value = "x"
    },
    {
      name  = "$bar$"
      type  = "string"
      value = "y"
    },
  ]
}

resource "trocco_pipeline_definition" "trocco_bigquery_datamart" {
  name = "trocco_bigquery_datamart"

  tasks = [
    {
      key  = "trocco_bigquery_datamart"
      type = "trocco_bigquery_datamart"

      trocco_bigquery_datamart_config = {
        definition_id = trocco_bigquery_datamart_definition.test_bigquery_datamart_notifications.id
        custom_variable_loop = {
          type = "bigquery"
          bigquery_config = {
            connection_id = trocco_connection.my_conn.id  # ← 動的に参照
            query         = "SELECT foo, bar FROM sample"
            variables     = ["$foo$", "$bar$"]
          }
        }
      }
    }
  ]
}
