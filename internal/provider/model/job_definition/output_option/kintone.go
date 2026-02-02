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

type KintoneOutputOption struct {
	KintoneConnectionID              types.Int64  `tfsdk:"kintone_connection_id"`
	AppID                            types.String `tfsdk:"app_id"`
	GuestSpaceID                     types.String `tfsdk:"guest_space_id"`
	Mode                             types.String `tfsdk:"mode"`
	UpdateKey                        types.String `tfsdk:"update_key"`
	IgnoreNulls                      types.Bool   `tfsdk:"ignore_nulls"`
	ReduceKey                        types.String `tfsdk:"reduce_key"`
	ChunkSize                        types.Int64  `tfsdk:"chunk_size"`
	KintoneOutputOptionColumnOptions types.List   `tfsdk:"kintone_output_option_column_options"`
}

type kintoneOutputOptionColumnOption struct {
	Name       types.String `tfsdk:"name"`
	FieldCode  types.String `tfsdk:"field_code"`
	Type       types.String `tfsdk:"type"`
	Timezone   types.String `tfsdk:"timezone"`
	SortColumn types.String `tfsdk:"sort_column"`
}

func NewKintoneOutputOption(ctx context.Context, kintoneOutputOption *output_option.KintoneOutputOption) *KintoneOutputOption {
	if kintoneOutputOption == nil {
		return nil
	}

	var updateKey types.String
	if kintoneOutputOption.Mode == "update" || kintoneOutputOption.Mode == "upsert" {
		updateKey = types.StringPointerValue(kintoneOutputOption.UpdateKey)
	} else {
		updateKey = types.StringNull()
	}

	result := &KintoneOutputOption{
		KintoneConnectionID: types.Int64Value(kintoneOutputOption.KintoneConnectionID),
		AppID:               types.StringValue(kintoneOutputOption.AppID),
		GuestSpaceID:        types.StringPointerValue(kintoneOutputOption.GuestSpaceID),
		Mode:                types.StringValue(kintoneOutputOption.Mode),
		UpdateKey:           updateKey,
		IgnoreNulls:         types.BoolValue(kintoneOutputOption.IgnoreNulls),
		ReduceKey:           types.StringPointerValue(kintoneOutputOption.ReduceKey),
		ChunkSize:           types.Int64Value(kintoneOutputOption.ChunkSize),
	}

	KintoneOutputOptionColumnOptions, err := newKintoneOutputOptionColumnOptions(ctx, kintoneOutputOption.KintoneOutputOptionColumnOptions)
	if err != nil {
		return nil
	}
	result.KintoneOutputOptionColumnOptions = KintoneOutputOptionColumnOptions

	return result
}

func newKintoneOutputOptionColumnOptions(ctx context.Context, inputOptions []output_option.KintoneOutputOptionColumnOption) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: kintoneOutputOptionColumnOption{}.attrTypes(),
	}

	if inputOptions == nil {
		return types.ListNull(objectType), nil
	}

	columnOptions := make([]kintoneOutputOptionColumnOption, 0, len(inputOptions))
	for _, input := range inputOptions {
		columnOption := kintoneOutputOptionColumnOption{
			Name:       types.StringValue(input.Name),
			FieldCode:  types.StringValue(input.FieldCode),
			Type:       types.StringValue(input.Type),
			Timezone:   types.StringPointerValue(input.Timezone),
			SortColumn: types.StringPointerValue(input.SortColumn),
		}
		columnOptions = append(columnOptions, columnOption)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, columnOptions)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert to ListValue: %v", diags)
	}
	return listValue, nil
}

func (k kintoneOutputOptionColumnOption) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":        types.StringType,
		"field_code":  types.StringType,
		"type":        types.StringType,
		"timezone":    types.StringType,
		"sort_column": types.StringType,
	}
}

func (kintoneOutputOption *KintoneOutputOption) ToInput(ctx context.Context) *outputOptionParameters.KintoneOutputOptionInput {
	if kintoneOutputOption == nil {
		return nil
	}

	var columnOptions *[]outputOptionParameters.KintoneOutputOptionColumnOptionInput
	if !kintoneOutputOption.KintoneOutputOptionColumnOptions.IsNull() && !kintoneOutputOption.KintoneOutputOptionColumnOptions.IsUnknown() {
		var columnOptionValues []kintoneOutputOptionColumnOption
		diags := kintoneOutputOption.KintoneOutputOptionColumnOptions.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}

		outputs := make([]outputOptionParameters.KintoneOutputOptionColumnOptionInput, 0, len(columnOptionValues))
		for _, input := range columnOptionValues {
			outputs = append(outputs, outputOptionParameters.KintoneOutputOptionColumnOptionInput{
				Name:       input.Name.ValueString(),
				FieldCode:  input.FieldCode.ValueString(),
				Type:       input.Type.ValueString(),
				Timezone:   input.Timezone.ValueStringPointer(),
				SortColumn: input.SortColumn.ValueStringPointer(),
			})
		}
		columnOptions = &outputs
	}

	return &outputOptionParameters.KintoneOutputOptionInput{
		KintoneConnectionID:              kintoneOutputOption.KintoneConnectionID.ValueInt64(),
		AppID:                            kintoneOutputOption.AppID.ValueString(),
		GuestSpaceID:                     kintoneOutputOption.GuestSpaceID.ValueStringPointer(),
		Mode:                             kintoneOutputOption.Mode.ValueString(),
		UpdateKey:                        kintoneOutputOption.UpdateKey.ValueStringPointer(),
		IgnoreNulls:                      kintoneOutputOption.IgnoreNulls.ValueBool(),
		ReduceKey:                        kintoneOutputOption.ReduceKey.ValueStringPointer(),
		ChunkSize:                        kintoneOutputOption.ChunkSize.ValueInt64(),
		KintoneOutputOptionColumnOptions: model.WrapObjectList(columnOptions),
	}
}

func (kintoneOutputOption *KintoneOutputOption) ToUpdateInput(ctx context.Context) *outputOptionParameters.UpdateKintoneOutputOptionInput {
	if kintoneOutputOption == nil {
		return nil
	}

	var columnOptions *[]outputOptionParameters.KintoneOutputOptionColumnOptionInput
	if !kintoneOutputOption.KintoneOutputOptionColumnOptions.IsNull() && !kintoneOutputOption.KintoneOutputOptionColumnOptions.IsUnknown() {
		var columnOptionValues []kintoneOutputOptionColumnOption
		diags := kintoneOutputOption.KintoneOutputOptionColumnOptions.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}

		outputs := make([]outputOptionParameters.KintoneOutputOptionColumnOptionInput, 0, len(columnOptionValues))
		for _, input := range columnOptionValues {
			outputs = append(outputs, outputOptionParameters.KintoneOutputOptionColumnOptionInput{
				Name:       input.Name.ValueString(),
				FieldCode:  input.FieldCode.ValueString(),
				Type:       input.Type.ValueString(),
				Timezone:   input.Timezone.ValueStringPointer(),
				SortColumn: input.SortColumn.ValueStringPointer(),
			})
		}
		columnOptions = &outputs
	}

	return &outputOptionParameters.UpdateKintoneOutputOptionInput{
		KintoneConnectionID:              kintoneOutputOption.KintoneConnectionID.ValueInt64Pointer(),
		AppID:                            kintoneOutputOption.AppID.ValueStringPointer(),
		GuestSpaceID:                     kintoneOutputOption.GuestSpaceID.ValueStringPointer(),
		Mode:                             kintoneOutputOption.Mode.ValueStringPointer(),
		UpdateKey:                        kintoneOutputOption.UpdateKey.ValueStringPointer(),
		IgnoreNulls:                      kintoneOutputOption.IgnoreNulls.ValueBoolPointer(),
		ReduceKey:                        kintoneOutputOption.ReduceKey.ValueStringPointer(),
		ChunkSize:                        kintoneOutputOption.ChunkSize.ValueInt64Pointer(),
		KintoneOutputOptionColumnOptions: model.WrapObjectList(columnOptions),
	}
}
