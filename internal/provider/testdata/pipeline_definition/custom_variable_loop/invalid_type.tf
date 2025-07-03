resource "trocco_pipeline_definition" "custom_variable_loop_invalid_type" {
  name = "custom_variable_loop_invalid_type"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
        custom_variable_loop = {
          type = "invalid_type" # Invalid type value
          string_config = {
            variables = [
              {
                name   = "$foo$"
                values = ["a", "b"]
              }
            ]
          }
        }
      }
    }
  ]
}
