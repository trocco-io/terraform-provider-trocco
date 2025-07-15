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
	// リソース作成時は無視しない
	if req.State.Raw.IsNull() {
		return
	}

	// unknown値または変化なしならスキップ
	if req.PlanValue.IsUnknown() || req.StateValue.IsUnknown() || req.PlanValue == req.StateValue {
		return
	}

	// Plan値をStateの値で上書きし、差分を無視する
	resp.PlanValue = req.StateValue
}
