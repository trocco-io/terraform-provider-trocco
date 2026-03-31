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

type GoogleDriveInputOption struct {
	GoogleDriveConnectionID types.Int64            `tfsdk:"google_drive_connection_id"`
	FolderID                types.String           `tfsdk:"folder_id"`
	FileMatchPattern        types.String           `tfsdk:"file_match_pattern"`
	IsSkipHeaderLine        types.Bool             `tfsdk:"is_skip_header_line"`
	StopWhenFileNotFound    types.Bool             `tfsdk:"stop_when_file_not_found"`
	DecompressionType       types.String           `tfsdk:"decompression_type"`
	CsvParser               *parser.CsvParser      `tfsdk:"csv_parser"`
	JsonlParser             *parser.JsonlParser    `tfsdk:"jsonl_parser"`
	JsonpathParser          *parser.JsonpathParser `tfsdk:"jsonpath_parser"`
	LtsvParser              *parser.LtsvParser     `tfsdk:"ltsv_parser"`
	ExcelParser             *parser.ExcelParser    `tfsdk:"excel_parser"`
	XmlParser               *parser.XmlParser      `tfsdk:"xml_parser"`
	CustomVariableSettings  types.List             `tfsdk:"custom_variable_settings"`
	Decoder                 *Decoder               `tfsdk:"decoder"`
}

func NewGoogleDriveInputOption(ctx context.Context, googleDriveInputOption *inputOptionEntities.GoogleDriveInputOption) *GoogleDriveInputOption {
	if googleDriveInputOption == nil {
		return nil
	}

	result := &GoogleDriveInputOption{
		GoogleDriveConnectionID: types.Int64Value(googleDriveInputOption.GoogleDriveConnectionID),
		FolderID:                types.StringValue(googleDriveInputOption.FolderID),
		FileMatchPattern:        types.StringValue(googleDriveInputOption.FileMatchPattern),
		IsSkipHeaderLine:        types.BoolValue(googleDriveInputOption.IsSkipHeaderLine),
		StopWhenFileNotFound:    types.BoolValue(googleDriveInputOption.StopWhenFileNotFound),
		DecompressionType:       types.StringPointerValue(googleDriveInputOption.DecompressionType),
		CsvParser:               parser.NewCsvParser(ctx, googleDriveInputOption.CsvParser),
		JsonlParser:             parser.NewJsonlParser(ctx, googleDriveInputOption.JsonlParser),
		JsonpathParser:          parser.NewJsonPathParser(ctx, googleDriveInputOption.JsonpathParser),
		LtsvParser:              parser.NewLtsvParser(ctx, googleDriveInputOption.LtsvParser),
		ExcelParser:             parser.NewExcelParser(ctx, googleDriveInputOption.ExcelParser),
		XmlParser:               parser.NewXmlParser(ctx, googleDriveInputOption.XmlParser),
		Decoder:                 NewDecoder(googleDriveInputOption.Decoder),
	}

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, googleDriveInputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func (googleDriveInputOption *GoogleDriveInputOption) ToInput(ctx context.Context) *inputOptionParameters.GoogleDriveInputOptionInput {
	if googleDriveInputOption == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, googleDriveInputOption.CustomVariableSettings)

	return &inputOptionParameters.GoogleDriveInputOptionInput{
		GoogleDriveConnectionID: googleDriveInputOption.GoogleDriveConnectionID.ValueInt64(),
		FolderID:                googleDriveInputOption.FolderID.ValueString(),
		FileMatchPattern:        googleDriveInputOption.FileMatchPattern.ValueString(),
		IsSkipHeaderLine:        googleDriveInputOption.IsSkipHeaderLine.ValueBool(),
		StopWhenFileNotFound:    googleDriveInputOption.StopWhenFileNotFound.ValueBool(),
		DecompressionType:       model.NewNullableString(googleDriveInputOption.DecompressionType),
		CsvParser:               googleDriveInputOption.CsvParser.ToCsvParserInput(ctx),
		JsonlParser:             googleDriveInputOption.JsonlParser.ToJsonlParserInput(ctx),
		JsonpathParser:          googleDriveInputOption.JsonpathParser.ToJsonpathParserInput(ctx),
		LtsvParser:              googleDriveInputOption.LtsvParser.ToLtsvParserInput(ctx),
		ExcelParser:             googleDriveInputOption.ExcelParser.ToExcelParserInput(ctx),
		XmlParser:               googleDriveInputOption.XmlParser.ToXmlParserInput(ctx),
		CustomVariableSettings:  model.ToCustomVariableSettingInputs(customVarSettings),
		Decoder:                 googleDriveInputOption.Decoder.ToDecoderInput(),
	}
}

func (googleDriveInputOption *GoogleDriveInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateGoogleDriveInputOptionInput {
	if googleDriveInputOption == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, googleDriveInputOption.CustomVariableSettings)

	return &inputOptionParameters.UpdateGoogleDriveInputOptionInput{
		GoogleDriveConnectionID: googleDriveInputOption.GoogleDriveConnectionID.ValueInt64Pointer(),
		FolderID:                googleDriveInputOption.FolderID.ValueStringPointer(),
		FileMatchPattern:        googleDriveInputOption.FileMatchPattern.ValueStringPointer(),
		IsSkipHeaderLine:        googleDriveInputOption.IsSkipHeaderLine.ValueBoolPointer(),
		StopWhenFileNotFound:    googleDriveInputOption.StopWhenFileNotFound.ValueBoolPointer(),
		DecompressionType:       model.NewNullableString(googleDriveInputOption.DecompressionType),
		CsvParser:               googleDriveInputOption.CsvParser.ToCsvParserInput(ctx),
		JsonlParser:             googleDriveInputOption.JsonlParser.ToJsonlParserInput(ctx),
		JsonpathParser:          googleDriveInputOption.JsonpathParser.ToJsonpathParserInput(ctx),
		LtsvParser:              googleDriveInputOption.LtsvParser.ToLtsvParserInput(ctx),
		ExcelParser:             googleDriveInputOption.ExcelParser.ToExcelParserInput(ctx),
		XmlParser:               googleDriveInputOption.XmlParser.ToXmlParserInput(ctx),
		CustomVariableSettings:  model.ToCustomVariableSettingInputs(customVarSettings),
		Decoder:                 googleDriveInputOption.Decoder.ToDecoderInput(),
	}
}
