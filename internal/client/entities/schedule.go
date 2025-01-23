package entities

type Schedule struct {
	Frequency string `json:"frequency"`
	Minute    int32  `json:"minute"`
	Hour      *int32 `json:"hour"`
	Day       *int32 `json:"day"`
	DayOfWeek *int32 `json:"day_of_week"`
	TimeZone  string `json:"time_zone"`
}
