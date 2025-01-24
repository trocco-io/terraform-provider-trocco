resource "trocco_job_definition" "bigquery_output_example" {
  output_option_type = "bigquery"
  output_option = {
    bigquery_output_option = {
      dataset                         = "test_dataset"
      table                           = "test_table"
      mode                            = "merge"
      auto_create_dataset             = true
      auto_create_table               = false
      timeout_sec                     = 300
      open_timeout_sec                = 300
      read_timeout_sec                = 300
      send_timeout_sec                = 300
      retries                         = 0
      bigquery_connection_id          = 1 # require your bigquery connection id
      partitioning_type               = "time_unit_column"
      time_partitioning_type          = "DAY"
      time_partitioning_field         = "created_at"
      time_partitioning_expiration_ms = 10000
      location                        = "US"
      bigquery_output_option_merge_keys = [
        "id"
      ]
    }
  }
}
