package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type SalesforceInputOption struct {
	Object                          string                          `json:"object"`
	ObjectAcquisitionMethod         string                          `json:"object_acquisition_method"`
	IsConvertTypeCustomColumns      bool                            `json:"is_convert_type_custom_columns"`
	IncludeDeletedOrArchivedRecords bool                            `json:"include_deleted_or_archived_records"`
	ApiVersion                      string                          `json:"api_version"`
	Soql                            *string                         `json:"soql"`
	SalesforceConnectionID          int64                           `json:"salesforce_connection_id"`
	Columns                         []SalesforceColumn              `json:"columns"`
	CustomVariableSettings          *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type SalesforceColumn struct {
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Format *string `json:"format"`
}
