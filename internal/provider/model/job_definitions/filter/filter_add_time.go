package filter

import (
	filterEntities "terraform-provider-trocco/internal/client/entities/job_definitions/filter"
	filter2 "terraform-provider-trocco/internal/client/parameter/job_definitions/filter"
	"terraform-provider-trocco/internal/provider/model"

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

func (t *FilterAddTime) ToInput() *filter2.FilterAddTimeInput {
	if t == nil {
		return nil
	}

	return &filter2.FilterAddTimeInput{
		ColumnName:      t.ColumnName.ValueString(),
		Type:            t.Type.ValueString(),
		TimestampFormat: model.NewNullableString(t.TimestampFormat),
		TimeZone:        model.NewNullableString(t.TimeZone),
	}
}
