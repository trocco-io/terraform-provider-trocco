package models

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities"
	"terraform-provider-trocco/internal/client/parameters"
)

type Schedule struct {
	Frequency types.String `tfsdk:"frequency"`
	Minute    types.Int32  `tfsdk:"minute"`
	Hour      types.Int32  `tfsdk:"hour"`
	Day       types.Int32  `tfsdk:"day"`
	DayOfWeek types.Int32  `tfsdk:"day_of_week"`
	TimeZone  types.String `tfsdk:"time_zone"`
}

func NewSchedules(schedules *[]entities.Schedule) *[]Schedule {
	if schedules == nil {
		return nil
	}
	outputs := make([]Schedule, 0, len(*schedules))
	for _, input := range *schedules {
		schedule := Schedule{
			Frequency: types.StringValue(input.Frequency),
			Minute:    types.Int32Value(input.Minute),
			Hour:      types.Int32PointerValue(input.Hour),
			Day:       types.Int32PointerValue(input.Day),
			DayOfWeek: types.Int32PointerValue(input.DayOfWeek),
			TimeZone:  types.StringValue(input.TimeZone),
		}
		outputs = append(outputs, schedule)
	}
	return &outputs
}

func (schedule Schedule) ToInput() parameters.ScheduleInput {
	return parameters.ScheduleInput{
		Frequency: schedule.Frequency.ValueString(),
		Minute:    schedule.Minute.ValueInt32(),
		Hour:      schedule.Hour.ValueInt32Pointer(),
		Day:       schedule.Day.ValueInt32Pointer(),
		DayOfWeek: schedule.DayOfWeek.ValueInt32Pointer(),
		TimeZone:  schedule.TimeZone.ValueString(),
	}
}