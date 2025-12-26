resource "trocco_pipeline_definition" "if_else_test" {
  name        = "if_else_test"
  description = "Test pipeline for if-else task"

  tasks = [
    {
      key  = "transfer"
      type = "trocco_transfer"
      trocco_transfer_config = {
        definition_id = 1
      }
    },
    {
      key  = "success_task"
      type = "trocco_transfer"
      trocco_transfer_config = {
        definition_id = 1
      }
    },
    {
      key  = "failure_task"
      type = "trocco_transfer"
      trocco_transfer_config = {
        definition_id = 1
      }
    },
    {
      key  = "if_else"
      type = "if_else"
      if_else_config = {
        name = "Check transfer status"
        condition_groups = {
          set_type = "and"
          conditions = [
            {
              variable = "status"
              task_key = "transfer"
              operator = "equal"
              value    = "succeeded"
            }
          ]
        }
        destinations = {
          if   = ["success_task"]
          else = ["failure_task"]
        }
      }
    }
  ]

  task_dependencies = [
    {
      source      = "transfer"
      destination = "if_else"
    },
    {
      source      = "if_else"
      destination = "success_task"
    },
    {
      source      = "if_else"
      destination = "failure_task"
    }
  ]
}
