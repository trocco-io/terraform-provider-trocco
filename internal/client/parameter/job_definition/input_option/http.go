package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
	jobDefinitions "terraform-provider-trocco/internal/client/parameter/job_definition"
)

type RequestParamInput struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Masking *bool  `json:"masking,omitempty"`
}

type RequestHeaderInput struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Masking *bool  `json:"masking,omitempty"`
}

type HTTPInputOptionInput struct {
	URL                                   string                                  `json:"url"`
	Method                                string                                  `json:"method"`
	UserAgent                             *parameter.NullableString               `json:"user_agent,omitempty"`
	Charset                               *parameter.NullableString               `json:"charset,omitempty"`
	PagerType                             *parameter.NullableString               `json:"pager_type,omitempty"`
	PagerFromParam                        *parameter.NullableString               `json:"pager_from_param,omitempty"`
	PagerToParam                          *parameter.NullableString               `json:"pager_to_param,omitempty"`
	PagerPages                            *parameter.NullableInt64                `json:"pager_pages,omitempty"`
	PagerStart                            *parameter.NullableInt64                `json:"pager_start,omitempty"`
	PagerStep                             *parameter.NullableInt64                `json:"pager_step,omitempty"`
	CursorRequestParameterCursorName      *parameter.NullableString               `json:"cursor_request_parameter_cursor_name,omitempty"`
	CursorResponseParameterCursorJsonPath *parameter.NullableString               `json:"cursor_response_parameter_cursor_json_path,omitempty"`
	CursorRequestParameterLimitName       *parameter.NullableString               `json:"cursor_request_parameter_limit_name,omitempty"`
	CursorRequestParameterLimitValue      *parameter.NullableString               `json:"cursor_request_parameter_limit_value,omitempty"`
	RequestParams                         *[]RequestParamInput                    `json:"request_params,omitempty"`
	RequestBody                           *parameter.NullableString               `json:"request_body,omitempty"`
	RequestHeaders                        *[]RequestHeaderInput                   `json:"request_headers,omitempty"`
	SuccessCode                           *parameter.NullableString               `json:"success_code,omitempty"`
	OpenTimeout                           *parameter.NullableInt64                `json:"open_timeout,omitempty"`
	ReadTimeout                           *parameter.NullableInt64                `json:"read_timeout,omitempty"`
	MaxRetries                            *parameter.NullableInt64                `json:"max_retries,omitempty"`
	RetryInterval                         *parameter.NullableInt64                `json:"retry_interval,omitempty"`
	RequestInterval                       *parameter.NullableInt64                `json:"request_interval,omitempty"`
	CustomVariableSettings                *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	CsvParser                             *jobDefinitions.CsvParserInput          `json:"csv_parser,omitempty"`
	JsonlParser                           *jobDefinitions.JsonlParserInput        `json:"jsonl_parser,omitempty"`
	JsonpathParser                        *jobDefinitions.JsonpathParserInput     `json:"jsonpath_parser,omitempty"`
	LtsvParser                            *jobDefinitions.LtsvParserInput         `json:"ltsv_parser,omitempty"`
	ExcelParser                           *jobDefinitions.ExcelParserInput        `json:"excel_parser,omitempty"`
	XmlParser                             *jobDefinitions.XmlParserInput          `json:"xml_parser,omitempty"`
}

type UpdateHTTPInputOptionInput struct {
	URL                                   *string                                 `json:"url,omitempty"`
	Method                                *string                                 `json:"method,omitempty"`
	UserAgent                             *parameter.NullableString               `json:"user_agent,omitempty"`
	Charset                               *parameter.NullableString               `json:"charset,omitempty"`
	PagerType                             *parameter.NullableString               `json:"pager_type,omitempty"`
	PagerFromParam                        *parameter.NullableString               `json:"pager_from_param,omitempty"`
	PagerToParam                          *parameter.NullableString               `json:"pager_to_param,omitempty"`
	PagerPages                            *parameter.NullableInt64                `json:"pager_pages,omitempty"`
	PagerStart                            *parameter.NullableInt64                `json:"pager_start,omitempty"`
	PagerStep                             *parameter.NullableInt64                `json:"pager_step,omitempty"`
	CursorRequestParameterCursorName      *parameter.NullableString               `json:"cursor_request_parameter_cursor_name,omitempty"`
	CursorResponseParameterCursorJsonPath *parameter.NullableString               `json:"cursor_response_parameter_cursor_json_path,omitempty"`
	CursorRequestParameterLimitName       *parameter.NullableString               `json:"cursor_request_parameter_limit_name,omitempty"`
	CursorRequestParameterLimitValue      *parameter.NullableString               `json:"cursor_request_parameter_limit_value,omitempty"`
	RequestParams                         *[]RequestParamInput                    `json:"request_params,omitempty"`
	RequestBody                           *parameter.NullableString               `json:"request_body,omitempty"`
	RequestHeaders                        *[]RequestHeaderInput                   `json:"request_headers,omitempty"`
	SuccessCode                           *parameter.NullableString               `json:"success_code,omitempty"`
	OpenTimeout                           *parameter.NullableInt64                `json:"open_timeout,omitempty"`
	ReadTimeout                           *parameter.NullableInt64                `json:"read_timeout,omitempty"`
	MaxRetries                            *parameter.NullableInt64                `json:"max_retries,omitempty"`
	RetryInterval                         *parameter.NullableInt64                `json:"retry_interval,omitempty"`
	RequestInterval                       *parameter.NullableInt64                `json:"request_interval,omitempty"`
	CustomVariableSettings                *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	CsvParser                             *jobDefinitions.CsvParserInput          `json:"csv_parser,omitempty"`
	JsonlParser                           *jobDefinitions.JsonlParserInput        `json:"jsonl_parser,omitempty"`
	JsonpathParser                        *jobDefinitions.JsonpathParserInput     `json:"jsonpath_parser,omitempty"`
	LtsvParser                            *jobDefinitions.LtsvParserInput         `json:"ltsv_parser,omitempty"`
	ExcelParser                           *jobDefinitions.ExcelParserInput        `json:"excel_parser,omitempty"`
	XmlParser                             *jobDefinitions.XmlParserInput          `json:"xml_parser,omitempty"`
}
