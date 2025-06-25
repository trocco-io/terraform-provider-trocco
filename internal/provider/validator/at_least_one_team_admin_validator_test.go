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

func TestAtLeastOneTeamAdminValidator_ValidateSet(t *testing.T) {
	memberObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"user_id": types.Int64Type,
			"role":    types.StringType,
		},
	}

	tests := []struct {
		name          string
		configValue   types.Set
		expectedError bool
		expectedMsg   string
	}{
		{
			name:          "unknown value - skip validation",
			configValue:   types.SetUnknown(memberObjectType),
			expectedError: false,
		},
		{
			name:          "null value - error",
			configValue:   types.SetNull(memberObjectType),
			expectedError: true,
			expectedMsg:   "The `members` list cannot be null",
		},
		{
			name: "empty list - error",
			configValue: types.SetValueMust(
				memberObjectType,
				[]attr.Value{},
			),
			expectedError: true,
			expectedMsg:   "The `members` list cannot be empty",
		},
		{
			name: "has team_admin - valid",
			configValue: types.SetValueMust(
				memberObjectType,
				[]attr.Value{
					types.ObjectValueMust(
						memberObjectType.AttrTypes,
						map[string]attr.Value{
							"user_id": types.Int64Value(1),
							"role":    types.StringValue("team_admin"),
						},
					),
					types.ObjectValueMust(
						memberObjectType.AttrTypes,
						map[string]attr.Value{
							"user_id": types.Int64Value(2),
							"role":    types.StringValue("member"),
						},
					),
				},
			),
			expectedError: false,
		},
		{
			name: "no team_admin - error",
			configValue: types.SetValueMust(
				memberObjectType,
				[]attr.Value{
					types.ObjectValueMust(
						memberObjectType.AttrTypes,
						map[string]attr.Value{
							"user_id": types.Int64Value(1),
							"role":    types.StringValue("member"),
						},
					),
					types.ObjectValueMust(
						memberObjectType.AttrTypes,
						map[string]attr.Value{
							"user_id": types.Int64Value(2),
							"role":    types.StringValue("member"),
						},
					),
				},
			),
			expectedError: true,
			expectedMsg:   "must include at least one member with `role` set to `team_admin`",
		},
		{
			name: "unknown role - skip error",
			configValue: types.SetValueMust(
				memberObjectType,
				[]attr.Value{
					types.ObjectValueMust(
						memberObjectType.AttrTypes,
						map[string]attr.Value{
							"user_id": types.Int64Value(1),
							"role":    types.StringUnknown(),
						},
					),
					types.ObjectValueMust(
						memberObjectType.AttrTypes,
						map[string]attr.Value{
							"user_id": types.Int64Value(2),
							"role":    types.StringValue("member"),
						},
					),
				},
			),
			expectedError: false, // Don't error when there are unknown roles
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := AtLeastOneTeamAdminValidator{}
			req := validator.SetRequest{
				ConfigValue: tt.configValue,
				Path:        path.Root("members"),
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
