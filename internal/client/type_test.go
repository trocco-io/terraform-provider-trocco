package client

import (
	"terraform-provider-trocco/internal/client/parameters"
	"testing"
)

func TestNullableInt64MarshalJSON(t *testing.T) {
	cases := []struct {
		name     string
		value    parameters.NullableInt64
		expected string
	}{
		{
			name:     "valid",
			value:    parameters.NullableInt64{Value: 123, Valid: true},
			expected: "123",
		},
		{
			name:     "invalid",
			value:    parameters.NullableInt64{Valid: false},
			expected: "null",
		},
	}
	for _, c := range cases {
		t.Run("should marshal "+c.name+" NullableInt64", func(t *testing.T) {
			b, err := c.value.MarshalJSON()
			if err != nil {
				t.Errorf("Expected no error, got %s", err)
				return
			}
			if string(b) != c.expected {
				t.Errorf("Expected %s, got %s", c.expected, string(b))
			}
		})
	}
}

func TestNullableStringMarshalJSON(t *testing.T) {
	cases := []struct {
		name     string
		value    parameters.NullableString
		expected string
	}{
		{
			name:     "valid",
			value:    parameters.NullableString{Value: "foo", Valid: true},
			expected: `"foo"`,
		},
		{
			name:     "invalid",
			value:    parameters.NullableString{Valid: false},
			expected: "null",
		},
	}
	for _, c := range cases {
		t.Run("should marshal "+c.name+" NullableString", func(t *testing.T) {
			b, err := c.value.MarshalJSON()
			if err != nil {
				t.Errorf("Expected no error, got %s", err)
				return
			}
			if string(b) != c.expected {
				t.Errorf("Expected %s, got %s", c.expected, string(b))
			}
		})
	}
}
