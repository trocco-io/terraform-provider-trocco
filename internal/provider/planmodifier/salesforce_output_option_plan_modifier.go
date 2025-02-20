package planmodifier

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &SalesforceOutputOptionPlanModifier{}

type SalesforceOutputOptionPlanModifier struct{}

func (d *SalesforceOutputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating salesforce output option attributes"
}

func (d *SalesforceOutputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *SalesforceOutputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var upsertKey types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("upsert_key"), &upsertKey)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var actionType types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("action_type"), &actionType)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if actionType.ValueString() != "upsert" && !upsertKey.IsNull() {
		addSalesforceOutputOptionAttributeError(req, resp, "upsert_key can only be set when action_type is 'upsert'")
	}
}

func addSalesforceOutputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Salesforce Output Option Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
