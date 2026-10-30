resource "trocco_connection" "my_conn_notify_when" {
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

resource "trocco_notification_destination" "slack_notify_when" {
  type = "slack_channel"
  slack_channel_config = {
    channel     = "example"
    webhook_url = "https://hooks.slack.com/services/xxx/yyy/zzz"
  }
}

resource "trocco_bigquery_datamart_definition" "test_bigquery_datamart_notify_when" {
  name                     = "test_bigquery_datamart_notify_when"
  is_runnable_concurrently = false
  bigquery_connection_id   = trocco_connection.my_conn_notify_when.id
  query                    = <<SQL
    SELECT * FROM examples
  SQL
  query_mode               = "insert"
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "incremental"
  merge_keys               = ["id"]
  on_matched_action        = "upsert"
  schema_evolution_mode    = "detect_only"

  notifications = [
    {
      destination_type  = "slack"
      slack_channel_id  = trocco_notification_destination.slack_notify_when.id
      notification_type = "job"
      notify_when       = "schema_evolution_detected"
      message           = "A schema change was detected"
    },
    {
      destination_type  = "slack"
      slack_channel_id  = trocco_notification_destination.slack_notify_when.id
      notification_type = "job"
      notify_when       = "quality_check_failed"
      message           = "A quality check was violated"
    }
  ]
}
