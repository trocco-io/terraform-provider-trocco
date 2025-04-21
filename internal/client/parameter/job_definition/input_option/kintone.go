package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type KintoneInputOptionInput struct {
	AppID                  string                                  `json:"app_id"`
	GuestSpaceID           *parameter.NullableString               `json:"guest_space_id,omitempty"`
	ExpandSubtable         *parameter.NullableBool                 `json:"expand_subtable"`
	Query                  *parameter.NullableString               `json:"query,omitempty"`
	KintoneConnectionID    int64                                   `json:"kintone_connection_id"`
	InputOptionColumns     []KintoneInputOptionColumn              `json:"input_option_columns"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateKintoneInputOptionInput struct {
	AppID                  *parameter.NullableString               `json:"app_id,omitempty"`
	GuestSpaceID           *parameter.NullableString               `json:"guest_space_id,omitempty"`
	ExpandSubtable         *parameter.NullableBool                 `json:"expand_subtable,omitempty"`
	Query                  *parameter.NullableString               `json:"query,omitempty"`
	KintoneConnectionID    *parameter.NullableInt64                `json:"kintone_connection_id,omitempty"`
	InputOptionColumns     []KintoneInputOptionColumn              `json:"input_option_columns,omitempty"`
	CustomVariableSettings *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type KintoneInputOptionColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
