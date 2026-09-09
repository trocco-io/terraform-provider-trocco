resource "trocco_snowflake_datamart_definition" "with_labels" {
  name                     = "example_with_labels"
  is_runnable_concurrently = false
  snowflake_connection_id  = 1
  query_mode               = "insert"
  query                    = "SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE"
  warehouse                = "EXAMPLE_WH"
  destination_database     = "DEST_DATABASE"
  destination_schema       = "DEST_SCHEMA"
  destination_table        = "DEST_TABLE"
  write_disposition        = "append"
  labels = [
    {
      name = "test_label_1"
    },
    {
      name = "test_label_2"
    }
  ]
}
