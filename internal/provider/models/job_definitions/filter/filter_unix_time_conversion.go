package filter

import "github.com/hashicorp/terraform-plugin-framework/types"

type FilterUnixTimeConversion struct {
	ColumnName       types.String `tfsdk:"column_name"`
	Kind             types.String `tfsdk:"kind"`
	UnixtimeUnit     types.String `tfsdk:"unixtime_unit"`
	DatetimeFormat   types.String `tfsdk:"datetime_format"`
	DatetimeTimezone types.String `tfsdk:"datetime_timezone"`
}
