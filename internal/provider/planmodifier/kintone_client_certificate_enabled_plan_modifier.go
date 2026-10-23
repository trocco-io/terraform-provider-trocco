package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Bool = KintoneClientCertificateEnabledPlanModifier{}

// KintoneClientCertificateEnabledPlanModifier derives the read-only
// `client_certificate_enabled` of a kintone connection from whether
// `client_certificate` is set in the configuration.
//
// The attribute is kept (rather than dropping it like the `ssl` / `gateway`
// blocks do) because it is the only certificate-related value the API returns:
// Read stores the API value, so when the certificate was enabled outside
// Terraform (in the TROCCO UI, or on a connection imported without
// `client_certificate` in the configuration) the plan shows
// `client_certificate_enabled: true -> false` instead of silently removing the
// certificate on the next unrelated update.
//
// For other connection types the planned value is null; carrying null forward
// explicitly avoids "(known after apply)" on every plan, see the comment at the
// top of connection_plan_modifier.go.
//
// When `connection_type` itself is unknown, the value is planned as unknown
// ("known after apply"), as it is for an unknown `client_certificate`.
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
