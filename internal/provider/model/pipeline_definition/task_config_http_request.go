package pipeline_definition

import (
	"context"
	"fmt"
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
	RequestHeaders    types.List   `tfsdk:"request_headers"`
	RequestParameters types.List   `tfsdk:"request_parameters"`
	CustomVariables   types.Set    `tfsdk:"custom_variables"`
}

func NewHTTPRequestTaskConfig(en *we.HTTPRequestTaskConfig, previous *HTTPRequestTaskConfig) *HTTPRequestTaskConfig {
	if en == nil {
		return nil
	}

	var previousRequestHeaders types.List
	var previousRequestParameters types.List
	if previous != nil {
		previousRequestHeaders = previous.RequestHeaders
		previousRequestParameters = previous.RequestParameters
	}

	CustomVariables, err := NewCustomVariables(en.CustomVariables)
	if err != nil {
		return nil
	}

	requestHeaders, err := NewHTTPRequestHeaders(en.RequestHeaders, previousRequestHeaders)
	if err != nil {
		return nil
	}

	requestParameters, err := NewHTTPRequestParameters(en.RequestParameters, previousRequestParameters)
	if err != nil {
		return nil
	}

	return &HTTPRequestTaskConfig{
		Name:              types.StringValue(en.Name),
		ConnectionID:      types.Int64PointerValue(en.ConnectionID),
		Method:            types.StringValue(en.HTTPMethod),
		URL:               types.StringValue(en.URL),
		RequestBody:       types.StringPointerValue(en.RequestBody),
		RequestHeaders:    requestHeaders,
		RequestParameters: requestParameters,
		CustomVariables:   CustomVariables,
	}
}

func (c *HTTPRequestTaskConfig) ToInput() *wp.HTTPRequestTaskConfig {
	requestHeaders := []wp.RequestHeader{}
	if !c.RequestHeaders.IsNull() {
		var headers []HTTPRequestHeader
		diags := c.RequestHeaders.ElementsAs(context.Background(), &headers, false)
		if !diags.HasError() {
			for _, header := range headers {
				requestHeaders = append(requestHeaders, wp.RequestHeader{
					Key:     header.Key.ValueString(),
					Value:   header.Value.ValueString(),
					Masking: model.NewNullableBool(header.Masking),
				})
			}
		}
	}

	requestParameters := []wp.RequestParameter{}
	if !c.RequestParameters.IsNull() {
		var parameters []HTTPRequestParameter
		diags := c.RequestParameters.ElementsAs(context.Background(), &parameters, false)
		if !diags.HasError() {
			for _, parameter := range parameters {
				requestParameters = append(requestParameters, wp.RequestParameter{
					Key:     parameter.Key.ValueString(),
					Value:   parameter.Value.ValueString(),
					Masking: model.NewNullableBool(parameter.Masking),
				})
			}
		}
	}

	return &wp.HTTPRequestTaskConfig{
		Name:              c.Name.ValueString(),
		ConnectionID:      &p.NullableInt64{Valid: !c.ConnectionID.IsNull(), Value: c.ConnectionID.ValueInt64()},
		HTTPMethod:        c.Method.ValueString(),
		URL:               c.URL.ValueString(),
		RequestBody:       c.RequestBody.ValueStringPointer(),
		RequestHeaders:    requestHeaders,
		RequestParameters: requestParameters,
		CustomVariables:   CustomVariablesToInput(c.CustomVariables),
	}
}

type HTTPRequestHeader struct {
	Key     types.String `tfsdk:"key"`
	Value   types.String `tfsdk:"value"`
	Masking types.Bool   `tfsdk:"masking"`
}

func NewHTTPRequestHeaders(ens []we.RequestHeader, previous types.List) (types.List, error) {
	if len(ens) == 0 {
		return types.ListNull(types.ObjectType{AttrTypes: HTTPRequestHeadersAttrTypes()}), nil
	}

	var headers []*HTTPRequestHeader
	for _, en := range ens {
		headers = append(headers, NewHTTPRequestHeader(en, nil))
	}

	if !previous.IsNull() {
		var previousHeaders []HTTPRequestHeader
		diags := previous.ElementsAs(context.Background(), &previousHeaders, false)
		if !diags.HasError() {
			for i := range headers {
				if len(previousHeaders) > i {
					headers[i].Value = previousHeaders[i].Value
				}
			}
		}
	}

	ctx := context.Background()
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: HTTPRequestHeadersAttrTypes()}, headers)
	if diags.HasError() {
		return types.ListNull(types.ObjectType{AttrTypes: HTTPRequestHeadersAttrTypes()}), fmt.Errorf("failed to convert HTTPRequestHeader to ListValue: %v", diags)
	}

	return list, nil
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

func NewHTTPRequestParameters(ens []we.RequestParameter, previous types.List) (types.List, error) {
	if len(ens) == 0 {
		return types.ListNull(types.ObjectType{AttrTypes: HTTPRequestParametersAttrTypes()}), nil
	}

	var parameters []*HTTPRequestParameter
	for _, en := range ens {
		parameters = append(parameters, NewHTTPRequestParameter(en, nil))
	}

	if !previous.IsNull() {
		var previousParameters []HTTPRequestParameter
		diags := previous.ElementsAs(context.Background(), &previousParameters, false)
		if !diags.HasError() {
			for i := range parameters {
				if len(previousParameters) > i {
					parameters[i].Value = previousParameters[i].Value
				}
			}
		}
	}

	ctx := context.Background()
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: HTTPRequestParametersAttrTypes()}, parameters)
	if diags.HasError() {
		return types.ListNull(types.ObjectType{AttrTypes: HTTPRequestParametersAttrTypes()}), fmt.Errorf("failed to convert HTTPRequestParameter to ListValue: %v", diags)
	}

	return list, nil
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
