package input_options

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	param "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PostgreSQLInputOption struct {
	PostgreSQLConnectionID    types.Int64                    `tfsdk:"postgresql_connection_id"`
	Database                  types.String                   `tfsdk:"database"`
	Schema                    types.String                   `tfsdk:"schema"`
	Query                     types.String                   `tfsdk:"query"`
	IncrementalLoadingEnabled types.Bool                     `tfsdk:"incremental_loading_enabled"`
	Table                     types.String                   `tfsdk:"table"`
	IncrementalColumns        types.String                   `tfsdk:"incremental_columns"`
	LastRecord                types.String                   `tfsdk:"last_record"`
	FetchRows                 types.Int64                    `tfsdk:"fetch_rows"`
	ConnectTimeout            types.Int64                    `tfsdk:"connect_timeout"`
	SocketTimeout             types.Int64                    `tfsdk:"socket_timeout"`
	DefaultTimeZone           types.String                   `tfsdk:"default_time_zone"`
	InputOptionColumns        []PostgreSQLInputOptionColumn  `tfsdk:"input_option_columns"`
	CustomVariableSettings    *[]model.CustomVariableSetting `tfsdk:"custom_variable_settings"`
	InputOptionColumnOptions  []InputOptionColumnOptions     `tfsdk:"input_option_column_options"`
}

type InputOptionColumnOptions struct {
	ColumnName      types.String `tfsdk:"column_name"`
	ColumnValueType types.String `tfsdk:"column_value_type"`
}

type PostgreSQLInputOptionColumn struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

func NewPostgreSQLInputOption(postgresqlInputOption *input_option.PostgreSQLInputOption) *PostgreSQLInputOption {
	if postgresqlInputOption == nil {
		return nil
	}

	return &PostgreSQLInputOption{
		Database:                  types.StringValue(postgresqlInputOption.Database),
		Schema:                    types.StringValue(postgresqlInputOption.Schema),
		Table:                     types.StringPointerValue(postgresqlInputOption.Table),
		Query:                     types.StringPointerValue(postgresqlInputOption.Query),
		IncrementalColumns:        types.StringPointerValue(postgresqlInputOption.IncrementalColumns),
		LastRecord:                types.StringPointerValue(postgresqlInputOption.LastRecord),
		IncrementalLoadingEnabled: types.BoolValue(postgresqlInputOption.IncrementalLoadingEnabled),
		FetchRows:                 types.Int64Value(postgresqlInputOption.FetchRows),
		ConnectTimeout:            types.Int64Value(postgresqlInputOption.ConnectTimeout),
		SocketTimeout:             types.Int64Value(postgresqlInputOption.SocketTimeout),
		DefaultTimeZone:           types.StringValue(postgresqlInputOption.DefaultTimeZone),
		PostgreSQLConnectionID:    types.Int64Value(postgresqlInputOption.PostgreSQLConnectionID),
		InputOptionColumns:        newPostgresqlInputOptionColumns(postgresqlInputOption.InputOptionColumns),
		InputOptionColumnOptions:  newInputOptionColumnOptions(postgresqlInputOption.InputOptionColumnOptions),
		CustomVariableSettings:    model.NewCustomVariableSettings(postgresqlInputOption.CustomVariableSettings),
	}
}

func newPostgresqlInputOptionColumns(inputOptionColumns []input_option.PostgreSQLInputOptionColumn) []PostgreSQLInputOptionColumn {
	if inputOptionColumns == nil {
		return nil
	}
	columns := make([]PostgreSQLInputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := PostgreSQLInputOptionColumn{
			Name: types.StringValue(input.Name),
			Type: types.StringValue(input.Type),
		}
		columns = append(columns, column)
	}
	return columns
}

func newInputOptionColumnOptions(InputOptions []input_option.InputOptionColumnOptions) []InputOptionColumnOptions {
	columns := make([]InputOptionColumnOptions, 0, len(InputOptions))
	for _, input := range InputOptions {
		column := InputOptionColumnOptions{
			ColumnName:      types.StringValue(input.ColumnName),
			ColumnValueType: types.StringValue(input.ColumnValueType),
		}
		columns = append(columns, column)
	}
	return columns
}

func (postgresqlInputOption *PostgreSQLInputOption) ToInput() *param.PostgreSQLInputOptionInput {
	if postgresqlInputOption == nil {
		return nil
	}

	return &param.PostgreSQLInputOptionInput{
		Database:                  postgresqlInputOption.Database.ValueString(),
		Schema:                    model.NewNullableString(postgresqlInputOption.Schema),
		Table:                     model.NewNullableString(postgresqlInputOption.Table),
		Query:                     model.NewNullableString(postgresqlInputOption.Query),
		IncrementalColumns:        model.NewNullableString(postgresqlInputOption.IncrementalColumns),
		LastRecord:                model.NewNullableString(postgresqlInputOption.LastRecord),
		IncrementalLoadingEnabled: model.NewNullableBool(postgresqlInputOption.IncrementalLoadingEnabled),
		FetchRows:                 model.NewNullableInt64(postgresqlInputOption.FetchRows),
		ConnectTimeout:            model.NewNullableInt64(postgresqlInputOption.ConnectTimeout),
		SocketTimeout:             model.NewNullableInt64(postgresqlInputOption.SocketTimeout),
		DefaultTimeZone:           model.NewNullableString(postgresqlInputOption.DefaultTimeZone),
		PostgreSQLConnectionID:    postgresqlInputOption.PostgreSQLConnectionID.ValueInt64(),
		InputOptionColumns:        toPostgresqlInputOptionColumnsInput(postgresqlInputOption.InputOptionColumns),
		InputOptionColumnOptions:  model.WrapObjectList(toInputOptionColumnOptions(postgresqlInputOption.InputOptionColumnOptions)),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(postgresqlInputOption.CustomVariableSettings),
	}
}

func (postgresqlInputOption *PostgreSQLInputOption) ToUpdateInput() *param.UpdatePostgreSQLInputOptionInput {
	if postgresqlInputOption == nil {
		return nil
	}
	return &param.UpdatePostgreSQLInputOptionInput{
		Database:                  postgresqlInputOption.Database.ValueStringPointer(),
		Schema:                    model.NewNullableString(postgresqlInputOption.Schema),
		Table:                     model.NewNullableString(postgresqlInputOption.Table),
		Query:                     model.NewNullableString(postgresqlInputOption.Query),
		IncrementalColumns:        model.NewNullableString(postgresqlInputOption.IncrementalColumns),
		LastRecord:                model.NewNullableString(postgresqlInputOption.LastRecord),
		IncrementalLoadingEnabled: postgresqlInputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		FetchRows:                 model.NewNullableInt64(postgresqlInputOption.FetchRows),
		ConnectTimeout:            model.NewNullableInt64(postgresqlInputOption.ConnectTimeout),
		SocketTimeout:             model.NewNullableInt64(postgresqlInputOption.SocketTimeout),
		DefaultTimeZone:           model.NewNullableString(postgresqlInputOption.DefaultTimeZone),
		PostgreSQLConnectionID:    postgresqlInputOption.PostgreSQLConnectionID.ValueInt64Pointer(),
		InputOptionColumns:        toPostgresqlInputOptionColumnsInput(postgresqlInputOption.InputOptionColumns),
		InputOptionColumnOptions:  model.WrapObjectList(toInputOptionColumnOptions(postgresqlInputOption.InputOptionColumnOptions)),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(postgresqlInputOption.CustomVariableSettings),
	}
}

func toPostgresqlInputOptionColumnsInput(columns []PostgreSQLInputOptionColumn) []param.PostgreSQLInputOptionColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]param.PostgreSQLInputOptionColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, param.PostgreSQLInputOptionColumn{
			Name: column.Name.ValueString(),
			Type: column.Type.ValueString(),
		})
	}
	return inputs
}

func toInputOptionColumnOptions(options []InputOptionColumnOptions) *[]param.InputOptionColumnOptions {
	if options == nil {
		return nil
	}

	inputs := make([]param.InputOptionColumnOptions, len(options))
	for i, option := range options {
		inputs[i] = param.InputOptionColumnOptions{
			ColumnName:      option.ColumnName.ValueString(),
			ColumnValueType: option.ColumnValueType.ValueString(),
		}
	}

	return &inputs
}
