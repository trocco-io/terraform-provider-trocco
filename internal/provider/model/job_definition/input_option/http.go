package input_options

import (
	"context"
	entity "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	parameter "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"
	"terraform-provider-trocco/internal/provider/model/job_definition/input_option/parser"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	URL                                   types.String           `tfsdk:"url"`
	Method                                types.String           `tfsdk:"method"`
	UserAgent                             types.String           `tfsdk:"user_agent"`
	Charset                               types.String           `tfsdk:"charset"`
	PagerType                             types.String           `tfsdk:"pager_type"`
	PagerFromParam                        types.String           `tfsdk:"pager_from_param"`
	PagerToParam                          types.String           `tfsdk:"pager_to_param"`
	PagerPages                            types.Int64            `tfsdk:"pager_pages"`
	PagerStart                            types.Int64            `tfsdk:"pager_start"`
	PagerStep                             types.Int64            `tfsdk:"pager_step"`
	CursorRequestParameterCursorName      types.String           `tfsdk:"cursor_request_parameter_cursor_name"`
	CursorResponseParameterCursorJsonPath types.String           `tfsdk:"cursor_response_parameter_cursor_json_path"`
	CursorRequestParameterLimitName       types.String           `tfsdk:"cursor_request_parameter_limit_name"`
	CursorRequestParameterLimitValue      types.String           `tfsdk:"cursor_request_parameter_limit_value"`
	RequestParams                         types.Set              `tfsdk:"request_params"`
	RequestBody                           types.String           `tfsdk:"request_body"`
	RequestHeaders                        types.Set              `tfsdk:"request_headers"`
	SuccessCode                           types.String           `tfsdk:"success_code"`
	OpenTimeout                           types.Int64            `tfsdk:"open_timeout"`
	ReadTimeout                           types.Int64            `tfsdk:"read_timeout"`
	MaxRetries                            types.Int64            `tfsdk:"max_retries"`
	RetryInterval                         types.Int64            `tfsdk:"retry_interval"`
	RequestInterval                       types.Int64            `tfsdk:"request_interval"`
	CsvParser                             *parser.CsvParser      `tfsdk:"csv_parser"`
	JsonlParser                           *parser.JsonlParser    `tfsdk:"jsonl_parser"`
	JsonpathParser                        *parser.JsonpathParser `tfsdk:"jsonpath_parser"`
	LtsvParser                            *parser.LtsvParser     `tfsdk:"ltsv_parser"`
	ExcelParser                           *parser.ExcelParser    `tfsdk:"excel_parser"`
	XmlParser                             *parser.XmlParser      `tfsdk:"xml_parser"`
	ParquetParser                         *parser.ParquetParser  `tfsdk:"parquet_parser"`
	CustomVariableSettings                types.List             `tfsdk:"custom_variable_settings"`
}

func NewHttpInputOption(ctx context.Context, httpInputOption *entity.HttpInputOption, previous *HttpInputOption) (*HttpInputOption, diag.Diagnostics) {
	if httpInputOption == nil {
		return nil, nil
	}

	var diags diag.Diagnostics
	var previousRequestHeaders []RequestHeader
	var previousRequestParameters []RequestParam
	if previous != nil {
		diags.Append(previous.RequestHeaders.ElementsAs(ctx, &previousRequestHeaders, false)...)
		diags.Append(previous.RequestParams.ElementsAs(ctx, &previousRequestParameters, false)...)
		if diags.HasError() {
			return nil, diags
		}
	}

	requestParams, d := NewRequestParams(ctx, httpInputOption.RequestParams, previousRequestParameters)
	diags.Append(d...)
	requestHeaders, d := NewRequestHeaders(ctx, httpInputOption.RequestHeaders, previousRequestHeaders)
	diags.Append(d...)

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, httpInputOption.CustomVariableSettings)
	if err != nil {
		diags.AddError("Error converting custom variable settings", err.Error())
		return nil, diags
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
		CursorRequestParameterLimitValue:      types.StringPointerValue(httpInputOption.CursorRequestParameterLimitValue),
		RequestParams:                         requestParams,
		RequestBody:                           types.StringPointerValue(httpInputOption.RequestBody),
		RequestHeaders:                        requestHeaders,
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
		ParquetParser:                         nil, // ParquetParser is not supported in Http input
		CustomVariableSettings:                customVariableSettings,
	}, diags
}

func NewRequestParams(ctx context.Context, params *[]entity.RequestParam, previous []RequestParam) (types.Set, diag.Diagnostics) {
	if params == nil || len(*params) == 0 {
		return types.SetNull(headerParamType()), nil
	}
	var ret []RequestParam
	for i, param := range *params {
		var previousRequestParameters RequestParam
		if len(previous) > i {
			previousRequestParameters = previous[i]
		}

		ret = append(ret, NewRequestParam(param, previousRequestParameters))
	}

	setValue, diags := types.SetValueFrom(ctx, headerParamType(), ret)

	if diags.HasError() {
		return types.SetNull(headerParamType()), diags
	}
	return setValue, diags
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

func NewRequestHeaders(ctx context.Context, headers *[]entity.RequestHeader, previous []RequestHeader) (types.Set, diag.Diagnostics) {
	if headers == nil || len(*headers) == 0 {
		return types.SetNull(headerParamType()), nil
	}
	var ret []RequestHeader
	for i, header := range *headers {
		var previousRequestHeader RequestHeader
		if len(previous) > i {
			previousRequestHeader = previous[i]
		}

		ret = append(ret, NewRequestHeader(header, previousRequestHeader))
	}

	setValue, diags := types.SetValueFrom(ctx, headerParamType(), ret)

	if diags.HasError() {
		return types.SetNull(headerParamType()), diags
	}
	return setValue, diags
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

func headerParamType() types.ObjectType {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"key":     types.StringType,
			"value":   types.StringType,
			"masking": types.BoolType,
		},
	}
}

func (httpInputOption *HttpInputOption) ToInput(ctx context.Context) (*parameter.HttpInputOptionInput, diag.Diagnostics) {
	if httpInputOption == nil {
		return nil, nil
	}

	var diags diag.Diagnostics
	var httpInputParams []RequestParam
	diags.Append(httpInputOption.RequestParams.ElementsAs(ctx, &httpInputParams, false)...)
	if diags.HasError() {
		return nil, diags
	}
	var requestParams []parameter.RequestParamInput
	for _, param := range httpInputParams {
		requestParams = append(requestParams, parameter.RequestParamInput{
			Key:     param.Key.ValueString(),
			Value:   param.Value.ValueString(),
			Masking: param.Masking.ValueBoolPointer(),
		})
	}

	var httpInputHeaders []RequestHeader
	diags.Append(httpInputOption.RequestHeaders.ElementsAs(ctx, &httpInputHeaders, false)...)
	if diags.HasError() {
		return nil, diags
	}
	var requestHeaders []parameter.RequestHeaderInput
	for _, header := range httpInputHeaders {
		requestHeaders = append(requestHeaders, parameter.RequestHeaderInput{
			Key:     header.Key.ValueString(),
			Value:   header.Value.ValueString(),
			Masking: header.Masking.ValueBoolPointer(),
		})
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, httpInputOption.CustomVariableSettings)

	return &parameter.HttpInputOptionInput{
		URL:                                   httpInputOption.URL.ValueString(),
		Method:                                httpInputOption.Method.ValueString(),
		UserAgent:                             model.NewNullableString(httpInputOption.UserAgent),
		Charset:                               model.NewNullableString(httpInputOption.Charset),
		PagerType:                             model.NewNullableString(httpInputOption.PagerType),
		PagerFromParam:                        model.NewNullableString(httpInputOption.PagerFromParam),
		PagerToParam:                          model.NewNullableString(httpInputOption.PagerToParam),
		PagerPages:                            model.NewNullableInt64(httpInputOption.PagerPages),
		PagerStart:                            model.NewNullableInt64(httpInputOption.PagerStart),
		PagerStep:                             model.NewNullableInt64(httpInputOption.PagerStep),
		CursorRequestParameterCursorName:      model.NewNullableString(httpInputOption.CursorRequestParameterCursorName),
		CursorResponseParameterCursorJsonPath: model.NewNullableString(httpInputOption.CursorResponseParameterCursorJsonPath),
		CursorRequestParameterLimitName:       model.NewNullableString(httpInputOption.CursorRequestParameterLimitName),
		CursorRequestParameterLimitValue:      model.NewNullableString(httpInputOption.CursorRequestParameterLimitValue),
		RequestParams:                         &requestParams,
		RequestBody:                           model.NewNullableString(httpInputOption.RequestBody),
		RequestHeaders:                        &requestHeaders,
		SuccessCode:                           model.NewNullableString(httpInputOption.SuccessCode),
		OpenTimeout:                           model.NewNullableInt64(httpInputOption.OpenTimeout),
		ReadTimeout:                           model.NewNullableInt64(httpInputOption.ReadTimeout),
		MaxRetries:                            model.NewNullableInt64(httpInputOption.MaxRetries),
		RetryInterval:                         model.NewNullableInt64(httpInputOption.RetryInterval),
		RequestInterval:                       model.NewNullableInt64(httpInputOption.RequestInterval),
		CsvParser:                             httpInputOption.CsvParser.ToCsvParserInput(),
		JsonlParser:                           httpInputOption.JsonlParser.ToJsonlParserInput(),
		JsonpathParser:                        httpInputOption.JsonpathParser.ToJsonpathParserInput(),
		LtsvParser:                            httpInputOption.LtsvParser.ToLtsvParserInput(),
		ExcelParser:                           httpInputOption.ExcelParser.ToExcelParserInput(),
		XmlParser:                             httpInputOption.XmlParser.ToXmlParserInput(),
		CustomVariableSettings:                model.ToCustomVariableSettingInputs(customVarSettings),
	}, diags
}

func (httpInputOption *HttpInputOption) ToUpdateInput(ctx context.Context) (*parameter.UpdateHttpInputOptionInput, diag.Diagnostics) {
	if httpInputOption == nil {
		return nil, nil
	}

	var diags diag.Diagnostics
	var httpInputParams []RequestParam
	diags.Append(httpInputOption.RequestParams.ElementsAs(ctx, &httpInputParams, false)...)
	if diags.HasError() {
		return nil, diags
	}
	var requestParams []parameter.RequestParamInput
	for _, param := range httpInputParams {
		requestParams = append(requestParams, parameter.RequestParamInput{
			Key:     param.Key.ValueString(),
			Value:   param.Value.ValueString(),
			Masking: param.Masking.ValueBoolPointer(),
		})
	}

	var httpInputHeaders []RequestHeader
	diags.Append(httpInputOption.RequestHeaders.ElementsAs(ctx, &httpInputHeaders, false)...)
	if diags.HasError() {
		return nil, diags
	}
	var requestHeaders []parameter.RequestHeaderInput
	for _, header := range httpInputHeaders {
		requestHeaders = append(requestHeaders, parameter.RequestHeaderInput{
			Key:     header.Key.ValueString(),
			Value:   header.Value.ValueString(),
			Masking: header.Masking.ValueBoolPointer(),
		})
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, httpInputOption.CustomVariableSettings)

	return &parameter.UpdateHttpInputOptionInput{
		URL:                                   httpInputOption.URL.ValueStringPointer(),
		Method:                                httpInputOption.Method.ValueStringPointer(),
		UserAgent:                             model.NewNullableString(httpInputOption.UserAgent),
		Charset:                               model.NewNullableString(httpInputOption.Charset),
		PagerType:                             model.NewNullableString(httpInputOption.PagerType),
		PagerFromParam:                        model.NewNullableString(httpInputOption.PagerFromParam),
		PagerToParam:                          model.NewNullableString(httpInputOption.PagerToParam),
		PagerPages:                            model.NewNullableInt64(httpInputOption.PagerPages),
		PagerStart:                            model.NewNullableInt64(httpInputOption.PagerStart),
		PagerStep:                             model.NewNullableInt64(httpInputOption.PagerStep),
		CursorRequestParameterCursorName:      model.NewNullableString(httpInputOption.CursorRequestParameterCursorName),
		CursorResponseParameterCursorJsonPath: model.NewNullableString(httpInputOption.CursorResponseParameterCursorJsonPath),
		CursorRequestParameterLimitName:       model.NewNullableString(httpInputOption.CursorRequestParameterLimitName),
		CursorRequestParameterLimitValue:      model.NewNullableString(httpInputOption.CursorRequestParameterLimitValue),
		RequestParams:                         &requestParams,
		RequestBody:                           model.NewNullableString(httpInputOption.RequestBody),
		RequestHeaders:                        &requestHeaders,
		SuccessCode:                           model.NewNullableString(httpInputOption.SuccessCode),
		OpenTimeout:                           model.NewNullableInt64(httpInputOption.OpenTimeout),
		ReadTimeout:                           model.NewNullableInt64(httpInputOption.ReadTimeout),
		MaxRetries:                            model.NewNullableInt64(httpInputOption.MaxRetries),
		RetryInterval:                         model.NewNullableInt64(httpInputOption.RetryInterval),
		RequestInterval:                       model.NewNullableInt64(httpInputOption.RequestInterval),
		CsvParser:                             httpInputOption.CsvParser.ToCsvParserInput(),
		JsonlParser:                           httpInputOption.JsonlParser.ToJsonlParserInput(),
		JsonpathParser:                        httpInputOption.JsonpathParser.ToJsonpathParserInput(),
		LtsvParser:                            httpInputOption.LtsvParser.ToLtsvParserInput(),
		ExcelParser:                           httpInputOption.ExcelParser.ToExcelParserInput(),
		XmlParser:                             httpInputOption.XmlParser.ToXmlParserInput(),
		CustomVariableSettings:                model.ToCustomVariableSettingInputs(customVarSettings),
	}, diags
}
