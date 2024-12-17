package workflow

import (
	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameters"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"

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

func NewHTTPRequestTaskConfig(c *we.HTTPRequestTaskConfig) *HTTPRequestTaskConfig {
	if c == nil {
		return nil
	}

	return &HTTPRequestTaskConfig{
		Name:              types.StringValue(c.Name),
		ConnectionID:      types.Int64PointerValue(c.ConnectionID),
		Method:            types.StringValue(c.HTTPMethod),
		URL:               types.StringValue(c.URL),
		RequestBody:       types.StringPointerValue(c.RequestBody),
		RequestHeaders:    NewHTTPRequestHeaders(c.RequestHeaders),
		RequestParameters: NewHTTPRequestParameters(c.RequestParameters),
		CustomVariables:   NewCustomVariables(c.CustomVariables),
	}
}

func (c *HTTPRequestTaskConfig) ToInput() *wp.HTTPRequestTaskConfig {
	requestHeaders := []wp.RequestHeader{}
	for _, e := range c.RequestHeaders {
		requestHeaders = append(requestHeaders, wp.RequestHeader{
			Key:     e.Key.ValueString(),
			Value:   e.Value.ValueString(),
			Masking: &p.NullableBool{Valid: !e.Masking.IsNull(), Value: e.Masking.ValueBool()},
		})
	}

	requestParameters := []wp.RequestParameter{}
	for _, e := range c.RequestParameters {
		requestParameters = append(requestParameters, wp.RequestParameter{
			Key:     e.Key.ValueString(),
			Value:   e.Value.ValueString(),
			Masking: &p.NullableBool{Valid: !e.Masking.IsNull(), Value: e.Masking.ValueBool()},
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

func NewHTTPRequestHeaders(ens []we.RequestHeader) []HTTPRequestHeader {
	if len(ens) == 0 {
		return nil
	}

	var mds []HTTPRequestHeader
	for _, en := range ens {
		mds = append(mds, NewHTTPRequestHeader(en))
	}

	return mds
}

func NewHTTPRequestHeader(en we.RequestHeader) HTTPRequestHeader {
	return HTTPRequestHeader{
		Key:     types.StringValue(en.Key),
		Value:   types.StringValue(en.Value),
		Masking: types.BoolValue(en.Masking),
	}
}

type HTTPRequestParameter struct {
	Key     types.String `tfsdk:"key"`
	Value   types.String `tfsdk:"value"`
	Masking types.Bool   `tfsdk:"masking"`
}

func NewHTTPRequestParameters(ens []we.RequestParameter) []HTTPRequestParameter {
	if len(ens) == 0 {
		return nil
	}

	var mds []HTTPRequestParameter
	for _, en := range ens {
		mds = append(mds, NewHTTPRequestParameter(en))
	}

	return mds
}

func NewHTTPRequestParameter(en we.RequestParameter) HTTPRequestParameter {
	return HTTPRequestParameter{
		Key:     types.StringValue(en.Key),
		Value:   types.StringValue(en.Value),
		Masking: types.BoolValue(en.Masking),
	}
}
