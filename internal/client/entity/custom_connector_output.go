package entity

type CustomConnectorOutput struct {
	ID               int64                           `json:"id"`
	Name             string                          `json:"name"`
	Description      *string                         `json:"description"`
	URL              string                          `json:"url"`
	AuthType         string                          `json:"auth_type"`
	AuthHeaderName   string                          `json:"auth_header_name"`
	AuthHeaderScheme *string                         `json:"auth_header_scheme"`
	GrantType        *string                         `json:"grant_type"`
	AuthURI          *string                         `json:"auth_uri"`
	AccessTokenURI   *string                         `json:"access_token_uri"`
	Endpoints        []CustomConnectorOutputEndpoint `json:"endpoints"`
	CreatedAt        string                          `json:"created_at"`
	UpdatedAt        string                          `json:"updated_at"`
}

// CustomConnectorOutputEndpoint has no paginator: transfer destination
// endpoints send records rather than paging through responses. `headers` and
// `query_parameters` reuse the FieldOption shape shared with the transfer
// source definition.
type CustomConnectorOutputEndpoint struct {
	ID                int64                          `json:"id"`
	Name              string                         `json:"name"`
	RequestType       string                         `json:"request_type"`
	BatchSize         int64                          `json:"batch_size"`
	Method            string                         `json:"method"`
	Operation         string                         `json:"operation"`
	Path              string                         `json:"path"`
	PayloadType       string                         `json:"payload_type"`
	SuccessCodes      string                         `json:"success_codes"`
	NotRetryableCodes string                         `json:"not_retryable_codes"`
	RequestTimeoutSec int64                          `json:"request_timeout_sec"`
	Template          string                         `json:"template"`
	Headers           []CustomConnectorFieldOption   `json:"headers"`
	PathParameters    []CustomConnectorPathParameter `json:"path_parameters"`
	QueryParameters   []CustomConnectorFieldOption   `json:"query_parameters"`
}
