resource "trocco_connection" "my_conn" {
  connection_type = "bigquery"

  name        = "BigQuery Example"
  description = "This is a BigQuery connection example"

  project_id               = "example"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}

resource "trocco_notification_destination" "slack" {
  type = "slack_channel"
  slack_channel_config = {
    channel     = "example"
    webhook_url = "https://hooks.slack.com/services/xxx/yyy/zzz"
  }
}

resource "trocco_notification_destination" "email" {
  type = "email"
  email_config = {
    email = "example@example.com"
  }
}

resource "trocco_bigquery_datamart_definition" "test_bigquery_datamart_notifications" {
  name                     = "test_bigquery_datamart_notifications"
  is_runnable_concurrently = false
  bigquery_connection_id   = trocco_connection.my_conn.id
  query                    = <<SQL
    SELECT * FROM examples
  SQL
  query_mode               = "insert"
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "append"

  notifications = [
    {
      destination_type   = "slack"
      slack_channel_id   = trocco_notification_destination.slack.id
      notification_type  = "job"
      notify_when        = "finished"
      message            = <<MESSAGE
This is a multi-line message
with several lines
  and some indentation
    to test TrimmedStringType
MESSAGE
    },
    {
      destination_type   = "email"
      email_id           = trocco_notification_destination.email.id
      notification_type  = "record"
      record_count       = 100
      record_operator    = "above"
      message            = <<MESSAGE
  This is another multi-line message
with leading and trailing whitespace
  
  to test TrimmedStringType
  
MESSAGE
    }
  ]
}
