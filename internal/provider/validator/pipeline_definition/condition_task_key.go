package pipeline_definition

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
)

var _ validator.Object = ConditionTaskKey{}

// taskScopedVariables are variables that require a task_key to be specified
var taskScopedVariables = []string{
	"status",
	"response_status_code",
	"transfer_record_count",
	"check_result",
}

type ConditionTaskKey struct {
}

func (v ConditionTaskKey) Description(ctx context.Context) string {
	return "Ensures task_key is specified when variable is a task-scoped variable (status, response_status_code, transfer_record_count, check_result)."
}

func (v ConditionTaskKey) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ConditionTaskKey) ValidateObject(
	ctx context.Context,
	req validator.ObjectRequest,
	resp *validator.ObjectResponse,
) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	attrs := req.ConfigValue.Attributes()

	variableAttr, ok := attrs["variable"].(types.String)
	if !ok || variableAttr.IsNull() || variableAttr.IsUnknown() {
		return
	}

	variable := variableAttr.ValueString()

	// Check if this is a task-scoped variable
	if !lo.Contains(taskScopedVariables, variable) {
		return
	}

	// For task-scoped variables, task_key is required
	taskKeyAttr, ok := attrs["task_key"].(types.String)
	if !ok || taskKeyAttr.IsNull() || taskKeyAttr.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			req.Path.AtName("task_key"),
			"Missing Required Attribute",
			fmt.Sprintf("task_key is required when variable is '%s'. Please specify the task_key of the task to check.", variable),
		)
	}
}
