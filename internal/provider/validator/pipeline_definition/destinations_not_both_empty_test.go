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

func TestDestinationsNotBothEmpty_ValidateObject(t *testing.T) {
	destinationsObjectType := map[string]attr.Type{
		"if":   types.ListType{ElemType: types.StringType},
		"else": types.ListType{ElemType: types.StringType},
	}

	tests := []struct {
		name          string
		configValue   types.Object
		expectedError bool
		expectedMsg   string
	}{
		{
			name:          "null value - skip validation",
			configValue:   types.ObjectNull(destinationsObjectType),
			expectedError: false,
		},
		{
			name:          "unknown value - skip validation",
			configValue:   types.ObjectUnknown(destinationsObjectType),
			expectedError: false,
		},
		{
			name: "both empty - error",
			configValue: types.ObjectValueMust(
				destinationsObjectType,
				map[string]attr.Value{
					"if":   types.ListValueMust(types.StringType, []attr.Value{}),
					"else": types.ListValueMust(types.StringType, []attr.Value{}),
				},
			),
			expectedError: true,
			expectedMsg:   "At least one destination must be specified",
		},
		{
			name: "if has value, else empty - valid",
			configValue: types.ObjectValueMust(
				destinationsObjectType,
				map[string]attr.Value{
					"if":   types.ListValueMust(types.StringType, []attr.Value{types.StringValue("task1")}),
					"else": types.ListValueMust(types.StringType, []attr.Value{}),
				},
			),
			expectedError: false,
		},
		{
			name: "if empty, else has value - valid",
			configValue: types.ObjectValueMust(
				destinationsObjectType,
				map[string]attr.Value{
					"if":   types.ListValueMust(types.StringType, []attr.Value{}),
					"else": types.ListValueMust(types.StringType, []attr.Value{types.StringValue("task1")}),
				},
			),
			expectedError: false,
		},
		{
			name: "both have values - valid",
			configValue: types.ObjectValueMust(
				destinationsObjectType,
				map[string]attr.Value{
					"if":   types.ListValueMust(types.StringType, []attr.Value{types.StringValue("task1")}),
					"else": types.ListValueMust(types.StringType, []attr.Value{types.StringValue("task2")}),
				},
			),
			expectedError: false,
		},
		{
			name: "if has multiple values - valid",
			configValue: types.ObjectValueMust(
				destinationsObjectType,
				map[string]attr.Value{
					"if": types.ListValueMust(types.StringType, []attr.Value{
						types.StringValue("task1"),
						types.StringValue("task2"),
					}),
					"else": types.ListValueMust(types.StringType, []attr.Value{}),
				},
			),
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := DestinationsNotBothEmpty{}
			req := validator.ObjectRequest{
				ConfigValue: tt.configValue,
				Path:        path.Root("destinations"),
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
