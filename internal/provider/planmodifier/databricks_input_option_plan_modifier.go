package planmodifier

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

var _ planmodifier.Object = &DatabricksInputOptionPlanModifier{}

type DatabricksInputOptionPlanModifier struct{}

func (d *DatabricksInputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating databricks input option attributes"
}

func (d *DatabricksInputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *DatabricksInputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	// No specific validation needed for Databricks input options currently
	// This modifier serves as a placeholder for future validation logic if needed
}
