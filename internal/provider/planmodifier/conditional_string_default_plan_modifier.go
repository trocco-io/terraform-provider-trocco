package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.String = &ConditionalStringDefaultPlanModifier{}

type ConditionalStringDefaultPlanModifier struct {
	DefaultValue      string
	ConnectionTypes   []string // Connection types where this default should apply
	ConnectionTypeKey string   // Key to check for connection type (default: "connection_type")
}

func (d *ConditionalStringDefaultPlanModifier) Description(ctx context.Context) string {
	return "Applies default value for string fields only when connection_type matches specified types"
}

func (d *ConditionalStringDefaultPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *ConditionalStringDefaultPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Don't modify if there's already a planned value
	if !req.PlanValue.IsNull() && !req.PlanValue.IsUnknown() {
		return
	}

	// Get the connection_type key (default to "connection_type")
	connectionTypeKey := d.ConnectionTypeKey
	if connectionTypeKey == "" {
		connectionTypeKey = "connection_type"
	}

	// Get the connection_type from the plan
	var connectionType types.String
	connectionTypePath := req.Path.ParentPath().AtName(connectionTypeKey)
	diags := req.Plan.GetAttribute(ctx, connectionTypePath, &connectionType)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if current connection type matches any of the specified types
	currentType := connectionType.ValueString()
	shouldApplyDefault := false
	for _, validType := range d.ConnectionTypes {
		if currentType == validType {
			shouldApplyDefault = true
			break
		}
	}

	// Only set default value if connection type matches
	if shouldApplyDefault {
		resp.PlanValue = types.StringValue(d.DefaultValue)
	}
}

// ConditionalStringDefault creates a plan modifier that applies a string default for specific connection types.
func ConditionalStringDefault(defaultValue string, connectionTypes ...string) planmodifier.String {
	return &ConditionalStringDefaultPlanModifier{
		DefaultValue:      defaultValue,
		ConnectionTypes:   connectionTypes,
		ConnectionTypeKey: "connection_type",
	}
}
