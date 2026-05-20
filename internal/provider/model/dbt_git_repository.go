package model

import "github.com/hashicorp/terraform-plugin-framework/types"

type DbtGitRepositoryModel struct {
	ID              types.Int64  `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Description     types.String `tfsdk:"description"`
	AdapterType     types.String `tfsdk:"adapter_type"`
	DbtVersion      types.String `tfsdk:"dbt_version"`
	URL             types.String `tfsdk:"url"`
	Branch          types.String `tfsdk:"branch"`
	Subdirectory    types.String `tfsdk:"subdirectory"`
	ResourceGroupID types.Int64  `tfsdk:"resource_group_id"`
}
