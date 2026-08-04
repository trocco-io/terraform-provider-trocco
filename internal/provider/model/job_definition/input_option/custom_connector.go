package input_options

import (
	"context"
	jobDefinitionEntities "terraform-provider-trocco/internal/client/entity/job_definition"
	inputOptionEntities "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	jobDefinitionParameters "terraform-provider-trocco/internal/client/parameter/job_definition"
	inputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"
	"terraform-provider-trocco/internal/provider/model/job_definition/input_option/parser"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomConnectorInputOption struct {
	CustomConnectorEndpointID   types.Int64                    `tfsdk:"custom_connector_endpoint_id"`
	CustomConnectorConnectionID types.Int64                    `tfsdk:"custom_connector_connection_id"`
	QueryParameters             types.List                     `tfsdk:"query_parameters"`
	Headers                     types.List                     `tfsdk:"headers"`
	PathParameters              types.List                     `tfsdk:"path_parameters"`
	RequestBodyParameters       types.List                     `tfsdk:"request_body_parameters"`
	JsonpathParser              *CustomConnectorJsonpathParser `tfsdk:"jsonpath_parser"`
	CustomVariableSettings      types.List                     `tfsdk:"custom_variable_settings"`
	URL                         types.String                   `tfsdk:"url"`
	AuthType                    types.String                   `tfsdk:"auth_type"`
	AuthHeaderName              types.String                   `tfsdk:"auth_header_name"`
	AuthHeaderScheme            types.String                   `tfsdk:"auth_header_scheme"`
	EndpointPath                types.String                   `tfsdk:"endpoint_path"`
	EndpointMethod              types.String                   `tfsdk:"endpoint_method"`
	EndpointRequestBody         types.String                   `tfsdk:"endpoint_request_body"`
	SuccessCodes                types.String                   `tfsdk:"success_codes"`
	NotRetryableCodes           types.String                   `tfsdk:"not_retryable_codes"`
	RequestTimeoutSec           types.Int64                    `tfsdk:"request_timeout_sec"`
	Paginator                   types.Object                   `tfsdk:"paginator"`
}

// CustomConnectorNamedValue is the shared shape for `query_parameters`,
// `headers` and `path_parameters` elements.
type CustomConnectorNamedValue struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

func customConnectorNamedValueAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":  types.StringType,
		"value": types.StringType,
	}
}

// CustomConnectorJsonpathParser mirrors the shared parser.JsonpathParser
// shape, except `root` is server-derived (Computed only): the API always
// overwrites it from the referenced endpoint's jsonpath_root and ignores
// whatever is sent in the request.
type CustomConnectorJsonpathParser struct {
	Root            types.String `tfsdk:"root"`
	DefaultTimeZone types.String `tfsdk:"default_time_zone"`
	Columns         types.List   `tfsdk:"columns"`
}

func customConnectorJsonpathParserColumnAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":      types.StringType,
		"type":      types.StringType,
		"time_zone": types.StringType,
		"format":    types.StringType,
	}
}

type CustomConnectorPaginator struct {
	StrategyType            types.String                            `tfsdk:"strategy_type"`
	InjectInto              types.String                            `tfsdk:"inject_into"`
	PageIncrementStrategy   *CustomConnectorPageIncrementStrategy   `tfsdk:"page_increment_strategy"`
	OffsetIncrementStrategy *CustomConnectorOffsetIncrementStrategy `tfsdk:"offset_increment_strategy"`
	CursorBasedStrategy     *CustomConnectorCursorBasedStrategy     `tfsdk:"cursor_based_strategy"`
	PageSizeOption          *CustomConnectorFieldNameOption         `tfsdk:"page_size_option"`
	PageTokenOption         *CustomConnectorFieldNameOption         `tfsdk:"page_token_option"`
}

type CustomConnectorPageIncrementStrategy struct {
	PageSize             types.Int64  `tfsdk:"page_size"`
	StartFromPage        types.Int64  `tfsdk:"start_from_page"`
	FirstPage            types.Int64  `tfsdk:"first_page"`
	InjectOnFirstRequest types.Bool   `tfsdk:"inject_on_first_request"`
	LastPageSize         types.String `tfsdk:"last_page_size"`
	TotalPages           types.String `tfsdk:"total_pages"`
	StopOnPage           types.Int64  `tfsdk:"stop_on_page"`
	MaxRequestCount      types.Int64  `tfsdk:"max_request_count"`
}

type CustomConnectorOffsetIncrementStrategy struct {
	PageSize             types.Int64  `tfsdk:"page_size"`
	StartFromOffset      types.Int64  `tfsdk:"start_from_offset"`
	FirstOffset          types.Int64  `tfsdk:"first_offset"`
	InjectOnFirstRequest types.Bool   `tfsdk:"inject_on_first_request"`
	LastPageSize         types.String `tfsdk:"last_page_size"`
	TotalRecords         types.String `tfsdk:"total_records"`
	StopOnOffset         types.Int64  `tfsdk:"stop_on_offset"`
	MaxRequestCount      types.Int64  `tfsdk:"max_request_count"`
}

type CustomConnectorCursorBasedStrategy struct {
	PageSize     types.Int64  `tfsdk:"page_size"`
	CursorValue  types.String `tfsdk:"cursor_value"`
	LastPageSize types.String `tfsdk:"last_page_size"`
	StopOnBlank  types.String `tfsdk:"stop_on_blank"`
}

type CustomConnectorFieldNameOption struct {
	FieldName types.String `tfsdk:"field_name"`
}

func customConnectorFieldNameOptionAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"field_name": types.StringType,
	}
}

func customConnectorPageIncrementStrategyAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"page_size":               types.Int64Type,
		"start_from_page":         types.Int64Type,
		"first_page":              types.Int64Type,
		"inject_on_first_request": types.BoolType,
		"last_page_size":          types.StringType,
		"total_pages":             types.StringType,
		"stop_on_page":            types.Int64Type,
		"max_request_count":       types.Int64Type,
	}
}

func customConnectorOffsetIncrementStrategyAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"page_size":               types.Int64Type,
		"start_from_offset":       types.Int64Type,
		"first_offset":            types.Int64Type,
		"inject_on_first_request": types.BoolType,
		"last_page_size":          types.StringType,
		"total_records":           types.StringType,
		"stop_on_offset":          types.Int64Type,
		"max_request_count":       types.Int64Type,
	}
}

func customConnectorCursorBasedStrategyAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"page_size":      types.Int64Type,
		"cursor_value":   types.StringType,
		"last_page_size": types.StringType,
		"stop_on_blank":  types.StringType,
	}
}

// customConnectorPaginatorAttrTypes must be used (via types.Object) rather
// than a plain struct for the `paginator` field: it's an entirely Computed
// nested attribute, so the framework may need to represent it as Unknown
// while decoding a Plan/Config that doesn't (and can't) configure it, which
// a raw Go struct pointer cannot represent.
func customConnectorPaginatorAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"strategy_type":             types.StringType,
		"inject_into":               types.StringType,
		"page_increment_strategy":   types.ObjectType{AttrTypes: customConnectorPageIncrementStrategyAttrTypes()},
		"offset_increment_strategy": types.ObjectType{AttrTypes: customConnectorOffsetIncrementStrategyAttrTypes()},
		"cursor_based_strategy":     types.ObjectType{AttrTypes: customConnectorCursorBasedStrategyAttrTypes()},
		"page_size_option":          types.ObjectType{AttrTypes: customConnectorFieldNameOptionAttrTypes()},
		"page_token_option":         types.ObjectType{AttrTypes: customConnectorFieldNameOptionAttrTypes()},
	}
}

func NewCustomConnectorInputOption(
	ctx context.Context,
	inputOption *inputOptionEntities.CustomConnectorInputOption,
) (*CustomConnectorInputOption, diag.Diagnostics) {
	if inputOption == nil {
		return nil, nil
	}

	var diags diag.Diagnostics

	queryParameters, d := newCustomConnectorNamedValueList(ctx, inputOption.QueryParameters)
	diags.Append(d...)
	headers, d := newCustomConnectorNamedValueList(ctx, inputOption.Headers)
	diags.Append(d...)
	pathParameters, d := newCustomConnectorNamedValueList(ctx, inputOption.PathParameters)
	diags.Append(d...)
	requestBodyParameters, d := newCustomConnectorNamedValueList(ctx, inputOption.RequestBodyParameters)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, inputOption.CustomVariableSettings)
	if err != nil {
		diags.AddError("Error converting custom variable settings", err.Error())
		return nil, diags
	}

	jsonpathParser, d := newCustomConnectorJsonpathParser(ctx, inputOption.JsonpathParser)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	paginator, d := newCustomConnectorPaginator(ctx, inputOption.Paginator)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	return &CustomConnectorInputOption{
		CustomConnectorEndpointID:   types.Int64PointerValue(inputOption.CustomConnectorEndpointID),
		CustomConnectorConnectionID: types.Int64PointerValue(inputOption.CustomConnectorConnectionID),
		QueryParameters:             queryParameters,
		Headers:                     headers,
		PathParameters:              pathParameters,
		RequestBodyParameters:       requestBodyParameters,
		JsonpathParser:              jsonpathParser,
		CustomVariableSettings:      customVariableSettings,
		URL:                         types.StringValue(inputOption.URL),
		AuthType:                    types.StringValue(inputOption.AuthType),
		AuthHeaderName:              types.StringPointerValue(inputOption.AuthHeaderName),
		AuthHeaderScheme:            types.StringPointerValue(inputOption.AuthHeaderScheme),
		EndpointPath:                types.StringValue(inputOption.EndpointPath),
		EndpointMethod:              types.StringValue(inputOption.EndpointMethod),
		EndpointRequestBody:         types.StringPointerValue(inputOption.EndpointRequestBody),
		SuccessCodes:                types.StringValue(inputOption.SuccessCodes),
		NotRetryableCodes:           types.StringValue(inputOption.NotRetryableCodes),
		RequestTimeoutSec:           types.Int64PointerValue(inputOption.RequestTimeoutSec),
		Paginator:                   paginator,
	}, diags
}

func newCustomConnectorNamedValueList(
	ctx context.Context,
	values []inputOptionEntities.CustomConnectorNamedValue,
) (types.List, diag.Diagnostics) {
	objectType := types.ObjectType{AttrTypes: customConnectorNamedValueAttrTypes()}

	if values == nil {
		return types.ListNull(objectType), nil
	}

	elements := make([]CustomConnectorNamedValue, 0, len(values))
	for _, v := range values {
		elements = append(elements, CustomConnectorNamedValue{
			Name:  types.StringValue(v.Name),
			Value: types.StringValue(v.Value),
		})
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, elements)
	if diags.HasError() {
		return types.ListNull(objectType), diags
	}
	return listValue, nil
}

func newCustomConnectorJsonpathParser(
	ctx context.Context,
	jsonpathParser *jobDefinitionEntities.JsonpathParser,
) (*CustomConnectorJsonpathParser, diag.Diagnostics) {
	if jsonpathParser == nil {
		return nil, nil
	}

	var diags diag.Diagnostics

	columnElements := make([]parser.JsonpathParserColumn, 0, len(jsonpathParser.Columns))
	for _, c := range jsonpathParser.Columns {
		columnElements = append(columnElements, parser.JsonpathParserColumn{
			Name:     types.StringValue(c.Name),
			Type:     types.StringValue(c.Type),
			TimeZone: types.StringPointerValue(c.TimeZone),
			Format:   types.StringPointerValue(c.Format),
		})
	}

	columns, d := types.ListValueFrom(
		ctx,
		types.ObjectType{AttrTypes: customConnectorJsonpathParserColumnAttrTypes()},
		columnElements,
	)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	return &CustomConnectorJsonpathParser{
		Root:            types.StringValue(jsonpathParser.Root),
		DefaultTimeZone: types.StringValue(jsonpathParser.DefaultTimeZone),
		Columns:         columns,
	}, diags
}

// toJsonpathParserInput converts the jsonpath_parser block to its wire
// shape. Root is always sent empty: the server overwrites it from the
// referenced endpoint's jsonpath_root and ignores whatever is submitted.
func (jsonpathParser *CustomConnectorJsonpathParser) toJsonpathParserInput(ctx context.Context) *jobDefinitionParameters.JsonpathParserInput {
	if jsonpathParser == nil {
		return nil
	}

	var columnValues []parser.JsonpathParserColumn
	diags := jsonpathParser.Columns.ElementsAs(ctx, &columnValues, false)
	if diags.HasError() {
		return nil
	}

	columns := make([]jobDefinitionParameters.JsonpathParserColumnInput, 0, len(columnValues))
	for _, c := range columnValues {
		columns = append(columns, jobDefinitionParameters.JsonpathParserColumnInput{
			Name:     c.Name.ValueString(),
			Type:     c.Type.ValueString(),
			TimeZone: c.TimeZone.ValueStringPointer(),
			Format:   c.Format.ValueStringPointer(),
		})
	}

	return &jobDefinitionParameters.JsonpathParserInput{
		Root:            "",
		DefaultTimeZone: jsonpathParser.DefaultTimeZone.ValueString(),
		Columns:         columns,
	}
}

func newCustomConnectorPaginator(
	ctx context.Context,
	paginator *inputOptionEntities.CustomConnectorPaginator,
) (types.Object, diag.Diagnostics) {
	attrTypes := customConnectorPaginatorAttrTypes()

	if paginator == nil {
		return types.ObjectNull(attrTypes), nil
	}

	value := CustomConnectorPaginator{
		StrategyType:    types.StringValue(paginator.PaginationStrategy.Type),
		InjectInto:      types.StringValue(paginator.InjectInto),
		PageSizeOption:  newCustomConnectorFieldNameOption(paginator.PageSizeOption),
		PageTokenOption: newCustomConnectorFieldNameOption(paginator.PageTokenOption),
	}

	s := paginator.PaginationStrategy
	switch s.Type {
	case "page_increment":
		value.PageIncrementStrategy = &CustomConnectorPageIncrementStrategy{
			PageSize:             types.Int64PointerValue(s.PageSize),
			StartFromPage:        types.Int64PointerValue(s.StartFromPage),
			FirstPage:            types.Int64PointerValue(s.FirstPage),
			InjectOnFirstRequest: types.BoolPointerValue(s.InjectOnFirstRequest),
			LastPageSize:         types.StringPointerValue(s.LastPageSize),
			TotalPages:           types.StringPointerValue(s.TotalPages),
			StopOnPage:           types.Int64PointerValue(s.StopOnPage),
			MaxRequestCount:      types.Int64PointerValue(s.MaxRequestCount),
		}
	case "offset_increment":
		value.OffsetIncrementStrategy = &CustomConnectorOffsetIncrementStrategy{
			PageSize:             types.Int64PointerValue(s.PageSize),
			StartFromOffset:      types.Int64PointerValue(s.StartFromOffset),
			FirstOffset:          types.Int64PointerValue(s.FirstOffset),
			InjectOnFirstRequest: types.BoolPointerValue(s.InjectOnFirstRequest),
			LastPageSize:         types.StringPointerValue(s.LastPageSize),
			TotalRecords:         types.StringPointerValue(s.TotalRecords),
			StopOnOffset:         types.Int64PointerValue(s.StopOnOffset),
			MaxRequestCount:      types.Int64PointerValue(s.MaxRequestCount),
		}
	case "cursor_based":
		value.CursorBasedStrategy = &CustomConnectorCursorBasedStrategy{
			PageSize:     types.Int64PointerValue(s.PageSize),
			CursorValue:  types.StringPointerValue(s.CursorValue),
			LastPageSize: types.StringPointerValue(s.LastPageSize),
			StopOnBlank:  types.StringPointerValue(s.StopOnBlank),
		}
	}

	objectValue, diags := types.ObjectValueFrom(ctx, attrTypes, value)
	if diags.HasError() {
		return types.ObjectNull(attrTypes), diags
	}
	return objectValue, diags
}

func newCustomConnectorFieldNameOption(o *inputOptionEntities.CustomConnectorFieldNameOption) *CustomConnectorFieldNameOption {
	if o == nil {
		return nil
	}
	return &CustomConnectorFieldNameOption{FieldName: types.StringValue(o.FieldName)}
}

func (inputOption *CustomConnectorInputOption) ToInput(ctx context.Context) (*inputOptionParameters.CustomConnectorInputOptionInput, diag.Diagnostics) {
	if inputOption == nil {
		return nil, nil
	}

	var diags diag.Diagnostics

	queryParameters, d := extractCustomConnectorNamedValues(ctx, inputOption.QueryParameters)
	diags.Append(d...)
	headers, d := extractCustomConnectorNamedValues(ctx, inputOption.Headers)
	diags.Append(d...)
	pathParameters, d := extractCustomConnectorNamedValues(ctx, inputOption.PathParameters)
	diags.Append(d...)
	requestBodyParameters, d := extractCustomConnectorNamedValues(ctx, inputOption.RequestBodyParameters)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	return &inputOptionParameters.CustomConnectorInputOptionInput{
		CustomConnectorEndpointID:   inputOption.CustomConnectorEndpointID.ValueInt64(),
		CustomConnectorConnectionID: inputOption.CustomConnectorConnectionID.ValueInt64(),
		QueryParameters:             queryParameters,
		Headers:                     headers,
		PathParameters:              pathParameters,
		RequestBodyParameters:       requestBodyParameters,
		JsonpathParser:              inputOption.JsonpathParser.toJsonpathParserInput(ctx),
		CustomVariableSettings:      model.ToCustomVariableSettingInputs(customVarSettings),
	}, diags
}

func (inputOption *CustomConnectorInputOption) ToUpdateInput(ctx context.Context) (*inputOptionParameters.UpdateCustomConnectorInputOptionInput, diag.Diagnostics) {
	if inputOption == nil {
		return nil, nil
	}

	var diags diag.Diagnostics

	queryParameters, d := extractCustomConnectorNamedValues(ctx, inputOption.QueryParameters)
	diags.Append(d...)
	headers, d := extractCustomConnectorNamedValues(ctx, inputOption.Headers)
	diags.Append(d...)
	pathParameters, d := extractCustomConnectorNamedValues(ctx, inputOption.PathParameters)
	diags.Append(d...)
	requestBodyParameters, d := extractCustomConnectorNamedValues(ctx, inputOption.RequestBodyParameters)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	return &inputOptionParameters.UpdateCustomConnectorInputOptionInput{
		CustomConnectorEndpointID:   model.NewNullableInt64(inputOption.CustomConnectorEndpointID),
		CustomConnectorConnectionID: model.NewNullableInt64(inputOption.CustomConnectorConnectionID),
		QueryParameters:             queryParameters,
		Headers:                     headers,
		PathParameters:              pathParameters,
		RequestBodyParameters:       requestBodyParameters,
		JsonpathParser:              inputOption.JsonpathParser.toJsonpathParserInput(ctx),
		CustomVariableSettings:      model.ToCustomVariableSettingInputs(customVarSettings),
	}, diags
}

// extractCustomConnectorNamedValues implements the API's null=keep /
// []=clear-all / values=full-replace update semantics for `query_parameters`,
// `headers` and `path_parameters`: a null (unconfigured) list yields a nil
// pointer, which is omitted from the JSON request entirely so the server
// preserves the existing values; a non-null list (including an explicitly
// empty one) always yields a non-nil pointer, so the key is sent and the
// server replaces the collection accordingly.
func extractCustomConnectorNamedValues(
	ctx context.Context,
	list types.List,
) (*[]inputOptionParameters.CustomConnectorNamedValueInput, diag.Diagnostics) {
	if list.IsNull() || list.IsUnknown() {
		return nil, nil
	}

	var values []CustomConnectorNamedValue
	diags := list.ElementsAs(ctx, &values, false)
	if diags.HasError() {
		return nil, diags
	}

	inputs := make([]inputOptionParameters.CustomConnectorNamedValueInput, 0, len(values))
	for _, v := range values {
		inputs = append(inputs, inputOptionParameters.CustomConnectorNamedValueInput{
			Name:  v.Name.ValueString(),
			Value: v.Value.ValueString(),
		})
	}
	return &inputs, nil
}
