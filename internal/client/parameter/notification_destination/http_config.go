package notification_destination

import "terraform-provider-trocco/internal/client/parameter"

type HTTPConfigInput struct {
	Name        *string              `json:"name,omitempty"`
	URL         *string              `json:"url,omitempty"`
	Description *parameter.NullableString `json:"description,omitempty"`
	Headers     *[]HTTPKeyValueInput `json:"headers,omitempty"`
	QueryParams *[]HTTPKeyValueInput `json:"query_params,omitempty"`
}

type HTTPKeyValueInput struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Masking bool   `json:"masking"`
}
