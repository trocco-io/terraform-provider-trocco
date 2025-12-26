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

type HubspotInputOption struct {
	HubspotConnectionID       types.Int64  `tfsdk:"hubspot_connection_id"`
	Target                    types.String `tfsdk:"target"`
	FromObjectType            types.String `tfsdk:"from_object_type"`
	ToObjectType              types.String `tfsdk:"to_object_type"`
	ObjectType                types.String `tfsdk:"object_type"`
	IncrementalLoadingEnabled types.Bool   `tfsdk:"incremental_loading_enabled"`
	LastRecordTime            types.String `tfsdk:"last_record_time"`
	EmailEventType            types.String `tfsdk:"email_event_type"`
	StartTimestamp            types.String `tfsdk:"start_timestamp"`
	EndTimestamp              types.String `tfsdk:"end_timestamp"`
	InputOptionColumns        types.List   `tfsdk:"input_option_columns"`
	CustomVariableSettings    types.List   `tfsdk:"custom_variable_settings"`
}

type HubspotColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func (HubspotColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   types.StringType,
		"type":   types.StringType,
		"format": types.StringType,
	}
}

func NewHubspotInputOption(ctx context.Context, hubspotInputOption *inputOptionEntities.HubspotInputOption) *HubspotInputOption {
	if hubspotInputOption == nil {
		return nil
	}

	result := &HubspotInputOption{
		HubspotConnectionID:       types.Int64Value(hubspotInputOption.HubspotConnectionID),
		Target:                    types.StringValue(hubspotInputOption.Target),
		FromObjectType:            types.StringPointerValue(hubspotInputOption.FromObjectType),
		ToObjectType:              types.StringPointerValue(hubspotInputOption.ToObjectType),
		ObjectType:                types.StringPointerValue(hubspotInputOption.ObjectType),
		IncrementalLoadingEnabled: types.BoolPointerValue(hubspotInputOption.IncrementalLoadingEnabled),
		LastRecordTime:            types.StringPointerValue(hubspotInputOption.LastRecordTime),
		EmailEventType:            types.StringPointerValue(hubspotInputOption.EmailEventType),
		StartTimestamp:            types.StringPointerValue(hubspotInputOption.StartTimestamp),
		EndTimestamp:              types.StringPointerValue(hubspotInputOption.EndTimestamp),
	}

	columns, err := newHubspotColumns(ctx, hubspotInputOption.InputOptionColumns)
	if err != nil {
		return nil
	}
	result.InputOptionColumns = columns

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, hubspotInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func newHubspotColumns(
	ctx context.Context,
	hubspotColumns []inputOptionEntities.HubspotColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: HubspotColumn{}.attrTypes(),
	}

	if hubspotColumns == nil {
		return types.ListNull(objectType), nil
	}

	columns := make([]HubspotColumn, 0, len(hubspotColumns))
	for _, input := range hubspotColumns {
		column := HubspotColumn{
			Name:   types.StringValue(input.Name),
			Type:   types.StringValue(input.Type),
			Format: types.StringPointerValue(input.Format),
		}
		columns = append(columns, column)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columns)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert hubspot columns to ListValue: %v", diags)
	}
	return listValue, nil
}

func (hubspotInputOption *HubspotInputOption) ToInput(ctx context.Context) *inputOptionParameters.HubspotInputOptionInput {
	if hubspotInputOption == nil {
		return nil
	}

	var columnValues []HubspotColumn
	if !hubspotInputOption.InputOptionColumns.IsNull() && !hubspotInputOption.InputOptionColumns.IsUnknown() {
		diags := hubspotInputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, hubspotInputOption.CustomVariableSettings)

	return &inputOptionParameters.HubspotInputOptionInput{
		HubspotConnectionID:       hubspotInputOption.HubspotConnectionID.ValueInt64(),
		Target:                    hubspotInputOption.Target.ValueString(),
		FromObjectType:            hubspotInputOption.FromObjectType.ValueStringPointer(),
		ToObjectType:              hubspotInputOption.ToObjectType.ValueStringPointer(),
		ObjectType:                hubspotInputOption.ObjectType.ValueStringPointer(),
		IncrementalLoadingEnabled: hubspotInputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		LastRecordTime:            hubspotInputOption.LastRecordTime.ValueStringPointer(),
		EmailEventType:            hubspotInputOption.EmailEventType.ValueStringPointer(),
		StartTimestamp:            hubspotInputOption.StartTimestamp.ValueStringPointer(),
		EndTimestamp:              hubspotInputOption.EndTimestamp.ValueStringPointer(),
		InputOptionColumns:        toHubspotColumnsInput(columnValues),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (hubspotInputOption *HubspotInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateHubspotInputOptionInput {
	if hubspotInputOption == nil {
		return nil
	}

	var columnValues []HubspotColumn
	if !hubspotInputOption.InputOptionColumns.IsNull() {
		if !hubspotInputOption.InputOptionColumns.IsUnknown() {
			diags := hubspotInputOption.InputOptionColumns.ElementsAs(ctx, &columnValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnValues = []HubspotColumn{}
		}
	} else {
		columnValues = []HubspotColumn{}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, hubspotInputOption.CustomVariableSettings)

	return &inputOptionParameters.UpdateHubspotInputOptionInput{
		HubspotConnectionID:       hubspotInputOption.HubspotConnectionID.ValueInt64Pointer(),
		Target:                    hubspotInputOption.Target.ValueStringPointer(),
		FromObjectType:            hubspotInputOption.FromObjectType.ValueStringPointer(),
		ToObjectType:              hubspotInputOption.ToObjectType.ValueStringPointer(),
		ObjectType:                hubspotInputOption.ObjectType.ValueStringPointer(),
		IncrementalLoadingEnabled: hubspotInputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		LastRecordTime:            hubspotInputOption.LastRecordTime.ValueStringPointer(),
		EmailEventType:            hubspotInputOption.EmailEventType.ValueStringPointer(),
		StartTimestamp:            hubspotInputOption.StartTimestamp.ValueStringPointer(),
		EndTimestamp:              hubspotInputOption.EndTimestamp.ValueStringPointer(),
		InputOptionColumns:        toHubspotColumnsInput(columnValues),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func toHubspotColumnsInput(columns []HubspotColumn) []inputOptionParameters.HubspotColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.HubspotColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, inputOptionParameters.HubspotColumn{
			Name:   column.Name.ValueString(),
			Type:   column.Type.ValueString(),
			Format: column.Format.ValueStringPointer(),
		})
	}
	return inputs
}
