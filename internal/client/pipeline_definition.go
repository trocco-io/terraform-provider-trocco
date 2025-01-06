package client

import (
	"fmt"
	"net/http"
	"net/url"

	entities "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameters"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
)

type ListPipelineDefinitionsInput struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type CreatePipelineDefinitionInput struct {
	ResourceGroupID              *p.NullableInt64     `json:"resource_group_id"`
	Name                         string               `json:"name"`
	Description                  *p.NullableString    `json:"description,omitempty"`
	MaxTaskParallelism           *p.NullableInt64     `json:"max_task_parallelism,omitempty"`
	ExecutionTimeout             *p.NullableInt64     `json:"execution_timeout,omitempty"`
	MaxRetries                   *p.NullableInt64     `json:"max_retries,omitempty"`
	MinRetryInterval             *p.NullableInt64     `json:"min_retry_interval,omitempty"`
	IsConcurrentExecutionSkipped *p.NullableBool      `json:"is_concurrent_execution_skipped,omitempty"`
	IsStoppedOnErrors            *p.NullableBool      `json:"is_stopped_on_errors,omitempty"`
	Labels                       *[]string            `json:"labels,omitempty"`
	Notifications                *[]*wp.Notification  `json:"notifications,omitempty"`
	Schedules                    *[]*wp.Schedule      `json:"schedules,omitempty"`
	Tasks                        *[]wp.Task           `json:"tasks,omitempty"`
	TaskDependencies             *[]wp.TaskDependency `json:"task_dependencies,omitempty"`
}

type UpdatePipelineDefinitionInput struct {
	ResourceGroupID              *p.NullableInt64     `json:"resource_group_id"`
	Name                         *string              `json:"name,omitempty"`
	Description                  *p.NullableString    `json:"description,omitempty"`
	MaxTaskParallelism           *p.NullableInt64     `json:"max_task_parallelism,omitempty"`
	ExecutionTimeout             *p.NullableInt64     `json:"execution_timeout,omitempty"`
	MaxRetries                   *p.NullableInt64     `json:"max_retries,omitempty"`
	MinRetryInterval             *p.NullableInt64     `json:"min_retry_interval,omitempty"`
	IsConcurrentExecutionSkipped *p.NullableBool      `json:"is_concurrent_execution_skipped,omitempty"`
	IsStoppedOnErrors            *p.NullableBool      `json:"is_stopped_on_errors,omitempty"`
	Labels                       *[]string            `json:"labels,omitempty"`
	Notifications                *[]*wp.Notification  `json:"notifications,omitempty"`
	Schedules                    *[]*wp.Schedule      `json:"schedules,omitempty"`
	Tasks                        *[]wp.Task           `json:"tasks,omitempty"`
	TaskDependencies             *[]wp.TaskDependency `json:"task_dependencies,omitempty"`
}

func (c *TroccoClient) ListPipelineDefinitions(in *ListPipelineDefinitionsInput) (*entities.PipelineDefinitionList, error) {
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

	out := &entities.PipelineDefinitionList{}
	if err := c.do(http.MethodGet, url, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) GetPipelineDefinition(id int64) (*entities.PipelineDefinition, error) {
	url := fmt.Sprintf("/api/pipeline_definitions/%d", id)

	out := &entities.PipelineDefinition{}
	if err := c.do(http.MethodGet, url, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) CreatePipelineDefinition(in *CreatePipelineDefinitionInput) (*entities.PipelineDefinition, error) {
	url := "/api/pipeline_definitions"

	out := &entities.PipelineDefinition{}
	if err := c.do(http.MethodPost, url, in, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) UpdatePipelineDefinition(id int64, in *UpdatePipelineDefinitionInput) (*entities.PipelineDefinition, error) {
	url := fmt.Sprintf("/api/pipeline_definitions/%d", id)

	out := &entities.PipelineDefinition{}
	if err := c.do(http.MethodPatch, url, in, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) DeletePipelineDefinition(id int64) error {
	url := fmt.Sprintf("/api/pipeline_definitions/%d", id)

	return c.do(http.MethodDelete, url, nil, nil)
}
