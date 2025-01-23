resource "trocco_pipeline_definition" "task_dependencies" {
  name = "task_dependencies"

  tasks = [
    {
      key  = "trocco_transfer_first"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
      }
    },
    {
      key  = "trocco_transfer_second"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 2
      }
    },
  ]


  task_dependencies = [
    {
      source      = "trocco_transfer_first"
      destination = "trocco_transfer_second"
    },
  ]
}
