package filter

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/filter"
	filter2 "terraform-provider-trocco/internal/client/parameter/job_definition/filter"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterUnixTimeConversion struct {
	ColumnName       types.String `tfsdk:"column_name"`
	Kind             types.String `tfsdk:"kind"`
	UnixtimeUnit     types.String `tfsdk:"unixtime_unit"`
	DatetimeFormat   types.String `tfsdk:"datetime_format"`
	DatetimeTimezone types.String `tfsdk:"datetime_timezone"`
}

func NewFilterUnixTimeConversions(filterUnixTimeConversions []filter.FilterUnixTimeConversion) []FilterUnixTimeConversion {
	if len(filterUnixTimeConversions) == 0 {
		return nil
	}

	outputs := make([]FilterUnixTimeConversion, 0, len(filterUnixTimeConversions))
	for _, input := range filterUnixTimeConversions {
		filterUnixTimeConversion := FilterUnixTimeConversion{
			ColumnName:       types.StringValue(input.ColumnName),
			Kind:             types.StringValue(input.Kind),
			UnixtimeUnit:     types.StringValue(input.UnixtimeUnit),
			DatetimeFormat:   types.StringValue(input.DatetimeFormat),
			DatetimeTimezone: types.StringValue(input.DatetimeTimezone),
		}
		outputs = append(outputs, filterUnixTimeConversion)
	}
	return outputs
}

func (filterUnixTimeConversion FilterUnixTimeConversion) ToInput() filter2.FilterUnixTimeConversionInput {
	return filter2.FilterUnixTimeConversionInput{
		ColumnName:       filterUnixTimeConversion.ColumnName.ValueString(),
		Kind:             filterUnixTimeConversion.Kind.ValueString(),
		UnixtimeUnit:     filterUnixTimeConversion.UnixtimeUnit.ValueString(),
		DatetimeFormat:   filterUnixTimeConversion.DatetimeFormat.ValueString(),
		DatetimeTimezone: filterUnixTimeConversion.DatetimeTimezone.ValueString(),
	}
}
