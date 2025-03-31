package notification_destination

type SlackChannelConfigInput struct {
	Channel    *string `json:"channel,omitempty"`
	WebhookURL *string `json:"webhook_url,omitempty"`
}
