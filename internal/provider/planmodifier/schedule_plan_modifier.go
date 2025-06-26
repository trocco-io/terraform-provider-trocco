package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &SchedulePlanModifier{}

type SchedulePlanModifier struct{}

func (d *SchedulePlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating schedule attributes"
}

func (d *SchedulePlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *SchedulePlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var frequency types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("frequency"), &frequency)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var hour types.Int64
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("hour"), &hour)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var dayOfWeek types.Int64
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("day_of_week"), &dayOfWeek)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var day types.Int64
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("day"), &day)...)
	if resp.Diagnostics.HasError() {
		return
	}

	switch frequency.ValueString() {
	case "daily":
		{
			if hour.IsNull() {
				addScheduleAttributeError(req, resp, "hour is required for daily frequency")
			}
		}
	case "weekly":
		{
			if hour.IsNull() {
				addScheduleAttributeError(req, resp, "hour is required for weekly frequency")
			}
			if dayOfWeek.IsNull() {
				addScheduleAttributeError(req, resp, "day_of_week is required for weekly frequency")
			}
		}
	case "monthly":
		{
			if hour.IsNull() {
				addScheduleAttributeError(req, resp, "hour is required for monthly frequency")
			}
			if day.IsNull() {
				addScheduleAttributeError(req, resp, "day is required for monthly frequency")
			}
		}
	}

}

func addScheduleAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Schedule Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
