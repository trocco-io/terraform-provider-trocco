resource "trocco_connection" "my_conn_quality_check" {
  connection_type = "bigquery"

  name        = "BigQuery Example Quality Check"
  description = "This is a BigQuery connection example for quality checks"

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

resource "trocco_notification_destination" "slack_quality_check" {
  type = "slack_channel"
  slack_channel_config = {
    channel     = "example"
    webhook_url = "https://hooks.slack.com/services/xxx/yyy/zzz"
  }
}

resource "trocco_bigquery_datamart_definition" "test_quality_check" {
  name                     = "test_quality_check"
  is_runnable_concurrently = false
  bigquery_connection_id   = trocco_connection.my_conn_quality_check.id
  query                    = <<SQL
    SELECT * FROM examples
  SQL
  query_mode               = "insert"
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "append"

  quality_check_enabled                     = true
  quality_check_on_violation                = "warn"
  quality_check_lookback_period_column      = "updated_at"
  quality_check_lookback_period_column_type = "TIMESTAMP"
  quality_check_lookback_period_timezone    = "Asia/Tokyo"
  quality_check_lookback_period_from        = 3
  quality_check_lookback_period_to          = 0
  quality_check_lookback_period_unit        = "days"
  quality_checks = [
    {
      check_type   = "not_null"
      column_names = ["id"]
    },
    {
      check_type   = "composite_unique"
      column_names = ["id", "updated_at"]
    }
  ]

  notifications = [
    {
      destination_type  = "slack"
      slack_channel_id  = trocco_notification_destination.slack_quality_check.id
      notification_type = "job"
      notify_when       = "quality_check_failed"
      message           = "A quality check was violated"
    }
  ]
}
