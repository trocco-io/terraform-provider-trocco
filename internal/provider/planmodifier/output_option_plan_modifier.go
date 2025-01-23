package planmodifier

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &OutputOptionPlanModifier{}

type OutputOptionPlanModifier struct{}

func (d *OutputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating schedule attributes"
}

func (d *OutputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *OutputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var outputOptionType types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("output_option_type"), &outputOptionType)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var outputOption types.Object
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("output_option"), &outputOption)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if output_option contains only attributes that start with output_option_type
	for attrName := range outputOption.Attributes() {
		if !outputOption.Attributes()[attrName].IsNull() && (outputOptionType.ValueString()+"_output_option") != attrName {
			addOutputOptionAttributeError(req, resp, fmt.Sprintf("Attribute output_option contains invalid attribute in output_option. attribute name: %s", attrName))
			return
		}
	}
}

func addOutputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"OutputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
