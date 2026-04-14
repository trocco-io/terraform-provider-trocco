package output_options

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	outputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RedshiftOutputOption struct {
	RedshiftConnectionID              types.Int64  `tfsdk:"redshift_connection_id"`
	Database                          types.String `tfsdk:"database"`
	Schema                            types.String `tfsdk:"schema"`
	Table                             types.String `tfsdk:"table"`
	CreateTableConstraint             types.String `tfsdk:"create_table_constraint"`
	CreateTableOption                 types.String `tfsdk:"create_table_option"`
	S3Bucket                          types.String `tfsdk:"s3_bucket"`
	S3KeyPrefix                       types.String `tfsdk:"s3_key_prefix"`
	DeleteS3TempFile                  types.Bool   `tfsdk:"delete_s3_temp_file"`
	CopyIAMRoleName                   types.String `tfsdk:"copy_iam_role_name"`
	RetryLimit                        types.Int64  `tfsdk:"retry_limit"`
	RetryWait                         types.Int64  `tfsdk:"retry_wait"`
	MaxRetryWait                      types.Int64  `tfsdk:"max_retry_wait"`
	Mode                              types.String `tfsdk:"mode"`
	DefaultTimeZone                   types.String `tfsdk:"default_time_zone"`
	BeforeLoad                        types.String `tfsdk:"before_load"`
	AfterLoad                         types.String `tfsdk:"after_load"`
	BatchSize                         types.Int64  `tfsdk:"batch_size"`
	RedshiftOutputOptionColumnOptions types.List   `tfsdk:"redshift_output_option_column_options"`
	RedshiftOutputOptionMergeKeys     types.Set    `tfsdk:"redshift_output_option_merge_keys"`
	CustomVariableSettings            types.List   `tfsdk:"custom_variable_settings"`
}

type redshiftOutputOptionColumnOption struct {
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	ValueType       types.String `tfsdk:"value_type"`
	TimestampFormat types.String `tfsdk:"timestamp_format"`
	Timezone        types.String `tfsdk:"timezone"`
}

func (r redshiftOutputOptionColumnOption) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":             types.StringType,
		"type":             types.StringType,
		"value_type":       types.StringType,
		"timestamp_format": types.StringType,
		"timezone":         types.StringType,
	}
}

func NewRedshiftOutputOption(ctx context.Context, entity *output_option.RedshiftOutputOption) *RedshiftOutputOption {
	if entity == nil {
		return nil
	}

	result := &RedshiftOutputOption{
		RedshiftConnectionID:  types.Int64Value(entity.RedshiftConnectionID),
		Database:              types.StringValue(entity.Database),
		Schema:                types.StringValue(entity.Schema),
		Table:                 types.StringValue(entity.Table),
		CreateTableConstraint: types.StringPointerValue(entity.CreateTableConstraint),
		CreateTableOption:     types.StringPointerValue(entity.CreateTableOption),
		S3Bucket:              types.StringValue(entity.S3Bucket),
		S3KeyPrefix:           types.StringValue(entity.S3KeyPrefix),
		DeleteS3TempFile:      types.BoolValue(entity.DeleteS3TempFile),
		CopyIAMRoleName:       types.StringPointerValue(entity.CopyIAMRoleName),
		RetryLimit:            types.Int64Value(entity.RetryLimit),
		RetryWait:             types.Int64Value(entity.RetryWait),
		MaxRetryWait:          types.Int64Value(entity.MaxRetryWait),
		Mode:                  types.StringValue(entity.Mode),
		DefaultTimeZone:       types.StringValue(entity.DefaultTimeZone),
		BeforeLoad:            types.StringPointerValue(entity.BeforeLoad),
		AfterLoad:             types.StringPointerValue(entity.AfterLoad),
		BatchSize:             types.Int64Value(entity.BatchSize),
	}

	columnOptions, err := newRedshiftOutputOptionColumnOptions(ctx, entity.RedshiftOutputOptionColumnOptions)
	if err != nil {
		return nil
	}
	result.RedshiftOutputOptionColumnOptions = columnOptions

	mergeKeys, err := newRedshiftMergeKeys(ctx, entity.RedshiftOutputOptionMergeKeys)
	if err != nil {
		return nil
	}
	result.RedshiftOutputOptionMergeKeys = mergeKeys

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, entity.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func newRedshiftOutputOptionColumnOptions(
	ctx context.Context,
	columnOptions []output_option.RedshiftOutputOptionColumnOption,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: redshiftOutputOptionColumnOption{}.attrTypes(),
	}

	if columnOptions == nil {
		return types.ListNull(objectType), nil
	}

	items := make([]redshiftOutputOptionColumnOption, 0, len(columnOptions))
	for _, input := range columnOptions {
		items = append(items, redshiftOutputOptionColumnOption{
			Name:            types.StringValue(input.Name),
			Type:            types.StringPointerValue(input.Type),
			ValueType:       types.StringPointerValue(input.ValueType),
			TimestampFormat: types.StringPointerValue(input.TimestampFormat),
			Timezone:        types.StringPointerValue(input.Timezone),
		})
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, items)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert to ListValue: %v", diags)
	}
	return listValue, nil
}

func newRedshiftMergeKeys(ctx context.Context, mergeKeys []string) (types.Set, error) {
	if len(mergeKeys) > 0 {
		values := make([]types.String, len(mergeKeys))
		for i, v := range mergeKeys {
			values[i] = types.StringValue(v)
		}
		setValue, diags := types.SetValueFrom(ctx, types.StringType, values)
		if diags.HasError() {
			return types.SetNull(types.StringType), fmt.Errorf("failed to convert to SetValue: %v", diags)
		}
		return setValue, nil
	}
	return types.SetNull(types.StringType), nil
}

func (o *RedshiftOutputOption) ToInput(ctx context.Context) *outputOptionParameters.RedshiftOutputOptionInput {
	if o == nil {
		return nil
	}

	columnOptions := toRedshiftColumnOptionInputs(ctx, o.RedshiftOutputOptionColumnOptions)
	mergeKeys := toRedshiftMergeKeyInputs(ctx, o.RedshiftOutputOptionMergeKeys)
	customVarSettings := common.ExtractCustomVariableSettings(ctx, o.CustomVariableSettings)

	return &outputOptionParameters.RedshiftOutputOptionInput{
		RedshiftConnectionID:              o.RedshiftConnectionID.ValueInt64(),
		Database:                          o.Database.ValueString(),
		Schema:                            o.Schema.ValueString(),
		Table:                             o.Table.ValueString(),
		CreateTableConstraint:             o.CreateTableConstraint.ValueStringPointer(),
		CreateTableOption:                 o.CreateTableOption.ValueStringPointer(),
		S3Bucket:                          o.S3Bucket.ValueString(),
		S3KeyPrefix:                       o.S3KeyPrefix.ValueString(),
		DeleteS3TempFile:                  o.DeleteS3TempFile.ValueBool(),
		CopyIAMRoleName:                   o.CopyIAMRoleName.ValueStringPointer(),
		RetryLimit:                        o.RetryLimit.ValueInt64(),
		RetryWait:                         o.RetryWait.ValueInt64(),
		MaxRetryWait:                      o.MaxRetryWait.ValueInt64(),
		Mode:                              o.Mode.ValueString(),
		DefaultTimeZone:                   o.DefaultTimeZone.ValueString(),
		BeforeLoad:                        o.BeforeLoad.ValueStringPointer(),
		AfterLoad:                         o.AfterLoad.ValueStringPointer(),
		BatchSize:                         o.BatchSize.ValueInt64(),
		RedshiftOutputOptionColumnOptions: model.WrapObjectList(columnOptions),
		RedshiftOutputOptionMergeKeys:     model.WrapObjectList(mergeKeys),
		CustomVariableSettings:            model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (o *RedshiftOutputOption) ToUpdateInput(ctx context.Context) *outputOptionParameters.UpdateRedshiftOutputOptionInput {
	if o == nil {
		return nil
	}

	columnOptions := toRedshiftColumnOptionInputs(ctx, o.RedshiftOutputOptionColumnOptions)
	mergeKeys := toRedshiftMergeKeyInputs(ctx, o.RedshiftOutputOptionMergeKeys)
	customVarSettings := common.ExtractCustomVariableSettings(ctx, o.CustomVariableSettings)

	return &outputOptionParameters.UpdateRedshiftOutputOptionInput{
		RedshiftConnectionID:              o.RedshiftConnectionID.ValueInt64Pointer(),
		Database:                          o.Database.ValueStringPointer(),
		Schema:                            o.Schema.ValueStringPointer(),
		Table:                             o.Table.ValueStringPointer(),
		CreateTableConstraint:             o.CreateTableConstraint.ValueStringPointer(),
		CreateTableOption:                 o.CreateTableOption.ValueStringPointer(),
		S3Bucket:                          o.S3Bucket.ValueStringPointer(),
		S3KeyPrefix:                       o.S3KeyPrefix.ValueStringPointer(),
		DeleteS3TempFile:                  o.DeleteS3TempFile.ValueBoolPointer(),
		CopyIAMRoleName:                   o.CopyIAMRoleName.ValueStringPointer(),
		RetryLimit:                        o.RetryLimit.ValueInt64Pointer(),
		RetryWait:                         o.RetryWait.ValueInt64Pointer(),
		MaxRetryWait:                      o.MaxRetryWait.ValueInt64Pointer(),
		Mode:                              o.Mode.ValueStringPointer(),
		DefaultTimeZone:                   o.DefaultTimeZone.ValueStringPointer(),
		BeforeLoad:                        o.BeforeLoad.ValueStringPointer(),
		AfterLoad:                         o.AfterLoad.ValueStringPointer(),
		BatchSize:                         o.BatchSize.ValueInt64Pointer(),
		RedshiftOutputOptionColumnOptions: model.WrapObjectList(columnOptions),
		RedshiftOutputOptionMergeKeys:     model.WrapObjectList(mergeKeys),
		CustomVariableSettings:            model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func toRedshiftColumnOptionInputs(
	ctx context.Context,
	list types.List,
) *[]outputOptionParameters.RedshiftOutputOptionColumnOptionInput {
	if list.IsNull() || list.IsUnknown() {
		return nil
	}

	var values []redshiftOutputOptionColumnOption
	diags := list.ElementsAs(ctx, &values, false)
	if diags.HasError() {
		return nil
	}

	outputs := make([]outputOptionParameters.RedshiftOutputOptionColumnOptionInput, 0, len(values))
	for _, v := range values {
		outputs = append(outputs, outputOptionParameters.RedshiftOutputOptionColumnOptionInput{
			Name:            v.Name.ValueString(),
			Type:            v.Type.ValueStringPointer(),
			ValueType:       v.ValueType.ValueStringPointer(),
			TimestampFormat: v.TimestampFormat.ValueStringPointer(),
			Timezone:        v.Timezone.ValueStringPointer(),
		})
	}
	return &outputs
}

func toRedshiftMergeKeyInputs(ctx context.Context, set types.Set) *[]string {
	if set.IsNull() || set.IsUnknown() {
		return nil
	}

	var values []types.String
	diags := set.ElementsAs(ctx, &values, false)
	if diags.HasError() {
		return nil
	}

	keys := make([]string, 0, len(values))
	for _, v := range values {
		keys = append(keys, v.ValueString())
	}
	return &keys
}
