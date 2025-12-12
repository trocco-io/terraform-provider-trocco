resource "trocco_job_definition" "databricks_output_example" {
  output_option_type = "databricks"

  output_option = {
    databricks_output_option = {
      databricks_connection_id = 1
      catalog_name             = "catalog_example"
      schema_name              = "schema_example"
      table                    = "table_example"
      batch_size               = 5000
      mode                     = "insert"
      default_time_zone        = "Etc/UTC"
      databricks_output_option_column_options = [
        {
          name             = "timestamp_column"
          type             = "TIMESTAMP"
          value_type       = "timestamp"
          timestamp_format = "%Y-%m-%d %H:%M:%S"
          timezone         = "Asia/Tokyo"
        }
      ]
      databricks_output_option_merge_keys = [
        "id"
      ]
    }
  }
}