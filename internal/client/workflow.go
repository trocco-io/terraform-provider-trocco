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
	ResourceID int64 `json:"resource_id"`
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
	ResourceID int64 `json:"resource_id,omitempty"`
}

type WorkflowTaskDependencyInput struct {
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
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
