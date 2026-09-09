resource "trocco_snowflake_datamart_definition" "with_notifications" {
  name                     = "example_with_notifications"
  is_runnable_concurrently = false
  snowflake_connection_id  = 1
  query_mode               = "insert"
  query                    = "SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE"
  warehouse                = "EXAMPLE_WH"
  destination_database     = "DEST_DATABASE"
  destination_schema       = "DEST_SCHEMA"
  destination_table        = "DEST_TABLE"
  write_disposition        = "append"
  notifications = [
    {
      destination_type  = "slack"
      slack_channel_id  = 1
      notification_type = "job"
      notify_when       = "finished"
      message           = "@here Job finished."
    },
    {
      destination_type  = "email"
      email_id          = 1
      notification_type = "record"
      record_count      = 100
      record_operator   = "below"
      message           = "Record count is below 100."
    }
  ]
}
