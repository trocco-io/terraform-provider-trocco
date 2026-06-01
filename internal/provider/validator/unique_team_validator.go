package validator

import (
	"context"
	"fmt"

	model "terraform-provider-trocco/internal/provider/model/resource_group"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Set = UniqueTeamValidator{}

type UniqueTeamValidator struct{}

func (v UniqueTeamValidator) Description(ctx context.Context) string {
	return "Ensures that team IDs are unique within the set."
}

func (v UniqueTeamValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v UniqueTeamValidator) ValidateSet(ctx context.Context, request validator.SetRequest, response *validator.SetResponse) {
	// Skip validation when the value is unknown.
	// During the first plan, `team_id` may reference another resource's `.id`
	// attribute that has not yet been created; uniqueness can only be checked
	// once the values become known.
	if request.ConfigValue.IsUnknown() {
		return
	}

	var teamIDs []model.TeamRoleResourceModel
	diags := request.ConfigValue.ElementsAs(ctx, &teamIDs, false)
	if diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	for i, teamI := range teamIDs {
		// Skip elements whose `team_id` is unknown or null; they cannot
		// meaningfully participate in equality checks until apply time.
		if teamI.TeamID.IsUnknown() || teamI.TeamID.IsNull() {
			continue
		}
		for j, teamJ := range teamIDs {
			if i == j {
				continue
			}
			if teamJ.TeamID.IsUnknown() || teamJ.TeamID.IsNull() {
				continue
			}

			if teamI.TeamID == teamJ.TeamID {
				response.Diagnostics.AddAttributeError(
					request.Path,
					"Duplicate Team ID",
					fmt.Sprintf("Team ID %q is duplicated in the list.", teamI.TeamID),
				)
				return
			}
		}
	}
}
