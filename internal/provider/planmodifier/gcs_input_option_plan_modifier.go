package planmodifier

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &GcsInputOptionPlanModifier{}

type GcsInputOptionPlanModifier struct{}

func (d *GcsInputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating schedule attributes"
}

func (d *GcsInputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *GcsInputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var lastPath types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("last_path"), &lastPath)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var incrementalLoadingEnabled types.Bool
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("incremental_loading_enabled"), &incrementalLoadingEnabled)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !incrementalLoadingEnabled.ValueBool() && !lastPath.IsNull() {
		addGcsInputOptionAttributeError(req, resp, "last_path is only valid when incremental_loading_enabled is true")
	}
}

func addGcsInputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Gcs InputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
