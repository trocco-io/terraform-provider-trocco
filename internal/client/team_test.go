package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// List Teams

func TestListTeams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/teams", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		resp := `
			{
				"items": [
					{
						"id": 1,
						"name": "Team 1",
						"description": "team 1 description",
						"created_at": "2023-10-16T18:24:51.806+09:00",
						"updated_at": "2023-10-16T18:24:51.806+09:00"
					},
					{
						"id": 2,
						"name": "Team 2",
						"description": "team 2 description",
						"created_at": "2023-10-16T18:24:51.806+09:00",
						"updated_at": "2023-10-16T18:24:51.806+09:00"
					}
				]
			}
		`
		_, err := w.Write([]byte(resp))
		assert.NoError(t, err)
	}))
	defer server.Close()

	output, err := NewDevTroccoClient("1234567890", server.URL).ListTeams(nil)

	require.NoError(t, err)
	assert.Len(t, output.Items, 2)
	assert.Equal(t, int64(1), output.Items[0].ID)
	assert.Equal(t, "Team 1", output.Items[0].Name)
	assert.Equal(t, "team 1 description", *output.Items[0].Description)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.Items[0].CreatedAt)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.Items[0].UpdatedAt)
	assert.Equal(t, int64(2), output.Items[1].ID)
	assert.Equal(t, "Team 2", output.Items[1].Name)
	assert.Equal(t, "team 2 description", *output.Items[1].Description)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.Items[1].CreatedAt)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.Items[1].UpdatedAt)
}

func TestListTeamsLimitAndCursor(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "1", r.URL.Query().Get("limit"))
		assert.Equal(t, "test_prev_cursor", r.URL.Query().Get("cursor"))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		resp := `
			{
				"items": [
					{
						"id": 1,
						"name": "Team 1",
						"description": "team 1 description",
						"created_at": "2023-10-16T18:24:51.806+09:00",
						"updated_at": "2023-10-16T18:24:51.806+09:00"
					}
				],
				"next_cursor": "test_next_cursor"
			}
		`
		_, err := w.Write([]byte(resp))
		assert.NoError(t, err)
	}))
	defer server.Close()

	input := &ListTeamsInput{}
	input.SetLimit(1)
	input.SetCursor("test_prev_cursor")

	output, err := NewDevTroccoClient("1234567890", server.URL).ListTeams(input)

	require.NoError(t, err)
	assert.Len(t, output.Items, 1)
	assert.Equal(t, int64(1), output.Items[0].ID)
	assert.Equal(t, "Team 1", output.Items[0].Name)
	assert.Equal(t, "team 1 description", *output.Items[0].Description)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.Items[0].CreatedAt)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.Items[0].UpdatedAt)
	assert.Equal(t, "test_next_cursor", *output.NextCursor)
}

// Get Team

func TestGetTeam(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/teams/1", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		resp := `
			{
				"id": 1,
				"name": "Team 1",
				"description": "team 1 description",
				"created_at": "2023-10-16T18:24:51.806+09:00",
				"updated_at": "2023-10-16T18:24:51.806+09:00",
				"members": [
								{
									"user_id": 1,
									"email": "admin@example.com",
									"role": "team_admin"
								},
								{
									"user_id": 2,
									"email": "member@example.com",
									"role": "team_member"
								}
							]
			}
		`
		_, err := w.Write([]byte(resp))
		assert.NoError(t, err)
	}))
	defer server.Close()

	output, err := NewDevTroccoClient("1234567890", server.URL).GetTeam(1)

	require.NoError(t, err)
	assert.Equal(t, int64(1), output.ID)
	assert.Equal(t, "Team 1", output.Name)
	assert.Equal(t, "team 1 description", *output.Description)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.CreatedAt)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.UpdatedAt)
	assert.Len(t, output.Members, 2)
	assert.Equal(t, int64(1), output.Members[0].UserID)
	assert.Equal(t, "admin@example.com", output.Members[0].Email)
	assert.Equal(t, "team_admin", output.Members[0].Role)
	assert.Equal(t, int64(2), output.Members[1].UserID)
	assert.Equal(t, "member@example.com", output.Members[1].Email)
	assert.Equal(t, "team_member", output.Members[1].Role)
}

// Create Team

func TestCreateTeam(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/teams", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		resp := `
			{
				"id": 1,
				"name": "Team",
				"description": "description",
				"created_at": "2023-10-16T18:24:51.806+09:00",
				"updated_at": "2023-10-16T18:24:51.806+09:00",
				"members": [
								{
									"user_id": 1,
									"email": "admin@example.com",
									"role": "team_admin"
								}
							]
			}
		`
		_, err := w.Write([]byte(resp))
		assert.NoError(t, err)
	}))
	defer server.Close()

	input := &CreateTeamInput{
		Name:        "Team",
		Description: lo.ToPtr("description"),
		Members: []MemberInput{
			{UserID: 1, Role: "team_admin"},
		},
	}

	output, err := NewDevTroccoClient("1234567890", server.URL).CreateTeam(input)

	require.NoError(t, err)
	assert.Equal(t, int64(1), output.ID)
	assert.Equal(t, "Team", output.Name)
	assert.Equal(t, "description", *output.Description)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.CreatedAt)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.UpdatedAt)
	assert.Len(t, output.Members, 1)
	assert.Equal(t, int64(1), output.Members[0].UserID)
	assert.Equal(t, "admin@example.com", output.Members[0].Email)
	assert.Equal(t, "team_admin", output.Members[0].Role)
}

// Update Team

func TestUpdateTeam(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/teams/1", r.URL.Path)
		assert.Equal(t, http.MethodPatch, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		resp := `
			{
				"id": 1,
				"name": "Team",
				"description": "description",
				"created_at": "2023-10-16T18:24:51.806+09:00",
				"updated_at": "2023-10-16T18:24:51.806+09:00",
				"members": [
								{
									"user_id": 1,
									"email": "admin@example.com",
									"role": "team_admin"
								},
								{
									"user_id": 2,
									"email": "member@example.com",
									"role": "team_member"
								}
							]
			}
		`
		_, err := w.Write([]byte(resp))
		assert.NoError(t, err)
	}))
	defer server.Close()

	input := &UpdateTeamInput{
		Name:        lo.ToPtr("Team"),
		Description: lo.ToPtr("description"),
		Members: []MemberInput{
			{UserID: 1, Role: "team_admin"},
			{UserID: 2, Role: "team_member"},
		},
	}

	output, err := NewDevTroccoClient("1234567890", server.URL).UpdateTeam(1, input)

	require.NoError(t, err)
	assert.Equal(t, int64(1), output.ID)
	assert.Equal(t, "Team", output.Name)
	assert.Equal(t, "description", *output.Description)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.CreatedAt)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.UpdatedAt)
	assert.Len(t, output.Members, 2)
	assert.Equal(t, int64(1), output.Members[0].UserID)
	assert.Equal(t, "admin@example.com", output.Members[0].Email)
	assert.Equal(t, "team_admin", output.Members[0].Role)
	assert.Equal(t, int64(2), output.Members[1].UserID)
	assert.Equal(t, "member@example.com", output.Members[1].Email)
	assert.Equal(t, "team_member", output.Members[1].Role)
}

// Delete Team

func TestDeleteTeam(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/teams/1", r.URL.Path)
		assert.Equal(t, http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	err := NewDevTroccoClient("1234567890", server.URL).DeleteTeam(1)

	assert.NoError(t, err)
}
