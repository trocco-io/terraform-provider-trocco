resource "trocco_snowflake_datamart_definition" "query_mode" {
  name                     = "example_query_mode"
  is_runnable_concurrently = false
  snowflake_connection_id  = 1
  query_mode               = "query"
  query                    = "CREATE OR REPLACE TABLE DEST_DATABASE.DEST_SCHEMA.DEST_TABLE AS SELECT * FROM SOURCE_DATABASE.SOURCE_SCHEMA.SOURCE_TABLE"
  warehouse                = "EXAMPLE_WH"
  statement_timeout        = 3600
}
