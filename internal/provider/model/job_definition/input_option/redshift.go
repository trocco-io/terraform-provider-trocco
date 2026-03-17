package input_options

import (
	"context"
	inputOptionEntities "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	inputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RedshiftInputOption struct {
	RedshiftConnectionID   types.Int64  `tfsdk:"redshift_connection_id"`
	Database               types.String `tfsdk:"database"`
	Schema                 types.String `tfsdk:"schema"`
	Query                  types.String `tfsdk:"query"`
	FetchRows              types.Int64  `tfsdk:"fetch_rows"`
	ConnectTimeout         types.Int64  `tfsdk:"connect_timeout"`
	SocketTimeout          types.Int64  `tfsdk:"socket_timeout"`
	CustomVariableSettings types.List   `tfsdk:"custom_variable_settings"`
}

func NewRedshiftInputOption(ctx context.Context, redshiftInputOption *inputOptionEntities.RedshiftInputOption) *RedshiftInputOption {
	if redshiftInputOption == nil {
		return nil
	}

	result := &RedshiftInputOption{
		RedshiftConnectionID: types.Int64Value(redshiftInputOption.RedshiftConnectionID),
		Database:             types.StringValue(redshiftInputOption.Database),
		Schema:               types.StringValue(redshiftInputOption.Schema),
		Query:                types.StringValue(redshiftInputOption.Query),
		FetchRows:            types.Int64Value(redshiftInputOption.FetchRows),
		ConnectTimeout:       types.Int64Value(redshiftInputOption.ConnectTimeout),
		SocketTimeout:        types.Int64Value(redshiftInputOption.SocketTimeout),
	}

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, redshiftInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func (o *RedshiftInputOption) ToInput(ctx context.Context) *inputOptionParameters.RedshiftInputOptionInput {
	if o == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, o.CustomVariableSettings)

	return &inputOptionParameters.RedshiftInputOptionInput{
		RedshiftConnectionID:   o.RedshiftConnectionID.ValueInt64(),
		Database:               o.Database.ValueString(),
		Schema:                 model.NewNullableString(o.Schema),
		Query:                  o.Query.ValueString(),
		FetchRows:              o.FetchRows.ValueInt64(),
		ConnectTimeout:         o.ConnectTimeout.ValueInt64(),
		SocketTimeout:          o.SocketTimeout.ValueInt64(),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (o *RedshiftInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateRedshiftInputOptionInput {
	if o == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, o.CustomVariableSettings)

	return &inputOptionParameters.UpdateRedshiftInputOptionInput{
		RedshiftConnectionID:   o.RedshiftConnectionID.ValueInt64Pointer(),
		Database:               o.Database.ValueStringPointer(),
		Schema:                 model.NewNullableString(o.Schema),
		Query:                  o.Query.ValueStringPointer(),
		FetchRows:              o.FetchRows.ValueInt64Pointer(),
		ConnectTimeout:         o.ConnectTimeout.ValueInt64Pointer(),
		SocketTimeout:          o.SocketTimeout.ValueInt64Pointer(),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
	}
}
