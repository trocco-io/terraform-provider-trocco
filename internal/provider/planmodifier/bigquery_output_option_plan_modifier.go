package planmodifier

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &BigqueryOutputOptionPlanModifier{}

type BigqueryOutputOptionPlanModifier struct{}

func (d *BigqueryOutputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating schedule attributes"
}

func (d *BigqueryOutputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *BigqueryOutputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var mode types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("mode"), &mode)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var bigqueryOutputOptionMergeKeys types.List
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("bigquery_output_option_merge_keys"), &bigqueryOutputOptionMergeKeys)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if mode.ValueString() != "merge" && len(bigqueryOutputOptionMergeKeys.Elements()) > 0 {
		addBigqueryOutputOptionAttributeError(req, resp, "bigquery_output_option_merge_keys can only be set when mode is 'merge'")
	}

	if mode.ValueString() == "merge" && len(bigqueryOutputOptionMergeKeys.Elements()) == 0 {
		addBigqueryOutputOptionAttributeError(req, resp, "bigquery_output_option_merge_keys must be set when mode is 'merge'")
	}
}

func addBigqueryOutputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Gcs InputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
