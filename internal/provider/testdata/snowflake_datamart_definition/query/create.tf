resource "trocco_connection" "snowflake_test_query" {
  connection_type = "snowflake"
  auth_method     = "key_pair"

  name        = "snowflake test query"
  host        = "example.snowflakecomputing.com"
  user_name   = "root"
  private_key = "-----BEGIN PRIVATE KEY-----\ndummy\n-----END PRIVATE KEY-----\n"
}

resource "trocco_snowflake_datamart_definition" "test_snowflake_datamart_query" {
  name                     = "test_snowflake_datamart_query"
  is_runnable_concurrently = false
  snowflake_connection_id  = trocco_connection.snowflake_test_query.id
  query_mode               = "query"
  query                    = <<SQL
    CREATE OR REPLACE TABLE EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE AS
    SELECT * FROM SOURCE_DATABASE.SOURCE_SCHEMA.SOURCE_TABLE
  SQL
  warehouse                = "EXAMPLE_WH"
  statement_timeout        = 3600
}
