package validator

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ validator.String = requiresAttributeValueOneOf{}
	_ validator.String = requiredIfAttributeValueOneOf{}
)

// RequiresAttributeValueOneOf returns a validator that rejects a configured
// value unless the attribute at targetPath holds one of the given values.
//
// Use it for an attribute the API keeps only while a discriminator attribute
// (such as `auth_type`) has a specific value, and silently discards otherwise.
// Without a plan-time error the discarded value makes `apply` fail with
// "Provider produced inconsistent result after apply", which names no
// attribute. Declaring the condition next to the attribute it guards also
// confines the work to the new attributes when a discriminator value is added.
func RequiresAttributeValueOneOf(targetPath path.Path, values ...string) validator.String {
	return requiresAttributeValueOneOf{targetPath: targetPath, values: values}
}

// RequiredIfAttributeValueOneOf returns a validator that requires a value
// whenever the attribute at targetPath holds one of the given values,
// mirroring a server-side conditional presence validation.
func RequiredIfAttributeValueOneOf(targetPath path.Path, values ...string) validator.String {
	return requiredIfAttributeValueOneOf{targetPath: targetPath, values: values}
}

type requiresAttributeValueOneOf struct {
	targetPath path.Path
	values     []string
}

func (v requiresAttributeValueOneOf) Description(ctx context.Context) string {
	return fmt.Sprintf("value can only be set when %s is %s", v.targetPath, quoteValues(v.values))
}

func (v requiresAttributeValueOneOf) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v requiresAttributeValueOneOf) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// An unknown value is still a configured one, so it is validated as well:
	// the target value decides whether the API keeps it, not this value.
	if req.ConfigValue.IsNull() {
		return
	}

	target, ok := targetValue(ctx, v.targetPath, req, resp)
	if !ok {
		return
	}
	// A null target cannot satisfy the condition either: the attribute is only
	// meaningful once the target selects the variant it belongs to.
	if target.IsNull() || !slices.Contains(v.values, target.ValueString()) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Attribute Combination",
			fmt.Sprintf("`%s` can only be set when `%s` is %s.", req.Path, v.targetPath, quoteValues(v.values)),
		)
	}
}

type requiredIfAttributeValueOneOf struct {
	targetPath path.Path
	values     []string
}

func (v requiredIfAttributeValueOneOf) Description(ctx context.Context) string {
	return fmt.Sprintf("value is required when %s is %s", v.targetPath, quoteValues(v.values))
}

func (v requiredIfAttributeValueOneOf) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v requiredIfAttributeValueOneOf) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}

	target, ok := targetValue(ctx, v.targetPath, req, resp)
	if !ok {
		return
	}
	if !target.IsNull() && slices.Contains(v.values, target.ValueString()) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Missing Attribute Configuration",
			fmt.Sprintf("`%s` is required when `%s` is %s.", req.Path, v.targetPath, quoteValues(v.values)),
		)
	}
}

// targetValue reads the attribute the condition depends on. It reports false
// when the value cannot decide the condition yet, either because reading it
// failed or because it is only known after apply.
func targetValue(ctx context.Context, targetPath path.Path, req validator.StringRequest, resp *validator.StringResponse) (types.String, bool) {
	var target types.String
	diags := req.Config.GetAttribute(ctx, targetPath, &target)
	resp.Diagnostics.Append(diags...)
	if diags.HasError() || target.IsUnknown() {
		return target, false
	}

	return target, true
}

func quoteValues(values []string) string {
	quoted := make([]string, 0, len(values))
	for _, v := range values {
		quoted = append(quoted, fmt.Sprintf("%q", v))
	}

	return strings.Join(quoted, " or ")
}
