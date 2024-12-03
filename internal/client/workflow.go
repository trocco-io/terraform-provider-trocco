package client

import (
	"fmt"
	"net/http"
	"net/url"
)

type WorkflowList struct {
	Workflows  []*Workflow `json:"workflows"`
	NextCursor string      `json:"next_cursor"`
}

type Workflow struct {
	ID               int64                    `json:"id"`
	Name             *string                  `json:"name"`
	Description      *string                  `json:"description"`
	Tasks            []WorkflowTask           `json:"tasks"`
	TaskDependencies []WorkflowTaskDependency `json:"task_dependencies"`
}

type WorkflowTask struct {
	Key            string             `json:"key"`
	TaskIdentifier int64              `json:"task_identifier"`
	Type           string             `json:"type"`
	Config         WorkflowTaskConfig `json:"config"`
}

type WorkflowTaskConfig struct {
	ResourceID      *int64                             `json:"resource_id"`
	Name            *string                            `json:"name"`
	Message         *string                            `json:"message"`
	Query           *string                            `json:"query"`
	Operator        *string                            `json:"operator"`
	QueryResult     *int64                             `json:"query_result"`
	AcceptsNull     *bool                              `json:"accepts_null"`
	Warehouse       *string                            `json:"warehouse"`
	Database        *string                            `json:"database"`
	TaskID          *string                            `json:"task_id"`
	CustomVariables []WorkflowTaskCustomVariableConfig `json:"custom_variables"`

	HTTPMethod        *string                              `json:"http_method"`
	URL               *string                              `json:"url"`
	RequestBody       *string                              `json:"request_body"`
	RequestHeaders    []WorkflowTaskRequestHeaderConfig    `json:"request_headers"`
	RequestParameters []WorkflowTaskRequestParameterConfig `json:"request_parameters"`
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

type WorkflowTaskCustomVariableConfig struct {
	Name      *string `json:"name"`
	Type      *string `json:"type"`
	Value     *string `json:"value"`
	Quantity  *int64  `json:"quantity"`
	Unit      *string `json:"unit"`
	Direction *string `json:"direction"`
	Format    *string `json:"format"`
	TimeZone  *string `json:"time_zone"`
}

type WorkflowTaskDependency struct {
	Source      int64 `json:"source"`
	Destination int64 `json:"destination"`
}

type GetWorkflowsInput struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type CreateWorkflowInput struct {
	Name             string                        `json:"name"`
	Description      *string                       `json:"description,omitempty"`
	Tasks            []WorkflowTaskInput           `json:"tasks,omitempty"`
	TaskDependencies []WorkflowTaskDependencyInput `json:"task_dependencies,omitempty"`
}

type UpdateWorkflowInput struct {
	Name             *string                       `json:"name,omitempty"`
	Description      *string                       `json:"description,omitempty"`
	Tasks            []WorkflowTaskInput           `json:"tasks,omitempty"`
	TaskDependencies []WorkflowTaskDependencyInput `json:"task_dependencies,omitempty"`
}

type WorkflowTaskInput struct {
	Key            string                  `json:"key,omitempty"`
	TaskIdentifier int64                   `json:"task_identifier,omitempty"`
	Type           string                  `json:"type,omitempty"`
	Config         WorkflowTaskConfigInput `json:"config,omitempty"`
}

type WorkflowTaskConfigInput struct {
	ResourceID      *NullableInt64                          `json:"resource_id,omitempty"`
	Name            *string                                 `json:"name,omitempty"`
	Message         *string                                 `json:"message,omitempty"`
	Query           *string                                 `json:"query,omitempty"`
	Operator        *string                                 `json:"operator,omitempty"`
	QueryResult     *NullableInt64                          `json:"query_result,omitempty"`
	AcceptsNull     *bool                                   `json:"accepts_null,omitempty"`
	Warehouse       *string                                 `json:"warehouse,omitempty"`
	Database        *string                                 `json:"database,omitempty"`
	TaskID          *string                                 `json:"task_id,omitempty"`
	CustomVarialbes []WorkflowTaskCustomVariableConfigInput `json:"custom_variables,omitempty"`

	HTTPMethod        *string                                   `json:"http_method,omitempty"`
	URL               *string                                   `json:"url,omitempty"`
	RequestBody       *string                                   `json:"request_body,omitempty"`
	RequestHeaders    []WorkflowTaskRequestHeaderConfigInput    `json:"request_headers,omitempty"`
	RequestParameters []WorkflowTaskRequestParameterConfigInput `json:"request_parameters,omitempty"`
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

type WorkflowTaskDependencyInput struct {
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
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

	out := &WorkflowList{}
	if err := c.do(
		http.MethodGet,
		fmt.Sprintf("/api/workflows?%s", params.Encode()),
		nil,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) GetWorkflow(id int64) (*Workflow, error) {
	out := &Workflow{}
	if err := c.do(
		http.MethodGet,
		fmt.Sprintf("/api/workflows/%d", id),
		nil,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) CreateWorkflow(in *CreateWorkflowInput) (*Workflow, error) {
	out := &Workflow{}
	if err := c.do(
		http.MethodPost,
		"/api/workflows",
		in,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) UpdateWorkflow(id int64, in *UpdateWorkflowInput) (*Workflow, error) {
	out := &Workflow{}
	if err := c.do(
		http.MethodPatch,
		fmt.Sprintf("/api/workflows/%d", id),
		in,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) DeleteWorkflow(id int64) error {
	return c.do(
		http.MethodDelete,
		fmt.Sprintf("/api/workflows/%d", id),
		nil,
		nil,
	)
}
