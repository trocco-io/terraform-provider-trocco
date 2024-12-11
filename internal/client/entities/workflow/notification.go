package workflow

type Notification struct {
	Type string `json:"type"`

	EmailConfig *EmailNotificationConfig `json:"email_config"`
	SlackConfig *SlackNotificationConfig `json:"slack_config"`
}

type EmailNotificationConfig struct {
	NotificationID int64  `json:"notification_id"`
	NotifyWhen     string `json:"notify_when"`
	Message        string `json:"message"`
}

type SlackNotificationConfig struct {
	NotificationID int64  `json:"notification_id"`
	NotifyWhen     string `json:"notify_when"`
	Message        string `json:"message"`
}
