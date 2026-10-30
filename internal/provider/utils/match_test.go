package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type matchElem struct {
	typ string
	dst string
	id  int
}

func matchPrimaryKey(e matchElem) string {
	return fmt.Sprintf("%s|%s|%d", e.typ, e.dst, e.id)
}

func matchFallbackKey(e matchElem) string {
	return fmt.Sprintf("%s|%s", e.typ, e.dst)
}

func TestMatchByKey(t *testing.T) {
	tests := []struct {
		name     string
		api      []matchElem
		ref      []matchElem
		fallback func(matchElem) string
		expected []matchElem
	}{
		{
			name:     "empty ref returns api unchanged",
			api:      []matchElem{{"job", "slack", 1}, {"job", "email", 2}},
			ref:      []matchElem{},
			fallback: matchFallbackKey,
			expected: []matchElem{{"job", "slack", 1}, {"job", "email", 2}},
		},
		{
			name:     "empty api with non-empty ref returns empty",
			api:      []matchElem{},
			ref:      []matchElem{{"job", "slack", 1}},
			fallback: matchFallbackKey,
			expected: []matchElem{},
		},
		{
			name:     "primary key reorders api to ref order",
			api:      []matchElem{{"job", "email", 2}, {"job", "slack", 1}},
			ref:      []matchElem{{"job", "slack", 1}, {"job", "email", 2}},
			fallback: matchFallbackKey,
			expected: []matchElem{{"job", "slack", 1}, {"job", "email", 2}},
		},
		{
			name:     "fallback key matches positionally when ref ids are unknown (CREATE)",
			api:      []matchElem{{"job", "email", 100}, {"job", "slack", 200}},
			ref:      []matchElem{{"job", "slack", 0}, {"job", "email", 0}},
			fallback: matchFallbackKey,
			expected: []matchElem{{"job", "slack", 200}, {"job", "email", 100}},
		},
		{
			name:     "mixed primary and fallback matches",
			api:      []matchElem{{"a", "y", 3}, {"a", "x", 1}, {"a", "x", 2}},
			ref:      []matchElem{{"a", "x", 1}, {"a", "x", 0}, {"a", "y", 3}},
			fallback: matchFallbackKey,
			expected: []matchElem{{"a", "x", 1}, {"a", "x", 2}, {"a", "y", 3}},
		},
		{
			name:     "nil fallback skips unmatched refs and appends leftovers",
			api:      []matchElem{{"a", "x", 1}, {"b", "y", 2}},
			ref:      []matchElem{{"a", "x", 1}, {"c", "z", 0}},
			fallback: nil,
			expected: []matchElem{{"a", "x", 1}, {"b", "y", 2}},
		},
		{
			name:     "api elements not in ref are appended in original order",
			api:      []matchElem{{"a", "x", 1}, {"b", "y", 2}, {"c", "z", 3}},
			ref:      []matchElem{{"b", "y", 2}},
			fallback: matchFallbackKey,
			expected: []matchElem{{"b", "y", 2}, {"a", "x", 1}, {"c", "z", 3}},
		},
		{
			name:     "duplicate fallback keys are consumed positionally",
			api:      []matchElem{{"a", "x", 1}, {"a", "x", 2}},
			ref:      []matchElem{{"a", "x", 0}, {"a", "x", 0}},
			fallback: matchFallbackKey,
			expected: []matchElem{{"a", "x", 1}, {"a", "x", 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MatchByKey(tt.api, tt.ref, matchPrimaryKey, tt.fallback)
			assert.Equal(t, tt.expected, got)
		})
	}
}
