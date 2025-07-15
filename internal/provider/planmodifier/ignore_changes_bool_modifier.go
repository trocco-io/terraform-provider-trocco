package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.Bool = &IgnoreChangesBoolPlanModifier{}

// IgnoreChangesBoolPlanModifier is a plan modifier that ignores changes to a bool attribute after creation.
type IgnoreChangesBoolPlanModifier struct{}

func (m IgnoreChangesBoolPlanModifier) Description(ctx context.Context) string {
	return "Ignore changes to boolean attribute after creation."
}

func (m IgnoreChangesBoolPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m IgnoreChangesBoolPlanModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	// Do not ignore during resource creation
	if req.State.Raw.IsNull() {
		return
	}

	// Skip if the value is unknown or there is no change
	if req.PlanValue.IsUnknown() || req.StateValue.IsUnknown() || req.PlanValue == req.StateValue {
		return
	}

	// Overwrite the Plan value with the State value to ignore the difference
	resp.PlanValue = req.StateValue
}
