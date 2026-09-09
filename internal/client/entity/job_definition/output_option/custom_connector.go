package output_option

import "terraform-provider-trocco/internal/client/entity"

// CustomConnectorOutputOption is the job definition side of a custom
// connector destination. Everything from URL onwards is a read-only snapshot
// taken from the referenced custom connector definition
// (`trocco_custom_connector_output`) at create/update time.
type CustomConnectorOutputOption struct {
	CustomConnectorConnectionID *int64  `json:"custom_connector_connection_id"`
	Mode                        string  `json:"mode"`
	UpdateKey                   *string `json:"update_key"`
	// Create/UpdateCustomConnectorOutputEndpointID are nullable because the
	// server nullifies them when the referenced endpoint is deleted from the
	// custom connector definition.
	CreateCustomConnectorOutputEndpointID *int64                                `json:"create_custom_connector_output_endpoint_id"`
	UpdateCustomConnectorOutputEndpointID *int64                                `json:"update_custom_connector_output_endpoint_id"`
	URL                                   string                                `json:"url"`
	AuthType                              string                                `json:"auth_type"`
	AuthHeaderName                        *string                               `json:"auth_header_name"`
	AuthHeaderScheme                      *string                               `json:"auth_header_scheme"`
	Endpoints                             []CustomConnectorOutputOptionEndpoint `json:"endpoints"`
	CustomVariableSettings                *[]entity.CustomVariableSetting       `json:"custom_variable_settings"`
}

// CustomConnectorOutputOptionEndpoint is the snapshot of one endpoint used by
// this job definition, distinguished by Operation (`create` or `update`). The
// `update` entry is absent while Mode is `insert`.
type CustomConnectorOutputOptionEndpoint struct {
	Operation         string                      `json:"operation"`
	RequestType       string                      `json:"request_type"`
	BatchSize         int64                       `json:"batch_size"`
	Method            string                      `json:"method"`
	Path              string                      `json:"path"`
	PayloadType       string                      `json:"payload_type"`
	SuccessCodes      string                      `json:"success_codes"`
	NotRetryableCodes string                      `json:"not_retryable_codes"`
	RequestTimeoutSec int64                       `json:"request_timeout_sec"`
	Template          *string                     `json:"template"`
	Headers           []CustomConnectorNamedValue `json:"headers"`
	PathParameters    []CustomConnectorNamedValue `json:"path_parameters"`
	QueryParameters   []CustomConnectorNamedValue `json:"query_parameters"`
}

// CustomConnectorNamedValue is the shared shape for the snapshot's `headers`,
// `path_parameters` and `query_parameters` entries (created in request order).
type CustomConnectorNamedValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
