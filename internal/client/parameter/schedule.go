package parameter

type ScheduleInput struct {
	Frequency string `json:"frequency"`
	Minute    int32  `json:"minute"`
	Hour      *int32 `json:"hour,omitempty"`
	Day       *int32 `json:"day,omitempty"`
	DayOfWeek *int32 `json:"day_of_week,omitempty"`
	TimeZone  string `json:"time_zone"`
}
