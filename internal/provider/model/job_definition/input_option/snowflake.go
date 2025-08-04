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

type SnowflakeInputOption struct {
	Warehouse              types.String `tfsdk:"warehouse"`
	Database               types.String `tfsdk:"database"`
	Schema                 types.String `tfsdk:"schema"`
	Query                  types.String `tfsdk:"query"`
	FetchRows              types.Int64  `tfsdk:"fetch_rows"`
	ConnectTimeout         types.Int64  `tfsdk:"connect_timeout"`
	SocketTimeout          types.Int64  `tfsdk:"socket_timeout"`
	SnowflakeConnectionID  types.Int64  `tfsdk:"snowflake_connection_id"`
	InputOptionColumns     types.List   `tfsdk:"input_option_columns"`
	CustomVariableSettings types.List   `tfsdk:"custom_variable_settings"`
}

type SnowflakeInputOptionColumn struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

func (SnowflakeInputOptionColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
		"type": types.StringType,
	}
}

func NewSnowflakeInputOption(ctx context.Context, snowflakeInputOption *inputOptionEntities.SnowflakeInputOption) *SnowflakeInputOption {
	if snowflakeInputOption == nil {
		return nil
	}

	result := &SnowflakeInputOption{
		Warehouse:             types.StringValue(snowflakeInputOption.Warehouse),
		Database:              types.StringValue(snowflakeInputOption.Database),
		Schema:                types.StringValue(snowflakeInputOption.Schema),
		Query:                 types.StringValue(snowflakeInputOption.Query),
		FetchRows:             types.Int64PointerValue(snowflakeInputOption.FetchRows),
		ConnectTimeout:        types.Int64PointerValue(snowflakeInputOption.ConnectTimeout),
		SocketTimeout:         types.Int64PointerValue(snowflakeInputOption.SocketTimeout),
		SnowflakeConnectionID: types.Int64Value(snowflakeInputOption.SnowflakeConnectionID),
	}

	columns, err := newSnowflakeInputOptionColumns(ctx, snowflakeInputOption.InputOptionColumns)
	if err != nil {
		return nil
	}
	result.InputOptionColumns = columns

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, snowflakeInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func newSnowflakeInputOptionColumns(
	ctx context.Context,
	inputOptionColumns []inputOptionEntities.SnowflakeInputOptionColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: SnowflakeInputOptionColumn{}.attrTypes(),
	}

	if inputOptionColumns == nil {
		return types.ListNull(objectType), nil
	}

	columns := make([]SnowflakeInputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := SnowflakeInputOptionColumn{
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

func (snowflakeInputOption *SnowflakeInputOption) ToInput(ctx context.Context) *inputOptionParameters.SnowflakeInputOptionInput {
	if snowflakeInputOption == nil {
		return nil
	}

	var columnValues []SnowflakeInputOptionColumn
	if !snowflakeInputOption.InputOptionColumns.IsNull() && !snowflakeInputOption.InputOptionColumns.IsUnknown() {
		diags := snowflakeInputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, snowflakeInputOption.CustomVariableSettings)

	return &inputOptionParameters.SnowflakeInputOptionInput{
		Warehouse:              snowflakeInputOption.Warehouse.ValueString(),
		Database:               snowflakeInputOption.Database.ValueString(),
		Schema:                 model.NewNullableString(snowflakeInputOption.Schema),
		Query:                  snowflakeInputOption.Query.ValueString(),
		FetchRows:              model.NewNullableInt64(snowflakeInputOption.FetchRows),
		ConnectTimeout:         model.NewNullableInt64(snowflakeInputOption.ConnectTimeout),
		SocketTimeout:          model.NewNullableInt64(snowflakeInputOption.SocketTimeout),
		SnowflakeConnectionID:  snowflakeInputOption.SnowflakeConnectionID.ValueInt64(),
		InputOptionColumns:     toSnowflakeInputOptionColumnsInput(columnValues),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (snowflakeInputOption *SnowflakeInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateSnowflakeInputOptionInput {
	if snowflakeInputOption == nil {
		return nil
	}

	var columnValues []SnowflakeInputOptionColumn
	if !snowflakeInputOption.InputOptionColumns.IsNull() {
		if !snowflakeInputOption.InputOptionColumns.IsUnknown() {
			diags := snowflakeInputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnValues = []SnowflakeInputOptionColumn{}
		}
	} else {
		columnValues = []SnowflakeInputOptionColumn{}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, snowflakeInputOption.CustomVariableSettings)

	return &inputOptionParameters.UpdateSnowflakeInputOptionInput{
		Warehouse:              snowflakeInputOption.Warehouse.ValueStringPointer(),
		Database:               snowflakeInputOption.Database.ValueStringPointer(),
		Schema:                 model.NewNullableString(snowflakeInputOption.Schema),
		Query:                  snowflakeInputOption.Query.ValueStringPointer(),
		FetchRows:              model.NewNullableInt64(snowflakeInputOption.FetchRows),
		ConnectTimeout:         model.NewNullableInt64(snowflakeInputOption.ConnectTimeout),
		SocketTimeout:          model.NewNullableInt64(snowflakeInputOption.SocketTimeout),
		SnowflakeConnectionID:  snowflakeInputOption.SnowflakeConnectionID.ValueInt64Pointer(),
		InputOptionColumns:     toSnowflakeInputOptionColumnsInput(columnValues),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func toSnowflakeInputOptionColumnsInput(columns []SnowflakeInputOptionColumn) []inputOptionParameters.SnowflakeInputOptionColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.SnowflakeInputOptionColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, inputOptionParameters.SnowflakeInputOptionColumn{
			Name: column.Name.ValueString(),
			Type: column.Type.ValueString(),
		})
	}
	return inputs
}
