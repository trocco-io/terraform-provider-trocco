package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type KintoneInputOption struct {
	AppID                  string                          `json:"app_id"`
	GuestSpaceID           *string                         `json:"guest_space_id"`
	ExpandSubtable         bool                            `json:"expand_subtable"`
	Query                  *string                         `json:"query"`
	KintoneConnectionID    int64                           `json:"kintone_connection_id"`
	InputOptionColumns     []KintoneInputOptionColumn      `json:"input_option_columns"`
	CustomVariableSettings *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type KintoneInputOptionColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
