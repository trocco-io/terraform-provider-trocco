package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"terraform-provider-trocco/internal/client/parameter"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestGetNotificationDestination(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/notification_destinations/email/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		resp := NotificationDestination{
			ID:    1,
			Email: lo.ToPtr("test@example.com"),
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	c := NewDevTroccoClient("dummy-token", ts.URL)

	result, err := c.GetNotificationDestination("email", 1)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "test@example.com", *result.Email)
}

func TestCreateNotificationDestination_Email(t *testing.T) {
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
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	out, err := c.CreateNotificationDestination("email", &CreateNotificationDestinationInput{
		EmailConfig: &EmailConfigInput{
			Email: &parameter.NullableString{Valid: true, Value: "test@example.com"},
		},
	})

	assert.NoError(t, err)

	assert.Equal(t, int64(8), out.ID)
	assert.Equal(t, "test@example.com", *out.Email)
}

func TestUpdateNotificationDestination(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, "/api/notification_destinations/email/1", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		resp := NotificationDestination{
			ID:    1,
			Email: lo.ToPtr("updated@example.com"),
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	c := NewDevTroccoClient("dummy-token", ts.URL)

	result, err := c.UpdateNotificationDestination("email", 1, &UpdateNotificationDestinationInput{
		EmailConfig: &EmailConfigInput{
			Email: &parameter.NullableString{Valid: true, Value: "test@example.com"},
		},
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(1), result.ID)
	assert.Equal(t, "updated@example.com", *result.Email)
}

func TestDeleteNotificationDestination(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/api/notification_destinations/email/1", r.URL.Path)
	}))
	defer ts.Close()

	c := NewDevTroccoClient("dummy-token", ts.URL)
	err := c.DeleteNotificationDestination("email", 1)
	assert.NoError(t, err)
}
