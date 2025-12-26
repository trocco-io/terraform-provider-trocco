package pipeline_definition

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestConditionsNotEmpty_ValidateList(t *testing.T) {
	conditionObjectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"variable": types.StringType,
			"task_key": types.StringType,
			"operator": types.StringType,
			"value":    types.StringType,
		},
	}

	tests := []struct {
		name          string
		configValue   types.List
		expectedError bool
		expectedMsg   string
	}{
		{
			name:          "null value - skip validation",
			configValue:   types.ListNull(conditionObjectType),
			expectedError: false,
		},
		{
			name:          "unknown value - skip validation",
			configValue:   types.ListUnknown(conditionObjectType),
			expectedError: false,
		},
		{
			name: "empty list - error",
			configValue: types.ListValueMust(
				conditionObjectType,
				[]attr.Value{},
			),
			expectedError: true,
			expectedMsg:   "At least one condition must be specified",
		},
		{
			name: "one condition - valid",
			configValue: types.ListValueMust(
				conditionObjectType,
				[]attr.Value{
					types.ObjectValueMust(
						conditionObjectType.AttrTypes,
						map[string]attr.Value{
							"variable": types.StringValue("current_time"),
							"task_key": types.StringNull(),
							"operator": types.StringValue("greater"),
							"value":    types.StringValue("2025-01-01T00:00:00Z"),
						},
					),
				},
			),
			expectedError: false,
		},
		{
			name: "multiple conditions - valid",
			configValue: types.ListValueMust(
				conditionObjectType,
				[]attr.Value{
					types.ObjectValueMust(
						conditionObjectType.AttrTypes,
						map[string]attr.Value{
							"variable": types.StringValue("current_time"),
							"task_key": types.StringNull(),
							"operator": types.StringValue("greater"),
							"value":    types.StringValue("2025-01-01T00:00:00Z"),
						},
					),
					types.ObjectValueMust(
						conditionObjectType.AttrTypes,
						map[string]attr.Value{
							"variable": types.StringValue("status"),
							"task_key": types.StringValue("transfer"),
							"operator": types.StringValue("equal"),
							"value":    types.StringValue("succeeded"),
						},
					),
				},
			),
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ConditionsNotEmpty{}
			req := validator.ListRequest{
				ConfigValue: tt.configValue,
				Path:        path.Root("conditions"),
			}
			resp := &validator.ListResponse{}

			v.ValidateList(context.Background(), req, resp)

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
