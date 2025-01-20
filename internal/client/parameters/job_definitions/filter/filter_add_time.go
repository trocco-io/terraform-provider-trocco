package filter

import "terraform-provider-trocco/internal/client/parameters"

type FilterAddTimeInput struct {
	ColumnName      string                     `json:"column_name"`
	Type            string                     `json:"type"`
	TimestampFormat *parameters.NullableString `json:"timestamp_format,omitempty"`
	TimeZone        *parameters.NullableString `json:"time_zone,omitempty"`
}
