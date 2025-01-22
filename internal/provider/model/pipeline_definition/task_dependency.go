package pipeline_definition

import (
	"strconv"
	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TaskDependency struct {
	Source      types.String `tfsdk:"source"`
	Destination types.String `tfsdk:"destination"`
}

func NewTaskDependencies(ens []*we.TaskDependency, keys map[int64]types.String, previous *PipelineDefinition) []*TaskDependency {
	if ens == nil {
		return nil
	}

	// If the attribute in the plan (or state) is nil, the provider should sets nil to the state.
	if previous.TaskDependencies == nil && len(ens) == 0 {
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

	// The keys are data that exist only in the provider, so they does not exist on import.
	// In this case, the provider uses the task identifiers as the keys.
	var source types.String
	var destination types.String
	if len(keys) == 0 {
		source = types.StringValue(strconv.FormatInt(en.Source, 10))
		destination = types.StringValue(strconv.FormatInt(en.Destination, 10))
	} else {
		source = keys[en.Source]
		destination = keys[en.Destination]
	}

	return &TaskDependency{
		Source:      source,
		Destination: destination,
	}
}

func (d *TaskDependency) ToInput() *wp.TaskDependency {
	return &wp.TaskDependency{
		Source:      d.Source.ValueString(),
		Destination: d.Destination.ValueString(),
	}
}
