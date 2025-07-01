resource "trocco_notification_destination" "slack_channel_test" {
  type = "slack_channel"
  slack_channel_config = {
    channel     = "trocco-log2"
    webhook_url = "https://hooks.slack.com/services/test"
  }
}