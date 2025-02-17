package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type SalesforceInputOptionInput struct {
	Object                          string                                  `json:"object"`
	ObjectAcquisitionMethod         *parameter.NullableString               `json:"object_acquisition_method,omitempty"`
	IsConvertTypeCustomColumns      *parameter.NullableBool                 `json:"is_convert_type_custom_columns,omitempty"`
	IncludeDeletedOrArchivedRecords *parameter.NullableBool                 `json:"include_deleted_or_archived_records,omitempty"`
	ApiVersion                      *parameter.NullableString               `json:"api_version,omitempty"`
	Soql                            *parameter.NullableString               `json:"soql,omitempty,omitempty"`
	SalesforceConnectionID          int64                                   `json:"salesforce_connection_id"`
	Columns                         []SalesforceColumn                      `json:"columns"`
	CustomVariableSettings          *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateSalesforceInputOptionInput struct {
	Object                          *string                                 `json:"object,omitempty"`
	ObjectAcquisitionMethod         *parameter.NullableString               `json:"object_acquisition_method,omitempty"`
	IsConvertTypeCustomColumns      *parameter.NullableBool                 `json:"is_convert_type_custom_columns,omitempty"`
	IncludeDeletedOrArchivedRecords *parameter.NullableBool                 `json:"include_deleted_or_archived_records,omitempty"`
	ApiVersion                      *parameter.NullableString               `json:"api_version,omitempty"`
	Soql                            *parameter.NullableString               `json:"soql,omitempty,omitempty"`
	SalesforceConnectionID          *int64                                  `json:"salesforce_connection_id,omitempty"`
	Columns                         []SalesforceColumn                      `json:"columns,omitempty"`
	CustomVariableSettings          *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type SalesforceColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
