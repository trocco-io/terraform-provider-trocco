resource "trocco_snowflake_datamart_definition" "with_optionals" {
  name                     = "example_with_optionals"
  description              = "This is an example with optional fields"
  is_runnable_concurrently = false
  resource_group_id        = 1
  custom_variable_settings = [
    {
      name  = "$string$",
      type  = "string",
      value = "foo",
    },
    {
      name      = "$timestamp$",
      type      = "timestamp",
      quantity  = 1,
      unit      = "hour",
      direction = "ago",
      format    = "%Y-%m-%d %H:%M:%S",
      time_zone = "Asia/Tokyo",
    }
  ]
  snowflake_connection_id = 1
  query_mode              = "insert"
  query                   = "SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE"
  warehouse               = "EXAMPLE_WH"
  statement_timeout       = 43200
  destination_database    = "DEST_DATABASE"
  destination_schema      = "DEST_SCHEMA"
  destination_table       = "DEST_TABLE"
  write_disposition       = "append"
}
