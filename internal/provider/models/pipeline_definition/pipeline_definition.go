package pipeline_definition

import (
	"terraform-provider-trocco/internal/client"
	pdp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
	"terraform-provider-trocco/internal/provider/models"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
)

type PipelineDefinition struct {
	ID                           types.Int64       `tfsdk:"id"`
	ResourceGroupID              types.Int64       `tfsdk:"resource_group_id"`
	Name                         types.String      `tfsdk:"name"`
	Description                  types.String      `tfsdk:"description"`
	MaxTaskParallelism           types.Int64       `tfsdk:"max_task_parallelism"`
	ExecutionTimeout             types.Int64       `tfsdk:"execution_timeout"`
	MaxRetries                   types.Int64       `tfsdk:"max_retries"`
	MinRetryInterval             types.Int64       `tfsdk:"min_retry_interval"`
	IsConcurrentExecutionSkipped types.Bool        `tfsdk:"is_concurrent_execution_skipped"`
	IsStoppedOnErrors            types.Bool        `tfsdk:"is_stopped_on_errors"`
	Labels                       []types.String    `tfsdk:"labels"`
	Notifications                []Notification    `tfsdk:"notifications"`
	Schedules                    []Schedule        `tfsdk:"schedules"`
	Tasks                        []*Task           `tfsdk:"tasks"`
	TaskDependencies             []*TaskDependency `tfsdk:"task_dependencies"`
}

func (m *PipelineDefinition) ToCreateInput() *client.CreateWorkflowInput {
	labels := []string{}
	for _, l := range m.Labels {
		labels = append(labels, l.ValueString())
	}

	notifications := []pdp.Notification{}
	for _, n := range m.Notifications {
		notifications = append(notifications, n.ToInput())
	}

	schedules := []pdp.Schedule{}
	for _, s := range m.Schedules {
		schedules = append(schedules, s.ToInput())
	}

	tasks := []pdp.Task{}
	for _, t := range m.Tasks {
		tasks = append(tasks, *t.ToInput(map[string]int64{}))
	}

	taskDependencies := []pdp.TaskDependency{}
	for _, d := range m.TaskDependencies {
		taskDependencies = append(taskDependencies, pdp.TaskDependency{
			Source:      d.Source.ValueString(),
			Destination: d.Destination.ValueString(),
		})
	}

	return &client.CreateWorkflowInput{
		ResourceGroupID:              models.NewNullableInt64(m.ResourceGroupID),
		Name:                         m.Name.ValueString(),
		Description:                  m.Description.ValueStringPointer(),
		MaxTaskParallelism:           models.NewNullableInt64(m.MaxTaskParallelism),
		ExecutionTimeout:             models.NewNullableInt64(m.ExecutionTimeout),
		MaxRetries:                   models.NewNullableInt64(m.MaxRetries),
		MinRetryInterval:             models.NewNullableInt64(m.MinRetryInterval),
		IsConcurrentExecutionSkipped: models.NewNullableBool(m.IsConcurrentExecutionSkipped),
		IsStoppedOnErrors:            models.NewNullableBool(m.IsStoppedOnErrors),
		Labels:                       lo.ToPtr(labels),
		Notifications:                lo.ToPtr(notifications),
		Schedules:                    lo.ToPtr(schedules),
		Tasks:                        tasks,
		TaskDependencies:             taskDependencies,
	}
}

func (m *PipelineDefinition) ToUpdateWorkflowInput(state *PipelineDefinition) *client.UpdateWorkflowInput {
	labels := []string{}
	for _, l := range m.Labels {
		labels = append(labels, l.ValueString())
	}

	notifications := []pdp.Notification{}
	for _, n := range m.Notifications {
		notifications = append(notifications, n.ToInput())
	}

	schedules := []pdp.Schedule{}
	for _, s := range m.Schedules {
		schedules = append(schedules, s.ToInput())
	}

	stateTaskIdentifiers := map[string]int64{}
	for _, s := range state.Tasks {
		stateTaskIdentifiers[s.Key.ValueString()] = s.TaskIdentifier.ValueInt64()
	}

	tasks := []pdp.Task{}
	for _, t := range m.Tasks {
		tasks = append(tasks, *t.ToInput(stateTaskIdentifiers))
	}

	taskDependencies := []pdp.TaskDependency{}
	for _, d := range m.TaskDependencies {
		taskDependencies = append(taskDependencies, pdp.TaskDependency{
			Source:      d.Source.ValueString(),
			Destination: d.Destination.ValueString(),
		})
	}

	return &client.UpdateWorkflowInput{
		ResourceGroupID:              models.NewNullableInt64(m.ResourceGroupID),
		Name:                         m.Name.ValueStringPointer(),
		Description:                  m.Description.ValueStringPointer(),
		MaxTaskParallelism:           models.NewNullableInt64(m.MaxTaskParallelism),
		ExecutionTimeout:             models.NewNullableInt64(m.ExecutionTimeout),
		MaxRetries:                   models.NewNullableInt64(m.MaxRetries),
		MinRetryInterval:             models.NewNullableInt64(m.MinRetryInterval),
		IsConcurrentExecutionSkipped: models.NewNullableBool(m.IsConcurrentExecutionSkipped),
		IsStoppedOnErrors:            models.NewNullableBool(m.IsStoppedOnErrors),
		Labels:                       lo.ToPtr(labels),
		Notifications:                lo.ToPtr(notifications),
		Schedules:                    lo.ToPtr(schedules),
		Tasks:                        tasks,
		TaskDependencies:             taskDependencies,
	}
}