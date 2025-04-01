resource "trocco_notification_destination" "slack_channel" {
  type = "slack_channel"

  slack_channel_config = {
    channel     = "#general"
    webhook_url = "https://hooks.slack.com/services/XXXX/YYYY/ZZZZ"
  }
}
