resource "trocco_job_definition" "databricks_to_bigquery" {
  input_option_type = "databricks"

  input_option = {
    databricks_input_option = {
      databricks_connection_id = 1
      catalog_name             = "catalog_example"
      schema_name              = "schema_example"
      query                    = "select * from example_table"
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
}
