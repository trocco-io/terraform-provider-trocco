package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	notification_parameter "terraform-provider-trocco/internal/client/parameter/notification_destination"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetNotificationDestinationEmail(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/notification_destinations/email/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		resp := NotificationDestination{
			ID:    1,
			Email: lo.ToPtr("test@example.com"),
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	result, err := c.GetNotificationDestination("email", 1)

	require.NoError(t, err)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "test@example.com", *result.Email)
}

func TestCreateNotificationDestinationEmail(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/notification_destinations/email", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		c := NotificationDestination{
			ID:    8,
			Email: lo.ToPtr("test@example.com"),
		}
		if err := json.NewEncoder(w).Encode(c); err != nil {
			panic(err)
		}
	}))

	c := NewDevTroccoClient("1234567890", s.URL)

	out, err := c.CreateNotificationDestination("email", &CreateNotificationDestinationInput{
		EmailConfig: &notification_parameter.EmailConfigInput{
			Email: lo.ToPtr("test@example.com"),
		},
	})

	require.NoError(t, err)

	assert.Equal(t, int64(8), out.ID)
	assert.Equal(t, "test@example.com", *out.Email)
}

func TestUpdateNotificationDestinationEmail(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, "/api/notification_destinations/email/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		resp := NotificationDestination{
			ID:    1,
			Email: lo.ToPtr("updated@example.com"),
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	result, err := c.UpdateNotificationDestination("email", 1, &UpdateNotificationDestinationInput{
		EmailConfig: &notification_parameter.EmailConfigInput{
			Email: lo.ToPtr("test@example.com"),
		},
	})

	require.NoError(t, err)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "updated@example.com", *result.Email)
}

func TestDeleteNotificationDestinationEmail(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/api/notification_destinations/email/1", r.URL.Path)
	}))
	defer s.Close()

	c := NewDevTroccoClient("dummy-token", s.URL)
	err := c.DeleteNotificationDestination("email", 1)
	assert.NoError(t, err)
}

func TestGetNotificationDestinationSlackChannel(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/notification_destinations/slack_channel/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		resp := NotificationDestination{
			ID:      1,
			Channel: lo.ToPtr("general"),
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	result, err := c.GetNotificationDestination("slack_channel", 1)

	require.NoError(t, err)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "general", *result.Channel)
}

func TestCreateNotificationDestinationSlackChannel(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/notification_destinations/slack_channel", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		c := NotificationDestination{
			ID:      8,
			Channel: lo.ToPtr("general"),
		}
		if err := json.NewEncoder(w).Encode(c); err != nil {
			panic(err)
		}
	}))

	c := NewDevTroccoClient("1234567890", s.URL)

	out, err := c.CreateNotificationDestination("slack_channel", &CreateNotificationDestinationInput{
		SlackChannelConfig: &notification_parameter.SlackChannelConfigInput{
			Channel:    lo.ToPtr("general"),
			WebhookURL: lo.ToPtr("https://slack-webhook-url.com"),
		},
	})

	require.NoError(t, err)
	assert.Equal(t, int64(8), out.ID)
	assert.Equal(t, "general", *out.Channel)
}

func TestUpdateNotificationDestinationSlackChannel(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, "/api/notification_destinations/slack_channel/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		resp := NotificationDestination{
			ID:      1,
			Channel: lo.ToPtr("updated-channel"),
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	result, err := c.UpdateNotificationDestination("slack_channel", 1, &UpdateNotificationDestinationInput{
		SlackChannelConfig: &notification_parameter.SlackChannelConfigInput{
			Channel:    lo.ToPtr("general"),
			WebhookURL: lo.ToPtr("https://slack-webhook-url.com"),
		},
	})

	require.NoError(t, err)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "updated-channel", *result.Channel)
}

func TestDeleteNotificationDestinationSlackChannel(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/api/notification_destinations/slack_channel/1", r.URL.Path)
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)
	err := c.DeleteNotificationDestination("slack_channel", 1)
	assert.NoError(t, err)
}
