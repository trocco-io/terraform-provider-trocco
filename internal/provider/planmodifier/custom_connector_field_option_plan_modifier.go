package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &CustomConnectorFieldOptionPlanModifier{}

// CustomConnectorFieldOptionPlanModifier validates the shared FieldOption
// shape used by an endpoint's `query_parameters`, `headers` and
// `request_body_parameters` attributes: default_value is required whenever
// is_editable is false.
type CustomConnectorFieldOptionPlanModifier struct{}

func (d *CustomConnectorFieldOptionPlanModifier) Description(ctx context.Context) string {
	return "Validates that default_value is set when is_editable is false"
}

func (d *CustomConnectorFieldOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *CustomConnectorFieldOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var isEditable types.Bool
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("is_editable"), &isEditable)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if isEditable.IsUnknown() || isEditable.ValueBool() {
		return
	}

	var defaultValue types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("default_value"), &defaultValue)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if defaultValue.IsNull() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"CustomConnector FieldOption Validation Error",
			fmt.Sprintf("Attribute %s default_value is required when is_editable is false", req.Path),
		)
	}
}
