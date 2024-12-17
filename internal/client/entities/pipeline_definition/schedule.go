package pipeline_definition

type Schedule struct {
	Type string `json:"type"`

	DailyConfig   *DailyScheduleConfig   `json:"daily_config"`
	HourlyConfig  *HourlyScheduleConfig  `json:"hourly_config"`
	MonthlyConfig *MonthlyScheduleConfig `json:"monthly_config"`
	WeeklyConfig  *WeeklyScheduleConfig  `json:"weekly_config"`
}

type DailyScheduleConfig struct {
	TimeZone string `json:"time_zone"`
	Hour     int64  `json:"hour"`
	Minute   int64  `json:"minute"`
}

type HourlyScheduleConfig struct {
	TimeZone string `json:"time_zone"`
	Minute   int64  `json:"minute"`
}

type MonthlyScheduleConfig struct {
	TimeZone string `json:"time_zone"`
	Day      int64  `json:"day"`
	Hour     int64  `json:"hour"`
	Minute   int64  `json:"minute"`
}

type WeeklyScheduleConfig struct {
	TimeZone  string `json:"time_zone"`
	DayOfWeek int64  `json:"day_of_week"`
	Hour      int64  `json:"hour"`
	Minute    int64  `json:"minute"`
}
