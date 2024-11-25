package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestGetConnections(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/connections/bigquery/", r.URL.Path)
		assert.Equal(t, "1", r.URL.Query().Get("limit"))
		assert.Equal(t, "Fk/RmZrNji8DOg6SefOy7A==", r.URL.Query().Get("cursor"))

		w.Header().Set("Content-Type", "application/json")

		c := ConnectionList{
			Connections: []*Connection{
				{
					ID:              8,
					Name:            lo.ToPtr("Foo"),
					Description:     lo.ToPtr("The quick brown fox jumps over the lazy dog."),
					ResourceGroupID: lo.ToPtr(int64(42)),
				},
			},
			NextCursor: "FkqWYxQrrVonxahG26lVQg==",
		}
		if err := json.NewEncoder(w).Encode(c); err != nil {
			panic(err)
		}
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	out, err := c.GetConnections("bigquery", &GetConnectionsInput{
		Limit:  1,
		Cursor: "Fk/RmZrNji8DOg6SefOy7A==",
	})

	assert.NoError(t, err)

	assert.Len(t, out.Connections, 1)
	assert.Equal(t, int64(8), out.Connections[0].ID)
	assert.Equal(t, "Foo", *out.Connections[0].Name)
	assert.Equal(t, "The quick brown fox jumps over the lazy dog.", *out.Connections[0].Description)
	assert.Equal(t, int64(42), *out.Connections[0].ResourceGroupID)
	assert.Equal(t, "FkqWYxQrrVonxahG26lVQg==", out.NextCursor)
}

func TestGetConnection(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/api/connections/bigquery/8", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		c := Connection{
			ID:              8,
			Name:            lo.ToPtr("Foo"),
			Description:     lo.ToPtr("The quick brown fox jumps over the lazy dog."),
			ResourceGroupID: lo.ToPtr(int64(42)),
		}
		if err := json.NewEncoder(w).Encode(c); err != nil {
			panic(err)
		}
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	out, err := c.GetConnection("bigquery", 8)

	assert.NoError(t, err)

	assert.Equal(t, int64(8), out.ID)
	assert.Equal(t, "Foo", *out.Name)
	assert.Equal(t, "The quick brown fox jumps over the lazy dog.", *out.Description)
	assert.Equal(t, int64(42), *out.ResourceGroupID)
}

func TestCreateConnection(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/api/connections/bigquery", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		c := Connection{
			ID:              8,
			Name:            lo.ToPtr("Foo"),
			Description:     lo.ToPtr("The quick brown fox jumps over the lazy dog."),
			ResourceGroupID: lo.ToPtr(int64(42)),
		}
		if err := json.NewEncoder(w).Encode(c); err != nil {
			panic(err)
		}
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	out, err := c.CreateConnection("bigquery", &CreateConnectionInput{
		Name:        "Foo",
		Description: lo.ToPtr("The quick brown fox jumps over the lazy dog."),
		ResourceGroupID: lo.ToPtr(NullableInt64{
			Valid: true,
			Value: int64(42),
		}),
	})

	assert.NoError(t, err)

	assert.Equal(t, int64(8), out.ID)
	assert.Equal(t, "Foo", *out.Name)
	assert.Equal(t, "The quick brown fox jumps over the lazy dog.", *out.Description)
	assert.Equal(t, int64(42), *out.ResourceGroupID)
}

func TestUpdateConnection(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPatch, r.Method)
		assert.Equal(t, "/api/connections/bigquery/8", r.URL.Path)

		w.Header().Set("Content-Type", "application/json")

		c := Connection{
			ID:              8,
			Name:            lo.ToPtr("Foo"),
			Description:     lo.ToPtr("The quick brown fox jumps over the lazy dog."),
			ResourceGroupID: lo.ToPtr(int64(42)),
		}
		if err := json.NewEncoder(w).Encode(c); err != nil {
			panic(err)
		}
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	out, err := c.UpdateConnection("bigquery", 8, &UpdateConnectionInput{
		Name:        lo.ToPtr("Foo"),
		Description: lo.ToPtr("The quick brown fox jumps over the lazy dog."),
		ResourceGroupID: lo.ToPtr(NullableInt64{
			Valid: true,
			Value: int64(42),
		}),
	})

	assert.NoError(t, err)

	assert.Equal(t, int64(8), out.ID)
	assert.Equal(t, "Foo", *out.Name)
	assert.Equal(t, "The quick brown fox jumps over the lazy dog.", *out.Description)
	assert.Equal(t, int64(42), *out.ResourceGroupID)
}

func TestDeleteConnection(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check request method and path are correct.
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/api/connections/bigquery/8", r.URL.Path)
	}))
	defer s.Close()

	c := NewDevTroccoClient("1234567890", s.URL)

	err := c.DeleteConnection("bigquery", 8)

	assert.NoError(t, err)
}
