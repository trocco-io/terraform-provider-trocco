resource "trocco_snowflake_datamart_definition" "test_snowflake_datamart_missing_quality_check_fields" {
  name                     = "test_snowflake_datamart_missing_quality_check_fields"
  is_runnable_concurrently = false
  snowflake_connection_id  = 1
  query_mode               = "insert"
  query                    = "SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE"
  warehouse                = "EXAMPLE_WH"
  destination_database     = "DEST_DATABASE"
  destination_schema       = "DEST_SCHEMA"
  destination_table        = "DEST_TABLE"
  write_disposition        = "truncate"

  quality_check_enabled = true
}
