resource "trocco_connection" "my_snowflake_conn" {
  connection_type = "snowflake"
  name = "Snowflake Test"
  host = "example.snowflakecomputing.com"
  auth_method = "user_password"
  user_name = "test_user"
  password = "test_password"
}


resource "trocco_pipeline_definition" "snowflake_data_check_query_check" {
  name = "snowflake_data_check"

  tasks = [
    {
      key  = "snowflake_data_check"
      type = "snowflake_data_check"

      snowflake_data_check_config = {
        name          = "Example"
        connection_id = trocco_connection.my_snowflake_conn.id
        query         = <<SQL
          SELECT COUNT(*) FROM examples
        SQL
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
        warehouse = "COMPUTE_WH"
      }
    }
  ]
}
