package input_options

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	param "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PostgreSQLInputOption struct {
	PostgreSQLConnectionID    types.Int64  `tfsdk:"postgresql_connection_id"`
	Database                  types.String `tfsdk:"database"`
	Schema                    types.String `tfsdk:"schema"`
	Query                     types.String `tfsdk:"query"`
	IncrementalLoadingEnabled types.Bool   `tfsdk:"incremental_loading_enabled"`
	Table                     types.String `tfsdk:"table"`
	IncrementalColumns        types.String `tfsdk:"incremental_columns"`
	LastRecord                types.String `tfsdk:"last_record"`
	FetchRows                 types.Int64  `tfsdk:"fetch_rows"`
	ConnectTimeout            types.Int64  `tfsdk:"connect_timeout"`
	SocketTimeout             types.Int64  `tfsdk:"socket_timeout"`
	DefaultTimeZone           types.String `tfsdk:"default_time_zone"`
	InputOptionColumns        types.List   `tfsdk:"input_option_columns"`
	CustomVariableSettings    types.List   `tfsdk:"custom_variable_settings"`
	InputOptionColumnOptions  types.List   `tfsdk:"input_option_column_options"`
}

type InputOptionColumnOptions struct {
	ColumnName      types.String `tfsdk:"column_name"`
	ColumnValueType types.String `tfsdk:"column_value_type"`
}

func (InputOptionColumnOptions) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"column_name":       types.StringType,
		"column_value_type": types.StringType,
	}
}

type PostgreSQLInputOptionColumn struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

func (PostgreSQLInputOptionColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
		"type": types.StringType,
	}
}

func NewPostgreSQLInputOption(postgresqlInputOption *input_option.PostgreSQLInputOption) *PostgreSQLInputOption {
	if postgresqlInputOption == nil {
		return nil
	}

	ctx := context.Background()

	result := &PostgreSQLInputOption{
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
	}

	inputOptionColumns, err := newPostgresqlInputOptionColumns(ctx, postgresqlInputOption.InputOptionColumns)
	if err != nil {
		return nil
	}
	result.InputOptionColumns = inputOptionColumns

	inputOptionColumnOptions, err := newInputOptionColumnOptions(ctx, postgresqlInputOption.InputOptionColumnOptions)
	if err != nil {
		return nil
	}
	result.InputOptionColumnOptions = inputOptionColumnOptions

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, postgresqlInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func newPostgresqlInputOptionColumns(
	ctx context.Context,
	inputOptionColumns []input_option.PostgreSQLInputOptionColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: PostgreSQLInputOptionColumn{}.attrTypes(),
	}

	if inputOptionColumns == nil {
		return types.ListNull(objectType), nil
	}

	columns := make([]PostgreSQLInputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := PostgreSQLInputOptionColumn{
			Name: types.StringValue(input.Name),
			Type: types.StringValue(input.Type),
		}
		columns = append(columns, column)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columns)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert input option columns to ListValue: %v", diags)
	}
	return listValue, nil
}

func newInputOptionColumnOptions(ctx context.Context, inputOptions []input_option.InputOptionColumnOptions) (types.List, error) {

	objectType := types.ObjectType{
		AttrTypes: InputOptionColumnOptions{}.attrTypes(),
	}

	if inputOptions == nil {
		return types.ListNull(objectType), nil
	}

	options := make([]InputOptionColumnOptions, 0, len(inputOptions))
	for _, input := range inputOptions {
		option := InputOptionColumnOptions{
			ColumnName:      types.StringValue(input.ColumnName),
			ColumnValueType: types.StringValue(input.ColumnValueType),
		}
		options = append(options, option)
	}

	setValue, diags := types.ListValueFrom(ctx, objectType, options)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert input option column options to ListValue: %v", diags)
	}
	return setValue, nil
}

func (postgresqlInputOption *PostgreSQLInputOption) ToInput() *param.PostgreSQLInputOptionInput {
	if postgresqlInputOption == nil {
		return nil
	}

	ctx := context.Background()

	var columnOptionValues []PostgreSQLInputOptionColumn
	if !postgresqlInputOption.InputOptionColumns.IsNull() && !postgresqlInputOption.InputOptionColumns.IsUnknown() {
		diags := postgresqlInputOption.InputOptionColumns.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}
	}

	var columnOptions []InputOptionColumnOptions
	if !postgresqlInputOption.InputOptionColumnOptions.IsNull() && !postgresqlInputOption.InputOptionColumnOptions.IsUnknown() {
		diags := postgresqlInputOption.InputOptionColumnOptions.ElementsAs(ctx, &columnOptions, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, postgresqlInputOption.CustomVariableSettings)

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
		InputOptionColumns:        toPostgresqlInputOptionColumnsInput(columnOptionValues),
		InputOptionColumnOptions:  model.WrapObjectList(toInputOptionColumnOptions(columnOptions)),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (postgresqlInputOption *PostgreSQLInputOption) ToUpdateInput() *param.UpdatePostgreSQLInputOptionInput {
	if postgresqlInputOption == nil {
		return nil
	}

	ctx := context.Background()

	var columnOptionValues []PostgreSQLInputOptionColumn
	if !postgresqlInputOption.InputOptionColumns.IsNull() {
		if !postgresqlInputOption.InputOptionColumns.IsUnknown() {
			diags := postgresqlInputOption.InputOptionColumns.ElementsAs(ctx, &columnOptionValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnOptionValues = []PostgreSQLInputOptionColumn{}
		}
	} else {
		columnOptionValues = []PostgreSQLInputOptionColumn{}
	}

	var columnOptions []InputOptionColumnOptions
	if !postgresqlInputOption.InputOptionColumnOptions.IsNull() {
		if !postgresqlInputOption.InputOptionColumnOptions.IsUnknown() {
			diags := postgresqlInputOption.InputOptionColumnOptions.ElementsAs(ctx, &columnOptions, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnOptions = []InputOptionColumnOptions{}
		}
	} else {
		columnOptions = []InputOptionColumnOptions{}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, postgresqlInputOption.CustomVariableSettings)

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
		InputOptionColumns:        toPostgresqlInputOptionColumnsInput(columnOptionValues),
		InputOptionColumnOptions:  model.WrapObjectList(toInputOptionColumnOptions(columnOptions)),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
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
