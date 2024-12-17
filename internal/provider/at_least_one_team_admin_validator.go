package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.List = AtLeastOneTeamAdminValidator{}

type AtLeastOneTeamAdminValidator struct{}

func (v AtLeastOneTeamAdminValidator) Description(ctx context.Context) string {
	return "Validates that at least one member has `role` set to `team_admin`."
}

func (v AtLeastOneTeamAdminValidator) MarkdownDescription(ctx context.Context) string {
	return "Validates that at least one member has `role` set to `team_admin`."

}

func (v AtLeastOneTeamAdminValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	var members []teamMemberResourceModel
	diags := req.ConfigValue.ElementsAs(ctx, &members, false)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	for _, member := range members {
		if member.Role.ValueString() == "team_admin" {
			return
		}
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Missing Team Admin",
		"The `members` list must include at least one member with `role` set to `team_admin`.",
	)
}
