package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &mysqlInputOptionPlanModifier{}

type mysqlInputOptionPlanModifier struct{}

func (d *mysqlInputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating schedule attributes"
}

func (d *mysqlInputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *mysqlInputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var lastRecord types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("last_record"), &lastRecord)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var incrementalLoadingEnabled types.Bool
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("incremental_loading_enabled"), &incrementalLoadingEnabled)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !incrementalLoadingEnabled.ValueBool() && !lastRecord.IsNull() {
		addMysqlInputOptionAttributeError(req, resp, "last_record is only valid when incremental_loading_enabled is true")
	}

}

func addMysqlInputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"MysqlInputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
