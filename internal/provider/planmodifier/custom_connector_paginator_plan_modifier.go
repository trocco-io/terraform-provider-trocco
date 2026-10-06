package planmodifier

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var liquidVariableNameRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

var _ planmodifier.Object = &CustomConnectorPaginatorPlanModifier{}

// CustomConnectorPaginatorPlanModifier validates that when paginator.inject_into
// is "request_body", the enclosing endpoint uses method = "POST" and has
// request_body set, since there is otherwise no request body to inject into.
type CustomConnectorPaginatorPlanModifier struct{}

func (d *CustomConnectorPaginatorPlanModifier) Description(ctx context.Context) string {
	return "Validates that method=POST and request_body are set when paginator.inject_into is request_body"
}

func (d *CustomConnectorPaginatorPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *CustomConnectorPaginatorPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	var injectInto types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("inject_into"), &injectInto)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if injectInto.IsUnknown() || injectInto.ValueString() != "request_body" {
		return
	}

	var method types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.ParentPath().AtName("method"), &method)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if !method.IsUnknown() && method.ValueString() != "POST" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"CustomConnector Paginator Validation Error",
			fmt.Sprintf("Attribute %s method must be POST when paginator.inject_into is \"request_body\"", req.Path.ParentPath().AtName("method")),
		)
	}

	var requestBody types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.ParentPath().AtName("request_body"), &requestBody)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if requestBody.IsNull() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"CustomConnector Paginator Validation Error",
			fmt.Sprintf("Attribute %s request_body is required when paginator.inject_into is \"request_body\"", req.Path.ParentPath().AtName("request_body")),
		)
	}

	var pageTokenOption types.Object
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("page_token_option"), &pageTokenOption)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if pageTokenOption.IsNull() {
		resp.Diagnostics.AddAttributeError(
			req.Path.AtName("page_token_option"),
			"CustomConnector Paginator Validation Error",
			fmt.Sprintf("Attribute %s is required when paginator.inject_into is \"request_body\"", req.Path.AtName("page_token_option")),
		)
	}

	var pageSizeFieldName, pageTokenFieldName types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("page_size_option").AtName("field_name"), &pageSizeFieldName)...)
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("page_token_option").AtName("field_name"), &pageTokenFieldName)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !pageSizeFieldName.IsNull() && !pageSizeFieldName.IsUnknown() && !liquidVariableNameRegex.MatchString(pageSizeFieldName.ValueString()) {
		resp.Diagnostics.AddAttributeError(
			req.Path.AtName("page_size_option").AtName("field_name"),
			"CustomConnector Paginator Validation Error",
			fmt.Sprintf("Attribute %s must be a valid Liquid variable name when paginator.inject_into is \"request_body\"", req.Path.AtName("page_size_option").AtName("field_name")),
		)
	}
	if !pageTokenFieldName.IsNull() && !pageTokenFieldName.IsUnknown() && !liquidVariableNameRegex.MatchString(pageTokenFieldName.ValueString()) {
		resp.Diagnostics.AddAttributeError(
			req.Path.AtName("page_token_option").AtName("field_name"),
			"CustomConnector Paginator Validation Error",
			fmt.Sprintf("Attribute %s must be a valid Liquid variable name when paginator.inject_into is \"request_body\"", req.Path.AtName("page_token_option").AtName("field_name")),
		)
	}
	// `Equal` reports true for two unknown values, so unknowns are skipped here
	// as well: the final values may still differ.
	if !pageSizeFieldName.IsNull() && !pageSizeFieldName.IsUnknown() &&
		!pageTokenFieldName.IsNull() && !pageTokenFieldName.IsUnknown() &&
		pageSizeFieldName.Equal(pageTokenFieldName) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"CustomConnector Paginator Validation Error",
			"page_size_option.field_name and page_token_option.field_name must not be the same when paginator.inject_into is \"request_body\"",
		)
	}
}
