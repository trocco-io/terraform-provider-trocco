package pipeline_definition

type HTTPRequestTaskConfig struct {
	Name              string             `json:"name"`
	ConnectionID      *int64             `json:"connection_id"`
	HTTPMethod        string             `json:"http_method"`
	URL               string             `json:"url"`
	RequestBody       *string            `json:"request_body"`
	RequestHeaders    []RequestHeader    `json:"request_headers"`
	RequestParameters []RequestParameter `json:"request_parameter"`
	CustomVariables   []CustomVariable   `json:"custom_variables"`
}

type RequestHeader struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Masking bool   `json:"masking"`
}

type RequestParameter struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Masking bool   `json:"masking"`
}
