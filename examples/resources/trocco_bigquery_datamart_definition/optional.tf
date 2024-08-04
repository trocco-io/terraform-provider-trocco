resource "trocco_bigquery_datamart_definition" "with_optionals" {
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
  bigquery_connection_id = 1
  query                  = "SELECT * FROM tables"
  query_mode             = "insert"
  destination_dataset    = "dist_datasets"
  destination_table      = "dist_tables"
  write_disposition      = "append"
}
