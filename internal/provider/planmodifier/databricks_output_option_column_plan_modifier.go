package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &DatabricksOutputOptionColumnPlanModifier{}

type DatabricksOutputOptionColumnPlanModifier struct{}

func (d *DatabricksOutputOptionColumnPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating databricks output option column attributes"
}

func (d *DatabricksOutputOptionColumnPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *DatabricksOutputOptionColumnPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var typ types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("type"), &typ)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var valueType types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("value_type"), &valueType)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var timestampFormat types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("timestamp_format"), &timestampFormat)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var timezone types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("timezone"), &timezone)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if typ.IsUnknown() || valueType.IsUnknown() || timestampFormat.IsUnknown() || timezone.IsUnknown() {
		return
	}

	// TROCCO decides whether it stores `timestamp_format` and `timezone` from
	// `value_type` alone, because that is what selects the embulk column setter
	// that reads them; `type` is only the DDL type, and is consulted for
	// `timezone` when `value_type` is omitted. A value TROCCO does not store is
	// dropped silently, which surfaces as "Provider produced inconsistent result
	// after apply", so reject the combination while planning instead.
	timestampFormatAvailable := valueType.ValueString() == "string" || valueType.ValueString() == "nstring"
	timezoneAvailable := timestampFormatAvailable || valueType.ValueString() == "date" || valueType.ValueString() == "time"
	if valueType.ValueString() == "" {
		timezoneAvailable = typ.ValueString() == "DATE" || typ.ValueString() == "TIMESTAMP"
	}

	if !timestampFormat.IsNull() && !timestampFormatAvailable {
		addDatabricksOutputOptionColumnAttributeError(req, resp, "timestamp_format can only be set when value_type is string or nstring")
	}

	if !timezone.IsNull() && !timezoneAvailable {
		addDatabricksOutputOptionColumnAttributeError(req, resp, "timezone can only be set when value_type is string, nstring, date or time, or, when value_type is omitted, when type is 'DATE' or 'TIMESTAMP'")
	}
}

func addDatabricksOutputOptionColumnAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Databricks output option column Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
