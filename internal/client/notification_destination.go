package client

import (
	"fmt"
	"net/http"
	"terraform-provider-trocco/internal/client/parameter"
)

type NotificationDestination struct {
	// Common Fields
	ID int64 `json:"id"`

	// Email Fields
	Email *string `json:"email"`

	// SlackChannel Fields
	Channel    *string `json:"channel"`
	WebhookURL *string `json:"webhook_url"`
}

type CreateNotificationDestinationInput struct {
	EmailConfig        *EmailConfigInput        `json:"email_config,omitempty"`
	SlackChannelConfig *SlackChannelConfigInput `json:"slack_channel_config,omitempty"`
}

type UpdateNotificationDestinationInput struct {
	EmailConfig        *EmailConfigInput        `json:"email_config,omitempty"`
	SlackChannelConfig *SlackChannelConfigInput `json:"slack_channel_config,omitempty"`
}

type EmailConfigInput struct {
	Email *parameter.NullableString `json:"email,omitempty"`
}

type SlackChannelConfigInput struct {
	Channel    *parameter.NullableString `json:"channel,omitempty"`
	WebhookURL *parameter.NullableString `json:"webhook_url,omitempty"`
}

func (c *TroccoClient) CreateNotificationDestination(notificationType string, in *CreateNotificationDestinationInput) (*NotificationDestination, error) {
	out := &NotificationDestination{}
	if err := c.do(
		http.MethodPost,
		fmt.Sprintf("/api/notification_destinations/%s", notificationType),
		in,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) UpdateNotificationDestination(notificationType string, id int64, in *UpdateNotificationDestinationInput) (*NotificationDestination, error) {
	out := &NotificationDestination{}
	if err := c.do(
		http.MethodPatch,
		fmt.Sprintf("/api/notification_destinations/%s/%d", notificationType, id),
		in,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) GetNotificationDestination(notificationType string, id int64) (*NotificationDestination, error) {
	out := &NotificationDestination{}
	if err := c.do(
		http.MethodGet,
		fmt.Sprintf("/api/notification_destinations/%s/%d", notificationType, id),
		nil,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) DeleteNotificationDestination(notificationType string, id int64) error {
	return c.do(
		http.MethodDelete,
		fmt.Sprintf("/api/notification_destinations/%s/%d", notificationType, id),
		nil,
		nil,
	)
}
