package pipeline_definition

import "terraform-provider-trocco/internal/client/parameter"

type Notification struct {
	Type            string                    `json:"type"`
	DestinationType string                    `json:"destination_type"`
	NotifyWhen      *parameter.NullableString `json:"notify_when"`
	Time            *parameter.NullableInt64  `json:"time"`
	EmailConfig     *EmailNotificationConfig  `json:"email_config"`
	SlackConfig     *SlackNotificationConfig  `json:"slack_config"`
}

type EmailNotificationConfig struct {
	NotificationID int64  `json:"notification_id"`
	Message        string `json:"message"`
}

type SlackNotificationConfig struct {
	NotificationID int64  `json:"notification_id"`
	Message        string `json:"message"`
}
