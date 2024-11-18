package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.String = &IgnoreChangesPlanModifier{}

// IgnoreChangesPlanModifier is a plan modifier that ignores changes attribute after creation.
type IgnoreChangesPlanModifier struct{}

func (m IgnoreChangesPlanModifier) Description(ctx context.Context) string {
	return "Ignore changes to attribute after creation."
}

func (m IgnoreChangesPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m IgnoreChangesPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// return if the resource is being created
	if req.State.Raw.IsNull() {
		return
	}

	if req.PlanValue.IsUnknown() || req.StateValue.IsUnknown() || req.PlanValue == req.StateValue {
		return
	}
	resp.PlanValue = req.StateValue
}
