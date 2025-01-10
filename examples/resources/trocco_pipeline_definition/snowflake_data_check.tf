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

resource "trocco_pipeline_definition" "snowflake_data_check_with_custom_variables" {
  name = "snowflake_data_check_with_custom_variables"

  tasks = [
    {
      key  = "snowflake_data_check_with_custom_variables"
      type = "snowflake_data_check"

      snowflake_data_check_config = {
        name          = "Example"
        connection_id = 1
        query         = "SELECT COUNT(id) FROM examples"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false

        custom_variables = [
          {
            name  = "$string$"
            type  = "string"
            value = "foo"
          },
          {
            name      = "$timestamp$"
            type      = "timestamp"
            quantity  = 1,
            unit      = "hour"
            direction = "ago"
            format    = "%Y-%m-%d %H:%M:%S"
            time_zone = "Asia/Tokyo"
          },
        ]
      }
    }
  ]
}
