package input_options

import (
	inputOptionEntity "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	inputOptionParams "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MySQLInputOption struct {
	Database                  types.String                   `tfsdk:"database"`
	Table                     types.String                   `tfsdk:"table"`
	Query                     types.String                   `tfsdk:"query"`
	IncrementalColumns        types.String                   `tfsdk:"incremental_columns"`
	LastRecord                types.String                   `tfsdk:"last_record"`
	IncrementalLoadingEnabled types.Bool                     `tfsdk:"incremental_loading_enabled"`
	FetchRows                 types.Int64                    `tfsdk:"fetch_rows"`
	ConnectTimeout            types.Int64                    `tfsdk:"connect_timeout"`
	SocketTimeout             types.Int64                    `tfsdk:"socket_timeout"`
	DefaultTimeZone           types.String                   `tfsdk:"default_time_zone"`
	UseLegacyDatetimeCode     types.Bool                     `tfsdk:"use_legacy_datetime_code"`
	MySQLConnectionID         types.Int64                    `tfsdk:"mysql_connection_id"`
	InputOptionColumns        []InputOptionColumn            `tfsdk:"input_option_columns"`
	CustomVariableSettings    *[]model.CustomVariableSetting `tfsdk:"custom_variable_settings"`
}

type InputOptionColumn struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

func NewMysqlInputOption(mysqlInputOption *inputOptionEntity.MySQLInputOption) *MySQLInputOption {
	if mysqlInputOption == nil {
		return nil
	}

	return &MySQLInputOption{
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
		InputOptionColumns:        newInputOptionColumns(mysqlInputOption.InputOptionColumns),
		CustomVariableSettings:    model.NewCustomVariableSettings(mysqlInputOption.CustomVariableSettings),
	}
}

func newInputOptionColumns(inputOptionColumns []inputOptionEntity.InputOptionColumn) []InputOptionColumn {
	if inputOptionColumns == nil {
		return nil
	}
	columns := make([]InputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := InputOptionColumn{
			Name: types.StringValue(input.Name),
			Type: types.StringValue(input.Type),
		}
		columns = append(columns, column)
	}
	return columns
}

func (mysqlInputOption *MySQLInputOption) ToInput() *inputOptionParams.MySQLInputOptionInput {
	if mysqlInputOption == nil {
		return nil
	}

	return &inputOptionParams.MySQLInputOptionInput{
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
		InputOptionColumns:        toMysqlInputOptionColumnsInput(mysqlInputOption.InputOptionColumns),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(mysqlInputOption.CustomVariableSettings),
	}
}

func (mysqlInputOption *MySQLInputOption) ToUpdateInput() *inputOptionParams.UpdateMySQLInputOptionInput {
	if mysqlInputOption == nil {
		return nil
	}

	inputOptionColumns := toMysqlInputOptionColumnsInput(mysqlInputOption.InputOptionColumns)

	return &inputOptionParams.UpdateMySQLInputOptionInput{
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
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(mysqlInputOption.CustomVariableSettings),
	}
}

func toMysqlInputOptionColumnsInput(columns []InputOptionColumn) []inputOptionParams.InputOptionColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]inputOptionParams.InputOptionColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, inputOptionParams.InputOptionColumn{
			Name: column.Name.ValueString(),
			Type: column.Type.ValueString(),
		})
	}
	return inputs
}
