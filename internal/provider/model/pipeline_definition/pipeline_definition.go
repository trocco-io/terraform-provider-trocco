package pipeline_definition

import (
	"terraform-provider-trocco/internal/client"
	entity "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	pdp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
	model "terraform-provider-trocco/internal/provider/model"

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
	Notifications                []*Notification   `tfsdk:"notifications"`
	Schedules                    []*Schedule       `tfsdk:"schedules"`
	Tasks                        []*Task           `tfsdk:"tasks"`
	TaskDependencies             []*TaskDependency `tfsdk:"task_dependencies"`
}

func NewPipelineDefinition(en *entity.PipelineDefinition, keys map[int64]types.String, previous *PipelineDefinition) *PipelineDefinition {
	return &PipelineDefinition{
		ID:                           types.Int64Value(en.ID),
		ResourceGroupID:              types.Int64PointerValue(en.ResourceGroupID),
		Name:                         types.StringPointerValue(en.Name),
		Description:                  types.StringPointerValue(en.Description),
		MaxTaskParallelism:           types.Int64PointerValue(en.MaxTaskParallelism),
		ExecutionTimeout:             types.Int64PointerValue(en.ExecutionTimeout),
		MaxRetries:                   types.Int64PointerValue(en.MaxRetries),
		MinRetryInterval:             types.Int64PointerValue(en.MinRetryInterval),
		IsConcurrentExecutionSkipped: types.BoolPointerValue(en.IsConcurrentExecutionSkipped),
		IsStoppedOnErrors:            types.BoolPointerValue(en.IsStoppedOnErrors),
		Labels:                       NewLabels(en.Labels, previous.Labels == nil),
		Notifications:                NewNotifications(en.Notifications, previous),
		Schedules:                    NewSchedules(en.Schedules, previous),
		Tasks:                        NewTasks(en.Tasks, keys, previous),
		TaskDependencies:             NewTaskDependencies(en.TaskDependencies, keys, previous),
	}
}

func (m *PipelineDefinition) ToCreateInput() *client.CreatePipelineDefinitionInput {
	labels := []string{}
	for _, l := range m.Labels {
		labels = append(labels, l.ValueString())
	}

	notifications := []*pdp.Notification{}
	for _, n := range m.Notifications {
		notifications = append(notifications, n.ToInput())
	}

	schedules := []*pdp.Schedule{}
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

	return &client.CreatePipelineDefinitionInput{
		ResourceGroupID:              model.NewNullableInt64(m.ResourceGroupID),
		Name:                         m.Name.ValueString(),
		Description:                  model.NewNullableString(m.Description),
		MaxTaskParallelism:           model.NewNullableInt64(m.MaxTaskParallelism),
		ExecutionTimeout:             model.NewNullableInt64(m.ExecutionTimeout),
		MaxRetries:                   model.NewNullableInt64(m.MaxRetries),
		MinRetryInterval:             model.NewNullableInt64(m.MinRetryInterval),
		IsConcurrentExecutionSkipped: model.NewNullableBool(m.IsConcurrentExecutionSkipped),
		IsStoppedOnErrors:            model.NewNullableBool(m.IsStoppedOnErrors),
		Labels:                       lo.ToPtr(labels),
		Notifications:                lo.ToPtr(notifications),
		Schedules:                    lo.ToPtr(schedules),
		Tasks:                        lo.ToPtr(tasks),
		TaskDependencies:             lo.ToPtr(taskDependencies),
	}
}

func (m *PipelineDefinition) ToUpdateWorkflowInput(state *PipelineDefinition) *client.UpdatePipelineDefinitionInput {
	labels := []string{}
	for _, l := range m.Labels {
		labels = append(labels, l.ValueString())
	}

	notifications := []*pdp.Notification{}
	for _, n := range m.Notifications {
		notifications = append(notifications, n.ToInput())
	}

	schedules := []*pdp.Schedule{}
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

	return &client.UpdatePipelineDefinitionInput{
		ResourceGroupID:              model.NewNullableInt64(m.ResourceGroupID),
		Name:                         m.Name.ValueStringPointer(),
		Description:                  model.NewNullableString(m.Description),
		MaxTaskParallelism:           model.NewNullableInt64(m.MaxTaskParallelism),
		ExecutionTimeout:             model.NewNullableInt64(m.ExecutionTimeout),
		MaxRetries:                   model.NewNullableInt64(m.MaxRetries),
		MinRetryInterval:             model.NewNullableInt64(m.MinRetryInterval),
		IsConcurrentExecutionSkipped: model.NewNullableBool(m.IsConcurrentExecutionSkipped),
		IsStoppedOnErrors:            model.NewNullableBool(m.IsStoppedOnErrors),
		Labels:                       lo.ToPtr(labels),
		Notifications:                lo.ToPtr(notifications),
		Schedules:                    lo.ToPtr(schedules),
		Tasks:                        lo.ToPtr(tasks),
		TaskDependencies:             lo.ToPtr(taskDependencies),
	}
}