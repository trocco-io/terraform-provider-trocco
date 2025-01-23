package filter

type FilterAddTime struct {
	ColumnName      string  `json:"column_name"`
	Type            string  `json:"type"`
	TimestampFormat *string `json:"timestamp_format"`
	TimeZone        *string `json:"time_zone"`
}
