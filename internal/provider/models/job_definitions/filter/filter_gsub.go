package filter

import "github.com/hashicorp/terraform-plugin-framework/types"

type FilterGsub struct {
	ColumnName types.String `tfsdk:"column_name"`
	Pattern    types.String `tfsdk:"pattern"`
	To         types.String `tfsdk:"to"`
}
