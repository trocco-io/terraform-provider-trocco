package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &YahooAdsApiYssInputOptionPlanModifier{}

type YahooAdsApiYssInputOptionPlanModifier struct{}

func (d *YahooAdsApiYssInputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating yahoo_ads_api_yss input option attributes"
}

func (d *YahooAdsApiYssInputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *YahooAdsApiYssInputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {

	var service, startDate, endDate, reportType types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("service"), &service)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("start_date"), &startDate)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("end_date"), &endDate)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("report_type"), &reportType)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if service.ValueString() == "report_definition_service" {
		if startDate.IsNull() {
			addYahooAdsApiYssInputOptionAttributeError(req, resp, "start_date is required for report_definition_service")
		}

		if endDate.IsNull() {
			addYahooAdsApiYssInputOptionAttributeError(req, resp, "end_date is required for report_definition_service")
		}

		if reportType.IsNull() {
			addYahooAdsApiYssInputOptionAttributeError(req, resp, "report_type is required for report_definition_service")
		}
	}

	if service.ValueString() == "campaign_export_service" {
		if !startDate.IsNull() && startDate.ValueString() != "" {
			addYahooAdsApiYssInputOptionAttributeError(req, resp, "start_date must not be set when service is campaign_export_service")
		}

		if !endDate.IsNull() && endDate.ValueString() != "" {
			addYahooAdsApiYssInputOptionAttributeError(req, resp, "end_date must not be set when service is campaign_export_service")
		}

		if !reportType.IsNull() && reportType.ValueString() != "" {
			addYahooAdsApiYssInputOptionAttributeError(req, resp, "report_type must not be set when service is campaign_export_service")
		}
	}
}
