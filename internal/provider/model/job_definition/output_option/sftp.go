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

type SftpOutputOption struct {
	SftpConnectionID       types.Int64  `tfsdk:"sftp_connection_id"`
	PathPrefix             types.String `tfsdk:"path_prefix"`
	FileExt                types.String `tfsdk:"file_ext"`
	IsMinimumOutputTasks   types.Bool   `tfsdk:"is_minimum_output_tasks"`
	EncoderType            types.String `tfsdk:"encoder_type"`
	CsvFormatter           types.Object `tfsdk:"csv_formatter"`
	JsonlFormatter         types.Object `tfsdk:"jsonl_formatter"`
	CustomVariableSettings types.List   `tfsdk:"custom_variable_settings"`
}

type sftpOutputOptionCsvFormatter struct {
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

type sftpOutputOptionCsvFormatterColumnOption struct {
	Name     types.String `tfsdk:"name"`
	Format   types.String `tfsdk:"format"`
	Timezone types.String `tfsdk:"timezone"`
}

type sftpOutputOptionJsonlFormatter struct {
	Encoding   types.String `tfsdk:"encoding"`
	Newline    types.String `tfsdk:"newline"`
	DateFormat types.String `tfsdk:"date_format"`
	Timezone   types.String `tfsdk:"timezone"`
}

func NewSftpOutputOption(ctx context.Context, sftpOutputOption *output_option.SftpOutputOption) *SftpOutputOption {
	if sftpOutputOption == nil {
		return nil
	}

	result := &SftpOutputOption{
		SftpConnectionID:     types.Int64PointerValue(sftpOutputOption.SftpConnectionID),
		PathPrefix:           types.StringValue(sftpOutputOption.PathPrefix),
		FileExt:              types.StringValue(sftpOutputOption.FileExt),
		IsMinimumOutputTasks: types.BoolValue(sftpOutputOption.IsMinimumOutputTasks),
		EncoderType:          types.StringValue(sftpOutputOption.EncoderType),
	}

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, sftpOutputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	// Read formatter data from the Formatter field if present, otherwise from top-level fields
	var csvFormatterData *output_option.SftpOutputOptionCsvFormatter
	var jsonlFormatterData *output_option.SftpOutputOptionJsonlFormatter

	if sftpOutputOption.Formatter != nil {
		// API response has formatter nested under Formatter field
		csvFormatterData = sftpOutputOption.Formatter.CsvFormatter
		jsonlFormatterData = sftpOutputOption.Formatter.JsonlFormatter
	} else {
		// Fallback to top-level fields (for older API versions or different endpoints)
		csvFormatterData = sftpOutputOption.CsvFormatter
		jsonlFormatterData = sftpOutputOption.JsonlFormatter
	}

	csvFormatter, err := newSftpOutputOptionCsvFormatter(ctx, csvFormatterData)
	if err != nil {
		return nil
	}
	result.CsvFormatter = csvFormatter

	jsonlFormatter, err := newSftpOutputOptionJsonlFormatter(ctx, jsonlFormatterData)
	if err != nil {
		return nil
	}
	result.JsonlFormatter = jsonlFormatter

	return result
}

func newSftpOutputOptionCsvFormatter(ctx context.Context, csvFormatter *output_option.SftpOutputOptionCsvFormatter) (types.Object, error) {
	objectType := types.ObjectType{
		AttrTypes: sftpOutputOptionCsvFormatter{}.attrTypes(),
	}

	if csvFormatter == nil {
		return types.ObjectNull(objectType.AttrTypes), nil
	}

	columnOptions, err := newSftpOutputOptionCsvFormatterColumnOptions(ctx, csvFormatter.CsvFormatterColumnOptionsAttributes)
	if err != nil {
		return types.ObjectNull(objectType.AttrTypes), err
	}

	formatter := sftpOutputOptionCsvFormatter{
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

func newSftpOutputOptionCsvFormatterColumnOptions(ctx context.Context, columnOptions *[]output_option.SftpOutputOptionCsvFormatterColumnOption) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: sftpOutputOptionCsvFormatterColumnOption{}.attrTypes(),
	}

	if columnOptions == nil {
		return types.ListNull(objectType), nil
	}

	options := make([]sftpOutputOptionCsvFormatterColumnOption, 0, len(*columnOptions))
	for _, opt := range *columnOptions {
		options = append(options, sftpOutputOptionCsvFormatterColumnOption{
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

// normalizeEncoding converts API encoding format to Terraform format
// API returns: utf_8, utf_16LE, utf_32BE, utf_32LE
// Terraform uses: UTF-8, UTF-16LE, UTF-32BE, UTF-32LE
func normalizeEncoding(apiEncoding string) string {
	switch apiEncoding {
	case "utf_8":
		return "UTF-8"
	case "utf_16LE":
		return "UTF-16LE"
	case "utf_32BE":
		return "UTF-32BE"
	case "utf_32LE":
		return "UTF-32LE"
	default:
		return apiEncoding
	}
}

// denormalizeEncoding converts Terraform encoding format to API format
// Terraform uses: UTF-8, UTF-16LE, UTF-32BE, UTF-32LE
// API expects: utf_8, utf_16LE, utf_32BE, utf_32LE
func denormalizeEncoding(terraformEncoding string) string {
	switch terraformEncoding {
	case "UTF-8":
		return "utf_8"
	case "UTF-16LE":
		return "utf_16LE"
	case "UTF-32BE":
		return "utf_32BE"
	case "UTF-32LE":
		return "utf_32LE"
	default:
		return terraformEncoding
	}
}

func newSftpOutputOptionJsonlFormatter(ctx context.Context, jsonlFormatter *output_option.SftpOutputOptionJsonlFormatter) (types.Object, error) {
	objectType := types.ObjectType{
		AttrTypes: sftpOutputOptionJsonlFormatter{}.attrTypes(),
	}

	if jsonlFormatter == nil {
		return types.ObjectNull(objectType.AttrTypes), nil
	}

	formatter := sftpOutputOptionJsonlFormatter{
		Encoding:   types.StringValue(normalizeEncoding(jsonlFormatter.Encoding)),
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

func (sftpOutputOptionCsvFormatter) attrTypes() map[string]attr.Type {
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
				AttrTypes: sftpOutputOptionCsvFormatterColumnOption{}.attrTypes(),
			},
		},
	}
}

func (sftpOutputOptionCsvFormatterColumnOption) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":     types.StringType,
		"format":   types.StringType,
		"timezone": types.StringType,
	}
}

func (sftpOutputOptionJsonlFormatter) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"encoding":    types.StringType,
		"newline":     types.StringType,
		"date_format": types.StringType,
		"timezone":    types.StringType,
	}
}

func (sftpOutputOption *SftpOutputOption) ToInput(ctx context.Context) *outputOptionParameters.SftpOutputOptionInput {
	if sftpOutputOption == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, sftpOutputOption.CustomVariableSettings)

	var csvFormatterInput *outputOptionParameters.SftpOutputOptionCsvFormatterInput
	if !sftpOutputOption.CsvFormatter.IsNull() && !sftpOutputOption.CsvFormatter.IsUnknown() {
		var csvFormatter sftpOutputOptionCsvFormatter
		diags := sftpOutputOption.CsvFormatter.As(ctx, &csvFormatter, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			return nil
		}
		csvFormatterInput = csvFormatter.toInput(ctx)
	}

	var jsonlFormatterInput *outputOptionParameters.SftpOutputOptionJsonlFormatterInput
	if !sftpOutputOption.JsonlFormatter.IsNull() && !sftpOutputOption.JsonlFormatter.IsUnknown() {
		var jsonlFormatter sftpOutputOptionJsonlFormatter
		diags := sftpOutputOption.JsonlFormatter.As(ctx, &jsonlFormatter, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			return nil
		}
		jsonlFormatterInput = jsonlFormatter.toInput(ctx)
	}

	// Auto-detect formatter_type based on which formatter is provided
	var formatterType string
	if csvFormatterInput != nil {
		formatterType = "csv"
	} else if jsonlFormatterInput != nil {
		formatterType = "jsonl"
	}

	return &outputOptionParameters.SftpOutputOptionInput{
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
		SftpConnectionID:       sftpOutputOption.SftpConnectionID.ValueInt64(),
		PathPrefix:             sftpOutputOption.PathPrefix.ValueString(),
		FileExt:                sftpOutputOption.FileExt.ValueString(),
		IsMinimumOutputTasks:   sftpOutputOption.IsMinimumOutputTasks.ValueBool(),
		FormatterType:          formatterType,
		EncoderType:            sftpOutputOption.EncoderType.ValueString(),
		CsvFormatter:           csvFormatterInput,
		JsonlFormatter:         jsonlFormatterInput,
	}
}

func (sftpOutputOption *SftpOutputOption) ToUpdateInput(ctx context.Context) *outputOptionParameters.UpdateSftpOutputOptionInput {
	if sftpOutputOption == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, sftpOutputOption.CustomVariableSettings)

	var csvFormatterInput *outputOptionParameters.SftpOutputOptionCsvFormatterInput
	if !sftpOutputOption.CsvFormatter.IsNull() {
		var csvFormatter sftpOutputOptionCsvFormatter
		if !sftpOutputOption.CsvFormatter.IsUnknown() {
			diags := sftpOutputOption.CsvFormatter.As(ctx, &csvFormatter, basetypes.ObjectAsOptions{})
			if diags.HasError() {
				return nil
			}
		}
		csvFormatterInput = csvFormatter.toInput(ctx)
	}

	var jsonlFormatterInput *outputOptionParameters.SftpOutputOptionJsonlFormatterInput
	if !sftpOutputOption.JsonlFormatter.IsNull() {
		var jsonlFormatter sftpOutputOptionJsonlFormatter
		if !sftpOutputOption.JsonlFormatter.IsUnknown() {
			diags := sftpOutputOption.JsonlFormatter.As(ctx, &jsonlFormatter, basetypes.ObjectAsOptions{})
			if diags.HasError() {
				return nil
			}
		}
		jsonlFormatterInput = jsonlFormatter.toInput(ctx)
	}

	// Auto-detect formatter_type based on which formatter is provided
	var formatterType *string
	if csvFormatterInput != nil {
		formatterTypeStr := "csv"
		formatterType = &formatterTypeStr
	} else if jsonlFormatterInput != nil {
		formatterTypeStr := "jsonl"
		formatterType = &formatterTypeStr
	}

	return &outputOptionParameters.UpdateSftpOutputOptionInput{
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
		SftpConnectionID:       sftpOutputOption.SftpConnectionID.ValueInt64Pointer(),
		PathPrefix:             sftpOutputOption.PathPrefix.ValueStringPointer(),
		FileExt:                sftpOutputOption.FileExt.ValueStringPointer(),
		IsMinimumOutputTasks:   sftpOutputOption.IsMinimumOutputTasks.ValueBoolPointer(),
		FormatterType:          formatterType,
		EncoderType:            sftpOutputOption.EncoderType.ValueStringPointer(),
		CsvFormatter:           csvFormatterInput,
		JsonlFormatter:         jsonlFormatterInput,
	}
}

func (csvFormatter *sftpOutputOptionCsvFormatter) toInput(ctx context.Context) *outputOptionParameters.SftpOutputOptionCsvFormatterInput {
	if csvFormatter == nil {
		return nil
	}

	var columnOptions *[]outputOptionParameters.SftpOutputOptionCsvFormatterColumnOptionInput
	if !csvFormatter.CsvFormatterColumnOptionsAttributes.IsNull() && !csvFormatter.CsvFormatterColumnOptionsAttributes.IsUnknown() {
		var columnOptionValues []sftpOutputOptionCsvFormatterColumnOption
		diags := csvFormatter.CsvFormatterColumnOptionsAttributes.ElementsAs(ctx, &columnOptionValues, false)
		if diags.HasError() {
			return nil
		}

		opts := make([]outputOptionParameters.SftpOutputOptionCsvFormatterColumnOptionInput, 0, len(columnOptionValues))
		for _, opt := range columnOptionValues {
			opts = append(opts, outputOptionParameters.SftpOutputOptionCsvFormatterColumnOptionInput{
				Name:     opt.Name.ValueString(),
				Format:   opt.Format.ValueString(),
				Timezone: model.NewNullableString(opt.Timezone),
			})
		}
		columnOptions = &opts
	}

	return &outputOptionParameters.SftpOutputOptionCsvFormatterInput{
		Delimiter:                           csvFormatter.Delimiter.ValueString(),
		Newline:                             csvFormatter.Newline.ValueString(),
		NewlineInField:                      csvFormatter.NewlineInField.ValueString(),
		Charset:                             csvFormatter.Charset.ValueString(),
		QuotePolicy:                         csvFormatter.QuotePolicy.ValueString(),
		Escape:                              csvFormatter.Escape.ValueString(),
		HeaderLine:                          csvFormatter.HeaderLine.ValueBool(),
		NullStringEnabled:                   csvFormatter.NullStringEnabled.ValueBool(),
		NullString:                          model.NewNullableString(csvFormatter.NullString),
		DefaultTimeZone:                     csvFormatter.DefaultTimeZone.ValueString(),
		CsvFormatterColumnOptionsAttributes: columnOptions,
	}
}

func (jsonlFormatter *sftpOutputOptionJsonlFormatter) toInput(ctx context.Context) *outputOptionParameters.SftpOutputOptionJsonlFormatterInput {
	if jsonlFormatter == nil {
		return nil
	}

	return &outputOptionParameters.SftpOutputOptionJsonlFormatterInput{
		Encoding:   denormalizeEncoding(jsonlFormatter.Encoding.ValueString()),
		Newline:    jsonlFormatter.Newline.ValueString(),
		DateFormat: model.NewNullableString(jsonlFormatter.DateFormat),
		Timezone:   model.NewNullableString(jsonlFormatter.Timezone),
	}
}
