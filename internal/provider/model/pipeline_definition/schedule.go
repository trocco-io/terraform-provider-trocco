package pipeline_definition

import (
	"context"
	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameter"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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

func NewSchedules(ctx context.Context, ens []*we.Schedule, previous *PipelineDefinition) types.Set {
	objectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"frequency":   types.StringType,
			"time_zone":   types.StringType,
			"minute":      types.Int64Type,
			"day":         types.Int64Type,
			"day_of_week": types.Int64Type,
			"hour":        types.Int64Type,
		},
	}

	if ens == nil {
		return types.SetNull(objectType)
	}

	// If the attribute in the plan (or state) is null, the provider should set null to the state.
	var isSchedulesNull bool
	if previous == nil {
		isSchedulesNull = true
	} else {
		isSchedulesNull = previous.Schedules.IsNull()
	}

	if isSchedulesNull && len(ens) == 0 {
		return types.SetNull(objectType)
	}

	mds := []*Schedule{}
	for _, en := range ens {
		mds = append(mds, NewSchedule(en))
	}

	setValue, diags := types.SetValueFrom(ctx, objectType, mds)
	if diags.HasError() {
		return types.SetNull(objectType)
	}

	return setValue
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
