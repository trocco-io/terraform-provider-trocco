package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &DatabricksOutputOptionPlanModifier{}

type DatabricksOutputOptionPlanModifier struct{}

func (d *DatabricksOutputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating databricks output option attributes"
}

func (d *DatabricksOutputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *DatabricksOutputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var mode types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("mode"), &mode)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var databricksOutputOptionMergeKeys types.Set
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("databricks_output_option_merge_keys"), &databricksOutputOptionMergeKeys)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if mode.ValueString() != "merge" && len(databricksOutputOptionMergeKeys.Elements()) > 0 {
		addDatabricksOutputOptionAttributeError(req, resp, "databricks_output_option_merge_keys can only be set when mode is 'merge'")
	}
}

func addDatabricksOutputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Databricks Output Option Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
