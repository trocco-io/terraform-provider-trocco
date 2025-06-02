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

type SnowflakeOutputOption struct {
	Warehouse                          types.String `tfsdk:"warehouse"`
	Database                           types.String `tfsdk:"database"`
	Schema                             types.String `tfsdk:"schema"`
	Table                              types.String `tfsdk:"table"`
	Mode                               types.String `tfsdk:"mode"`
	EmptyFieldAsNull                   types.Bool   `tfsdk:"empty_field_as_null"`
	DeleteStageOnError                 types.Bool   `tfsdk:"delete_stage_on_error"`
	BatchSize                          types.Int64  `tfsdk:"batch_size"`
	RetryLimit                         types.Int64  `tfsdk:"retry_limit"`
	RetryWait                          types.Int64  `tfsdk:"retry_wait"`
	MaxRetryWait                       types.Int64  `tfsdk:"max_retry_wait"`
	DefaultTimeZone                    types.String `tfsdk:"default_time_zone"`
	SnowflakeConnectionId              types.Int64  `tfsdk:"snowflake_connection_id"`
	SnowflakeOutputOptionColumnOptions types.List   `tfsdk:"snowflake_output_option_column_options"`
	SnowflakeOutputOptionMergeKeys     types.Set    `tfsdk:"snowflake_output_option_merge_keys"`
	CustomVariableSettings             types.List   `tfsdk:"custom_variable_settings"`
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

	ctx := context.Background()

	result := &SnowflakeOutputOption{
		Warehouse:             types.StringValue(snowflakeOutputOption.Warehouse),
		Database:              types.StringValue(snowflakeOutputOption.Database),
		Schema:                types.StringValue(snowflakeOutputOption.Schema),
		Table:                 types.StringValue(snowflakeOutputOption.Table),
		Mode:                  types.StringPointerValue(snowflakeOutputOption.Mode),
		EmptyFieldAsNull:      types.BoolPointerValue(snowflakeOutputOption.EmptyFieldAsNull),
		DeleteStageOnError:    types.BoolPointerValue(snowflakeOutputOption.DeleteStageOnError),
		BatchSize:             types.Int64PointerValue(snowflakeOutputOption.BatchSize),
		RetryLimit:            types.Int64PointerValue(snowflakeOutputOption.RetryLimit),
		RetryWait:             types.Int64PointerValue(snowflakeOutputOption.RetryWait),
		MaxRetryWait:          types.Int64PointerValue(snowflakeOutputOption.MaxRetryWait),
		DefaultTimeZone:       types.StringPointerValue(snowflakeOutputOption.DefaultTimeZone),
		SnowflakeConnectionId: types.Int64Value(snowflakeOutputOption.SnowflakeConnectionId),
	}

	SnowflakeOutputOptionMergeKeys, err := newSnowflakeOutputOptionMergeKeys(ctx, snowflakeOutputOption.SnowflakeOutputOptionMergeKeys)
	if err != nil {
		return nil
	}
	result.SnowflakeOutputOptionMergeKeys = SnowflakeOutputOptionMergeKeys

	SnowflakeOutputOptionColumnOptions, err := newSnowflakeOutputOptionColumnOptions(ctx, snowflakeOutputOption.SnowflakeOutputOptionColumnOptions)
	if err != nil {
		return nil
	}
	result.SnowflakeOutputOptionColumnOptions = SnowflakeOutputOptionColumnOptions

	CustomVariableSettings, err := ConvertCustomVariableSettingsToList(ctx, snowflakeOutputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = CustomVariableSettings

	return result
}

func newSnowflakeOutputOptionMergeKeys(ctx context.Context, mergeKeys []string) (types.Set, error) {
	if mergeKeys != nil {
		values := make([]types.String, len(mergeKeys))
		for i, v := range mergeKeys {
			values[i] = types.StringValue(v)
		}
		setValue, diags := types.SetValueFrom(ctx, types.StringType, values)
		if diags.HasError() {
			return types.SetNull(types.StringType), fmt.Errorf("failed to convert to ListValue: %v", diags)
		}
		return setValue, nil
	}
	return types.SetNull(types.StringType), nil
}

func newSnowflakeOutputOptionColumnOptions(ctx context.Context, inputOptions []output_option.SnowflakeOutputOptionColumnOption) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: snowflakeOutputOptionColumnOption{}.attrTypes(),
	}

	if inputOptions == nil {
		return types.ListNull(objectType), nil
	}

	columnOptions := make([]snowflakeOutputOptionColumnOption, 0, len(inputOptions))
	for _, input := range inputOptions {
		columnOption := snowflakeOutputOptionColumnOption{
			Name:            types.StringValue(input.Name),
			Type:            types.StringValue(input.Type),
			ValueType:       types.StringPointerValue(input.ValueType),
			TimestampFormat: types.StringPointerValue(input.TimestampFormat),
			Timezone:        types.StringPointerValue(input.Timezone),
		}
		columnOptions = append(columnOptions, columnOption)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columnOptions)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert to ListValue: %v", diags)
	}
	return listValue, nil
}

func (s snowflakeOutputOptionColumnOption) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":             types.StringType,
		"type":             types.StringType,
		"value_type":       types.StringType,
		"timestamp_format": types.StringType,
		"timezone":         types.StringType,
	}
}

func (snowflakeOutputOption *SnowflakeOutputOption) ToInput() *output_options2.SnowflakeOutputOptionInput {
	if snowflakeOutputOption == nil {
		return nil
	}

	ctx := context.Background()

	var mergeKeys *[]string
	if !snowflakeOutputOption.SnowflakeOutputOptionMergeKeys.IsNull() && !snowflakeOutputOption.SnowflakeOutputOptionMergeKeys.IsUnknown() {
		var mergeKeyValues []types.String
		diags := snowflakeOutputOption.SnowflakeOutputOptionMergeKeys.ElementsAs(ctx, &mergeKeyValues, false)
		if diags.HasError() {
			return nil
		}
		mk := make([]string, 0, len(mergeKeyValues))
		for _, input := range mergeKeyValues {
			mk = append(mk, input.ValueString())
		}
		mergeKeys = &mk
	}

	var columnOptions *[]output_options2.SnowflakeOutputOptionColumnOptionInput
	if !snowflakeOutputOption.SnowflakeOutputOptionColumnOptions.IsNull() && !snowflakeOutputOption.SnowflakeOutputOptionColumnOptions.IsUnknown() {
		var columnOptionValues []snowflakeOutputOptionColumnOption
		diags := snowflakeOutputOption.SnowflakeOutputOptionColumnOptions.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}

		outputs := make([]output_options2.SnowflakeOutputOptionColumnOptionInput, 0, len(columnOptionValues))
		for _, input := range columnOptionValues {
			outputs = append(outputs, output_options2.SnowflakeOutputOptionColumnOptionInput{
				Name:            input.Name.ValueString(),
				Type:            input.Type.ValueString(),
				ValueType:       input.ValueType.ValueStringPointer(),
				TimestampFormat: input.TimestampFormat.ValueStringPointer(),
				Timezone:        input.Timezone.ValueStringPointer(),
			})
		}
		columnOptions = &outputs
	}

	customVarSettings := ExtractCustomVariableSettings(ctx, snowflakeOutputOption.CustomVariableSettings)

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
		SnowflakeOutputOptionColumnOptions: model.WrapObjectList(columnOptions),
		SnowflakeOutputOptionMergeKeys:     model.WrapObjectList(mergeKeys),
		CustomVariableSettings:             model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (snowflakeOutputOption *SnowflakeOutputOption) ToUpdateInput() *output_options2.UpdateSnowflakeOutputOptionInput {
	if snowflakeOutputOption == nil {
		return nil
	}

	ctx := context.Background()

	var mergeKeys *[]string
	if !snowflakeOutputOption.SnowflakeOutputOptionMergeKeys.IsNull() && !snowflakeOutputOption.SnowflakeOutputOptionMergeKeys.IsUnknown() {
		var mergeKeyValues []types.String
		diags := snowflakeOutputOption.SnowflakeOutputOptionMergeKeys.ElementsAs(ctx, &mergeKeyValues, false)
		if diags.HasError() {
			return nil
		}

		mk := make([]string, 0, len(mergeKeyValues))
		for _, input := range mergeKeyValues {
			mk = append(mk, input.ValueString())
		}
		mergeKeys = &mk
	}

	var columnOptions *[]output_options2.SnowflakeOutputOptionColumnOptionInput
	if !snowflakeOutputOption.SnowflakeOutputOptionColumnOptions.IsNull() && !snowflakeOutputOption.SnowflakeOutputOptionColumnOptions.IsUnknown() {
		var columnOptionValues []snowflakeOutputOptionColumnOption
		diags := snowflakeOutputOption.SnowflakeOutputOptionColumnOptions.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}

		outputs := make([]output_options2.SnowflakeOutputOptionColumnOptionInput, 0, len(columnOptionValues))
		for _, input := range columnOptionValues {
			outputs = append(outputs, output_options2.SnowflakeOutputOptionColumnOptionInput{
				Name:            input.Name.ValueString(),
				Type:            input.Type.ValueString(),
				ValueType:       input.ValueType.ValueStringPointer(),
				TimestampFormat: input.TimestampFormat.ValueStringPointer(),
				Timezone:        input.Timezone.ValueStringPointer(),
			})
		}
		columnOptions = &outputs
	}

	customVarSettings := ExtractCustomVariableSettings(ctx, snowflakeOutputOption.CustomVariableSettings)

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
		SnowflakeOutputOptionColumnOptions: model.WrapObjectList(columnOptions),
		SnowflakeOutputOptionMergeKeys:     model.WrapObjectList(mergeKeys),
		CustomVariableSettings:             model.ToCustomVariableSettingInputs(customVarSettings),
	}
}
