package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// trocco_connection carries a few server-derived attributes (auth_type, scopes,
// authorized) that only apply to some connection types and are therefore null
// for the rest. The framework marks a Computed attribute as Unknown whenever
// its planned value is null, so those attributes show up as
// "(known after apply)" on every plan that touches the connection, even though
// nothing about them can change. The built-in UseStateForUnknown modifiers
// cannot fix that: they bail out precisely when the prior state value is null.
//
// The modifiers below carry the prior value forward instead, including null.

const connectionStateForUnknownDescription = "Once the connection exists, the prior state value of this server-derived attribute is used, including when it is null."

var _ planmodifier.String = ConnectionStateForUnknownStringPlanModifier{}
var _ planmodifier.List = ConnectionStateForUnknownListPlanModifier{}
var _ planmodifier.Bool = ConnectionAuthorizedPlanModifier{}

// ConnectionStateForUnknownStringPlanModifier is used for `auth_type`.
//
// Residual limitation: for a `custom_connector` connection, `auth_type` mirrors
// the referenced custom connector definition, so editing that definition's
// `auth_type` changes this value without changing anything on the connection
// itself. That alone does not trigger an update of the connection (and thus no
// re-read before apply), so it is only observable when an unrelated connection
// attribute changes in the same apply.
type ConnectionStateForUnknownStringPlanModifier struct{}

func (m ConnectionStateForUnknownStringPlanModifier) Description(context.Context) string {
	return connectionStateForUnknownDescription
}

func (m ConnectionStateForUnknownStringPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m ConnectionStateForUnknownStringPlanModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.State.Raw.IsNull() || !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	resp.PlanValue = req.StateValue
}

// ConnectionStateForUnknownListPlanModifier is used for `scopes`. Omitting the
// attribute means "keep the scopes stored on the server", so the prior state is
// always the right planned value.
type ConnectionStateForUnknownListPlanModifier struct{}

func (m ConnectionStateForUnknownListPlanModifier) Description(context.Context) string {
	return connectionStateForUnknownDescription
}

func (m ConnectionStateForUnknownListPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m ConnectionStateForUnknownListPlanModifier) PlanModifyList(_ context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if req.State.Raw.IsNull() || !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	resp.PlanValue = req.StateValue
}

// connectionAuthorizedResetAttributes are the attributes whose modification
// makes the server drop `authorized` back to false, because the previously
// verified credentials no longer apply.
var connectionAuthorizedResetAttributes = []string{
	"oauth2_client_id",
	"oauth2_client_secret",
}

// ConnectionAuthorizedPlanModifier is used for `authorized`.
//
// Unlike the modifiers above it cannot reuse the prior state unconditionally:
// the server resets `authorized` to false whenever the OAuth2 credentials
// change, so a plan that carried a prior `true` forward would be contradicted
// by the apply. It therefore reuses the prior value - including null, which the
// built-in UseStateForUnknown refuses to do - only while those credentials are
// untouched, and falls back to Unknown otherwise.
type ConnectionAuthorizedPlanModifier struct{}

func (m ConnectionAuthorizedPlanModifier) Description(context.Context) string {
	return "While the OAuth2 credentials are unchanged, the prior state value is used, including when it is null; otherwise it is re-derived on apply."
}

func (m ConnectionAuthorizedPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m ConnectionAuthorizedPlanModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if req.State.Raw.IsNull() {
		return
	}

	if connectionOAuth2CredentialsUnchanged(ctx, req.Plan, req.State, &resp.Diagnostics) {
		resp.PlanValue = req.StateValue
		return
	}
	resp.PlanValue = types.BoolUnknown()
}

// connectionOAuth2CredentialsUnchanged reports whether the planned OAuth2
// credentials are identical to the ones the prior state was produced with, in
// which case the server leaves `authorized` as it is.
func connectionOAuth2CredentialsUnchanged(
	ctx context.Context,
	plan tfsdk.Plan,
	state tfsdk.State,
	diags *diag.Diagnostics,
) bool {
	for _, name := range connectionAuthorizedResetAttributes {
		var planValue, stateValue types.String
		diags.Append(plan.GetAttribute(ctx, path.Root(name), &planValue)...)
		diags.Append(state.GetAttribute(ctx, path.Root(name), &stateValue)...)
		if diags.HasError() {
			return false
		}
		if !planValue.Equal(stateValue) {
			return false
		}
	}
	return true
}
