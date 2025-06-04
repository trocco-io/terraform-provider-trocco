package validator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ validator.Set = mergeKeysRequiredOnlyForMergeModeValidator{}

type mergeKeysRequiredOnlyForMergeModeValidator struct{}

func (v mergeKeysRequiredOnlyForMergeModeValidator) Description(ctx context.Context) string {
	return "Ensures `merge_keys` is set only when `mode` is `merge`."
}

func (v mergeKeysRequiredOnlyForMergeModeValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v mergeKeysRequiredOnlyForMergeModeValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	modePath := req.Path.ParentPath().AtName("mode")
	var mode types.String

	diags := req.Config.GetAttribute(ctx, modePath, &mode)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() || mode.IsNull() || mode.IsUnknown() {
		return
	}

	isMergeKeysSet := false

	if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
		var mergeKeys []types.String
		diags := req.ConfigValue.ElementsAs(ctx, &mergeKeys, false)
		resp.Diagnostics.Append(diags...)
		if !resp.Diagnostics.HasError() && len(mergeKeys) > 0 {
			isMergeKeysSet = true
		}
	}

	if mode.ValueString() == "merge" {
		if !isMergeKeysSet {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Missing Required Merge Keys",
				"The `merge_keys` field must be set and contain at least one element when `mode` is `merge`.",
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

func MergeKeysRequiredOnlyForMergeMode() validator.Set {
	return mergeKeysRequiredOnlyForMergeModeValidator{}
}
