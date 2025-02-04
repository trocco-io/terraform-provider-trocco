package input_options

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	param "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnowflakeInputOption struct {
	Warehouse              types.String                   `tfsdk:"warehouse"`
	Database               types.String                   `tfsdk:"database"`
	Schema                 types.String                   `tfsdk:"schema"`
	Query                  types.String                   `tfsdk:"query"`
	FetchRows              types.Int64                    `tfsdk:"fetch_rows"`
	ConnectTimeout         types.Int64                    `tfsdk:"connect_timeout"`
	SocketTimeout          types.Int64                    `tfsdk:"socket_timeout"`
	SnowflakeConnectionID  types.Int64                    `tfsdk:"snowflake_connection_id"`
	InputOptionColumns     []SnowflakeInputOptionColumn   `tfsdk:"input_option_columns"`
	CustomVariableSettings *[]model.CustomVariableSetting `tfsdk:"custom_variable_settings"`
}

type SnowflakeInputOptionColumn struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

func NewSnowflakeInputOption(snowflakeInputOption *input_option.SnowflakeInputOption) *SnowflakeInputOption {
	if snowflakeInputOption == nil {
		return nil
	}

	return &SnowflakeInputOption{
		Warehouse:              types.StringValue(snowflakeInputOption.Warehouse),
		Database:               types.StringValue(snowflakeInputOption.Database),
		Schema:                 types.StringValue(snowflakeInputOption.Schema),
		Query:                  types.StringValue(snowflakeInputOption.Query),
		FetchRows:              types.Int64PointerValue(snowflakeInputOption.FetchRows),
		ConnectTimeout:         types.Int64PointerValue(snowflakeInputOption.ConnectTimeout),
		SocketTimeout:          types.Int64PointerValue(snowflakeInputOption.SocketTimeout),
		SnowflakeConnectionID:  types.Int64Value(snowflakeInputOption.SnowflakeConnectionID),
		InputOptionColumns:     newSnowflakeInputOptionColumns(snowflakeInputOption.InputOptionColumns),
		CustomVariableSettings: model.NewCustomVariableSettings(snowflakeInputOption.CustomVariableSettings),
	}
}

func newSnowflakeInputOptionColumns(inputOptionColumns []input_option.SnowflakeInputOptionColumn) []SnowflakeInputOptionColumn {
	if inputOptionColumns == nil {
		return nil
	}
	columns := make([]SnowflakeInputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := SnowflakeInputOptionColumn{
			Name: types.StringValue(input.Name),
			Type: types.StringValue(input.Type),
		}
		columns = append(columns, column)
	}
	return columns
}

func (snowflakeInputOption *SnowflakeInputOption) ToInput() *param.SnowflakeInputOptionInput {
	if snowflakeInputOption == nil {
		return nil
	}

	return &param.SnowflakeInputOptionInput{
		Warehouse:              snowflakeInputOption.Warehouse.ValueString(),
		Database:               snowflakeInputOption.Database.ValueString(),
		Schema:                 model.NewNullableString(snowflakeInputOption.Schema),
		Query:                  snowflakeInputOption.Query.ValueString(),
		FetchRows:              model.NewNullableInt64(snowflakeInputOption.FetchRows),
		ConnectTimeout:         model.NewNullableInt64(snowflakeInputOption.ConnectTimeout),
		SocketTimeout:          model.NewNullableInt64(snowflakeInputOption.SocketTimeout),
		SnowflakeConnectionID:  snowflakeInputOption.SnowflakeConnectionID.ValueInt64(),
		InputOptionColumns:     toSnowflakeInputOptionColumnsInput(snowflakeInputOption.InputOptionColumns),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(snowflakeInputOption.CustomVariableSettings),
	}
}

func (snowflakeInputOption *SnowflakeInputOption) ToUpdateInput() *param.UpdateSnowflakeInputOptionInput {
	if snowflakeInputOption == nil {
		return nil
	}

	inputOptionColumns := toSnowflakeInputOptionColumnsInput(snowflakeInputOption.InputOptionColumns)

	return &param.UpdateSnowflakeInputOptionInput{
		Warehouse:              snowflakeInputOption.Warehouse.ValueStringPointer(),
		Database:               snowflakeInputOption.Database.ValueStringPointer(),
		Schema:                 model.NewNullableString(snowflakeInputOption.Schema),
		Query:                  snowflakeInputOption.Query.ValueStringPointer(),
		FetchRows:              model.NewNullableInt64(snowflakeInputOption.FetchRows),
		ConnectTimeout:         model.NewNullableInt64(snowflakeInputOption.ConnectTimeout),
		SocketTimeout:          model.NewNullableInt64(snowflakeInputOption.SocketTimeout),
		SnowflakeConnectionID:  snowflakeInputOption.SnowflakeConnectionID.ValueInt64Pointer(),
		InputOptionColumns:     inputOptionColumns,
		CustomVariableSettings: model.ToCustomVariableSettingInputs(snowflakeInputOption.CustomVariableSettings),
	}
}

func toSnowflakeInputOptionColumnsInput(columns []SnowflakeInputOptionColumn) []param.SnowflakeInputOptionColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]param.SnowflakeInputOptionColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, param.SnowflakeInputOptionColumn{
			Name: column.Name.ValueString(),
			Type: column.Type.ValueString(),
		})
	}
	return inputs
}
