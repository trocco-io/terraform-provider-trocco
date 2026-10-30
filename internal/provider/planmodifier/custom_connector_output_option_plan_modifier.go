package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &CustomConnectorOutputOptionPlanModifier{}

// CustomConnectorOutputOptionPlanModifier enforces the API's mode-dependent
// rules for custom_connector_output_option at plan time. The server accepts
// `update_key` / `update_custom_connector_output_endpoint_id` only while mode
// is `upsert`, and requires both in that case.
type CustomConnectorOutputOptionPlanModifier struct{}

func (d *CustomConnectorOutputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating custom connector output option attributes"
}

func (d *CustomConnectorOutputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *CustomConnectorOutputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	// mode is read from the plan so the schema default (`insert`) applies when
	// the configuration omits it; the update-only attributes are read from the
	// configuration, since the rules are about what the user specified.
	var mode types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("mode"), &mode)...)

	var updateKey types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.AtName("update_key"), &updateKey)...)

	var updateEndpointID types.Int64
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.AtName("update_custom_connector_output_endpoint_id"), &updateEndpointID)...)

	var updateEndpointSettings types.Object
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.AtName("update_endpoint_settings"), &updateEndpointSettings)...)

	if resp.Diagnostics.HasError() || mode.IsUnknown() {
		return
	}

	switch mode.ValueString() {
	case "upsert":
		if updateKey.IsNull() {
			addCustomConnectorOutputOptionAttributeError(req, resp, "update_key is required when mode is 'upsert'")
		}
		if updateEndpointID.IsNull() {
			addCustomConnectorOutputOptionAttributeError(req, resp, "update_custom_connector_output_endpoint_id is required when mode is 'upsert'")
		}
	default:
		// A null mode is treated as `insert`: that is the schema default, and
		// the server applies the same fallback.
		if !updateKey.IsNull() {
			addCustomConnectorOutputOptionAttributeError(req, resp, "update_key must not be specified when mode is 'insert'")
		}
		if !updateEndpointID.IsNull() {
			addCustomConnectorOutputOptionAttributeError(req, resp, "update_custom_connector_output_endpoint_id must not be specified when mode is 'insert'")
		}
		if !updateEndpointSettings.IsNull() {
			addCustomConnectorOutputOptionAttributeError(req, resp, "update_endpoint_settings must not be specified when mode is 'insert'")
		}
	}
}

func addCustomConnectorOutputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"CustomConnector OutputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
