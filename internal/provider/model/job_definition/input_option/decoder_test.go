package input_options

import (
	"testing"

	jobDefinitionEntities "terraform-provider-trocco/internal/client/entity/job_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestNewDecoder(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		decoder  *jobDefinitionEntities.Decoder
		previous *Decoder
		expected *Decoder
	}{
		{
			name:     "returns nil when api decoder is nil and previous is nil",
			decoder:  nil,
			previous: nil,
			expected: nil,
		},
		{
			name:     "keeps previous decoder with empty match_name when api decoder is nil",
			decoder:  nil,
			previous: &Decoder{MatchName: types.StringValue("")},
			expected: &Decoder{MatchName: types.StringValue("")},
		},
		{
			name:     "keeps previous decoder with null match_name when api decoder is nil",
			decoder:  nil,
			previous: &Decoder{MatchName: types.StringNull()},
			expected: &Decoder{MatchName: types.StringNull()},
		},
		{
			name:     "returns nil when api decoder is nil and previous match_name is not empty",
			decoder:  nil,
			previous: &Decoder{MatchName: types.StringValue("foo")},
			expected: nil,
		},
		{
			name:     "returns api decoder when it is not nil",
			decoder:  &jobDefinitionEntities.Decoder{MatchName: "foo"},
			previous: nil,
			expected: &Decoder{MatchName: types.StringValue("foo")},
		},
		{
			name:     "returns api decoder even when previous differs",
			decoder:  &jobDefinitionEntities.Decoder{MatchName: "foo"},
			previous: &Decoder{MatchName: types.StringValue("")},
			expected: &Decoder{MatchName: types.StringValue("foo")},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, c.expected, NewDecoder(c.decoder, c.previous))
		})
	}
}
