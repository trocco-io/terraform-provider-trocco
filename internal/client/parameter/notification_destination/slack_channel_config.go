package notification_destination

import "terraform-provider-trocco/internal/client/parameter"

type SlackChannelConfigInput struct {
	Channel    *parameter.NullableString `json:"channel,omitempty"`
	WebhookURL *parameter.NullableString `json:"webhook_url,omitempty"`
}
