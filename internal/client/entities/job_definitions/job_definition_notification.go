package job_definitions

type JobDefinitionNotification struct {
	DestinationType  string  `json:"destination_type"`
	SlackChannelID   *int64  `json:"slack_channel_id"`
	EmailID          *int64  `json:"email_id"`
	NotificationType string  `json:"notification_type"`
	NotifyWhen       *string `json:"notify_when"`
	Message          string  `json:"message"`
	RecordCount      *int64  `json:"record_count"`
	RecordOperator   *string `json:"record_operator"`
	RecordType       *string `json:"record_type"`
	Minutes          *int64  `json:"minutes"`
}
