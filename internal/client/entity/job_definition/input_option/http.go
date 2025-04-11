package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
	job_definitions "terraform-provider-trocco/internal/client/entity/job_definition"
)

type RequestParam struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Masking *bool  `json:"masking"`
}
type RequestHeader struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Masking *bool  `json:"masking"`
}

type HttpInputOption struct {
	URL                                   string                          `json:"url"`
	Method                                string                          `json:"method"`
	UserAgent                             *string                         `json:"user_agent"`
	Charset                               *string                         `json:"charset"`
	PagerType                             string                          `json:"pager_type"`
	PagerFromParam                        *string                         `json:"pager_from_param"`
	PagerToParam                          *string                         `json:"pager_to_param"`
	PagerPages                            *int64                          `json:"pager_pages"`
	PagerStart                            *int64                          `json:"pager_start"`
	PagerStep                             *int64                          `json:"pager_step"`
	CursorRequestParameterCursorName      *string                         `json:"cursor_request_parameter_cursor_name"`
	CursorResponseParameterCursorJsonPath *string                         `json:"cursor_response_parameter_cursor_json_path"`
	CursorRequestParameterLimitName       *string                         `json:"cursor_request_parameter_limit_name"`
	CursorRequestParameterLimitValue      *int64                          `json:"cursor_request_parameter_limit_value"`
	RequestParams                         *[]RequestParam                 `json:"request_params"`
	RequestBody                           *string                         `json:"request_body"`
	RequestHeaders                        *[]RequestHeader                `json:"request_headers"`
	SuccessCode                           *string                         `json:"success_code"`
	OpenTimeout                           *int64                          `json:"open_timeout"`
	ReadTimeout                           *int64                          `json:"read_timeout"`
	MaxRetries                            *int64                          `json:"max_retries"`
	RetryInterval                         *int64                          `json:"retry_interval"`
	RequestInterval                       *int64                          `json:"request_interval"`
	CustomVariableSettings                *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
	CsvParser                             *job_definitions.CsvParser      `json:"csv_parser"`
	JsonlParser                           *job_definitions.JsonlParser    `json:"jsonl_parser"`
	JsonpathParser                        *job_definitions.JsonpathParser `json:"jsonpath_parser"`
	LtsvParser                            *job_definitions.LtsvParser     `json:"ltsv_parser"`
	ExcelParser                           *job_definitions.ExcelParser    `json:"excel_parser"`
	XmlParser                             *job_definitions.XmlParser      `json:"xml_parser"`
}
