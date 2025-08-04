package pipeline_definition

import (
	parameter "terraform-provider-trocco/internal/client/parameter"
)

type HTTPRequestTaskConfig struct {
	Name              string                   `json:"name,omitempty"`
	ConnectionID      *parameter.NullableInt64 `json:"connection_id,omitempty"`
	HTTPMethod        string                   `json:"http_method,omitempty"`
	URL               string                   `json:"url,omitempty"`
	RequestBody       *string                  `json:"request_body,omitempty"`
	RequestHeaders    []RequestHeader          `json:"request_headers,omitempty"`
	RequestParameters []RequestParameter       `json:"request_parameters,omitempty"`
	CustomVariables   []CustomVariable         `json:"custom_variables,omitempty"`
}

type RequestHeader struct {
	Key     string                  `json:"key,omitempty"`
	Value   string                  `json:"value,omitempty"`
	Masking *parameter.NullableBool `json:"masking,omitempty"`
}

type RequestParameter struct {
	Key     string                  `json:"key,omitempty"`
	Value   string                  `json:"value,omitempty"`
	Masking *parameter.NullableBool `json:"masking,omitempty"`
}
