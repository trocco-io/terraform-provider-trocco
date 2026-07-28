resource "trocco_snowflake_datamart_definition" "incremental" {
  name                     = "example_incremental"
  is_runnable_concurrently = false
  snowflake_connection_id  = 1
  query_mode               = "insert"
  query                    = "SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE"
  warehouse                = "EXAMPLE_WH"
  destination_database     = "DEST_DATABASE"
  destination_schema       = "DEST_SCHEMA"
  destination_table        = "DEST_TABLE"

  write_disposition           = "incremental"
  merge_keys                  = ["id"]
  on_matched_action           = "upsert"
  schema_evolution_mode       = "detect_only"
  lookback_period_column      = "updated_at"
  lookback_period_column_type = "TIMESTAMP_NTZ"
  lookback_period_timezone    = "Asia/Tokyo"
  lookback_period_from        = 3
  lookback_period_to          = 0
  lookback_period_unit        = "days"
}
