package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type GoogleSpreadsheetsInputOption struct {
	SpreadsheetsURL                string                                `json:"spreadsheets_url"`
	WorkSheetTitle                 string                                `json:"work_sheet_title"`
	StartRow                       int64                                 `json:"start_row"`
	StartColumn                    string                                `json:"start_column"`
	DefaultTimeZone                string                                `json:"default_time_zone"`
	NullString                     string                                `json:"null_string"`
	GoogleSpreadsheetsConnectionID int64                                 `json:"snowflake_connection_id"`
	InputOptionColumns             []GoogleSpreadsheetsInputOptionColumn `json:"input_option_columns"`
	CustomVariableSettings         *[]entity.CustomVariableSetting       `json:"custom_variable_settings"`
}

type GoogleSpreadsheetsInputOptionColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
