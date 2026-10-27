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

  quality_check_enabled = false
}
