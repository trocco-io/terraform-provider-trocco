package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Object = DestinationsNotBothEmpty{}

type DestinationsNotBothEmpty struct {
}

func (v DestinationsNotBothEmpty) Description(ctx context.Context) string {
	return "Ensures at least one of 'if' or 'else' has destinations specified."
}

func (v DestinationsNotBothEmpty) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v DestinationsNotBothEmpty) ValidateObject(
	ctx context.Context,
	req validator.ObjectRequest,
	resp *validator.ObjectResponse,
) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	attrs := req.ConfigValue.Attributes()

	ifAttr, okIf := attrs["if"].(types.List)
	elseAttr, okElse := attrs["else"].(types.List)

	if !okIf || !okElse {
		return
	}

	// Check if both are empty
	ifEmpty := ifAttr.IsNull() || ifAttr.IsUnknown() || len(ifAttr.Elements()) == 0
	elseEmpty := elseAttr.IsNull() || elseAttr.IsUnknown() || len(elseAttr.Elements()) == 0

	if ifEmpty && elseEmpty {
		resp.Diagnostics.AddError(
			"Invalid Destinations Configuration",
			"At least one destination must be specified in 'if' or 'else'. Both cannot be empty.",
		)
	}
}
