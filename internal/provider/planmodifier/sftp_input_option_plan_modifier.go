package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &SftpInputOptionPlanModifier{}

type SftpInputOptionPlanModifier struct{}

func (d *SftpInputOptionPlanModifier) Description(ctx context.Context) string {
	return "modifier for validating sftp input option attributes"
}

func (d *SftpInputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *SftpInputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var lastPath types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("last_path"), &lastPath)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var incrementalLoadingEnabled types.Bool
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("incremental_loading_enabled"), &incrementalLoadingEnabled)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !incrementalLoadingEnabled.ValueBool() && !lastPath.IsNull() && lastPath.ValueString() != "" {
		addSftpInputOptionAttributeError(req, resp, "last_path is only valid when incremental_loading_enabled is true")
	}
}

func addSftpInputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"SftpInputOption Validation Error",
		fmt.Sprintf("attribute %s %s", req.Path, message),
	)
}
