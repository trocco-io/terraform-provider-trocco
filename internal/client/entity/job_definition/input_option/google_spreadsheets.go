package input_option

import (
	"strings"
	"terraform-provider-trocco/internal/client/entity"
)

type GoogleSpreadsheetsInputOption struct {
	SpreadsheetsURL                string                                `json:"spreadsheets_url"`
	WorksheetTitle                 string                                `json:"worksheet_title"`
	StartRow                       int64                                 `json:"start_row"`
	StartColumn                    string                                `json:"start_column"`
	DefaultTimeZone                string                                `json:"default_time_zone"`
	NullString                     string                                `json:"null_string"`
	GoogleSpreadsheetsConnectionID int64                                 `json:"google_spreadsheets_connection_id"`
	InputOptionColumns             []GoogleSpreadsheetsInputOptionColumn `json:"input_option_columns"`
	CustomVariableSettings         *[]entity.CustomVariableSetting       `json:"custom_variable_settings"`
}

func (inputOption *GoogleSpreadsheetsInputOption) SpreadsheetsID() string {
	if inputOption.SpreadsheetsURL == "" {
		return ""
	}
	// ex) "https://docs.google.com/spreadsheets/d/MY_SHEETS_ID/edit#gid=0" --> "MY_SHEETS_ID"
	return strings.Split(strings.Split(inputOption.SpreadsheetsURL, "https://docs.google.com/spreadsheets/d/")[1], "/")[0]
}

type GoogleSpreadsheetsInputOptionColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
