package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &PostgresqlOutputOptionPlanModifier{}

type PostgresqlOutputOptionPlanModifier struct{}

func (d *PostgresqlOutputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating postgresql output option attributes"
}

func (d *PostgresqlOutputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *PostgresqlOutputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var mode types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("mode"), &mode)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var mergeKeys types.Set
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("merge_keys"), &mergeKeys)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if mode.ValueString() != "merge" && len(mergeKeys.Elements()) > 0 {
		addPostgresqlOutputOptionAttributeError(req, resp, "merge_keys can only be set when mode is 'merge'")
	}
}

func addPostgresqlOutputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"PostgreSQL Output Option Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
