package input_options

import (
	entity "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	parameter "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/input_option/parser"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RequestParam struct {
	Key     types.String `tfsdk:"key"`
	Value   types.String `tfsdk:"value"`
	Masking types.Bool   `tfsdk:"masking"`
}

type RequestHeader struct {
	Key     types.String `tfsdk:"key"`
	Value   types.String `tfsdk:"value"`
	Masking types.Bool   `tfsdk:"masking"`
}

type HttpInputOption struct {
	URL                                   types.String                   `tfsdk:"url"`
	Method                                types.String                   `tfsdk:"method"`
	UserAgent                             types.String                   `tfsdk:"user_agent"`
	Charset                               types.String                   `tfsdk:"charset"`
	PagerType                             types.String                   `tfsdk:"pager_type"`
	PagerFromParam                        types.String                   `tfsdk:"pager_from_param"`
	PagerToParam                          types.String                   `tfsdk:"pager_to_param"`
	PagerPages                            types.Int64                    `tfsdk:"pager_pages"`
	PagerStart                            types.Int64                    `tfsdk:"pager_start"`
	PagerStep                             types.Int64                    `tfsdk:"pager_step"`
	CursorRequestParameterCursorName      types.String                   `tfsdk:"cursor_request_parameter_cursor_name"`
	CursorResponseParameterCursorJsonPath types.String                   `tfsdk:"cursor_response_parameter_cursor_json_path"`
	CursorRequestParameterLimitName       types.String                   `tfsdk:"cursor_request_parameter_limit_name"`
	CursorRequestParameterLimitValue      types.Int64                    `tfsdk:"cursor_request_parameter_limit_value"`
	RequestParams                         []RequestParam                 `tfsdk:"request_params"`
	RequestBody                           types.String                   `tfsdk:"request_body"`
	RequestHeaders                        []RequestHeader                `tfsdk:"request_headers"`
	SuccessCode                           types.String                   `tfsdk:"success_code"`
	OpenTimeout                           types.Int64                    `tfsdk:"open_timeout"`
	ReadTimeout                           types.Int64                    `tfsdk:"read_timeout"`
	MaxRetries                            types.Int64                    `tfsdk:"max_retries"`
	RetryInterval                         types.Int64                    `tfsdk:"retry_interval"`
	RequestInterval                       types.Int64                    `tfsdk:"request_interval"`
	CsvParser                             *parser.CsvParser              `tfsdk:"csv_parser"`
	JsonlParser                           *parser.JsonlParser            `tfsdk:"jsonl_parser"`
	JsonpathParser                        *parser.JsonpathParser         `tfsdk:"jsonpath_parser"`
	LtsvParser                            *parser.LtsvParser             `tfsdk:"ltsv_parser"`
	ExcelParser                           *parser.ExcelParser            `tfsdk:"excel_parser"`
	XmlParser                             *parser.XmlParser              `tfsdk:"xml_parser"`
	ParquetParser                         *parser.ParquetParser          `tfsdk:"parquet_parser"`
	CustomVariableSettings                *[]model.CustomVariableSetting `tfsdk:"custom_variable_settings"`
}

func NewHttpInputOption(httpInputOption *entity.HttpInputOption, previous *HttpInputOption) *HttpInputOption {
	if httpInputOption == nil {
		return nil
	}

	var previousRequestHeaders []RequestHeader
	var previousRequestParameters []RequestParam
	if previous != nil {
		previousRequestHeaders = previous.RequestHeaders
		previousRequestParameters = previous.RequestParams
	}

	return &HttpInputOption{
		URL:                                   types.StringValue(httpInputOption.URL),
		Method:                                types.StringValue(httpInputOption.Method),
		UserAgent:                             types.StringPointerValue(httpInputOption.UserAgent),
		Charset:                               types.StringPointerValue(httpInputOption.Charset),
		PagerType:                             types.StringPointerValue(httpInputOption.PagerType),
		PagerFromParam:                        types.StringPointerValue(httpInputOption.PagerFromParam),
		PagerToParam:                          types.StringPointerValue(httpInputOption.PagerToParam),
		PagerPages:                            types.Int64PointerValue(httpInputOption.PagerPages),
		PagerStart:                            types.Int64PointerValue(httpInputOption.PagerStart),
		PagerStep:                             types.Int64PointerValue(httpInputOption.PagerStep),
		CursorRequestParameterCursorName:      types.StringPointerValue(httpInputOption.CursorRequestParameterCursorName),
		CursorResponseParameterCursorJsonPath: types.StringPointerValue(httpInputOption.CursorResponseParameterCursorJsonPath),
		CursorRequestParameterLimitName:       types.StringPointerValue(httpInputOption.CursorRequestParameterLimitName),
		CursorRequestParameterLimitValue:      types.Int64PointerValue(httpInputOption.CursorRequestParameterLimitValue),
		RequestParams:                         NewRequestParams(httpInputOption.RequestParams, previousRequestParameters),
		RequestBody:                           types.StringPointerValue(httpInputOption.RequestBody),
		RequestHeaders:                        NewRequestHeaders(httpInputOption.RequestHeaders, previousRequestHeaders),
		SuccessCode:                           types.StringPointerValue(httpInputOption.SuccessCode),
		OpenTimeout:                           types.Int64PointerValue(httpInputOption.OpenTimeout),
		ReadTimeout:                           types.Int64PointerValue(httpInputOption.ReadTimeout),
		MaxRetries:                            types.Int64PointerValue(httpInputOption.MaxRetries),
		RetryInterval:                         types.Int64PointerValue(httpInputOption.RetryInterval),
		RequestInterval:                       types.Int64PointerValue(httpInputOption.RequestInterval),
		CsvParser:                             parser.NewCsvParser(httpInputOption.CsvParser),
		JsonlParser:                           parser.NewJsonlParser(httpInputOption.JsonlParser),
		JsonpathParser:                        parser.NewJsonPathParser(httpInputOption.JsonpathParser),
		LtsvParser:                            parser.NewLtsvParser(httpInputOption.LtsvParser),
		ExcelParser:                           parser.NewExcelParser(httpInputOption.ExcelParser),
		XmlParser:                             parser.NewXmlParser(httpInputOption.XmlParser),
		CustomVariableSettings:                model.NewCustomVariableSettings(httpInputOption.CustomVariableSettings),
	}
}

func NewRequestParams(params *[]entity.RequestParam, previous []RequestParam) []RequestParam {
	if params == nil || len(*params) == 0 {
		return nil
	}
	var ret []RequestParam
	for i, param := range *params {
		var previousRequestParameters RequestParam
		if len(previous) > i {
			previousRequestParameters = previous[i]
		}

		ret = append(ret, NewRequestParam(param, previousRequestParameters))
	}
	return ret
}

func NewRequestParam(param entity.RequestParam, previous RequestParam) RequestParam {
	value := types.StringValue(param.Value)
	if param.Masking != nil && *param.Masking && previous != (RequestParam{}) {
		value = previous.Value
	}

	return RequestParam{
		Key:     types.StringValue(param.Key),
		Value:   value,
		Masking: types.BoolPointerValue(param.Masking),
	}
}

func NewRequestHeaders(headers *[]entity.RequestHeader, previous []RequestHeader) []RequestHeader {
	if headers == nil || len(*headers) == 0 {
		return nil
	}
	var ret []RequestHeader
	for i, header := range *headers {
		var previousRequestHeader RequestHeader
		if len(previous) > i {
			previousRequestHeader = previous[i]
		}

		ret = append(ret, NewRequestHeader(header, previousRequestHeader))
	}
	return ret
}

func NewRequestHeader(header entity.RequestHeader, previous RequestHeader) RequestHeader {
	value := types.StringValue(header.Value)
	if header.Masking != nil && *header.Masking && previous != (RequestHeader{}) {
		value = previous.Value
	}

	return RequestHeader{
		Key:     types.StringValue(header.Key),
		Value:   value,
		Masking: types.BoolPointerValue(header.Masking),
	}
}

func (httpInputOption *HttpInputOption) ToInput() *parameter.HttpInputOptionInput {
	if httpInputOption == nil {
		return nil
	}

	var requestParams []parameter.RequestParamInput
	for _, param := range httpInputOption.RequestParams {
		requestParams = append(requestParams, parameter.RequestParamInput{
			Key:     param.Key.ValueString(),
			Value:   param.Value.ValueString(),
			Masking: param.Masking.ValueBoolPointer(),
		})
	}

	var requestHeaders []parameter.RequestHeaderInput
	for _, header := range httpInputOption.RequestHeaders {
		requestHeaders = append(requestHeaders, parameter.RequestHeaderInput{
			Key:     header.Key.ValueString(),
			Value:   header.Value.ValueString(),
			Masking: header.Masking.ValueBoolPointer(),
		})
	}

	return &parameter.HttpInputOptionInput{
		URL:                                   httpInputOption.URL.ValueString(),
		Method:                                httpInputOption.Method.ValueString(),
		UserAgent:                             httpInputOption.UserAgent.ValueStringPointer(),
		Charset:                               httpInputOption.Charset.ValueStringPointer(),
		PagerType:                             httpInputOption.PagerType.ValueStringPointer(),
		PagerFromParam:                        httpInputOption.PagerFromParam.ValueStringPointer(),
		PagerToParam:                          httpInputOption.PagerToParam.ValueStringPointer(),
		PagerPages:                            httpInputOption.PagerPages.ValueInt64Pointer(),
		PagerStart:                            httpInputOption.PagerStart.ValueInt64Pointer(),
		PagerStep:                             httpInputOption.PagerStep.ValueInt64Pointer(),
		CursorRequestParameterCursorName:      httpInputOption.CursorRequestParameterCursorName.ValueStringPointer(),
		CursorResponseParameterCursorJsonPath: httpInputOption.CursorResponseParameterCursorJsonPath.ValueStringPointer(),
		CursorRequestParameterLimitName:       httpInputOption.CursorRequestParameterLimitName.ValueStringPointer(),
		CursorRequestParameterLimitValue:      httpInputOption.CursorRequestParameterLimitValue.ValueInt64Pointer(),
		RequestParams:                         &requestParams,
		RequestBody:                           httpInputOption.RequestBody.ValueStringPointer(),
		RequestHeaders:                        &requestHeaders,
		SuccessCode:                           httpInputOption.SuccessCode.ValueStringPointer(),
		OpenTimeout:                           httpInputOption.OpenTimeout.ValueInt64Pointer(),
		ReadTimeout:                           httpInputOption.ReadTimeout.ValueInt64Pointer(),
		MaxRetries:                            httpInputOption.MaxRetries.ValueInt64Pointer(),
		RetryInterval:                         httpInputOption.RetryInterval.ValueInt64Pointer(),
		RequestInterval:                       httpInputOption.RequestInterval.ValueInt64Pointer(),
		CsvParser:                             httpInputOption.CsvParser.ToCsvParserInput(),
		JsonlParser:                           httpInputOption.JsonlParser.ToJsonlParserInput(),
		JsonpathParser:                        httpInputOption.JsonpathParser.ToJsonpathParserInput(),
		LtsvParser:                            httpInputOption.LtsvParser.ToLtsvParserInput(),
		ExcelParser:                           httpInputOption.ExcelParser.ToExcelParserInput(),
		XmlParser:                             httpInputOption.XmlParser.ToXmlParserInput(),
		CustomVariableSettings:                model.ToCustomVariableSettingInputs(httpInputOption.CustomVariableSettings),
	}
}

func (httpInputOption *HttpInputOption) ToUpdateInput() *parameter.UpdateHttpInputOptionInput {
	if httpInputOption == nil {
		return nil
	}

	var requestParams []parameter.RequestParamInput
	for _, param := range httpInputOption.RequestParams {
		requestParams = append(requestParams, parameter.RequestParamInput{
			Key:     param.Key.ValueString(),
			Value:   param.Value.ValueString(),
			Masking: param.Masking.ValueBoolPointer(),
		})
	}

	var requestHeaders []parameter.RequestHeaderInput
	for _, header := range httpInputOption.RequestHeaders {
		requestHeaders = append(requestHeaders, parameter.RequestHeaderInput{
			Key:     header.Key.ValueString(),
			Value:   header.Value.ValueString(),
			Masking: header.Masking.ValueBoolPointer(),
		})
	}

	return &parameter.UpdateHttpInputOptionInput{
		URL:                                   httpInputOption.URL.ValueStringPointer(),
		Method:                                httpInputOption.Method.ValueStringPointer(),
		UserAgent:                             httpInputOption.UserAgent.ValueStringPointer(),
		Charset:                               httpInputOption.Charset.ValueStringPointer(),
		PagerType:                             httpInputOption.PagerType.ValueStringPointer(),
		PagerFromParam:                        httpInputOption.PagerFromParam.ValueStringPointer(),
		PagerToParam:                          httpInputOption.PagerToParam.ValueStringPointer(),
		PagerPages:                            httpInputOption.PagerPages.ValueInt64Pointer(),
		PagerStart:                            httpInputOption.PagerStart.ValueInt64Pointer(),
		PagerStep:                             httpInputOption.PagerStep.ValueInt64Pointer(),
		CursorRequestParameterCursorName:      httpInputOption.CursorRequestParameterCursorName.ValueStringPointer(),
		CursorResponseParameterCursorJsonPath: httpInputOption.CursorResponseParameterCursorJsonPath.ValueStringPointer(),
		CursorRequestParameterLimitName:       httpInputOption.CursorRequestParameterLimitName.ValueStringPointer(),
		CursorRequestParameterLimitValue:      httpInputOption.CursorRequestParameterLimitValue.ValueInt64Pointer(),
		RequestParams:                         &requestParams,
		RequestBody:                           httpInputOption.RequestBody.ValueStringPointer(),
		RequestHeaders:                        &requestHeaders,
		SuccessCode:                           httpInputOption.SuccessCode.ValueStringPointer(),
		OpenTimeout:                           httpInputOption.OpenTimeout.ValueInt64Pointer(),
		ReadTimeout:                           httpInputOption.ReadTimeout.ValueInt64Pointer(),
		MaxRetries:                            httpInputOption.MaxRetries.ValueInt64Pointer(),
		RetryInterval:                         httpInputOption.RetryInterval.ValueInt64Pointer(),
		RequestInterval:                       httpInputOption.RequestInterval.ValueInt64Pointer(),
		CsvParser:                             httpInputOption.CsvParser.ToCsvParserInput(),
		JsonlParser:                           httpInputOption.JsonlParser.ToJsonlParserInput(),
		JsonpathParser:                        httpInputOption.JsonpathParser.ToJsonpathParserInput(),
		LtsvParser:                            httpInputOption.LtsvParser.ToLtsvParserInput(),
		ExcelParser:                           httpInputOption.ExcelParser.ToExcelParserInput(),
		XmlParser:                             httpInputOption.XmlParser.ToXmlParserInput(),
		CustomVariableSettings:                model.ToCustomVariableSettingInputs(httpInputOption.CustomVariableSettings),
	}
}
