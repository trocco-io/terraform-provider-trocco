package filter

type FilterUnixTimeConversion struct {
	ColumnName       string `json:"column_name"`
	Kind             string `json:"kind"`
	UnixtimeUnit     string `json:"unixtime_unit"`
	DatetimeFormat   string `json:"datetime_format"`
	DatetimeTimezone string `json:"datetime_timezone"`
}
