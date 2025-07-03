resource "trocco_pipeline_definition" "notifications" {
  name = "notifications"

  notifications = [
    {
      type             = "job_execution"
      destination_type = "slack"
      notify_when      = "finished"

      slack_config = {
        notification_id = 1
        message         = "The quick brown fox jumps over the lazy dog."
      }
    },
    {
      type             = "job_time_alert"
      destination_type = "email"
      time             = 5

      email_config = {
        notification_id = 1
        message         = "The quick brown fox jumps over the lazy dog."
      }
    },
  ]
}
