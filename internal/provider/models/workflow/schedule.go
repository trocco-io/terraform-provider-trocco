package workflow

import (
	we "terraform-provider-trocco/internal/client/entities/workflow"
	wp "terraform-provider-trocco/internal/client/parameters/workflow"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

//
// Schedule
//

type Schedule struct {
	Type types.String `tfsdk:"type"`

	DailyConfig   *DailyScheduleConfig   `tfsdk:"daily_config"`
	HourlyConfig  *HourlyScheduleConfig  `tfsdk:"hourly_config"`
	MonthlyConfig *MonthlyScheduleConfig `tfsdk:"monthly_config"`
	WeeklyConfig  *WeeklyScheduleConfig  `tfsdk:"weekly_config"`
}

func NewSchedules(ens []we.Schedule) []Schedule {
	if ens == nil {
		return nil
	}

	var mds []Schedule
	for _, en := range ens {
		mds = append(mds, NewSchedule(en))
	}

	// If no schedules are present, the API returns an empty array but the provider should set `null`.
	if len(mds) == 0 {
		return nil
	}

	return mds
}

func NewSchedule(en we.Schedule) Schedule {
	return Schedule{
		Type:          types.StringValue(en.Type),
		DailyConfig:   NewDailyScheduleConfig(en.DailyConfig),
		HourlyConfig:  NewHourlyScheduleConfig(en.HourlyConfig),
		MonthlyConfig: NewMonthlyScheduleConfig(en.MonthlyConfig),
		WeeklyConfig:  NewWeeklyScheduleConfig(en.WeeklyConfig),
	}
}

func (m *Schedule) ToInput() wp.Schedule {
	p := wp.Schedule{
		Type: m.Type.ValueString(),
	}

	if m.DailyConfig != nil {
		p.DailyConfig = m.DailyConfig.ToInput()
	}
	if m.HourlyConfig != nil {
		p.HourlyConfig = m.HourlyConfig.ToInput()
	}
	if m.MonthlyConfig != nil {
		p.MonthlyConfig = m.MonthlyConfig.ToInput()
	}
	if m.WeeklyConfig != nil {
		p.WeeklyConfig = m.WeeklyConfig.ToInput()
	}

	return p
}

//
// DailyScheduleConfig
//

type DailyScheduleConfig struct {
	TimeZone types.String `tfsdk:"time_zone"`
	Hour     types.Int64  `tfsdk:"hour"`
	Minute   types.Int64  `tfsdk:"minute"`
}

func NewDailyScheduleConfig(en *we.DailyScheduleConfig) *DailyScheduleConfig {
	if en == nil {
		return nil
	}

	return &DailyScheduleConfig{
		TimeZone: types.StringValue(en.TimeZone),
		Hour:     types.Int64Value(en.Hour),
		Minute:   types.Int64Value(en.Minute),
	}
}

func (c *DailyScheduleConfig) ToInput() *wp.DailyScheduleConfig {
	return &wp.DailyScheduleConfig{
		TimeZone: c.TimeZone.ValueString(),
		Hour:     c.Hour.ValueInt64(),
		Minute:   c.Minute.ValueInt64(),
	}
}

//
// HourlyScheduleConfig
//

type HourlyScheduleConfig struct {
	TimeZone types.String `tfsdk:"time_zone"`
	Minute   types.Int64  `tfsdk:"minute"`
}

func NewHourlyScheduleConfig(en *we.HourlyScheduleConfig) *HourlyScheduleConfig {
	if en == nil {
		return nil
	}

	return &HourlyScheduleConfig{
		TimeZone: types.StringValue(en.TimeZone),
		Minute:   types.Int64Value(en.Minute),
	}
}

func (c *HourlyScheduleConfig) ToInput() *wp.HourlyScheduleConfig {
	return &wp.HourlyScheduleConfig{
		TimeZone: c.TimeZone.ValueString(),
		Minute:   c.Minute.ValueInt64(),
	}
}

//
// MonthlyScheduleConfig
//

type MonthlyScheduleConfig struct {
	TimeZone types.String `tfsdk:"time_zone"`
	Day      types.Int64  `tfsdk:"day"`
	Hour     types.Int64  `tfsdk:"hour"`
	Minute   types.Int64  `tfsdk:"minute"`
}

func NewMonthlyScheduleConfig(en *we.MonthlyScheduleConfig) *MonthlyScheduleConfig {
	if en == nil {
		return nil
	}

	return &MonthlyScheduleConfig{
		TimeZone: types.StringValue(en.TimeZone),
		Day:      types.Int64Value(en.Day),
		Hour:     types.Int64Value(en.Hour),
		Minute:   types.Int64Value(en.Minute),
	}
}

func (c *MonthlyScheduleConfig) ToInput() *wp.MonthlyScheduleConfig {
	return &wp.MonthlyScheduleConfig{
		TimeZone: c.TimeZone.ValueString(),
		Day:      c.Day.ValueInt64(),
		Hour:     c.Hour.ValueInt64(),
		Minute:   c.Minute.ValueInt64(),
	}
}

//
// WeeklyScheduleConfig
//

type WeeklyScheduleConfig struct {
	TimeZone  types.String `tfsdk:"time_zone"`
	DayOfWeek types.Int64  `tfsdk:"day_of_week"`
	Hour      types.Int64  `tfsdk:"hour"`
	Minute    types.Int64  `tfsdk:"minute"`
}

func NewWeeklyScheduleConfig(en *we.WeeklyScheduleConfig) *WeeklyScheduleConfig {
	if en == nil {
		return nil
	}

	return &WeeklyScheduleConfig{
		TimeZone:  types.StringValue(en.TimeZone),
		DayOfWeek: types.Int64Value(en.DayOfWeek),
		Hour:      types.Int64Value(en.Hour),
		Minute:    types.Int64Value(en.Minute),
	}
}

func (c *WeeklyScheduleConfig) ToInput() *wp.WeeklyScheduleConfig {
	return &wp.WeeklyScheduleConfig{
		TimeZone:  c.TimeZone.ValueString(),
		DayOfWeek: c.DayOfWeek.ValueInt64(),
		Hour:      c.Hour.ValueInt64(),
		Minute:    c.Minute.ValueInt64(),
	}
}
