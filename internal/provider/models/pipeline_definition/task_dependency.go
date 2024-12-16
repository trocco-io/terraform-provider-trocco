package workflow

import (
	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TaskDependency struct {
	Source      types.String `tfsdk:"source"`
	Destination types.String `tfsdk:"destination"`
}

func NewTaskDependencies(ens []*we.TaskDependency, keys map[int64]types.String) []*TaskDependency {
	if ens == nil {
		return nil
	}

	mds := []*TaskDependency{}
	for _, en := range ens {
		mds = append(mds, NewTaskDependency(en, keys))
	}

	return mds
}

func NewTaskDependency(en *we.TaskDependency, keys map[int64]types.String) *TaskDependency {
	if en == nil {
		return nil
	}

	return &TaskDependency{
		Source:      keys[en.Source],
		Destination: keys[en.Destination],
	}
}

// taskDependencies := []*wm.TaskDependency{}
// for _, d := range workflow.TaskDependencies {
// 	taskDependencies = append(taskDependencies, &wm.TaskDependency{
// 		Source:      keys[d.Source],
// 		Destination: keys[d.Destination],
// 	})
// }
