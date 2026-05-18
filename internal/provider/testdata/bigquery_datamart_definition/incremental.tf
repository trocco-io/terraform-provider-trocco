resource "trocco_connection" "my_conn_incremental" {
  connection_type = "bigquery"

  name        = "BigQuery Example Incremental"
  description = "This is a BigQuery connection example for incremental"

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

resource "trocco_bigquery_datamart_definition" "test_incremental" {
  name                        = "test_incremental"
  is_runnable_concurrently    = false
  bigquery_connection_id      = trocco_connection.my_conn_incremental.id
  query                       = <<SQL
    SELECT * FROM examples
  SQL
  query_mode                  = "insert"
  destination_dataset         = "dist_datasets"
  destination_table           = "dist_tables"
  write_disposition           = "incremental"
  merge_keys                  = ["id"]
  on_matched_action           = "upsert"
  schema_evolution_mode       = "detect_only"
  lookback_period_column      = "updated_at"
  lookback_period_column_type = "TIMESTAMP"
  lookback_period_timezone    = "Asia/Tokyo"
  lookback_period_from        = 3
  lookback_period_to          = 0
  lookback_period_unit        = "days"
}
