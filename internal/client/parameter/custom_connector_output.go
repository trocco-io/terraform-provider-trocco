package parameter

type CreateCustomConnectorOutputInput struct {
	Name             string                               `json:"name"`
	Description      *NullableString                      `json:"description,omitempty"`
	URL              string                               `json:"url"`
	AuthType         string                               `json:"auth_type"`
	AuthHeaderName   *string                              `json:"auth_header_name,omitempty"`
	AuthHeaderScheme *string                              `json:"auth_header_scheme,omitempty"`
	GrantType        *string                              `json:"grant_type,omitempty"`
	AuthURI          *string                              `json:"auth_uri,omitempty"`
	AccessTokenURI   *string                              `json:"access_token_uri,omitempty"`
	Endpoints        []CustomConnectorOutputEndpointInput `json:"endpoints"`
}

func (input *CreateCustomConnectorOutputInput) SetAuthHeaderName(v string) {
	input.AuthHeaderName = &v
}
func (input *CreateCustomConnectorOutputInput) SetAuthHeaderScheme(v string) {
	input.AuthHeaderScheme = &v
}
func (input *CreateCustomConnectorOutputInput) SetGrantType(v string) { input.GrantType = &v }
func (input *CreateCustomConnectorOutputInput) SetAuthURI(v string)   { input.AuthURI = &v }
func (input *CreateCustomConnectorOutputInput) SetAccessTokenURI(v string) {
	input.AccessTokenURI = &v
}

type UpdateCustomConnectorOutputInput struct {
	Name             *string                              `json:"name,omitempty"`
	Description      *NullableString                      `json:"description,omitempty"`
	URL              *string                              `json:"url,omitempty"`
	AuthType         *string                              `json:"auth_type,omitempty"`
	AuthHeaderName   *string                              `json:"auth_header_name,omitempty"`
	AuthHeaderScheme *string                              `json:"auth_header_scheme,omitempty"`
	GrantType        *string                              `json:"grant_type,omitempty"`
	AuthURI          *string                              `json:"auth_uri,omitempty"`
	AccessTokenURI   *string                              `json:"access_token_uri,omitempty"`
	Endpoints        []CustomConnectorOutputEndpointInput `json:"endpoints"`
}

func (input *UpdateCustomConnectorOutputInput) SetName(v string)     { input.Name = &v }
func (input *UpdateCustomConnectorOutputInput) SetURL(v string)      { input.URL = &v }
func (input *UpdateCustomConnectorOutputInput) SetAuthType(v string) { input.AuthType = &v }
func (input *UpdateCustomConnectorOutputInput) SetAuthHeaderName(v string) {
	input.AuthHeaderName = &v
}
func (input *UpdateCustomConnectorOutputInput) SetAuthHeaderScheme(v string) {
	input.AuthHeaderScheme = &v
}
func (input *UpdateCustomConnectorOutputInput) SetGrantType(v string) { input.GrantType = &v }
func (input *UpdateCustomConnectorOutputInput) SetAuthURI(v string)   { input.AuthURI = &v }
func (input *UpdateCustomConnectorOutputInput) SetAccessTokenURI(v string) {
	input.AccessTokenURI = &v
}

// CustomConnectorOutputEndpointInput is shared by create/update requests. ID is
// set only when preserving an existing endpoint (resolved by matching `name`
// against prior state); left nil, the server creates a new endpoint.
//
// Unlike the transfer source endpoint, every attribute below is always sent:
// the server assigns them unconditionally, so a missing key is written as NULL
// and rejected by the presence validations rather than falling back to the
// column default.
type CustomConnectorOutputEndpointInput struct {
	ID                *int64                              `json:"id,omitempty"`
	Name              string                              `json:"name"`
	RequestType       string                              `json:"request_type"`
	BatchSize         int64                               `json:"batch_size"`
	Method            string                              `json:"method"`
	Operation         string                              `json:"operation"`
	Path              string                              `json:"path"`
	PayloadType       string                              `json:"payload_type"`
	SuccessCodes      string                              `json:"success_codes"`
	NotRetryableCodes string                              `json:"not_retryable_codes"`
	RequestTimeoutSec int64                               `json:"request_timeout_sec"`
	Template          string                              `json:"template"`
	Headers           []CustomConnectorFieldOptionInput   `json:"headers"`
	PathParameters    []CustomConnectorPathParameterInput `json:"path_parameters"`
	QueryParameters   []CustomConnectorFieldOptionInput   `json:"query_parameters"`
}

func (e *CustomConnectorOutputEndpointInput) SetID(v int64) { e.ID = &v }
