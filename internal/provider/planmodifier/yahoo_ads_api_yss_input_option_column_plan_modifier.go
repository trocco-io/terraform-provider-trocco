package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &YahooAdsApiYssInputOptionColumnPlanModifier{}

type YahooAdsApiYssInputOptionColumnPlanModifier struct{}

func (d *YahooAdsApiYssInputOptionColumnPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating yahoo_ads_api_yss input option column attributes"
}

func (d *YahooAdsApiYssInputOptionColumnPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *YahooAdsApiYssInputOptionColumnPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
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
		addYahooAdsApiYssInputOptionAttributeError(req, resp, "format is required when type is 'timestamp'")
	}

	if typ.ValueString() != "timestamp" && !format.IsNull() {
		addYahooAdsApiYssInputOptionAttributeError(req, resp, "format is only allowed when type is 'timestamp'")
	}
}

func addYahooAdsApiYssInputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"YahooAdsApiYssInputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
