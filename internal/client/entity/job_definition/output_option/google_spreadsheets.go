package output_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type GoogleSpreadsheetsOutputOption struct {
	GoogleSpreadsheetsConnectionID      int64                                  `json:"google_spreadsheets_connection_id"`
	SpreadsheetsID                      string                                 `json:"spreadsheets_id"`
	WorksheetTitle                      string                                 `json:"worksheet_title"`
	Timezone                            string                                 `json:"timezone"`
	ValueInputOption                    string                                 `json:"value_input_option"`
	Mode                                string                                 `json:"mode"`
	GoogleSpreadsheetsOutputOptionSorts *[]GoogleSpreadsheetsOutputOptionSorts `json:"google_spreadsheets_output_option_sorts"`
	CustomVariableSettings              *[]entity.CustomVariableSetting        `json:"custom_variable_settings"`
}

type GoogleSpreadsheetsOutputOptionSorts struct {
	Column string `json:"column"`
	Order  string `json:"order"`
}
