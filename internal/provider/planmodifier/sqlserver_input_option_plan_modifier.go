package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &SqlserverInputOptionPlanModifier{}

type SqlserverInputOptionPlanModifier struct{}

func (d *SqlserverInputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating sqlserver input option attributes"
}

func (d *SqlserverInputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *SqlserverInputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var inputOptionType types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("input_option_type"), &inputOptionType)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if inputOptionType.ValueString() != "sqlserver" {
		return
	}

	var incrementalLoadingEnabled types.Bool
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("incremental_loading_enabled"), &incrementalLoadingEnabled)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var query types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("query"), &query)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var lastRecord types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("last_record"), &lastRecord)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var table types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("table"), &table)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !incrementalLoadingEnabled.ValueBool() {
		if query.IsNull() {
			addSqlserverInputOptionAttributeError(req, resp, "query is required when incremental_loading_enabled is false")
		}

		if !lastRecord.IsNull() {
			addSqlserverInputOptionAttributeError(req, resp, "last_record is only valid when incremental_loading_enabled is true")
		}
	} else if table.IsNull() {
		addSqlserverInputOptionAttributeError(req, resp, "table is required when incremental_loading_enabled is true")
	}
}

func addSqlserverInputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"SqlserverInputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
