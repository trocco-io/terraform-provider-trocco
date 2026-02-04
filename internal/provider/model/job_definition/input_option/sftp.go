package input_options

import (
	"context"
	inputOptionEntities "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	inputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"
	"terraform-provider-trocco/internal/provider/model/job_definition/input_option/parser"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type SftpInputOption struct {
	SftpConnectionID          types.Int64            `tfsdk:"sftp_connection_id"`
	PathPrefix                types.String           `tfsdk:"path_prefix"`
	PathMatchPattern          types.String           `tfsdk:"path_match_pattern"`
	IncrementalLoadingEnabled types.Bool             `tfsdk:"incremental_loading_enabled"`
	LastPath                  types.String           `tfsdk:"last_path"`
	StopWhenFileNotFound      types.Bool             `tfsdk:"stop_when_file_not_found"`
	DecompressionType         types.String           `tfsdk:"decompression_type"`
	Decoder                   *Decoder               `tfsdk:"decoder"`
	CustomVariableSettings    types.List             `tfsdk:"custom_variable_settings"`
	CsvParser                 *parser.CsvParser      `tfsdk:"csv_parser"`
	JsonlParser               *parser.JsonlParser    `tfsdk:"jsonl_parser"`
	JsonpathParser            *parser.JsonpathParser `tfsdk:"jsonpath_parser"`
	LtsvParser                *parser.LtsvParser     `tfsdk:"ltsv_parser"`
	ExcelParser               *parser.ExcelParser    `tfsdk:"excel_parser"`
	XmlParser                 *parser.XmlParser      `tfsdk:"xml_parser"`
}

func NewSftpInputOption(ctx context.Context, sftpInputOption *inputOptionEntities.SftpInputOption) *SftpInputOption {
	if sftpInputOption == nil {
		return nil
	}

	result := &SftpInputOption{
		SftpConnectionID:          types.Int64Value(sftpInputOption.SftpConnectionID),
		PathPrefix:                types.StringValue(sftpInputOption.PathPrefix),
		PathMatchPattern:          types.StringPointerValue(sftpInputOption.PathMatchPattern),
		IncrementalLoadingEnabled: types.BoolValue(sftpInputOption.IncrementalLoadingEnabled),
		LastPath:                  types.StringPointerValue(sftpInputOption.LastPath),
		StopWhenFileNotFound:      types.BoolValue(sftpInputOption.StopWhenFileNotFound),
		DecompressionType:         types.StringValue(sftpInputOption.DecompressionType),
		Decoder:                   NewDecoder(sftpInputOption.Decoder),
		CsvParser:                 parser.NewCsvParser(ctx, sftpInputOption.CsvParser),
		JsonlParser:               parser.NewJsonlParser(ctx, sftpInputOption.JsonlParser),
		JsonpathParser:            parser.NewJsonPathParser(ctx, sftpInputOption.JsonpathParser),
		LtsvParser:                parser.NewLtsvParser(ctx, sftpInputOption.LtsvParser),
		ExcelParser:               parser.NewExcelParser(ctx, sftpInputOption.ExcelParser),
		XmlParser:                 parser.NewXmlParser(ctx, sftpInputOption.XmlParser),
	}

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, sftpInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func (sftpInputOption *SftpInputOption) ToInput(ctx context.Context) *inputOptionParameters.SftpInputOptionInput {
	if sftpInputOption == nil {
		return nil
	}
	customVarSettings := common.ExtractCustomVariableSettings(ctx, sftpInputOption.CustomVariableSettings)

	return &inputOptionParameters.SftpInputOptionInput{
		SftpConnectionID:          sftpInputOption.SftpConnectionID.ValueInt64(),
		PathPrefix:                sftpInputOption.PathPrefix.ValueString(),
		PathMatchPattern:          model.NewNullableString(sftpInputOption.PathMatchPattern),
		IncrementalLoadingEnabled: sftpInputOption.IncrementalLoadingEnabled.ValueBool(),
		LastPath:                  model.NewNullableString(sftpInputOption.LastPath),
		StopWhenFileNotFound:      sftpInputOption.StopWhenFileNotFound.ValueBool(),
		DecompressionType:         sftpInputOption.DecompressionType.ValueString(),
		Decoder:                   sftpInputOption.Decoder.ToDecoderInput(),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
		CsvParser:                 sftpInputOption.CsvParser.ToCsvParserInput(ctx),
		JsonlParser:               sftpInputOption.JsonlParser.ToJsonlParserInput(ctx),
		JsonpathParser:            sftpInputOption.JsonpathParser.ToJsonpathParserInput(ctx),
		LtsvParser:                sftpInputOption.LtsvParser.ToLtsvParserInput(ctx),
		ExcelParser:               sftpInputOption.ExcelParser.ToExcelParserInput(ctx),
		XmlParser:                 sftpInputOption.XmlParser.ToXmlParserInput(ctx),
	}
}

func (sftpInputOption *SftpInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateSftpInputOptionInput {
	if sftpInputOption == nil {
		return nil
	}
	customVarSettings := common.ExtractCustomVariableSettings(ctx, sftpInputOption.CustomVariableSettings)

	return &inputOptionParameters.UpdateSftpInputOptionInput{
		SftpConnectionID:          sftpInputOption.SftpConnectionID.ValueInt64Pointer(),
		PathPrefix:                sftpInputOption.PathPrefix.ValueStringPointer(),
		PathMatchPattern:          model.NewNullableString(sftpInputOption.PathMatchPattern),
		IncrementalLoadingEnabled: sftpInputOption.IncrementalLoadingEnabled.ValueBoolPointer(),
		LastPath:                  model.NewNullableString(sftpInputOption.LastPath),
		StopWhenFileNotFound:      sftpInputOption.StopWhenFileNotFound.ValueBoolPointer(),
		DecompressionType:         sftpInputOption.DecompressionType.ValueStringPointer(),
		Decoder:                   sftpInputOption.Decoder.ToDecoderInput(),
		CustomVariableSettings:    model.ToCustomVariableSettingInputs(customVarSettings),
		CsvParser:                 sftpInputOption.CsvParser.ToCsvParserInput(ctx),
		JsonlParser:               sftpInputOption.JsonlParser.ToJsonlParserInput(ctx),
		JsonpathParser:            sftpInputOption.JsonpathParser.ToJsonpathParserInput(ctx),
		LtsvParser:                sftpInputOption.LtsvParser.ToLtsvParserInput(ctx),
		ExcelParser:               sftpInputOption.ExcelParser.ToExcelParserInput(ctx),
		XmlParser:                 sftpInputOption.XmlParser.ToXmlParserInput(ctx),
	}
}
