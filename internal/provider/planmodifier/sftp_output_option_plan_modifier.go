package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ planmodifier.Object = &SftpOutputOptionPlanModifier{}

type SftpOutputOptionPlanModifier struct{}

func (d *SftpOutputOptionPlanModifier) Description(ctx context.Context) string {
	return "modifier for validating sftp output option attributes"
}

func (d *SftpOutputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *SftpOutputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	// If the entire sftp_output_option is null or unknown, skip validation
	if req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
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

	// Validate that exactly one formatter is provided
	csvProvided := !csvFormatter.IsNull() && !csvFormatter.IsUnknown()
	jsonlProvided := !jsonlFormatter.IsNull() && !jsonlFormatter.IsUnknown()

	if !csvProvided && !jsonlProvided {
		addSftpOutputOptionAttributeError(req, resp, "either csv_formatter or jsonl_formatter must be provided")
		return
	}

	if csvProvided && jsonlProvided {
		addSftpOutputOptionAttributeError(req, resp, "only one of csv_formatter or jsonl_formatter can be provided")
		return
	}
}

func addSftpOutputOptionAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"SftpOutputOption Validation Error",
		fmt.Sprintf("attribute %s %s", req.Path, message),
	)
}
