resource "trocco_connection" "my_conn_scd" {
  connection_type = "bigquery"

  name        = "BigQuery Example SCD"
  description = "This is a BigQuery connection example for scd_type_2"

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

resource "trocco_bigquery_datamart_definition" "test_scd_type_2" {
  name                     = "test_scd_type_2"
  is_runnable_concurrently = false
  bigquery_connection_id   = trocco_connection.my_conn_scd.id
  query                    = <<SQL
    SELECT * FROM examples
  SQL
  query_mode               = "insert"
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "scd_type_2"
  merge_keys               = ["id"]
  incremental_column       = "updated_at"
  schema_evolution_mode    = "detect_only"
}
