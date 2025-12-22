package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Bool = &ConditionalBooleanDefaultPlanModifier{}

type ConditionalBooleanDefaultPlanModifier struct {
	DefaultValue      bool
	ConnectionTypes   []string // Connection types where this default should apply
	ConnectionTypeKey string   // Key to check for connection type (default: "connection_type")
}

func (d *ConditionalBooleanDefaultPlanModifier) Description(ctx context.Context) string {
	return "Applies default value for boolean fields only when connection_type matches specified types"
}

func (d *ConditionalBooleanDefaultPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *ConditionalBooleanDefaultPlanModifier) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
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
	// Build path to connection_type field (go up to parent, then access connection_type)
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

	if shouldApplyDefault {
		resp.PlanValue = types.BoolValue(d.DefaultValue)
	} else {
		// For non-matching connections, set to null
		resp.PlanValue = types.BoolNull()
	}
}

// Creates a plan modifier that applies defaults for specific connection types.
func ConditionalBooleanDefault(defaultValue bool, connectionTypes ...string) planmodifier.Bool {
	return &ConditionalBooleanDefaultPlanModifier{
		DefaultValue:      defaultValue,
		ConnectionTypes:   connectionTypes,
		ConnectionTypeKey: "connection_type",
	}
}
