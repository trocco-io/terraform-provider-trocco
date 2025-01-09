package validator

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = WrappingDollarValidator{}

type WrappingDollarValidator struct{}

func (v WrappingDollarValidator) Description(ctx context.Context) string {
	return "value must start and end with `$`"
}

func (v WrappingDollarValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v WrappingDollarValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	s := req.ConfigValue.ValueString()

	if !isValid(s) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Wrapping Dollar Validator Error",
			fmt.Sprintf("Attribute %s %s, got: %s", req.Path, v.Description(ctx), s),
		)

		return
	}
}

func isValid(s string) bool {
	return strings.HasPrefix(s, "$") && strings.HasSuffix(s, "$")
}
