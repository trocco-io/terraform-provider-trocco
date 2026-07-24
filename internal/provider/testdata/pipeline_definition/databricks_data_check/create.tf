resource "trocco_connection" "my_databricks_conn" {
  connection_type = "databricks"

  name                  = "Databricks Test"
  server_hostname       = "example.databricks.com"
  http_path             = "/sql/1.0/warehouses/xxxx-xxxx-xxxx-xxxx"
  auth_type             = "pat"
  personal_access_token = "dapiXXXXXXXXXXXXXXXXXXXX"
}

resource "trocco_pipeline_definition" "databricks_data_check_query_check" {
  name = "databricks_data_check"

  tasks = [
    {
      key  = "databricks_data_check"
      type = "databricks_data_check"

      databricks_data_check_config = {
        name          = "Example"
        connection_id = trocco_connection.my_databricks_conn.id
        query         = <<SQL
          SELECT COUNT(*) FROM examples
        SQL
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
      }
    }
  ]
}
