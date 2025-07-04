package input_options

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	input_options2 "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MySQLInputOption struct {
	Database                  types.String `tfsdk:"database"`
	Table                     types.String `tfsdk:"table"`
	Query                     types.String `tfsdk:"query"`
	IncrementalColumns        types.String `tfsdk:"incremental_columns"`
	LastRecord                types.String `tfsdk:"last_record"`
	IncrementalLoadingEnabled types.Bool   `tfsdk:"incremental_loading_enabled"`
	FetchRows                 types.Int64  `tfsdk:"fetch_rows"`
	ConnectTimeout            types.Int64  `tfsdk:"connect_timeout"`
	SocketTimeout             types.Int64  `tfsdk:"socket_timeout"`
	DefaultTimeZone           types.String `tfsdk:"default_time_zone"`
	UseLegacyDatetimeCode     types.Bool   `tfsdk:"use_legacy_datetime_code"`
	MySQLConnectionID         types.Int64  `tfsdk:"mysql_connection_id"`
	InputOptionColumns        types.List   `tfsdk:"input_option_columns"`
	CustomVariableSettings    types.List   `tfsdk:"custom_variable_settings"`
}

type InputOptionColumn struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

func NewMysqlInputOption(ctx context.Context, mysqlInputOption *input_option.MySQLInputOption) *MySQLInputOption {
	if mysqlInputOption == nil {
		return nil
	}

	result := &MySQLInputOption{
		Database:                  types.StringValue(mysqlInputOption.Database),
		Table:                     types.StringPointerValue(mysqlInputOption.Table),
		Query:                     types.StringPointerValue(mysqlInputOption.Query),
		IncrementalColumns:        types.StringPointerValue(mysqlInputOption.IncrementalColumns),
		LastRecord:                types.StringPointerValue(mysqlInputOption.LastRecord),
		IncrementalLoadingEnabled: types.BoolValue(mysqlInputOption.IncrementalLoadingEnabled),
		FetchRows:                 types.Int64Value(mysqlInputOption.FetchRows),
		ConnectTimeout:            types.Int64Value(mysqlInputOption.ConnectTimeout),
		SocketTimeout:             types.Int64Value(mysqlInputOption.SocketTimeout),
		DefaultTimeZone:           types.StringPointerValue(mysqlInputOption.DefaultTimeZone),
		UseLegacyDatetimeCode:     types.BoolPointerValue(mysqlInputOption.UseLegacyDatetimeCode),
		MySQLConnectionID:         types.Int64Value(mysqlInputOption.MySQLConnectionID),
	}

	inputOptionColumns, err := newInputOptionColumns(ctx, mysqlInputOption.InputOptionColumns)
	if err != nil {
		return nil
	}
	result.InputOptionColumns = inputOptionColumns

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, mysqlInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func newInputOptionColumns(
	ctx context.Context,
	inputOptionColumns []input_option.InputOptionColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: InputOptionColumn{}.attrTypes(),
	}

	if inputOptionColumns == nil {
		return types.ListNull(objectType), nil
	}

	columns := make([]InputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := InputOptionColumn{
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

func (InputOptionColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
		"type": types.StringType,
	}
}

func (mysqlInputOption *MySQLInputOption) ToInput(ctx context.Context) *input_options2.MySQLInputOptionInput {
	if mysqlInputOption == nil {
		return nil
	}

	var columnOptionValues []InputOptionColumn
	if !mysqlInputOption.InputOptionColumns.IsNull() && !mysqlInputOption.InputOptionColumns.IsUnknown() {
		diags := mysqlInputOption.InputOptionColumns.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, mysqlInputOption.CustomVariableSettings)

	return &input_options2.MySQLInputOptionInput{
		Database:                  mysqlInputOption.Database.ValueString(),
		Table:                     model.NewNullableString(mysqlInputOption.Table),
		Query:                     model.NewNullableString(mysqlInputOption.Query),
		IncrementalColumns:        model.NewNullableString(mysqlInputOption.IncrementalColumns),
		LastRecord:                model.NewNullableString(mysqlInputOption.LastRecord),
		IncrementalLoadingEnabled: mysqlInputOption.IncrementalLoadingEnabled.ValueBool(),
		FetchRows:                 mysqlInputOption.FetchRows.ValueInt64(),
		ConnectTimeout:            mysqlInputOption.ConnectTimeout.ValueInt64(),
		SocketTimeout:             mysqlInputOption.SocketTimeout.ValueInt64(),
		DefaultTimeZone:           model.NewNullableString(mysqlInputOption.DefaultTimeZone),
		UseLegacyDatetimeCode:     mysqlInputOption.UseLegacyDatetimeCode.ValueBool(),
		MySQLConnectionID:         mysqlInputOption.MySQLConnectionID.ValueInt64(),
		InputOptionColumns:        toMysqlInputOptionColumnsInput(columnOptionValues),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (mysqlInputOption *MySQLInputOption) ToUpdateInput(ctx context.Context) *input_options2.UpdateMySQLInputOptionInput {
	if mysqlInputOption == nil {
		return nil
	}

	var columnOptionValues []InputOptionColumn
	if !mysqlInputOption.InputOptionColumns.IsNull() {
		if !mysqlInputOption.InputOptionColumns.IsUnknown() {
			diags := mysqlInputOption.InputOptionColumns.ElementsAs(ctx, &columnOptionValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnOptionValues = []InputOptionColumn{}
		}
	} else {
		columnOptionValues = []InputOptionColumn{}
	}

	inputOptionColumns := toMysqlInputOptionColumnsInput(columnOptionValues)
	customVarSettings := common.ExtractCustomVariableSettings(ctx, mysqlInputOption.CustomVariableSettings)

	return &input_options2.UpdateMySQLInputOptionInput{
		Database:                  mysqlInputOption.Database.ValueStringPointer(),
		Table:                     model.NewNullableString(mysqlInputOption.Table),
		Query:                     model.NewNullableString(mysqlInputOption.Query),
		IncrementalColumns:        model.NewNullableString(mysqlInputOption.IncrementalColumns),
		LastRecord:                model.NewNullableString(mysqlInputOption.LastRecord),
		IncrementalLoadingEnabled: mysqlInputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		FetchRows:                 mysqlInputOption.FetchRows.ValueInt64Pointer(),
		ConnectTimeout:            mysqlInputOption.ConnectTimeout.ValueInt64Pointer(),
		SocketTimeout:             mysqlInputOption.SocketTimeout.ValueInt64Pointer(),
		DefaultTimeZone:           model.NewNullableString(mysqlInputOption.DefaultTimeZone),
		UseLegacyDatetimeCode:     mysqlInputOption.UseLegacyDatetimeCode.ValueBoolPointer(),
		MySQLConnectionID:         mysqlInputOption.MySQLConnectionID.ValueInt64Pointer(),
		InputOptionColumns:        &inputOptionColumns,
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func toMysqlInputOptionColumnsInput(columns []InputOptionColumn) []input_options2.InputOptionColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]input_options2.InputOptionColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, input_options2.InputOptionColumn{
			Name: column.Name.ValueString(),
			Type: column.Type.ValueString(),
		})
	}
	return inputs
}
