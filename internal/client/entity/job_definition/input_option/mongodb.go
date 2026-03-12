package input_option

import "terraform-provider-trocco/internal/client/entity"

type MongoDBInputOption struct {
	Database                  string                          `json:"database"`
	Collection                string                          `json:"collection"`
	Query                     *string                         `json:"query"`
	IncrementalLoadingEnabled bool                            `json:"incremental_loading_enabled"`
	IncrementalColumns        *string                         `json:"incremental_columns"`
	LastRecord                interface{}                     `json:"last_record"`
	MongoDBConnectionID       int64                           `json:"mongodb_connection_id"`
	InputOptionColumns        []MongodbInputOptionColumn      `json:"input_option_columns"`
	CustomVariableSettings    *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type MongodbInputOptionColumn struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Format   *string `json:"format,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}
