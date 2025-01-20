package pipeline_definition

type Schedule struct {
	Frequency string `json:"frequency"`
	TimeZone  string `json:"time_zone"`
	Minute    int64  `json:"minute"`
	Day       *int64 `json:"day"`
	DayOfWeek *int64 `json:"day_of_week"`
	Hour      *int64 `json:"hour"`
}
