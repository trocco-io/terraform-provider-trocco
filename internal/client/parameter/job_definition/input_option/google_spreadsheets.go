package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type GoogleSpreadsheetsInputOptionInput struct {
	GoogleSpreadsheetsConnectionID int64                                   `json:"google_spreadsheets_connection_id"`
	SpreadsheetsURL                string                                  `json:"spreadsheets_url"`
	WorksheetTitle                 string                                  `json:"worksheet_title"`
	StartRow                       int64                                   `json:"start_row"`
	StartColumn                    string                                  `json:"start_column"`
	DefaultTimeZone                string                                  `json:"default_time_zone"`
	NullString                     string                                  `json:"null_string"`
	InputOptionColumns             []GoogleSpreadsheetsInputOptionColumn   `json:"input_option_columns"`
	CustomVariableSettings         *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateGoogleSpreadsheetsInputOptionInput struct {
	GoogleSpreadsheetsConnectionID *parameter.NullableInt64                `json:"google_spreadsheets_connection_id,omitempty"`
	SpreadsheetsURL                *parameter.NullableString               `json:"spreadsheets_url,omitempty"`
	WorksheetTitle                 *parameter.NullableString               `json:"worksheet_title,omitempty"`
	StartRow                       *parameter.NullableInt64                `json:"start_row,omitempty"`
	StartColumn                    *parameter.NullableString               `json:"start_column,omitempty"`
	DefaultTimeZone                *parameter.NullableString               `json:"default_time_zone,omitempty"`
	NullString                     *parameter.NullableString               `json:"null_string,omitempty"`
	InputOptionColumns             []GoogleSpreadsheetsInputOptionColumn   `json:"input_option_columns,omitempty"`
	CustomVariableSettings         *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type GoogleSpreadsheetsInputOptionColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
