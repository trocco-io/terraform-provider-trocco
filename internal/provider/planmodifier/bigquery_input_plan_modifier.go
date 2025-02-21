package planmodifier

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &BigqueryInputOptionPlanModifier{}

type BigqueryInputOptionPlanModifier struct{}

func (d *BigqueryInputOptionPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating bigquery input option attributes"
}

func (d *BigqueryInputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *BigqueryInputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var gcsUri types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("gcs_uri"), &gcsUri)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var gcsUriFormat types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("gcs_uri_format"), &gcsUriFormat)...)

	if gcsUriFormat.IsNull() || gcsUriFormat.ValueString() == "bucket" {
		if !gcsUri.IsNull() {
			regex := regexp.MustCompile(`^[a-z0-9._-]+$`)
			if !regex.MatchString(gcsUri.ValueString()) {
				addBigqueryInputOptionAttributeError(req, resp, "gcs_uri must match the regular expression ^[a-z0-9._-]+$")
			}
		}
	}
}

func addBigqueryInputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Bigquery InputOption Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}
