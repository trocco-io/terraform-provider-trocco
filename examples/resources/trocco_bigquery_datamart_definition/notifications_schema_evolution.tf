resource "trocco_bigquery_datamart_definition" "with_schema_evolution_notifications" {
  name                     = "example_with_schema_evolution_notifications"
  is_runnable_concurrently = false
  bigquery_connection_id   = 1
  query                    = "SELECT * FROM tables"
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
      slack_channel_id  = 1
      notification_type = "job"
      notify_when       = "schema_evolution_detected"
      message           = "@here A schema change was detected."
    }
  ]
}
