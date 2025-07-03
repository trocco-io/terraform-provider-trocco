package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &SnowflakeOutputOptionColumnPlanModifier{}

type SnowflakeOutputOptionColumnPlanModifier struct{}

func (d *SnowflakeOutputOptionColumnPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating snowflake output option column attributes"
}

func (d *SnowflakeOutputOptionColumnPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *SnowflakeOutputOptionColumnPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
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

	if typ.ValueString() != "TIMESTAMP" && !timezone.IsNull() {
		addSnowflakeOutputOptionColumnAttributeError(req, resp, "timezone can only be set when type is 'TIMESTAMP'")
	}

	if !timestampFormat.IsNull() && (typ.ValueString() != "TIMESTAMP" || (valueType.ValueString() != "string" && valueType.ValueString() != "nstring")) {
		addSnowflakeOutputOptionColumnAttributeError(req, resp, "timestamp_format can only be set when type is 'TIMESTAMP' and value_type is string or nstring")
	}
}

func addSnowflakeOutputOptionColumnAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Snowflake output option column Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
