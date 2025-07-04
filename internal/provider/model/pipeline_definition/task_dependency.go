package pipeline_definition

import (
	"strconv"
	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var taskDependencyAttrTypes = map[string]attr.Type{
	"source":      types.StringType,
	"destination": types.StringType,
}

var taskDependencyObjectType = types.ObjectType{
	AttrTypes: taskDependencyAttrTypes,
}

type TaskDependency struct {
	Source      types.String `tfsdk:"source"`
	Destination types.String `tfsdk:"destination"`
}

func NewTaskDependencies(ens []*we.TaskDependency, keys map[int64]types.String, previous *PipelineDefinition) types.Set {
	if ens == nil {
		return types.SetNull(taskDependencyObjectType)
	}

	// If the attribute in the plan (or state) is nil, the provider should sets nil to the state.
	var isTaskDependenciesNull bool
	if previous == nil {
		isTaskDependenciesNull = true
	} else {
		isTaskDependenciesNull = previous.TaskDependencies.IsNull()
	}

	if isTaskDependenciesNull && len(ens) == 0 {
		return types.SetNull(taskDependencyObjectType)
	}

	elements := []attr.Value{}
	for _, en := range ens {
		td := NewTaskDependency(en, keys)
		if td != nil {
			obj, _ := types.ObjectValue(
				taskDependencyAttrTypes,
				map[string]attr.Value{
					"source":      td.Source,
					"destination": td.Destination,
				},
			)
			elements = append(elements, obj)
		}
	}

	set, diags := types.SetValue(taskDependencyObjectType, elements)
	if diags.HasError() {
		return types.SetNull(taskDependencyObjectType)
	}
	return set
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
