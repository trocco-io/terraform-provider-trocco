package model

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity"
	"terraform-provider-trocco/internal/provider/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomConnectorInputModel struct {
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

type CustomConnectorEndpointModel struct {
	ID                    types.Int64                    `tfsdk:"id"`
	Name                  types.String                   `tfsdk:"name"`
	Path                  types.String                   `tfsdk:"path"`
	Method                types.String                   `tfsdk:"method"`
	RequestBody           types.String                   `tfsdk:"request_body"`
	JsonpathRoot          types.String                   `tfsdk:"jsonpath_root"`
	SuccessCodes          types.String                   `tfsdk:"success_codes"`
	NotRetryableCodes     types.String                   `tfsdk:"not_retryable_codes"`
	RequestTimeoutSec     types.Int64                    `tfsdk:"request_timeout_sec"`
	QueryParameters       types.List                     `tfsdk:"query_parameters"`
	Headers               types.List                     `tfsdk:"headers"`
	PathParameters        types.List                     `tfsdk:"path_parameters"`
	RequestBodyParameters types.List                     `tfsdk:"request_body_parameters"`
	Paginator             *CustomConnectorPaginatorModel `tfsdk:"paginator"`
}

// CustomConnectorFieldOptionModel is the shared shape for an endpoint's
// `query_parameters`, `headers` and `request_body_parameters`.
type CustomConnectorFieldOptionModel struct {
	Name         types.String `tfsdk:"name"`
	DisplayName  types.String `tfsdk:"display_name"`
	DefaultValue types.String `tfsdk:"default_value"`
	IsEditable   types.Bool   `tfsdk:"is_editable"`
	IsRequired   types.Bool   `tfsdk:"is_required"`
}

type CustomConnectorPathParameterModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type CustomConnectorPaginatorModel struct {
	InjectInto              types.String                                 `tfsdk:"inject_into"`
	PageIncrementStrategy   *CustomConnectorPageIncrementStrategyModel   `tfsdk:"page_increment_strategy"`
	OffsetIncrementStrategy *CustomConnectorOffsetIncrementStrategyModel `tfsdk:"offset_increment_strategy"`
	CursorBasedStrategy     *CustomConnectorCursorBasedStrategyModel     `tfsdk:"cursor_based_strategy"`
	PageSizeOption          *CustomConnectorFieldNameOptionModel         `tfsdk:"page_size_option"`
	PageTokenOption         *CustomConnectorFieldNameOptionModel         `tfsdk:"page_token_option"`
}

type CustomConnectorPageIncrementStrategyModel struct {
	PageSize        types.Int64  `tfsdk:"page_size"`
	LastPageSize    types.String `tfsdk:"last_page_size"`
	StartFromPage   types.Int64  `tfsdk:"start_from_page"`
	StopOnPage      types.Int64  `tfsdk:"stop_on_page"`
	TotalPages      types.String `tfsdk:"total_pages"`
	MaxRequestCount types.Int64  `tfsdk:"max_request_count"`
}

type CustomConnectorOffsetIncrementStrategyModel struct {
	PageSize        types.Int64  `tfsdk:"page_size"`
	LastPageSize    types.String `tfsdk:"last_page_size"`
	StartFromOffset types.Int64  `tfsdk:"start_from_offset"`
	StopOnOffset    types.Int64  `tfsdk:"stop_on_offset"`
	TotalRecords    types.String `tfsdk:"total_records"`
	MaxRequestCount types.Int64  `tfsdk:"max_request_count"`
}

type CustomConnectorCursorBasedStrategyModel struct {
	PageSize    types.Int64  `tfsdk:"page_size"`
	CursorValue types.String `tfsdk:"cursor_value"`
}

type CustomConnectorFieldNameOptionModel struct {
	FieldName types.String `tfsdk:"field_name"`
}

func CustomConnectorFieldOptionAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":          types.StringType,
		"display_name":  types.StringType,
		"default_value": types.StringType,
		"is_editable":   types.BoolType,
		"is_required":   types.BoolType,
	}
}

func CustomConnectorPathParameterAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":  types.StringType,
		"value": types.StringType,
	}
}

func CustomConnectorFieldNameOptionAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"field_name": types.StringType,
	}
}

func CustomConnectorPageIncrementStrategyAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"page_size":         types.Int64Type,
		"last_page_size":    types.StringType,
		"start_from_page":   types.Int64Type,
		"stop_on_page":      types.Int64Type,
		"total_pages":       types.StringType,
		"max_request_count": types.Int64Type,
	}
}

func CustomConnectorOffsetIncrementStrategyAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"page_size":         types.Int64Type,
		"last_page_size":    types.StringType,
		"start_from_offset": types.Int64Type,
		"stop_on_offset":    types.Int64Type,
		"total_records":     types.StringType,
		"max_request_count": types.Int64Type,
	}
}

func CustomConnectorCursorBasedStrategyAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"page_size":    types.Int64Type,
		"cursor_value": types.StringType,
	}
}

func CustomConnectorPaginatorAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"inject_into":               types.StringType,
		"page_increment_strategy":   types.ObjectType{AttrTypes: CustomConnectorPageIncrementStrategyAttrTypes()},
		"offset_increment_strategy": types.ObjectType{AttrTypes: CustomConnectorOffsetIncrementStrategyAttrTypes()},
		"cursor_based_strategy":     types.ObjectType{AttrTypes: CustomConnectorCursorBasedStrategyAttrTypes()},
		"page_size_option":          types.ObjectType{AttrTypes: CustomConnectorFieldNameOptionAttrTypes()},
		"page_token_option":         types.ObjectType{AttrTypes: CustomConnectorFieldNameOptionAttrTypes()},
	}
}

func CustomConnectorEndpointAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                      types.Int64Type,
		"name":                    types.StringType,
		"path":                    types.StringType,
		"method":                  types.StringType,
		"request_body":            types.StringType,
		"jsonpath_root":           types.StringType,
		"success_codes":           types.StringType,
		"not_retryable_codes":     types.StringType,
		"request_timeout_sec":     types.Int64Type,
		"query_parameters":        types.ListType{ElemType: types.ObjectType{AttrTypes: CustomConnectorFieldOptionAttrTypes()}},
		"headers":                 types.ListType{ElemType: types.ObjectType{AttrTypes: CustomConnectorFieldOptionAttrTypes()}},
		"path_parameters":         types.ListType{ElemType: types.ObjectType{AttrTypes: CustomConnectorPathParameterAttrTypes()}},
		"request_body_parameters": types.ListType{ElemType: types.ObjectType{AttrTypes: CustomConnectorFieldOptionAttrTypes()}},
		"paginator":               types.ObjectType{AttrTypes: CustomConnectorPaginatorAttrTypes()},
	}
}

// customConnectorEndpointKey is the stable identity used to match endpoints
// across API responses, plan, and prior state. Endpoint names are required to
// be unique within a connector (enforced in ValidateConfig).
func customConnectorEndpointKey(e CustomConnectorEndpointModel) string {
	return e.Name.ValueString()
}

func newCustomConnectorFieldOptionModels(opts []entity.CustomConnectorFieldOption) []CustomConnectorFieldOptionModel {
	out := make([]CustomConnectorFieldOptionModel, 0, len(opts))
	for _, o := range opts {
		out = append(out, CustomConnectorFieldOptionModel{
			Name:         types.StringValue(o.Name),
			DisplayName:  types.StringValue(o.DisplayName),
			DefaultValue: types.StringPointerValue(o.DefaultValue),
			IsEditable:   types.BoolValue(o.IsEditable),
			IsRequired:   types.BoolValue(o.IsRequired),
		})
	}
	return out
}

func newCustomConnectorPathParameterModels(params []entity.CustomConnectorPathParameter) []CustomConnectorPathParameterModel {
	out := make([]CustomConnectorPathParameterModel, 0, len(params))
	for _, p := range params {
		out = append(out, CustomConnectorPathParameterModel{
			Name:  types.StringValue(p.Name),
			Value: types.StringValue(p.Value),
		})
	}
	return out
}

func newCustomConnectorPaginatorModel(p *entity.CustomConnectorPaginator) *CustomConnectorPaginatorModel {
	if p == nil {
		return nil
	}
	m := &CustomConnectorPaginatorModel{
		InjectInto: types.StringValue(p.InjectInto),
	}
	s := p.PaginationStrategy
	switch s.Type {
	case "page_increment":
		m.PageIncrementStrategy = &CustomConnectorPageIncrementStrategyModel{
			PageSize:        types.Int64PointerValue(s.PageSize),
			LastPageSize:    types.StringPointerValue(s.LastPageSize),
			StartFromPage:   types.Int64PointerValue(s.StartFromPage),
			StopOnPage:      types.Int64PointerValue(s.StopOnPage),
			TotalPages:      types.StringPointerValue(s.TotalPages),
			MaxRequestCount: types.Int64PointerValue(s.MaxRequestCount),
		}
	case "offset_increment":
		m.OffsetIncrementStrategy = &CustomConnectorOffsetIncrementStrategyModel{
			PageSize:        types.Int64PointerValue(s.PageSize),
			LastPageSize:    types.StringPointerValue(s.LastPageSize),
			StartFromOffset: types.Int64PointerValue(s.StartFromOffset),
			StopOnOffset:    types.Int64PointerValue(s.StopOnOffset),
			TotalRecords:    types.StringPointerValue(s.TotalRecords),
			MaxRequestCount: types.Int64PointerValue(s.MaxRequestCount),
		}
	case "cursor_based":
		m.CursorBasedStrategy = &CustomConnectorCursorBasedStrategyModel{
			PageSize:    types.Int64PointerValue(s.PageSize),
			CursorValue: types.StringPointerValue(s.CursorValue),
		}
	}
	if p.PageSizeOption != nil {
		m.PageSizeOption = &CustomConnectorFieldNameOptionModel{FieldName: types.StringValue(p.PageSizeOption.FieldName)}
	}
	if p.PageTokenOption != nil {
		m.PageTokenOption = &CustomConnectorFieldNameOptionModel{FieldName: types.StringValue(p.PageTokenOption.FieldName)}
	}
	return m
}

// NewCustomConnectorInputModel hydrates the TF model from the API entity.
// refEndpoints is the endpoint list from the plan (Create/Update) or prior
// state (Read); it carries no authority over content, only the desired
// ordering/identity for utils.MatchByKey, since the API returns endpoints
// sorted by id rather than in the user's configured order.
func NewCustomConnectorInputModel(ctx context.Context, def *entity.CustomConnectorInput, refEndpoints []CustomConnectorEndpointModel) (CustomConnectorInputModel, error) {
	m := CustomConnectorInputModel{
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

	endpoints := make([]CustomConnectorEndpointModel, 0, len(def.Endpoints))
	for _, e := range def.Endpoints {
		em := CustomConnectorEndpointModel{
			ID:                types.Int64Value(e.ID),
			Name:              types.StringValue(e.Name),
			Path:              types.StringValue(e.Path),
			Method:            types.StringValue(e.Method),
			RequestBody:       types.StringPointerValue(e.RequestBody),
			JsonpathRoot:      types.StringValue(e.JsonpathRoot),
			SuccessCodes:      types.StringValue(e.SuccessCodes),
			NotRetryableCodes: types.StringValue(e.NotRetryableCodes),
			RequestTimeoutSec: types.Int64Value(e.RequestTimeoutSec),
			Paginator:         newCustomConnectorPaginatorModel(e.Paginator),
		}

		queryParams, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: CustomConnectorFieldOptionAttrTypes()}, newCustomConnectorFieldOptionModels(e.QueryParameters))
		if diags.HasError() {
			return m, fmt.Errorf("failed to convert query_parameters to ListValue: %v", diags)
		}
		em.QueryParameters = queryParams

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

		requestBodyParams, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: CustomConnectorFieldOptionAttrTypes()}, newCustomConnectorFieldOptionModels(e.RequestBodyParameters))
		if diags.HasError() {
			return m, fmt.Errorf("failed to convert request_body_parameters to ListValue: %v", diags)
		}
		em.RequestBodyParameters = requestBodyParams

		endpoints = append(endpoints, em)
	}
	endpoints = utils.MatchByKey(endpoints, refEndpoints, customConnectorEndpointKey, nil)

	endpointsList, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: CustomConnectorEndpointAttrTypes()}, endpoints)
	if diags.HasError() {
		return m, fmt.Errorf("failed to convert endpoints to ListValue: %v", diags)
	}
	m.Endpoints = endpointsList

	return m, nil
}
