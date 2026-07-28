resource "trocco_connection" "snowflake_transition" {
  connection_type = "snowflake"
  auth_method     = "key_pair"

  name        = "snowflake transition test"
  host        = "example.snowflakecomputing.com"
  user_name   = "root"
  private_key = "-----BEGIN PRIVATE KEY-----\ndummy\n-----END PRIVATE KEY-----\n"
}

resource "trocco_snowflake_datamart_definition" "test_snowflake_datamart_transition" {
  name                     = "test_snowflake_datamart_transition"
  is_runnable_concurrently = false
  snowflake_connection_id  = trocco_connection.snowflake_transition.id
  query_mode               = "insert"
  query                    = <<SQL
    SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE
  SQL
  warehouse                = "EXAMPLE_WH"
  destination_database     = "DEST_DATABASE"
  destination_schema       = "DEST_SCHEMA"
  destination_table        = "DEST_TABLE"

  write_disposition = "truncate"
}
