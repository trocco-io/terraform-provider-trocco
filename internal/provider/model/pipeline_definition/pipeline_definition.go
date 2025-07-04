package pipeline_definition

import (
	"context"
	"terraform-provider-trocco/internal/client"
	entity "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	pdp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
	"terraform-provider-trocco/internal/provider/custom_type"
	model "terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
)

type PipelineDefinition struct {
	ID                           types.Int64                    `tfsdk:"id"`
	ResourceGroupID              types.Int64                    `tfsdk:"resource_group_id"`
	Name                         types.String                   `tfsdk:"name"`
	Description                  custom_type.TrimmedStringValue `tfsdk:"description"`
	MaxTaskParallelism           types.Int64                    `tfsdk:"max_task_parallelism"`
	ExecutionTimeout             types.Int64                    `tfsdk:"execution_timeout"`
	MaxRetries                   types.Int64                    `tfsdk:"max_retries"`
	MinRetryInterval             types.Int64                    `tfsdk:"min_retry_interval"`
	IsConcurrentExecutionSkipped types.Bool                     `tfsdk:"is_concurrent_execution_skipped"`
	IsStoppedOnErrors            types.Bool                     `tfsdk:"is_stopped_on_errors"`
	Labels                       types.Set                      `tfsdk:"labels"`
	Notifications                types.Set                      `tfsdk:"notifications"`
	Schedules                    types.Set                      `tfsdk:"schedules"`
	Tasks                        types.Set                      `tfsdk:"tasks"`
	TaskDependencies             types.Set                      `tfsdk:"task_dependencies"`
}

func NewPipelineDefinition(ctx context.Context, en *entity.PipelineDefinition, keys map[int64]types.String, previous *PipelineDefinition) *PipelineDefinition {
	var notifications types.Set
	if previous == nil {
		notifications = NewNotifications(ctx, en.Notifications, true)
	} else {
		notifications = NewNotifications(ctx, en.Notifications, previous.Notifications.IsNull())
	}

	var labels types.Set
	if previous == nil {
		labels = NewLabels(ctx, en.Labels, true)
	} else {
		labels = NewLabels(ctx, en.Labels, previous.Labels.IsNull())
	}

	return &PipelineDefinition{
		ID:                           types.Int64Value(en.ID),
		ResourceGroupID:              types.Int64PointerValue(en.ResourceGroupID),
		Name:                         types.StringPointerValue(en.Name),
		Description:                  custom_type.TrimmedStringValue{StringValue: types.StringPointerValue(en.Description)},
		MaxTaskParallelism:           types.Int64PointerValue(en.MaxTaskParallelism),
		ExecutionTimeout:             types.Int64PointerValue(en.ExecutionTimeout),
		MaxRetries:                   types.Int64PointerValue(en.MaxRetries),
		MinRetryInterval:             types.Int64PointerValue(en.MinRetryInterval),
		IsConcurrentExecutionSkipped: types.BoolPointerValue(en.IsConcurrentExecutionSkipped),
		IsStoppedOnErrors:            types.BoolPointerValue(en.IsStoppedOnErrors),
		Labels:                       labels,
		Notifications:                notifications,
		Schedules:                    NewSchedules(ctx, en.Schedules, previous),
		Tasks:                        NewTasks(ctx, en.Tasks, keys, previous),
		TaskDependencies:             NewTaskDependencies(en.TaskDependencies, keys, previous),
	}
}

func (m *PipelineDefinition) ToCreateInput(ctx context.Context) *client.CreatePipelineDefinitionInput {
	labels := []string{}
	if !m.Labels.IsNull() && !m.Labels.IsUnknown() {
		var labelValues []types.String
		diags := m.Labels.ElementsAs(ctx, &labelValues, false)
		if !diags.HasError() {
			for _, l := range labelValues {
				labels = append(labels, l.ValueString())
			}
		}
	}

	notifications := []*pdp.Notification{}
	if !m.Notifications.IsNull() && !m.Notifications.IsUnknown() {
		var notificationValues []*Notification
		diags := m.Notifications.ElementsAs(ctx, &notificationValues, false)
		if !diags.HasError() {
			for _, n := range notificationValues {
				notifications = append(notifications, n.ToInput())
			}
		}
	}

	schedules := []*pdp.Schedule{}
	if !m.Schedules.IsNull() && !m.Schedules.IsUnknown() {
		var scheduleValues []*Schedule
		diags := m.Schedules.ElementsAs(ctx, &scheduleValues, false)
		if !diags.HasError() {
			for _, s := range scheduleValues {
				schedules = append(schedules, s.ToInput())
			}
		}
	}

	tasks := []pdp.Task{}
	if !m.Tasks.IsNull() && !m.Tasks.IsUnknown() {
		var tfTasks []*Task
		if diags := m.Tasks.ElementsAs(ctx, &tfTasks, false); diags.HasError() {
			return nil
		}

		for _, t := range tfTasks {
			tasks = append(tasks, *t.ToInput(ctx, map[string]int64{}))
		}
	}

	taskDependencies := []pdp.TaskDependency{}
	if !m.TaskDependencies.IsNull() && !m.TaskDependencies.IsUnknown() {
		var taskDependencyValues []*TaskDependency
		diags := m.TaskDependencies.ElementsAs(ctx, &taskDependencyValues, false)
		if !diags.HasError() {
			for _, d := range taskDependencyValues {
				taskDependencies = append(taskDependencies, pdp.TaskDependency{
					Source:      d.Source.ValueString(),
					Destination: d.Destination.ValueString(),
				})
			}
		}
	}

	return &client.CreatePipelineDefinitionInput{
		ResourceGroupID:              model.NewNullableInt64(m.ResourceGroupID),
		Name:                         m.Name.ValueString(),
		Description:                  model.NewNullableString(m.Description.StringValue),
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

func (m *PipelineDefinition) ToUpdateWorkflowInput(ctx context.Context, state *PipelineDefinition) *client.UpdatePipelineDefinitionInput {
	labels := []string{}
	if !m.Labels.IsNull() && !m.Labels.IsUnknown() {
		var labelValues []types.String
		diags := m.Labels.ElementsAs(ctx, &labelValues, false)
		if !diags.HasError() {
			for _, l := range labelValues {
				labels = append(labels, l.ValueString())
			}
		}
	}

	notifications := []*pdp.Notification{}
	if !m.Notifications.IsNull() && !m.Notifications.IsUnknown() {
		var notificationValues []*Notification
		diags := m.Notifications.ElementsAs(ctx, &notificationValues, false)
		if !diags.HasError() {
			for _, n := range notificationValues {
				notifications = append(notifications, n.ToInput())
			}
		}
	}

	schedules := []*pdp.Schedule{}
	if !m.Schedules.IsNull() && !m.Schedules.IsUnknown() {
		var scheduleValues []*Schedule
		diags := m.Schedules.ElementsAs(ctx, &scheduleValues, false)
		if !diags.HasError() {
			for _, s := range scheduleValues {
				schedules = append(schedules, s.ToInput())
			}
		}
	}

	stateTaskIdentifiers := map[string]int64{}
	if !state.Tasks.IsNull() && !state.Tasks.IsUnknown() {
		var stateTasks []*Task
		if diags := state.Tasks.ElementsAs(ctx, &stateTasks, false); diags.HasError() {
			return nil
		}

		for _, s := range stateTasks {
			stateTaskIdentifiers[s.Key.ValueString()] = s.TaskIdentifier.ValueInt64()
		}
	}

	tasks := []pdp.Task{}
	if !m.Tasks.IsNull() && !m.Tasks.IsUnknown() {
		var tfTasks []*Task
		if diags := m.Tasks.ElementsAs(ctx, &tfTasks, false); diags.HasError() {
			return nil
		}

		for _, t := range tfTasks {
			tasks = append(tasks, *t.ToInput(ctx, stateTaskIdentifiers))
		}
	}

	taskDependencies := []pdp.TaskDependency{}
	if !m.TaskDependencies.IsNull() && !m.TaskDependencies.IsUnknown() {
		var taskDependencyValues []*TaskDependency
		diags := m.TaskDependencies.ElementsAs(ctx, &taskDependencyValues, false)
		if !diags.HasError() {
			for _, d := range taskDependencyValues {
				taskDependencies = append(taskDependencies, pdp.TaskDependency{
					Source:      d.Source.ValueString(),
					Destination: d.Destination.ValueString(),
				})
			}
		}
	}

	return &client.UpdatePipelineDefinitionInput{
		ResourceGroupID:              model.NewNullableInt64(m.ResourceGroupID),
		Name:                         m.Name.ValueStringPointer(),
		Description:                  model.NewNullableString(m.Description.StringValue),
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
