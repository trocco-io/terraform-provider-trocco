resource "trocco_snowflake_datamart_definition" "scd_type_2" {
  name                     = "example_scd_type_2"
  is_runnable_concurrently = false
  snowflake_connection_id  = 1
  query_mode               = "insert"
  query                    = "SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE"
  warehouse                = "EXAMPLE_WH"
  destination_database     = "DEST_DATABASE"
  destination_schema       = "DEST_SCHEMA"
  destination_table        = "DEST_TABLE"

  write_disposition     = "scd_type_2"
  merge_keys            = ["id"]
  incremental_column    = "updated_at"
  schema_evolution_mode = "detect_only"
}
