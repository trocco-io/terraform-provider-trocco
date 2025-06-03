package validator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.List = mergeKeysRequiredOnlyForMergeModeValidator{}

type mergeKeysRequiredOnlyForMergeModeValidator struct{}

func (v mergeKeysRequiredOnlyForMergeModeValidator) Description(ctx context.Context) string {
	return "Ensures `merge_keys` is set only when `mode` is `merge`."
}

func (v mergeKeysRequiredOnlyForMergeModeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v mergeKeysRequiredOnlyForMergeModeValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	modePath := req.Path.ParentPath().AtName("mode")
	var mode types.String

	diags := req.Config.GetAttribute(ctx, modePath, &mode)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() || mode.IsNull() || mode.IsUnknown() {
		return
	}

	isMergeKeysSet := !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown()

	if mode.ValueString() == "merge" {
		if !isMergeKeysSet {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing Required Merge Keys",
				"The `merge_keys` field must be set when `mode` is `merge`.",
			)
		}
	} else {
		if isMergeKeysSet {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Merge Keys Setting",
				"The `merge_keys` field must not be set when `mode` is not `merge`.",
			)
		}
	}
}

func MergeKeysRequiredOnlyForMergeMode() validator.List {
	return mergeKeysRequiredOnlyForMergeModeValidator{}
}
