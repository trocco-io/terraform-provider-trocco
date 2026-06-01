package validator

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestUniqueTeamValidator_ValidateSet(t *testing.T) {
	teamObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"team_id": types.Int64Type,
			"role":    types.StringType,
		},
	}

	team := func(id attr.Value, role string) attr.Value {
		return types.ObjectValueMust(
			teamObjectType.AttrTypes,
			map[string]attr.Value{
				"team_id": id,
				"role":    types.StringValue(role),
			},
		)
	}

	tests := []struct {
		name          string
		configValue   types.Set
		expectedError bool
		expectedMsg   string
	}{
		{
			name:          "unknown set - skip validation",
			configValue:   types.SetUnknown(teamObjectType),
			expectedError: false,
		},
		{
			name:          "null set - skip validation",
			configValue:   types.SetNull(teamObjectType),
			expectedError: false,
		},
		{
			name: "empty list - no error",
			configValue: types.SetValueMust(
				teamObjectType,
				[]attr.Value{},
			),
			expectedError: false,
		},
		{
			name: "all unique team IDs - no error",
			configValue: types.SetValueMust(
				teamObjectType,
				[]attr.Value{
					team(types.Int64Value(1), "administrator"),
					team(types.Int64Value(2), "operator"),
				},
			),
			expectedError: false,
		},
		{
			name: "duplicate known team IDs - error",
			configValue: types.SetValueMust(
				teamObjectType,
				[]attr.Value{
					team(types.Int64Value(1), "administrator"),
					team(types.Int64Value(1), "operator"),
				},
			),
			expectedError: true,
			expectedMsg:   "is duplicated in the list",
		},
		{
			name: "all unknown team IDs - skip pair comparison",
			configValue: types.SetValueMust(
				teamObjectType,
				[]attr.Value{
					team(types.Int64Unknown(), "administrator"),
					team(types.Int64Unknown(), "operator"),
				},
			),
			expectedError: false,
		},
		{
			name: "mix of known and unknown team IDs - no error when uniques differ",
			configValue: types.SetValueMust(
				teamObjectType,
				[]attr.Value{
					team(types.Int64Value(1), "administrator"),
					team(types.Int64Unknown(), "operator"),
				},
			),
			expectedError: false,
		},
		{
			name: "null team_id in one element - skip that pair",
			configValue: types.SetValueMust(
				teamObjectType,
				[]attr.Value{
					team(types.Int64Value(1), "administrator"),
					team(types.Int64Null(), "operator"),
				},
			),
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := UniqueTeamValidator{}
			req := validator.SetRequest{
				ConfigValue: tt.configValue,
				Path:        path.Root("teams"),
			}
			resp := &validator.SetResponse{}

			v.ValidateSet(context.Background(), req, resp)

			if tt.expectedError {
				assert.True(t, resp.Diagnostics.HasError(), "expected error but got none")
				if tt.expectedMsg != "" {
					assert.Contains(t, resp.Diagnostics.Errors()[0].Detail(), tt.expectedMsg)
				}
			} else {
				assert.False(t, resp.Diagnostics.HasError(), "unexpected error: %v", resp.Diagnostics)
			}
		})
	}
}
