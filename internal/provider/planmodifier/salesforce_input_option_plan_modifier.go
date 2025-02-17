package planmodifier

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
)

var _ planmodifier.Object = &SalesforceInputOptionPlanModifier{}

type SalesforceInputOptionPlanModifier struct{}

func (d *SalesforceInputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating salesforce input option attributes"
}

func (d *SalesforceInputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *SalesforceInputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var soql types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("soql"), &soql)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var objectAcquisitionMethod types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("object_acquisition_method"), &objectAcquisitionMethod)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var isConvertTypeCustomColumns types.Bool
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("is_convert_type_custom_columns"), &isConvertTypeCustomColumns)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var columns types.List
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("columns"), &columns)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Verify that soql is not specified for all_columns
	if objectAcquisitionMethod.ValueString() == "all_columns" && !soql.IsNull() {
		addSalesforceInputOptionAttributeError(req, resp, "soql cannot be specified when object_acquisition_method is 'all_columns'")
	}

	// Verify that soql is specified for soql
	if objectAcquisitionMethod.ValueString() == "soql" && soql.IsNull() {
		addSalesforceInputOptionAttributeError(req, resp, "soql must be specified when object_acquisition_method is 'soql'")
	}

	// Verify that column type is valid
	if isConvertTypeCustomColumns.ValueBool() {
		for _, column := range columns.Elements() {
			if columnObj, ok := column.(types.Object); ok {
				attributes := columnObj.Attributes()
				colType, ok := attributes["type"].(types.String)
				if !ok {
					resp.Diagnostics.AddError(
						"Unexpected error",
						"column type is not a string",
					)
				}
				colName, ok := attributes["name"].(types.String)
				if !ok {
					resp.Diagnostics.AddError(
						"Unexpected error",
						"column name is not a string",
					)
				}
				if (colType.ValueString() != "boolean" && colType.ValueString() != "string") && strings.HasSuffix(colName.ValueString(), "__c") {
					addSalesforceInputOptionAttributeError(req, resp, "column type must be 'boolean' or 'string' when is_convert_type_custom_columns is true and column name end with '__c'")
				} else if colType.ValueString() != "string" && !strings.HasSuffix(colName.ValueString(), "__c") {
					addSalesforceInputOptionAttributeError(req, resp, "column type must be 'string' when is_convert_type_custom_columns is true and column name does not end with '__c'")
				}
			} else {
				resp.Diagnostics.AddError(
					"Invalid Column Type",
					"Expected column to be an object",
				)
				return
			}
		}
	}

}

func addSalesforceInputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"SalesforceInputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
