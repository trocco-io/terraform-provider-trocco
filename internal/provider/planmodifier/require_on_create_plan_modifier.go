package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.String = &RequiredOnCreatePlanModifier{}

// RequiredOnCreatePlanModifier is a plan modifier that ensures that a field is only required on resource creation.
type RequiredOnCreatePlanModifier struct {
	AttributeName string
}

func (m RequiredOnCreatePlanModifier) Description(ctx context.Context) string {
	return "This field is required on resource creation."
}

func (m RequiredOnCreatePlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m RequiredOnCreatePlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Check if the resource is being created.
	if !req.State.Raw.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() || req.ConfigValue.ValueString() == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root(m.AttributeName),
			"Missing Required Attribute",
			"The attribute '"+m.AttributeName+"' is required on resource creation.",
		)
	}
}
