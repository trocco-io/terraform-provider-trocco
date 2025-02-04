package output_options

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	output_options2 "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SnowflakeOutputOption struct {
	Warehouse                          types.String                        `tfsdk:"warehouse"`
	Database                           types.String                        `tfsdk:"database"`
	Schema                             types.String                        `tfsdk:"schema"`
	Table                              types.String                        `tfsdk:"table"`
	Mode                               types.String                        `tfsdk:"mode"`
	EmptyFieldAsNull                   types.Bool                          `tfsdk:"empty_field_as_null"`
	DeleteStageOnError                 types.Bool                          `tfsdk:"delete_stage_on_error"`
	BatchSize                          types.Int64                         `tfsdk:"batch_size"`
	RetryLimit                         types.Int64                         `tfsdk:"retry_limit"`
	RetryWait                          types.Int64                         `tfsdk:"retry_wait"`
	MaxRetryWait                       types.Int64                         `tfsdk:"max_retry_wait"`
	DefaultTimeZone                    types.String                        `tfsdk:"default_time_zone"`
	SnowflakeConnectionId              types.Int64                         `tfsdk:"snowflake_connection_id"`
	SnowflakeOutputOptionColumnOptions []snowflakeOutputOptionColumnOption `tfsdk:"snowflake_output_option_column_options"`
	SnowflakeOutputOptionMergeKeys     []types.String                      `tfsdk:"snowflake_output_option_merge_keys"`
	CustomVariableSettings             *[]model.CustomVariableSetting      `tfsdk:"custom_variable_settings"`
}

type snowflakeOutputOptionColumnOption struct {
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	ValueType       types.String `tfsdk:"value_type"`
	TimestampFormat types.String `tfsdk:"timestamp_format"`
	Timezone        types.String `tfsdk:"timezone"`
}

func NewSnowflakeOutputOption(snowflakeOutputOption *output_option.SnowflakeOutputOption) *SnowflakeOutputOption {
	if snowflakeOutputOption == nil {
		return nil
	}

	return &SnowflakeOutputOption{
		Warehouse:                          types.StringValue(snowflakeOutputOption.Warehouse),
		Database:                           types.StringValue(snowflakeOutputOption.Database),
		Schema:                             types.StringValue(snowflakeOutputOption.Schema),
		Table:                              types.StringValue(snowflakeOutputOption.Table),
		Mode:                               types.StringPointerValue(snowflakeOutputOption.Mode),
		EmptyFieldAsNull:                   types.BoolPointerValue(snowflakeOutputOption.EmptyFieldAsNull),
		DeleteStageOnError:                 types.BoolPointerValue(snowflakeOutputOption.DeleteStageOnError),
		BatchSize:                          types.Int64PointerValue(snowflakeOutputOption.BatchSize),
		RetryLimit:                         types.Int64PointerValue(snowflakeOutputOption.RetryLimit),
		RetryWait:                          types.Int64PointerValue(snowflakeOutputOption.RetryWait),
		MaxRetryWait:                       types.Int64PointerValue(snowflakeOutputOption.MaxRetryWait),
		DefaultTimeZone:                    types.StringPointerValue(snowflakeOutputOption.DefaultTimeZone),
		SnowflakeConnectionId:              types.Int64Value(snowflakeOutputOption.SnowflakeConnectionId),
		SnowflakeOutputOptionColumnOptions: newSnowflakeOutputOptionColumnOptions(snowflakeOutputOption.SnowflakeOutputOptionColumnOptions),
		SnowflakeOutputOptionMergeKeys:     newSnowflakeOutputOptionMergeKeys(snowflakeOutputOption.SnowflakeOutputOptionMergeKeys),
		CustomVariableSettings:             model.NewCustomVariableSettings(snowflakeOutputOption.CustomVariableSettings),
	}
}

func newSnowflakeOutputOptionMergeKeys(mergeKeys []string) []types.String {
	if mergeKeys == nil {
		return nil
	}

	outputs := make([]types.String, 0, len(mergeKeys))
	for _, input := range mergeKeys {
		outputs = append(outputs, types.StringValue(input))
	}
	return outputs
}

func newSnowflakeOutputOptionColumnOptions(snowflakeOutputOptionColumnOptions []output_option.SnowflakeOutputOptionColumnOption) []snowflakeOutputOptionColumnOption {
	if snowflakeOutputOptionColumnOptions == nil {
		return nil
	}

	outputs := make([]snowflakeOutputOptionColumnOption, 0, len(snowflakeOutputOptionColumnOptions))
	for _, input := range snowflakeOutputOptionColumnOptions {
		columnOption := snowflakeOutputOptionColumnOption{
			Name:            types.StringValue(input.Name),
			Type:            types.StringValue(input.Type),
			ValueType:       types.StringPointerValue(input.ValueType),
			TimestampFormat: types.StringPointerValue(input.TimestampFormat),
			Timezone:        types.StringPointerValue(input.Timezone),
		}
		outputs = append(outputs, columnOption)
	}
	return outputs
}

func (snowflakeOutputOption *SnowflakeOutputOption) ToInput() *output_options2.SnowflakeOutputOptionInput {
	if snowflakeOutputOption == nil {
		return nil
	}

	var mergeKeys *[]string
	if snowflakeOutputOption.SnowflakeOutputOptionMergeKeys != nil {
		mk := make([]string, 0, len(snowflakeOutputOption.SnowflakeOutputOptionMergeKeys))
		for _, input := range snowflakeOutputOption.SnowflakeOutputOptionMergeKeys {
			mk = append(mk, input.ValueString())
		}
		mergeKeys = &mk
	}

	return &output_options2.SnowflakeOutputOptionInput{
		Warehouse:                          snowflakeOutputOption.Warehouse.ValueString(),
		Database:                           snowflakeOutputOption.Database.ValueString(),
		Schema:                             snowflakeOutputOption.Schema.ValueString(),
		Table:                              snowflakeOutputOption.Table.ValueString(),
		Mode:                               model.NewNullableString(snowflakeOutputOption.Mode),
		EmptyFieldAsNull:                   model.NewNullableBool(snowflakeOutputOption.EmptyFieldAsNull),
		DeleteStageOnError:                 model.NewNullableBool(snowflakeOutputOption.DeleteStageOnError),
		BatchSize:                          model.NewNullableInt64(snowflakeOutputOption.BatchSize),
		RetryLimit:                         model.NewNullableInt64(snowflakeOutputOption.RetryLimit),
		RetryWait:                          model.NewNullableInt64(snowflakeOutputOption.RetryWait),
		MaxRetryWait:                       model.NewNullableInt64(snowflakeOutputOption.MaxRetryWait),
		DefaultTimeZone:                    model.NewNullableString(snowflakeOutputOption.DefaultTimeZone),
		SnowflakeConnectionId:              snowflakeOutputOption.SnowflakeConnectionId.ValueInt64(),
		SnowflakeOutputOptionColumnOptions: model.WrapObjectList(toInputSnowflakeOutputOptionColumnOptions(snowflakeOutputOption.SnowflakeOutputOptionColumnOptions)),
		SnowflakeOutputOptionMergeKeys:     model.WrapObjectList(mergeKeys),
		CustomVariableSettings:             model.ToCustomVariableSettingInputs(snowflakeOutputOption.CustomVariableSettings),
	}
}

func (snowflakeOutputOption *SnowflakeOutputOption) ToUpdateInput() *output_options2.UpdateSnowflakeOutputOptionInput {
	if snowflakeOutputOption == nil {
		return nil
	}

	var mergeKeys *[]string
	if snowflakeOutputOption.SnowflakeOutputOptionMergeKeys != nil {
		mk := make([]string, 0, len(snowflakeOutputOption.SnowflakeOutputOptionMergeKeys))
		for _, input := range snowflakeOutputOption.SnowflakeOutputOptionMergeKeys {
			mk = append(mk, input.ValueString())
		}
		mergeKeys = &mk
	}

	return &output_options2.UpdateSnowflakeOutputOptionInput{
		Warehouse:                          snowflakeOutputOption.Warehouse.ValueStringPointer(),
		Database:                           snowflakeOutputOption.Database.ValueStringPointer(),
		Schema:                             snowflakeOutputOption.Schema.ValueStringPointer(),
		Table:                              snowflakeOutputOption.Table.ValueStringPointer(),
		Mode:                               model.NewNullableString(snowflakeOutputOption.Mode),
		EmptyFieldAsNull:                   model.NewNullableBool(snowflakeOutputOption.EmptyFieldAsNull),
		DeleteStageOnError:                 model.NewNullableBool(snowflakeOutputOption.DeleteStageOnError),
		BatchSize:                          model.NewNullableInt64(snowflakeOutputOption.BatchSize),
		RetryLimit:                         model.NewNullableInt64(snowflakeOutputOption.RetryLimit),
		RetryWait:                          model.NewNullableInt64(snowflakeOutputOption.RetryWait),
		MaxRetryWait:                       model.NewNullableInt64(snowflakeOutputOption.MaxRetryWait),
		DefaultTimeZone:                    model.NewNullableString(snowflakeOutputOption.DefaultTimeZone),
		SnowflakeConnectionId:              snowflakeOutputOption.SnowflakeConnectionId.ValueInt64Pointer(),
		SnowflakeOutputOptionColumnOptions: model.WrapObjectList(toInputSnowflakeOutputOptionColumnOptions(snowflakeOutputOption.SnowflakeOutputOptionColumnOptions)),
		SnowflakeOutputOptionMergeKeys:     model.WrapObjectList(mergeKeys),
		CustomVariableSettings:             model.ToCustomVariableSettingInputs(snowflakeOutputOption.CustomVariableSettings),
	}
}

func toInputSnowflakeOutputOptionColumnOptions(snowflakeOutputOptionColumnOptions []snowflakeOutputOptionColumnOption) *[]output_options2.SnowflakeOutputOptionColumnOptionInput {
	if snowflakeOutputOptionColumnOptions == nil {
		return nil
	}

	outputs := make([]output_options2.SnowflakeOutputOptionColumnOptionInput, 0, len(snowflakeOutputOptionColumnOptions))
	for _, input := range snowflakeOutputOptionColumnOptions {
		outputs = append(outputs, output_options2.SnowflakeOutputOptionColumnOptionInput{
			Name:            input.Name.ValueString(),
			Type:            input.Type.ValueString(),
			ValueType:       input.ValueType.ValueStringPointer(),
			TimestampFormat: input.TimestampFormat.ValueStringPointer(),
			Timezone:        input.Timezone.ValueStringPointer(),
		})
	}
	return &outputs
}
