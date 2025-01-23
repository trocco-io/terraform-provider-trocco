resource "trocco_pipeline_definition" "redshift_data_check" {
  name = "redshift_data_check"

  tasks = [
    {
      key  = "redshift_data_check"
      type = "redshift_data_check"

      redshift_data_check_config = {
        name          = "Example"
        connection_id = 1
        query         = "SELECT COUNT(id) FROM examples"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
        database      = "EXAMPLE"
      }
    }
  ]
}

resource "trocco_pipeline_definition" "redshift_data_check_with_custom_variables" {
  name = "redshift_data_check_with_custom_variables"

  tasks = [
    {
      key  = "redshift_data_check_with_custom_variables"
      type = "redshift_data_check"

      redshift_data_check_config = {
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
