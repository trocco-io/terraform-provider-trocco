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

type SalesforceInputOption struct {
	Object                          types.String `tfsdk:"object"`
	ObjectAcquisitionMethod         types.String `tfsdk:"object_acquisition_method"`
	IsConvertTypeCustomColumns      types.Bool   `tfsdk:"is_convert_type_custom_columns"`
	IncludeDeletedOrArchivedRecords types.Bool   `tfsdk:"include_deleted_or_archived_records"`
	ApiVersion                      types.String `tfsdk:"api_version"`
	Soql                            types.String `tfsdk:"soql"`
	SalesforceConnectionID          types.Int64  `tfsdk:"salesforce_connection_id"`
	Columns                         types.List   `tfsdk:"columns"`
	CustomVariableSettings          types.List   `tfsdk:"custom_variable_settings"`
}

type SalesforceColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func (SalesforceColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   types.StringType,
		"type":   types.StringType,
		"format": types.StringType,
	}
}

func NewSalesforceInputOption(ctx context.Context, salesforceInputOption *inputOptionEntities.SalesforceInputOption) *SalesforceInputOption {
	if salesforceInputOption == nil {
		return nil
	}

	result := &SalesforceInputOption{
		Object:                          types.StringValue(salesforceInputOption.Object),
		ObjectAcquisitionMethod:         types.StringValue(salesforceInputOption.ObjectAcquisitionMethod),
		IsConvertTypeCustomColumns:      types.BoolValue(salesforceInputOption.IsConvertTypeCustomColumns),
		IncludeDeletedOrArchivedRecords: types.BoolValue(salesforceInputOption.IncludeDeletedOrArchivedRecords),
		ApiVersion:                      types.StringValue(salesforceInputOption.ApiVersion),
		Soql:                            types.StringPointerValue(salesforceInputOption.Soql),
		SalesforceConnectionID:          types.Int64Value(salesforceInputOption.SalesforceConnectionID),
	}

	columns, err := newColumns(ctx, salesforceInputOption.Columns)
	if err != nil {
		return nil
	}
	result.Columns = columns

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, salesforceInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func newColumns(
	ctx context.Context,
	salesforceColumns []inputOptionEntities.SalesforceColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: SalesforceColumn{}.attrTypes(),
	}

	if salesforceColumns == nil {
		return types.ListNull(objectType), nil
	}

	columns := make([]SalesforceColumn, 0, len(salesforceColumns))
	for _, input := range salesforceColumns {
		column := SalesforceColumn{
			Name:   types.StringValue(input.Name),
			Type:   types.StringValue(input.Type),
			Format: types.StringPointerValue(input.Format),
		}
		columns = append(columns, column)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columns)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert salesforce columns to ListValue: %v", diags)
	}
	return listValue, nil
}

func (salesforceInputOption *SalesforceInputOption) ToInput(ctx context.Context) *inputOptionParameters.SalesforceInputOptionInput {
	if salesforceInputOption == nil {
		return nil
	}

	var columnValues []SalesforceColumn
	if !salesforceInputOption.Columns.IsNull() && !salesforceInputOption.Columns.IsUnknown() {
		diags := salesforceInputOption.Columns.ElementsAs(ctx, &columnValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, salesforceInputOption.CustomVariableSettings)

	return &inputOptionParameters.SalesforceInputOptionInput{
		Object:                          salesforceInputOption.Object.ValueString(),
		ObjectAcquisitionMethod:         model.NewNullableString(salesforceInputOption.ObjectAcquisitionMethod),
		IsConvertTypeCustomColumns:      model.NewNullableBool(salesforceInputOption.IsConvertTypeCustomColumns),
		IncludeDeletedOrArchivedRecords: model.NewNullableBool(salesforceInputOption.IncludeDeletedOrArchivedRecords),
		ApiVersion:                      model.NewNullableString(salesforceInputOption.ApiVersion),
		Soql:                            model.NewNullableString(salesforceInputOption.Soql),
		SalesforceConnectionID:          salesforceInputOption.SalesforceConnectionID.ValueInt64(),
		Columns:                         toSalesforceColumnsInput(columnValues),
		CustomVariableSettings:          model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (salesforceInputOption *SalesforceInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateSalesforceInputOptionInput {
	if salesforceInputOption == nil {
		return nil
	}

	var columnValues []SalesforceColumn
	if !salesforceInputOption.Columns.IsNull() {
		if !salesforceInputOption.Columns.IsUnknown() {
			diags := salesforceInputOption.Columns.ElementsAs(ctx, &columnValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnValues = []SalesforceColumn{}
		}
	} else {
		columnValues = []SalesforceColumn{}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, salesforceInputOption.CustomVariableSettings)

	return &inputOptionParameters.UpdateSalesforceInputOptionInput{
		Object:                          salesforceInputOption.Object.ValueStringPointer(),
		ObjectAcquisitionMethod:         model.NewNullableString(salesforceInputOption.ObjectAcquisitionMethod),
		IsConvertTypeCustomColumns:      model.NewNullableBool(salesforceInputOption.IsConvertTypeCustomColumns),
		IncludeDeletedOrArchivedRecords: model.NewNullableBool(salesforceInputOption.IncludeDeletedOrArchivedRecords),
		ApiVersion:                      model.NewNullableString(salesforceInputOption.ApiVersion),
		Soql:                            model.NewNullableString(salesforceInputOption.Soql),
		SalesforceConnectionID:          salesforceInputOption.SalesforceConnectionID.ValueInt64Pointer(),
		Columns:                         toSalesforceColumnsInput(columnValues),
		CustomVariableSettings:          model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func toSalesforceColumnsInput(columns []SalesforceColumn) []inputOptionParameters.SalesforceColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.SalesforceColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, inputOptionParameters.SalesforceColumn{
			Name:   column.Name.ValueString(),
			Type:   column.Type.ValueString(),
			Format: column.Format.ValueStringPointer(),
		})
	}
	return inputs
}
