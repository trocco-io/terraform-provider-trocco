package validator

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type UserPasswordValidator struct{}

func (v UserPasswordValidator) Description(ctx context.Context) string {
	return "The password must be at least 8 characters long and contain at least one letter and one number. It is required when creating a new user but optional during updates."
}

func (v UserPasswordValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v UserPasswordValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	password := request.ConfigValue.ValueString()
	// see: https://documents.trocco.io/docs/password-policy

	if password == "" {
		return
	}

	if len(password) < 8 || len(password) > 128 {
		response.Diagnostics.AddError(
			"Invalid Password",
			"Password must be between 8 and 128 characters.",
		)
	}

	if !regexp.MustCompile(`[a-zA-Z]`).MatchString(password) {
		response.Diagnostics.AddError(
			"Invalid Password",
			"Password must contain at least one letter.",
		)
	}

	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		response.Diagnostics.AddError(
			"Invalid Password",
			"Password must contain at least one number.",
		)
	}
}
