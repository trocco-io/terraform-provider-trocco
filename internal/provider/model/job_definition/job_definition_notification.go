package job_definitions

import (
	jobDefinitionEntities "terraform-provider-trocco/internal/client/entity/job_definition"
	jobDefinitionParameters "terraform-provider-trocco/internal/client/parameter/job_definition"
	"terraform-provider-trocco/internal/provider/custom_type"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type JobDefinitionNotification struct {
	DestinationType  types.String                   `tfsdk:"destination_type"`
	SlackChannelID   types.Int64                    `tfsdk:"slack_channel_id"`
	EmailID          types.Int64                    `tfsdk:"email_id"`
	NotificationType types.String                   `tfsdk:"notification_type"`
	NotifyWhen       types.String                   `tfsdk:"notify_when"`
	Message          custom_type.TrimmedStringValue `tfsdk:"message"`
	RecordCount      types.Int64                    `tfsdk:"record_count"`
	RecordOperator   types.String                   `tfsdk:"record_operator"`
	RecordType       types.String                   `tfsdk:"record_type"`
	Minutes          types.Int64                    `tfsdk:"minutes"`
}

func NewJobDefinitionNotifications(jobDefinitionNotifications []jobDefinitionEntities.JobDefinitionNotification) []JobDefinitionNotification {
	if jobDefinitionNotifications == nil {
		return nil
	}
	notifications := make([]JobDefinitionNotification, 0, len(jobDefinitionNotifications))
	for _, input := range jobDefinitionNotifications {
		notification := JobDefinitionNotification{
			DestinationType:  types.StringValue(input.DestinationType),
			SlackChannelID:   types.Int64PointerValue(input.SlackChannelID),
			EmailID:          types.Int64PointerValue(input.EmailID),
			NotificationType: types.StringValue(input.NotificationType),
			NotifyWhen:       types.StringPointerValue(input.NotifyWhen),
			Message:          custom_type.TrimmedStringValue{StringValue: types.StringValue(input.Message)},
			RecordCount:      types.Int64PointerValue(input.RecordCount),
			RecordOperator:   types.StringPointerValue(input.RecordOperator),
			RecordType:       types.StringPointerValue(input.RecordType),
			Minutes:          types.Int64PointerValue(input.Minutes),
		}
		notifications = append(notifications, notification)
	}
	return notifications
}

func (notification JobDefinitionNotification) ToInput() jobDefinitionParameters.JobDefinitionNotificationInput {
	input := jobDefinitionParameters.JobDefinitionNotificationInput{
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

func (n JobDefinitionNotification) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"destination_type":  types.StringType,
		"slack_channel_id":  types.Int64Type,
		"email_id":          types.Int64Type,
		"notification_type": types.StringType,
		"notify_when":       types.StringType,
		"message":           types.StringType,
		"record_count":      types.Int64Type,
		"record_operator":   types.StringType,
		"record_type":       types.StringType,
		"minutes":           types.Int64Type,
	}
}
