package pipeline_definition

import (
	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameters"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
	model "terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HTTPRequestTaskConfig struct {
	Name              types.String           `tfsdk:"name"`
	ConnectionID      types.Int64            `tfsdk:"connection_id"`
	Method            types.String           `tfsdk:"http_method"`
	URL               types.String           `tfsdk:"url"`
	RequestBody       types.String           `tfsdk:"request_body"`
	RequestHeaders    []HTTPRequestHeader    `tfsdk:"request_headers"`
	RequestParameters []HTTPRequestParameter `tfsdk:"request_parameters"`
	CustomVariables   []CustomVariable       `tfsdk:"custom_variables"`
}

func NewHTTPRequestTaskConfig(en *we.HTTPRequestTaskConfig, previous *HTTPRequestTaskConfig) *HTTPRequestTaskConfig {
	if en == nil {
		return nil
	}

	return &HTTPRequestTaskConfig{
		Name:              types.StringValue(en.Name),
		ConnectionID:      types.Int64PointerValue(en.ConnectionID),
		Method:            types.StringValue(en.HTTPMethod),
		URL:               types.StringValue(en.URL),
		RequestBody:       types.StringPointerValue(en.RequestBody),
		RequestHeaders:    NewHTTPRequestHeaders(en.RequestHeaders, previous.RequestHeaders),
		RequestParameters: NewHTTPRequestParameters(en.RequestParameters, previous.RequestParameters),
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

func NewHTTPRequestHeaders(ens []we.RequestHeader, previous []HTTPRequestHeader) []HTTPRequestHeader {
	if len(ens) == 0 {
		return nil
	}

	var mds []HTTPRequestHeader
	for i, en := range ens {
		mds = append(mds, NewHTTPRequestHeader(en, previous[i]))
	}

	return mds
}

func NewHTTPRequestHeader(en we.RequestHeader, previous HTTPRequestHeader) HTTPRequestHeader {
	value := types.StringValue(en.Value)
	if en.Masking {
		value = previous.Value
	}

	return HTTPRequestHeader{
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

func NewHTTPRequestParameters(ens []we.RequestParameter, previous []HTTPRequestParameter) []HTTPRequestParameter {
	if len(ens) == 0 {
		return nil
	}

	var mds []HTTPRequestParameter
	for i, en := range ens {
		mds = append(mds, NewHTTPRequestParameter(en, previous[i]))
	}

	return mds
}

func NewHTTPRequestParameter(en we.RequestParameter, previous HTTPRequestParameter) HTTPRequestParameter {
	value := types.StringValue(en.Value)
	if en.Masking {
		value = previous.Value
	}

	return HTTPRequestParameter{
		Key:     types.StringValue(en.Key),
		Value:   value,
		Masking: types.BoolValue(en.Masking),
	}
}
