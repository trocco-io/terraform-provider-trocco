package planmodifier

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &FilterColumnPlanModifier{}

type FilterColumnPlanModifier struct{}

func (d *FilterColumnPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating filter column attributes"
}

func (d *FilterColumnPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *FilterColumnPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var jsonExpandEnabled types.Bool
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("json_expand_enabled"), &jsonExpandEnabled)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var typeProp types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("type"), &typeProp)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var jsonExpandKeepBaseColumn types.Bool
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("json_expand_keep_base_column"), &jsonExpandKeepBaseColumn)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var jsonExpandColumns types.List
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("json_expand_columns"), &jsonExpandColumns)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if jsonExpandEnabled.ValueBool() && typeProp.ValueString() != "json" {
		addFilterColumnAttributeError(req, resp, "If json_expand_enabled is true, type must be json.")
	}

	if !jsonExpandEnabled.ValueBool() && jsonExpandKeepBaseColumn.ValueBool() {
		addFilterColumnAttributeError(req, resp, "If json_expand_enabled is false, json_expand_keep_base_column must be false.")
	}

	if !jsonExpandEnabled.ValueBool() && !jsonExpandColumns.IsNull() {
		addFilterColumnAttributeError(req, resp, "If json_expand_enabled is false, json_expand_columns must be null.")
	}

	if jsonExpandEnabled.ValueBool() && len(jsonExpandColumns.Elements()) < 1 {
		addFilterColumnAttributeError(req, resp, "If json_expand_enabled is true, json_expand_columns must not be empty.")
	}
}

func addFilterColumnAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"FilterColumn Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
