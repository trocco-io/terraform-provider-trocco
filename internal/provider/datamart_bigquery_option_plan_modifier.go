package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &datamartBigqueryOptionPlanModifier{}

type datamartBigqueryOptionPlanModifier struct{}

func (d *datamartBigqueryOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating datamart_bigquery_option"
}

func (d *datamartBigqueryOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *datamartBigqueryOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var queryMode types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("query_mode"), &queryMode)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var destinationDataset types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("destination_dataset"), &destinationDataset)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var destinationTable types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("destination_table"), &destinationTable)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var writeDisposition types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("write_disposition"), &writeDisposition)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var partitioning types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("partitioning"), &partitioning)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var partitioningTime types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("partitioning_time"), &partitioningTime)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var partitioningField types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("partitioning_field"), &partitioningField)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if queryMode.ValueString() == "insert" {
		if destinationDataset.IsNull() {
			addDatamartBigqueryOptionAttributeError(req, resp, "destination_dataset is required for insert query mode")
		}
		if destinationTable.IsNull() {
			addDatamartBigqueryOptionAttributeError(req, resp, "destination_table is required for insert query mode")
		}
		if writeDisposition.IsNull() {
			addDatamartBigqueryOptionAttributeError(req, resp, "write_disposition is required for insert query mode")
		}
		if !partitioning.IsNull() {
			if partitioningTime.IsNull() {
				addDatamartBigqueryOptionAttributeError(req, resp, "partitioning_time is required when partitioning is set")
			}
			if partitioning.ValueString() == "time_unit_column" && partitioningField.IsNull() {
				addDatamartBigqueryOptionAttributeError(req, resp, "partitioning_field is required when partitioning is set to time_unit_column")
			}
		}
	}
}

func addDatamartBigqueryOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"DatamartBigqueryOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
