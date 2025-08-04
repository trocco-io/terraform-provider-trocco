package client

import (
	"fmt"
	"net/http"
	"net/url"

	entity "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	parameter "terraform-provider-trocco/internal/client/parameter"
	pipelineDefinitionParameters "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
)

type ListPipelineDefinitionsInput struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type CreatePipelineDefinitionInput struct {
	ResourceGroupID              *parameter.NullableInt64                       `json:"resource_group_id"`
	Name                         string                                         `json:"name"`
	Description                  *parameter.NullableString                      `json:"description,omitempty"`
	MaxTaskParallelism           *parameter.NullableInt64                       `json:"max_task_parallelism,omitempty"`
	ExecutionTimeout             *parameter.NullableInt64                       `json:"execution_timeout,omitempty"`
	MaxRetries                   *parameter.NullableInt64                       `json:"max_retries,omitempty"`
	MinRetryInterval             *parameter.NullableInt64                       `json:"min_retry_interval,omitempty"`
	IsConcurrentExecutionSkipped *parameter.NullableBool                        `json:"is_concurrent_execution_skipped,omitempty"`
	IsStoppedOnErrors            *parameter.NullableBool                        `json:"is_stopped_on_errors,omitempty"`
	Labels                       *[]string                                      `json:"labels,omitempty"`
	Notifications                *[]*pipelineDefinitionParameters.Notification  `json:"notifications,omitempty"`
	Schedules                    *[]*pipelineDefinitionParameters.Schedule      `json:"schedules,omitempty"`
	Tasks                        *[]pipelineDefinitionParameters.Task           `json:"tasks,omitempty"`
	TaskDependencies             *[]pipelineDefinitionParameters.TaskDependency `json:"task_dependencies,omitempty"`
}

type UpdatePipelineDefinitionInput struct {
	ResourceGroupID              *parameter.NullableInt64                       `json:"resource_group_id"`
	Name                         *string                                        `json:"name,omitempty"`
	Description                  *parameter.NullableString                      `json:"description,omitempty"`
	MaxTaskParallelism           *parameter.NullableInt64                       `json:"max_task_parallelism,omitempty"`
	ExecutionTimeout             *parameter.NullableInt64                       `json:"execution_timeout,omitempty"`
	MaxRetries                   *parameter.NullableInt64                       `json:"max_retries,omitempty"`
	MinRetryInterval             *parameter.NullableInt64                       `json:"min_retry_interval,omitempty"`
	IsConcurrentExecutionSkipped *parameter.NullableBool                        `json:"is_concurrent_execution_skipped,omitempty"`
	IsStoppedOnErrors            *parameter.NullableBool                        `json:"is_stopped_on_errors,omitempty"`
	Labels                       *[]string                                      `json:"labels,omitempty"`
	Notifications                *[]*pipelineDefinitionParameters.Notification  `json:"notifications,omitempty"`
	Schedules                    *[]*pipelineDefinitionParameters.Schedule      `json:"schedules,omitempty"`
	Tasks                        *[]pipelineDefinitionParameters.Task           `json:"tasks,omitempty"`
	TaskDependencies             *[]pipelineDefinitionParameters.TaskDependency `json:"task_dependencies,omitempty"`
}

func (c *TroccoClient) ListPipelineDefinitions(in *ListPipelineDefinitionsInput) (*entity.PipelineDefinitionList, error) {
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

	out := &entity.PipelineDefinitionList{}
	if err := c.do(http.MethodGet, url, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) GetPipelineDefinition(id int64) (*entity.PipelineDefinition, error) {
	url := fmt.Sprintf("/api/pipeline_definitions/%d", id)

	out := &entity.PipelineDefinition{}
	if err := c.do(http.MethodGet, url, nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) CreatePipelineDefinition(in *CreatePipelineDefinitionInput) (*entity.PipelineDefinition, error) {
	url := "/api/pipeline_definitions"

	out := &entity.PipelineDefinition{}
	if err := c.do(http.MethodPost, url, in, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) UpdatePipelineDefinition(id int64, in *UpdatePipelineDefinitionInput) (*entity.PipelineDefinition, error) {
	url := fmt.Sprintf("/api/pipeline_definitions/%d", id)

	out := &entity.PipelineDefinition{}
	if err := c.do(http.MethodPatch, url, in, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) DeletePipelineDefinition(id int64) error {
	url := fmt.Sprintf("/api/pipeline_definitions/%d", id)

	return c.do(http.MethodDelete, url, nil, nil)
}
