package team

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamResourceModel struct {
	ID          types.Int64               `tfsdk:"id"`
	Name        types.String              `tfsdk:"name"`
	Description types.String              `tfsdk:"description"`
	Members     []TeamMemberResourceModel `tfsdk:"members"`
}

type TeamMemberResourceModel struct {
	UserID types.Int64  `tfsdk:"user_id"`
	Role   types.String `tfsdk:"role"`
}
