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
      mysql_connection_id         = trocco_connection.mysql.id
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
      bigquery_connection_id                   = trocco_connection.bigquery.id
      bigquery_output_option_clustering_fields = []
      bigquery_output_option_column_options    = []
      bigquery_output_option_merge_keys        = []
    }
  }

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
