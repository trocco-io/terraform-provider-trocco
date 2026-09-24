package input_options

import (
	"context"
	inputOptionEntities "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	inputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OracleInputOption struct {
	OracleConnectionID        types.Int64  `tfsdk:"oracle_connection_id"`
	Database                  types.String `tfsdk:"database"`
	ConnectionType            types.String `tfsdk:"connection_type"`
	NetServiceName            types.String `tfsdk:"net_service_name"`
	Schema                    types.String `tfsdk:"schema"`
	Query                     types.String `tfsdk:"query"`
	IncrementalLoadingEnabled types.Bool   `tfsdk:"incremental_loading_enabled"`
	Table                     types.String `tfsdk:"table"`
	IncrementalColumns        types.String `tfsdk:"incremental_columns"`
	DefaultTimeZone           types.String `tfsdk:"default_time_zone"`
	InputOptionColumns        types.List   `tfsdk:"input_option_columns"`
	InputOptionColumnOptions  types.List   `tfsdk:"input_option_column_options"`
	CustomVariableSettings    types.List   `tfsdk:"custom_variable_settings"`
}

type OracleInputOptionColumn struct {
	Name     types.String `tfsdk:"name"`
	Type     types.String `tfsdk:"type"`
	Format   types.String `tfsdk:"format"`
	Timezone types.String `tfsdk:"timezone"`
}

type OracleInputOptionColumnOption struct {
	ColumnName      types.String `tfsdk:"column_name"`
	ColumnValueType types.String `tfsdk:"column_value_type"`
}

func NewOracleInputOption(ctx context.Context, oracleInputOption *inputOptionEntities.OracleInputOption) *OracleInputOption {
	if oracleInputOption == nil {
		return nil
	}

	result := &OracleInputOption{
		OracleConnectionID:        types.Int64Value(oracleInputOption.OracleConnectionID),
		Database:                  types.StringPointerValue(oracleInputOption.Database),
		ConnectionType:            types.StringValue(oracleInputOption.ConnectionType),
		NetServiceName:            types.StringPointerValue(oracleInputOption.NetServiceName),
		Schema:                    types.StringPointerValue(oracleInputOption.Schema),
		Query:                     types.StringPointerValue(oracleInputOption.Query),
		IncrementalLoadingEnabled: types.BoolValue(oracleInputOption.IncrementalLoadingEnabled),
		Table:                     types.StringPointerValue(oracleInputOption.Table),
		IncrementalColumns:        types.StringPointerValue(oracleInputOption.IncrementalColumns),
		DefaultTimeZone:           types.StringValue(oracleInputOption.DefaultTimeZone),
	}

	if oracleInputOption.InputOptionColumns != nil {
		inputOptionColumnList, diags := oracleInputOptionColumnsToList(ctx, *oracleInputOption.InputOptionColumns)
		if diags.HasError() {
			return nil
		}
		result.InputOptionColumns = inputOptionColumnList
	} else {
		result.InputOptionColumns = types.ListNull(types.ObjectType{})
	}

	if oracleInputOption.InputOptionColumnOptions != nil {
		inputOptionColumnOptionList, diags := oracleInputOptionColumnOptionsToList(ctx, *oracleInputOption.InputOptionColumnOptions)
		if diags.HasError() {
			return nil
		}
		result.InputOptionColumnOptions = inputOptionColumnOptionList
	} else {
		result.InputOptionColumnOptions = types.ListNull(types.ObjectType{})
	}

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, oracleInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func (o *OracleInputOption) ToInput(ctx context.Context) *inputOptionParameters.OracleInputOptionInput {
	if o == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, o.CustomVariableSettings)
	inputOptionCols := extractOracleInputOptionColumns(ctx, o.InputOptionColumns)
	inputOptionColOpts := extractOracleInputOptionColumnOptions(ctx, o.InputOptionColumnOptions)

	return &inputOptionParameters.OracleInputOptionInput{
		OracleConnectionID:        o.OracleConnectionID.ValueInt64(),
		Database:                  model.NewNullableString(o.Database),
		ConnectionType:            model.NewNullableString(o.ConnectionType),
		NetServiceName:            model.NewNullableString(o.NetServiceName),
		Schema:                    model.NewNullableString(o.Schema),
		Query:                     o.Query.ValueStringPointer(),
		IncrementalLoadingEnabled: o.IncrementalLoadingEnabled.ValueBool(),
		Table:                     o.Table.ValueStringPointer(),
		IncrementalColumns:        model.NewNullableString(o.IncrementalColumns),
		DefaultTimeZone:           o.DefaultTimeZone.ValueStringPointer(),
		InputOptionColumns:        inputOptionCols,
		InputOptionColumnOptions:  inputOptionColOpts,
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (o *OracleInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateOracleInputOptionInput {
	if o == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, o.CustomVariableSettings)
	inputOptionCols := extractOracleInputOptionColumns(ctx, o.InputOptionColumns)
	inputOptionColOpts := extractOracleInputOptionColumnOptions(ctx, o.InputOptionColumnOptions)

	return &inputOptionParameters.UpdateOracleInputOptionInput{
		OracleConnectionID:        o.OracleConnectionID.ValueInt64Pointer(),
		Database:                  model.NewNullableString(o.Database),
		ConnectionType:            model.NewNullableString(o.ConnectionType),
		NetServiceName:            model.NewNullableString(o.NetServiceName),
		Schema:                    model.NewNullableString(o.Schema),
		Query:                     o.Query.ValueStringPointer(),
		IncrementalLoadingEnabled: o.IncrementalLoadingEnabled.ValueBoolPointer(),
		Table:                     o.Table.ValueStringPointer(),
		IncrementalColumns:        model.NewNullableString(o.IncrementalColumns),
		DefaultTimeZone:           o.DefaultTimeZone.ValueStringPointer(),
		InputOptionColumns:        inputOptionCols,
		InputOptionColumnOptions:  inputOptionColOpts,
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func oracleInputOptionColumnsToList(ctx context.Context, columns []inputOptionEntities.OracleInputOptionColumn) (types.List, diag.Diagnostics) {
	columnModels := make([]OracleInputOptionColumn, len(columns))
	for i, col := range columns {
		columnModels[i] = OracleInputOptionColumn{
			Name:     types.StringValue(col.Name),
			Type:     types.StringValue(col.Type),
			Format:   types.StringPointerValue(col.Format),
			Timezone: types.StringPointerValue(col.Timezone),
		}
	}
	return types.ListValueFrom(ctx, types.ObjectType{}, columnModels)
}

func oracleInputOptionColumnOptionsToList(ctx context.Context, columnOptions []inputOptionEntities.OracleInputOptionColumnOption) (types.List, diag.Diagnostics) {
	columnOptionModels := make([]OracleInputOptionColumnOption, len(columnOptions))
	for i, colOpt := range columnOptions {
		columnOptionModels[i] = OracleInputOptionColumnOption{
			ColumnName:      types.StringValue(colOpt.ColumnName),
			ColumnValueType: types.StringValue(colOpt.ColumnValueType),
		}
	}
	return types.ListValueFrom(ctx, types.ObjectType{}, columnOptionModels)
}

func extractOracleInputOptionColumns(ctx context.Context, columns types.List) *[]inputOptionParameters.OracleInputOptionColumnInput {
	if columns.IsNull() || columns.IsUnknown() {
		return nil
	}

	var columnModels []OracleInputOptionColumn
	columns.ElementsAs(ctx, &columnModels, false)

	columnInputs := make([]inputOptionParameters.OracleInputOptionColumnInput, len(columnModels))
	for i, col := range columnModels {
		columnInputs[i] = inputOptionParameters.OracleInputOptionColumnInput{
			Name:     col.Name.ValueString(),
			Type:     col.Type.ValueString(),
			Format:   model.NewNullableString(col.Format),
			Timezone: model.NewNullableString(col.Timezone),
		}
	}

	return &columnInputs
}

func extractOracleInputOptionColumnOptions(ctx context.Context, columnOptions types.List) *[]inputOptionParameters.OracleInputOptionColumnOptionInput {
	if columnOptions.IsNull() || columnOptions.IsUnknown() {
		return nil
	}

	var columnOptionModels []OracleInputOptionColumnOption
	columnOptions.ElementsAs(ctx, &columnOptionModels, false)

	columnOptionInputs := make([]inputOptionParameters.OracleInputOptionColumnOptionInput, len(columnOptionModels))
	for i, colOpt := range columnOptionModels {
		columnOptionInputs[i] = inputOptionParameters.OracleInputOptionColumnOptionInput{
			ColumnName:      colOpt.ColumnName.ValueString(),
			ColumnValueType: colOpt.ColumnValueType.ValueString(),
		}
	}

	return &columnOptionInputs
}
