package planmodifier

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NormalizeSQLQuery returns a plan modifier that normalizes SQL query strings
// by standardizing line endings and whitespace.
func NormalizeSQLQuery() planmodifier.String {
	return &normalizeSQLQueryModifier{}
}

// normalizeSQLQueryModifier implements the plan modifier.
type normalizeSQLQueryModifier struct{}

// Description returns a human-readable description of the plan modifier.
func (m *normalizeSQLQueryModifier) Description(ctx context.Context) string {
	return "Normalizes SQL query strings by standardizing line endings and whitespace"
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m *normalizeSQLQueryModifier) MarkdownDescription(ctx context.Context) string {
	return "Normalizes SQL query strings by standardizing line endings and whitespace"
}

// PlanModifyString implements the plan modification logic.
func (m *normalizeSQLQueryModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// If the plan doesn't have a value, do nothing
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	// If the state doesn't have a value, do nothing
	if req.StateValue.IsNull() {
		return
	}

	planQuery := req.PlanValue.ValueString()
	stateQuery := req.StateValue.ValueString()

	// Normalize both queries
	normalizedPlanQuery := normalizeQuery(planQuery)
	normalizedStateQuery := normalizeQuery(stateQuery)

	// If the normalized queries are the same, use the state value
	if normalizedPlanQuery == normalizedStateQuery {
		resp.PlanValue = req.StateValue
	} else {
		// Otherwise, use the normalized plan value
		resp.PlanValue = types.StringValue(normalizedPlanQuery)
	}
}

// normalizeQuery normalizes an SQL query string by standardizing line endings and whitespace.
func normalizeQuery(query string) string {
	// 1. Replace Windows line endings (\r\n) with Unix line endings (\n)
	normalized := strings.ReplaceAll(query, "\r\n", "\n")

	// 2. Trim leading and trailing whitespace of the entire query
	normalized = strings.TrimSpace(normalized)

	// 3. Remove spaces after semicolons
	normalized = strings.ReplaceAll(normalized, "; ", ";")

	// 4. Process each line separately
	lines := strings.Split(normalized, "\n")
	for i, line := range lines {
		// Trim trailing whitespace for each line
		lines[i] = strings.TrimRight(line, " \t")
	}
	normalized = strings.Join(lines, "\n")

	// 5. Remove trailing newlines at the end of the query
	normalized = strings.TrimRight(normalized, "\n")

	// 6. Handle semicolons with newlines
	normalized = strings.ReplaceAll(normalized, "\n;", ";")

	return normalized
}
