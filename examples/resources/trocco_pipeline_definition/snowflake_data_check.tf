resource "trocco_pipeline_definition" "snowflake_data_check" {
  name = "snowflake_data_check"

  tasks = [
    {
      key  = "snowflake_data_check"
      type = "snowflake_data_check"

      snowflake_data_check_config = {
        name          = "Example"
        connection_id = 1
        query         = "SELECT COUNT(id) FROM examples"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
        warehouse     = "EXAMPLE"
      }
    }
  ]
}
