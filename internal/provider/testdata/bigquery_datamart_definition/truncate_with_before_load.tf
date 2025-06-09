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

resource "trocco_bigquery_datamart_definition" "test_truncate_with_before_load" {
  name                     = "test_truncate_with_before_load"
  is_runnable_concurrently = false
  bigquery_connection_id   = trocco_connection.my_conn.id
  query                    = <<SQL
    SELECT * FROM examples
  SQL
  query_mode               = "insert"
  before_load              = <<SQL
    DELETE FROM examples
    WHERE created_at < '2024-01-01'
  SQL
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "truncate"
}
