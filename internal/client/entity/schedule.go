package entity

type Schedule struct {
	Frequency string `json:"frequency"`
	Minute    int64  `json:"minute"`
	Hour      *int64 `json:"hour"`
	Day       *int64 `json:"day"`
	DayOfWeek *int64 `json:"day_of_week"`
	TimeZone  string `json:"time_zone"`
}
