package client

import (
	"fmt"
	"net/http"
	"net/url"

	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
)

// -----------------------------------------------------------------------------
// Client-side data types
// -----------------------------------------------------------------------------

type WorkflowList struct {
	Workflows  []*Workflow `json:"workflows"`
	NextCursor string      `json:"next_cursor"`
}

type Workflow struct {
	ID               int64                    `json:"id"`
	Name             *string                  `json:"name"`
	Description      *string                  `json:"description"`
	Labels           []string                 `json:"labels"`
	Notifications    []we.Notification        `json:"notifications"`
	Schedules        []we.Schedule            `json:"schedules"`
	Tasks            []WorkflowTask           `json:"tasks"`
	TaskDependencies []WorkflowTaskDependency `json:"task_dependencies"`
}

type WorkflowTask struct {
	Key            string `json:"key"`
	TaskIdentifier int64  `json:"task_identifier"`
	Type           string `json:"type"`

	TroccoTransferConfig          *we.TroccoTransferTaskConfig          `json:"trocco_transfer_config"`
	TroccoTransferBulkConfig      *we.TroccoTransferBulkTaskConfig      `json:"trocco_transfer_bulk_config"`
	DBTConfig                     *we.DBTTaskConfig                     `json:"dbt_config"`
	TroccoAgentConfig             *we.TroccoAgentTaskConfig             `json:"trocco_agent_config"`
	WorkflowConfig                *we.WorkflowTaskConfig                `json:"workflow_config"`
	TroccoBigQueryDatamartConfig  *we.TroccoBigQueryDatamartTaskConfig  `json:"trocco_bigquery_datamart_config"`
	TroccoRedshiftDatamartConfig  *we.TroccoRedshiftDatamartTaskConfig  `json:"trocco_redshift_datamart_config"`
	TroccoSnowflakeDatamartConfig *we.TroccoSnowflakeDatamartTaskConfig `json:"trocco_snowflake_datamart_config"`
	SlackNotificationConfig       *we.SlackNotificationTaskConfig       `json:"slack_notification_config"`
	TableauDataExtractionConfig   *we.TableauDataExtractionTaskConfig   `json:"tableau_data_extraction_config"`
	BigqueryDataCheckConfig       *WorkflowBigqueryDataCheckTaskConfig  `json:"bigquery_data_check_config"`
	SnowflakeDataCheckConfig      *WorkflowSnowflakeDataCheckTaskConfig `json:"snowflake_data_check_config"`
	RedshiftDataCheckConfig       *WorkflowRedshiftDataCheckTaskConfig  `json:"redshift_data_check_config"`
	HTTPRequestConfig             *WorkflowHTTPRequestTaskConfig        `json:"http_request_config"`
}

type WorkflowHTTPRequestTaskConfig struct {
	Name              string                               `json:"name"`
	ConnectionID      *int64                               `json:"connection_id"`
	HTTPMethod        string                               `json:"http_method"`
	URL               string                               `json:"url"`
	RequestBody       *string                              `json:"request_body"`
	RequestHeaders    []WorkflowTaskRequestHeaderConfig    `json:"request_headers"`
	RequestParameters []WorkflowTaskRequestParameterConfig `json:"request_parameters"`
	CustomVariables   []we.CustomVariable                  `json:"custom_variables"`
}

type WorkflowTaskRequestHeaderConfig struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Masking bool   `json:"masking"`
}

type WorkflowTaskRequestParameterConfig struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Masking bool   `json:"masking"`
}

type WorkflowBigqueryDataCheckTaskConfig struct {
	Name            string              `json:"name"`
	ConnectionID    int64               `json:"connection_id"`
	Query           string              `json:"query"`
	Operator        string              `json:"operator"`
	QueryResult     int64               `json:"query_result"`
	AcceptsNull     bool                `json:"accepts_null"`
	CustomVariables []we.CustomVariable `json:"custom_variables"`
}

type WorkflowSnowflakeDataCheckTaskConfig struct {
	Name            string              `json:"name"`
	ConnectionID    int64               `json:"connection_id"`
	Query           string              `json:"query"`
	Operator        string              `json:"operator"`
	QueryResult     int64               `json:"query_result"`
	AcceptsNull     bool                `json:"accepts_null"`
	Warehouse       string              `json:"warehouse"`
	CustomVariables []we.CustomVariable `json:"custom_variables"`
}

type WorkflowRedshiftDataCheckTaskConfig struct {
	Name            string              `json:"name"`
	ConnectionID    int64               `json:"connection_id"`
	Query           string              `json:"query"`
	Operator        string              `json:"operator"`
	QueryResult     int64               `json:"query_result"`
	AcceptsNull     bool                `json:"accepts_null"`
	Database        string              `json:"database"`
	CustomVariables []we.CustomVariable `json:"custom_variables"`
}

type WorkflowTaskDependency struct {
	Source      int64 `json:"source"`
	Destination int64 `json:"destination"`
}

// -----------------------------------------------------------------------------
// Parameters
// -----------------------------------------------------------------------------

type GetWorkflowsInput struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type CreateWorkflowInput struct {
	Name             string                        `json:"name"`
	Description      *string                       `json:"description,omitempty"`
	Labels           *[]string                     `json:"labels,omitempty"`
	Notifications    *[]wp.Notification            `json:"notifications,omitempty"`
	Schedules        *[]wp.Schedule                `json:"schedules,omitempty"`
	Tasks            []WorkflowTaskInput           `json:"tasks,omitempty"`
	TaskDependencies []WorkflowTaskDependencyInput `json:"task_dependencies,omitempty"`
}

type UpdateWorkflowInput struct {
	Name             *string                       `json:"name,omitempty"`
	Description      *string                       `json:"description,omitempty"`
	Labels           *[]string                     `json:"labels,omitempty"`
	Notifications    *[]wp.Notification            `json:"notifications,omitempty"`
	Schedules        *[]wp.Schedule                `json:"schedules,omitempty"`
	Tasks            []WorkflowTaskInput           `json:"tasks,omitempty"`
	TaskDependencies []WorkflowTaskDependencyInput `json:"task_dependencies,omitempty"`
}

type WorkflowTaskInput struct {
	Key            string `json:"key,omitempty"`
	TaskIdentifier int64  `json:"task_identifier,omitempty"`
	Type           string `json:"type,omitempty"`

	TroccoTransferConfig          *wp.TroccoTransferTaskConfig                  `json:"trocco_transfer_config,omitempty"`
	TroccoTransferBulkConfig      *wp.TroccoTransferBulkTaskConfig              `json:"trocco_transfer_bulk_config,omitempty"`
	DBTConfig                     *wp.DBTTaskConfig                             `json:"dbt_config,omitempty"`
	TroccoAgentConfig             *wp.TroccoAgentTaskConfig                     `json:"trocco_agent_config,omitempty"`
	TroccoBigQueryDatamartConfig  *wp.TroccoBigQueryDatamartTaskConfig          `json:"trocco_bigquery_datamart_config,omitempty"`
	TroccoRedshiftDatamartConfig  *wp.TroccoRedshiftDatamartTaskConfig          `json:"trocco_redshift_datamart_config,omitempty"`
	TroccoSnowflakeDatamartConfig *wp.TroccoSnowflakeDatamartTaskConfig         `json:"trocco_snowflake_datamart_config,omitempty"`
	WorkflowConfig                *wp.WorkflowTaskConfig                        `json:"workflow_config,omitempty"`
	SlackNotificationConfig       *wp.SlackNotificationTaskConfig               `json:"slack_notification_config,omitempty"`
	TableauDataExtractionConfig   *WorkflowTableauDataExtractionTaskConfigInput `json:"tableau_data_extraction_config,omitempty"`
	BigqueryDataCheckConfig       *WorkflowBigqueryDataCheckTaskConfigInput     `json:"bigquery_data_check_config,omitempty"`
	SnowflakeDataCheckConfig      *WorkflowSnowflakeDataCheckTaskConfigInput    `json:"snowflake_data_check_config,omitempty"`
	RedshiftDataCheckConfig       *WorkflowRedshiftDataCheckTaskConfigInput     `json:"redshift_data_check_config,omitempty"`
	HTTPRequestConfig             *WorkflowHTTPRequestTaskConfigInput           `json:"http_request_config,omitempty"`
}

type WorkflowTableauDataExtractionTaskConfigInput struct {
	Name         string `json:"name,omitempty"`
	ConnectionID int64  `json:"connection_id,omitempty"`
	TaskID       string `json:"task_id,omitempty"`
}

type WorkflowBigqueryDataCheckTaskConfigInput struct {
	Name            string                                  `json:"name,omitempty"`
	ConnectionID    int64                                   `json:"connection_id,omitempty"`
	Query           string                                  `json:"query,omitempty"`
	Operator        string                                  `json:"operator,omitempty"`
	QueryResult     *NullableInt64                          `json:"query_result,omitempty"`
	AcceptsNull     *NullableBool                           `json:"accepts_null,omitempty"`
	CustomVariables []WorkflowTaskCustomVariableConfigInput `json:"custom_variables,omitempty"`
}

type WorkflowSnowflakeDataCheckTaskConfigInput struct {
	Name            string                                  `json:"name,omitempty"`
	ConnectionID    int64                                   `json:"connection_id,omitempty"`
	Query           string                                  `json:"query,omitempty"`
	Operator        string                                  `json:"operator,omitempty"`
	QueryResult     *NullableInt64                          `json:"query_result,omitempty"`
	AcceptsNull     *NullableBool                           `json:"accepts_null,omitempty"`
	Warehouse       string                                  `json:"warehouse,omitempty"`
	CustomVariables []WorkflowTaskCustomVariableConfigInput `json:"custom_variables,omitempty"`
}

type WorkflowRedshiftDataCheckTaskConfigInput struct {
	Name            string                                  `json:"name,omitempty"`
	ConnectionID    int64                                   `json:"connection_id,omitempty"`
	Query           string                                  `json:"query,omitempty"`
	Operator        string                                  `json:"operator,omitempty"`
	QueryResult     *NullableInt64                          `json:"query_result,omitempty"`
	AcceptsNull     *NullableBool                           `json:"accepts_null,omitempty"`
	Database        string                                  `json:"database,omitempty"`
	CustomVariables []WorkflowTaskCustomVariableConfigInput `json:"custom_variables,omitempty"`
}

type WorkflowHTTPRequestTaskConfigInput struct {
	Name              string                                    `json:"name,omitempty"`
	ConnectionID      *NullableInt64                            `json:"connection_id,omitempty"`
	HTTPMethod        string                                    `json:"http_method,omitempty"`
	URL               string                                    `json:"url,omitempty"`
	RequestBody       *string                                   `json:"request_body,omitempty"`
	RequestHeaders    []WorkflowTaskRequestHeaderConfigInput    `json:"request_headers,omitempty"`
	RequestParameters []WorkflowTaskRequestParameterConfigInput `json:"request_parameters,omitempty"`
	CustomVariables   []WorkflowTaskCustomVariableConfigInput   `json:"custom_variables,omitempty"`
}

type WorkflowTaskRequestHeaderConfigInput struct {
	Key     string        `json:"key,omitempty"`
	Value   string        `json:"value,omitempty"`
	Masking *NullableBool `json:"masking,omitempty"`
}

type WorkflowTaskRequestParameterConfigInput struct {
	Key     string        `json:"key,omitempty"`
	Value   string        `json:"value,omitempty"`
	Masking *NullableBool `json:"masking,omitempty"`
}

type WorkflowTaskCustomVariableConfigInput struct {
	Name      *string        `json:"name,omitempty"`
	Type      *string        `json:"type,omitempty"`
	Value     *string        `json:"value,omitempty"`
	Quantity  *NullableInt64 `json:"quantity,omitempty"`
	Unit      *string        `json:"unit,omitempty"`
	Direction *string        `json:"direction,omitempty"`
	Format    *string        `json:"format,omitempty"`
	TimeZone  *string        `json:"time_zone,omitempty"`
}

type WorkflowTaskDependencyInput struct {
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
}

// -----------------------------------------------------------------------------
// Operations
// -----------------------------------------------------------------------------

func (c *TroccoClient) GetWorkflows(in *GetWorkflowsInput) (*WorkflowList, error) {
	params := url.Values{}
	if in != nil {
		if in.Limit != 0 {
			params.Add("limit", fmt.Sprintf("%d", in.Limit))
		}

		if in.Cursor != "" {
			params.Add("cursor", in.Cursor)
		}
	}

	url := fmt.Sprintf("/api/pipeline_definitions?%s", params.Encode())

	out := &WorkflowList{}
	if err := c.do(http.MethodGet, url, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) GetWorkflow(id int64) (*Workflow, error) {
	url := fmt.Sprintf("/api/pipeline_definitions/%d", id)

	out := &Workflow{}
	if err := c.do(http.MethodGet, url, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) CreateWorkflow(in *CreateWorkflowInput) (*Workflow, error) {
	url := "/api/pipeline_definitions"

	out := &Workflow{}
	if err := c.do(http.MethodPost, url, in, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) UpdateWorkflow(id int64, in *UpdateWorkflowInput) (*Workflow, error) {
	url := fmt.Sprintf("/api/pipeline_definitions/%d", id)

	out := &Workflow{}
	if err := c.do(http.MethodPatch, url, in, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) DeleteWorkflow(id int64) error {
	url := fmt.Sprintf("/api/pipeline_definitions/%d", id)

	return c.do(http.MethodDelete, url, nil, nil)
}
