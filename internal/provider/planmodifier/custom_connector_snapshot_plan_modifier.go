package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// The custom_connector_input_option snapshot attributes (url, auth_type,
// endpoint_request_body, paginator, etc.) are entirely server-derived and
// can legitimately be null (e.g. an endpoint with no request body). The
// framework's built-in UseStateForUnknown modifiers leave the planned value
// Unknown whenever the prior state value is null, which would make such
// attributes show as "(known after apply)" on every plan even though
// nothing changed. These modifiers instead always carry the prior state
// value forward (including null) once the resource exists, matching the
// snapshot semantics of these attributes.

var _ planmodifier.String = CustomConnectorSnapshotStringPlanModifier{}
var _ planmodifier.Int64 = CustomConnectorSnapshotInt64PlanModifier{}
var _ planmodifier.Bool = CustomConnectorSnapshotBoolPlanModifier{}
var _ planmodifier.Object = CustomConnectorSnapshotObjectPlanModifier{}
var _ planmodifier.List = CustomConnectorSnapshotListPlanModifier{}

const customConnectorSnapshotPlanModifierDescription = "Once the resource exists, the prior state value of this server-derived attribute is always used, including when it is null."

type CustomConnectorSnapshotStringPlanModifier struct{}

func (m CustomConnectorSnapshotStringPlanModifier) Description(context.Context) string {
	return customConnectorSnapshotPlanModifierDescription
}

func (m CustomConnectorSnapshotStringPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m CustomConnectorSnapshotStringPlanModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.State.Raw.IsNull() || !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	resp.PlanValue = req.StateValue
}

type CustomConnectorSnapshotInt64PlanModifier struct{}

func (m CustomConnectorSnapshotInt64PlanModifier) Description(context.Context) string {
	return customConnectorSnapshotPlanModifierDescription
}

func (m CustomConnectorSnapshotInt64PlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m CustomConnectorSnapshotInt64PlanModifier) PlanModifyInt64(_ context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if req.State.Raw.IsNull() || !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	resp.PlanValue = req.StateValue
}

type CustomConnectorSnapshotBoolPlanModifier struct{}

func (m CustomConnectorSnapshotBoolPlanModifier) Description(context.Context) string {
	return customConnectorSnapshotPlanModifierDescription
}

func (m CustomConnectorSnapshotBoolPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m CustomConnectorSnapshotBoolPlanModifier) PlanModifyBool(_ context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	if req.State.Raw.IsNull() || !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	resp.PlanValue = req.StateValue
}

type CustomConnectorSnapshotObjectPlanModifier struct{}

func (m CustomConnectorSnapshotObjectPlanModifier) Description(context.Context) string {
	return customConnectorSnapshotPlanModifierDescription
}

func (m CustomConnectorSnapshotObjectPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m CustomConnectorSnapshotObjectPlanModifier) PlanModifyObject(_ context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if req.State.Raw.IsNull() || !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	resp.PlanValue = req.StateValue
}

type CustomConnectorSnapshotListPlanModifier struct{}

func (m CustomConnectorSnapshotListPlanModifier) Description(context.Context) string {
	return customConnectorSnapshotPlanModifierDescription
}

func (m CustomConnectorSnapshotListPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m CustomConnectorSnapshotListPlanModifier) PlanModifyList(_ context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if req.State.Raw.IsNull() || !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}
	resp.PlanValue = req.StateValue
}
