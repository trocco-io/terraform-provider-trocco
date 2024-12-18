package filter

import "github.com/hashicorp/terraform-plugin-framework/types"

type FilterRows struct {
	Condition           types.String         `tfsdk:"condition"`
	FilterRowConditions []filterRowCondition `tfsdk:"filter_row_conditions"`
}

type filterRowCondition struct {
	Column   types.String `tfsdk:"column"`
	Operator types.String `tfsdk:"operator"`
	Argument types.String `tfsdk:"argument"`
}
