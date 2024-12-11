package workflow

import (
	we "terraform-provider-trocco/internal/client/entities/workflow"
	wp "terraform-provider-trocco/internal/client/parameters/workflow"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

//
// Notification
//

type Notification struct {
	Type types.String `tfsdk:"type"`

	EmailConfig *EmailNotificationConfig `tfsdk:"email_config"`
	SlackConfig *SlackNotificationConfig `tfsdk:"slack_config"`
}

func NewNotification(en *we.Notification) *Notification {
	if en == nil {
		return nil
	}

	return &Notification{
		Type:        types.StringValue(en.Type),
		EmailConfig: NewEmailNotificationConfig(en.EmailConfig),
		SlackConfig: NewSlackNotificationConfig(en.SlackConfig),
	}
}

func (n *Notification) ToInput() wp.Notification {
	param := wp.Notification{
		Type: n.Type.ValueString(),
	}

	if n.EmailConfig != nil {
		param.EmailConfig = n.EmailConfig.ToInput()
	}
	if n.SlackConfig != nil {
		param.SlackConfig = n.SlackConfig.ToInput()
	}

	return param
}

//
// EmailNotificationConfig
//

type EmailNotificationConfig struct {
	NotificationID types.Int64  `tfsdk:"notification_id"`
	NotifyWhen     types.String `tfsdk:"notify_when"`
	Message        types.String `tfsdk:"message"`
}

func NewEmailNotificationConfig(en *we.EmailNotificationConfig) *EmailNotificationConfig {
	if en == nil {
		return nil
	}

	return &EmailNotificationConfig{
		NotificationID: types.Int64Value(en.NotificationID),
		NotifyWhen:     types.StringValue(en.NotifyWhen),
		Message:        types.StringValue(en.Message),
	}
}

func (c *EmailNotificationConfig) ToInput() *wp.EmailNotificationConfig {
	return &wp.EmailNotificationConfig{
		NotificationID: c.NotificationID.ValueInt64(),
		NotifyWhen:     c.NotifyWhen.ValueString(),
		Message:        c.Message.ValueString(),
	}
}

//
// SlackNotificationConfig
//

type SlackNotificationConfig struct {
	NotificationID types.Int64  `tfsdk:"notification_id"`
	NotifyWhen     types.String `tfsdk:"notify_when"`
	Message        types.String `tfsdk:"message"`
}

func NewSlackNotificationConfig(en *we.SlackNotificationConfig) *SlackNotificationConfig {
	if en == nil {
		return nil
	}

	return &SlackNotificationConfig{
		NotificationID: types.Int64Value(en.NotificationID),
		NotifyWhen:     types.StringValue(en.NotifyWhen),
		Message:        types.StringValue(en.Message),
	}
}

func (c *SlackNotificationConfig) ToInput() *wp.SlackNotificationConfig {
	return &wp.SlackNotificationConfig{
		NotificationID: c.NotificationID.ValueInt64(),
		NotifyWhen:     c.NotifyWhen.ValueString(),
		Message:        c.Message.ValueString(),
	}
}
