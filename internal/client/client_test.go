package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewTroccoClient(t *testing.T) {
	t.Run("should return a new TroccoClient for japan region", func(t *testing.T) {
		token := "1234567890"
		client := NewTroccoClient(token)
		if client == nil {
			t.Error("Expected a new TroccoClient, got nil")
		}
		if client.BaseURL != "https://trocco.io" {
			t.Errorf("Expected BaseURL to be https://trocco.io, got %s", client.BaseURL)
		}
		if client.APIKey != "1234567890" {
			t.Errorf("Expected APIKey to be 1234567890, got %s", client.APIKey)
		}
	})
}

func TestNewTroccoClientWithRegion(t *testing.T) {
	cases := []struct {
		region          string
		expectedBaseURL string
	}{
		{"japan", "https://trocco.io"},
		{"india", "https://in.trocco.io"},
		{"korea", "https://kr.trocco.io"},
	}
	for _, c := range cases {
		t.Run("should return a new TroccoClient for "+c.region+" region", func(t *testing.T) {
			client, err := NewTroccoClientWithRegion("1234567890", c.region)
			if err != nil {
				t.Errorf("Expected no error, got %s", err)
			}
			if client == nil {
				t.Error("Expected a new TroccoClient, got nil")
			}
			if client.BaseURL != c.expectedBaseURL {
				t.Errorf("Expected BaseURL to be %s, got %s", c.expectedBaseURL, client.BaseURL)
			}
		})
	}
}

func TestDo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cases := []struct {
			name     string
			value    string
			expected string
		}{
			{"Method", r.Method, "GET"},
			{"Path", r.URL.Path, "/api"},
			{"Header Authorization", r.Header.Get("Authorization"), "Token 1234567890"},
			{"Header Accept", r.Header.Get("Accept"), "application/json"},
			{"Header Content-Type", r.Header.Get("Content-Type"), "application/json"},
		}
		for _, c := range cases {
			if c.value != c.expected {
				t.Errorf("Expected %s to be %s, got %s", c.name, c.expected, c.value)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()
	client := NewDevTroccoClient("1234567890", server.URL)
	err := client.Do("GET", "/api", nil, nil)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func TestDoWithInput(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		if string(body) != `{"name":"test"}` {
			t.Errorf("Expected body to be {\"name\":\"test\"}, got %s", string(body))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()
	client := NewDevTroccoClient("1234567890", server.URL)
	type TestBody struct {
		Name string `json:"name"`
	}
	input := TestBody{Name: "test"}
	err := client.Do("POST", "", &input, nil)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
}

func TestDoWithOutput(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name":"test"}`))
	}))
	defer server.Close()
	client := NewDevTroccoClient("1234567890", server.URL)
	type TestBody struct {
		Name string `json:"name"`
	}
	output := TestBody{}
	err := client.Do("POST", "", nil, &output)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}
	if output.Name != "test" {
		t.Errorf("Expected output.Name to be test, got %s", output.Name)
	}
}
