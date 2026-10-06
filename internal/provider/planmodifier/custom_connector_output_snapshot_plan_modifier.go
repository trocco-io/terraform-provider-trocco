package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// The custom_connector_output_option snapshot attributes (url, auth_type,
// endpoints, ...) are re-taken from the referenced custom connector definition
// on every create/update, so their planned value can only be carried over from
// the prior state while the request is unchanged. `endpoints` in particular
// embeds the values submitted through create_endpoint_settings /
// update_endpoint_settings, so unconditionally reusing the prior state (as
// CustomConnectorSnapshot*PlanModifier does for the transfer source, whose
// snapshots are not derived from the configuration) would plan a value the
// apply then contradicts.
//
// These modifiers therefore reuse the prior state value - including null, which
// the framework's built-in UseStateForUnknown modifiers refuse to do - only
// while every request-affecting attribute matches the prior state, and fall
// back to Unknown ("known after apply") otherwise.
//
// Residual limitation: editing the custom connector definition itself changes
// the snapshot without changing any of these attributes. That alone does not
// trigger an update of the job definition (and thus no re-snapshot), so it is
// only observable when an unrelated job definition attribute changes in the
// same apply.

var _ planmodifier.String = CustomConnectorOutputSnapshotStringPlanModifier{}
var _ planmodifier.List = CustomConnectorOutputSnapshotListPlanModifier{}

const customConnectorOutputSnapshotPlanModifierDescription = "While the request is unchanged, the prior state value of this server-derived attribute is used, including when it is null; otherwise it is re-derived on apply."

// customConnectorOutputOptionRequestAttributes are the attributes of
// custom_connector_output_option whose value can change the snapshot the API
// returns: either they select which endpoints are snapshotted, or they supply
// values that the snapshot embeds.
//
// custom_variable_settings is deliberately absent even though it is part of the
// request, since it does not participate in the snapshot.
var customConnectorOutputOptionRequestAttributes = []string{
	"custom_connector_connection_id",
	"mode",
	"update_key",
	"create_custom_connector_output_endpoint_id",
	"update_custom_connector_output_endpoint_id",
	"create_endpoint_settings",
	"update_endpoint_settings",
}

type CustomConnectorOutputSnapshotStringPlanModifier struct{}

func (m CustomConnectorOutputSnapshotStringPlanModifier) Description(context.Context) string {
	return customConnectorOutputSnapshotPlanModifierDescription
}

func (m CustomConnectorOutputSnapshotStringPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m CustomConnectorOutputSnapshotStringPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.State.Raw.IsNull() {
		return
	}

	if customConnectorOutputSnapshotKeepsState(ctx, req.Plan, req.State, req.Path.ParentPath(), &resp.Diagnostics) {
		resp.PlanValue = req.StateValue
		return
	}
	resp.PlanValue = types.StringUnknown()
}

type CustomConnectorOutputSnapshotListPlanModifier struct{}

func (m CustomConnectorOutputSnapshotListPlanModifier) Description(context.Context) string {
	return customConnectorOutputSnapshotPlanModifierDescription
}

func (m CustomConnectorOutputSnapshotListPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m CustomConnectorOutputSnapshotListPlanModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if req.State.Raw.IsNull() {
		return
	}

	if customConnectorOutputSnapshotKeepsState(ctx, req.Plan, req.State, req.Path.ParentPath(), &resp.Diagnostics) {
		resp.PlanValue = req.StateValue
		return
	}
	resp.PlanValue = types.ListUnknown(req.PlanValue.ElementType(ctx))
}

// customConnectorOutputSnapshotKeepsState reports whether the planned request
// is identical to the one the prior state was produced by, in which case the
// server returns the very same snapshot.
func customConnectorOutputSnapshotKeepsState(
	ctx context.Context,
	plan tfsdk.Plan,
	state tfsdk.State,
	optionPath path.Path,
	diags *diag.Diagnostics,
) bool {
	var planOption, stateOption types.Object
	diags.Append(plan.GetAttribute(ctx, optionPath, &planOption)...)
	diags.Append(state.GetAttribute(ctx, optionPath, &stateOption)...)
	if diags.HasError() || planOption.IsNull() || planOption.IsUnknown() || stateOption.IsNull() || stateOption.IsUnknown() {
		return false
	}

	planAttributes := planOption.Attributes()
	stateAttributes := stateOption.Attributes()
	for _, name := range customConnectorOutputOptionRequestAttributes {
		planValue, ok := planAttributes[name]
		if !ok {
			return false
		}
		stateValue, ok := stateAttributes[name]
		if !ok || !planValue.Equal(stateValue) {
			return false
		}
	}
	return true
}
