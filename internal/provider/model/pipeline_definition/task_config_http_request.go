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
	Name              types.String `tfsdk:"name"`
	ConnectionID      types.Int64  `tfsdk:"connection_id"`
	Method            types.String `tfsdk:"http_method"`
	URL               types.String `tfsdk:"url"`
	RequestBody       types.String `tfsdk:"request_body"`
	RequestHeaders    types.Set    `tfsdk:"request_headers"`
	RequestParameters types.Set    `tfsdk:"request_parameters"`
	CustomVariables   types.Set    `tfsdk:"custom_variables"`
}

func NewHTTPRequestTaskConfig(en *we.HTTPRequestTaskConfig, previous *HTTPRequestTaskConfig) *HTTPRequestTaskConfig {
	if en == nil {
		return nil
	}

	var previousRequestHeaders types.Set
	var previousRequestParameters types.Set
	if previous != nil {
		previousRequestHeaders = previous.RequestHeaders
		previousRequestParameters = previous.RequestParameters
	} else {
		objectHeaderType := types.ObjectType{AttrTypes: HTTPRequestHeader{}.AttrTypes()}
		objectParamType := types.ObjectType{AttrTypes: HTTPRequestParameter{}.AttrTypes()}
		previousRequestHeaders = types.SetNull(objectHeaderType)
		previousRequestParameters = types.SetNull(objectParamType)
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
	// RequestHeaders は元の配列型
	requestHeaders := []wp.RequestHeader{}
	for _, e := range c.RequestHeaders {
		requestHeaders = append(requestHeaders, wp.RequestHeader{
			Key:     e.Key.ValueString(),
			Value:   e.Value.ValueString(),
			Masking: model.NewNullableBool(e.Masking),
		})
	}

	// RequestParameters は元の配列型
	requestParameters := []wp.RequestParameter{}
	for _, e := range c.RequestParameters {
		requestParameters = append(requestParameters, wp.RequestParameter{
			Key:     e.Key.ValueString(),
			Value:   e.Value.ValueString(),
			Masking: model.NewNullableBool(e.Masking),
		})
	}

	// CustomVariables は元の配列型
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

func NewHTTPRequestHeaders(ens []we.RequestHeader, previous types.Set) []*HTTPRequestHeader {
	if len(ens) == 0 {
		return nil
	}

	// For simplicity, we're not handling the previous values optimization with Sets
	// This could be improved to extract previous values from the Set if needed
	models := make([]*HTTPRequestHeader, len(ens))
	for i, en := range ens {
		models[i] = NewHTTPRequestHeader(en, nil)
	}

	return models
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

func NewHTTPRequestParameters(ens []we.RequestParameter, previous types.Set) []*HTTPRequestParameter {
	if len(ens) == 0 {
		return nil
	}

	// For simplicity, we're not handling the previous values optimization with Sets
	// This could be improved to extract previous values from the Set if needed
	models := make([]*HTTPRequestParameter, len(ens))
	for i, en := range ens {
		models[i] = NewHTTPRequestParameter(en, nil)
	}

	return models
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

func (HTTPRequestHeader) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key":     types.StringType,
		"value":   types.StringType,
		"masking": types.BoolType,
	}
}

func (HTTPRequestParameter) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"key":     types.StringType,
		"value":   types.StringType,
		"masking": types.BoolType,
	}
}
