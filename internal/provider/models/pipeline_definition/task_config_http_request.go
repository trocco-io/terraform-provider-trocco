package workflow

import (
	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
