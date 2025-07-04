package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &CustomVariableSettingPlanModifier{}

type CustomVariableSettingPlanModifier struct{}

func (d *CustomVariableSettingPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating custom_variable_setting"
}

func (d *CustomVariableSettingPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *CustomVariableSettingPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var typ types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("type"), &typ)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var value types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("value"), &value)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var quantity types.Int64
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("quantity"), &quantity)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var unit types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("unit"), &unit)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var direction types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("direction"), &direction)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var format types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("format"), &format)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var timeZone types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("time_zone"), &timeZone)...)
	if resp.Diagnostics.HasError() {
		return
	}

	switch typ.ValueString() {
	case "string":
		{
			if value.IsNull() {
				addCustomVariableSettingAttributeError(req, resp, "value is required for string type")
			}
		}
	case "timestamp", "timestamp_runtime":
		{
			if quantity.IsNull() {
				addCustomVariableSettingAttributeError(req, resp, "quantity is required for timestamp/timestamp_runtime type")
			}
			if unit.IsNull() {
				addCustomVariableSettingAttributeError(req, resp, "unit is required for timestamp/timestamp_runtime type")
			}
			if direction.IsNull() {
				addCustomVariableSettingAttributeError(req, resp, "direction is required for timestamp/timestamp_runtime type")
			}
			if format.IsNull() {
				addCustomVariableSettingAttributeError(req, resp, "format is required for timestamp/timestamp_runtime type")
			}
			if timeZone.IsNull() {
				addCustomVariableSettingAttributeError(req, resp, "time_zone is required for timestamp/timestamp_runtime type")
			}
		}
	}
}

func addCustomVariableSettingAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"CustomVariableSetting Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
