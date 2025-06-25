package entity

type CustomVariableSetting struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Value     *string `json:"value"`
	Quantity  *int64  `json:"quantity"`
	Unit      *string `json:"unit"`
	Direction *string `json:"direction"`
	Format    *string `json:"format"`
	TimeZone  *string `json:"time_zone"`
}
