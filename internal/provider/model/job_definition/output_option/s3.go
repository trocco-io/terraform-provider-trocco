package output_options

import (
	"context"
	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	outputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type S3OutputOption struct {
	S3ConnectionID         types.Int64     `tfsdk:"s3_connection_id"`
	Bucket                 types.String    `tfsdk:"bucket"`
	PathPrefix             types.String    `tfsdk:"path_prefix"`
	Region                 types.String    `tfsdk:"region"`
	FileExt                types.String    `tfsdk:"file_ext"`
	SequenceFormat         types.String    `tfsdk:"sequence_format"`
	CannedAcl              types.String    `tfsdk:"canned_acl"`
	IsMinimumOutputTasks   types.Bool      `tfsdk:"is_minimum_output_tasks"`
	MultipartUploadEnabled types.Bool      `tfsdk:"multipart_upload_enabled"`
	FormatterType          types.String    `tfsdk:"formatter_type"`
	EncoderType            types.String    `tfsdk:"encoder_type"`
	CsvFormatter           *CsvFormatter   `tfsdk:"csv_formatter"`
	JsonlFormatter         *JsonlFormatter `tfsdk:"jsonl_formatter"`
	CustomVariableSettings types.List      `tfsdk:"custom_variable_settings"`
}

type CsvFormatter struct {
	Delimiter                           types.String `tfsdk:"delimiter"`
	Escape                              types.String `tfsdk:"escape"`
	HeaderLine                          types.Bool   `tfsdk:"header_line"`
	Charset                             types.String `tfsdk:"charset"`
	QuotePolicy                         types.String `tfsdk:"quote_policy"`
	Newline                             types.String `tfsdk:"newline"`
	NewlineInField                      types.String `tfsdk:"newline_in_field"`
	NullStringEnabled                   types.Bool   `tfsdk:"null_string_enabled"`
	NullString                          types.String `tfsdk:"null_string"`
	DefaultTimeZone                     types.String `tfsdk:"default_time_zone"`
	CsvFormatterColumnOptionsAttributes types.List   `tfsdk:"csv_formatter_column_options_attributes"`
}

type CsvFormatterColumnOption struct {
	Name     types.String `tfsdk:"name"`
	Format   types.String `tfsdk:"format"`
	Timezone types.String `tfsdk:"timezone"`
}

type JsonlFormatter struct {
	Encoding   types.String `tfsdk:"encoding"`
	Newline    types.String `tfsdk:"newline"`
	DateFormat types.String `tfsdk:"date_format"`
	Timezone   types.String `tfsdk:"timezone"`
}

func NewS3OutputOption(ctx context.Context, s3OutputOption *output_option.S3OutputOption) *S3OutputOption {
	if s3OutputOption == nil {
		return nil
	}

	result := &S3OutputOption{
		S3ConnectionID:         types.Int64Value(s3OutputOption.S3ConnectionID),
		Bucket:                 types.StringValue(s3OutputOption.Bucket),
		PathPrefix:             types.StringValue(s3OutputOption.PathPrefix),
		Region:                 types.StringValue(s3OutputOption.Region),
		FileExt:                types.StringValue(s3OutputOption.FileExt),
		SequenceFormat:         types.StringValue(s3OutputOption.SequenceFormat),
		CannedAcl:              types.StringValue(s3OutputOption.CannedAcl),
		IsMinimumOutputTasks:   types.BoolValue(s3OutputOption.IsMinimumOutputTasks),
		MultipartUploadEnabled: types.BoolValue(s3OutputOption.MultipartUploadEnabled),
		EncoderType:            types.StringPointerValue(s3OutputOption.EncoderType),
	}

	// Convert custom variable settings
	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, s3OutputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	// Handle formatter conversion: Entity has nested structure, Model has flat formatter_type
	if s3OutputOption.Formatter != nil {
		result.FormatterType = types.StringValue(s3OutputOption.Formatter.Type)
		result.CsvFormatter = NewCsvFormatter(ctx, s3OutputOption.Formatter.CsvFormatter)
		result.JsonlFormatter = NewJsonlFormatter(ctx, s3OutputOption.Formatter.JsonlFormatter)
	} else {
		result.FormatterType = types.StringNull()
		result.CsvFormatter = nil
		result.JsonlFormatter = nil
	}

	return result
}

func NewCsvFormatter(ctx context.Context, csvFormatter *output_option.CsvFormatter) *CsvFormatter {
	if csvFormatter == nil {
		return nil
	}

	columnElements := make([]CsvFormatterColumnOption, 0)
	if csvFormatter.CsvFormatterColumnOptionsAttributes != nil {
		for _, opt := range *csvFormatter.CsvFormatterColumnOptionsAttributes {
			option := CsvFormatterColumnOption{
				Name:     types.StringValue(opt.Name),
				Format:   types.StringValue(opt.Format),
				Timezone: types.StringPointerValue(opt.Timezone),
			}
			columnElements = append(columnElements, option)
		}
	}

	columnOptions, diags := types.ListValueFrom(
		ctx,
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":     types.StringType,
				"format":   types.StringType,
				"timezone": types.StringType,
			},
		},
		columnElements,
	)
	if diags.HasError() {
		return nil
	}

	return &CsvFormatter{
		Delimiter:                           types.StringValue(csvFormatter.Delimiter),
		Escape:                              types.StringValue(csvFormatter.Escape),
		HeaderLine:                          types.BoolValue(csvFormatter.HeaderLine),
		Charset:                             types.StringValue(csvFormatter.Charset),
		QuotePolicy:                         types.StringValue(csvFormatter.QuotePolicy),
		Newline:                             types.StringValue(csvFormatter.Newline),
		NewlineInField:                      types.StringValue(csvFormatter.NewlineInField),
		NullStringEnabled:                   types.BoolValue(csvFormatter.NullStringEnabled),
		NullString:                          types.StringPointerValue(csvFormatter.NullString),
		DefaultTimeZone:                     types.StringValue(csvFormatter.DefaultTimeZone),
		CsvFormatterColumnOptionsAttributes: columnOptions,
	}
}

func NewJsonlFormatter(ctx context.Context, jsonlFormatter *output_option.JsonlFormatter) *JsonlFormatter {
	if jsonlFormatter == nil {
		return nil
	}

	return &JsonlFormatter{
		Encoding:   types.StringValue(jsonlFormatter.Encoding),
		Newline:    types.StringValue(jsonlFormatter.Newline),
		DateFormat: types.StringPointerValue(jsonlFormatter.DateFormat),
		Timezone:   types.StringPointerValue(jsonlFormatter.Timezone),
	}
}

func (s3OutputOption *S3OutputOption) ToInput(ctx context.Context) *outputOptionParameters.S3OutputOptionInput {
	if s3OutputOption == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, s3OutputOption.CustomVariableSettings)

	return &outputOptionParameters.S3OutputOptionInput{
		S3ConnectionID:         s3OutputOption.S3ConnectionID.ValueInt64(),
		Bucket:                 s3OutputOption.Bucket.ValueString(),
		PathPrefix:             s3OutputOption.PathPrefix.ValueString(),
		Region:                 s3OutputOption.Region.ValueString(),
		FileExt:                s3OutputOption.FileExt.ValueString(),
		SequenceFormat:         s3OutputOption.SequenceFormat.ValueString(),
		CannedAcl:              s3OutputOption.CannedAcl.ValueString(),
		IsMinimumOutputTasks:   s3OutputOption.IsMinimumOutputTasks.ValueBool(),
		MultipartUploadEnabled: s3OutputOption.MultipartUploadEnabled.ValueBool(),
		FormatterType:          s3OutputOption.FormatterType.ValueString(),
		EncoderType:            s3OutputOption.EncoderType.ValueString(),
		CsvFormatter:           s3OutputOption.CsvFormatter.ToCsvFormatterInput(ctx),
		JsonlFormatter:         s3OutputOption.JsonlFormatter.ToJsonlFormatterInput(ctx),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (s3OutputOption *S3OutputOption) ToUpdateInput(ctx context.Context) *outputOptionParameters.UpdateS3OutputOptionInput {
	if s3OutputOption == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, s3OutputOption.CustomVariableSettings)

	return &outputOptionParameters.UpdateS3OutputOptionInput{
		S3ConnectionID:         s3OutputOption.S3ConnectionID.ValueInt64Pointer(),
		Bucket:                 s3OutputOption.Bucket.ValueStringPointer(),
		PathPrefix:             s3OutputOption.PathPrefix.ValueStringPointer(),
		Region:                 s3OutputOption.Region.ValueStringPointer(),
		FileExt:                s3OutputOption.FileExt.ValueStringPointer(),
		SequenceFormat:         s3OutputOption.SequenceFormat.ValueStringPointer(),
		CannedAcl:              s3OutputOption.CannedAcl.ValueStringPointer(),
		IsMinimumOutputTasks:   s3OutputOption.IsMinimumOutputTasks.ValueBoolPointer(),
		MultipartUploadEnabled: s3OutputOption.MultipartUploadEnabled.ValueBoolPointer(),
		FormatterType:          s3OutputOption.FormatterType.ValueStringPointer(),
		EncoderType:            s3OutputOption.EncoderType.ValueStringPointer(),
		CsvFormatter:           s3OutputOption.CsvFormatter.ToCsvFormatterInput(ctx),
		JsonlFormatter:         s3OutputOption.JsonlFormatter.ToJsonlFormatterInput(ctx),
		CustomVariableSettings: model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (csvFormatter *CsvFormatter) ToCsvFormatterInput(ctx context.Context) *outputOptionParameters.CsvFormatterInput {
	if csvFormatter == nil {
		return nil
	}

	input := &outputOptionParameters.CsvFormatterInput{
		Delimiter:         csvFormatter.Delimiter.ValueString(),
		Escape:            csvFormatter.Escape.ValueString(),
		HeaderLine:        csvFormatter.HeaderLine.ValueBool(),
		Charset:           csvFormatter.Charset.ValueString(),
		QuotePolicy:       csvFormatter.QuotePolicy.ValueString(),
		Newline:           csvFormatter.Newline.ValueString(),
		NewlineInField:    csvFormatter.NewlineInField.ValueString(),
		NullStringEnabled: csvFormatter.NullStringEnabled.ValueBool(),
		NullString:        csvFormatter.NullString.ValueStringPointer(),
		DefaultTimeZone:   csvFormatter.DefaultTimeZone.ValueString(),
	}

	// Convert column options
	if !csvFormatter.CsvFormatterColumnOptionsAttributes.IsNull() && !csvFormatter.CsvFormatterColumnOptionsAttributes.IsUnknown() {
		var columnOptions []CsvFormatterColumnOption
		diags := csvFormatter.CsvFormatterColumnOptionsAttributes.ElementsAs(ctx, &columnOptions, false)
		if diags.HasError() {
			return input
		}

		columnInputs := make([]outputOptionParameters.CsvFormatterColumnOptionInput, 0, len(columnOptions))
		for _, opt := range columnOptions {
			columnInputs = append(columnInputs, outputOptionParameters.CsvFormatterColumnOptionInput{
				Name:     opt.Name.ValueString(),
				Format:   opt.Format.ValueString(),
				Timezone: opt.Timezone.ValueStringPointer(),
			})
		}
		input.CsvFormatterColumnOptionsAttributes = &columnInputs
	}

	return input
}

func (jsonlFormatter *JsonlFormatter) ToJsonlFormatterInput(ctx context.Context) *outputOptionParameters.JsonlFormatterInput {
	if jsonlFormatter == nil {
		return nil
	}

	return &outputOptionParameters.JsonlFormatterInput{
		Encoding:   jsonlFormatter.Encoding.ValueString(),
		Newline:    jsonlFormatter.Newline.ValueString(),
		DateFormat: jsonlFormatter.DateFormat.ValueStringPointer(),
		Timezone:   jsonlFormatter.Timezone.ValueStringPointer(),
	}
}
