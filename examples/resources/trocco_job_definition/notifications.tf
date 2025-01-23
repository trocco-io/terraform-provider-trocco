resource "trocco_job_definition" "notifications" {
  notifications = [
    {
      destination_type  = "email"
      email_id          = 1 # require your email id
      message           = "email failed"
      notification_type = "job"
      notify_when       = "failed"
    },
    {
      destination_type  = "email"
      email_id          = 1 # require your email id
      message           = "email1"
      notification_type = "job"
      notify_when       = "finished"
    },
    {
      destination_type  = "email"
      email_id          = 1 # require your email id
      message           = "record count email skipped"
      notification_type = "record"
      record_count      = 10
      record_operator   = "below"
      record_type       = "skipped"
    },
    {
      destination_type  = "email"
      email_id          = 1 # require your email id
      message           = "time alert email"
      minutes           = 10
      notification_type = "exec_time"
    },
    {
      destination_type  = "slack"
      message           = "record count slack transfer"
      notification_type = "record"
      record_count      = 10
      record_operator   = "below"
      record_type       = "transfer"
      slack_channel_id  = 1 # require your slack id
    },
    {
      destination_type  = "slack"
      message           = "slack 1"
      notification_type = "job"
      notify_when       = "finished"
      slack_channel_id  = 1 # require your slack id
    },
    {
      destination_type  = "slack"
      message           = "slack failed"
      notification_type = "job"
      notify_when       = "failed"
      slack_channel_id  = 1 # require your slack id
    },
    {
      destination_type  = "slack"
      message           = "time alert slack"
      minutes           = 10
      notification_type = "exec_time"
      slack_channel_id  = 1 # require your slack id
    },
  ]
}