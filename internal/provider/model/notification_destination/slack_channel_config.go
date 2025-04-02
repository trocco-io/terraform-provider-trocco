package notification_destination

import "github.com/hashicorp/terraform-plugin-framework/types"

type SlackChannelConfig struct {
	Channel    types.String `tfsdk:"channel"`
	WebhookURL types.String `tfsdk:"webhook_url"`
}
