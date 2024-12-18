package job_definitions

import "github.com/hashicorp/terraform-plugin-framework/types"

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
