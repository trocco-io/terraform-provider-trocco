resource "trocco_datamart_definition" "with_notifications" {
  name                     = "example_with_notifications"
  data_warehouse_type      = "bigquery"
  is_runnable_concurrently = false
  datamart_bigquery_option = {
    bigquery_connection_id = 1
    query                  = "SELECT * FROM tables"
    query_mode             = "insert"
    destination_dataset    = "dist_datasets"
    destination_table      = "dist_tables"
    write_disposition      = "append"
  }
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
