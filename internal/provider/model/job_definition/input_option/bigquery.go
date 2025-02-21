package input_options

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	param "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BigqueryInputOption struct {
	BigqueryConnectionID  types.Int64  `tfsdk:"bigquery_connection_id"`
	GcsUri                types.String `tfsdk:"gcs_uri"`
	GcsUriFormat          types.String `tfsdk:"gcs_uri_format"`
	Query                 types.String `tfsdk:"query"`
	TempDataset           types.String `tfsdk:"temp_dataset"`
	IsStandardSQL         types.Bool   `tfsdk:"is_standard_sql"`
	CleanupGcsFiles       types.Bool   `tfsdk:"cleanup_gcs_files"`
	FileFormat            types.String `tfsdk:"file_format"`
	Location              types.String `tfsdk:"location"`
	Cache                 types.Bool   `tfsdk:"cache"`
	BigqueryJobWaitSecond types.Int64  `tfsdk:"bigquery_job_wait_second"`

	Columns                []BigqueryColumn               `tfsdk:"columns"`
	CustomVariableSettings *[]model.CustomVariableSetting `tfsdk:"custom_variable_settings"`
	Decoder                *Decoder                       `tfsdk:"decoder"`
}

type BigqueryColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func NewBigqueryInputOption(bigqueryInputOption *input_option.BigqueryInputOption) *BigqueryInputOption {
	if bigqueryInputOption == nil {
		return nil
	}
	return &BigqueryInputOption{
		BigqueryConnectionID:  types.Int64Value(bigqueryInputOption.BigqueryConnectionID),
		GcsUri:                types.StringValue(bigqueryInputOption.GcsUri),
		GcsUriFormat:          types.StringPointerValue(bigqueryInputOption.GcsUriFormat),
		Query:                 types.StringValue(bigqueryInputOption.Query),
		TempDataset:           types.StringValue(bigqueryInputOption.TempDataset),
		IsStandardSQL:         types.BoolPointerValue(bigqueryInputOption.IsStandardSQL),
		CleanupGcsFiles:       types.BoolPointerValue(bigqueryInputOption.CleanupGcsFiles),
		FileFormat:            types.StringPointerValue(bigqueryInputOption.FileFormat),
		Location:              types.StringPointerValue(bigqueryInputOption.Location),
		Cache:                 types.BoolPointerValue(bigqueryInputOption.Cache),
		BigqueryJobWaitSecond: types.Int64PointerValue(bigqueryInputOption.BigqueryJobWaitSecond),

		Columns:                newBigqueryColumns(bigqueryInputOption.Columns),
		CustomVariableSettings: model.NewCustomVariableSettings(bigqueryInputOption.CustomVariableSettings),
	}
}

func (bigqueryInputOption *BigqueryInputOption) ToInput() *param.BigqueryInputOptionInput {
	if bigqueryInputOption == nil {
		return nil
	}

	return &param.BigqueryInputOptionInput{
		BigqueryConnectionID:  bigqueryInputOption.BigqueryConnectionID.ValueInt64(),
		GcsUri:                bigqueryInputOption.GcsUri.ValueString(),
		GcsUriFormat:          model.NewNullableString(bigqueryInputOption.GcsUriFormat),
		Query:                 bigqueryInputOption.Query.ValueString(),
		TempDataset:           bigqueryInputOption.TempDataset.ValueString(),
		IsStandardSQL:         bigqueryInputOption.IsStandardSQL.ValueBoolPointer(),
		CleanupGcsFiles:       bigqueryInputOption.CleanupGcsFiles.ValueBoolPointer(),
		FileFormat:            model.NewNullableString(bigqueryInputOption.FileFormat),
		Location:              model.NewNullableString(bigqueryInputOption.Location),
		Cache:                 bigqueryInputOption.Cache.ValueBoolPointer(),
		BigqueryJobWaitSecond: bigqueryInputOption.BigqueryJobWaitSecond.ValueInt64Pointer(),

		Columns:                toBigqueryColumnsInput(bigqueryInputOption.Columns),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(bigqueryInputOption.CustomVariableSettings),
		Decoder:                bigqueryInputOption.Decoder.ToDecoderInput(),
	}
}

func (bigqueryInputOption *BigqueryInputOption) ToUpdateInput() *param.UpdateBigqueryInputOptionInput {
	if bigqueryInputOption == nil {
		return nil
	}

	return &param.UpdateBigqueryInputOptionInput{
		BigqueryConnectionID:  bigqueryInputOption.BigqueryConnectionID.ValueInt64Pointer(),
		GcsUri:                model.NewNullableString(bigqueryInputOption.GcsUri),
		GcsUriFormat:          model.NewNullableString(bigqueryInputOption.GcsUriFormat),
		Query:                 model.NewNullableString(bigqueryInputOption.Query),
		TempDataset:           model.NewNullableString(bigqueryInputOption.TempDataset),
		IsStandardSQL:         bigqueryInputOption.IsStandardSQL.ValueBoolPointer(),
		CleanupGcsFiles:       bigqueryInputOption.CleanupGcsFiles.ValueBoolPointer(),
		FileFormat:            model.NewNullableString(bigqueryInputOption.FileFormat),
		Location:              model.NewNullableString(bigqueryInputOption.Location),
		Cache:                 bigqueryInputOption.Cache.ValueBoolPointer(),
		BigqueryJobWaitSecond: bigqueryInputOption.BigqueryJobWaitSecond.ValueInt64Pointer(),

		Columns:                toBigqueryColumnsInput(bigqueryInputOption.Columns),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(bigqueryInputOption.CustomVariableSettings),
		Decoder:                bigqueryInputOption.Decoder.ToDecoderInput(),
	}
}

func newBigqueryColumns(bigqueryColumns []input_option.BigqueryColumn) []BigqueryColumn {
	if bigqueryColumns == nil {
		return nil
	}
	columns := make([]BigqueryColumn, 0, len(bigqueryColumns))
	for _, input := range bigqueryColumns {
		column := BigqueryColumn{
			Name:   types.StringValue(input.Name),
			Type:   types.StringValue(input.Type),
			Format: types.StringPointerValue(input.Format),
		}
		columns = append(columns, column)
	}
	return columns
}

func toBigqueryColumnsInput(columns []BigqueryColumn) []param.BigqueryColumn {
	if columns == nil {
		return nil
	}

	inputs := make([]param.BigqueryColumn, 0, len(columns))
	for _, column := range columns {
		inputs = append(inputs, param.BigqueryColumn{
			Name:   column.Name.ValueString(),
			Type:   column.Type.ValueString(),
			Format: column.Format.ValueStringPointer(),
		})
	}
	return inputs
}
