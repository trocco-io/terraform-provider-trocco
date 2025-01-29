package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

// List Teams

func TestListResourceGroup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/resource_groups", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		resp := `
			{
				"items": [
					{
						"id": 1,
						"name": "ResourceGroup 1",
						"description": "ResourceGroup 1 description",
						"created_at": "2023-10-16T18:24:51.806+09:00",
						"updated_at": "2023-10-16T18:24:51.806+09:00"
					},
					{
						"id": 2,
						"name": "ResourceGroup 2",
						"description": "ResourceGroup 2 description",
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

	output, err := NewDevTroccoClient("1234567890", server.URL).ListResourceGroups(nil)

	assert.NoError(t, err)
	assert.Len(t, output.Items, 2)
	assert.Equal(t, int64(1), output.Items[0].ID)
	assert.Equal(t, "ResourceGroup 1", output.Items[0].Name)
	assert.Equal(t, "ResourceGroup 1 description", *output.Items[0].Description)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.Items[0].CreatedAt)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.Items[0].UpdatedAt)
	assert.Equal(t, int64(2), output.Items[1].ID)
	assert.Equal(t, "ResourceGroup 2", output.Items[1].Name)
	assert.Equal(t, "ResourceGroup 2 description", *output.Items[1].Description)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.Items[1].CreatedAt)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.Items[1].UpdatedAt)
}

func TestListResourceGroupsLimitAndCursor(t *testing.T) {
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
						"name": "ResourceGroup 1",
						"description": "ResourceGroup 1 description",
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

	input := &ListResourceGroupInput{}
	input.SetLimit(1)
	input.SetCursor("test_prev_cursor")

	output, err := NewDevTroccoClient("1234567890", server.URL).ListResourceGroups(input)

	assert.NoError(t, err)
	assert.Len(t, output.Items, 1)
	assert.Equal(t, int64(1), output.Items[0].ID)
	assert.Equal(t, "ResourceGroup 1", output.Items[0].Name)
	assert.Equal(t, "ResourceGroup 1 description", *output.Items[0].Description)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.Items[0].CreatedAt)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.Items[0].UpdatedAt)
	assert.Equal(t, "test_next_cursor", *output.NextCursor)
}

// Get Team

func TestGetResourceGroup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/resource_groups/1", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		resp := `
			{
				"id": 1,
				"name": "ResourceGroup 1",
				"description": "ResourceGroup 1 description",
				"created_at": "2023-10-16T18:24:51.806+09:00",
				"updated_at": "2023-10-16T18:24:51.806+09:00",
				"teams": [
								{
									"team_id": 1,
									"role": "administrator"
								},
								{
									"team_id": 2,
									"role": "operator"
								}
							]
			}
		`
		_, err := w.Write([]byte(resp))
		assert.NoError(t, err)
	}))
	defer server.Close()

	output, err := NewDevTroccoClient("1234567890", server.URL).GetResourceGroup(1)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), output.ID)
	assert.Equal(t, "ResourceGroup 1", output.Name)
	assert.Equal(t, "ResourceGroup 1 description", *output.Description)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.CreatedAt)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.UpdatedAt)
	assert.Len(t, output.Teams, 2)
	assert.Equal(t, int64(1), output.Teams[0].TeamID)
	assert.Equal(t, "administrator", output.Teams[0].Role)
	assert.Equal(t, int64(2), output.Teams[1].TeamID)
	assert.Equal(t, "operator", output.Teams[1].Role)
}

// Create Team

func TestCreateResourceGroup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/resource_groups", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		resp := `
			{
				"id": 1,
				"name": "ResourceGroup",
				"description": "description",
				"created_at": "2023-10-16T18:24:51.806+09:00",
				"updated_at": "2023-10-16T18:24:51.806+09:00",
				"teams": [
								{
									"team_id": 1,
									"role": "administrator"
								}
							]
			}
		`
		_, err := w.Write([]byte(resp))
		assert.NoError(t, err)
	}))
	defer server.Close()

	input := &CreateResourceGroupInput{
		Name:        lo.ToPtr("ResourceGroup"),
		Description: lo.ToPtr("description"),
		Teams: []TeamRoleInput{
			{TeamID: 1, Role: "administrator"},
		},
	}

	output, err := NewDevTroccoClient("1234567890", server.URL).CreateResourceGroup(input)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), output.ID)
	assert.Equal(t, "ResourceGroup", output.Name)
	assert.Equal(t, "description", *output.Description)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.CreatedAt)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.UpdatedAt)
	assert.Len(t, output.Teams, 1)
	assert.Equal(t, int64(1), output.Teams[0].TeamID)
	assert.Equal(t, "administrator", output.Teams[0].Role)
}

// Update Team

func TestUpdateResourceGroup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/resource_groups/1", r.URL.Path)
		assert.Equal(t, http.MethodPatch, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		resp := `
			{
				"id": 1,
				"name": "ResourceGroup",
				"description": "description",
				"created_at": "2023-10-16T18:24:51.806+09:00",
				"updated_at": "2023-10-16T18:24:51.806+09:00",
				"teams": [
								{
									"team_id": 1,
									"role": "administrator"
								},
								{
									"team_id": 2,
									"role": "operator"
								}
							]
			}
		`
		_, err := w.Write([]byte(resp))
		assert.NoError(t, err)
	}))
	defer server.Close()

	input := &UpdateResourceGroupInput{
		Name:        lo.ToPtr("ResourceGroup"),
		Description: lo.ToPtr("description"),
		Teams: []TeamRoleInput{
			{TeamID: 1, Role: "administrator"},
			{TeamID: 2, Role: "operator"},
		},
	}

	output, err := NewDevTroccoClient("1234567890", server.URL).UpdateResourceGroup(1, input)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), output.ID)
	assert.Equal(t, "ResourceGroup", output.Name)
	assert.Equal(t, "description", *output.Description)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.CreatedAt)
	assert.Equal(t, "2023-10-16T18:24:51.806+09:00", output.UpdatedAt)
	assert.Len(t, output.Teams, 2)
	assert.Equal(t, int64(1), output.Teams[0].TeamID)
	assert.Equal(t, "administrator", output.Teams[0].Role)
	assert.Equal(t, int64(2), output.Teams[1].TeamID)
	assert.Equal(t, "operator", output.Teams[1].Role)
}

// Delete Team

func TestDeleteResourceGroup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/resource_groups/1", r.URL.Path)
		assert.Equal(t, http.MethodDelete, r.Method)

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	err := NewDevTroccoClient("1234567890", server.URL).DeleteResourceGroup(1)

	assert.NoError(t, err)
}
