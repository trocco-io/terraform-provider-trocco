package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &GoogleSpreadsheetsInputOptionColumnPlanModifier{}

type GoogleSpreadsheetsInputOptionColumnPlanModifier struct{}

func (d *GoogleSpreadsheetsInputOptionColumnPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating snowflake output option attributes"
}

func (d *GoogleSpreadsheetsInputOptionColumnPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *GoogleSpreadsheetsInputOptionColumnPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
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
		attributeError(req, resp, "format is required when type is 'timestamp'")
	}

	if typ.ValueString() != "timestamp" && !format.IsNull() {
		attributeError(req, resp, "format is only allowed when type is 'timestamp'")
	}
}

func attributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Google Spreadsheets Input Option Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
