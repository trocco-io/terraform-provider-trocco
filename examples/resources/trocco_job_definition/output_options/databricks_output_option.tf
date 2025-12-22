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
    }
  }
}