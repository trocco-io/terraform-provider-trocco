package client

import (
	"fmt"
	"net/http"
	"terraform-provider-trocco/internal/client/parameter/notification_destination"
)

type NotificationDestination struct {
	// Common Fields
	ID int64 `json:"id"`

	// Email Fields
	Email *string `json:"email"`

	// SlackChannel Fields
	Channel *string `json:"channel"`

	// HTTP Fields
	Name        *string        `json:"name"`
	URL         *string        `json:"url"`
	Description *string        `json:"description"`
	Headers     []HTTPKeyValue `json:"headers"`
	QueryParams []HTTPKeyValue `json:"query_params"`
}

// HTTPKeyValue is a header or a query parameter of an HTTP notification destination.
// The API returns an empty Value when Masking is true.
type HTTPKeyValue struct {
	ID      *int64 `json:"id"`
	Key     string `json:"key"`
	Value   string `json:"value"`
	Masking bool   `json:"masking"`
}

type CreateNotificationDestinationInput struct {
	EmailConfig        *notification_destination.EmailConfigInput        `json:"email_config,omitempty"`
	SlackChannelConfig *notification_destination.SlackChannelConfigInput `json:"slack_channel_config,omitempty"`
	HTTPConfig         *notification_destination.HTTPConfigInput         `json:"http_config,omitempty"`
}

type UpdateNotificationDestinationInput struct {
	EmailConfig        *notification_destination.EmailConfigInput        `json:"email_config,omitempty"`
	SlackChannelConfig *notification_destination.SlackChannelConfigInput `json:"slack_channel_config,omitempty"`
	HTTPConfig         *notification_destination.HTTPConfigInput         `json:"http_config,omitempty"`
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
