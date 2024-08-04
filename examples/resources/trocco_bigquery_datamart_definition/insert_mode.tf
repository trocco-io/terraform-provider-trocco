resource "trocco_bigquery_datamart_definition" "insert_mode" {
  name                     = "example_insert_mode"
  is_runnable_concurrently = false
  bigquery_connection_id   = 1
  query                    = "SELECT * FROM tables"
  query_mode               = "insert"
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "append"
  before_load              = "DELETE FROM tables WHERE created_at < '2024-01-01'"
  partitioning             = "time_unit_column"
  partitioning_time        = "DAY"
  partitioning_field       = "created_at"
  clustering_fields        = ["id", "name"]
}
