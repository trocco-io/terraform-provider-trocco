package parameter

type CreateCustomConnectorInputInput struct {
	Name             string                         `json:"name"`
	Description      *NullableString                `json:"description,omitempty"`
	URL              string                         `json:"url"`
	AuthType         string                         `json:"auth_type"`
	AuthHeaderName   *string                        `json:"auth_header_name,omitempty"`
	AuthHeaderScheme *string                        `json:"auth_header_scheme,omitempty"`
	GrantType        *string                        `json:"grant_type,omitempty"`
	AuthURI          *string                        `json:"auth_uri,omitempty"`
	AccessTokenURI   *string                        `json:"access_token_uri,omitempty"`
	Endpoints        []CustomConnectorEndpointInput `json:"endpoints"`
}

func (input *CreateCustomConnectorInputInput) SetAuthHeaderName(v string) { input.AuthHeaderName = &v }
func (input *CreateCustomConnectorInputInput) SetAuthHeaderScheme(v string) {
	input.AuthHeaderScheme = &v
}
func (input *CreateCustomConnectorInputInput) SetGrantType(v string)      { input.GrantType = &v }
func (input *CreateCustomConnectorInputInput) SetAuthURI(v string)        { input.AuthURI = &v }
func (input *CreateCustomConnectorInputInput) SetAccessTokenURI(v string) { input.AccessTokenURI = &v }

type UpdateCustomConnectorInputInput struct {
	Name             *string                        `json:"name,omitempty"`
	Description      *NullableString                `json:"description,omitempty"`
	URL              *string                        `json:"url,omitempty"`
	AuthType         *string                        `json:"auth_type,omitempty"`
	AuthHeaderName   *string                        `json:"auth_header_name,omitempty"`
	AuthHeaderScheme *string                        `json:"auth_header_scheme,omitempty"`
	GrantType        *string                        `json:"grant_type,omitempty"`
	AuthURI          *string                        `json:"auth_uri,omitempty"`
	AccessTokenURI   *string                        `json:"access_token_uri,omitempty"`
	Endpoints        []CustomConnectorEndpointInput `json:"endpoints"`
}

func (input *UpdateCustomConnectorInputInput) SetName(v string)           { input.Name = &v }
func (input *UpdateCustomConnectorInputInput) SetURL(v string)            { input.URL = &v }
func (input *UpdateCustomConnectorInputInput) SetAuthType(v string)       { input.AuthType = &v }
func (input *UpdateCustomConnectorInputInput) SetAuthHeaderName(v string) { input.AuthHeaderName = &v }
func (input *UpdateCustomConnectorInputInput) SetAuthHeaderScheme(v string) {
	input.AuthHeaderScheme = &v
}
func (input *UpdateCustomConnectorInputInput) SetGrantType(v string)      { input.GrantType = &v }
func (input *UpdateCustomConnectorInputInput) SetAuthURI(v string)        { input.AuthURI = &v }
func (input *UpdateCustomConnectorInputInput) SetAccessTokenURI(v string) { input.AccessTokenURI = &v }

// CustomConnectorEndpointInput is shared by create/update requests. ID is set
// only when preserving an existing endpoint (resolved by matching `name`
// against prior state); left nil, the server creates a new endpoint.
//
// `request_body` and `paginator` deliberately omit `omitempty`: the API keeps
// the current value when a key is absent, so a nil pointer has to be sent as an
// explicit null to clear it. Terraform cannot distinguish "unset" from
// "explicit null", so a null here always means "remove".
type CustomConnectorEndpointInput struct {
	ID                *int64                              `json:"id,omitempty"`
	Name              string                              `json:"name"`
	Path              string                              `json:"path"`
	Method            string                              `json:"method"`
	RequestBody       *string                             `json:"request_body"`
	JsonpathRoot      string                              `json:"jsonpath_root"`
	SuccessCodes      string                              `json:"success_codes"`
	NotRetryableCodes string                              `json:"not_retryable_codes"`
	RequestTimeoutSec *int64                              `json:"request_timeout_sec,omitempty"`
	QueryParameters   []CustomConnectorFieldOptionInput   `json:"query_parameters"`
	Headers           []CustomConnectorFieldOptionInput   `json:"headers"`
	PathParameters    []CustomConnectorPathParameterInput `json:"path_parameters"`
	// RequestBodyParameters declares the placeholders (`{name}`) that
	// `request_body` may contain, so that a transfer setting can fill them in.
	RequestBodyParameters []CustomConnectorFieldOptionInput `json:"request_body_parameters"`
	Paginator             *CustomConnectorPaginatorInput    `json:"paginator"`
}

func (e *CustomConnectorEndpointInput) SetPaginator(v CustomConnectorPaginatorInput) {
	e.Paginator = &v
}

func (e *CustomConnectorEndpointInput) SetID(v int64)                { e.ID = &v }
func (e *CustomConnectorEndpointInput) SetRequestBody(v string)      { e.RequestBody = &v }
func (e *CustomConnectorEndpointInput) SetRequestTimeoutSec(v int64) { e.RequestTimeoutSec = &v }

type CustomConnectorFieldOptionInput struct {
	Name         string  `json:"name"`
	DisplayName  string  `json:"display_name"`
	DefaultValue *string `json:"default_value,omitempty"`
	IsEditable   bool    `json:"is_editable"`
	IsRequired   bool    `json:"is_required"`
}

type CustomConnectorPathParameterInput struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CustomConnectorPaginatorInput struct {
	InjectInto         *string                                `json:"inject_into,omitempty"`
	PaginationStrategy CustomConnectorPaginationStrategyInput `json:"pagination_strategy"`
	PageSizeOption     *CustomConnectorFieldNameOptionInput   `json:"page_size_option,omitempty"`
	PageTokenOption    *CustomConnectorFieldNameOptionInput   `json:"page_token_option,omitempty"`
}

// CustomConnectorPaginationStrategyInput flattens the 3 mutually exclusive TF
// blocks (page_increment_strategy / offset_increment_strategy /
// cursor_based_strategy) into the API's single oneOf payload, discriminated by
// Type. Only the fields relevant to Type should be set by the caller.
type CustomConnectorPaginationStrategyInput struct {
	Type            string  `json:"type"`
	PageSize        *int64  `json:"page_size,omitempty"`
	LastPageSize    *string `json:"last_page_size,omitempty"`
	StartFromPage   *int64  `json:"start_from_page,omitempty"`
	StopOnPage      *int64  `json:"stop_on_page,omitempty"`
	TotalPages      *string `json:"total_pages,omitempty"`
	StartFromOffset *int64  `json:"start_from_offset,omitempty"`
	StopOnOffset    *int64  `json:"stop_on_offset,omitempty"`
	TotalRecords    *string `json:"total_records,omitempty"`
	MaxRequestCount *int64  `json:"max_request_count,omitempty"`
	CursorValue     *string `json:"cursor_value,omitempty"`
}

type CustomConnectorFieldNameOptionInput struct {
	FieldName string `json:"field_name"`
}
