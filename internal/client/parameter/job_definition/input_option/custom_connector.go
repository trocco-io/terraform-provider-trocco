package input_options

import (
	jobDefinitionParameters "terraform-provider-trocco/internal/client/parameter/job_definition"

	"terraform-provider-trocco/internal/client/parameter"
)

type CustomConnectorInputOptionInput struct {
	CustomConnectorEndpointID   int64                                        `json:"custom_connector_endpoint_id"`
	CustomConnectorConnectionID int64                                        `json:"custom_connector_connection_id"`
	QueryParameters             *[]CustomConnectorNamedValueInput            `json:"query_parameters,omitempty"`
	Headers                     *[]CustomConnectorNamedValueInput            `json:"headers,omitempty"`
	PathParameters              *[]CustomConnectorNamedValueInput            `json:"path_parameters,omitempty"`
	RequestBodyParameters       *[]CustomConnectorNamedValueInput            `json:"request_body_parameters,omitempty"`
	JsonpathParser              *jobDefinitionParameters.JsonpathParserInput `json:"jsonpath_parser"`
	CustomVariableSettings      *[]parameter.CustomVariableSettingInput      `json:"custom_variable_settings,omitempty"`
}

type UpdateCustomConnectorInputOptionInput struct {
	CustomConnectorEndpointID   *parameter.NullableInt64                     `json:"custom_connector_endpoint_id,omitempty"`
	CustomConnectorConnectionID *parameter.NullableInt64                     `json:"custom_connector_connection_id,omitempty"`
	QueryParameters             *[]CustomConnectorNamedValueInput            `json:"query_parameters,omitempty"`
	Headers                     *[]CustomConnectorNamedValueInput            `json:"headers,omitempty"`
	PathParameters              *[]CustomConnectorNamedValueInput            `json:"path_parameters,omitempty"`
	RequestBodyParameters       *[]CustomConnectorNamedValueInput            `json:"request_body_parameters,omitempty"`
	JsonpathParser              *jobDefinitionParameters.JsonpathParserInput `json:"jsonpath_parser,omitempty"`
	CustomVariableSettings      *[]parameter.CustomVariableSettingInput      `json:"custom_variable_settings,omitempty"`
}

// CustomConnectorNamedValueInput is the shared shape for `query_parameters`,
// `headers` and `path_parameters` request items.
type CustomConnectorNamedValueInput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
