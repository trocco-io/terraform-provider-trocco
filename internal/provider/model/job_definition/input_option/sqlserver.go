package input_options

import (
	"context"
	"fmt"
	inputOptionEntities "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	inputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SQLServerInputOption struct {
	SQLServerConnectionID     types.Int64  `tfsdk:"sqlserver_connection_id"`
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

type SQLServerInputOptionColumn struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

func (SQLServerInputOptionColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
		"type": types.StringType,
	}
}

func NewSQLServerInputOption(ctx context.Context, sqlserverInputOption *inputOptionEntities.SQLServerInputOption) *SQLServerInputOption {
	if sqlserverInputOption == nil {
		return nil
	}

	result := &SQLServerInputOption{
		Database:                  types.StringValue(sqlserverInputOption.Database),
		Schema:                    types.StringPointerValue(sqlserverInputOption.Schema),
		Table:                     types.StringPointerValue(sqlserverInputOption.Table),
		Query:                     types.StringPointerValue(sqlserverInputOption.Query),
		IncrementalColumns:        types.StringPointerValue(sqlserverInputOption.IncrementalColumns),
		LastRecord:                types.StringPointerValue(sqlserverInputOption.LastRecord),
		IncrementalLoadingEnabled: types.BoolValue(sqlserverInputOption.IncrementalLoadingEnabled),
		FetchRows:                 types.Int64Value(sqlserverInputOption.FetchRows),
		ConnectTimeout:            types.Int64Value(sqlserverInputOption.ConnectTimeout),
		SocketTimeout:             types.Int64Value(sqlserverInputOption.SocketTimeout),
		DefaultTimeZone:           types.StringValue(sqlserverInputOption.DefaultTimeZone),
		SQLServerConnectionID:     types.Int64Value(sqlserverInputOption.SQLServerConnectionID),
	}

	inputOptionColumns, err := newSQLServerInputOptionColumns(ctx, sqlserverInputOption.InputOptionColumns)
	if err != nil {
		return nil
	}
	result.InputOptionColumns = inputOptionColumns

	inputOptionColumnOptions, err := newInputOptionColumnOptions(ctx, sqlserverInputOption.InputOptionColumnOptions)
	if err != nil {
		return nil
	}
	result.InputOptionColumnOptions = inputOptionColumnOptions

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, sqlserverInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func newSQLServerInputOptionColumns(
	ctx context.Context,
	inputOptionColumns []inputOptionEntities.SQLServerInputOptionColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: SQLServerInputOptionColumn{}.attrTypes(),
	}

	if inputOptionColumns == nil {
		return types.ListNull(objectType), nil
	}

	columns := make([]SQLServerInputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		columns = append(columns, SQLServerInputOptionColumn{
			Name: types.StringValue(input.Name),
			Type: types.StringValue(input.Type),
		})
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columns)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert input option columns to ListValue: %v", diags)
	}
	return listValue, nil
}

func (sqlserverInputOption *SQLServerInputOption) ToInput(ctx context.Context) *inputOptionParameters.SQLServerInputOptionInput {
	if sqlserverInputOption == nil {
		return nil
	}

	var columnValues []SQLServerInputOptionColumn
	if !sqlserverInputOption.InputOptionColumns.IsNull() && !sqlserverInputOption.InputOptionColumns.IsUnknown() {
		diags := sqlserverInputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
		if diags.HasError() {
			return nil
		}
	}

	var columnOptions []InputOptionColumnOptions
	if !sqlserverInputOption.InputOptionColumnOptions.IsNull() && !sqlserverInputOption.InputOptionColumnOptions.IsUnknown() {
		diags := sqlserverInputOption.InputOptionColumnOptions.ElementsAs(ctx, &columnOptions, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, sqlserverInputOption.CustomVariableSettings)

	return &inputOptionParameters.SQLServerInputOptionInput{
		Database:                  sqlserverInputOption.Database.ValueString(),
		Schema:                    model.NewNullableString(sqlserverInputOption.Schema),
		Table:                     model.NewNullableString(sqlserverInputOption.Table),
		Query:                     model.NewNullableString(sqlserverInputOption.Query),
		IncrementalColumns:        model.NewNullableString(sqlserverInputOption.IncrementalColumns),
		LastRecord:                model.NewNullableString(sqlserverInputOption.LastRecord),
		IncrementalLoadingEnabled: model.NewNullableBool(sqlserverInputOption.IncrementalLoadingEnabled),
		FetchRows:                 model.NewNullableInt64(sqlserverInputOption.FetchRows),
		ConnectTimeout:            model.NewNullableInt64(sqlserverInputOption.ConnectTimeout),
		SocketTimeout:             model.NewNullableInt64(sqlserverInputOption.SocketTimeout),
		DefaultTimeZone:           model.NewNullableString(sqlserverInputOption.DefaultTimeZone),
		SQLServerConnectionID:     sqlserverInputOption.SQLServerConnectionID.ValueInt64(),
		InputOptionColumns:        toSQLServerInputOptionColumnsInput(columnValues),
		InputOptionColumnOptions:  model.WrapObjectList(toInputOptionColumnOptions(columnOptions)),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (sqlserverInputOption *SQLServerInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateSQLServerInputOptionInput {
	if sqlserverInputOption == nil {
		return nil
	}

	columnValues := []SQLServerInputOptionColumn{}
	if !sqlserverInputOption.InputOptionColumns.IsNull() && !sqlserverInputOption.InputOptionColumns.IsUnknown() {
		diags := sqlserverInputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
		if diags.HasError() {
			return nil
		}
	}

	columnOptions := []InputOptionColumnOptions{}
	if !sqlserverInputOption.InputOptionColumnOptions.IsNull() && !sqlserverInputOption.InputOptionColumnOptions.IsUnknown() {
		diags := sqlserverInputOption.InputOptionColumnOptions.ElementsAs(ctx, &columnOptions, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, sqlserverInputOption.CustomVariableSettings)

	return &inputOptionParameters.UpdateSQLServerInputOptionInput{
		Database:                  sqlserverInputOption.Database.ValueStringPointer(),
		Schema:                    model.NewNullableString(sqlserverInputOption.Schema),
		Table:                     model.NewNullableString(sqlserverInputOption.Table),
		Query:                     model.NewNullableString(sqlserverInputOption.Query),
		IncrementalColumns:        model.NewNullableString(sqlserverInputOption.IncrementalColumns),
		LastRecord:                model.NewNullableString(sqlserverInputOption.LastRecord),
		IncrementalLoadingEnabled: sqlserverInputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		FetchRows:                 model.NewNullableInt64(sqlserverInputOption.FetchRows),
		ConnectTimeout:            model.NewNullableInt64(sqlserverInputOption.ConnectTimeout),
		SocketTimeout:             model.NewNullableInt64(sqlserverInputOption.SocketTimeout),
		DefaultTimeZone:           model.NewNullableString(sqlserverInputOption.DefaultTimeZone),
		SQLServerConnectionID:     sqlserverInputOption.SQLServerConnectionID.ValueInt64Pointer(),
		InputOptionColumns:        toSQLServerInputOptionColumnsInput(columnValues),
		InputOptionColumnOptions:  model.WrapObjectList(toInputOptionColumnOptions(columnOptions)),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func toSQLServerInputOptionColumnsInput(columns []SQLServerInputOptionColumn) []inputOptionParameters.SQLServerInputOptionColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.SQLServerInputOptionColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, inputOptionParameters.SQLServerInputOptionColumn{
			Name: column.Name.ValueString(),
			Type: column.Type.ValueString(),
		})
	}
	return inputs
}
