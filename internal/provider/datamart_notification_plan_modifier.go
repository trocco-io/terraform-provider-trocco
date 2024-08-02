package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &datamartNotificationPlanModifier{}

type datamartNotificationPlanModifier struct{}

func (d *datamartNotificationPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating notification attributes"
}

func (d *datamartNotificationPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *datamartNotificationPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var destination_type types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("destination_type"), &destination_type)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var slackChannelID types.Int64
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("slack_channel_id"), &slackChannelID)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var emailID types.Int64
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("email_id"), &emailID)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var notificationType types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("notification_type"), &notificationType)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var notifyWhen types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("notify_when"), &notifyWhen)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var recordCount types.Int64
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("record_count"), &recordCount)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var recordOperator types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("record_operator"), &recordOperator)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if destination_type.ValueString() == "slack" && slackChannelID.IsNull() {
		addNotificationAttributeError(req, resp, "slack_channel_id is required for slack destination type")
	}

	if destination_type.ValueString() == "email" && emailID.IsNull() {
		addNotificationAttributeError(req, resp, "email_id is required for email destination type")
	}

	if notificationType.ValueString() == "job" && notifyWhen.IsNull() {
		addNotificationAttributeError(req, resp, "notify_when is required for job notification type")
	}

	if notificationType.ValueString() == "record_count" && recordCount.IsNull() {
		addNotificationAttributeError(req, resp, "record_count is required for record_count notification type")
	}

	if notificationType.ValueString() == "record_count" && recordOperator.IsNull() {
		addNotificationAttributeError(req, resp, "record_operator is required for record_count notification type")
	}
}

func addNotificationAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Notification Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
