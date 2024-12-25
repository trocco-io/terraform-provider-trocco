package parameters

type CustomVariableSettingInput struct {
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	Value     *string        `json:"value,omitempty"`
	Quantity  *NullableInt64 `json:"quantity,omitempty"`
	Unit      *string        `json:"unit,omitempty"`
	Direction *string        `json:"direction,omitempty"`
	Format    *string        `json:"format,omitempty"`
	TimeZone  *string        `json:"time_zone,omitempty"`
}
