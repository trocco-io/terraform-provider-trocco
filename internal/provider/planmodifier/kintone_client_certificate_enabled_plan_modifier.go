package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Bool = KintoneClientCertificateEnabledPlanModifier{}

// KintoneClientCertificateEnabledPlanModifier plans the read-only
// `client_certificate_enabled` of a kintone connection: true when
// `client_certificate` is in the configuration, false otherwise.
//
// The attribute is kept, unlike the `ssl` / `gateway` blocks, because it is the
// only certificate-related value the API returns. Read stores it, so a certificate
// enabled outside Terraform (UI, or import without `client_certificate`) shows up
// as `true -> false` in the plan instead of being removed by the next update.
//
// Other connection types are planned as null; see connection_plan_modifier.go for
// why null must be carried forward explicitly.
type KintoneClientCertificateEnabledPlanModifier struct{}

func (m KintoneClientCertificateEnabledPlanModifier) Description(context.Context) string {
	return "For kintone connections, true when client_certificate is set in the configuration and false otherwise; null for other connection types."
}

func (m KintoneClientCertificateEnabledPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m KintoneClientCertificateEnabledPlanModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	var connectionType types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.ParentPath().AtName("connection_type"), &connectionType)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if connectionType.IsUnknown() {
		resp.PlanValue = types.BoolUnknown()
		return
	}
	if connectionType.ValueString() != "kintone" {
		resp.PlanValue = types.BoolNull()
		return
	}

	var certificate types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, req.Path.ParentPath().AtName("client_certificate"), &certificate)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if certificate.IsUnknown() {
		resp.PlanValue = types.BoolUnknown()
		return
	}
	resp.PlanValue = types.BoolValue(!certificate.IsNull())
}
