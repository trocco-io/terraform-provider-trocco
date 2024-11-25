package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

// ListUsers

func TestListUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/users", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := `
			{
				"items": [
					{
						"id": 1,
						"email": "test1@example.com",
						"role": "admin",
						"can_use_audit_log": true,
						"is_restricted_connection_modify": false,
						"last_sign_in_at": "2024-07-29T19:00:00.000+09:00",
						"created_at": "2024-07-29T19:00:00.000+09:00",
						"updated_at": "2024-07-29T20:00:00.000+09:00"
    				},
    				{
						"id": 2,
						"email": "test2@example.com",
						"role": "member",
						"can_use_audit_log": false,
						"is_restricted_connection_modify": true,
						"last_sign_in_at": "2024-07-29T19:00:00.000+09:00",
						"created_at": "2024-07-29T21:00:00.000+09:00",
						"updated_at": "2024-07-29T22:00:00.000+09:00"
    				}
				]
    		}
    	`
		_, err := w.Write([]byte(resp))

		assert.NoError(t, err)
	}))
	defer server.Close()

	output, err := NewDevTroccoClient("1234567890", server.URL).ListUsers(nil)

	assert.NoError(t, err)
	assert.Len(t, output.Items, 2)
	assert.Equal(t, int64(1), output.Items[0].ID)
	assert.Equal(t, "test1@example.com", output.Items[0].Email)
	assert.Equal(t, "admin", output.Items[0].Role)
	assert.True(t, output.Items[0].CanUseAuditLog)
	assert.False(t, output.Items[0].IsRestrictedConnectionModify)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.Items[0].LastSignInAt)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.Items[0].CreatedAt)
	assert.Equal(t, "2024-07-29T20:00:00.000+09:00", output.Items[0].UpdatedAt)
	assert.Equal(t, int64(2), output.Items[1].ID)
	assert.Equal(t, "test2@example.com", output.Items[1].Email)
	assert.Equal(t, "member", output.Items[1].Role)
	assert.False(t, output.Items[1].CanUseAuditLog)
	assert.True(t, output.Items[1].IsRestrictedConnectionModify)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.Items[1].LastSignInAt)
	assert.Equal(t, "2024-07-29T21:00:00.000+09:00", output.Items[1].CreatedAt)
	assert.Equal(t, "2024-07-29T22:00:00.000+09:00", output.Items[1].UpdatedAt)
}

func TestListUsersLimitAndCursor(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "1", r.URL.Query().Get("limit"))
		assert.Equal(t, "test_prev_cursor", r.URL.Query().Get("cursor"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
				"items": [],
				"next_cursor": "test_next_cursor"
			}
		`))

		assert.NoError(t, err)
	}))
	defer server.Close()

	input := ListUsersInput{}
	input.SetLimit(1)
	input.SetCursor("test_prev_cursor")
	output, err := NewDevTroccoClient("1234567890", server.URL).ListUsers(&input)

	assert.NoError(t, err)
	assert.Equal(t, "test_next_cursor", *output.NextCursor)
}

// GetUser

func TestGetUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/users/1", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
				"id": 1,
				"email": "test1@example.com",
				"role": "admin",
				"can_use_audit_log": true,
				"is_restricted_connection_modify": false,
				"last_sign_in_at": "2024-07-29T19:00:00.000+09:00",
				"created_at": "2024-07-29T19:00:00.000+09:00",
				"updated_at": "2024-07-29T20:00:00.000+09:00"
    		}
		`))

		assert.NoError(t, err)
	}))
	defer server.Close()

	output, err := NewDevTroccoClient("1234567890", server.URL).GetUser(1)

	assert.NoError(t, err)

	assert.Equal(t, int64(1), output.ID)
	assert.Equal(t, "test1@example.com", output.Email)
	assert.Equal(t, "admin", output.Role)
	assert.True(t, output.CanUseAuditLog)
	assert.False(t, output.IsRestrictedConnectionModify)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.LastSignInAt)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.CreatedAt)
	assert.Equal(t, "2024-07-29T20:00:00.000+09:00", output.UpdatedAt)
}

// CreateUser

func TestCreateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/users", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err := w.Write([]byte(`
			{
				"id": 1,
				"email": "test@example.com",
				"role": "admin",
				"can_use_audit_log": true,
				"is_restricted_connection_modify": false,
				"last_sign_in_at": "2024-07-29T19:00:00.000+09:00",
				"created_at": "2024-07-29T19:00:00.000+09:00",
				"updated_at": "2024-07-29T20:00:00.000+09:00"
    		}
		`))

		assert.NoError(t, err)
	}))
	defer server.Close()

	output, err := NewDevTroccoClient("1234567890", server.URL).CreateUser(&CreateUserInput{
		Email:                        "test@example.com",
		Role:                         "admin",
		CanUseAuditLog:               lo.ToPtr(true),
		IsRestrictedConnectionModify: lo.ToPtr(false),
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(1), output.ID)
	assert.Equal(t, "test@example.com", output.Email)
	assert.Equal(t, "admin", output.Role)
	assert.True(t, output.CanUseAuditLog)
	assert.False(t, output.IsRestrictedConnectionModify)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.LastSignInAt)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.CreatedAt)
	assert.Equal(t, "2024-07-29T20:00:00.000+09:00", output.UpdatedAt)
}

// UpdateUser

func TestUpdateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"path", r.URL.Path, "/api/users/1"},
			{"method", r.Method, http.MethodPatch},
		}
		testCases(t, cases)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := `
			{
				"id": 1,
				"email": "test@example.com",
				"role": "admin",
				"can_use_audit_log": true,
				"is_restricted_connection_modify": false,
				"last_sign_in_at": "2024-07-29T19:00:00.000+09:00",
				"created_at": "2024-07-29T19:00:00.000+09:00",
				"updated_at": "2024-07-29T20:00:00.000+09:00"
    		}
		`
		_, err := w.Write([]byte(resp))

		assert.NoError(t, err)
	}))
	defer server.Close()

	output, err := NewDevTroccoClient("1234567890", server.URL).UpdateUser(1, &UpdateUserInput{
		Role:                         lo.ToPtr("admin"),
		CanUseAuditLog:               lo.ToPtr(true),
		IsRestrictedConnectionModify: lo.ToPtr(false),
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(1), output.ID)
	assert.Equal(t, "test@example.com", output.Email)
	assert.Equal(t, "admin", output.Role)
	assert.True(t, output.CanUseAuditLog)
	assert.False(t, output.IsRestrictedConnectionModify)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.LastSignInAt)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.CreatedAt)
	assert.Equal(t, "2024-07-29T20:00:00.000+09:00", output.UpdatedAt)
}

// DeleteUser

func TestDeleteUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"path", r.URL.Path, "/api/users/1"},
			{"method", r.Method, http.MethodDelete},
		}
		testCases(t, cases)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	err := NewDevTroccoClient("1234567890", server.URL).DeleteUser(1)

	assert.NoError(t, err)
}
