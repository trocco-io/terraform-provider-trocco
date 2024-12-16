package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ResourceTaskDependency struct {
	Source      types.String `tfsdk:"source"`
	Destination types.String `tfsdk:"destination"`
}

// func NewResourceTaskDependencies(
// 	tasks: []Task,
// ) []ResourceTaskDependency {
// 	keys := map[int64]types.String{}
// 	for _, t := range workflow.Tasks {
// 		keys[t.TaskIdentifier] = types.StringValue(t.Key)
// 	}

// 	return []ResourceTaskDependency{}
// }

// taskDependencies := []workflowResourceTaskDependencyModel{}
// for _, d := range workflow.TaskDependencies {
// 	taskDependencies = append(taskDependencies, workflowResourceTaskDependencyModel{
// 		Source:      keys[d.Source],
// 		Destination: keys[d.Destination],
// 	})
// }
