package planmodifier

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &InputOptionPlanModifier{}

type InputOptionPlanModifier struct{}

func (d *InputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating input option attributes"
}

func (d *InputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *InputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var inputOptionType types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("input_option_type"), &inputOptionType)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var inputOption types.Object
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("input_option"), &inputOption)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Check if input_option contains only attributes that start with input_option_type
	for attrName := range inputOption.Attributes() {
		if !inputOption.Attributes()[attrName].IsNull() && (inputOptionType.ValueString()+"_input_option") != attrName {
			addInputOptionAttributeError(req, resp, fmt.Sprintf("Attribute input_option contains invalid attribute in input_option. attribute name: %s", attrName))
			return
		}
	}
}

func addInputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"InputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
