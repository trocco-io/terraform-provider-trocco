package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &S3OutputOptionPlanModifier{}

type S3OutputOptionPlanModifier struct{}

func (d *S3OutputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating s3 output option attributes"
}

func (d *S3OutputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *S3OutputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var formatterType types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("formatter_type"), &formatterType)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var csvFormatter types.Object
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("csv_formatter"), &csvFormatter)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var jsonlFormatter types.Object
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("jsonl_formatter"), &jsonlFormatter)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate formatter_type and formatter consistency
	if formatterType.ValueString() == "csv" {
		if csvFormatter.IsNull() {
			addS3OutputOptionAttributeError(req, resp, "csv_formatter must be set when formatter_type is 'csv'")
		}
		if !jsonlFormatter.IsNull() {
			addS3OutputOptionAttributeError(req, resp, "jsonl_formatter cannot be set when formatter_type is 'csv'")
		}
	}

	if formatterType.ValueString() == "jsonl" {
		if jsonlFormatter.IsNull() {
			addS3OutputOptionAttributeError(req, resp, "jsonl_formatter must be set when formatter_type is 'jsonl'")
		}
		if !csvFormatter.IsNull() {
			addS3OutputOptionAttributeError(req, resp, "csv_formatter cannot be set when formatter_type is 'jsonl'")
		}
	}
}

func addS3OutputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"S3 OutputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
