package filter

import (
	filterEntities "terraform-provider-trocco/internal/client/entities/job_definitions/filter"
	filter2 "terraform-provider-trocco/internal/client/parameter/job_definitions/filter"
	"terraform-provider-trocco/internal/provider/models"

	"github.com/hashicorp/terraform-plugin-framework/types"
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

func (filterAddTime *FilterAddTime) ToInput() *filter2.FilterAddTimeInput {
	if filterAddTime == nil {
		return nil
	}

	return &filter2.FilterAddTimeInput{
		ColumnName:      filterAddTime.ColumnName.ValueString(),
		Type:            filterAddTime.Type.ValueString(),
		TimestampFormat: models.NewNullableString(filterAddTime.TimestampFormat),
		TimeZone:        models.NewNullableString(filterAddTime.TimeZone),
	}
}
