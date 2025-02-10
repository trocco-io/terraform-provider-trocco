package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type GoogleSpreadsheetsInputOptionInput struct {
	SpreadsheetsURL                string                                  `json:"spreadsheets_url"`
	WorksheetTitle                 string                                  `json:"worksheet_title"`
	StartRow                       int64                                   `json:"start_row"`
	StartColumn                    string                                  `json:"start_column"`
	DefaultTimeZone                string                                  `json:"default_time_zone"`
	NullString                     string                                  `json:"null_string"`
	GoogleSpreadsheetsConnectionID int64                                   `json:"google_spreadsheets_connection_id"`
	InputOptionColumns             []GoogleSpreadsheetsInputOptionColumn   `json:"input_option_columns"`
	CustomVariableSettings         *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateGoogleSpreadsheetsInputOptionInput struct {
	SpreadsheetsURL                *string                                 `json:"spreadsheets_url"`
	WorksheetTitle                 *string                                 `json:"worksheet_title"`
	StartRow                       *int64                                  `json:"start_row"`
	StartColumn                    *string                                 `json:"start_column"`
	DefaultTimeZone                *string                                 `json:"default_time_zone"`
	NullString                     *string                                 `json:"null_string"`
	GoogleSpreadsheetsConnectionID *int64                                  `json:"google_spreadsheets_connection_id,omitempty"`
	InputOptionColumns             []GoogleSpreadsheetsInputOptionColumn   `json:"input_option_columns,omitempty"`
	CustomVariableSettings         *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type GoogleSpreadsheetsInputOptionColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
