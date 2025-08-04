package filter

import (
	filterEntities "terraform-provider-trocco/internal/client/entity/job_definition/filter"
	filterParameters "terraform-provider-trocco/internal/client/parameter/job_definition/filter"
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

func (t *FilterAddTime) ToInput() *filterParameters.FilterAddTimeInput {
	if t == nil {
		return nil
	}

	return &filterParameters.FilterAddTimeInput{
		ColumnName:      t.ColumnName.ValueString(),
		Type:            t.Type.ValueString(),
		TimestampFormat: model.NewNullableString(t.TimestampFormat),
		TimeZone:        model.NewNullableString(t.TimeZone),
	}
}
