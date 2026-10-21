resource "trocco_connection" "snowflake_test_invalid" {
  connection_type = "snowflake"
  auth_method     = "key_pair"

  name        = "snowflake test invalid"
  host        = "example.snowflakecomputing.com"
  user_name   = "root"
  private_key = "-----BEGIN PRIVATE KEY-----\ndummy\n-----END PRIVATE KEY-----\n"
}

resource "trocco_snowflake_datamart_definition" "test_missing_required_insert_fields" {
  name                     = "test_missing_required_insert_fields"
  is_runnable_concurrently = false
  snowflake_connection_id  = trocco_connection.snowflake_test_invalid.id
  query_mode               = "insert"
  query                    = "SELECT * FROM examples"
  warehouse                = "EXAMPLE_WH"
  # Missing: destination_database, destination_schema, destination_table, write_disposition
}
