package output_options

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	output_options2 "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SalesforceOutputOption struct {
	Object                 types.String `tfsdk:"object"`
	ActionType             types.String `tfsdk:"action_type"`
	ApiVersion             types.String `tfsdk:"api_version"`
	UpsertKey              types.String `tfsdk:"upsert_key"`
	IgnoreNulls            types.Bool   `tfsdk:"ignore_nulls"`
	ThrowIfFailed          types.Bool   `tfsdk:"throw_if_failed"`
	SalesforceConnectionId types.Int64  `tfsdk:"salesforce_connection_id"`
}

func NewSalesforceOutputOption(salesforceOutputOption *output_option.SalesforceOutputOption) *SalesforceOutputOption {
	if salesforceOutputOption == nil {
		return nil
	}

	return &SalesforceOutputOption{
		Object:                 types.StringValue(salesforceOutputOption.Object),
		ActionType:             types.StringValue(salesforceOutputOption.ActionType),
		ApiVersion:             types.StringValue(salesforceOutputOption.ApiVersion),
		UpsertKey:              types.StringPointerValue(salesforceOutputOption.UpsertKey),
		IgnoreNulls:            types.BoolValue(salesforceOutputOption.IgnoreNulls),
		ThrowIfFailed:          types.BoolValue(salesforceOutputOption.ThrowIfFailed),
		SalesforceConnectionId: types.Int64Value(salesforceOutputOption.SalesforceConnectionId),
	}
}

func (salesforceOutputOption *SalesforceOutputOption) ToInput() *output_options2.SalesforceOutputOptionInput {
	if salesforceOutputOption == nil {
		return nil
	}

	return &output_options2.SalesforceOutputOptionInput{
		Object:                 salesforceOutputOption.Object.ValueString(),
		ActionType:             model.NewNullableString(salesforceOutputOption.ActionType),
		ApiVersion:             model.NewNullableString(salesforceOutputOption.ApiVersion),
		UpsertKey:              model.NewNullableString(salesforceOutputOption.UpsertKey),
		IgnoreNulls:            model.NewNullableBool(salesforceOutputOption.IgnoreNulls),
		ThrowIfFailed:          model.NewNullableBool(salesforceOutputOption.ThrowIfFailed),
		SalesforceConnectionId: salesforceOutputOption.SalesforceConnectionId.ValueInt64(),
	}
}

func (salesforceOutputOption *SalesforceOutputOption) ToUpdateInput() *output_options2.UpdateSalesforceOutputOptionInput {
	if salesforceOutputOption == nil {
		return nil
	}

	return &output_options2.UpdateSalesforceOutputOptionInput{
		Object:                 salesforceOutputOption.Object.ValueStringPointer(),
		ActionType:             model.NewNullableString(salesforceOutputOption.ActionType),
		ApiVersion:             model.NewNullableString(salesforceOutputOption.ApiVersion),
		UpsertKey:              model.NewNullableString(salesforceOutputOption.UpsertKey),
		IgnoreNulls:            model.NewNullableBool(salesforceOutputOption.IgnoreNulls),
		ThrowIfFailed:          model.NewNullableBool(salesforceOutputOption.ThrowIfFailed),
		SalesforceConnectionId: salesforceOutputOption.SalesforceConnectionId.ValueInt64Pointer(),
	}
}
