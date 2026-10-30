package entity

type CustomConnectorInput struct {
	ID               int64                     `json:"id"`
	Name             string                    `json:"name"`
	Description      *string                   `json:"description"`
	URL              string                    `json:"url"`
	AuthType         string                    `json:"auth_type"`
	AuthHeaderName   string                    `json:"auth_header_name"`
	AuthHeaderScheme *string                   `json:"auth_header_scheme"`
	GrantType        *string                   `json:"grant_type"`
	AuthURI          *string                   `json:"auth_uri"`
	AccessTokenURI   *string                   `json:"access_token_uri"`
	Endpoints        []CustomConnectorEndpoint `json:"endpoints"`
	CreatedAt        string                    `json:"created_at"`
	UpdatedAt        string                    `json:"updated_at"`
}

type CustomConnectorEndpoint struct {
	ID                int64                          `json:"id"`
	Name              string                         `json:"name"`
	Path              string                         `json:"path"`
	Method            string                         `json:"method"`
	RequestBody       *string                        `json:"request_body"`
	JsonpathRoot      string                         `json:"jsonpath_root"`
	SuccessCodes      string                         `json:"success_codes"`
	NotRetryableCodes string                         `json:"not_retryable_codes"`
	RequestTimeoutSec int64                          `json:"request_timeout_sec"`
	QueryParameters   []CustomConnectorFieldOption   `json:"query_parameters"`
	Headers           []CustomConnectorFieldOption   `json:"headers"`
	PathParameters    []CustomConnectorPathParameter `json:"path_parameters"`
	// RequestBodyParameters declares the placeholders (`{name}`) that
	// `request_body` may contain, so that a transfer setting can fill them in.
	RequestBodyParameters []CustomConnectorFieldOption `json:"request_body_parameters"`
	Paginator             *CustomConnectorPaginator    `json:"paginator"`
}

// CustomConnectorFieldOption is the shared shape for an endpoint's
// `query_parameters`, `headers` and `request_body_parameters`.
type CustomConnectorFieldOption struct {
	Name         string  `json:"name"`
	DisplayName  string  `json:"display_name"`
	DefaultValue *string `json:"default_value"`
	IsEditable   bool    `json:"is_editable"`
	IsRequired   bool    `json:"is_required"`
}

type CustomConnectorPathParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CustomConnectorPaginator struct {
	InjectInto         string                            `json:"inject_into"`
	PaginationStrategy CustomConnectorPaginationStrategy `json:"pagination_strategy"`
	PageSizeOption     *CustomConnectorFieldNameOption   `json:"page_size_option"`
	PageTokenOption    *CustomConnectorFieldNameOption   `json:"page_token_option"`
}

// CustomConnectorPaginationStrategy flattens the API's oneOf(page_increment,
// offset_increment, cursor_based) into a single struct, discriminated by Type.
// Only the fields relevant to Type are populated by the server.
type CustomConnectorPaginationStrategy struct {
	Type            string  `json:"type"`
	PageSize        *int64  `json:"page_size"`
	LastPageSize    *string `json:"last_page_size"`
	StartFromPage   *int64  `json:"start_from_page"`
	StopOnPage      *int64  `json:"stop_on_page"`
	TotalPages      *string `json:"total_pages"`
	StartFromOffset *int64  `json:"start_from_offset"`
	StopOnOffset    *int64  `json:"stop_on_offset"`
	TotalRecords    *string `json:"total_records"`
	MaxRequestCount *int64  `json:"max_request_count"`
	CursorValue     *string `json:"cursor_value"`
}

type CustomConnectorFieldNameOption struct {
	FieldName string `json:"field_name"`
}
