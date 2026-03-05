package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.List = &ConditionalEmptyListDefaultPlanModifier{}

type ConditionalEmptyListDefaultPlanModifier struct {
	ConnectionTypes   []string
	ConnectionTypeKey string
}

func (d *ConditionalEmptyListDefaultPlanModifier) Description(ctx context.Context) string {
	return "Applies empty list default only when connection_type matches specified types"
}

func (d *ConditionalEmptyListDefaultPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *ConditionalEmptyListDefaultPlanModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	if !req.PlanValue.IsNull() && !req.PlanValue.IsUnknown() {
		return
	}

	connectionTypeKey := d.ConnectionTypeKey
	if connectionTypeKey == "" {
		connectionTypeKey = "connection_type"
	}

	var connectionType types.String
	connectionTypePath := req.Path.ParentPath().AtName(connectionTypeKey)
	diags := req.Plan.GetAttribute(ctx, connectionTypePath, &connectionType)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	currentType := connectionType.ValueString()
	for _, validType := range d.ConnectionTypes {
		if currentType == validType {
			resp.PlanValue, _ = types.ListValue(req.PlanValue.ElementType(ctx), []attr.Value{})
			return
		}
	}
}

func ConditionalEmptyListDefault(connectionTypes ...string) planmodifier.List {
	return &ConditionalEmptyListDefaultPlanModifier{
		ConnectionTypes:   connectionTypes,
		ConnectionTypeKey: "connection_type",
	}
}
