resource "trocco_connection" "snowflake_transition" {
  connection_type = "snowflake"
  auth_method     = "key_pair"

  name        = "snowflake transition test"
  host        = "example.snowflakecomputing.com"
  user_name   = "root"
  private_key = "-----BEGIN PRIVATE KEY-----\ndummy\n-----END PRIVATE KEY-----\n"
}

resource "trocco_notification_destination" "transition_slack" {
  type = "slack_channel"
  slack_channel_config = {
    channel     = "transition"
    webhook_url = "https://hooks.slack.com/services/xxx/yyy/zzz"
  }
}

resource "trocco_notification_destination" "transition_email" {
  type = "email"
  email_config = {
    email = "transition@example.com"
  }
}

resource "trocco_snowflake_datamart_definition" "test_snowflake_datamart_transition" {
  name                     = "test_snowflake_datamart_transition"
  is_runnable_concurrently = false
  snowflake_connection_id  = trocco_connection.snowflake_transition.id
  query_mode               = "insert"
  query                    = <<SQL
    SELECT * FROM EXAMPLE_DATABASE.EXAMPLE_SCHEMA.EXAMPLE_TABLE
  SQL
  warehouse                = "EXAMPLE_WH"
  destination_database     = "DEST_DATABASE"
  destination_schema       = "DEST_SCHEMA"
  destination_table        = "DEST_TABLE"

  notifications = [
    {
      destination_type  = "slack"
      slack_channel_id  = trocco_notification_destination.transition_slack.id
      notification_type = "job"
      notify_when       = "finished"
      message           = "transition job notification"
    },
    {
      destination_type  = "email"
      email_id          = trocco_notification_destination.transition_email.id
      notification_type = "record"
      record_count      = 10
      record_operator   = "below"
      message           = "transition record notification"
    }
  ]

  schedules = [
    {
      frequency = "hourly"
      minute    = 5
      time_zone = "Asia/Tokyo"
    },
    {
      frequency = "daily"
      hour      = 3
      minute    = 30
      time_zone = "Asia/Tokyo"
    },
    {
      frequency   = "weekly"
      day_of_week = 1
      hour        = 6
      minute      = 15
      time_zone   = "Asia/Tokyo"
    }
  ]

  write_disposition     = "scd_type_2"
  merge_keys            = ["id"]
  incremental_column    = "updated_at"
  schema_evolution_mode = "auto_add_column"
}
