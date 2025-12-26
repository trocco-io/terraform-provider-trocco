package pipeline_definition

import (
	"context"
	"strconv"
	pipelineDefinitionEntities "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	pipelineDefinitionParameters "terraform-provider-trocco/internal/client/parameter/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
)

type IfElseTaskConfig struct {
	Name            types.String     `tfsdk:"name"`
	ConditionGroups *ConditionGroups `tfsdk:"condition_groups"`
	Destinations    *Destinations    `tfsdk:"destinations"`
}

func NewIfElseTaskConfig(ctx context.Context, en *pipelineDefinitionEntities.IfElseTaskConfig, keys map[int64]types.String) *IfElseTaskConfig {
	if en == nil {
		return nil
	}

	return &IfElseTaskConfig{
		Name:            types.StringValue(en.Name),
		ConditionGroups: NewConditionGroups(en.ConditionGroups, keys),
		Destinations:    NewDestinations(en.Destinations, keys),
	}
}

func (c *IfElseTaskConfig) ToInput(ctx context.Context) *pipelineDefinitionParameters.IfElseTaskConfig {
	if c == nil {
		return nil
	}

	return &pipelineDefinitionParameters.IfElseTaskConfig{
		Name:            c.Name.ValueString(),
		ConditionGroups: c.ConditionGroups.ToInput(),
		Destinations:    c.Destinations.ToInput(ctx),
	}
}

type ConditionGroups struct {
	SetType    types.String `tfsdk:"set_type"`
	Conditions []Condition  `tfsdk:"conditions"`
}

func NewConditionGroups(en *pipelineDefinitionEntities.ConditionGroups, keys map[int64]types.String) *ConditionGroups {
	if en == nil {
		return nil
	}

	conditions := []Condition{}
	for _, c := range en.Conditions {
		if c != nil {
			conditions = append(conditions, *NewCondition(c, keys))
		}
	}

	return &ConditionGroups{
		SetType:    types.StringValue(en.SetType),
		Conditions: conditions,
	}
}

func (c *ConditionGroups) ToInput() *pipelineDefinitionParameters.ConditionGroups {
	if c == nil {
		return nil
	}

	conditions := []*pipelineDefinitionParameters.Condition{}
	for _, condition := range c.Conditions {
		conditions = append(conditions, condition.ToInput())
	}

	return &pipelineDefinitionParameters.ConditionGroups{
		SetType:    c.SetType.ValueString(),
		Conditions: conditions,
	}
}

type Condition struct {
	Variable types.String `tfsdk:"variable"`
	TaskKey  types.String `tfsdk:"task_key"`
	Operator types.String `tfsdk:"operator"`
	Value    types.String `tfsdk:"value"`
}

func NewCondition(en *pipelineDefinitionEntities.Condition, keys map[int64]types.String) *Condition {
	if en == nil {
		return nil
	}

	// API response returns task identifier as "identifier" field (string like "3")
	// Convert it to task key using keys map
	taskKey := types.StringNull()
	if en.Identifier != nil {
		if id, err := strconv.ParseInt(*en.Identifier, 10, 64); err == nil {
			if key, ok := keys[id]; ok {
				taskKey = key
			}
		}
	}

	return &Condition{
		Variable: types.StringValue(en.Variable),
		TaskKey:  taskKey,
		Operator: types.StringValue(en.Operator),
		Value:    types.StringValue(en.Value),
	}
}

func (c *Condition) ToInput() *pipelineDefinitionParameters.Condition {
	var taskKey *string
	if !c.TaskKey.IsNull() {
		taskKey = lo.ToPtr(c.TaskKey.ValueString())
	}

	return &pipelineDefinitionParameters.Condition{
		Variable: c.Variable.ValueString(),
		TaskKey:  taskKey,
		Operator: c.Operator.ValueString(),
		Value:    c.Value.ValueString(),
	}
}

type Destinations struct {
	If   types.List `tfsdk:"if"`
	Else types.List `tfsdk:"else"`
}

func NewDestinations(en *pipelineDefinitionEntities.Destinations, keys map[int64]types.String) *Destinations {
	if en == nil {
		return nil
	}

	// API response returns task identifiers as strings (e.g., "3")
	// Convert them to task keys using keys map
	//
	// When API response has values, convert them.
	// When API response is nil/empty, return empty list (user must specify [] explicitly).
	ifList := types.ListValueMust(types.StringType, []attr.Value{})
	if en.If != nil && len(en.If) > 0 {
		ifValues := []attr.Value{}
		for _, v := range en.If {
			if id, err := strconv.ParseInt(v, 10, 64); err == nil {
				if key, ok := keys[id]; ok {
					ifValues = append(ifValues, key)
				}
			}
		}
		ifList = types.ListValueMust(types.StringType, ifValues)
	}

	elseList := types.ListValueMust(types.StringType, []attr.Value{})
	if en.Else != nil && len(en.Else) > 0 {
		elseValues := []attr.Value{}
		for _, v := range en.Else {
			if id, err := strconv.ParseInt(v, 10, 64); err == nil {
				if key, ok := keys[id]; ok {
					elseValues = append(elseValues, key)
				}
			}
		}
		elseList = types.ListValueMust(types.StringType, elseValues)
	}

	return &Destinations{
		If:   ifList,
		Else: elseList,
	}
}

func (d *Destinations) ToInput(ctx context.Context) *pipelineDefinitionParameters.Destinations {
	if d == nil {
		return nil
	}

	// Only set non-empty slices; nil slices are omitted from JSON due to omitempty
	var ifDest []string
	if !d.If.IsNull() && !d.If.IsUnknown() {
		d.If.ElementsAs(ctx, &ifDest, false)
		if len(ifDest) == 0 {
			ifDest = nil
		}
	}

	var elseDest []string
	if !d.Else.IsNull() && !d.Else.IsUnknown() {
		d.Else.ElementsAs(ctx, &elseDest, false)
		if len(elseDest) == 0 {
			elseDest = nil
		}
	}

	return &pipelineDefinitionParameters.Destinations{
		If:   ifDest,
		Else: elseDest,
	}
}

func ConditionAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"variable": types.StringType,
		"task_key": types.StringType,
		"operator": types.StringType,
		"value":    types.StringType,
	}
}

func ConditionGroupsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"set_type":   types.StringType,
		"conditions": types.ListType{ElemType: types.ObjectType{AttrTypes: ConditionAttrTypes()}},
	}
}

func DestinationsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"if":   types.ListType{ElemType: types.StringType},
		"else": types.ListType{ElemType: types.StringType},
	}
}

func IfElseTaskConfigAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":             types.StringType,
		"condition_groups": types.ObjectType{AttrTypes: ConditionGroupsAttrTypes()},
		"destinations":     types.ObjectType{AttrTypes: DestinationsAttrTypes()},
	}
}
