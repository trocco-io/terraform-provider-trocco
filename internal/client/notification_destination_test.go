package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	notificationDestinationParameters "terraform-provider-trocco/internal/client/parameter/notification_destination"
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
		EmailConfig: &notificationDestinationParameters.EmailConfigInput{
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
		EmailConfig: &notificationDestinationParameters.EmailConfigInput{
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
		SlackChannelConfig: &notificationDestinationParameters.SlackChannelConfigInput{
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
		SlackChannelConfig: &notificationDestinationParameters.SlackChannelConfigInput{
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

func TestGetNotificationDestinationHTTP(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/notification_destinations/http/3", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		// The API returns an empty value for masked headers and query parameters.
		_, _ = w.Write([]byte(`{"id":3,"type":"http","name":"my-webhook","url":"https://example.com/webhook","description":"desc","headers":[{"id":10,"key":"Authorization","value":"","masking":true},{"id":11,"key":"Content-Type","value":"application/json","masking":false}],"query_params":[{"id":20,"key":"token","value":"","masking":true}]}`))
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	result, err := c.GetNotificationDestination("http", 3)

	require.NoError(t, err)
	assert.Equal(t, int64(3), result.ID)
	assert.Equal(t, "my-webhook", *result.Name)
	assert.Equal(t, "https://example.com/webhook", *result.URL)
	assert.Equal(t, "desc", *result.Description)
	require.Len(t, result.Headers, 2)
	assert.Equal(t, "Authorization", result.Headers[0].Key)
	assert.Empty(t, result.Headers[0].Value)
	assert.True(t, result.Headers[0].Masking)
	assert.Equal(t, "application/json", result.Headers[1].Value)
	assert.False(t, result.Headers[1].Masking)
	require.Len(t, result.QueryParams, 1)
	assert.Equal(t, "token", result.QueryParams[0].Key)
	assert.True(t, result.QueryParams[0].Masking)
}

func TestCreateNotificationDestinationHTTP(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/notification_destinations/http", r.URL.Path)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.JSONEq(t, `{"http_config":{"name":"my-webhook","url":"https://example.com/webhook","description":"desc","headers":[{"key":"Authorization","value":"Bearer xxx","masking":true}],"query_params":[]}}`, string(body))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":8,"type":"http","name":"my-webhook","url":"https://example.com/webhook","description":"desc","headers":[{"id":1,"key":"Authorization","value":"","masking":true}],"query_params":[]}`))
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	out, err := c.CreateNotificationDestination("http", &CreateNotificationDestinationInput{
		HTTPConfig: &notificationDestinationParameters.HTTPConfigInput{
			Name:        lo.ToPtr("my-webhook"),
			URL:         lo.ToPtr("https://example.com/webhook"),
			Description: lo.ToPtr("desc"),
			Headers:     &[]notificationDestinationParameters.HTTPKeyValueInput{{Key: "Authorization", Value: "Bearer xxx", Masking: true}},
			QueryParams: &[]notificationDestinationParameters.HTTPKeyValueInput{},
		},
	})

	require.NoError(t, err)
	assert.Equal(t, int64(8), out.ID)
	assert.Equal(t, "my-webhook", *out.Name)
	require.Len(t, out.Headers, 1)
	assert.Empty(t, out.Headers[0].Value)
}

func TestUpdateNotificationDestinationHTTP(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, "/api/notification_destinations/http/3", r.URL.Path)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.JSONEq(t, `{"http_config":{"name":"renamed","url":"https://example.com/webhook/v2","description":"","headers":[],"query_params":[{"key":"token","value":"yyy","masking":false}]}}`, string(body))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":3,"type":"http","name":"renamed","url":"https://example.com/webhook/v2","description":"","headers":[],"query_params":[{"id":21,"key":"token","value":"yyy","masking":false}]}`))
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	out, err := c.UpdateNotificationDestination("http", 3, &UpdateNotificationDestinationInput{
		HTTPConfig: &notificationDestinationParameters.HTTPConfigInput{
			Name:        lo.ToPtr("renamed"),
			URL:         lo.ToPtr("https://example.com/webhook/v2"),
			Description: lo.ToPtr(""),
			Headers:     &[]notificationDestinationParameters.HTTPKeyValueInput{},
			QueryParams: &[]notificationDestinationParameters.HTTPKeyValueInput{{Key: "token", Value: "yyy", Masking: false}},
		},
	})

	require.NoError(t, err)
	assert.Equal(t, "renamed", *out.Name)
	require.Len(t, out.QueryParams, 1)
	assert.Equal(t, "yyy", out.QueryParams[0].Value)
}
