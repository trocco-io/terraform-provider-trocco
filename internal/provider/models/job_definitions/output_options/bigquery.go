package output_options

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions/output_options"
	output_options2 "terraform-provider-trocco/internal/client/parameters/job_definitions/output_options"
	"terraform-provider-trocco/internal/provider/models"
)

type BigQueryOutputOption struct {
	Dataset                                types.String                        `tfsdk:"dataset"`
	Table                                  types.String                        `tfsdk:"table"`
	AutoCreateDataset                      types.Bool                          `tfsdk:"auto_create_dataset"`
	AutoCreateTable                        types.Bool                          `tfsdk:"auto_create_table"`
	OpenTimeoutSec                         types.Int64                         `tfsdk:"open_timeout_sec"`
	TimeoutSec                             types.Int64                         `tfsdk:"timeout_sec"`
	SendTimeoutSec                         types.Int64                         `tfsdk:"send_timeout_sec"`
	ReadTimeoutSec                         types.Int64                         `tfsdk:"read_timeout_sec"`
	Retries                                types.Int64                         `tfsdk:"retries"`
	Mode                                   types.String                        `tfsdk:"mode"`
	PartitioningType                       types.String                        `tfsdk:"partitioning_type"`
	TimePartitioningType                   types.String                        `tfsdk:"time_partitioning_type"`
	TimePartitioningField                  types.String                        `tfsdk:"time_partitioning_field"`
	TimePartitioningExpirationMs           types.Int64                         `tfsdk:"time_partitioning_expiration_ms"`
	TimePartitioningRequirePartitionFilter types.Bool                          `tfsdk:"time_partitioning_require_partition_filter"`
	Location                               types.String                        `tfsdk:"location"`
	TemplateTable                          types.String                        `tfsdk:"template_table"`
	BigQueryConnectionID                   types.Int64                         `tfsdk:"bigquery_connection_id"`
	BeforeLoad                             types.String                        `tfsdk:"before_load"`
	CustomVariableSettings                 *[]models.CustomVariableSetting     `tfsdk:"custom_variable_settings"`
	BigQueryOutputOptionColumnOptions      *[]bigQueryOutputOptionColumnOption `tfsdk:"bigquery_output_option_column_options"`
	BigQueryOutputOptionClusteringFields   *[]types.String                     `tfsdk:"bigquery_output_option_clustering_fields"`
	BigQueryOutputOptionMergeKeys          *[]types.String                     `tfsdk:"bigquery_output_option_merge_keys"`
}

type bigQueryOutputOptionColumnOption struct {
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	Mode            types.String `tfsdk:"mode"`
	TimestampFormat types.String `tfsdk:"timestamp_format"`
	Timezone        types.String `tfsdk:"timezone"`
	Description     types.String `tfsdk:"description"`
}

func NewBigQueryOutputOption(bigQueryOutputOption *output_options.BigQueryOutputOption) *BigQueryOutputOption {
	if bigQueryOutputOption == nil {
		return nil
	}
	columnOptions := make([]bigQueryOutputOptionColumnOption, 0, len(*bigQueryOutputOption.BigQueryOutputOptionColumnOptions))
	for _, input := range *bigQueryOutputOption.BigQueryOutputOptionColumnOptions {
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

	return &BigQueryOutputOption{
		CustomVariableSettings:                 models.NewCustomVariableSettings(bigQueryOutputOption.CustomVariableSettings),
		Dataset:                                types.StringValue(bigQueryOutputOption.Dataset),
		Table:                                  types.StringValue(bigQueryOutputOption.Table),
		AutoCreateDataset:                      types.BoolValue(bigQueryOutputOption.AutoCreateDataset),
		AutoCreateTable:                        types.BoolValue(bigQueryOutputOption.AutoCreateTable),
		OpenTimeoutSec:                         types.Int64Value(bigQueryOutputOption.OpenTimeoutSec),
		TimeoutSec:                             types.Int64Value(bigQueryOutputOption.TimeoutSec),
		SendTimeoutSec:                         types.Int64Value(bigQueryOutputOption.SendTimeoutSec),
		ReadTimeoutSec:                         types.Int64Value(bigQueryOutputOption.ReadTimeoutSec),
		Retries:                                types.Int64Value(bigQueryOutputOption.Retries),
		Mode:                                   types.StringValue(bigQueryOutputOption.Mode),
		PartitioningType:                       types.StringPointerValue(bigQueryOutputOption.PartitioningType),
		TimePartitioningType:                   types.StringPointerValue(bigQueryOutputOption.TimePartitioningType),
		TimePartitioningField:                  types.StringPointerValue(bigQueryOutputOption.TimePartitioningField),
		TimePartitioningExpirationMs:           types.Int64PointerValue(bigQueryOutputOption.TimePartitioningExpirationMs),
		TimePartitioningRequirePartitionFilter: types.BoolPointerValue(bigQueryOutputOption.TimePartitioningRequirePartitionFilter),
		Location:                               types.StringPointerValue(bigQueryOutputOption.Location),
		TemplateTable:                          types.StringPointerValue(bigQueryOutputOption.TemplateTable),
		BeforeLoad:                             types.StringValue(bigQueryOutputOption.BeforeLoad),
		BigQueryConnectionID:                   types.Int64Value(bigQueryOutputOption.BigQueryConnectionID),
		BigQueryOutputOptionColumnOptions:      newBigqueryOutputOptionColumnOptions(bigQueryOutputOption.BigQueryOutputOptionColumnOptions),
		BigQueryOutputOptionClusteringFields:   nil,
		BigQueryOutputOptionMergeKeys:          nil,
	}
}

func newBigqueryOutputOptionColumnOptions(bigQueryOutputOptionColumnOptions *[]output_options.BigQueryOutputOptionColumnOption) *[]bigQueryOutputOptionColumnOption {
	if bigQueryOutputOptionColumnOptions == nil {
		return nil
	}

	outputs := make([]bigQueryOutputOptionColumnOption, 0, len(*bigQueryOutputOptionColumnOptions))
	for _, input := range *bigQueryOutputOptionColumnOptions {
		columnOption := bigQueryOutputOptionColumnOption{
			Name:            types.StringValue(input.Name),
			Type:            types.StringValue(input.Type),
			Mode:            types.StringValue(input.Mode),
			TimestampFormat: types.StringPointerValue(input.TimestampFormat),
			Timezone:        types.StringPointerValue(input.Timezone),
			Description:     types.StringPointerValue(input.Description),
		}
		outputs = append(outputs, columnOption)
	}
	return &outputs
}

func (bigqueryOutputOption *BigQueryOutputOption) ToInput() *output_options2.BigQueryOutputOptionInput {
	if bigqueryOutputOption == nil {
		return nil
	}

	clusteringFields := make([]string, 0, len(*bigqueryOutputOption.BigQueryOutputOptionClusteringFields))
	for _, input := range *bigqueryOutputOption.BigQueryOutputOptionClusteringFields {
		clusteringFields = append(clusteringFields, input.ValueString())
	}
	mergeKeys := make([]string, 0, len(*bigqueryOutputOption.BigQueryOutputOptionMergeKeys))
	for _, input := range *bigqueryOutputOption.BigQueryOutputOptionMergeKeys {
		mergeKeys = append(mergeKeys, input.ValueString())
	}

	return &output_options2.BigQueryOutputOptionInput{
		Dataset:                                bigqueryOutputOption.Dataset.ValueString(),
		Table:                                  bigqueryOutputOption.Table.ValueString(),
		AutoCreateDataset:                      bigqueryOutputOption.AutoCreateDataset.ValueBool(),
		AutoCreateTable:                        bigqueryOutputOption.AutoCreateTable.ValueBool(),
		OpenTimeoutSec:                         bigqueryOutputOption.OpenTimeoutSec.ValueInt64(),
		TimeoutSec:                             bigqueryOutputOption.TimeoutSec.ValueInt64(),
		SendTimeoutSec:                         bigqueryOutputOption.SendTimeoutSec.ValueInt64(),
		ReadTimeoutSec:                         bigqueryOutputOption.ReadTimeoutSec.ValueInt64(),
		Retries:                                bigqueryOutputOption.Retries.ValueInt64(),
		Mode:                                   bigqueryOutputOption.Mode.ValueString(),
		PartitioningType:                       bigqueryOutputOption.PartitioningType.ValueStringPointer(),
		TimePartitioningType:                   bigqueryOutputOption.TimePartitioningType.ValueStringPointer(),
		TimePartitioningField:                  bigqueryOutputOption.TimePartitioningField.ValueStringPointer(),
		TimePartitioningExpirationMs:           bigqueryOutputOption.TimePartitioningExpirationMs.ValueInt64Pointer(),
		TimePartitioningRequirePartitionFilter: bigqueryOutputOption.TimePartitioningRequirePartitionFilter.ValueBoolPointer(),
		Location:                               bigqueryOutputOption.Location.ValueStringPointer(),
		TemplateTable:                          bigqueryOutputOption.TemplateTable.ValueStringPointer(),
		BigQueryConnectionID:                   bigqueryOutputOption.BigQueryConnectionID.ValueInt64(),
		BeforeLoad:                             bigqueryOutputOption.BeforeLoad.ValueString(),
		CustomVariableSettings:                 models.ToCustomVariableSettingInputs(bigqueryOutputOption.CustomVariableSettings),
		BigQueryOutputOptionColumnOptions:      toInputBigqueryOutputOptionColumnOptions(bigqueryOutputOption.BigQueryOutputOptionColumnOptions),
		BigQueryOutputOptionClusteringFields:   &clusteringFields,
		BigQueryOutputOptionMergeKeys:          &mergeKeys,
	}
}

func toInputBigqueryOutputOptionColumnOptions(bigqueryOutputOptionColumnOptions *[]bigQueryOutputOptionColumnOption) *[]output_options2.BigQueryOutputOptionColumnOptionInput {
	if bigqueryOutputOptionColumnOptions == nil {
		return nil
	}

	outputs := make([]output_options2.BigQueryOutputOptionColumnOptionInput, 0, len(*bigqueryOutputOptionColumnOptions))
	for _, input := range *bigqueryOutputOptionColumnOptions {
		columnOption := output_options2.BigQueryOutputOptionColumnOptionInput{
			Name:            input.Name.String(),
			Type:            input.Type.String(),
			Mode:            input.Mode.String(),
			TimestampFormat: input.TimestampFormat.ValueStringPointer(),
			Timezone:        input.Timezone.ValueStringPointer(),
			Description:     input.Description.ValueStringPointer(),
		}
		outputs = append(outputs, columnOption)
	}
	return &outputs
}