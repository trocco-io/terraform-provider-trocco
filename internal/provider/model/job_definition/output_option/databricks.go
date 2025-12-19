package output_options

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	outputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DatabricksOutputOption struct {
	DatabricksConnectionID              types.Int64  `tfsdk:"databricks_connection_id"`
	CatalogName                         types.String `tfsdk:"catalog_name"`
	SchemaName                          types.String `tfsdk:"schema_name"`
	Table                               types.String `tfsdk:"table"`
	BatchSize                           types.Int64  `tfsdk:"batch_size"`
	Mode                                types.String `tfsdk:"mode"`
	DefaultTimeZone                     types.String `tfsdk:"default_time_zone"`
	DatabricksOutputOptionColumnOptions types.List   `tfsdk:"databricks_output_option_column_options"`
	DatabricksOutputOptionMergeKeys     types.Set    `tfsdk:"databricks_output_option_merge_keys"`
}

type databricksOutputOptionColumnOption struct {
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	ValueType       types.String `tfsdk:"value_type"`
	TimestampFormat types.String `tfsdk:"timestamp_format"`
	Timezone        types.String `tfsdk:"timezone"`
}

func NewDatabricksOutputOption(ctx context.Context, entity *output_option.DatabricksOutputOption) *DatabricksOutputOption {
	if entity == nil {
		return nil
	}

	result := &DatabricksOutputOption{
		DatabricksConnectionID: types.Int64Value(entity.DatabricksConnectionID),
		CatalogName:            types.StringValue(entity.CatalogName),
		SchemaName:             types.StringValue(entity.SchemaName),
		Table:                  types.StringValue(entity.Table),
		BatchSize:              types.Int64Value(entity.BatchSize),
		Mode:                   types.StringValue(entity.Mode),
		DefaultTimeZone:        types.StringValue(entity.DefaultTimeZone),
	}

	DatabricksOutputOptionMergeKeys, err := newDatabricksOutputOptionMergeKeys(ctx, entity.DatabricksOutputOptionMergeKeys)
	if err != nil {
		return nil
	}
	result.DatabricksOutputOptionMergeKeys = DatabricksOutputOptionMergeKeys

	DatabricksOutputOptionColumnOptions, err := newDatabricksOutputOptionColumnOptions(ctx, entity.DatabricksOutputOptionColumnOptions)
	if err != nil {
		return nil
	}
	result.DatabricksOutputOptionColumnOptions = DatabricksOutputOptionColumnOptions

	return result
}

func newDatabricksOutputOptionMergeKeys(ctx context.Context, mergeKeys []string) (types.Set, error) {
	if mergeKeys != nil {
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

func newDatabricksOutputOptionColumnOptions(ctx context.Context, inputOptions []output_option.DatabricksOutputOptionColumnOption) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: databricksOutputOptionColumnOption{}.attrTypes(),
	}

	if inputOptions == nil {
		return types.ListNull(objectType), nil
	}

	columnOptions := make([]databricksOutputOptionColumnOption, 0, len(inputOptions))
	for _, input := range inputOptions {
		columnOption := databricksOutputOptionColumnOption{
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

func (d databricksOutputOptionColumnOption) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":             types.StringType,
		"type":             types.StringType,
		"value_type":       types.StringType,
		"timestamp_format": types.StringType,
		"timezone":         types.StringType,
	}
}

func (o *DatabricksOutputOption) ToInput(ctx context.Context) *outputOptionParameters.DatabricksOutputOptionInput {
	if o == nil {
		return nil
	}

	var mergeKeys *[]string
	if !o.DatabricksOutputOptionMergeKeys.IsNull() {
		var mergeKeyValues []types.String
		diags := o.DatabricksOutputOptionMergeKeys.ElementsAs(ctx, &mergeKeyValues, false)
		if diags.HasError() {
			return nil
		}
		mk := make([]string, 0, len(mergeKeyValues))
		for _, input := range mergeKeyValues {
			mk = append(mk, input.ValueString())
		}
		mergeKeys = &mk
	}

	var columnOptions *[]outputOptionParameters.DatabricksOutputOptionColumnOptionInput
	if !o.DatabricksOutputOptionColumnOptions.IsNull() && !o.DatabricksOutputOptionColumnOptions.IsUnknown() {
		var columnOptionValues []databricksOutputOptionColumnOption
		diags := o.DatabricksOutputOptionColumnOptions.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}

		outputs := make([]outputOptionParameters.DatabricksOutputOptionColumnOptionInput, 0, len(columnOptionValues))
		for _, input := range columnOptionValues {
			outputs = append(outputs, outputOptionParameters.DatabricksOutputOptionColumnOptionInput{
				Name:            input.Name.ValueString(),
				Type:            input.Type.ValueString(),
				ValueType:       input.ValueType.ValueStringPointer(),
				TimestampFormat: input.TimestampFormat.ValueStringPointer(),
				Timezone:        input.Timezone.ValueStringPointer(),
			})
		}
		columnOptions = &outputs
	}

	return &outputOptionParameters.DatabricksOutputOptionInput{
		DatabricksConnectionID:              o.DatabricksConnectionID.ValueInt64(),
		CatalogName:                         o.CatalogName.ValueString(),
		SchemaName:                          o.SchemaName.ValueString(),
		Table:                               o.Table.ValueString(),
		BatchSize:                           o.BatchSize.ValueInt64(),
		Mode:                                o.Mode.ValueString(),
		DefaultTimeZone:                     o.DefaultTimeZone.ValueString(),
		DatabricksOutputOptionColumnOptions: model.WrapObjectList(columnOptions),
		DatabricksOutputOptionMergeKeys:     model.WrapObjectList(mergeKeys),
	}
}

func (o *DatabricksOutputOption) ToUpdateInput(ctx context.Context) *outputOptionParameters.UpdateDatabricksOutputOptionInput {
	if o == nil {
		return nil
	}

	var mergeKeys *[]string
	if !o.DatabricksOutputOptionMergeKeys.IsNull() {
		var mergeKeyValues []types.String
		diags := o.DatabricksOutputOptionMergeKeys.ElementsAs(ctx, &mergeKeyValues, false)
		if diags.HasError() {
			return nil
		}

		mk := make([]string, 0, len(mergeKeyValues))
		for _, input := range mergeKeyValues {
			mk = append(mk, input.ValueString())
		}
		mergeKeys = &mk
	}

	var columnOptions *[]outputOptionParameters.DatabricksOutputOptionColumnOptionInput
	if !o.DatabricksOutputOptionColumnOptions.IsNull() && !o.DatabricksOutputOptionColumnOptions.IsUnknown() {
		var columnOptionValues []databricksOutputOptionColumnOption
		diags := o.DatabricksOutputOptionColumnOptions.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}

		outputs := make([]outputOptionParameters.DatabricksOutputOptionColumnOptionInput, 0, len(columnOptionValues))
		for _, input := range columnOptionValues {
			outputs = append(outputs, outputOptionParameters.DatabricksOutputOptionColumnOptionInput{
				Name:            input.Name.ValueString(),
				Type:            input.Type.ValueString(),
				ValueType:       input.ValueType.ValueStringPointer(),
				TimestampFormat: input.TimestampFormat.ValueStringPointer(),
				Timezone:        input.Timezone.ValueStringPointer(),
			})
		}
		columnOptions = &outputs
	}

	return &outputOptionParameters.UpdateDatabricksOutputOptionInput{
		DatabricksConnectionID:              o.DatabricksConnectionID.ValueInt64Pointer(),
		CatalogName:                         o.CatalogName.ValueStringPointer(),
		SchemaName:                          o.SchemaName.ValueStringPointer(),
		Table:                               o.Table.ValueStringPointer(),
		BatchSize:                           o.BatchSize.ValueInt64Pointer(),
		Mode:                                o.Mode.ValueStringPointer(),
		DefaultTimeZone:                     o.DefaultTimeZone.ValueStringPointer(),
		DatabricksOutputOptionColumnOptions: model.WrapObjectList(columnOptions),
		DatabricksOutputOptionMergeKeys:     model.WrapObjectList(mergeKeys),
	}
}
