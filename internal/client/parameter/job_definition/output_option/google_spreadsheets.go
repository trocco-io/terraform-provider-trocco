package output_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type GoogleSpreadsheetsOutputOptionInput struct {
	GoogleSpreadsheetsConnectionId      int64                                                                   `json:"google_spreadsheets_connection_id"`
	SpreadsheetsID                      string                                                                  `json:"spreadsheets_id"`
	WorksheetTitle                      string                                                                  `json:"worksheet_title"`
	Timezone                            string                                                                  `json:"timezone"`
	ValueInputOption                    string                                                                  `json:"value_input_option"`
	Mode                                string                                                                  `json:"mode,omitempty"`
	GoogleSpreadsheetsOutputOptionSorts *parameter.NullableObjectList[GoogleSpreadsheetsOutputOptionSortsInput] `json:"google_spreadsheets_output_option_sorts"`
	CustomVariableSettings              *[]parameter.CustomVariableSettingInput                                 `json:"custom_variable_settings,omitempty"`
}

type UpdateGoogleSpreadsheetsOutputOptionInput struct {
	GoogleSpreadsheetsConnectionId      *int64                                                                  `json:"google_spreadsheets_connection_id"`
	SpreadsheetsID                      *string                                                                 `json:"spreadsheets_id"`
	WorksheetTitle                      *string                                                                 `json:"worksheet_title"`
	Timezone                            *string                                                                 `json:"timezone"`
	ValueInputOption                    *string                                                                 `json:"value_input_option"`
	Mode                                *string                                                                 `json:"mode,omitempty"`
	GoogleSpreadsheetsOutputOptionSorts *parameter.NullableObjectList[GoogleSpreadsheetsOutputOptionSortsInput] `json:"google_spreadsheets_output_option_sorts"`
	CustomVariableSettings              *[]parameter.CustomVariableSettingInput                                 `json:"custom_variable_settings,omitempty"`
}

type GoogleSpreadsheetsOutputOptionSortsInput struct {
	Column string `json:"column"`
	Order  string `json:"order"`
}
