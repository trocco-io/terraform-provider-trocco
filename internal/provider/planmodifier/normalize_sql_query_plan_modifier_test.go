package planmodifier

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestNormalizeQuery(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Windows line endings",
			input:    "SELECT COUNT() AS count\r\nFROM table\r\nWHERE id = 1",
			expected: "SELECT COUNT() AS count\nFROM table\nWHERE id = 1",
		},
		{
			name:     "Unix line endings",
			input:    "SELECT COUNT() AS count\nFROM table\nWHERE id = 1",
			expected: "SELECT COUNT() AS count\nFROM table\nWHERE id = 1",
		},
		{
			name:     "Mixed line endings",
			input:    "SELECT COUNT() AS count\r\nFROM table\nWHERE id = 1",
			expected: "SELECT COUNT() AS count\nFROM table\nWHERE id = 1",
		},
		{
			name:     "Trailing spaces",
			input:    "SELECT COUNT() AS count FROM table WHERE id = 1; ",
			expected: "SELECT COUNT() AS count FROM table WHERE id = 1;",
		},
		{
			name:     "Semicolon with space",
			input:    "SELECT COUNT() AS count FROM table WHERE id = 1; ",
			expected: "SELECT COUNT() AS count FROM table WHERE id = 1;",
		},
		{
			name:     "Leading spaces preserved",
			input:    "SELECT COUNT() AS count\n  FROM table\n    WHERE id = 1",
			expected: "SELECT COUNT() AS count\n  FROM table\n    WHERE id = 1",
		},
		{
			name:     "Trailing spaces removed",
			input:    "SELECT COUNT() AS count  \nFROM table  \nWHERE id = 1  ",
			expected: "SELECT COUNT() AS count\nFROM table\nWHERE id = 1",
		},
		{
			name:     "Complex query with variables",
			input:    "SELECT COUNT() AS count_of_time\r\nFROM sample.test.test\r\nWHERE time = '$day$'\r\n",
			expected: "SELECT COUNT() AS count_of_time\nFROM sample.test.test\nWHERE time = '$day$'",
		},
		{
			name:     "Trailing newlines",
			input:    "SELECT COUNT() AS count\nFROM table\nWHERE id = 1\n\n",
			expected: "SELECT COUNT() AS count\nFROM table\nWHERE id = 1",
		},
		{
			name:     "Semicolon with newline",
			input:    "SELECT COUNT() AS count\nFROM table\nWHERE id = 1\n;",
			expected: "SELECT COUNT() AS count\nFROM table\nWHERE id = 1;",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := normalizeQuery(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %q, got: %q", tc.expected, result)
			}
		})
	}
}

func TestNormalizeSQLQueryModifier(t *testing.T) {
	testCases := []struct {
		name          string
		stateValue    string
		planValue     string
		expectedValue string
	}{
		{
			name:          "Different line endings",
			stateValue:    "SELECT * FROM table\nWHERE id = 1",
			planValue:     "SELECT * FROM table\r\nWHERE id = 1",
			expectedValue: "SELECT * FROM table\nWHERE id = 1", // State value should be used
		},
		{
			name:          "Different whitespace at line end",
			stateValue:    "SELECT * FROM table\nWHERE id = 1",
			planValue:     "SELECT * FROM table  \nWHERE id = 1  ",
			expectedValue: "SELECT * FROM table\nWHERE id = 1", // State value should be used
		},
		{
			name:          "Different whitespace within line",
			stateValue:    "SELECT * FROM table WHERE id = 1",
			planValue:     "SELECT   *   FROM   table   WHERE   id = 1",
			expectedValue: "SELECT   *   FROM   table   WHERE   id = 1", // Plan value should be used with preserved whitespace
		},
		{
			name:          "Different query",
			stateValue:    "SELECT * FROM table1",
			planValue:     "SELECT * FROM table2",
			expectedValue: "SELECT * FROM table2", // Plan value should be used
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			modifier := NormalizeSQLQuery()

			req := planmodifier.StringRequest{
				StateValue: types.StringValue(tc.stateValue),
				PlanValue:  types.StringValue(tc.planValue),
			}

			resp := planmodifier.StringResponse{
				PlanValue: req.PlanValue,
			}

			modifier.PlanModifyString(ctx, req, &resp)

			if resp.PlanValue.ValueString() != tc.expectedValue {
				t.Errorf("Expected plan value: %q, got: %q", tc.expectedValue, resp.PlanValue.ValueString())
			}
		})
	}
}
