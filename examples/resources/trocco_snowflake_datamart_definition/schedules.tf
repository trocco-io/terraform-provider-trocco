resource "trocco_snowflake_datamart_definition" "with_schedules" {
  name                     = "example_with_schedules"
  is_runnable_concurrently = false
  snowflake_connection_id  = 1
  query_mode               = "insert"
  query                    = "SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE"
  warehouse                = "EXAMPLE_WH"
  destination_database     = "DEST_DATABASE"
  destination_schema       = "DEST_SCHEMA"
  destination_table        = "DEST_TABLE"
  write_disposition        = "append"
  schedules = [
    {
      frequency = "hourly"
      minute    = 0
      time_zone = "Asia/Tokyo"
    },
    {
      frequency = "daily"
      hour      = 0
      minute    = 0
      time_zone = "Asia/Tokyo"
    },
    {
      frequency   = "weekly"
      day_of_week = 0
      hour        = 0
      minute      = 0
      time_zone   = "Asia/Tokyo"
    },
    {
      frequency = "monthly"
      day       = 1
      hour      = 0
      minute    = 0
      time_zone = "Asia/Tokyo"
    }
  ]
}
