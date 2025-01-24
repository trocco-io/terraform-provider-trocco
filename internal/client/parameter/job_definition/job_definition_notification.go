package job_definitions

type JobDefinitionNotificationInput struct {
	DestinationType  string  `json:"destination_type"`
	SlackChannelID   *int64  `json:"slack_channel_id,omitempty"`
	EmailID          *int64  `json:"email_id,omitempty"`
	NotificationType string  `json:"notification_type"`
	NotifyWhen       *string `json:"notify_when,omitempty"`
	Message          string  `json:"message"`
	RecordCount      *int64  `json:"record_count,omitempty"`
	RecordOperator   *string `json:"record_operator,omitempty"`
	RecordType       *string `json:"record_type,omitempty"`
	Minutes          *int64  `json:"minutes,omitempty"`
}
