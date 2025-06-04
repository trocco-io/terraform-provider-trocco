package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Set = emptySetForNullModifier{}

type emptySetForNullModifier struct{}

func (m emptySetForNullModifier) Description(ctx context.Context) string {
	return "Treat null set as an empty set during planning to avoid unnecessary diffs"
}

func (m emptySetForNullModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m emptySetForNullModifier) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	if req.PlanValue.IsNull() {
		resp.PlanValue, _ = types.SetValue(req.PlanValue.ElementType(ctx), []attr.Value{})
	}
}

func EmptySetForNull() planmodifier.Set {
	return emptySetForNullModifier{}
}
