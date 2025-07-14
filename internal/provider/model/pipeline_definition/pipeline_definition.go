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
	labels, labelsOk := convertStringSet(ctx, m.Labels)
	notifications, notificationsOk := convertNotificationSet(ctx, m.Notifications)
	schedules, schedulesOk := convertScheduleSet(ctx, m.Schedules)
	taskDependencies, taskDepsOk := convertTaskDependencySet(ctx, m.TaskDependencies)

	// Return nil if any conversion failed
	if !labelsOk || !notificationsOk || !schedulesOk || !taskDepsOk {
		return nil
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
	labels, labelsOk := convertStringSet(ctx, m.Labels)
	notifications, notificationsOk := convertNotificationSet(ctx, m.Notifications)
	schedules, schedulesOk := convertScheduleSet(ctx, m.Schedules)
	taskDependencies, taskDepsOk := convertTaskDependencySet(ctx, m.TaskDependencies)

	// Return nil if any conversion failed
	if !labelsOk || !notificationsOk || !schedulesOk || !taskDepsOk {
		return nil
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

// Helper functions for pipeline definition to reduce code duplication

// convertStringSet converts types.Set to string slice  
func convertStringSet(ctx context.Context, source types.Set) ([]string, bool) {
	if source.IsNull() || source.IsUnknown() {
		return []string{}, true
	}

	var values []types.String
	if diags := source.ElementsAs(ctx, &values, false); !diags.HasError() {
		result := make([]string, 0, len(values))
		for _, v := range values {
			result = append(result, v.ValueString())
		}
		return result, true
	}
	return []string{}, false
}

// convertNotificationSet converts notification set
func convertNotificationSet(ctx context.Context, source types.Set) ([]*pdp.Notification, bool) {
	if source.IsNull() || source.IsUnknown() {
		return []*pdp.Notification{}, true
	}

	var notificationValues []*Notification
	if diags := source.ElementsAs(ctx, &notificationValues, false); !diags.HasError() {
		result := make([]*pdp.Notification, 0, len(notificationValues))
		for _, n := range notificationValues {
			result = append(result, n.ToInput())
		}
		return result, true
	}
	return []*pdp.Notification{}, false
}

// convertScheduleSet converts schedule set
func convertScheduleSet(ctx context.Context, source types.Set) ([]*pdp.Schedule, bool) {
	if source.IsNull() || source.IsUnknown() {
		return []*pdp.Schedule{}, true
	}

	var scheduleValues []*Schedule
	if diags := source.ElementsAs(ctx, &scheduleValues, false); !diags.HasError() {
		result := make([]*pdp.Schedule, 0, len(scheduleValues))
		for _, s := range scheduleValues {
			result = append(result, s.ToInput())
		}
		return result, true
	}
	return []*pdp.Schedule{}, false
}

// convertTaskDependencySet converts task dependency set
func convertTaskDependencySet(ctx context.Context, source types.Set) ([]pdp.TaskDependency, bool) {
	if source.IsNull() || source.IsUnknown() {
		return []pdp.TaskDependency{}, true
	}

	var taskDependencyValues []*TaskDependency
	if diags := source.ElementsAs(ctx, &taskDependencyValues, false); !diags.HasError() {
		result := make([]pdp.TaskDependency, 0, len(taskDependencyValues))
		for _, d := range taskDependencyValues {
			result = append(result, pdp.TaskDependency{
				Source:      d.Source.ValueString(),
				Destination: d.Destination.ValueString(),
			})
		}
		return result, true
	}
	return []pdp.TaskDependency{}, false
}
