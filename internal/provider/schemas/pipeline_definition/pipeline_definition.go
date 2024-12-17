package workflow

import (
	"terraform-provider-trocco/internal/client"

	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
	wm "terraform-provider-trocco/internal/provider/models/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
)

type PipelineDefinition struct {
	ID               types.Int64          `tfsdk:"id"`
	Name             types.String         `tfsdk:"name"`
	Description      types.String         `tfsdk:"description"`
	Labels           []types.String       `tfsdk:"labels"`
	Notifications    []wm.Notification    `tfsdk:"notifications"`
	Schedules        []wm.Schedule        `tfsdk:"schedules"`
	Tasks            []*wm.Task           `tfsdk:"tasks"`
	TaskDependencies []*wm.TaskDependency `tfsdk:"task_dependencies"`
}

func (m *PipelineDefinition) ToCreateWorkflowInput() *client.CreateWorkflowInput {
	labels := []string{}
	for _, l := range m.Labels {
		labels = append(labels, l.ValueString())
	}

	notifications := []wp.Notification{}
	for _, n := range m.Notifications {
		notifications = append(notifications, n.ToInput())
	}

	schedules := []wp.Schedule{}
	for _, s := range m.Schedules {
		schedules = append(schedules, s.ToInput())
	}

	tasks := []wp.Task{}
	for _, t := range m.Tasks {
		tasks = append(tasks, *t.ToInput(map[string]int64{}))
	}

	taskDependencies := []wp.TaskDependency{}
	for _, d := range m.TaskDependencies {
		taskDependencies = append(taskDependencies, *d.ToInput())
	}

	return &client.CreateWorkflowInput{
		Name:             m.Name.ValueString(),
		Description:      m.Description.ValueStringPointer(),
		Labels:           lo.ToPtr(labels),
		Notifications:    lo.ToPtr(notifications),
		Schedules:        lo.ToPtr(schedules),
		Tasks:            tasks,
		TaskDependencies: taskDependencies,
	}
}

func (m *PipelineDefinition) ToUpdateWorkflowInput(st *PipelineDefinition) *client.UpdateWorkflowInput {
	labels := []string{}
	for _, l := range m.Labels {
		labels = append(labels, l.ValueString())
	}

	notifications := []wp.Notification{}
	for _, n := range m.Notifications {
		notifications = append(notifications, n.ToInput())
	}

	schedules := []wp.Schedule{}
	for _, s := range m.Schedules {
		schedules = append(schedules, s.ToInput())
	}

	identifiers := map[string]int64{}
	for _, t := range st.Tasks {
		identifiers[t.Key.ValueString()] = t.TaskIdentifier.ValueInt64()
	}

	tasks := []wp.Task{}
	for _, t := range m.Tasks {
		tasks = append(tasks, *t.ToInput(identifiers))
	}

	taskDependencies := []wp.TaskDependency{}
	for _, d := range m.TaskDependencies {
		taskDependencies = append(taskDependencies, *d.ToInput())
	}

	return &client.UpdateWorkflowInput{
		Name:             m.Name.ValueStringPointer(),
		Description:      m.Description.ValueStringPointer(),
		Labels:           lo.ToPtr(labels),
		Notifications:    lo.ToPtr(notifications),
		Schedules:        lo.ToPtr(schedules),
		Tasks:            tasks,
		TaskDependencies: taskDependencies,
	}
}
