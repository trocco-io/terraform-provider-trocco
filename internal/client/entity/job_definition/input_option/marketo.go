package input_option

import (
	"terraform-provider-trocco/internal/client/entity"
)

type MarketoInputOption struct {
	MarketoConnectionID             int64                           `json:"marketo_connection_id"`
	Target                          string                          `json:"target"`
	FromDate                        *string                         `json:"from_date"`
	EndDate                         *string                         `json:"end_date"`
	UseUpdatedAt                    *bool                           `json:"use_updated_at"`
	PollingIntervalSecond           *int64                          `json:"polling_interval_second"`
	BulkJobTimeoutSecond            *int64                          `json:"bulk_job_timeout_second"`
	ActivityTypeIDs                 *[]int64                        `json:"activity_type_ids"`
	CustomObjectAPIName             *string                         `json:"custom_object_api_name"`
	CustomObjectFilterType          *string                         `json:"custom_object_filter_type"`
	CustomObjectFilterFromValue     *int64                          `json:"custom_object_filter_from_value"`
	CustomObjectFilterToValue       *int64                          `json:"custom_object_filter_to_value"`
	CustomObjectFields              *[]MarketoCustomObjectField     `json:"custom_object_fields"`
	ListIDs                         *string                         `json:"list_ids"`
	ProgramIDs                      *string                         `json:"program_ids"`
	RootID                          *int64                          `json:"root_id"`
	RootType                        *string                         `json:"root_type"`
	MaxDepth                        *int64                          `json:"max_depth"`
	Workspace                       *string                         `json:"workspace"`
	MarketoInputOptionColumns       *[]MarketoColumn                `json:"input_option_columns"`
	MarketoInputOptionFilterColumns *[]MarketoFilterColumn          `json:"input_option_filter_columns"`
	CustomVariableSettings          *[]entity.CustomVariableSetting `json:"custom_variable_settings"`
}

type MarketoColumn struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type MarketoFilterColumn struct {
	Name string `json:"name"`
}

type MarketoCustomObjectField struct {
	Name string `json:"name"`
}
