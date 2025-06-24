package pipeline_definition

import (
	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameter"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

//
// Schedule
//

type Schedule struct {
	Frequency types.String `tfsdk:"frequency"`
	TimeZone  types.String `tfsdk:"time_zone"`
	Day       types.Int32  `tfsdk:"day"`
	DayOfWeek types.Int32  `tfsdk:"day_of_week"`
	Hour      types.Int32  `tfsdk:"hour"`
	Minute    types.Int32  `tfsdk:"minute"`
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
		Minute:    types.Int32Value(int32(en.Minute)),
		Day:       types.Int32PointerValue(int32PtrFromInt64Ptr(en.Day)),
		DayOfWeek: types.Int32PointerValue(int32PtrFromInt64Ptr(en.DayOfWeek)),
		Hour:      types.Int32PointerValue(int32PtrFromInt64Ptr(en.Hour)),
	}
}

func int32PtrFromInt64Ptr(i64ptr *int64) *int32 {
	if i64ptr == nil {
		return nil
	}
	i32 := int32(*i64ptr)
	return &i32
}

func (m *Schedule) ToInput() *wp.Schedule {
	return &wp.Schedule{
		Frequency: m.Frequency.ValueString(),
		TimeZone:  m.TimeZone.ValueString(),
		Minute:    int64(m.Minute.ValueInt32()),
		Day:       &p.NullableInt64{Valid: !m.Day.IsNull(), Value: int64(m.Day.ValueInt32())},
		DayOfWeek: &p.NullableInt64{Valid: !m.DayOfWeek.IsNull(), Value: int64(m.DayOfWeek.ValueInt32())},
		Hour:      &p.NullableInt64{Valid: !m.Hour.IsNull(), Value: int64(m.Hour.ValueInt32())},
	}
}
