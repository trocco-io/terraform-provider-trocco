package input_options

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	param "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PostgreSQLInputOption struct {
	PostgreSQLConnectionID             types.Int64                           `tfsdk:"postgresql_connection_id"`
	Database                           types.String                          `tfsdk:"database"`
	Schema                             types.String                          `tfsdk:"schema"`
	Query                              types.String                          `tfsdk:"query"`
	IncrementalLoadingEnabled          types.Bool                            `tfsdk:"incremental_loading_enabled"`
	Table                              types.String                          `tfsdk:"table"`
	IncrementalColumns                 types.String                          `tfsdk:"incremental_columns"`
	LastRecord                         types.String                          `tfsdk:"last_record"`
	FetchRows                          types.Int64                           `tfsdk:"fetch_rows"`
	ConnectTimeout                     types.Int64                           `tfsdk:"connect_timeout"`
	SocketTimeout                      types.Int64                           `tfsdk:"socket_timeout"`
	DefaultTimeZone                    types.String                          `tfsdk:"default_time_zone"`
	PostgreSQLInputOptionColumnOptions *[]PostgreSQLInputOptionColumnOptions `tfsdk:"postgresql_input_option_column_options"`
	CustomVariableSettings             *[]model.CustomVariableSetting        `tfsdk:"custom_variable_settings"`
}

type PostgreSQLInputOptionColumnOptions struct {
	ColumnName      types.String `tfsdk:"column_name"`
	ColumnValueType types.String `tfsdk:"column_value_type"`
}

func NewPostgreSQLInputOption(postgresqlInputOption *input_option.PostgreSQLInputOption) *PostgreSQLInputOption {
	if postgresqlInputOption == nil {
		return nil
	}

	return &PostgreSQLInputOption{
		Database:                           types.StringValue(postgresqlInputOption.Database),
		Schema:                             types.StringValue(postgresqlInputOption.Schema),
		Table:                              types.StringPointerValue(postgresqlInputOption.Table),
		Query:                              types.StringPointerValue(postgresqlInputOption.Query),
		IncrementalColumns:                 types.StringPointerValue(postgresqlInputOption.IncrementalColumns),
		LastRecord:                         types.StringPointerValue(postgresqlInputOption.LastRecord),
		IncrementalLoadingEnabled:          types.BoolValue(postgresqlInputOption.IncrementalLoadingEnabled),
		FetchRows:                          types.Int64Value(postgresqlInputOption.FetchRows),
		ConnectTimeout:                     types.Int64Value(postgresqlInputOption.ConnectTimeout),
		SocketTimeout:                      types.Int64Value(postgresqlInputOption.SocketTimeout),
		DefaultTimeZone:                    types.StringValue(postgresqlInputOption.DefaultTimeZone),
		PostgreSQLConnectionID:             types.Int64Value(postgresqlInputOption.PostgreSQLConnectionID),
		PostgreSQLInputOptionColumnOptions: newPostgreSQLInputOptionColumnOptions(postgresqlInputOption.PostgreSQLInputOptionColumnOptions),
		CustomVariableSettings:             model.NewCustomVariableSettings(postgresqlInputOption.CustomVariableSettings),
	}
}

func newPostgreSQLInputOptionColumnOptions(postgreSQLInputOptionColumnOptions *[]input_option.PostgreSQLInputOptionColumnOptions) *[]PostgreSQLInputOptionColumnOptions {
	if postgreSQLInputOptionColumnOptions == nil {
		return nil
	}
	columns := make([]PostgreSQLInputOptionColumnOptions, 0, len(*postgreSQLInputOptionColumnOptions))
	for _, input := range *postgreSQLInputOptionColumnOptions {
		column := PostgreSQLInputOptionColumnOptions{
			ColumnName:      types.StringValue(input.ColumnName),
			ColumnValueType: types.StringValue(input.ColumnValueType),
		}
		columns = append(columns, column)
	}
	return &columns
}

func (postgresqlInputOption *PostgreSQLInputOption) ToInput() *param.PostgreSQLInputOptionInput {
	if postgresqlInputOption == nil {
		return nil
	}

	return &param.PostgreSQLInputOptionInput{
		Database:                           postgresqlInputOption.Database.ValueString(),
		Schema:                             postgresqlInputOption.Schema.ValueString(),
		Table:                              model.NewNullableString(postgresqlInputOption.Table),
		Query:                              model.NewNullableString(postgresqlInputOption.Query),
		IncrementalColumns:                 model.NewNullableString(postgresqlInputOption.IncrementalColumns),
		LastRecord:                         model.NewNullableString(postgresqlInputOption.LastRecord),
		IncrementalLoadingEnabled:          postgresqlInputOption.IncrementalLoadingEnabled.ValueBool(),
		FetchRows:                          postgresqlInputOption.FetchRows.ValueInt64(),
		ConnectTimeout:                     postgresqlInputOption.ConnectTimeout.ValueInt64(),
		SocketTimeout:                      postgresqlInputOption.SocketTimeout.ValueInt64(),
		DefaultTimeZone:                    postgresqlInputOption.DefaultTimeZone.ValueString(),
		PostgreSQLConnectionID:             postgresqlInputOption.PostgreSQLConnectionID.ValueInt64(),
		PostgreSQLInputOptionColumnOptions: toPostgreSQLInputOptionColumnOptions(postgresqlInputOption.PostgreSQLInputOptionColumnOptions),
		CustomVariableSettings:             model.ToCustomVariableSettingInputs(postgresqlInputOption.CustomVariableSettings),
	}
}

func (postgresqlInputOption *PostgreSQLInputOption) ToUpdateInput() *param.UpdatePostgreSQLInputOptionInput {
	if postgresqlInputOption == nil {
		return nil
	}
	return &param.UpdatePostgreSQLInputOptionInput{
		Database:                           postgresqlInputOption.Database.ValueStringPointer(),
		Schema:                             postgresqlInputOption.Schema.ValueStringPointer(),
		Table:                              model.NewNullableString(postgresqlInputOption.Table),
		Query:                              model.NewNullableString(postgresqlInputOption.Query),
		IncrementalColumns:                 model.NewNullableString(postgresqlInputOption.IncrementalColumns),
		LastRecord:                         model.NewNullableString(postgresqlInputOption.LastRecord),
		IncrementalLoadingEnabled:          postgresqlInputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		FetchRows:                          postgresqlInputOption.FetchRows.ValueInt64Pointer(),
		ConnectTimeout:                     postgresqlInputOption.ConnectTimeout.ValueInt64Pointer(),
		SocketTimeout:                      postgresqlInputOption.SocketTimeout.ValueInt64Pointer(),
		DefaultTimeZone:                    postgresqlInputOption.DefaultTimeZone.ValueStringPointer(),
		PostgreSQLConnectionID:             postgresqlInputOption.PostgreSQLConnectionID.ValueInt64Pointer(),
		PostgreSQLInputOptionColumnOptions: toPostgreSQLInputOptionColumnOptions(postgresqlInputOption.PostgreSQLInputOptionColumnOptions),
		CustomVariableSettings:             model.ToCustomVariableSettingInputs(postgresqlInputOption.CustomVariableSettings),
	}
}

func toPostgreSQLInputOptionColumnOptions(columns *[]PostgreSQLInputOptionColumnOptions) *[]param.PostgreSQLInputOptionColumnOptions {
	if columns == nil {
		return nil
	}

	inputs := make([]param.PostgreSQLInputOptionColumnOptions, 0, len(*columns))
	for _, column := range *columns {
		inputs = append(inputs, param.PostgreSQLInputOptionColumnOptions{
			ColumnName:      column.ColumnName.ValueString(),
			ColumnValueType: column.ColumnValueType.ValueString(),
		})
	}
	return &inputs
}
