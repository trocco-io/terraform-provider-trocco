package parameters

type ScheduleInput struct {
	Frequency string `json:"frequency"`
	Minute    int    `json:"minute"`
	Hour      *int   `json:"hour,omitempty"`
	Day       *int   `json:"day,omitempty"`
	DayOfWeek *int   `json:"day_of_week,omitempty"`
	TimeZone  string `json:"time_zone"`
}
