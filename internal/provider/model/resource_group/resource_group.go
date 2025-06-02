package resource_group

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ResourceGroupResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Teams       types.Set    `tfsdk:"teams"`
}

type TeamRoleResourceModel struct {
	TeamID types.Int64  `tfsdk:"team_id"`
	Role   types.String `tfsdk:"role"`
}

func (TeamRoleResourceModel) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"team_id": types.Int64Type,
		"role":    types.StringType,
	}
}
