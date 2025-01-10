resource "trocco_pipeline_definition" "bigquery_data_check" {
  name = "bigquery_data_check"

  tasks = [
    {
      key  = "bigquery_data_check"
      type = "bigquery_data_check"

      bigquery_data_check_config = {
        name          = "Example"
        connection_id = 1
        query         = "SELECT COUNT(id) FROM examples"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
      }
    }
  ]
}

resource "trocco_pipeline_definition" "bigquery_data_check_with_custom_variables" {
  name = "bigquery_data_check_with_custom_variables"

  tasks = [
    {
      key  = "bigquery_data_check_with_custom_variables"
      type = "bigquery_data_check"

      bigquery_data_check_config = {
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

