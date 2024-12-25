package entities

type Schedule struct {
	Frequency string `json:"frequency"`
	Minute    int    `json:"minute"`
	Hour      *int   `json:"hour"`
	Day       *int   `json:"day"`
	DayOfWeek *int   `json:"day_of_week"`
	TimeZone  string `json:"time_zone"`
}
