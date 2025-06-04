package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Int64 = &ConditionalInt64DefaultPlanModifier{}

type ConditionalInt64DefaultPlanModifier struct {
	CondAttrPath path.Path
	TargetValue  string
	DefaultValue int64
}

func (m ConditionalInt64DefaultPlanModifier) Description(ctx context.Context) string {
	return fmt.Sprintf("Set default %d when %s == %q, otherwise null",
		m.DefaultValue, m.CondAttrPath.String(), m.TargetValue)
}

func (m ConditionalInt64DefaultPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

// If the string at CondAttrPath == TargetValue,
//
//	set *this* int64 attribute to DefaultValue (else null).
//
// Handy for cases like:
//
//	pager_type == "offset" → pager_pages = 1
//	otherwise              → pager_pages = null
func (m ConditionalInt64DefaultPlanModifier) PlanModifyInt64(
	ctx context.Context,
	req planmodifier.Int64Request,
	resp *planmodifier.Int64Response,
) {
	if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
		return
	}

	var cond types.String
	req.Config.GetAttribute(ctx, m.CondAttrPath, &cond)

	if cond.ValueString() == m.TargetValue {
		resp.PlanValue = types.Int64Value(m.DefaultValue)
	} else {
		resp.PlanValue = types.Int64Null()
	}
}
