package input_option

import (
	"encoding/json"
	"testing"
)

func TestOracleInputOptionUnmarshalJSON(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected *string
	}{
		{
			name:     "last_record as string",
			input:    `{"oracle_connection_id":1,"connection_type":"sid","incremental_loading_enabled":true,"last_record":"2026-01-01 00:00:00"}`,
			expected: strPtr("2026-01-01 00:00:00"),
		},
		{
			name:     "last_record as array with a single element",
			input:    `{"oracle_connection_id":1,"connection_type":"sid","incremental_loading_enabled":true,"last_record":["2026-01-01 00:00:00"]}`,
			expected: strPtr("2026-01-01 00:00:00"),
		},
		{
			name:     "last_record as array with multiple elements",
			input:    `{"oracle_connection_id":1,"connection_type":"sid","incremental_loading_enabled":true,"last_record":["foo","bar"]}`,
			expected: strPtr("foo,bar"),
		},
		{
			name:     "last_record is null",
			input:    `{"oracle_connection_id":1,"connection_type":"sid","incremental_loading_enabled":true,"last_record":null}`,
			expected: nil,
		},
		{
			name:     "last_record is absent",
			input:    `{"oracle_connection_id":1,"connection_type":"sid","incremental_loading_enabled":true}`,
			expected: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got OracleInputOption
			if err := json.Unmarshal([]byte(c.input), &got); err != nil {
				t.Fatalf("expected no error, got %s", err)
			}
			if (c.expected == nil) != (got.LastRecord == nil) {
				t.Fatalf("expected %v, got %v", c.expected, got.LastRecord)
			}
			if c.expected != nil && *c.expected != *got.LastRecord {
				t.Fatalf("expected %s, got %s", *c.expected, *got.LastRecord)
			}
		})
	}
}

func TestOracleInputOptionUnmarshalJSONPreservesOtherFields(t *testing.T) {
	input := `{"oracle_connection_id":42,"connection_type":"service_name","incremental_loading_enabled":true,"table":"my_table","last_record":["foo"]}`

	var got OracleInputOption
	if err := json.Unmarshal([]byte(input), &got); err != nil {
		t.Fatalf("expected no error, got %s", err)
	}
	if got.OracleConnectionID != 42 {
		t.Errorf("expected OracleConnectionID 42, got %d", got.OracleConnectionID)
	}
	if got.ConnectionType != "service_name" {
		t.Errorf("expected ConnectionType service_name, got %s", got.ConnectionType)
	}
	if got.Table == nil || *got.Table != "my_table" {
		t.Errorf("expected Table my_table, got %v", got.Table)
	}
}

func strPtr(s string) *string {
	return &s
}
