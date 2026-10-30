resource "trocco_connection" "snowflake_scd" {
  connection_type = "snowflake"
  auth_method     = "key_pair"

  name        = "snowflake scd_type_2 test"
  host        = "example.snowflakecomputing.com"
  user_name   = "root"
  private_key = "-----BEGIN PRIVATE KEY-----\ndummy\n-----END PRIVATE KEY-----\n"
}

resource "trocco_snowflake_datamart_definition" "test_snowflake_datamart_scd_type_2" {
  name                     = "test_snowflake_datamart_scd_type_2"
  is_runnable_concurrently = false
  snowflake_connection_id  = trocco_connection.snowflake_scd.id
  query_mode               = "insert"
  query                    = <<SQL
    SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE
  SQL
  warehouse                = "EXAMPLE_WH"
  destination_database     = "DEST_DATABASE"
  destination_schema       = "DEST_SCHEMA"
  destination_table        = "DEST_TABLE"

  write_disposition     = "scd_type_2"
  merge_keys            = ["id"]
  incremental_column    = "updated_at"
  schema_evolution_mode = "detect_only"
}
