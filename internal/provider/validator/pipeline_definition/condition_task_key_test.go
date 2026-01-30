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

func TestConditionTaskKey_ValidateObject(t *testing.T) {
	conditionObjectType := map[string]attr.Type{
		"variable": types.StringType,
		"task_key": types.StringType,
		"operator": types.StringType,
		"value":    types.StringType,
	}

	tests := []struct {
		name          string
		configValue   types.Object
		expectedError bool
		expectedMsg   string
	}{
		{
			name:          "null value - skip validation",
			configValue:   types.ObjectNull(conditionObjectType),
			expectedError: false,
		},
		{
			name:          "unknown value - skip validation",
			configValue:   types.ObjectUnknown(conditionObjectType),
			expectedError: false,
		},
		{
			name: "current_time without task_key - valid",
			configValue: types.ObjectValueMust(
				conditionObjectType,
				map[string]attr.Value{
					"variable": types.StringValue("current_time"),
					"task_key": types.StringNull(),
					"operator": types.StringValue("greater"),
					"value":    types.StringValue("2025-01-01T00:00:00Z"),
				},
			),
			expectedError: false,
		},
		{
			name: "environment without task_key - valid",
			configValue: types.ObjectValueMust(
				conditionObjectType,
				map[string]attr.Value{
					"variable": types.StringValue("environment"),
					"task_key": types.StringNull(),
					"operator": types.StringValue("equal"),
					"value":    types.StringValue("production"),
				},
			),
			expectedError: false,
		},
		{
			name: "status with task_key - valid",
			configValue: types.ObjectValueMust(
				conditionObjectType,
				map[string]attr.Value{
					"variable": types.StringValue("status"),
					"task_key": types.StringValue("transfer"),
					"operator": types.StringValue("equal"),
					"value":    types.StringValue("succeeded"),
				},
			),
			expectedError: false,
		},
		{
			name: "status without task_key - error",
			configValue: types.ObjectValueMust(
				conditionObjectType,
				map[string]attr.Value{
					"variable": types.StringValue("status"),
					"task_key": types.StringNull(),
					"operator": types.StringValue("equal"),
					"value":    types.StringValue("succeeded"),
				},
			),
			expectedError: true,
			expectedMsg:   "task_key is required when variable is 'status'",
		},
		{
			name: "response_status_code without task_key - error",
			configValue: types.ObjectValueMust(
				conditionObjectType,
				map[string]attr.Value{
					"variable": types.StringValue("response_status_code"),
					"task_key": types.StringNull(),
					"operator": types.StringValue("equal"),
					"value":    types.StringValue("200"),
				},
			),
			expectedError: true,
			expectedMsg:   "task_key is required when variable is 'response_status_code'",
		},
		{
			name: "transfer_record_count without task_key - error",
			configValue: types.ObjectValueMust(
				conditionObjectType,
				map[string]attr.Value{
					"variable": types.StringValue("transfer_record_count"),
					"task_key": types.StringNull(),
					"operator": types.StringValue("greater"),
					"value":    types.StringValue("0"),
				},
			),
			expectedError: true,
			expectedMsg:   "task_key is required when variable is 'transfer_record_count'",
		},
		{
			name: "check_result without task_key - error",
			configValue: types.ObjectValueMust(
				conditionObjectType,
				map[string]attr.Value{
					"variable": types.StringValue("check_result"),
					"task_key": types.StringNull(),
					"operator": types.StringValue("equal"),
					"value":    types.StringValue("true"),
				},
			),
			expectedError: true,
			expectedMsg:   "task_key is required when variable is 'check_result'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ConditionTaskKey{}
			req := validator.ObjectRequest{
				ConfigValue: tt.configValue,
				Path:        path.Root("condition"),
			}
			resp := &validator.ObjectResponse{}

			v.ValidateObject(context.Background(), req, resp)

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
