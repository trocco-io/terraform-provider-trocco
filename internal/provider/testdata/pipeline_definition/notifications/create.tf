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

resource "trocco_pipeline_definition" "notifications_test" {
  name = "notifications_test"

  notifications = [
    {
      type             = "job_execution"
      destination_type = "slack"
      notify_when      = "finished"

      slack_config = {
        notification_id = trocco_notification_destination.slack.id
        message         = <<MESSAGE
This is a multi-line message
with several lines
  and some indentation
    to test TrimmedStringType
MESSAGE
      }
    },
    {
      type             = "job_time_alert"
      destination_type = "email"
      time             = 5

      email_config = {
        notification_id = trocco_notification_destination.email.id
        message         = <<MESSAGE
  This is another multi-line message
with leading and trailing whitespace
  
  to test TrimmedStringType
  
MESSAGE
      }
    },
  ]

  tasks = [
    {
      key  = "bigquery_data_check"
      type = "bigquery_data_check"

      bigquery_data_check_config = {
        name          = "Example"
        connection_id = trocco_connection.my_conn.id
        query         = <<SQL
          SELECT COUNT(*) FROM examples
        SQL
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
      }
    }
  ]
}
