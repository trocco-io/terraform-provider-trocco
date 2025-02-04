package planmodifier

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &SnowflakeOutputOptionPlanModifier{}

type SnowflakeOutputOptionPlanModifier struct{}

func (d *SnowflakeOutputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating snowflake output option attributes"
}

func (d *SnowflakeOutputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *SnowflakeOutputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var mode types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("mode"), &mode)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var snowflakeOutputOptionMergeKeys types.List
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("snowflake_output_option_merge_keys"), &snowflakeOutputOptionMergeKeys)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if mode.ValueString() != "merge" && len(snowflakeOutputOptionMergeKeys.Elements()) > 0 {
		addSnowflakeOutputOptionAttributeError(req, resp, "snowflake_output_option_merge_keys can only be set when mode is 'merge'")
	}
}

func addSnowflakeOutputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Snowflake Output Option Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
