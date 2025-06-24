package model

import (
	"terraform-provider-trocco/internal/client/entity"
	"terraform-provider-trocco/internal/client/parameter"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Schedule struct {
	Frequency types.String `tfsdk:"frequency"`
	Minute    types.Int64  `tfsdk:"minute"`
	Hour      types.Int64  `tfsdk:"hour"`
	Day       types.Int64  `tfsdk:"day"`
	DayOfWeek types.Int64  `tfsdk:"day_of_week"`
	TimeZone  types.String `tfsdk:"time_zone"`
}

func NewSchedules(schedules []entity.Schedule) []Schedule {
	if schedules == nil {
		return nil
	}
	outputs := make([]Schedule, 0, len(schedules))
	for _, input := range schedules {
		schedule := Schedule{
			Frequency: types.StringValue(input.Frequency),
			Minute:    types.Int64Value(int64(input.Minute)),
			Hour:      types.Int64PointerValue(int64PtrFromInt32Ptr(input.Hour)),
			Day:       types.Int64PointerValue(int64PtrFromInt32Ptr(input.Day)),
			DayOfWeek: types.Int64PointerValue(int64PtrFromInt32Ptr(input.DayOfWeek)),
			TimeZone:  types.StringValue(input.TimeZone),
		}
		outputs = append(outputs, schedule)
	}
	return outputs
}

func (schedule Schedule) ToInput() parameter.ScheduleInput {
	return parameter.ScheduleInput{
		Frequency: schedule.Frequency.ValueString(),
		Minute:    int32(schedule.Minute.ValueInt64()),
		Hour:      int32PtrFromInt64Ptr(schedule.Hour.ValueInt64Pointer()),
		Day:       int32PtrFromInt64Ptr(schedule.Day.ValueInt64Pointer()),
		DayOfWeek: int32PtrFromInt64Ptr(schedule.DayOfWeek.ValueInt64Pointer()),
		TimeZone:  schedule.TimeZone.ValueString(),
	}
}
