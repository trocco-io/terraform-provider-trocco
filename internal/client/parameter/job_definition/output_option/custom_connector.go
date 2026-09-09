package output_options

import "terraform-provider-trocco/internal/client/parameter"

type CustomConnectorOutputOptionInput struct {
	CustomConnectorConnectionID           int64                                   `json:"custom_connector_connection_id"`
	Mode                                  string                                  `json:"mode"`
	UpdateKey                             *string                                 `json:"update_key,omitempty"`
	CreateCustomConnectorOutputEndpointID int64                                   `json:"create_custom_connector_output_endpoint_id"`
	UpdateCustomConnectorOutputEndpointID *int64                                  `json:"update_custom_connector_output_endpoint_id,omitempty"`
	CreateEndpointSettings                *CustomConnectorEndpointSettingsInput   `json:"create_endpoint_settings,omitempty"`
	UpdateEndpointSettings                *CustomConnectorEndpointSettingsInput   `json:"update_endpoint_settings,omitempty"`
	CustomVariableSettings                *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

// UpdateCustomConnectorOutputOptionInput omits every key the configuration
// leaves unset, since the server falls back to the stored value for any key
// absent from the request. `update_key` and
// `update_custom_connector_output_endpoint_id` are therefore never sent while
// `mode` is `insert`: the server clears both on its own once it sees that mode,
// and sending them explicitly is rejected.
type UpdateCustomConnectorOutputOptionInput struct {
	CustomConnectorConnectionID           *int64                                  `json:"custom_connector_connection_id,omitempty"`
	Mode                                  *string                                 `json:"mode,omitempty"`
	UpdateKey                             *string                                 `json:"update_key,omitempty"`
	CreateCustomConnectorOutputEndpointID *int64                                  `json:"create_custom_connector_output_endpoint_id,omitempty"`
	UpdateCustomConnectorOutputEndpointID *int64                                  `json:"update_custom_connector_output_endpoint_id,omitempty"`
	CreateEndpointSettings                *CustomConnectorEndpointSettingsInput   `json:"create_endpoint_settings,omitempty"`
	UpdateEndpointSettings                *CustomConnectorEndpointSettingsInput   `json:"update_endpoint_settings,omitempty"`
	CustomVariableSettings                *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

// CustomConnectorEndpointSettingsInput carries the editable values of one
// endpoint. Each collection is a pointer so the API's null=keep / []=clear-all
// / values=full-replace update semantics can be expressed: a nil pointer omits
// the key entirely and the server preserves the stored values.
type CustomConnectorEndpointSettingsInput struct {
	Headers         *[]CustomConnectorNamedValueInput `json:"headers,omitempty"`
	PathParameters  *[]CustomConnectorNamedValueInput `json:"path_parameters,omitempty"`
	QueryParameters *[]CustomConnectorNamedValueInput `json:"query_parameters,omitempty"`
}

type CustomConnectorNamedValueInput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
