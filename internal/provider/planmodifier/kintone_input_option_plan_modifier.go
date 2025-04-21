package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &KintoneInputOptionPlanModifier{}

type KintoneInputOptionPlanModifier struct{}

func (d *KintoneInputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating kintone input option attributes"
}

func (d *KintoneInputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *KintoneInputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var typ types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("type"), &typ)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var format types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("format"), &format)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if typ.ValueString() == "timestamp" && format.IsNull() {
		addKintoneInputOptionAttributeError(req, resp, "format is required when type is 'timestamp'")
	}

	if typ.ValueString() != "timestamp" && !format.IsNull() {
		addKintoneInputOptionAttributeError(req, resp, "format is only allowed when type is 'timestamp'")
	}
}

func addKintoneInputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"KintoneInputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
