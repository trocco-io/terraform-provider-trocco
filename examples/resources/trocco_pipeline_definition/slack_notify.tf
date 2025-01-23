resource "trocco_pipeline_definition" "slack_notify" {
  name = "slack_notify"

  tasks = [
    {
      key  = "slack_notify"
      type = "slack_notify"

      slack_notification_config = {
        name          = "Example"
        connection_id = 1
        message       = "The quick brown fox jumps over the lazy dog."
        ignore_error  = false
      }
    }
  ]
}
