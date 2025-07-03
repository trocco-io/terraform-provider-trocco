package parameter

type ScheduleInput struct {
	Frequency string `json:"frequency"`
	Minute    int64  `json:"minute"`
	Hour      *int64 `json:"hour,omitempty"`
	Day       *int64 `json:"day,omitempty"`
	DayOfWeek *int64 `json:"day_of_week,omitempty"`
	TimeZone  string `json:"time_zone"`
}
