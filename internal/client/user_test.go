package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// ListUsers

func TestListUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"path", r.URL.Path, "/api/users"},
			{"method", r.Method, http.MethodGet},
		}
		testCases(t, cases)
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
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	output, err := client.ListUsers(nil)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if len(output.Items) != 2 {
		t.Errorf("Expected output.Items to have 2 items, got %d", len(output.Items))
	}
	cases := []Case{
		{"first item's ID", output.Items[0].ID, int64(1)},
		{"first item's email", output.Items[0].Email, "test1@example.com"},
		{"first item's role", output.Items[0].Role, "admin"},
		{"first item's can_use_audit_log", output.Items[0].CanUseAuditLog, true},
		{"first item's is_restricted_connection_modify", output.Items[0].IsRestrictedConnectionModify, false},
		{"first item's last_sign_in_at", output.Items[0].LastSignInAt, "2024-07-29T19:00:00.000+09:00"},
		{"first item's created_at", output.Items[0].CreatedAt, "2024-07-29T19:00:00.000+09:00"},
		{"first item's updated_at", output.Items[0].UpdatedAt, "2024-07-29T20:00:00.000+09:00"},
		{"second item's ID", output.Items[1].ID, int64(2)},
		{"second item's email", output.Items[1].Email, "test2@example.com"},
		{"second item's role", output.Items[1].Role, "member"},
		{"second item's can_use_audit_log", output.Items[1].CanUseAuditLog, false},
		{"second item's is_restricted_connection_modify", output.Items[1].IsRestrictedConnectionModify, true},
		{"second item's last_sign_in_at", output.Items[1].LastSignInAt, "2024-07-29T19:00:00.000+09:00"},
		{"second item's created_at", output.Items[1].CreatedAt, "2024-07-29T21:00:00.000+09:00"},
		{"second item's updated_at", output.Items[1].UpdatedAt, "2024-07-29T22:00:00.000+09:00"},
	}
	testCases(t, cases)
}

func TestListUsersLimitAndCursor(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"query parameter limit", r.URL.Query().Get("limit"), "1"},
			{"query parameter cursor", r.URL.Query().Get("cursor"), "test_prev_cursor"},
		}
		testCases(t, cases)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := `
			{
				"items": [],
				"next_cursor": "test_next_cursor"
			}
		`
		_, err := w.Write([]byte(resp))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := ListUsersInput{}
	input.SetLimit(1)
	input.SetCursor("test_prev_cursor")
	output, err := client.ListUsers(&input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	cases := []Case{
		{"next_cursor", *output.NextCursor, "test_next_cursor"},
	}
	testCases(t, cases)
}

// GetUser

func TestGetUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"path", r.URL.Path, "/api/users/1"},
			{"method", r.Method, http.MethodGet},
		}
		testCases(t, cases)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := `
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
		`
		_, err := w.Write([]byte(resp))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	output, err := client.GetUser(1)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	cases := []Case{
		{"ID", output.ID, int64(1)},
		{"email", output.Email, "test1@example.com"},
		{"role", output.Role, "admin"},
		{"can_use_audit_log", output.CanUseAuditLog, true},
		{"is_restricted_connection_modify", output.IsRestrictedConnectionModify, false},
		{"last_sign_in_at", output.LastSignInAt, "2024-07-29T19:00:00.000+09:00"},
		{"created_at", output.CreatedAt, "2024-07-29T19:00:00.000+09:00"},
		{"updated_at", output.UpdatedAt, "2024-07-29T20:00:00.000+09:00"},
	}
	testCases(t, cases)
}

// CreateUser

func TestCreateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"path", r.URL.Path, "/api/users"},
			{"method", r.Method, http.MethodPost},
		}
		testCases(t, cases)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		resp := `
			{
				"id": 1,
				"email": "test@example.com",
				"role": "admin",
				"can_use_audit_log": true,
				"is_restricted_connection_modify": false
				"last_sign_in_at": "2024-07-29T19:00:00.000+09:00",
				"created_at": "2024-07-29T19:00:00.000+09:00",
				"updated_at": "2024-07-29T20:00:00.000+09:00"
			}
		`
		_, err := w.Write([]byte(resp))
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
	}))
	defer server.Close()

	client := NewDevTroccoClient("1234567890", server.URL)
	input := CreateUserInput{
		Email:                        "test@example.com",
		Role:                         "admin",
		CanUseAuditLog:               true,
		IsRestrictedConnectionModify: false,
	}
	output, err := client.CreateUser(&input)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	cases := []Case{
		{"ID", output.ID, int64(1)},
		{"email", output.Email, "test@exmaple.com"},
		{"role", output.Role, "admin"},
		{"can_use_audit_log", output.CanUseAuditLog, true},
		{"is_restricted_connection_modify", output.IsRestrictedConnectionModify, false},
		{"last_sign_in_at", output.LastSignInAt, "2024-07-29T19:00:00.000+09:00"},
		{"created_at", output.CreatedAt, "2024-07-29T19:00:00.000+09:00"},
		{"updated_at", output.UpdatedAt, "2024-07-29T20:00:00.000+09:00"},
	}
	testCases(t, cases)
}
