resource "trocco_connection" "snowflake_incremental" {
  connection_type = "snowflake"
  auth_method     = "key_pair"

  name        = "snowflake incremental test"
  host        = "example.snowflakecomputing.com"
  user_name   = "root"
  private_key = "-----BEGIN PRIVATE KEY-----\ndummy\n-----END PRIVATE KEY-----\n"
}

resource "trocco_snowflake_datamart_definition" "test_snowflake_datamart_incremental" {
  name                     = "test_snowflake_datamart_incremental"
  is_runnable_concurrently = false
  snowflake_connection_id  = trocco_connection.snowflake_incremental.id
  query_mode               = "insert"
  query                    = <<SQL
    SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE
  SQL
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
