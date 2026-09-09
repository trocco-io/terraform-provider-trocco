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
    {
      type             = "job_execution"
      destination_type = "http"
      notify_when      = "failed"

      http_config = {
        notification_id = 1                                      # ID of a trocco_notification_destination with type = "http"
        message         = "{\"text\": \"The workflow failed.\"}" # sent as the request body; must be a JSON string
      }
    },
  ]
}
