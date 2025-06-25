resource "trocco_pipeline_definition" "trocco_transfer_with_custom_variable_loop_with_period_config" {
  name = "trocco_transfer_with_custom_variable_loop_with_period_config"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
        custom_variable_loop = {
          type = "period"
          period_config = {
            interval  = "day"
            time_zone = "Asia/Tokyo"
            from = {
              value = 7
              unit  = "day"
            }
            to = {
              value = 1
              unit  = "day"
            }
            variables = [
              {
                name = "$date$"
                offset = {
                  value = 0
                  unit  = "day"
                }
              }
            ]
          }
        }
      }
    }
  ]
}
