package model

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity"
	"terraform-provider-trocco/internal/provider/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomConnectorOutputModel struct {
	ID               types.Int64  `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	URL              types.String `tfsdk:"url"`
	AuthType         types.String `tfsdk:"auth_type"`
	AuthHeaderName   types.String `tfsdk:"auth_header_name"`
	AuthHeaderScheme types.String `tfsdk:"auth_header_scheme"`
	GrantType        types.String `tfsdk:"grant_type"`
	AuthURI          types.String `tfsdk:"auth_uri"`
	AccessTokenURI   types.String `tfsdk:"access_token_uri"`
	Endpoints        types.List   `tfsdk:"endpoints"`
	CreatedAt        types.String `tfsdk:"created_at"`
	UpdatedAt        types.String `tfsdk:"updated_at"`
}

// CustomConnectorOutputEndpointModel reuses CustomConnectorFieldOptionModel for
// `headers`/`query_parameters` and CustomConnectorPathParameterModel for
// `path_parameters`, which share their shape with the transfer source
// definition. There is no `paginator`.
type CustomConnectorOutputEndpointModel struct {
	ID                types.Int64  `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	RequestType       types.String `tfsdk:"request_type"`
	BatchSize         types.Int64  `tfsdk:"batch_size"`
	Method            types.String `tfsdk:"method"`
	Operation         types.String `tfsdk:"operation"`
	Path              types.String `tfsdk:"path"`
	PayloadType       types.String `tfsdk:"payload_type"`
	SuccessCodes      types.String `tfsdk:"success_codes"`
	NotRetryableCodes types.String `tfsdk:"not_retryable_codes"`
	RequestTimeoutSec types.Int64  `tfsdk:"request_timeout_sec"`
	Template          types.String `tfsdk:"template"`
	Headers           types.List   `tfsdk:"headers"`
	PathParameters    types.List   `tfsdk:"path_parameters"`
	QueryParameters   types.List   `tfsdk:"query_parameters"`
}

func CustomConnectorOutputEndpointAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                  types.Int64Type,
		"name":                types.StringType,
		"request_type":        types.StringType,
		"batch_size":          types.Int64Type,
		"method":              types.StringType,
		"operation":           types.StringType,
		"path":                types.StringType,
		"payload_type":        types.StringType,
		"success_codes":       types.StringType,
		"not_retryable_codes": types.StringType,
		"request_timeout_sec": types.Int64Type,
		"template":            types.StringType,
		"headers":             types.ListType{ElemType: types.ObjectType{AttrTypes: CustomConnectorFieldOptionAttrTypes()}},
		"path_parameters":     types.ListType{ElemType: types.ObjectType{AttrTypes: CustomConnectorPathParameterAttrTypes()}},
		"query_parameters":    types.ListType{ElemType: types.ObjectType{AttrTypes: CustomConnectorFieldOptionAttrTypes()}},
	}
}

// customConnectorOutputEndpointKey is the stable identity used to match
// endpoints across API responses, plan, and prior state. Endpoint names are
// required to be unique within a connector (enforced in ValidateConfig).
func customConnectorOutputEndpointKey(e CustomConnectorOutputEndpointModel) string {
	return e.Name.ValueString()
}

// NewCustomConnectorOutputModel hydrates the TF model from the API entity.
// refEndpoints is the endpoint list from the plan (Create/Update) or prior
// state (Read); it carries no authority over content, only the desired
// ordering/identity for utils.MatchByKey, since the API returns endpoints
// sorted by id rather than in the user's configured order.
func NewCustomConnectorOutputModel(ctx context.Context, def *entity.CustomConnectorOutput, refEndpoints []CustomConnectorOutputEndpointModel) (CustomConnectorOutputModel, error) {
	m := CustomConnectorOutputModel{
		ID:               types.Int64Value(def.ID),
		Name:             types.StringValue(def.Name),
		Description:      types.StringPointerValue(def.Description),
		URL:              types.StringValue(def.URL),
		AuthType:         types.StringValue(def.AuthType),
		AuthHeaderName:   types.StringValue(def.AuthHeaderName),
		AuthHeaderScheme: types.StringPointerValue(def.AuthHeaderScheme),
		GrantType:        types.StringPointerValue(def.GrantType),
		AuthURI:          types.StringPointerValue(def.AuthURI),
		AccessTokenURI:   types.StringPointerValue(def.AccessTokenURI),
		CreatedAt:        types.StringValue(def.CreatedAt),
		UpdatedAt:        types.StringValue(def.UpdatedAt),
	}

	endpoints := make([]CustomConnectorOutputEndpointModel, 0, len(def.Endpoints))
	for _, e := range def.Endpoints {
		em := CustomConnectorOutputEndpointModel{
			ID:                types.Int64Value(e.ID),
			Name:              types.StringValue(e.Name),
			RequestType:       types.StringValue(e.RequestType),
			BatchSize:         types.Int64Value(e.BatchSize),
			Method:            types.StringValue(e.Method),
			Operation:         types.StringValue(e.Operation),
			Path:              types.StringValue(e.Path),
			PayloadType:       types.StringValue(e.PayloadType),
			SuccessCodes:      types.StringValue(e.SuccessCodes),
			NotRetryableCodes: types.StringValue(e.NotRetryableCodes),
			RequestTimeoutSec: types.Int64Value(e.RequestTimeoutSec),
			Template:          types.StringValue(e.Template),
		}

		headers, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: CustomConnectorFieldOptionAttrTypes()}, newCustomConnectorFieldOptionModels(e.Headers))
		if diags.HasError() {
			return m, fmt.Errorf("failed to convert headers to ListValue: %v", diags)
		}
		em.Headers = headers

		pathParams, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: CustomConnectorPathParameterAttrTypes()}, newCustomConnectorPathParameterModels(e.PathParameters))
		if diags.HasError() {
			return m, fmt.Errorf("failed to convert path_parameters to ListValue: %v", diags)
		}
		em.PathParameters = pathParams

		queryParams, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: CustomConnectorFieldOptionAttrTypes()}, newCustomConnectorFieldOptionModels(e.QueryParameters))
		if diags.HasError() {
			return m, fmt.Errorf("failed to convert query_parameters to ListValue: %v", diags)
		}
		em.QueryParameters = queryParams

		endpoints = append(endpoints, em)
	}
	endpoints = utils.MatchByKey(endpoints, refEndpoints, customConnectorOutputEndpointKey, nil)

	endpointsList, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: CustomConnectorOutputEndpointAttrTypes()}, endpoints)
	if diags.HasError() {
		return m, fmt.Errorf("failed to convert endpoints to ListValue: %v", diags)
	}
	m.Endpoints = endpointsList

	return m, nil
}
