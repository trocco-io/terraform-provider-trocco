package filter

type FilterAddTimeInput struct {
	ColumnName      string  `json:"column_name"`
	Type            string  `json:"type"`
	TimestampFormat *string `json:"timestamp_format,omitempty"`
	TimeZone        *string `json:"time_zone,omitempty"`
}
