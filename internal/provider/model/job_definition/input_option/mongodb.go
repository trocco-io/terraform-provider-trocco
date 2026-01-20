package input_options

import (
	"context"
	"encoding/json"
	"fmt"
	inputOptionEntities "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	inputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MongoDBInputOption struct {
	Database                  types.String `tfsdk:"database"`
	Collection                types.String `tfsdk:"collection"`
	Query                     types.String `tfsdk:"query"`
	IncrementalLoadingEnabled types.Bool   `tfsdk:"incremental_loading_enabled"`
	IncrementalColumns        types.String `tfsdk:"incremental_columns"`
	LastRecord                types.String `tfsdk:"last_record"`
	MongoDBConnectionID       types.Int64  `tfsdk:"mongodb_connection_id"`
	InputOptionColumns        types.List   `tfsdk:"input_option_columns"`
	CustomVariableSettings    types.List   `tfsdk:"custom_variable_settings"`
}

type MongodbInputOptionColumn struct {
	Name     types.String `tfsdk:"name"`
	Type     types.String `tfsdk:"type"`
	Format   types.String `tfsdk:"format"`
	Timezone types.String `tfsdk:"timezone"`
}

func NewMongodbInputOption(ctx context.Context, mongodbInputOption *inputOptionEntities.MongoDBInputOption) *MongoDBInputOption {
	if mongodbInputOption == nil {
		return nil
	}

	result := &MongoDBInputOption{
		Database:                  types.StringValue(mongodbInputOption.Database),
		Collection:                types.StringValue(mongodbInputOption.Collection),
		Query:                     types.StringPointerValue(mongodbInputOption.Query),
		IncrementalLoadingEnabled: types.BoolValue(mongodbInputOption.IncrementalLoadingEnabled),
		IncrementalColumns:        types.StringPointerValue(mongodbInputOption.IncrementalColumns),
		MongoDBConnectionID:       types.Int64Value(mongodbInputOption.MongoDBConnectionID),
	}

	if mongodbInputOption.LastRecord != nil {
		if lastRecordMap, ok := mongodbInputOption.LastRecord.(map[string]interface{}); ok {
			jsonBytes, err := json.Marshal(lastRecordMap)
			if err == nil {
				result.LastRecord = types.StringValue(string(jsonBytes))
			} else {
				result.LastRecord = types.StringNull()
			}
		} else {
			result.LastRecord = types.StringNull()
		}
	} else {
		result.LastRecord = types.StringNull()
	}

	inputOptionColumns, err := newMongodbInputOptionColumns(ctx, mongodbInputOption.InputOptionColumns)
	if err != nil {
		return nil
	}
	result.InputOptionColumns = inputOptionColumns

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, mongodbInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func newMongodbInputOptionColumns(
	ctx context.Context,
	inputOptionColumns []inputOptionEntities.MongodbInputOptionColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: MongodbInputOptionColumn{}.attrTypes(),
	}

	if inputOptionColumns == nil {
		return types.ListNull(objectType), nil
	}

	columns := make([]MongodbInputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := MongodbInputOptionColumn{
			Name:     types.StringValue(input.Name),
			Type:     types.StringValue(input.Type),
			Format:   types.StringPointerValue(input.Format),
			Timezone: types.StringPointerValue(input.Timezone),
		}
		columns = append(columns, column)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columns)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert mongodb input option columns to ListValue: %v", diags)
	}
	return listValue, nil
}

func (MongodbInputOptionColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":     types.StringType,
		"type":     types.StringType,
		"format":   types.StringType,
		"timezone": types.StringType,
	}
}

func (mongodbInputOption *MongoDBInputOption) ToInput(ctx context.Context) *inputOptionParameters.MongoDBInputOptionInput {
	if mongodbInputOption == nil {
		return nil
	}

	var columnOptionValues []MongodbInputOptionColumn
	if !mongodbInputOption.InputOptionColumns.IsNull() && !mongodbInputOption.InputOptionColumns.IsUnknown() {
		diags := mongodbInputOption.InputOptionColumns.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, mongodbInputOption.CustomVariableSettings)

	var lastRecord interface{}
	if !mongodbInputOption.LastRecord.IsNull() && !mongodbInputOption.LastRecord.IsUnknown() {
		lastRecordStr := mongodbInputOption.LastRecord.ValueString()
		var lastRecordMap map[string]interface{}
		if err := json.Unmarshal([]byte(lastRecordStr), &lastRecordMap); err == nil {
			lastRecord = lastRecordMap
		} else {
			lastRecord = nil
		}
	}

	return &inputOptionParameters.MongoDBInputOptionInput{
		Database:                  mongodbInputOption.Database.ValueString(),
		Collection:                mongodbInputOption.Collection.ValueString(),
		Query:                     model.NewNullableString(mongodbInputOption.Query),
		IncrementalLoadingEnabled: mongodbInputOption.IncrementalLoadingEnabled.ValueBool(),
		IncrementalColumns:        model.NewNullableString(mongodbInputOption.IncrementalColumns),
		LastRecord:                lastRecord,
		MongoDBConnectionID:       mongodbInputOption.MongoDBConnectionID.ValueInt64(),
		InputOptionColumns:        toMongodbInputOptionColumnsInput(columnOptionValues),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (mongodbInputOption *MongoDBInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateMongoDBInputOptionInput {
	if mongodbInputOption == nil {
		return nil
	}

	var columnOptionValues []MongodbInputOptionColumn
	if !mongodbInputOption.InputOptionColumns.IsNull() {
		if !mongodbInputOption.InputOptionColumns.IsUnknown() {
			diags := mongodbInputOption.InputOptionColumns.ElementsAs(ctx, &columnOptionValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnOptionValues = []MongodbInputOptionColumn{}
		}
	} else {
		columnOptionValues = []MongodbInputOptionColumn{}
	}

	inputOptionColumns := toMongodbInputOptionColumnsInput(columnOptionValues)
	customVarSettings := common.ExtractCustomVariableSettings(ctx, mongodbInputOption.CustomVariableSettings)

	var lastRecord interface{}
	if !mongodbInputOption.LastRecord.IsNull() && !mongodbInputOption.LastRecord.IsUnknown() {
		lastRecordStr := mongodbInputOption.LastRecord.ValueString()
		var lastRecordMap map[string]interface{}
		if err := json.Unmarshal([]byte(lastRecordStr), &lastRecordMap); err == nil {
			lastRecord = lastRecordMap
		}
	}

	return &inputOptionParameters.UpdateMongoDBInputOptionInput{
		Database:                  mongodbInputOption.Database.ValueStringPointer(),
		Collection:                mongodbInputOption.Collection.ValueStringPointer(),
		Query:                     model.NewNullableString(mongodbInputOption.Query),
		IncrementalLoadingEnabled: mongodbInputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		IncrementalColumns:        model.NewNullableString(mongodbInputOption.IncrementalColumns),
		LastRecord:                lastRecord,
		MongoDBConnectionID:       mongodbInputOption.MongoDBConnectionID.ValueInt64Pointer(),
		InputOptionColumns:        &inputOptionColumns,
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func toMongodbInputOptionColumnsInput(columns []MongodbInputOptionColumn) []inputOptionParameters.MongodbInputOptionColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.MongodbInputOptionColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, inputOptionParameters.MongodbInputOptionColumn{
			Name:     column.Name.ValueString(),
			Type:     column.Type.ValueString(),
			Format:   model.NewNullableString(column.Format),
			Timezone: model.NewNullableString(column.Timezone),
		})
	}
	return inputs
}
