package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
	job_definitions "terraform-provider-trocco/internal/client/parameter/job_definition"
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

type HttpInputOptionInput struct {
	URL                                   string                                  `json:"url"`
	Method                                string                                  `json:"method"`
	UserAgent                             *parameter.NullableString               `json:"user_agent,omitempty"`
	Charset                               *parameter.NullableString               `json:"charset,omitempty"`
	PagerType                             string                                  `json:"pager_type"`
	PagerFromParam                        *parameter.NullableString               `json:"pager_from_param,omitempty"`
	PagerToParam                          *parameter.NullableString               `json:"pager_to_param,omitempty"`
	PagerPages                            *int64                                  `json:"pager_pages,omitempty"`
	PagerStart                            *int64                                  `json:"pager_start,omitempty"`
	PagerStep                             *int64                                  `json:"pager_step,omitempty"`
	CursorRequestParameterCursorName      *parameter.NullableString               `json:"cursor_request_parameter_cursor_name,omitempty"`
	CursorResponseParameterCursorJsonPath *parameter.NullableString               `json:"cursor_response_parameter_cursor_json_path,omitempty"`
	CursorRequestParameterLimitName       *parameter.NullableString               `json:"cursor_request_parameter_limit_name,omitempty"`
	CursorRequestParameterLimitValue      *int64                                  `json:"cursor_request_parameter_limit_value,omitempty"`
	RequestParams                         *[]RequestParamInput                    `json:"request_params,omitempty"`
	RequestBody                           *parameter.NullableString               `json:"request_body,omitempty"`
	RequestHeaders                        *[]RequestHeaderInput                   `json:"request_headers,omitempty"`
	SuccessCode                           *parameter.NullableString               `json:"success_code,omitempty"`
	OpenTimeout                           *int64                                  `json:"open_timeout,omitempty"`
	ReadTimeout                           *int64                                  `json:"read_timeout,omitempty"`
	MaxRetries                            *int64                                  `json:"max_retries,omitempty"`
	RetryInterval                         *int64                                  `json:"retry_interval,omitempty"`
	RequestInterval                       *int64                                  `json:"request_interval,omitempty"`
	CustomVariableSettings                *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	CsvParser                             *job_definitions.CsvParserInput         `json:"csv_parser,omitempty"`
	JsonlParser                           *job_definitions.JsonlParserInput       `json:"jsonl_parser,omitempty"`
	JsonpathParser                        *job_definitions.JsonpathParserInput    `json:"jsonpath_parser,omitempty"`
	LtsvParser                            *job_definitions.LtsvParserInput        `json:"ltsv_parser,omitempty"`
	ExcelParser                           *job_definitions.ExcelParserInput       `json:"excel_parser,omitempty"`
	XmlParser                             *job_definitions.XmlParserInput         `json:"xml_parser,omitempty"`
}

type UpdateHttpInputOptionInput struct {
	URL                                   *string                                 `json:"url,omitempty"`
	Method                                *string                                 `json:"method,omitempty"`
	UserAgent                             *parameter.NullableString               `json:"user_agent,omitempty"`
	Charset                               *parameter.NullableString               `json:"charset,omitempty"`
	PagerType                             *string                                 `json:"pager_type,omitempty"`
	PagerFromParam                        *parameter.NullableString               `json:"pager_from_param,omitempty"`
	PagerToParam                          *parameter.NullableString               `json:"pager_to_param,omitempty"`
	PagerPages                            *int64                                  `json:"pager_pages,omitempty"`
	PagerStart                            *int64                                  `json:"pager_start,omitempty"`
	PagerStep                             *int64                                  `json:"pager_step,omitempty"`
	CursorRequestParameterCursorName      *parameter.NullableString               `json:"cursor_request_parameter_cursor_name,omitempty"`
	CursorResponseParameterCursorJsonPath *parameter.NullableString               `json:"cursor_response_parameter_cursor_json_path,omitempty"`
	CursorRequestParameterLimitName       *parameter.NullableString               `json:"cursor_request_parameter_limit_name,omitempty"`
	CursorRequestParameterLimitValue      *int64                                  `json:"cursor_request_parameter_limit_value,omitempty"`
	RequestParams                         *[]RequestParamInput                    `json:"request_params,omitempty"`
	RequestBody                           *parameter.NullableString               `json:"request_body,omitempty"`
	RequestHeaders                        *[]RequestHeaderInput                   `json:"request_headers,omitempty"`
	SuccessCode                           *parameter.NullableString               `json:"success_code,omitempty"`
	OpenTimeout                           *int64                                  `json:"open_timeout,omitempty"`
	ReadTimeout                           *int64                                  `json:"read_timeout,omitempty"`
	MaxRetries                            *int64                                  `json:"max_retries,omitempty"`
	RetryInterval                         *int64                                  `json:"retry_interval,omitempty"`
	RequestInterval                       *int64                                  `json:"request_interval,omitempty"`
	CustomVariableSettings                *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
	CsvParser                             *job_definitions.CsvParserInput         `json:"csv_parser,omitempty"`
	JsonlParser                           *job_definitions.JsonlParserInput       `json:"jsonl_parser,omitempty"`
	JsonpathParser                        *job_definitions.JsonpathParserInput    `json:"jsonpath_parser,omitempty"`
	LtsvParser                            *job_definitions.LtsvParserInput        `json:"ltsv_parser,omitempty"`
	ExcelParser                           *job_definitions.ExcelParserInput       `json:"excel_parser,omitempty"`
	XmlParser                             *job_definitions.XmlParserInput         `json:"xml_parser,omitempty"`
}
