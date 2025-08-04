package pipeline_definition

import (
	parameter "terraform-provider-trocco/internal/client/parameter"
)

type Schedule struct {
	Frequency string                   `json:"frequency"`
	TimeZone  string                   `json:"time_zone"`
	Minute    int64                    `json:"minute"`
	Day       *parameter.NullableInt64 `json:"day,omitempty"`
	DayOfWeek *parameter.NullableInt64 `json:"day_of_week,omitempty"`
	Hour      *parameter.NullableInt64 `json:"hour,omitempty"`
}
