package filter

import "github.com/hashicorp/terraform-plugin-framework/types"

type FilterStringTransform struct {
	ColumnName types.String `tfsdk:"column_name"`
	Type       types.String `tfsdk:"type"`
}
