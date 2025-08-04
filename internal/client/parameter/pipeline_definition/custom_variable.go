package pipeline_definition

import (
	parameter "terraform-provider-trocco/internal/client/parameter"
)

type CustomVariable struct {
	Name      *string                  `json:"name,omitempty"`
	Type      *string                  `json:"type,omitempty"`
	Value     *string                  `json:"value,omitempty"`
	Quantity  *parameter.NullableInt64 `json:"quantity,omitempty"`
	Unit      *string                  `json:"unit,omitempty"`
	Direction *string                  `json:"direction,omitempty"`
	Format    *string                  `json:"format,omitempty"`
	TimeZone  *string                  `json:"time_zone,omitempty"`
}
