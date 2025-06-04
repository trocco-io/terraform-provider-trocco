package output_options

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	output_options2 "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BigQueryOutputOption struct {
	Dataset                              types.String                   `tfsdk:"dataset"`
	Table                                types.String                   `tfsdk:"table"`
	AutoCreateDataset                    types.Bool                     `tfsdk:"auto_create_dataset"`
	OpenTimeoutSec                       types.Int64                    `tfsdk:"open_timeout_sec"`
	TimeoutSec                           types.Int64                    `tfsdk:"timeout_sec"`
	SendTimeoutSec                       types.Int64                    `tfsdk:"send_timeout_sec"`
	ReadTimeoutSec                       types.Int64                    `tfsdk:"read_timeout_sec"`
	Retries                              types.Int64                    `tfsdk:"retries"`
	Mode                                 types.String                   `tfsdk:"mode"`
	PartitioningType                     types.String                   `tfsdk:"partitioning_type"`
	TimePartitioningType                 types.String                   `tfsdk:"time_partitioning_type"`
	TimePartitioningField                types.String                   `tfsdk:"time_partitioning_field"`
	TimePartitioningExpirationMs         types.Int64                    `tfsdk:"time_partitioning_expiration_ms"`
	Location                             types.String                   `tfsdk:"location"`
	TemplateTable                        types.String                   `tfsdk:"template_table"`
	BigQueryConnectionID                 types.Int64                    `tfsdk:"bigquery_connection_id"`
	CustomVariableSettings               *[]model.CustomVariableSetting `tfsdk:"custom_variable_settings"`
	BigQueryOutputOptionColumnOptions    types.List                     `tfsdk:"bigquery_output_option_column_options"`
	BigQueryOutputOptionClusteringFields types.Set                      `tfsdk:"bigquery_output_option_clustering_fields"`
	BigQueryOutputOptionMergeKeys        types.Set                      `tfsdk:"bigquery_output_option_merge_keys"`
}

type bigQueryOutputOptionColumnOption struct {
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	Mode            types.String `tfsdk:"mode"`
	TimestampFormat types.String `tfsdk:"timestamp_format"`
	Timezone        types.String `tfsdk:"timezone"`
	Description     types.String `tfsdk:"description"`
}

func NewBigQueryOutputOption(bigQueryOutputOption *output_option.BigQueryOutputOption) *BigQueryOutputOption {
	if bigQueryOutputOption == nil {
		return nil
	}

	ctx := context.Background()

	result := &BigQueryOutputOption{
		CustomVariableSettings:       model.NewCustomVariableSettings(bigQueryOutputOption.CustomVariableSettings),
		Dataset:                      types.StringValue(bigQueryOutputOption.Dataset),
		Table:                        types.StringValue(bigQueryOutputOption.Table),
		AutoCreateDataset:            types.BoolValue(bigQueryOutputOption.AutoCreateDataset),
		OpenTimeoutSec:               types.Int64Value(bigQueryOutputOption.OpenTimeoutSec),
		TimeoutSec:                   types.Int64Value(bigQueryOutputOption.TimeoutSec),
		SendTimeoutSec:               types.Int64Value(bigQueryOutputOption.SendTimeoutSec),
		ReadTimeoutSec:               types.Int64Value(bigQueryOutputOption.ReadTimeoutSec),
		Retries:                      types.Int64Value(bigQueryOutputOption.Retries),
		Mode:                         types.StringValue(bigQueryOutputOption.Mode),
		PartitioningType:             types.StringPointerValue(bigQueryOutputOption.PartitioningType),
		TimePartitioningType:         types.StringPointerValue(bigQueryOutputOption.TimePartitioningType),
		TimePartitioningField:        types.StringPointerValue(bigQueryOutputOption.TimePartitioningField),
		TimePartitioningExpirationMs: types.Int64PointerValue(bigQueryOutputOption.TimePartitioningExpirationMs),
		Location:                     types.StringPointerValue(bigQueryOutputOption.Location),
		TemplateTable:                types.StringPointerValue(bigQueryOutputOption.TemplateTable),
		BigQueryConnectionID:         types.Int64Value(bigQueryOutputOption.BigQueryConnectionID),
	}

	BigQueryOutputOptionColumnOptions, err := newBigqueryOutputOptionColumnOptions(ctx, *bigQueryOutputOption.BigQueryOutputOptionColumnOptions)
	if err != nil {
		return nil
	}
	result.BigQueryOutputOptionColumnOptions = BigQueryOutputOptionColumnOptions

	var clusteringFields []string
	if bigQueryOutputOption.BigQueryOutputOptionClusteringFields != nil {
		clusteringFields = *bigQueryOutputOption.BigQueryOutputOptionClusteringFields
	}
	BigQueryOutputOptionClusteringFields, err := newBigQueryOutputOptionClusteringFields(ctx, clusteringFields)
	if err != nil {
		return nil
	}
	result.BigQueryOutputOptionClusteringFields = BigQueryOutputOptionClusteringFields

	var mergeKeys []string
	if bigQueryOutputOption.BigQueryOutputOptionMergeKeys != nil {
		mergeKeys = *bigQueryOutputOption.BigQueryOutputOptionMergeKeys
	}
	BigQueryOutputOptionMergeKeys, err := newBigQueryOutputOptionMergeKeys(ctx, mergeKeys)
	if err != nil {
		return nil
	}
	result.BigQueryOutputOptionMergeKeys = BigQueryOutputOptionMergeKeys
	return result
}

func newBigQueryOutputOptionMergeKeys(ctx context.Context, mergeKeys []string) (types.Set, error) {
	if mergeKeys == nil {
		return types.SetNull(types.StringType), nil
	}

	values := make([]types.String, len(mergeKeys))
	for i, v := range mergeKeys {
		values[i] = types.StringValue(v)
	}

	setValue, diags := types.SetValueFrom(ctx, types.StringType, values)
	if diags.HasError() {
		return types.SetNull(types.StringType), fmt.Errorf("failed to convert mergeKeys to SetValue: %v", diags)
	}

	return setValue, nil
}

func newBigQueryOutputOptionClusteringFields(ctx context.Context, fields []string) (types.Set, error) {
	if fields == nil {
		return types.SetNull(types.StringType), nil
	}

	values := make([]types.String, len(fields))
	for i, v := range fields {
		values[i] = types.StringValue(v)
	}

	setValue, diags := types.SetValueFrom(ctx, types.StringType, values)
	if diags.HasError() {
		return types.SetNull(types.StringType), fmt.Errorf("failed to convert to SetValue: %v", diags)
	}

	return setValue, nil
}

func newBigqueryOutputOptionColumnOptions(
	ctx context.Context,
	bigQueryOutputOptionColumnOptions []output_option.BigQueryOutputOptionColumnOption,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: bigQueryOutputOptionColumnOption{}.attrTypes(),
	}

	if bigQueryOutputOptionColumnOptions == nil {
		return types.ListNull(objectType), nil
	}

	columnOptions := make([]bigQueryOutputOptionColumnOption, 0, len(bigQueryOutputOptionColumnOptions))
	for _, input := range bigQueryOutputOptionColumnOptions {
		columnOption := bigQueryOutputOptionColumnOption{
			Name:            types.StringValue(input.Name),
			Type:            types.StringValue(input.Type),
			Mode:            types.StringValue(input.Mode),
			TimestampFormat: types.StringPointerValue(input.TimestampFormat),
			Timezone:        types.StringPointerValue(input.Timezone),
			Description:     types.StringPointerValue(input.Description),
		}
		columnOptions = append(columnOptions, columnOption)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columnOptions)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert column options to ListValue: %v", diags)
	}
	return listValue, nil
}

func (bigQueryOutputOptionColumnOption) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":             types.StringType,
		"type":             types.StringType,
		"mode":             types.StringType,
		"timestamp_format": types.StringType,
		"timezone":         types.StringType,
		"description":      types.StringType,
	}
}

func (bigqueryOutputOption *BigQueryOutputOption) ToInput() *output_options2.BigQueryOutputOptionInput {
	if bigqueryOutputOption == nil {
		return nil
	}

	ctx := context.Background()

	var clusteringFields []string
	if !bigqueryOutputOption.BigQueryOutputOptionClusteringFields.IsNull() &&
		!bigqueryOutputOption.BigQueryOutputOptionClusteringFields.IsUnknown() {
		var clusteringFieldValues []types.String
		diags := bigqueryOutputOption.BigQueryOutputOptionClusteringFields.ElementsAs(ctx, &clusteringFieldValues, false)
		if diags.HasError() {
			return nil
		}

		clusteringFields = make([]string, 0, len(clusteringFieldValues))
		for _, input := range clusteringFieldValues {
			clusteringFields = append(clusteringFields, input.ValueString())
		}
	}

	var mergeKeys []string
	if !bigqueryOutputOption.BigQueryOutputOptionMergeKeys.IsNull() &&
		!bigqueryOutputOption.BigQueryOutputOptionMergeKeys.IsUnknown() {
		var mergeKeyValues []types.String
		diags := bigqueryOutputOption.BigQueryOutputOptionMergeKeys.ElementsAs(ctx, &mergeKeyValues, false)
		if diags.HasError() {
			return nil
		}

		mergeKeys = make([]string, 0, len(mergeKeyValues))
		for _, input := range mergeKeyValues {
			mergeKeys = append(mergeKeys, input.ValueString())
		}
	}

	var columnOptionValues []bigQueryOutputOptionColumnOption
	if !bigqueryOutputOption.BigQueryOutputOptionColumnOptions.IsNull() && !bigqueryOutputOption.BigQueryOutputOptionColumnOptions.IsUnknown() {
		diags := bigqueryOutputOption.BigQueryOutputOptionColumnOptions.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}
	}
	columnOptions := toInputBigqueryOutputOptionColumnOptions(&columnOptionValues)

	return &output_options2.BigQueryOutputOptionInput{
		Dataset:                              bigqueryOutputOption.Dataset.ValueString(),
		Table:                                bigqueryOutputOption.Table.ValueString(),
		AutoCreateDataset:                    bigqueryOutputOption.AutoCreateDataset.ValueBool(),
		OpenTimeoutSec:                       bigqueryOutputOption.OpenTimeoutSec.ValueInt64(),
		TimeoutSec:                           bigqueryOutputOption.TimeoutSec.ValueInt64(),
		SendTimeoutSec:                       bigqueryOutputOption.SendTimeoutSec.ValueInt64(),
		ReadTimeoutSec:                       bigqueryOutputOption.ReadTimeoutSec.ValueInt64(),
		Retries:                              bigqueryOutputOption.Retries.ValueInt64(),
		Mode:                                 bigqueryOutputOption.Mode.ValueString(),
		PartitioningType:                     model.NewNullableString(bigqueryOutputOption.PartitioningType),
		TimePartitioningType:                 model.NewNullableString(bigqueryOutputOption.TimePartitioningType),
		TimePartitioningField:                model.NewNullableString(bigqueryOutputOption.TimePartitioningField),
		TimePartitioningExpirationMs:         model.NewNullableInt64(bigqueryOutputOption.TimePartitioningExpirationMs),
		Location:                             bigqueryOutputOption.Location.ValueString(),
		TemplateTable:                        model.NewNullableString(bigqueryOutputOption.TemplateTable),
		BigQueryConnectionID:                 bigqueryOutputOption.BigQueryConnectionID.ValueInt64(),
		CustomVariableSettings:               model.ToCustomVariableSettingInputs(bigqueryOutputOption.CustomVariableSettings),
		BigQueryOutputOptionColumnOptions:    columnOptions,
		BigQueryOutputOptionClusteringFields: clusteringFields,
		BigQueryOutputOptionMergeKeys:        mergeKeys,
	}
}

func (bigqueryOutputOption *BigQueryOutputOption) ToUpdateInput() *output_options2.UpdateBigQueryOutputOptionInput {
	if bigqueryOutputOption == nil {
		return nil
	}

	ctx := context.Background()

	var clusteringFields []string
	if !bigqueryOutputOption.BigQueryOutputOptionClusteringFields.IsNull() {
		var clusteringFieldValues []types.String
		if !bigqueryOutputOption.BigQueryOutputOptionClusteringFields.IsUnknown() {
			diags := bigqueryOutputOption.BigQueryOutputOptionClusteringFields.ElementsAs(ctx, &clusteringFieldValues, false)
			if diags.HasError() {
				return nil
			}
		}
		clusteringFields = make([]string, 0, len(clusteringFieldValues))
		for _, input := range clusteringFieldValues {
			clusteringFields = append(clusteringFields, input.ValueString())
		}
	}

	var mergeKeys []string
	if !bigqueryOutputOption.BigQueryOutputOptionMergeKeys.IsNull() {
		var mergeKeyValues []types.String
		if !bigqueryOutputOption.BigQueryOutputOptionMergeKeys.IsUnknown() {
			diags := bigqueryOutputOption.BigQueryOutputOptionMergeKeys.ElementsAs(ctx, &mergeKeyValues, false)
			if diags.HasError() {
				return nil
			}
		}
		mergeKeys = make([]string, 0, len(mergeKeyValues))
		for _, input := range mergeKeyValues {
			mergeKeys = append(mergeKeys, input.ValueString())
		}
	}

	var columnOptionValues []bigQueryOutputOptionColumnOption
	if !bigqueryOutputOption.BigQueryOutputOptionColumnOptions.IsNull() {
		if !bigqueryOutputOption.BigQueryOutputOptionColumnOptions.IsUnknown() {
			diags := bigqueryOutputOption.BigQueryOutputOptionColumnOptions.ElementsAs(ctx, &columnOptionValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnOptionValues = []bigQueryOutputOptionColumnOption{}
		}
	} else {
		columnOptionValues = []bigQueryOutputOptionColumnOption{}
	}
	columnOptions := toInputBigqueryOutputOptionColumnOptions(&columnOptionValues)

	return &output_options2.UpdateBigQueryOutputOptionInput{
		Dataset:                              bigqueryOutputOption.Dataset.ValueStringPointer(),
		Table:                                bigqueryOutputOption.Table.ValueStringPointer(),
		AutoCreateDataset:                    bigqueryOutputOption.AutoCreateDataset.ValueBoolPointer(),
		OpenTimeoutSec:                       bigqueryOutputOption.OpenTimeoutSec.ValueInt64Pointer(),
		TimeoutSec:                           bigqueryOutputOption.TimeoutSec.ValueInt64Pointer(),
		SendTimeoutSec:                       bigqueryOutputOption.SendTimeoutSec.ValueInt64Pointer(),
		ReadTimeoutSec:                       bigqueryOutputOption.ReadTimeoutSec.ValueInt64Pointer(),
		Retries:                              bigqueryOutputOption.Retries.ValueInt64Pointer(),
		Mode:                                 bigqueryOutputOption.Mode.ValueStringPointer(),
		PartitioningType:                     model.NewNullableString(bigqueryOutputOption.PartitioningType),
		TimePartitioningType:                 model.NewNullableString(bigqueryOutputOption.TimePartitioningType),
		TimePartitioningField:                model.NewNullableString(bigqueryOutputOption.TimePartitioningField),
		TimePartitioningExpirationMs:         model.NewNullableInt64(bigqueryOutputOption.TimePartitioningExpirationMs),
		Location:                             bigqueryOutputOption.Location.ValueStringPointer(),
		TemplateTable:                        model.NewNullableString(bigqueryOutputOption.TemplateTable),
		BigQueryConnectionID:                 bigqueryOutputOption.BigQueryConnectionID.ValueInt64Pointer(),
		CustomVariableSettings:               model.ToCustomVariableSettingInputs(bigqueryOutputOption.CustomVariableSettings),
		BigQueryOutputOptionColumnOptions:    columnOptions,
		BigQueryOutputOptionClusteringFields: &clusteringFields,
		BigQueryOutputOptionMergeKeys:        &mergeKeys,
	}
}

func toInputBigqueryOutputOptionColumnOptions(bigqueryOutputOptionColumnOptions *[]bigQueryOutputOptionColumnOption) *[]output_options2.BigQueryOutputOptionColumnOptionInput {
	if bigqueryOutputOptionColumnOptions == nil {
		return nil
	}

	outputs := make([]output_options2.BigQueryOutputOptionColumnOptionInput, 0, len(*bigqueryOutputOptionColumnOptions))
	for _, input := range *bigqueryOutputOptionColumnOptions {
		outputs = append(outputs, output_options2.BigQueryOutputOptionColumnOptionInput{
			Name:            input.Name.ValueString(),
			Type:            input.Type.ValueString(),
			Mode:            input.Mode.ValueString(),
			TimestampFormat: input.TimestampFormat.ValueStringPointer(),
			Timezone:        input.Timezone.ValueStringPointer(),
			Description:     input.Description.ValueStringPointer(),
		})
	}
	return &outputs
}
