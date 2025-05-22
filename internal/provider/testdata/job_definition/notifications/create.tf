resource "trocco_connection" "test_mysql" {
  connection_type = "mysql"
  name            = "MySQL Example"
  host            = "db.example.com"
  port            = 65535
  user_name       = "root"
  password        = "password"
}

resource "trocco_connection" "test_bq" {
  connection_type          = "bigquery"
  name                     = "BigQuery Example"
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

resource "trocco_job_definition" "notifications_test" {
  name                     = "notifications_test"
  description              = "Test job definition with notifications"
  resource_enhancement     = "medium"
  retry_limit              = 0
  is_runnable_concurrently = false
  
  filter_columns = [
    {
      name                         = "id"
      src                          = "id"
      type                         = "long"
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    },
    {
      name                         = "name"
      src                          = "name"
      type                         = "string"
      default                      = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
    }
  ]
  
  input_option_type = "mysql"
  input_option = {
    mysql_input_option = {
      database                    = "test_database"
      table                       = "test_table"
      fetch_rows                  = 1000
      default_time_zone           = "Asia/Tokyo"
      incremental_loading_enabled = false
      mysql_connection_id         = trocco_connection.test_mysql.id
      input_option_columns = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "name"
          type = "string"
        }
      ]
    }
  }
  
  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                                  = "test_dataset"
      table                                    = "test_table"
      mode                                     = "append"
      location                                 = "US"
      bigquery_connection_id                   = trocco_connection.test_bq.id
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }

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
      notification_type  = "exec_time"
      minutes            = 30
      message            = <<MESSAGE
  This is another multi-line message
with leading and trailing whitespace
  
  to test TrimmedStringType
  
MESSAGE
    }
  ]
}
