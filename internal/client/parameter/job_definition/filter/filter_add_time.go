package filter

import "terraform-provider-trocco/internal/client/parameter"

type FilterAddTimeInput struct {
	ColumnName      string                    `json:"column_name"`
	Type            string                    `json:"type"`
	TimestampFormat *parameter.NullableString `json:"timestamp_format,omitempty"`
	TimeZone        *parameter.NullableString `json:"time_zone,omitempty"`
}
