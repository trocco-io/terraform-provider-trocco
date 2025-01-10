package job_definitions

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	job_definitions2 "terraform-provider-trocco/internal/client/parameters/job_definitions"
)

type JobDefinitionNotification struct {
	DestinationType  types.String `tfsdk:"destination_type"`
	SlackChannelID   types.Int64  `tfsdk:"slack_channel_id"`
	EmailID          types.Int64  `tfsdk:"email_id"`
	NotificationType types.String `tfsdk:"notification_type"`
	NotifyWhen       types.String `tfsdk:"notify_when"`
	Message          types.String `tfsdk:"message"`
	RecordCount      types.Int64  `tfsdk:"record_count"`
	RecordOperator   types.String `tfsdk:"record_operator"`
	RecordType       types.String `tfsdk:"record_type"`
	Minutes          types.Int64  `tfsdk:"minutes"`
}

func NewJobDefinitionNotifications(jobDefinitionNotifications *[]job_definitions.JobDefinitionNotification) []JobDefinitionNotification {
	if jobDefinitionNotifications == nil {
		return nil
	}
	notifications := make([]JobDefinitionNotification, 0, len(*jobDefinitionNotifications))
	for _, input := range *jobDefinitionNotifications {
		notification := JobDefinitionNotification{
			DestinationType:  types.StringValue(input.DestinationType),
			SlackChannelID:   types.Int64PointerValue(input.SlackChannelID),
			EmailID:          types.Int64PointerValue(input.EmailID),
			NotificationType: types.StringValue(input.NotificationType),
			NotifyWhen:       types.StringPointerValue(input.NotifyWhen),
			Message:          types.StringValue(input.Message),
			RecordCount:      types.Int64PointerValue(input.RecordCount),
			RecordOperator:   types.StringPointerValue(input.RecordOperator),
			RecordType:       types.StringPointerValue(input.RecordType),
			Minutes:          types.Int64PointerValue(input.Minutes),
		}
		notifications = append(notifications, notification)
	}
	return notifications
}

func (notification JobDefinitionNotification) ToInput() job_definitions2.JobDefinitionNotificationInput {
	input := job_definitions2.JobDefinitionNotificationInput{
		DestinationType:  notification.DestinationType.ValueString(),
		SlackChannelID:   notification.SlackChannelID.ValueInt64Pointer(),
		EmailID:          notification.EmailID.ValueInt64Pointer(),
		NotificationType: notification.NotificationType.ValueString(),
		NotifyWhen:       notification.NotifyWhen.ValueStringPointer(),
		Message:          notification.Message.ValueString(),
		RecordCount:      notification.RecordCount.ValueInt64Pointer(),
		RecordOperator:   notification.RecordOperator.ValueStringPointer(),
		RecordType:       notification.RecordType.ValueStringPointer(),
		Minutes:          notification.Minutes.ValueInt64Pointer(),
	}
	return input
}
