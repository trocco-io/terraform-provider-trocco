package filter

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	filterEntities "terraform-provider-trocco/internal/client/entities/job_definitions/filter"
)

type FilterAddTime struct {
	ColumnName      types.String `tfsdk:"column_name"`
	Type            types.String `tfsdk:"type"`
	TimestampFormat types.String `tfsdk:"timestamp_format"`
	TimeZone        types.String `tfsdk:"time_zone"`
}

func NewFilterAddTime(filterAddTime *filterEntities.FilterAddTime) *FilterAddTime {
	if filterAddTime == nil {
		return nil
	}
	return &FilterAddTime{
		ColumnName:      types.StringValue(filterAddTime.ColumnName),
		Type:            types.StringValue(filterAddTime.Type),
		TimestampFormat: types.StringPointerValue(filterAddTime.TimestampFormat),
		TimeZone:        types.StringPointerValue(filterAddTime.TimeZone),
	}
}
