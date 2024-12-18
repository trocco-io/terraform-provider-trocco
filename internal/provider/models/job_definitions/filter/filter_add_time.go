package filter

import "github.com/hashicorp/terraform-plugin-framework/types"

type FilterAddTime struct {
	ColumnName      types.String `tfsdk:"column_name"`
	Type            types.String `tfsdk:"type"`
	TimestampFormat types.String `tfsdk:"timestamp_format"`
	TimeZone        types.String `tfsdk:"time_zone"`
}
