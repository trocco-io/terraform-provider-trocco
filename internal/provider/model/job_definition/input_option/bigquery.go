package input_options

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	parmas "terraform-provider-trocco/internal/client/parameter/job_definition"
	param "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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

	Columns                types.List `tfsdk:"columns"`
	CustomVariableSettings types.List `tfsdk:"custom_variable_settings"`
	Decoder                *Decoder   `tfsdk:"decoder"`
}

type BigqueryColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func (BigqueryColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   types.StringType,
		"type":   types.StringType,
		"format": types.StringType,
	}
}

func NewBigqueryInputOption(bigqueryInputOption *input_option.BigqueryInputOption) *BigqueryInputOption {
	if bigqueryInputOption == nil {
		return nil
	}

	ctx := context.Background()
	result := &BigqueryInputOption{
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
	}

	columns, err := newBigqueryColumns(ctx, bigqueryInputOption.Columns)
	if err != nil {
		return nil
	}
	result.Columns = columns

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, bigqueryInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func (bigqueryInputOption *BigqueryInputOption) ToInput() *param.BigqueryInputOptionInput {
	if bigqueryInputOption == nil {
		return nil
	}

	ctx := context.Background()

	var columnValues []BigqueryColumn
	if !bigqueryInputOption.Columns.IsNull() && !bigqueryInputOption.Columns.IsUnknown() {
		diags := bigqueryInputOption.Columns.ElementsAs(ctx, &columnValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, bigqueryInputOption.CustomVariableSettings)

	var decoder *parmas.DecoderInput
	if bigqueryInputOption.Decoder != nil {
		decoder = bigqueryInputOption.Decoder.ToDecoderInput()
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

		Columns:                toBigqueryColumnsInput(columnValues),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
		Decoder:                decoder,
	}
}

func (bigqueryInputOption *BigqueryInputOption) ToUpdateInput() *param.UpdateBigqueryInputOptionInput {
	if bigqueryInputOption == nil {
		return nil
	}

	ctx := context.Background()

	var columnValues []BigqueryColumn
	if !bigqueryInputOption.Columns.IsNull() {
		if !bigqueryInputOption.Columns.IsUnknown() {
			diags := bigqueryInputOption.Columns.ElementsAs(ctx, &columnValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnValues = []BigqueryColumn{}
		}
	} else {
		columnValues = []BigqueryColumn{}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, bigqueryInputOption.CustomVariableSettings)

	var decoder *parmas.DecoderInput
	if bigqueryInputOption.Decoder != nil {
		decoder = bigqueryInputOption.Decoder.ToDecoderInput()
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

		Columns:                toBigqueryColumnsInput(columnValues),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
		Decoder:                decoder,
	}
}

func newBigqueryColumns(
	ctx context.Context,
	bigqueryColumns []input_option.BigqueryColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: BigqueryColumn{}.attrTypes(),
	}

	if bigqueryColumns == nil {
		return types.ListNull(objectType), nil
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

	listValue, diags := types.ListValueFrom(ctx, objectType, columns)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert bigquery columns to ListValue: %v", diags)
	}
	return listValue, nil
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
