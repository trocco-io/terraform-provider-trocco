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

type MarketoInputOption struct {
	MarketoConnectionID             types.Int64  `tfsdk:"marketo_connection_id"`
	Target                          types.String `tfsdk:"target"`
	FromDate                        types.String `tfsdk:"from_date"`
	EndDate                         types.String `tfsdk:"end_date"`
	UseUpdatedAt                    types.Bool   `tfsdk:"use_updated_at"`
	PollingIntervalSecond           types.Int64  `tfsdk:"polling_interval_second"`
	BulkJobTimeoutSecond            types.Int64  `tfsdk:"bulk_job_timeout_second"`
	ActivityTypeIDs                 types.List   `tfsdk:"activity_type_ids"`
	CustomObjectAPIName             types.String `tfsdk:"custom_object_api_name"`
	CustomObjectFilterType          types.String `tfsdk:"custom_object_filter_type"`
	CustomObjectFilterFromValue     types.Int64  `tfsdk:"custom_object_filter_from_value"`
	CustomObjectFilterToValue       types.Int64  `tfsdk:"custom_object_filter_to_value"`
	CustomObjectFields              types.List   `tfsdk:"custom_object_fields"`
	ListIDs                         types.String `tfsdk:"list_ids"`
	ProgramIDs                      types.String `tfsdk:"program_ids"`
	RootID                          types.Int64  `tfsdk:"root_id"`
	RootType                        types.String `tfsdk:"root_type"`
	MaxDepth                        types.Int64  `tfsdk:"max_depth"`
	Workspace                       types.String `tfsdk:"workspace"`
	InputOptionColumns              types.List   `tfsdk:"input_option_columns"`
	MarketoInputOptionFilterColumns types.List   `tfsdk:"marketo_input_option_filter_columns"`
	CustomVariableSettings          types.List   `tfsdk:"custom_variable_settings"`
}

type MarketoColumn struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

type MarketoFilterColumn struct {
	Name types.String `tfsdk:"name"`
}

type MarketoCustomObjectField struct {
	Name types.String `tfsdk:"name"`
}

func (MarketoColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
		"type": types.StringType,
	}
}

func (MarketoFilterColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
	}
}

func (MarketoCustomObjectField) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
	}
}

func NewMarketoInputOption(ctx context.Context, marketoInputOption *inputOptionEntities.MarketoInputOption) *MarketoInputOption {
	if marketoInputOption == nil {
		return nil
	}

	result := &MarketoInputOption{
		MarketoConnectionID:         types.Int64Value(marketoInputOption.MarketoConnectionID),
		Target:                      types.StringValue(marketoInputOption.Target),
		FromDate:                    types.StringPointerValue(marketoInputOption.FromDate),
		EndDate:                     types.StringPointerValue(marketoInputOption.EndDate),
		UseUpdatedAt:                types.BoolPointerValue(marketoInputOption.UseUpdatedAt),
		PollingIntervalSecond:       types.Int64PointerValue(marketoInputOption.PollingIntervalSecond),
		BulkJobTimeoutSecond:        types.Int64PointerValue(marketoInputOption.BulkJobTimeoutSecond),
		CustomObjectAPIName:         types.StringPointerValue(marketoInputOption.CustomObjectAPIName),
		CustomObjectFilterType:      types.StringPointerValue(marketoInputOption.CustomObjectFilterType),
		CustomObjectFilterFromValue: types.Int64PointerValue(marketoInputOption.CustomObjectFilterFromValue),
		CustomObjectFilterToValue:   types.Int64PointerValue(marketoInputOption.CustomObjectFilterToValue),
		ListIDs:                     types.StringPointerValue(marketoInputOption.ListIDs),
		ProgramIDs:                  types.StringPointerValue(marketoInputOption.ProgramIDs),
		RootID:                      types.Int64PointerValue(marketoInputOption.RootID),
		RootType:                    types.StringPointerValue(marketoInputOption.RootType),
		MaxDepth:                    types.Int64PointerValue(marketoInputOption.MaxDepth),
		Workspace:                   types.StringPointerValue(marketoInputOption.Workspace),
	}

	// Convert activity type IDs
	if marketoInputOption.ActivityTypeIDs != nil {
		activityTypeIDList, diags := types.ListValueFrom(ctx, types.Int64Type, *marketoInputOption.ActivityTypeIDs)
		if diags.HasError() {
			return nil
		}
		result.ActivityTypeIDs = activityTypeIDList
	} else {
		result.ActivityTypeIDs = types.ListNull(types.Int64Type)
	}

	columns, err := newMarketoColumns(ctx, marketoInputOption.InputOptionColumns)
	if err != nil {
		return nil
	}
	result.InputOptionColumns = columns

	filterColumns, err := newMarketoFilterColumns(ctx, marketoInputOption.MarketoInputOptionFilterColumns)
	if err != nil {
		return nil
	}
	result.MarketoInputOptionFilterColumns = filterColumns

	customObjectFields, err := newMarketoCustomObjectFields(ctx, marketoInputOption.CustomObjectFields)
	if err != nil {
		return nil
	}
	result.CustomObjectFields = customObjectFields

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, marketoInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func newMarketoColumns(
	ctx context.Context,
	columns *[]inputOptionEntities.MarketoColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: MarketoColumn{}.attrTypes(),
	}

	if columns == nil {
		return types.ListNull(objectType), nil
	}

	columnsResult := make([]MarketoColumn, 0, len(*columns))
	for _, col := range *columns {
		column := MarketoColumn{
			Name: types.StringValue(col.Name),
			Type: types.StringValue(col.Type),
		}
		columnsResult = append(columnsResult, column)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columnsResult)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert marketo columns to ListValue: %v", diags)
	}
	return listValue, nil
}

func newMarketoFilterColumns(
	ctx context.Context,
	filterColumns *[]inputOptionEntities.MarketoFilterColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: MarketoFilterColumn{}.attrTypes(),
	}

	if filterColumns == nil || len(*filterColumns) == 0 {
		return types.ListNull(objectType), nil
	}

	filterColumnsResult := make([]MarketoFilterColumn, 0, len(*filterColumns))
	for _, col := range *filterColumns {
		filterColumn := MarketoFilterColumn{
			Name: types.StringValue(col.Name),
		}
		filterColumnsResult = append(filterColumnsResult, filterColumn)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, filterColumnsResult)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert marketo filter columns to ListValue: %v", diags)
	}
	return listValue, nil
}

func newMarketoCustomObjectFields(
	ctx context.Context,
	fields *[]inputOptionEntities.MarketoCustomObjectField,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: MarketoCustomObjectField{}.attrTypes(),
	}

	if fields == nil || len(*fields) == 0 {
		return types.ListNull(objectType), nil
	}

	fieldsResult := make([]MarketoCustomObjectField, 0, len(*fields))
	for _, field := range *fields {
		customObjectField := MarketoCustomObjectField{
			Name: types.StringValue(field.Name),
		}
		fieldsResult = append(fieldsResult, customObjectField)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, fieldsResult)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert marketo custom object fields to ListValue: %v", diags)
	}
	return listValue, nil
}

func (marketoInputOption *MarketoInputOption) ToInput(ctx context.Context) *inputOptionParameters.MarketoInputOptionInput {
	if marketoInputOption == nil {
		return nil
	}

	input := &inputOptionParameters.MarketoInputOptionInput{
		MarketoConnectionID:         marketoInputOption.MarketoConnectionID.ValueInt64(),
		Target:                      marketoInputOption.Target.ValueString(),
		FromDate:                    marketoInputOption.FromDate.ValueStringPointer(),
		EndDate:                     marketoInputOption.EndDate.ValueStringPointer(),
		UseUpdatedAt:                model.NewNullableBool(marketoInputOption.UseUpdatedAt),
		PollingIntervalSecond:       model.NewNullableInt64(marketoInputOption.PollingIntervalSecond),
		BulkJobTimeoutSecond:        model.NewNullableInt64(marketoInputOption.BulkJobTimeoutSecond),
		CustomObjectAPIName:         marketoInputOption.CustomObjectAPIName.ValueStringPointer(),
		CustomObjectFilterType:      marketoInputOption.CustomObjectFilterType.ValueStringPointer(),
		CustomObjectFilterFromValue: marketoInputOption.CustomObjectFilterFromValue.ValueInt64Pointer(),
		CustomObjectFilterToValue:   marketoInputOption.CustomObjectFilterToValue.ValueInt64Pointer(),
		ListIDs:                     marketoInputOption.ListIDs.ValueStringPointer(),
		ProgramIDs:                  marketoInputOption.ProgramIDs.ValueStringPointer(),
		RootID:                      marketoInputOption.RootID.ValueInt64Pointer(),
		RootType:                    marketoInputOption.RootType.ValueStringPointer(),
		MaxDepth:                    model.NewNullableInt64(marketoInputOption.MaxDepth),
		Workspace:                   marketoInputOption.Workspace.ValueStringPointer(),
	}

	// Convert activity type IDs
	if !marketoInputOption.ActivityTypeIDs.IsNull() && !marketoInputOption.ActivityTypeIDs.IsUnknown() {
		var activityTypeIDs []int64
		diags := marketoInputOption.ActivityTypeIDs.ElementsAs(ctx, &activityTypeIDs, false)
		if !diags.HasError() && len(activityTypeIDs) > 0 {
			input.ActivityTypeIDs = &activityTypeIDs
		}
	}

	// Convert columns
	var columnValues []MarketoColumn
	if !marketoInputOption.InputOptionColumns.IsNull() && !marketoInputOption.InputOptionColumns.IsUnknown() {
		diags := marketoInputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
		if !diags.HasError() {
			input.MarketoInputOptionColumns = toMarketoColumnsInput(columnValues)
		}
	}

	// Convert filter columns
	var filterColumnValues []MarketoFilterColumn
	if !marketoInputOption.MarketoInputOptionFilterColumns.IsNull() && !marketoInputOption.MarketoInputOptionFilterColumns.IsUnknown() {
		diags := marketoInputOption.MarketoInputOptionFilterColumns.ElementsAs(ctx, &filterColumnValues, false)
		if !diags.HasError() {
			input.MarketoInputOptionFilterColumns = toMarketoFilterColumnsInput(filterColumnValues)
		}
	}

	// Convert custom object fields
	var customObjectFieldValues []MarketoCustomObjectField
	if !marketoInputOption.CustomObjectFields.IsNull() && !marketoInputOption.CustomObjectFields.IsUnknown() {
		diags := marketoInputOption.CustomObjectFields.ElementsAs(ctx, &customObjectFieldValues, false)
		if !diags.HasError() {
			input.CustomObjectFields = toMarketoCustomObjectFieldsInput(customObjectFieldValues)
		}
	}

	// Convert custom variable settings
	customVarSettings := common.ExtractCustomVariableSettings(ctx, marketoInputOption.CustomVariableSettings)
	input.CustomVariableSettings = model.ToCustomVariableSettingInputs(customVarSettings)

	return input
}

func (marketoInputOption *MarketoInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateMarketoInputOptionInput {
	if marketoInputOption == nil {
		return nil
	}

	input := &inputOptionParameters.UpdateMarketoInputOptionInput{
		MarketoConnectionID:         marketoInputOption.MarketoConnectionID.ValueInt64Pointer(),
		Target:                      marketoInputOption.Target.ValueStringPointer(),
		FromDate:                    marketoInputOption.FromDate.ValueStringPointer(),
		EndDate:                     marketoInputOption.EndDate.ValueStringPointer(),
		UseUpdatedAt:                model.NewNullableBool(marketoInputOption.UseUpdatedAt),
		PollingIntervalSecond:       model.NewNullableInt64(marketoInputOption.PollingIntervalSecond),
		BulkJobTimeoutSecond:        model.NewNullableInt64(marketoInputOption.BulkJobTimeoutSecond),
		CustomObjectAPIName:         marketoInputOption.CustomObjectAPIName.ValueStringPointer(),
		CustomObjectFilterType:      marketoInputOption.CustomObjectFilterType.ValueStringPointer(),
		CustomObjectFilterFromValue: marketoInputOption.CustomObjectFilterFromValue.ValueInt64Pointer(),
		CustomObjectFilterToValue:   marketoInputOption.CustomObjectFilterToValue.ValueInt64Pointer(),
		ListIDs:                     marketoInputOption.ListIDs.ValueStringPointer(),
		ProgramIDs:                  marketoInputOption.ProgramIDs.ValueStringPointer(),
		RootID:                      marketoInputOption.RootID.ValueInt64Pointer(),
		RootType:                    marketoInputOption.RootType.ValueStringPointer(),
		MaxDepth:                    model.NewNullableInt64(marketoInputOption.MaxDepth),
		Workspace:                   marketoInputOption.Workspace.ValueStringPointer(),
	}

	// Convert activity type IDs
	if !marketoInputOption.ActivityTypeIDs.IsNull() {
		var activityTypeIDs []int64
		if !marketoInputOption.ActivityTypeIDs.IsUnknown() {
			diags := marketoInputOption.ActivityTypeIDs.ElementsAs(ctx, &activityTypeIDs, false)
			if !diags.HasError() && len(activityTypeIDs) > 0 {
				input.ActivityTypeIDs = &activityTypeIDs
			}
		}
	}

	// Convert columns
	var columnValues []MarketoColumn
	if !marketoInputOption.InputOptionColumns.IsNull() {
		if !marketoInputOption.InputOptionColumns.IsUnknown() {
			diags := marketoInputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
			if !diags.HasError() {
				input.MarketoInputOptionColumns = toMarketoColumnsInput(columnValues)
			}
		} else {
			input.MarketoInputOptionColumns = &[]inputOptionParameters.MarketoColumn{}
		}
	}

	// Convert filter columns
	var filterColumnValues []MarketoFilterColumn
	if !marketoInputOption.MarketoInputOptionFilterColumns.IsNull() {
		if !marketoInputOption.MarketoInputOptionFilterColumns.IsUnknown() {
			diags := marketoInputOption.MarketoInputOptionFilterColumns.ElementsAs(ctx, &filterColumnValues, false)
			if !diags.HasError() {
				input.MarketoInputOptionFilterColumns = toMarketoFilterColumnsInput(filterColumnValues)
			}
		} else {
			input.MarketoInputOptionFilterColumns = &[]inputOptionParameters.MarketoFilterColumn{}
		}
	}

	// Convert custom object fields
	var customObjectFieldValues []MarketoCustomObjectField
	if !marketoInputOption.CustomObjectFields.IsNull() {
		if !marketoInputOption.CustomObjectFields.IsUnknown() {
			diags := marketoInputOption.CustomObjectFields.ElementsAs(ctx, &customObjectFieldValues, false)
			if !diags.HasError() {
				input.CustomObjectFields = toMarketoCustomObjectFieldsInput(customObjectFieldValues)
			}
		} else {
			input.CustomObjectFields = &[]inputOptionParameters.MarketoCustomObjectField{}
		}
	}

	// Convert custom variable settings
	customVarSettings := common.ExtractCustomVariableSettings(ctx, marketoInputOption.CustomVariableSettings)
	input.CustomVariableSettings = model.ToCustomVariableSettingInputs(customVarSettings)

	return input
}

func toMarketoColumnsInput(columns []MarketoColumn) *[]inputOptionParameters.MarketoColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.MarketoColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, inputOptionParameters.MarketoColumn{
			Name: column.Name.ValueString(),
			Type: column.Type.ValueString(),
		})
	}
	return &inputs
}

func toMarketoFilterColumnsInput(filterColumns []MarketoFilterColumn) *[]inputOptionParameters.MarketoFilterColumn {
	if filterColumns == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.MarketoFilterColumn, 0, len(filterColumns))
	for _, filterColumn := range filterColumns {
		inputs = append(inputs, inputOptionParameters.MarketoFilterColumn{
			Name: filterColumn.Name.ValueString(),
		})
	}
	return &inputs
}

func toMarketoCustomObjectFieldsInput(fields []MarketoCustomObjectField) *[]inputOptionParameters.MarketoCustomObjectField {
	if fields == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.MarketoCustomObjectField, 0, len(fields))
	for _, field := range fields {
		inputs = append(inputs, inputOptionParameters.MarketoCustomObjectField{
			Name: field.Name.ValueString(),
		})
	}
	return &inputs
}
