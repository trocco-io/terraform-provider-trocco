package validator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	model "terraform-provider-trocco/internal/provider/model/team"
)

var _ validator.Set = AtLeastOneTeamAdminValidator{}

type AtLeastOneTeamAdminValidator struct{}

func (v AtLeastOneTeamAdminValidator) Description(ctx context.Context) string {
	return "Validates that at least one member has `role` set to `team_admin`."
}

func (v AtLeastOneTeamAdminValidator) MarkdownDescription(ctx context.Context) string {
	return "Validates that at least one member has `role` set to `team_admin`."

}

func (v AtLeastOneTeamAdminValidator) ValidateSet(ctx context.Context, req validator.SetRequest, resp *validator.SetResponse) {
	// Skip validation only if the value is unknown
	// This allows the validation to run during apply phase when the value becomes known
	if req.ConfigValue.IsUnknown() {
		return
	}

	// If the value is null, it's an error - we need at least one member
	if req.ConfigValue.IsNull() {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Missing Team Members",
			"The `members` list cannot be null. It must include at least one member with `role` set to `team_admin`.",
		)
		return
	}

	var members []model.TeamMemberResourceModel
	diags := req.ConfigValue.ElementsAs(ctx, &members, false)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Check if the list is empty
	if len(members) == 0 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Empty Team Members",
			"The `members` list cannot be empty. It must include at least one member with `role` set to `team_admin`.",
		)
		return
	}

	// Check for at least one team_admin
	hasTeamAdmin := false
	for _, member := range members {
		// Skip members with unknown role during plan phase
		if member.Role.IsUnknown() {
			continue
		}
		if member.Role.ValueString() == "team_admin" {
			hasTeamAdmin = true
			break
		}
	}

	// Only report error if we've checked all members and found no team_admin
	// and there are no unknown roles that might become team_admin
	allRolesKnown := true
	for _, member := range members {
		if member.Role.IsUnknown() {
			allRolesKnown = false
			break
		}
	}

	if !hasTeamAdmin && allRolesKnown {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Missing Team Admin",
			"The `members` list must include at least one member with `role` set to `team_admin`.",
		)
	}
}
