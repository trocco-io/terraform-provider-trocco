resource "trocco_bigquery_datamart_definition" "with_quality_checks" {
  name                     = "example_with_quality_checks"
  is_runnable_concurrently = false
  bigquery_connection_id   = 1
  query                    = "SELECT * FROM tables"
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
      slack_channel_id  = 1
      notification_type = "job"
      notify_when       = "quality_check_failed"
      message           = "@here A quality check was violated."
    }
  ]
}
