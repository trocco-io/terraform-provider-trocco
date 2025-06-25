resource "trocco_pipeline_definition" "trocco_transfer_with_custom_variable_loop_with_string_config" {
  name = "trocco_transfer_with_custom_variable_loop_with_string_config"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1

        custom_variable_loop = {
          type = "string"
          string_config = {
            variables = [
              {
                name   = "$foo$"
                values = ["a", "b"]
              },
              {
                name = "$bar$"
                values : ["", "c"]
              }
            ]
          }
        }
      }
    }
  ]
}
