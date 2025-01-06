package pipeline_definition

import (
	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameters"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

//
// Schedule
//

type Schedule struct {
	Frequency types.String `tfsdk:"frequency"`
	TimeZone  types.String `tfsdk:"time_zone"`
	Day       types.Int64  `tfsdk:"day"`
	DayOfWeek types.Int64  `tfsdk:"day_of_week"`
	Hour      types.Int64  `tfsdk:"hour"`
	Minute    types.Int64  `tfsdk:"minute"`
}

func NewSchedules(ens []*we.Schedule, previous *PipelineDefinition) []*Schedule {
	if ens == nil {
		return nil
	}

	// If the attribute in the plan (or state) is nil, the provider should sets nil to the state.
	if previous.Schedules == nil && len(ens) == 0 {
		return nil
	}

	mds := []*Schedule{}
	for _, en := range ens {
		mds = append(mds, NewSchedule(en))
	}

	return mds
}

func NewSchedule(en *we.Schedule) *Schedule {
	return &Schedule{
		Frequency: types.StringValue(en.Frequency),
		TimeZone:  types.StringValue(en.TimeZone),
		Minute:    types.Int64Value(en.Minute),
		Day:       types.Int64PointerValue(en.Day),
		DayOfWeek: types.Int64PointerValue(en.DayOfWeek),
		Hour:      types.Int64PointerValue(en.Hour),
	}
}

func (m *Schedule) ToInput() *wp.Schedule {
	return &wp.Schedule{
		Frequency: m.Frequency.ValueString(),
		TimeZone:  m.TimeZone.ValueString(),
		Minute:    m.Minute.ValueInt64(),
		Day:       &p.NullableInt64{Valid: !m.Day.IsNull(), Value: m.Day.ValueInt64()},
		DayOfWeek: &p.NullableInt64{Valid: !m.DayOfWeek.IsNull(), Value: m.DayOfWeek.ValueInt64()},
		Hour:      &p.NullableInt64{Valid: !m.Hour.IsNull(), Value: m.Hour.ValueInt64()},
	}
}
