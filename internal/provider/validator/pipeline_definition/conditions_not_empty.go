package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.List = ConditionsNotEmpty{}

type ConditionsNotEmpty struct {
}

func (v ConditionsNotEmpty) Description(ctx context.Context) string {
	return "Ensures at least one condition is specified."
}

func (v ConditionsNotEmpty) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ConditionsNotEmpty) ValidateList(
	ctx context.Context,
	req validator.ListRequest,
	resp *validator.ListResponse,
) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	if len(req.ConfigValue.Elements()) == 0 {
		resp.Diagnostics.AddError(
			"Invalid Conditions Configuration",
			"At least one condition must be specified in 'conditions'. The list cannot be empty.",
		)
	}
}
