package pipeline_definition

import (
	"context"
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

func NewIfElseTaskConfig(ctx context.Context, en *pipelineDefinitionEntities.IfElseTaskConfig) *IfElseTaskConfig {
	if en == nil {
		return nil
	}

	return &IfElseTaskConfig{
		Name:            types.StringValue(en.Name),
		ConditionGroups: NewConditionGroups(en.ConditionGroups),
		Destinations:    NewDestinations(en.Destinations),
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

func NewConditionGroups(en *pipelineDefinitionEntities.ConditionGroups) *ConditionGroups {
	if en == nil {
		return nil
	}

	conditions := []Condition{}
	for _, c := range en.Conditions {
		if c != nil {
			conditions = append(conditions, *NewCondition(c))
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

func NewCondition(en *pipelineDefinitionEntities.Condition) *Condition {
	if en == nil {
		return nil
	}

	taskKey := types.StringNull()
	if en.TaskKey != nil {
		taskKey = types.StringValue(*en.TaskKey)
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

func NewDestinations(en *pipelineDefinitionEntities.Destinations) *Destinations {
	if en == nil {
		return nil
	}

	ifList := types.ListNull(types.StringType)
	if en.If != nil {
		ifValues := []attr.Value{}
		for _, v := range en.If {
			ifValues = append(ifValues, types.StringValue(v))
		}
		ifList, _ = types.ListValue(types.StringType, ifValues)
	}

	elseList := types.ListNull(types.StringType)
	if en.Else != nil {
		elseValues := []attr.Value{}
		for _, v := range en.Else {
			elseValues = append(elseValues, types.StringValue(v))
		}
		elseList, _ = types.ListValue(types.StringType, elseValues)
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

	var ifDest []string
	if !d.If.IsNull() && !d.If.IsUnknown() {
		d.If.ElementsAs(ctx, &ifDest, false)
	}

	var elseDest []string
	if !d.Else.IsNull() && !d.Else.IsUnknown() {
		d.Else.ElementsAs(ctx, &elseDest, false)
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
