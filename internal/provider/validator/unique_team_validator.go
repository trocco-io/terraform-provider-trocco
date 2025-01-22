package validator

import (
	"context"
	"fmt"

	model "terraform-provider-trocco/internal/provider/model/resource_group"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type UniqueTeamValidator struct{}

func (v UniqueTeamValidator) Description(ctx context.Context) string {
	return "Ensures that team IDs are unique within the set."
}

func (v UniqueTeamValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v UniqueTeamValidator) ValidateSet(ctx context.Context, request validator.SetRequest, response *validator.SetResponse) {
	var teamIDs []model.TeamRoleResourceModel
	diags := request.ConfigValue.ElementsAs(ctx, &teamIDs, false)
	if diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	for i, teamID := range teamIDs {
		for j, otherTeamID := range teamIDs {
			if i == j {
				continue
			}

			if teamID.TeamID == otherTeamID.TeamID {
				response.Diagnostics.AddAttributeError(
					request.Path,
					"Duplicate Team ID",
					fmt.Sprintf("Team ID %q is duplicated in the list.", teamID.TeamID),
				)
				return
			}
		}
	}
}
