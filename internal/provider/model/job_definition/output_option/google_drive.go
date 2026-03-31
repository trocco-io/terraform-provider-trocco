package output_options

import (
	"context"
	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	outputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GoogleDriveOutputOption struct {
	GoogleDriveConnectionID types.Int64   `tfsdk:"google_drive_connection_id"`
	MainFolderID            types.String  `tfsdk:"main_folder_id"`
	ChildFolderName         types.String  `tfsdk:"child_folder_name"`
	FileName                types.String  `tfsdk:"file_name"`
	FormatterType           types.String  `tfsdk:"formatter_type"`
	CsvFormatter            *csvFormatter `tfsdk:"csv_formatter"`
	CustomVariableSettings  types.List    `tfsdk:"custom_variable_settings"`
}

func NewGoogleDriveOutputOption(ctx context.Context, googleDriveOutputOption *output_option.GoogleDriveOutputOption) *GoogleDriveOutputOption {
	if googleDriveOutputOption == nil {
		return nil
	}

	result := &GoogleDriveOutputOption{
		GoogleDriveConnectionID: types.Int64Value(googleDriveOutputOption.GoogleDriveConnectionID),
		MainFolderID:            types.StringValue(googleDriveOutputOption.MainFolderID),
		ChildFolderName:         types.StringPointerValue(googleDriveOutputOption.ChildFolderName),
		FileName:                types.StringValue(googleDriveOutputOption.FileName),
	}

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, googleDriveOutputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	if googleDriveOutputOption.Formatter != nil {
		result.FormatterType = types.StringValue(googleDriveOutputOption.Formatter.Type)
		result.CsvFormatter = newCsvFormatter(ctx, googleDriveOutputOption.Formatter.CsvFormatter)
	}

	return result
}

func (o *GoogleDriveOutputOption) ToInput(ctx context.Context) *outputOptionParameters.GoogleDriveOutputOptionInput {
	if o == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, o.CustomVariableSettings)

	return &outputOptionParameters.GoogleDriveOutputOptionInput{
		GoogleDriveConnectionID: o.GoogleDriveConnectionID.ValueInt64(),
		MainFolderID:            o.MainFolderID.ValueString(),
		ChildFolderName:         model.NewNullableString(o.ChildFolderName),
		FileName:                o.FileName.ValueString(),
		FormatterType:           o.FormatterType.ValueString(),
		CsvFormatter:            o.CsvFormatter.toCsvFormatterInput(ctx),
		CustomVariableSettings:  model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (o *GoogleDriveOutputOption) ToUpdateInput(ctx context.Context) *outputOptionParameters.UpdateGoogleDriveOutputOptionInput {
	if o == nil {
		return nil
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, o.CustomVariableSettings)

	return &outputOptionParameters.UpdateGoogleDriveOutputOptionInput{
		GoogleDriveConnectionID: o.GoogleDriveConnectionID.ValueInt64Pointer(),
		MainFolderID:            o.MainFolderID.ValueStringPointer(),
		ChildFolderName:         model.NewNullableString(o.ChildFolderName),
		FileName:                o.FileName.ValueStringPointer(),
		FormatterType:           o.FormatterType.ValueStringPointer(),
		CsvFormatter:            o.CsvFormatter.toCsvFormatterInput(ctx),
		CustomVariableSettings:  model.ToCustomVariableSettingInputs(customVarSettings),
	}
}
