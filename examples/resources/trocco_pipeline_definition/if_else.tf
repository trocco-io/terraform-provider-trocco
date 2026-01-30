# Example: if-else task to branch workflow based on task status
resource "trocco_pipeline_definition" "if_else" {
  name        = "if_else_example"
  description = "Pipeline with if-else branching based on task status"

  tasks = [
    {
      key  = "transfer"
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
    },
    {
      key  = "success_task"
      type = "trocco_bigquery_datamart"

      trocco_bigquery_datamart_config = {
        definition_id = 2
      }
    },
    {
      key  = "failure_task"
      type = "slack_notify"

      slack_notification_config = {
        name          = "Notify failure"
        connection_id = 1
        message       = "Transfer task failed."
        ignore_error  = false
      }
    },
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
    },
  ]
}
