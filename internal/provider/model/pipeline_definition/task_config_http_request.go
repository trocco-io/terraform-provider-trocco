package pipeline_definition

import (
	"context"
	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameter"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
	model "terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HTTPRequestTaskConfig struct {
	Name              types.String            `tfsdk:"name"`
	ConnectionID      types.Int64             `tfsdk:"connection_id"`
	Method            types.String            `tfsdk:"http_method"`
	URL               types.String            `tfsdk:"url"`
	RequestBody       types.String            `tfsdk:"request_body"`
	RequestHeaders    []*HTTPRequestHeader    `tfsdk:"request_headers"`
	RequestParameters []*HTTPRequestParameter `tfsdk:"request_parameters"`
	CustomVariables   []CustomVariable        `tfsdk:"custom_variables"`
}

func NewHTTPRequestTaskConfig(ctx context.Context, en *we.HTTPRequestTaskConfig, previous *HTTPRequestTaskConfig) *HTTPRequestTaskConfig {
	if en == nil {
		return nil
	}

	var previousRequestHeaders []*HTTPRequestHeader
	var previousRequestParameters []*HTTPRequestParameter
	if previous != nil {
		previousRequestHeaders = previous.RequestHeaders
		previousRequestParameters = previous.RequestParameters
	}

	return &HTTPRequestTaskConfig{
		Name:              types.StringValue(en.Name),
		ConnectionID:      types.Int64PointerValue(en.ConnectionID),
		Method:            types.StringValue(en.HTTPMethod),
		URL:               types.StringValue(en.URL),
		RequestBody:       types.StringPointerValue(en.RequestBody),
		RequestHeaders:    NewHTTPRequestHeaders(en.RequestHeaders, previousRequestHeaders),
		RequestParameters: NewHTTPRequestParameters(en.RequestParameters, previousRequestParameters),
		CustomVariables:   NewCustomVariables(en.CustomVariables),
	}
}

func (c *HTTPRequestTaskConfig) ToInput() *wp.HTTPRequestTaskConfig {
	requestHeaders := []wp.RequestHeader{}
	for _, e := range c.RequestHeaders {
		requestHeaders = append(requestHeaders, wp.RequestHeader{
			Key:     e.Key.ValueString(),
			Value:   e.Value.ValueString(),
			Masking: model.NewNullableBool(e.Masking),
		})
	}

	requestParameters := []wp.RequestParameter{}
	for _, e := range c.RequestParameters {
		requestParameters = append(requestParameters, wp.RequestParameter{
			Key:     e.Key.ValueString(),
			Value:   e.Value.ValueString(),
			Masking: model.NewNullableBool(e.Masking),
		})
	}

	customVariables := []wp.CustomVariable{}
	for _, v := range c.CustomVariables {
		customVariables = append(customVariables, v.ToInput())
	}

	return &wp.HTTPRequestTaskConfig{
		Name:              c.Name.ValueString(),
		ConnectionID:      &p.NullableInt64{Valid: !c.ConnectionID.IsNull(), Value: c.ConnectionID.ValueInt64()},
		HTTPMethod:        c.Method.ValueString(),
		URL:               c.URL.ValueString(),
		RequestBody:       c.RequestBody.ValueStringPointer(),
		RequestHeaders:    requestHeaders,
		RequestParameters: requestParameters,
		CustomVariables:   customVariables,
	}
}

type HTTPRequestHeader struct {
	Key     types.String `tfsdk:"key"`
	Value   types.String `tfsdk:"value"`
	Masking types.Bool   `tfsdk:"masking"`
}

func NewHTTPRequestHeaders(ens []we.RequestHeader, previous []*HTTPRequestHeader) []*HTTPRequestHeader {
	if len(ens) == 0 {
		return nil
	}

	var mds []*HTTPRequestHeader
	for i, en := range ens {
		var previousHTTPRequestHeader *HTTPRequestHeader
		if len(previous) > i {
			previousHTTPRequestHeader = previous[i]
		}

		mds = append(mds, NewHTTPRequestHeader(en, previousHTTPRequestHeader))
	}

	return mds
}

func NewHTTPRequestHeader(en we.RequestHeader, previous *HTTPRequestHeader) *HTTPRequestHeader {
	value := types.StringValue(en.Value)
	if en.Masking && previous != nil {
		value = previous.Value
	}

	return &HTTPRequestHeader{
		Key:     types.StringValue(en.Key),
		Value:   value,
		Masking: types.BoolValue(en.Masking),
	}
}

type HTTPRequestParameter struct {
	Key     types.String `tfsdk:"key"`
	Value   types.String `tfsdk:"value"`
	Masking types.Bool   `tfsdk:"masking"`
}

func NewHTTPRequestParameters(ens []we.RequestParameter, previous []*HTTPRequestParameter) []*HTTPRequestParameter {
	if len(ens) == 0 {
		return nil
	}

	var mds []*HTTPRequestParameter
	for i, en := range ens {
		var previousHTTPRequestParameter *HTTPRequestParameter
		if len(previous) > i {
			previousHTTPRequestParameter = previous[i]
		}

		mds = append(mds, NewHTTPRequestParameter(en, previousHTTPRequestParameter))
	}

	return mds
}

func NewHTTPRequestParameter(en we.RequestParameter, previous *HTTPRequestParameter) *HTTPRequestParameter {
	value := types.StringValue(en.Value)
	if en.Masking && previous != nil {
		value = previous.Value
	}

	return &HTTPRequestParameter{
		Key:     types.StringValue(en.Key),
		Value:   value,
		Masking: types.BoolValue(en.Masking),
	}
}

func NewHTTPRequestParameterValue(en we.RequestParameter, previous *HTTPRequestParameter) types.String {
	value := types.StringValue(en.Value)
	if en.Masking && previous != nil {
		value = previous.Value
	}
	return value
}

func HTTPRequestHeadersAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key":     types.StringType,
		"value":   types.StringType,
		"masking": types.BoolType,
	}
}

func HTTPRequestParametersAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key":     types.StringType,
		"value":   types.StringType,
		"masking": types.BoolType,
	}
}

func HTTPRequestTaskConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":               types.StringType,
		"connection_id":      types.Int64Type,
		"http_method":        types.StringType,
		"url":                types.StringType,
		"request_body":       types.StringType,
		"request_headers":    types.ListType{ElemType: types.ObjectType{AttrTypes: HTTPRequestHeadersAttrTypes()}},
		"request_parameters": types.ListType{ElemType: types.ObjectType{AttrTypes: HTTPRequestParametersAttrTypes()}},
		"custom_variables":   types.SetType{ElemType: types.ObjectType{AttrTypes: CustomVariableAttrTypes()}},
	}
}
