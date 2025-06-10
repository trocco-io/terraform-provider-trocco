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
	if !c.RequestHeaders.IsNull() && !c.RequestHeaders.IsUnknown() {
		var headers []HTTPRequestHeader
		diags := c.RequestHeaders.ElementsAs(context.Background(), &headers, false)
		if !diags.HasError() {
			for _, e := range headers {
				requestHeaders = append(requestHeaders, wp.RequestHeader{
					Key:     e.Key.ValueString(),
					Value:   e.Value.ValueString(),
					Masking: model.NewNullableBool(e.Masking),
				})
			}
		}
	}

	requestParameters := []wp.RequestParameter{}
	if !c.RequestParameters.IsNull() && !c.RequestParameters.IsUnknown() {
		var parameters []HTTPRequestParameter
		diags := c.RequestParameters.ElementsAs(context.Background(), &parameters, false)
		if !diags.HasError() {
			for _, e := range parameters {
				requestParameters = append(requestParameters, wp.RequestParameter{
					Key:     e.Key.ValueString(),
					Value:   e.Value.ValueString(),
					Masking: model.NewNullableBool(e.Masking),
				})
			}
		}
	}

	customVariables := []wp.CustomVariable{}
	if !c.CustomVariables.IsNull() && !c.CustomVariables.IsUnknown() {
		var variables []CustomVariable
		diags := c.CustomVariables.ElementsAs(context.Background(), &variables, false)
		if diags.HasError() {
			return nil
		}
		for _, v := range variables {
			customVariables = append(customVariables, v.ToInput())
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
		CustomVariables:   customVariables,
	}
}

type HTTPRequestHeader struct {
	Key     types.String `tfsdk:"key"`
	Value   types.String `tfsdk:"value"`
	Masking types.Bool   `tfsdk:"masking"`
}

func NewHTTPRequestHeaders(ens []we.RequestHeader, previous types.List) types.List {
	if len(ens) == 0 {
		return types.ListNull(types.ObjectType{AttrTypes: HTTPRequestHeadersAttrTypes()})
	}

	var previousHeaders []HTTPRequestHeader
	if !previous.IsNull() && !previous.IsUnknown() {
		previous.ElementsAs(context.Background(), &previousHeaders, false)
	}

	var elements []attr.Value
	for i, en := range ens {
		var previousHTTPRequestHeader *HTTPRequestHeader
		if len(previousHeaders) > i {
			header := previousHeaders[i]
			previousHTTPRequestHeader = &header
		}

		elements = append(elements, types.ObjectValueMust(
			map[string]attr.Type{
				"key":     types.StringType,
				"value":   types.StringType,
				"masking": types.BoolType,
			},
			map[string]attr.Value{
				"key":     types.StringValue(en.Key),
				"value":   NewHTTPRequestHeaderValue(en, previousHTTPRequestHeader),
				"masking": types.BoolValue(en.Masking),
			},
		))
	}

	return types.ListValueMust(types.ObjectType{AttrTypes: HTTPRequestHeadersAttrTypes()}, elements)
}

func NewHTTPRequestHeaderValue(en we.RequestHeader, previous *HTTPRequestHeader) types.String {
	value := types.StringValue(en.Value)
	if en.Masking && previous != nil {
		value = previous.Value
	}
	return value
}

type HTTPRequestParameter struct {
	Key     types.String `tfsdk:"key"`
	Value   types.String `tfsdk:"value"`
	Masking types.Bool   `tfsdk:"masking"`
}

func NewHTTPRequestParameters(ens []we.RequestParameter, previous types.List) types.List {
	if len(ens) == 0 {
		return types.ListNull(types.ObjectType{AttrTypes: map[string]attr.Type{
			"key":     types.StringType,
			"value":   types.StringType,
			"masking": types.BoolType,
		}})
	}

	var previousParameters []HTTPRequestParameter
	if !previous.IsNull() && !previous.IsUnknown() {
		previous.ElementsAs(context.Background(), &previousParameters, false)
	}

	var elements []attr.Value
	for i, en := range ens {
		var previousHTTPRequestParameter *HTTPRequestParameter
		if len(previousParameters) > i {
			param := previousParameters[i]
			previousHTTPRequestParameter = &param
		}

		elements = append(elements, types.ObjectValueMust(
			map[string]attr.Type{
				"key":     types.StringType,
				"value":   types.StringType,
				"masking": types.BoolType,
			},
			map[string]attr.Value{
				"key":     types.StringValue(en.Key),
				"value":   NewHTTPRequestParameterValue(en, previousHTTPRequestParameter),
				"masking": types.BoolValue(en.Masking),
			},
		))
	}

	return types.ListValueMust(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key":     types.StringType,
				"value":   types.StringType,
				"masking": types.BoolType,
			},
		},
		elements,
	)
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
