package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type GoogleDriveOutputOptionInput struct {
	GoogleDriveConnectionID int64                                   `json:"google_drive_connection_id"`
	MainFolderID            string                                  `json:"main_folder_id"`
	ChildFolderName         *parameter.NullableString               `json:"child_folder_name,omitempty"`
	FileName                string                                  `json:"file_name"`
	FormatterType           string                                  `json:"formatter_type"`
	CsvFormatter            *CsvFormatterInput                      `json:"csv_formatter,omitempty"`
	CustomVariableSettings  *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateGoogleDriveOutputOptionInput struct {
	GoogleDriveConnectionID *int64                                  `json:"google_drive_connection_id,omitempty"`
	MainFolderID            *string                                 `json:"main_folder_id,omitempty"`
	ChildFolderName         *parameter.NullableString               `json:"child_folder_name,omitempty"`
	FileName                *string                                 `json:"file_name,omitempty"`
	FormatterType           *string                                 `json:"formatter_type,omitempty"`
	CsvFormatter            *CsvFormatterInput                      `json:"csv_formatter,omitempty"`
	CustomVariableSettings  *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}
