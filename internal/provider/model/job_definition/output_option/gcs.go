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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type GcsOutputOption struct {
	GcsConnectionID        types.Int64  `tfsdk:"gcs_connection_id"`
	Bucket                 types.String `tfsdk:"bucket"`
	PathPrefix             types.String `tfsdk:"path_prefix"`
	FileExt                types.String `tfsdk:"file_ext"`
	SequenceFormat         types.String `tfsdk:"sequence_format"`
	IsMinimumOutputTasks   types.Bool   `tfsdk:"is_minimum_output_tasks"`
	EncoderType            types.String `tfsdk:"encoder_type"`
	FormatterType          types.String `tfsdk:"formatter_type"`
	CsvFormatter           types.Object `tfsdk:"csv_formatter"`
	JsonlFormatter         types.Object `tfsdk:"jsonl_formatter"`
	CustomVariableSettings types.List   `tfsdk:"custom_variable_settings"`
}

type gcsOutputOptionCsvFormatter struct {
	Delimiter                           types.String `tfsdk:"delimiter"`
	Newline                             types.String `tfsdk:"newline"`
	NewlineInField                      types.String `tfsdk:"newline_in_field"`
	Charset                             types.String `tfsdk:"charset"`
	QuotePolicy                         types.String `tfsdk:"quote_policy"`
	Escape                              types.String `tfsdk:"escape"`
	HeaderLine                          types.Bool   `tfsdk:"header_line"`
	NullStringEnabled                   types.Bool   `tfsdk:"null_string_enabled"`
	NullString                          types.String `tfsdk:"null_string"`
	DefaultTimeZone                     types.String `tfsdk:"default_time_zone"`
	CsvFormatterColumnOptionsAttributes types.List   `tfsdk:"csv_formatter_column_options_attributes"`
}

type gcsOutputOptionCsvFormatterColumnOption struct {
	Name     types.String `tfsdk:"name"`
	Format   types.String `tfsdk:"format"`
	Timezone types.String `tfsdk:"timezone"`
}

type gcsOutputOptionJsonlFormatter struct {
	Encoding   types.String `tfsdk:"encoding"`
	Newline    types.String `tfsdk:"newline"`
	DateFormat types.String `tfsdk:"date_format"`
	Timezone   types.String `tfsdk:"timezone"`
}

func NewGcsOutputOption(ctx context.Context, gcsOutputOption *output_option.GcsOutputOption) *GcsOutputOption {
	if gcsOutputOption == nil {
		return nil
	}

	result := &GcsOutputOption{
		GcsConnectionID:      types.Int64PointerValue(gcsOutputOption.GcsConnectionID),
		Bucket:               types.StringValue(gcsOutputOption.Bucket),
		PathPrefix:           types.StringValue(gcsOutputOption.PathPrefix),
		FileExt:              types.StringValue(gcsOutputOption.FileExt),
		SequenceFormat:       types.StringPointerValue(gcsOutputOption.SequenceFormat),
		IsMinimumOutputTasks: types.BoolValue(gcsOutputOption.IsMinimumOutputTasks),
		EncoderType:          types.StringValue(gcsOutputOption.EncoderType),
	}

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, gcsOutputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	if gcsOutputOption.Formatter != nil {
		result.FormatterType = types.StringValue(gcsOutputOption.Formatter.Type)

		csvFormatter, err := newGcsOutputOptionCsvFormatter(ctx, gcsOutputOption.Formatter.CsvFormatter)
		if err != nil {
			return nil
		}
		result.CsvFormatter = csvFormatter

		jsonlFormatter, err := newGcsOutputOptionJsonlFormatter(ctx, gcsOutputOption.Formatter.JsonlFormatter)
		if err != nil {
			return nil
		}
		result.JsonlFormatter = jsonlFormatter
	} else {
		result.FormatterType = types.StringNull()
		result.CsvFormatter = types.ObjectNull(gcsOutputOptionCsvFormatter{}.attrTypes())
		result.JsonlFormatter = types.ObjectNull(gcsOutputOptionJsonlFormatter{}.attrTypes())
	}

	return result
}

func newGcsOutputOptionCsvFormatter(ctx context.Context, csvFormatter *output_option.CsvFormatter) (types.Object, error) {
	objectType := types.ObjectType{
		AttrTypes: gcsOutputOptionCsvFormatter{}.attrTypes(),
	}

	if csvFormatter == nil {
		return types.ObjectNull(objectType.AttrTypes), nil
	}

	columnOptions, err := newGcsOutputOptionCsvFormatterColumnOptions(ctx, csvFormatter.CsvFormatterColumnOptionsAttributes)
	if err != nil {
		return types.ObjectNull(objectType.AttrTypes), err
	}

	formatter := gcsOutputOptionCsvFormatter{
		Delimiter:                           types.StringValue(csvFormatter.Delimiter),
		Newline:                             types.StringValue(csvFormatter.Newline),
		NewlineInField:                      types.StringValue(csvFormatter.NewlineInField),
		Charset:                             types.StringValue(csvFormatter.Charset),
		QuotePolicy:                         types.StringValue(csvFormatter.QuotePolicy),
		Escape:                              types.StringValue(csvFormatter.Escape),
		HeaderLine:                          types.BoolValue(csvFormatter.HeaderLine),
		NullStringEnabled:                   types.BoolValue(csvFormatter.NullStringEnabled),
		NullString:                          types.StringPointerValue(csvFormatter.NullString),
		DefaultTimeZone:                     types.StringValue(csvFormatter.DefaultTimeZone),
		CsvFormatterColumnOptionsAttributes: columnOptions,
	}

	objectValue, diags := types.ObjectValueFrom(ctx, objectType.AttrTypes, formatter)
	if diags.HasError() {
		return types.ObjectNull(objectType.AttrTypes), fmt.Errorf("failed to convert csv formatter to ObjectValue: %v", diags)
	}

	return objectValue, nil
}

func newGcsOutputOptionCsvFormatterColumnOptions(ctx context.Context, columnOptions *[]output_option.CsvFormatterColumnOption) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: gcsOutputOptionCsvFormatterColumnOption{}.attrTypes(),
	}

	if columnOptions == nil {
		return types.ListNull(objectType), nil
	}

	options := make([]gcsOutputOptionCsvFormatterColumnOption, 0, len(*columnOptions))
	for _, opt := range *columnOptions {
		options = append(options, gcsOutputOptionCsvFormatterColumnOption{
			Name:     types.StringValue(opt.Name),
			Format:   types.StringValue(opt.Format),
			Timezone: types.StringPointerValue(opt.Timezone),
		})
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, options)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert column options to ListValue: %v", diags)
	}

	return listValue, nil
}

func newGcsOutputOptionJsonlFormatter(ctx context.Context, jsonlFormatter *output_option.JsonlFormatter) (types.Object, error) {
	objectType := types.ObjectType{
		AttrTypes: gcsOutputOptionJsonlFormatter{}.attrTypes(),
	}

	if jsonlFormatter == nil {
		return types.ObjectNull(objectType.AttrTypes), nil
	}

	formatter := gcsOutputOptionJsonlFormatter{
		Encoding:   types.StringValue(common.NormalizeEncoding(jsonlFormatter.Encoding)),
		Newline:    types.StringValue(jsonlFormatter.Newline),
		DateFormat: types.StringPointerValue(jsonlFormatter.DateFormat),
		Timezone:   types.StringPointerValue(jsonlFormatter.Timezone),
	}

	objectValue, diags := types.ObjectValueFrom(ctx, objectType.AttrTypes, formatter)
	if diags.HasError() {
		return types.ObjectNull(objectType.AttrTypes), fmt.Errorf("failed to convert jsonl formatter to ObjectValue: %v", diags)
	}

	return objectValue, nil
}

func (gcsOutputOptionCsvFormatter) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"delimiter":           types.StringType,
		"newline":             types.StringType,
		"newline_in_field":    types.StringType,
		"charset":             types.StringType,
		"quote_policy":        types.StringType,
		"escape":              types.StringType,
		"header_line":         types.BoolType,
		"null_string_enabled": types.BoolType,
		"null_string":         types.StringType,
		"default_time_zone":   types.StringType,
		"csv_formatter_column_options_attributes": types.ListType{
			ElemType: types.ObjectType{
				AttrTypes: gcsOutputOptionCsvFormatterColumnOption{}.attrTypes(),
			},
		},
	}
}

func (gcsOutputOptionCsvFormatterColumnOption) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":     types.StringType,
		"format":   types.StringType,
		"timezone": types.StringType,
	}
}

func (gcsOutputOptionJsonlFormatter) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"encoding":    types.StringType,
		"newline":     types.StringType,
		"date_format": types.StringType,
		"timezone":    types.StringType,
	}
}

func (gcsOutputOption *GcsOutputOption) ToInput(ctx context.Context) *outputOptionParameters.GcsOutputOptionInput {
	if gcsOutputOption == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, gcsOutputOption.CustomVariableSettings)

	var csvFormatterInput *outputOptionParameters.GcsOutputOptionCsvFormatterInput
	if !gcsOutputOption.CsvFormatter.IsNull() && !gcsOutputOption.CsvFormatter.IsUnknown() {
		var csvFormatter gcsOutputOptionCsvFormatter
		diags := gcsOutputOption.CsvFormatter.As(ctx, &csvFormatter, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			return nil
		}
		csvFormatterInput = csvFormatter.toInput(ctx)
	}

	var jsonlFormatterInput *outputOptionParameters.GcsOutputOptionJsonlFormatterInput
	if !gcsOutputOption.JsonlFormatter.IsNull() && !gcsOutputOption.JsonlFormatter.IsUnknown() {
		var jsonlFormatter gcsOutputOptionJsonlFormatter
		diags := gcsOutputOption.JsonlFormatter.As(ctx, &jsonlFormatter, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			return nil
		}
		jsonlFormatterInput = jsonlFormatter.toInput()
	}

	return &outputOptionParameters.GcsOutputOptionInput{
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
		GcsConnectionID:        gcsOutputOption.GcsConnectionID.ValueInt64(),
		Bucket:                 gcsOutputOption.Bucket.ValueString(),
		PathPrefix:             gcsOutputOption.PathPrefix.ValueString(),
		FileExt:                gcsOutputOption.FileExt.ValueString(),
		SequenceFormat:         model.NewNullableString(gcsOutputOption.SequenceFormat),
		IsMinimumOutputTasks:   gcsOutputOption.IsMinimumOutputTasks.ValueBool(),
		FormatterType:          gcsOutputOption.FormatterType.ValueString(),
		EncoderType:            gcsOutputOption.EncoderType.ValueString(),
		CsvFormatter:           csvFormatterInput,
		JsonlFormatter:         jsonlFormatterInput,
	}
}

func (gcsOutputOption *GcsOutputOption) ToUpdateInput(ctx context.Context) *outputOptionParameters.UpdateGcsOutputOptionInput {
	if gcsOutputOption == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, gcsOutputOption.CustomVariableSettings)

	var csvFormatterInput *outputOptionParameters.GcsOutputOptionCsvFormatterInput
	if !gcsOutputOption.CsvFormatter.IsNull() {
		var csvFormatter gcsOutputOptionCsvFormatter
		if !gcsOutputOption.CsvFormatter.IsUnknown() {
			diags := gcsOutputOption.CsvFormatter.As(ctx, &csvFormatter, basetypes.ObjectAsOptions{})
			if diags.HasError() {
				return nil
			}
		}
		csvFormatterInput = csvFormatter.toInput(ctx)
	}

	var jsonlFormatterInput *outputOptionParameters.GcsOutputOptionJsonlFormatterInput
	if !gcsOutputOption.JsonlFormatter.IsNull() {
		var jsonlFormatter gcsOutputOptionJsonlFormatter
		if !gcsOutputOption.JsonlFormatter.IsUnknown() {
			diags := gcsOutputOption.JsonlFormatter.As(ctx, &jsonlFormatter, basetypes.ObjectAsOptions{})
			if diags.HasError() {
				return nil
			}
		}
		jsonlFormatterInput = jsonlFormatter.toInput()
	}

	return &outputOptionParameters.UpdateGcsOutputOptionInput{
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
		GcsConnectionID:        gcsOutputOption.GcsConnectionID.ValueInt64Pointer(),
		Bucket:                 gcsOutputOption.Bucket.ValueStringPointer(),
		PathPrefix:             gcsOutputOption.PathPrefix.ValueStringPointer(),
		FileExt:                gcsOutputOption.FileExt.ValueStringPointer(),
		SequenceFormat:         model.NewNullableString(gcsOutputOption.SequenceFormat),
		IsMinimumOutputTasks:   gcsOutputOption.IsMinimumOutputTasks.ValueBoolPointer(),
		FormatterType:          gcsOutputOption.FormatterType.ValueStringPointer(),
		EncoderType:            gcsOutputOption.EncoderType.ValueStringPointer(),
		CsvFormatter:           csvFormatterInput,
		JsonlFormatter:         jsonlFormatterInput,
	}
}

func (csvFormatter *gcsOutputOptionCsvFormatter) toInput(ctx context.Context) *outputOptionParameters.GcsOutputOptionCsvFormatterInput {
	if csvFormatter == nil {
		return nil
	}

	var columnOptions *[]outputOptionParameters.GcsOutputOptionCsvFormatterColumnOptionInput
	if !csvFormatter.CsvFormatterColumnOptionsAttributes.IsNull() && !csvFormatter.CsvFormatterColumnOptionsAttributes.IsUnknown() {
		var columnOptionValues []gcsOutputOptionCsvFormatterColumnOption
		diags := csvFormatter.CsvFormatterColumnOptionsAttributes.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}

		opts := make([]outputOptionParameters.GcsOutputOptionCsvFormatterColumnOptionInput, 0, len(columnOptionValues))
		for _, opt := range columnOptionValues {
			opts = append(opts, outputOptionParameters.GcsOutputOptionCsvFormatterColumnOptionInput{
				Name:     opt.Name.ValueString(),
				Format:   opt.Format.ValueString(),
				Timezone: opt.Timezone.ValueStringPointer(),
			})
		}
		columnOptions = &opts
	}

	return &outputOptionParameters.GcsOutputOptionCsvFormatterInput{
		Delimiter:                           csvFormatter.Delimiter.ValueString(),
		Newline:                             csvFormatter.Newline.ValueString(),
		NewlineInField:                      csvFormatter.NewlineInField.ValueString(),
		Charset:                             csvFormatter.Charset.ValueString(),
		QuotePolicy:                         csvFormatter.QuotePolicy.ValueString(),
		Escape:                              csvFormatter.Escape.ValueString(),
		HeaderLine:                          csvFormatter.HeaderLine.ValueBool(),
		NullStringEnabled:                   csvFormatter.NullStringEnabled.ValueBool(),
		NullString:                          csvFormatter.NullString.ValueStringPointer(),
		DefaultTimeZone:                     csvFormatter.DefaultTimeZone.ValueString(),
		CsvFormatterColumnOptionsAttributes: columnOptions,
	}
}

func (jsonlFormatter *gcsOutputOptionJsonlFormatter) toInput() *outputOptionParameters.GcsOutputOptionJsonlFormatterInput {
	if jsonlFormatter == nil {
		return nil
	}

	return &outputOptionParameters.GcsOutputOptionJsonlFormatterInput{
		Encoding:   common.DenormalizeEncoding(jsonlFormatter.Encoding.ValueString()),
		Newline:    jsonlFormatter.Newline.ValueString(),
		DateFormat: jsonlFormatter.DateFormat.ValueStringPointer(),
		Timezone:   jsonlFormatter.Timezone.ValueStringPointer(),
	}
}
