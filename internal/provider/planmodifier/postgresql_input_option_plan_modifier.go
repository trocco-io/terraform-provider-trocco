package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &PostgresqlInputOptionPlanModifier{}

type PostgresqlInputOptionPlanModifier struct{}

func (d *PostgresqlInputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating postgresql input option attributes"
}

func (d *PostgresqlInputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *PostgresqlInputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
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

	var incrementalColumns types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("incremental_columns"), &incrementalColumns)...)
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
			addPostgresqlInputOptionAttributeError(req, resp, "query is required when incremental_loading_enabled is false")
		}

		if !lastRecord.IsNull() || !incrementalColumns.IsNull() || !table.IsNull() {
			addPostgresqlInputOptionAttributeError(req, resp, "last_record, incremental_columns, and table must be null when incremental_loading_enabled is false")
		}
	} else {
		if !query.IsNull() {
			addPostgresqlInputOptionAttributeError(req, resp, "query must be null when incremental_loading_enabled is true")
		}

		if table.IsNull() {
			addPostgresqlInputOptionAttributeError(req, resp, "table is required when incremental_loading_enabled is true")
		}
	}
}

func addPostgresqlInputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"PostgresqlInputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
