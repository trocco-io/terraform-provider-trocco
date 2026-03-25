package output_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type GoogleDriveOutputOption struct {
	GoogleDriveConnectionID int64                           `json:"google_drive_connection_id"`
	MainFolderID            string                          `json:"main_folder_id"`
	ChildFolderName         *string                         `json:"child_folder_name"`
	FileName                string                          `json:"file_name"`
	Formatter               *GoogleDriveFormatter           `json:"formatter"`
	CustomVariableSettings  *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type GoogleDriveFormatter struct {
	Type         string        `json:"type"`
	CsvFormatter *CsvFormatter `json:"csv_formatter"`
}
