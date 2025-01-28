package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

// ListLabels

func TestListLabels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/labels", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := `
			{
				"items": [
					{
						"id": 1,
						"name": "Label 1",
						"description": "Description 1",
						"color": "#FFFFFF",
						"created_at": "2024-07-29T19:00:00.000+09:00",
						"updated_at": "2024-07-29T20:00:00.000+09:00"
    				},
    				{
						"id": 2,
						"name": "Label 2",
						"description": "Description 2",
						"color": "#000000",
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

	output, err := NewDevTroccoClient("1234567890", server.URL).ListLabels(nil)

	assert.NoError(t, err)
	assert.Len(t, output.Items, 2)
	assert.Equal(t, int64(1), output.Items[0].ID)
	assert.Equal(t, "Label 1", output.Items[0].Name)
	assert.Equal(t, "Description 1", *output.Items[0].Description)
	assert.Equal(t, "#FFFFFF", output.Items[0].Color)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.Items[0].CreatedAt)
	assert.Equal(t, "2024-07-29T20:00:00.000+09:00", output.Items[0].UpdatedAt)
	assert.Equal(t, int64(2), output.Items[1].ID)
	assert.Equal(t, "Label 2", output.Items[1].Name)
	assert.Equal(t, "Description 2", *output.Items[1].Description)
	assert.Equal(t, "#000000", output.Items[1].Color)
	assert.Equal(t, "2024-07-29T21:00:00.000+09:00", output.Items[1].CreatedAt)
	assert.Equal(t, "2024-07-29T22:00:00.000+09:00", output.Items[1].UpdatedAt)
}

func TestListLabelsLimitAndCursor(t *testing.T) {
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

	input := ListLabelsInput{}
	input.SetLimit(1)
	input.SetCursor("test_prev_cursor")
	output, err := NewDevTroccoClient("1234567890", server.URL).ListLabels(&input)

	assert.NoError(t, err)
	assert.Equal(t, "test_next_cursor", *output.NextCursor)
}

// GetLabel

func TestGetLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/labels/1", r.URL.Path)
		assert.Equal(t, http.MethodGet, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
				"id": 1,
				"name": "Label 1",
				"description": "Description 1",
				"color": "#FFFFFF",
				"created_at": "2024-07-29T19:00:00.000+09:00",
				"updated_at": "2024-07-29T20:00:00.000+09:00"
    		}
		`))

		assert.NoError(t, err)
	}))
	defer server.Close()

	output, err := NewDevTroccoClient("1234567890", server.URL).GetLabel(1)

	assert.NoError(t, err)

	assert.Equal(t, int64(1), output.ID)
	assert.Equal(t, "Label 1", output.Name)
	assert.Equal(t, "Description 1", *output.Description)
	assert.Equal(t, "#FFFFFF", output.Color)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.CreatedAt)
	assert.Equal(t, "2024-07-29T20:00:00.000+09:00", output.UpdatedAt)
}

// CreateLabel

func TestCreateLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/labels", r.URL.Path)
		assert.Equal(t, http.MethodPost, r.Method)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err := w.Write([]byte(`
			{
				"id": 1,
				"name": "Label 1",
				"description": "Description 1",
				"color": "#FFFFFF",
				"created_at": "2024-07-29T19:00:00.000+09:00",
				"updated_at": "2024-07-29T20:00:00.000+09:00"
    		}
		`))

		assert.NoError(t, err)
	}))
	defer server.Close()

	output, err := NewDevTroccoClient("1234567890", server.URL).CreateLabel(&CreateLabelInput{
		Name:        "Label 1",
		Description: lo.ToPtr("Description 1"),
		Color:       "#FFFFFF",
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(1), output.ID)
	assert.Equal(t, "Label 1", output.Name)
	assert.Equal(t, "Description 1", *output.Description)
	assert.Equal(t, "#FFFFFF", output.Color)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.CreatedAt)
	assert.Equal(t, "2024-07-29T20:00:00.000+09:00", output.UpdatedAt)
}

// UpdateLabel

func TestUpdateLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"path", r.URL.Path, "/api/labels/1"},
			{"method", r.Method, http.MethodPatch},
		}
		testCases(t, cases)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := `
			{
				"id": 1,
				"name": "Updated Label",
				"description": "Updated Description",
				"color": "#000000",
				"created_at": "2024-07-29T19:00:00.000+09:00",
				"updated_at": "2024-07-29T20:00:00.000+09:00"
    		}
		`
		_, err := w.Write([]byte(resp))

		assert.NoError(t, err)
	}))
	defer server.Close()

	output, err := NewDevTroccoClient("1234567890", server.URL).UpdateLabel(1, &UpdateLabelInput{
		Name:        lo.ToPtr("Updated Label"),
		Description: lo.ToPtr("Updated Description"),
		Color:       lo.ToPtr("#000000"),
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(1), output.ID)
	assert.Equal(t, "Updated Label", output.Name)
	assert.Equal(t, "Updated Description", *output.Description)
	assert.Equal(t, "#000000", output.Color)
	assert.Equal(t, "2024-07-29T19:00:00.000+09:00", output.CreatedAt)
	assert.Equal(t, "2024-07-29T20:00:00.000+09:00", output.UpdatedAt)
}

// DeleteLabel

func TestDeleteLabel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []Case{
			{"path", r.URL.Path, "/api/labels/1"},
			{"method", r.Method, http.MethodDelete},
		}
		testCases(t, cases)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	err := NewDevTroccoClient("1234567890", server.URL).DeleteLabel(1)

	assert.NoError(t, err)
}
