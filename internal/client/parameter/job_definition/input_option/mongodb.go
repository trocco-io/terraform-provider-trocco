package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type MongoDBInputOptionInput struct {
	Database                  string                                  `json:"database"`
	Collection                string                                  `json:"collection"`
	Query                     *parameter.NullableString               `json:"query"`
	IncrementalLoadingEnabled bool                                    `json:"incremental_loading_enabled"`
	IncrementalColumns        *parameter.NullableString               `json:"incremental_columns,omitempty"`
	LastRecord                interface{}                             `json:"last_record,omitempty"`
	MongoDBConnectionID       int64                                   `json:"mongodb_connection_id"`
	InputOptionColumns        []MongodbInputOptionColumn              `json:"input_option_columns"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateMongoDBInputOptionInput struct {
	Database                  *string                                 `json:"database,omitempty"`
	Collection                *string                                 `json:"collection,omitempty"`
	Query                     *parameter.NullableString               `json:"query,omitempty"`
	IncrementalLoadingEnabled *bool                                   `json:"incremental_loading_enabled,omitempty"`
	IncrementalColumns        *parameter.NullableString               `json:"incremental_columns,omitempty"`
	LastRecord                interface{}                             `json:"last_record,omitempty"`
	MongoDBConnectionID       *int64                                  `json:"mongodb_connection_id,omitempty"`
	InputOptionColumns        *[]MongodbInputOptionColumn             `json:"input_option_columns,omitempty"`
	CustomVariableSettings    *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type MongodbInputOptionColumn struct {
	Name     string                    `json:"name"`
	Type     string                    `json:"type"`
	Format   *parameter.NullableString `json:"format,omitempty"`
	Timezone *parameter.NullableString `json:"timezone,omitempty"`
}
