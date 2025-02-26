package input_options

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	param "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
)

type SalesforceInputOption struct {
	Object                          types.String                   `tfsdk:"object"`
	ObjectAcquisitionMethod         types.String                   `tfsdk:"object_acquisition_method"`
	IsConvertTypeCustomColumns      types.Bool                     `tfsdk:"is_convert_type_custom_columns"`
	IncludeDeletedOrArchivedRecords types.Bool                     `tfsdk:"include_deleted_or_archived_records"`
	ApiVersion                      types.String                   `tfsdk:"api_version"`
	Soql                            types.String                   `tfsdk:"soql"`
	SalesforceConnectionID          types.Int64                    `tfsdk:"salesforce_connection_id"`
	Columns                         []SalesforceColumn             `tfsdk:"columns"`
	CustomVariableSettings          *[]model.CustomVariableSetting `tfsdk:"custom_variable_settings"`
}

type SalesforceColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func NewSalesforceInputOption(salesforceInputOption *input_option.SalesforceInputOption) *SalesforceInputOption {
	if salesforceInputOption == nil {
		return nil
	}

	return &SalesforceInputOption{
		Object:                          types.StringValue(salesforceInputOption.Object),
		ObjectAcquisitionMethod:         types.StringValue(salesforceInputOption.ObjectAcquisitionMethod),
		IsConvertTypeCustomColumns:      types.BoolValue(salesforceInputOption.IsConvertTypeCustomColumns),
		IncludeDeletedOrArchivedRecords: types.BoolValue(salesforceInputOption.IncludeDeletedOrArchivedRecords),
		ApiVersion:                      types.StringValue(salesforceInputOption.ApiVersion),
		Soql:                            types.StringPointerValue(salesforceInputOption.Soql),
		SalesforceConnectionID:          types.Int64Value(salesforceInputOption.SalesforceConnectionID),
		Columns:                         newColumns(salesforceInputOption.Columns),
		CustomVariableSettings:          model.NewCustomVariableSettings(salesforceInputOption.CustomVariableSettings),
	}
}

func newColumns(salesforceColumns []input_option.SalesforceColumn) []SalesforceColumn {
	if salesforceColumns == nil {
		return nil
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
	return columns
}

func (salesforceInputOption *SalesforceInputOption) ToInput() *param.SalesforceInputOptionInput {
	if salesforceInputOption == nil {
		return nil
	}

	return &param.SalesforceInputOptionInput{
		Object:                          salesforceInputOption.Object.ValueString(),
		ObjectAcquisitionMethod:         model.NewNullableString(salesforceInputOption.ObjectAcquisitionMethod),
		IsConvertTypeCustomColumns:      model.NewNullableBool(salesforceInputOption.IsConvertTypeCustomColumns),
		IncludeDeletedOrArchivedRecords: model.NewNullableBool(salesforceInputOption.IncludeDeletedOrArchivedRecords),
		ApiVersion:                      model.NewNullableString(salesforceInputOption.ApiVersion),
		Soql:                            model.NewNullableString(salesforceInputOption.Soql),
		SalesforceConnectionID:          salesforceInputOption.SalesforceConnectionID.ValueInt64(),
		Columns:                         toSalesforceColumnsInput(salesforceInputOption.Columns),
		CustomVariableSettings:          model.ToCustomVariableSettingInputs(salesforceInputOption.CustomVariableSettings),
	}
}

func (salesforceInputOption *SalesforceInputOption) ToUpdateInput() *param.UpdateSalesforceInputOptionInput {
	if salesforceInputOption == nil {
		return nil
	}

	return &param.UpdateSalesforceInputOptionInput{
		Object:                          salesforceInputOption.Object.ValueStringPointer(),
		ObjectAcquisitionMethod:         model.NewNullableString(salesforceInputOption.ObjectAcquisitionMethod),
		IsConvertTypeCustomColumns:      model.NewNullableBool(salesforceInputOption.IsConvertTypeCustomColumns),
		IncludeDeletedOrArchivedRecords: model.NewNullableBool(salesforceInputOption.IncludeDeletedOrArchivedRecords),
		ApiVersion:                      model.NewNullableString(salesforceInputOption.ApiVersion),
		Soql:                            model.NewNullableString(salesforceInputOption.Soql),
		SalesforceConnectionID:          salesforceInputOption.SalesforceConnectionID.ValueInt64Pointer(),
		Columns:                         toSalesforceColumnsInput(salesforceInputOption.Columns),
		CustomVariableSettings:          model.ToCustomVariableSettingInputs(salesforceInputOption.CustomVariableSettings),
	}
}

func toSalesforceColumnsInput(columns []SalesforceColumn) []param.SalesforceColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]param.SalesforceColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, param.SalesforceColumn{
			Name:   column.Name.ValueString(),
			Type:   column.Type.ValueString(),
			Format: column.Format.ValueStringPointer(),
		})
	}
	return inputs
}
