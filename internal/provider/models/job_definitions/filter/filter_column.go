package filter

import "github.com/hashicorp/terraform-plugin-framework/types"

type FilterColumn struct {
	Name                     types.String       `tfsdk:"name"`
	Src                      types.String       `tfsdk:"src"`
	Type                     types.String       `tfsdk:"type"`
	Default                  types.String       `tfsdk:"default"`
	HasParser                types.Bool         `tfsdk:"has_parser"`
	Format                   types.String       `tfsdk:"format"`
	JSONExpandEnabled        types.Bool         `tfsdk:"json_expand_enabled"`
	JSONExpandKeepBaseColumn types.Bool         `tfsdk:"json_expand_keep_base_column"`
	JSONExpandColumns        []jsonExpandColumn `tfsdk:"json_expand_columns"`
}

type jsonExpandColumn struct {
	Name     types.String `tfsdk:"name"`
	JSONPath types.String `tfsdk:"json_path"`
	Type     types.String `tfsdk:"type"`
	Format   types.String `tfsdk:"format"`
	Timezone types.String `tfsdk:"timezone"`
}
