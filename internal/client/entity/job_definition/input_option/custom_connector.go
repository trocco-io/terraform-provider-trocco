package input_option

import (
	jobDefinitionEntities "terraform-provider-trocco/internal/client/entity/job_definition"

	"terraform-provider-trocco/internal/client/entity"
)

// CustomConnectorInputOption models every snapshot field as a pointer: the
// underlying columns are all nullable, and job definitions created before the
// snapshot feature existed carry null in them (the model applies no fallback).
// Decoding those into plain strings would turn null into "", putting an empty
// string in state where null is correct.
type CustomConnectorInputOption struct {
	CustomConnectorEndpointID   *int64  `json:"custom_connector_endpoint_id"`
	CustomConnectorConnectionID *int64  `json:"custom_connector_connection_id"`
	URL                         *string `json:"url"`
	AuthType                    *string `json:"auth_type"`
	AuthHeaderName              *string `json:"auth_header_name"`
	AuthHeaderScheme            *string `json:"auth_header_scheme"`
	EndpointPath                *string `json:"endpoint_path"`
	EndpointMethod              *string `json:"endpoint_method"`
	EndpointRequestBody         *string `json:"endpoint_request_body"`
	SuccessCodes                *string `json:"success_codes"`
	NotRetryableCodes           *string `json:"not_retryable_codes"`
	// RequestTimeoutSec is nullable here (unlike the definition side's
	// NOT NULL column) for the same reason, and null is treated as 30s.
	RequestTimeoutSec      *int64                                `json:"request_timeout_sec"`
	QueryParameters        []CustomConnectorNamedValue           `json:"query_parameters"`
	Headers                []CustomConnectorNamedValue           `json:"headers"`
	PathParameters         []CustomConnectorNamedValue           `json:"path_parameters"`
	RequestBodyParameters  []CustomConnectorNamedValue           `json:"request_body_parameters"`
	Paginator              *CustomConnectorPaginator             `json:"paginator"`
	JsonpathParser         *jobDefinitionEntities.JsonpathParser `json:"jsonpath_parser"`
	CustomVariableSettings *[]entity.CustomVariableSetting       `json:"custom_variable_settings"`
}

// CustomConnectorNamedValue is the shared shape for `query_parameters`,
// `headers` and `path_parameters` (created in request order).
type CustomConnectorNamedValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// CustomConnectorPaginator is a read-only snapshot of the referenced
// endpoint's pagination settings, taken at create/update time. As of
// n-transfer-ui#47352, its wire shape matches the custom connector
// definition's own paginator (internal/client/entity's
// CustomConnectorPaginator): the oneOf(page_increment, offset_increment,
// cursor_based) strategy is flattened into a single discriminated
// pagination_strategy object rather than three separate nullable keys.
// This side still carries a superset of fields (first_offset/first_page/
// inject_on_first_request/stop_on_blank) that the definition side doesn't
// have, so the type isn't reused as-is.
// CustomConnectorPaginator is the snapshot of the referenced endpoint's
// pagination settings.
//
// The API also serializes a `page_end_option` here, which is deliberately not
// decoded: the custom connector definition API does not expose it, so
// `trocco_custom_connector_input` cannot manage it, and surfacing a snapshot
// field with no counterpart on the definition side would be misleading. An
// endpoint configured with it in the web UI therefore reports an incomplete
// snapshot. Add it here (and to the schema/model) once the definition API
// supports it.
type CustomConnectorPaginator struct {
	InjectInto         string                            `json:"inject_into"`
	PaginationStrategy CustomConnectorPaginationStrategy `json:"pagination_strategy"`
	PageSizeOption     *CustomConnectorFieldNameOption   `json:"page_size_option"`
	PageTokenOption    *CustomConnectorFieldNameOption   `json:"page_token_option"`
}

// CustomConnectorPaginationStrategy flattens the oneOf(page_increment,
// offset_increment, cursor_based) into a single struct, discriminated by
// Type. Only the fields relevant to Type are populated by the server.
// last_page_size/total_pages/total_records/cursor_value are stored as
// strings in TROCCO's DB (they may hold a JSONPath expression instead of a
// literal number), so they're modeled as *string here despite the OpenAPI
// document's (inaccurate) integer/boolean typing.
type CustomConnectorPaginationStrategy struct {
	Type                 string  `json:"type"`
	PageSize             *int64  `json:"page_size"`
	StartFromPage        *int64  `json:"start_from_page"`
	FirstPage            *int64  `json:"first_page"`
	InjectOnFirstRequest *bool   `json:"inject_on_first_request"`
	LastPageSize         *string `json:"last_page_size"`
	TotalPages           *string `json:"total_pages"`
	StopOnPage           *int64  `json:"stop_on_page"`
	StartFromOffset      *int64  `json:"start_from_offset"`
	FirstOffset          *int64  `json:"first_offset"`
	TotalRecords         *string `json:"total_records"`
	StopOnOffset         *int64  `json:"stop_on_offset"`
	MaxRequestCount      *int64  `json:"max_request_count"`
	CursorValue          *string `json:"cursor_value"`
	StopOnBlank          *string `json:"stop_on_blank"`
}

type CustomConnectorFieldNameOption struct {
	FieldName string `json:"field_name"`
}
