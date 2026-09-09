resource "trocco_connection" "snowflake_quality_checks" {
  connection_type = "snowflake"
  auth_method     = "key_pair"

  name        = "snowflake quality checks test"
  host        = "example.snowflakecomputing.com"
  user_name   = "root"
  private_key = "-----BEGIN PRIVATE KEY-----\ndummy\n-----END PRIVATE KEY-----\n"
}

resource "trocco_snowflake_datamart_definition" "test_snowflake_datamart_quality_checks" {
  name                     = "test_snowflake_datamart_quality_checks"
  is_runnable_concurrently = false
  snowflake_connection_id  = trocco_connection.snowflake_quality_checks.id
  query_mode               = "insert"
  query                    = <<SQL
    SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE
  SQL
  warehouse                = "EXAMPLE_WH"
  destination_database     = "DEST_DATABASE"
  destination_schema       = "DEST_SCHEMA"
  destination_table        = "DEST_TABLE"
  write_disposition        = "append"

  quality_check_enabled                     = true
  quality_check_on_violation                = "fail"
  quality_check_lookback_period_column      = "updated_at"
  quality_check_lookback_period_column_type = "TIMESTAMP_NTZ"
  quality_check_lookback_period_timezone    = "Asia/Tokyo"
  quality_check_lookback_period_from        = 3
  quality_check_lookback_period_to          = 0
  quality_check_lookback_period_unit        = "days"
  quality_checks = [
    {
      check_type   = "not_null"
      column_names = ["id"]
    },
    {
      check_type   = "composite_unique"
      column_names = ["id", "updated_at"]
    },
  ]
}
