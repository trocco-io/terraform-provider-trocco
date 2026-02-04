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

type MysqlOutputOption struct {
	MysqlConnectionID              types.Int64  `tfsdk:"mysql_connection_id"`
	Database                       types.String `tfsdk:"database"`
	Table                          types.String `tfsdk:"table"`
	Mode                           types.String `tfsdk:"mode"`
	RetryLimit                     types.Int64  `tfsdk:"retry_limit"`
	RetryWait                      types.Int64  `tfsdk:"retry_wait"`
	MaxRetryWait                   types.Int64  `tfsdk:"max_retry_wait"`
	DefaultTimeZone                types.String `tfsdk:"default_time_zone"`
	BeforeLoad                     types.String `tfsdk:"before_load"`
	AfterLoad                      types.String `tfsdk:"after_load"`
	MysqlOutputOptionColumnOptions types.List   `tfsdk:"mysql_output_option_column_options"`
	CustomVariableSettings         types.List   `tfsdk:"custom_variable_settings"`
}

type mysqlOutputOptionColumnOption struct {
	Name      types.String `tfsdk:"name"`
	Type      types.String `tfsdk:"type"`
	Scale     types.Int64  `tfsdk:"scale"`
	Precision types.Int64  `tfsdk:"precision"`
}

func NewMysqlOutputOption(ctx context.Context, entity *output_option.MysqlOutputOption) *MysqlOutputOption {
	if entity == nil {
		return nil
	}

	result := &MysqlOutputOption{
		MysqlConnectionID: types.Int64Value(entity.MysqlConnectionID),
		Database:          types.StringValue(entity.Database),
		Table:             types.StringValue(entity.Table),
		Mode:              types.StringValue(entity.Mode),
		RetryLimit:        types.Int64Value(entity.RetryLimit),
		RetryWait:         types.Int64Value(entity.RetryWait),
		MaxRetryWait:      types.Int64Value(entity.MaxRetryWait),
		DefaultTimeZone:   types.StringValue(entity.DefaultTimeZone),
		BeforeLoad:        types.StringPointerValue(entity.BeforeLoad),
		AfterLoad:         types.StringPointerValue(entity.AfterLoad),
	}

	CustomVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, entity.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = CustomVariableSettings

	var columnOptions []output_option.MysqlOutputOptionColumnOption
	if entity.MysqlOutputOptionColumnOptions != nil {
		columnOptions = *entity.MysqlOutputOptionColumnOptions
	}
	MysqlOutputOptionColumnOptions, err := newMysqlOutputOptionColumnOptions(ctx, columnOptions)
	if err != nil {
		return nil
	}
	result.MysqlOutputOptionColumnOptions = MysqlOutputOptionColumnOptions

	return result
}

func newMysqlOutputOptionColumnOptions(
	ctx context.Context,
	mysqlOutputOptionColumnOptions []output_option.MysqlOutputOptionColumnOption,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: mysqlOutputOptionColumnOption{}.attrTypes(),
	}

	if mysqlOutputOptionColumnOptions == nil {
		return types.ListNull(objectType), nil
	}

	columnOptions := make([]mysqlOutputOptionColumnOption, 0, len(mysqlOutputOptionColumnOptions))
	for _, input := range mysqlOutputOptionColumnOptions {
		columnOption := mysqlOutputOptionColumnOption{
			Name:      types.StringValue(input.Name),
			Type:      types.StringValue(input.Type),
			Scale:     types.Int64PointerValue(input.Scale),
			Precision: types.Int64PointerValue(input.Precision),
		}
		columnOptions = append(columnOptions, columnOption)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columnOptions)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert to ListValue: %v", diags)
	}
	return listValue, nil
}

func (m mysqlOutputOptionColumnOption) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":      types.StringType,
		"type":      types.StringType,
		"scale":     types.Int64Type,
		"precision": types.Int64Type,
	}
}

func (o *MysqlOutputOption) ToInput(ctx context.Context) *outputOptionParameters.MysqlOutputOptionInput {
	if o == nil {
		return nil
	}

	var columnOptions *[]outputOptionParameters.MysqlOutputOptionColumnOptionInput
	if !o.MysqlOutputOptionColumnOptions.IsNull() && !o.MysqlOutputOptionColumnOptions.IsUnknown() {
		var columnOptionValues []mysqlOutputOptionColumnOption
		diags := o.MysqlOutputOptionColumnOptions.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}

		outputs := make([]outputOptionParameters.MysqlOutputOptionColumnOptionInput, 0, len(columnOptionValues))
		for _, input := range columnOptionValues {
			outputs = append(outputs, outputOptionParameters.MysqlOutputOptionColumnOptionInput{
				Name:      input.Name.ValueString(),
				Type:      input.Type.ValueString(),
				Scale:     input.Scale.ValueInt64Pointer(),
				Precision: input.Precision.ValueInt64Pointer(),
			})
		}
		columnOptions = &outputs
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, o.CustomVariableSettings)

	return &outputOptionParameters.MysqlOutputOptionInput{
		MysqlConnectionID:              o.MysqlConnectionID.ValueInt64(),
		Database:                       o.Database.ValueString(),
		Table:                          o.Table.ValueString(),
		Mode:                           o.Mode.ValueString(),
		RetryLimit:                     o.RetryLimit.ValueInt64(),
		RetryWait:                      o.RetryWait.ValueInt64(),
		MaxRetryWait:                   o.MaxRetryWait.ValueInt64(),
		DefaultTimeZone:                o.DefaultTimeZone.ValueString(),
		BeforeLoad:                     o.BeforeLoad.ValueStringPointer(),
		AfterLoad:                      o.AfterLoad.ValueStringPointer(),
		MysqlOutputOptionColumnOptions: model.WrapObjectList(columnOptions),
		CustomVariableSettings:         model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (o *MysqlOutputOption) ToUpdateInput(ctx context.Context) *outputOptionParameters.UpdateMysqlOutputOptionInput {
	if o == nil {
		return nil
	}

	var columnOptions *[]outputOptionParameters.MysqlOutputOptionColumnOptionInput
	if !o.MysqlOutputOptionColumnOptions.IsNull() && !o.MysqlOutputOptionColumnOptions.IsUnknown() {
		var columnOptionValues []mysqlOutputOptionColumnOption
		diags := o.MysqlOutputOptionColumnOptions.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}

		outputs := make([]outputOptionParameters.MysqlOutputOptionColumnOptionInput, 0, len(columnOptionValues))
		for _, input := range columnOptionValues {
			outputs = append(outputs, outputOptionParameters.MysqlOutputOptionColumnOptionInput{
				Name:      input.Name.ValueString(),
				Type:      input.Type.ValueString(),
				Scale:     input.Scale.ValueInt64Pointer(),
				Precision: input.Precision.ValueInt64Pointer(),
			})
		}
		columnOptions = &outputs
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, o.CustomVariableSettings)

	return &outputOptionParameters.UpdateMysqlOutputOptionInput{
		MysqlConnectionID:              o.MysqlConnectionID.ValueInt64Pointer(),
		Database:                       o.Database.ValueStringPointer(),
		Table:                          o.Table.ValueStringPointer(),
		Mode:                           o.Mode.ValueStringPointer(),
		RetryLimit:                     o.RetryLimit.ValueInt64Pointer(),
		RetryWait:                      o.RetryWait.ValueInt64Pointer(),
		MaxRetryWait:                   o.MaxRetryWait.ValueInt64Pointer(),
		DefaultTimeZone:                o.DefaultTimeZone.ValueStringPointer(),
		BeforeLoad:                     o.BeforeLoad.ValueStringPointer(),
		AfterLoad:                      o.AfterLoad.ValueStringPointer(),
		MysqlOutputOptionColumnOptions: model.WrapObjectList(columnOptions),
		CustomVariableSettings:         model.ToCustomVariableSettingInputs(customVarSettings),
	}
}
