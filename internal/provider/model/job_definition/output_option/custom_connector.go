package output_options

import (
	"context"
	outputOptionEntities "terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	outputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomConnectorOutputOption struct {
	CustomConnectorConnectionID           types.Int64                      `tfsdk:"custom_connector_connection_id"`
	Mode                                  types.String                     `tfsdk:"mode"`
	UpdateKey                             types.String                     `tfsdk:"update_key"`
	CreateCustomConnectorOutputEndpointID types.Int64                      `tfsdk:"create_custom_connector_output_endpoint_id"`
	UpdateCustomConnectorOutputEndpointID types.Int64                      `tfsdk:"update_custom_connector_output_endpoint_id"`
	CreateEndpointSettings                *CustomConnectorEndpointSettings `tfsdk:"create_endpoint_settings"`
	UpdateEndpointSettings                *CustomConnectorEndpointSettings `tfsdk:"update_endpoint_settings"`
	CustomVariableSettings                types.List                       `tfsdk:"custom_variable_settings"`
	URL                                   types.String                     `tfsdk:"url"`
	AuthType                              types.String                     `tfsdk:"auth_type"`
	AuthHeaderName                        types.String                     `tfsdk:"auth_header_name"`
	AuthHeaderScheme                      types.String                     `tfsdk:"auth_header_scheme"`
	Endpoints                             types.List                       `tfsdk:"endpoints"`
}

// CustomConnectorEndpointSettings holds the editable values of one endpoint.
// The API never echoes these back (the resulting values are exposed through
// the `endpoints` snapshot instead, merged with the definition's non-editable
// defaults), so they are configuration-only and are carried over from the
// prior plan/state in NewCustomConnectorOutputOption.
type CustomConnectorEndpointSettings struct {
	Headers         types.List `tfsdk:"headers"`
	PathParameters  types.List `tfsdk:"path_parameters"`
	QueryParameters types.List `tfsdk:"query_parameters"`
}

// CustomConnectorNamedValue is the shared shape for `headers`,
// `path_parameters` and `query_parameters` elements.
type CustomConnectorNamedValue struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

func CustomConnectorNamedValueAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":  types.StringType,
		"value": types.StringType,
	}
}

type CustomConnectorOutputOptionEndpoint struct {
	Operation         types.String `tfsdk:"operation"`
	RequestType       types.String `tfsdk:"request_type"`
	BatchSize         types.Int64  `tfsdk:"batch_size"`
	Method            types.String `tfsdk:"method"`
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

func CustomConnectorOutputOptionEndpointAttrTypes() map[string]attr.Type {
	namedValueList := types.ListType{ElemType: types.ObjectType{AttrTypes: CustomConnectorNamedValueAttrTypes()}}

	return map[string]attr.Type{
		"operation":           types.StringType,
		"request_type":        types.StringType,
		"batch_size":          types.Int64Type,
		"method":              types.StringType,
		"path":                types.StringType,
		"payload_type":        types.StringType,
		"success_codes":       types.StringType,
		"not_retryable_codes": types.StringType,
		"request_timeout_sec": types.Int64Type,
		"template":            types.StringType,
		"headers":             namedValueList,
		"path_parameters":     namedValueList,
		"query_parameters":    namedValueList,
	}
}

// NewCustomConnectorOutputOption hydrates the TF model from the API entity.
// previous is the plan (Create/Update) or the prior state (Read); it supplies
// the values of `create_endpoint_settings`/`update_endpoint_settings`, which
// the API accepts but never returns.
func NewCustomConnectorOutputOption(
	ctx context.Context,
	outputOption *outputOptionEntities.CustomConnectorOutputOption,
	previous *CustomConnectorOutputOption,
) (*CustomConnectorOutputOption, diag.Diagnostics) {
	if outputOption == nil {
		return nil, nil
	}

	var diags diag.Diagnostics

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, outputOption.CustomVariableSettings)
	if err != nil {
		diags.AddError("Error converting custom variable settings", err.Error())
		return nil, diags
	}

	endpoints, d := newCustomConnectorOutputOptionEndpoints(ctx, outputOption.Endpoints)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	result := &CustomConnectorOutputOption{
		CustomConnectorConnectionID:           types.Int64PointerValue(outputOption.CustomConnectorConnectionID),
		Mode:                                  types.StringValue(outputOption.Mode),
		UpdateKey:                             newCustomConnectorUpdateKey(outputOption.UpdateKey),
		CreateCustomConnectorOutputEndpointID: types.Int64PointerValue(outputOption.CreateCustomConnectorOutputEndpointID),
		UpdateCustomConnectorOutputEndpointID: types.Int64PointerValue(outputOption.UpdateCustomConnectorOutputEndpointID),
		CustomVariableSettings:                customVariableSettings,
		URL:                                   types.StringValue(outputOption.URL),
		AuthType:                              types.StringValue(outputOption.AuthType),
		AuthHeaderName:                        types.StringPointerValue(outputOption.AuthHeaderName),
		AuthHeaderScheme:                      types.StringPointerValue(outputOption.AuthHeaderScheme),
		Endpoints:                             endpoints,
	}

	if previous != nil {
		result.CreateEndpointSettings = previous.CreateEndpointSettings
		result.UpdateEndpointSettings = previous.UpdateEndpointSettings
	}

	return result, diags
}

// newCustomConnectorUpdateKey maps the server's empty string to null. The API
// reports an unset update_key as "" rather than null, including after it clears
// the value when mode is `insert`; keeping it as "" would conflict with an
// unset (null) configuration. Empty strings are rejected at the schema level,
// so no configuration can legitimately round-trip to "".
func newCustomConnectorUpdateKey(updateKey *string) types.String {
	if updateKey == nil || *updateKey == "" {
		return types.StringNull()
	}
	return types.StringValue(*updateKey)
}

func newCustomConnectorOutputOptionEndpoints(
	ctx context.Context,
	endpoints []outputOptionEntities.CustomConnectorOutputOptionEndpoint,
) (types.List, diag.Diagnostics) {
	objectType := types.ObjectType{AttrTypes: CustomConnectorOutputOptionEndpointAttrTypes()}

	if endpoints == nil {
		return types.ListNull(objectType), nil
	}

	var diags diag.Diagnostics

	elements := make([]CustomConnectorOutputOptionEndpoint, 0, len(endpoints))
	for _, e := range endpoints {
		headers, d := newCustomConnectorNamedValueList(ctx, e.Headers)
		diags.Append(d...)
		pathParameters, d := newCustomConnectorNamedValueList(ctx, e.PathParameters)
		diags.Append(d...)
		queryParameters, d := newCustomConnectorNamedValueList(ctx, e.QueryParameters)
		diags.Append(d...)
		if diags.HasError() {
			return types.ListNull(objectType), diags
		}

		elements = append(elements, CustomConnectorOutputOptionEndpoint{
			Operation:         types.StringValue(e.Operation),
			RequestType:       types.StringValue(e.RequestType),
			BatchSize:         types.Int64Value(e.BatchSize),
			Method:            types.StringValue(e.Method),
			Path:              types.StringValue(e.Path),
			PayloadType:       types.StringValue(e.PayloadType),
			SuccessCodes:      types.StringValue(e.SuccessCodes),
			NotRetryableCodes: types.StringValue(e.NotRetryableCodes),
			RequestTimeoutSec: types.Int64Value(e.RequestTimeoutSec),
			Template:          types.StringPointerValue(e.Template),
			Headers:           headers,
			PathParameters:    pathParameters,
			QueryParameters:   queryParameters,
		})
	}

	listValue, d := types.ListValueFrom(ctx, objectType, elements)
	diags.Append(d...)
	if diags.HasError() {
		return types.ListNull(objectType), diags
	}
	return listValue, diags
}

func newCustomConnectorNamedValueList(
	ctx context.Context,
	values []outputOptionEntities.CustomConnectorNamedValue,
) (types.List, diag.Diagnostics) {
	objectType := types.ObjectType{AttrTypes: CustomConnectorNamedValueAttrTypes()}

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

func (outputOption *CustomConnectorOutputOption) ToInput(ctx context.Context) (*outputOptionParameters.CustomConnectorOutputOptionInput, diag.Diagnostics) {
	if outputOption == nil {
		return nil, nil
	}

	var diags diag.Diagnostics

	createEndpointSettings, d := outputOption.CreateEndpointSettings.toInput(ctx)
	diags.Append(d...)
	updateEndpointSettings, d := outputOption.UpdateEndpointSettings.toInput(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, outputOption.CustomVariableSettings)

	return &outputOptionParameters.CustomConnectorOutputOptionInput{
		CustomConnectorConnectionID:           outputOption.CustomConnectorConnectionID.ValueInt64(),
		Mode:                                  outputOption.Mode.ValueString(),
		UpdateKey:                             outputOption.UpdateKey.ValueStringPointer(),
		CreateCustomConnectorOutputEndpointID: outputOption.CreateCustomConnectorOutputEndpointID.ValueInt64(),
		UpdateCustomConnectorOutputEndpointID: outputOption.UpdateCustomConnectorOutputEndpointID.ValueInt64Pointer(),
		CreateEndpointSettings:                createEndpointSettings,
		UpdateEndpointSettings:                updateEndpointSettings,
		CustomVariableSettings:                model.ToCustomVariableSettingInputs(customVarSettings),
	}, diags
}

func (outputOption *CustomConnectorOutputOption) ToUpdateInput(ctx context.Context) (*outputOptionParameters.UpdateCustomConnectorOutputOptionInput, diag.Diagnostics) {
	if outputOption == nil {
		return nil, nil
	}

	var diags diag.Diagnostics

	createEndpointSettings, d := outputOption.CreateEndpointSettings.toInput(ctx)
	diags.Append(d...)
	updateEndpointSettings, d := outputOption.UpdateEndpointSettings.toInput(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, outputOption.CustomVariableSettings)

	return &outputOptionParameters.UpdateCustomConnectorOutputOptionInput{
		CustomConnectorConnectionID:           outputOption.CustomConnectorConnectionID.ValueInt64Pointer(),
		Mode:                                  outputOption.Mode.ValueStringPointer(),
		UpdateKey:                             outputOption.UpdateKey.ValueStringPointer(),
		CreateCustomConnectorOutputEndpointID: outputOption.CreateCustomConnectorOutputEndpointID.ValueInt64Pointer(),
		UpdateCustomConnectorOutputEndpointID: outputOption.UpdateCustomConnectorOutputEndpointID.ValueInt64Pointer(),
		CreateEndpointSettings:                createEndpointSettings,
		UpdateEndpointSettings:                updateEndpointSettings,
		CustomVariableSettings:                model.ToCustomVariableSettingInputs(customVarSettings),
	}, diags
}

func (settings *CustomConnectorEndpointSettings) toInput(ctx context.Context) (*outputOptionParameters.CustomConnectorEndpointSettingsInput, diag.Diagnostics) {
	if settings == nil {
		return nil, nil
	}

	var diags diag.Diagnostics

	headers, d := extractCustomConnectorNamedValues(ctx, settings.Headers)
	diags.Append(d...)
	pathParameters, d := extractCustomConnectorNamedValues(ctx, settings.PathParameters)
	diags.Append(d...)
	queryParameters, d := extractCustomConnectorNamedValues(ctx, settings.QueryParameters)
	diags.Append(d...)
	if diags.HasError() {
		return nil, diags
	}

	return &outputOptionParameters.CustomConnectorEndpointSettingsInput{
		Headers:         headers,
		PathParameters:  pathParameters,
		QueryParameters: queryParameters,
	}, diags
}

// extractCustomConnectorNamedValues implements the API's null=keep /
// []=clear-all / values=full-replace update semantics: a null (unconfigured)
// list yields a nil pointer, which is omitted from the JSON request entirely so
// the server preserves the existing values; a non-null list (including an
// explicitly empty one) always yields a non-nil pointer, so the key is sent and
// the server replaces the collection accordingly.
func extractCustomConnectorNamedValues(
	ctx context.Context,
	list types.List,
) (*[]outputOptionParameters.CustomConnectorNamedValueInput, diag.Diagnostics) {
	if list.IsNull() || list.IsUnknown() {
		return nil, nil
	}

	var values []CustomConnectorNamedValue
	diags := list.ElementsAs(ctx, &values, false)
	if diags.HasError() {
		return nil, diags
	}

	inputs := make([]outputOptionParameters.CustomConnectorNamedValueInput, 0, len(values))
	for _, v := range values {
		inputs = append(inputs, outputOptionParameters.CustomConnectorNamedValueInput{
			Name:  v.Name.ValueString(),
			Value: v.Value.ValueString(),
		})
	}
	return &inputs, nil
}
