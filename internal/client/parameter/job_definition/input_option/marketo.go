package input_options

import (
	"terraform-provider-trocco/internal/client/parameter"
)

type MarketoInputOptionInput struct {
	MarketoConnectionID             int64                                   `json:"marketo_connection_id"`
	Target                          string                                  `json:"target"`
	FromDate                        *string                                 `json:"from_date,omitempty"`
	EndDate                         *string                                 `json:"end_date,omitempty"`
	UseUpdatedAt                    *parameter.NullableBool                 `json:"use_updated_at,omitempty"`
	PollingIntervalSecond           *parameter.NullableInt64                `json:"polling_interval_second,omitempty"`
	BulkJobTimeoutSecond            *parameter.NullableInt64                `json:"bulk_job_timeout_second,omitempty"`
	ActivityTypeIDs                 *[]int64                                `json:"activity_type_ids,omitempty"`
	CustomObjectAPIName             *string                                 `json:"custom_object_api_name,omitempty"`
	CustomObjectFilterType          *string                                 `json:"custom_object_filter_type,omitempty"`
	CustomObjectFilterFromValue     *int64                                  `json:"custom_object_filter_from_value,omitempty"`
	CustomObjectFilterToValue       *int64                                  `json:"custom_object_filter_to_value,omitempty"`
	CustomObjectFields              *[]MarketoCustomObjectField             `json:"custom_object_fields,omitempty"`
	ListIDs                         *string                                 `json:"list_ids,omitempty"`
	ProgramIDs                      *string                                 `json:"program_ids,omitempty"`
	RootID                          *int64                                  `json:"root_id,omitempty"`
	RootType                        *string                                 `json:"root_type,omitempty"`
	MaxDepth                        *parameter.NullableInt64                `json:"max_depth,omitempty"`
	Workspace                       *string                                 `json:"workspace,omitempty"`
	MarketoInputOptionColumns       *[]MarketoColumn                        `json:"input_option_columns,omitempty"`
	MarketoInputOptionFilterColumns *[]MarketoFilterColumn                  `json:"marketo_input_option_filter_columns,omitempty"`
	CustomVariableSettings          *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
}

type UpdateMarketoInputOptionInput struct {
	MarketoConnectionID             *int64                                  `json:"marketo_connection_id,omitempty"`
	Target                          *string                                 `json:"target,omitempty"`
	FromDate                        *string                                 `json:"from_date,omitempty"`
	EndDate                         *string                                 `json:"end_date,omitempty"`
	UseUpdatedAt                    *parameter.NullableBool                 `json:"use_updated_at,omitempty"`
	PollingIntervalSecond           *parameter.NullableInt64                `json:"polling_interval_second,omitempty"`
	BulkJobTimeoutSecond            *parameter.NullableInt64                `json:"bulk_job_timeout_second,omitempty"`
	ActivityTypeIDs                 *[]int64                                `json:"activity_type_ids,omitempty"`
	CustomObjectAPIName             *string                                 `json:"custom_object_api_name,omitempty"`
	CustomObjectFilterType          *string                                 `json:"custom_object_filter_type,omitempty"`
	CustomObjectFilterFromValue     *int64                                  `json:"custom_object_filter_from_value,omitempty"`
	CustomObjectFilterToValue       *int64                                  `json:"custom_object_filter_to_value,omitempty"`
	CustomObjectFields              *[]MarketoCustomObjectField             `json:"custom_object_fields,omitempty"`
	ListIDs                         *string                                 `json:"list_ids,omitempty"`
	ProgramIDs                      *string                                 `json:"program_ids,omitempty"`
	RootID                          *int64                                  `json:"root_id,omitempty"`
	RootType                        *string                                 `json:"root_type,omitempty"`
	MaxDepth                        *parameter.NullableInt64                `json:"max_depth,omitempty"`
	Workspace                       *string                                 `json:"workspace,omitempty"`
	MarketoInputOptionColumns       *[]MarketoColumn                        `json:"input_option_columns,omitempty"`
	MarketoInputOptionFilterColumns *[]MarketoFilterColumn                  `json:"marketo_input_option_filter_columns,omitempty"`
	CustomVariableSettings          *[]parameter.CustomVariableSettingInput `json:"custom_variable_settings,omitempty"`
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
