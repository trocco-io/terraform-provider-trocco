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

type DatabricksInputOption struct {
	DatabricksConnectionID types.Int64  `tfsdk:"databricks_connection_id"`
	CatalogName            types.String `tfsdk:"catalog_name"`
	SchemaName             types.String `tfsdk:"schema_name"`
	Query                  types.String `tfsdk:"query"`
	FetchRows              types.Int64  `tfsdk:"fetch_rows"`

	InputOptionColumns     types.List `tfsdk:"input_option_columns"`
	CustomVariableSettings types.List `tfsdk:"custom_variable_settings"`
}

type DatabricksInputOptionColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func NewDatabricksInputOption(ctx context.Context, inputOption *inputOptionEntities.DatabricksInputOption) *DatabricksInputOption {
	if inputOption == nil {
		return nil
	}

	result := &DatabricksInputOption{
		DatabricksConnectionID: types.Int64Value(inputOption.DatabricksConnectionID),
		CatalogName:            types.StringValue(inputOption.CatalogName),
		SchemaName:             types.StringValue(inputOption.SchemaName),
		Query:                  types.StringValue(inputOption.Query),
		FetchRows:              types.Int64Value(inputOption.FetchRows),
	}

	inputOptionColumns, err := newDatabricksInputOptionColumns(ctx, inputOption.InputOptionColumns)
	if err != nil {
		return nil
	}
	result.InputOptionColumns = inputOptionColumns

	CustomVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, inputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = CustomVariableSettings

	return result
}

func newDatabricksInputOptionColumns(
	ctx context.Context,
	inputOptionColumns []inputOptionEntities.DatabricksColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: DatabricksInputOptionColumn{}.attrTypes(),
	}

	if inputOptionColumns == nil {
		return types.ListNull(objectType), nil
	}

	columns := make([]DatabricksInputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := DatabricksInputOptionColumn{
			Name:   types.StringValue(input.Name),
			Type:   types.StringValue(input.Type),
			Format: types.StringPointerValue(input.Format),
		}
		columns = append(columns, column)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columns)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert input option columns to ListValue: %v", diags)
	}
	return listValue, nil
}

func (DatabricksInputOptionColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   types.StringType,
		"type":   types.StringType,
		"format": types.StringType,
	}
}

func (inputOption *DatabricksInputOption) ToInput(ctx context.Context) *inputOptionParameters.DatabricksInputOptionInput {
	if inputOption == nil {
		return nil
	}

	var columnOptionValues []DatabricksInputOptionColumn
	if !inputOption.InputOptionColumns.IsNull() && !inputOption.InputOptionColumns.IsUnknown() {
		diags := inputOption.InputOptionColumns.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	return &inputOptionParameters.DatabricksInputOptionInput{
		DatabricksConnectionID: inputOption.DatabricksConnectionID.ValueInt64(),
		CatalogName:            inputOption.CatalogName.ValueString(),
		SchemaName:             inputOption.SchemaName.ValueString(),
		Query:                  inputOption.Query.ValueString(),
		FetchRows:              model.NewNullableInt64(inputOption.FetchRows),

		InputOptionColumns:     toDatabricksInputOptionColumnsInput(columnOptionValues),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (inputOption *DatabricksInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateDatabricksInputOptionInput {
	if inputOption == nil {
		return nil
	}

	var columnOptionValues []DatabricksInputOptionColumn
	if !inputOption.InputOptionColumns.IsNull() {
		if !inputOption.InputOptionColumns.IsUnknown() {
			diags := inputOption.InputOptionColumns.ElementsAs(ctx, &columnOptionValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnOptionValues = []DatabricksInputOptionColumn{}
		}
	} else {
		columnOptionValues = []DatabricksInputOptionColumn{}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	return &inputOptionParameters.UpdateDatabricksInputOptionInput{
		DatabricksConnectionID: inputOption.DatabricksConnectionID.ValueInt64(),
		CatalogName:            inputOption.CatalogName.ValueString(),
		SchemaName:             inputOption.SchemaName.ValueString(),
		Query:                  inputOption.Query.ValueString(),
		FetchRows:              model.NewNullableInt64(inputOption.FetchRows),

		InputOptionColumns:     toDatabricksInputOptionColumnsInput(columnOptionValues),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func toDatabricksInputOptionColumnsInput(columns []DatabricksInputOptionColumn) []inputOptionParameters.DatabricksInputOptionColumn {
	if columns == nil {
		return nil
	}
	result := make([]inputOptionParameters.DatabricksInputOptionColumn, 0, len(columns))
	for _, column := range columns {
		result = append(result, inputOptionParameters.DatabricksInputOptionColumn{
			Name:   column.Name.ValueString(),
			Type:   column.Type.ValueString(),
			Format: column.Format.ValueStringPointer(),
		})
	}
	return result
}
