resource "trocco_connection" "snowflake_notifications" {
  connection_type = "snowflake"
  auth_method     = "key_pair"

  name        = "snowflake notifications test"
  host        = "example.snowflakecomputing.com"
  user_name   = "root"
  private_key = "-----BEGIN PRIVATE KEY-----\ndummy\n-----END PRIVATE KEY-----\n"
}

resource "trocco_notification_destination" "slack" {
  type = "slack_channel"
  slack_channel_config = {
    channel     = "example"
    webhook_url = "https://hooks.slack.com/services/xxx/yyy/zzz"
  }
}

resource "trocco_notification_destination" "slack_b" {
  type = "slack_channel"
  slack_channel_config = {
    channel     = "channel-b"
    webhook_url = "https://hooks.slack.com/services/ddd/eee/fff"
  }
}

resource "trocco_notification_destination" "email" {
  type = "email"
  email_config = {
    email = "example@example.com"
  }
}

resource "trocco_snowflake_datamart_definition" "test_snowflake_datamart_notifications" {
  name                     = "test_snowflake_datamart_notifications"
  is_runnable_concurrently = false
  snowflake_connection_id  = trocco_connection.snowflake_notifications.id
  query_mode               = "insert"
  query                    = <<SQL
    SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE
  SQL
  warehouse                = "EXAMPLE_WH"
  destination_database     = "DEST_DATABASE"
  destination_schema       = "DEST_SCHEMA"
  destination_table        = "DEST_TABLE"
  write_disposition        = "append"

  notifications = [
    {
      destination_type  = "slack"
      slack_channel_id  = trocco_notification_destination.slack.id
      notification_type = "job"
      notify_when       = "finished"
      message           = "message for slack channel"
    },
    {
      destination_type  = "slack"
      slack_channel_id  = trocco_notification_destination.slack_b.id
      notification_type = "job"
      notify_when       = "failed"
      message           = "message for slack channel b"
    }
  ]
}
