package pipeline_definition

import (
	"context"
	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
	"terraform-provider-trocco/internal/provider/custom_type"
	model "terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

//
// Notification
//

type Notification struct {
	Type            types.String             `tfsdk:"type"`
	DestinationType types.String             `tfsdk:"destination_type"`
	NotifyWhen      types.String             `tfsdk:"notify_when"`
	Time            types.Int64              `tfsdk:"time"`
	EmailConfig     *EmailNotificationConfig `tfsdk:"email_config"`
	SlackConfig     *SlackNotificationConfig `tfsdk:"slack_config"`
}

func NewNotifications(ens []*we.Notification, previousIsNull bool) types.Set {
	ctx := context.Background()
	objectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"type":             types.StringType,
			"destination_type": types.StringType,
			"notify_when":      types.StringType,
			"time":             types.Int64Type,
			"email_config": types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"notification_id": types.Int64Type,
					"message":         types.StringType,
				},
			},
			"slack_config": types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"notification_id": types.Int64Type,
					"message":         types.StringType,
				},
			},
		},
	}

	if ens == nil {
		return types.SetNull(objectType)
	}

	if previousIsNull && len(ens) == 0 {
		return types.SetNull(objectType)
	}

	mds := []*Notification{}
	for _, en := range ens {
		mds = append(mds, NewNotification(en))
	}

	setValue, diags := types.SetValueFrom(ctx, objectType, mds)
	if diags.HasError() {
		return types.SetNull(objectType)
	}

	return setValue
}

func NewNotification(en *we.Notification) *Notification {
	return &Notification{
		Type:            types.StringValue(en.Type),
		DestinationType: types.StringValue(en.DestinationType),
		NotifyWhen:      types.StringPointerValue(en.NotifyWhen),
		Time:            types.Int64PointerValue(en.Time),
		EmailConfig:     NewEmailNotificationConfig(en.EmailConfig),
		SlackConfig:     NewSlackNotificationConfig(en.SlackConfig),
	}
}

func (n *Notification) ToInput() *wp.Notification {
	param := &wp.Notification{
		Type:            n.Type.ValueString(),
		DestinationType: n.DestinationType.ValueString(),
		NotifyWhen:      model.NewNullableString(n.NotifyWhen),
		Time:            model.NewNullableInt64(n.Time),
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
	NotificationID types.Int64                    `tfsdk:"notification_id"`
	Message        custom_type.TrimmedStringValue `tfsdk:"message"`
}

func NewEmailNotificationConfig(en *we.EmailNotificationConfig) *EmailNotificationConfig {
	if en == nil {
		return nil
	}

	return &EmailNotificationConfig{
		NotificationID: types.Int64Value(en.NotificationID),
		Message:        custom_type.TrimmedStringValue{StringValue: types.StringValue(en.Message)},
	}
}

func (c *EmailNotificationConfig) ToInput() *wp.EmailNotificationConfig {
	return &wp.EmailNotificationConfig{
		NotificationID: c.NotificationID.ValueInt64(),
		Message:        c.Message.ValueString(),
	}
}

//
// SlackNotificationConfig
//

type SlackNotificationConfig struct {
	NotificationID types.Int64                    `tfsdk:"notification_id"`
	Message        custom_type.TrimmedStringValue `tfsdk:"message"`
}

func NewSlackNotificationConfig(en *we.SlackNotificationConfig) *SlackNotificationConfig {
	if en == nil {
		return nil
	}

	return &SlackNotificationConfig{
		NotificationID: types.Int64Value(en.NotificationID),
		Message:        custom_type.TrimmedStringValue{StringValue: types.StringValue(en.Message)},
	}
}

func (c *SlackNotificationConfig) ToInput() *wp.SlackNotificationConfig {
	return &wp.SlackNotificationConfig{
		NotificationID: c.NotificationID.ValueInt64(),
		Message:        c.Message.ValueString(),
	}
}
