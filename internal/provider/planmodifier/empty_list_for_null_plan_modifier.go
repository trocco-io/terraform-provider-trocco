package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.List = emptyListForNullModifier{}

type emptyListForNullModifier struct{}

func (m emptyListForNullModifier) Description(ctx context.Context) string {
	return "Treat null list as an empty list during planning to avoid unnecessary diffs"
}

func (m emptyListForNullModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m emptyListForNullModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if req.PlanValue.IsNull() {
		resp.PlanValue, _ = types.ListValue(req.PlanValue.ElementType(ctx), []attr.Value{})
	}
}

func EmptyListForNull() planmodifier.List {
	return emptyListForNullModifier{}
}
