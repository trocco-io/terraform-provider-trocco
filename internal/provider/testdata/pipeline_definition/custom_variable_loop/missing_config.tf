resource "trocco_pipeline_definition" "custom_variable_loop_missing_config" {
  name = "custom_variable_loop_missing_config"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
        custom_variable_loop = {
          type = "string"
          # string_config is missing
        }
      }
    }
  ]
}

resource "trocco_pipeline_definition" "custom_variable_loop_wrong_config" {
  name = "custom_variable_loop_wrong_config"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
        custom_variable_loop = {
          type = "string"
          # Using bigquery_config instead of string_config
          bigquery_config = {
            connection_id = 1
            query         = "SELECT foo, bar FROM sample"
            variables = [
              "$foo$",
              "$bar$"
            ]
          }
        }
      }
    }
  ]
}
