package pipeline_definition

import (
	p "terraform-provider-trocco/internal/client/parameter"
)

type Schedule struct {
	Frequency string           `json:"frequency"`
	TimeZone  string           `json:"time_zone"`
	Minute    int64            `json:"minute"`
	Day       *p.NullableInt64 `json:"day,omitempty"`
	DayOfWeek *p.NullableInt64 `json:"day_of_week,omitempty"`
	Hour      *p.NullableInt64 `json:"hour,omitempty"`
}
