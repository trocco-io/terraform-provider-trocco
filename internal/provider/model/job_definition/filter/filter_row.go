package filter

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity/job_definition/filter"
	filterParameters "terraform-provider-trocco/internal/client/parameter/job_definition/filter"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterRows struct {
	Condition           types.String `tfsdk:"condition"`
	FilterRowConditions types.List   `tfsdk:"filter_row_conditions"`
}

type filterRowCondition struct {
	Column   types.String `tfsdk:"column"`
	Operator types.String `tfsdk:"operator"`
	Argument types.String `tfsdk:"argument"`
}

func (filterRowCondition) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"column":   types.StringType,
		"operator": types.StringType,
		"argument": types.StringType,
	}
}

func NewFilterRows(ctx context.Context, filterRows *filter.FilterRows) *FilterRows {
	if filterRows == nil {
		return nil
	}

	result := &FilterRows{
		Condition: types.StringValue(filterRows.Condition),
	}

	conditions, err := newFilterRowConditions(ctx, filterRows.FilterRowConditions)
	if err != nil {
		return nil
	}
	result.FilterRowConditions = conditions

	return result
}

func newFilterRowConditions(
	ctx context.Context,
	conditions []filter.FilterRowCondition,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: filterRowCondition{}.attrTypes(),
	}

	if conditions == nil {
		return types.ListNull(objectType), nil
	}

	filterConditions := make([]filterRowCondition, 0, len(conditions))
	for _, input := range conditions {
		condition := filterRowCondition{
			Column:   types.StringValue(input.Column),
			Operator: types.StringValue(input.Operator),
			Argument: types.StringValue(input.Argument),
		}
		filterConditions = append(filterConditions, condition)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, filterConditions)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert filter row conditions to ListValue: %v", diags)
	}
	return listValue, nil
}

func (filterRows *FilterRows) ToInput(ctx context.Context) *filterParameters.FilterRowsInput {
	if filterRows == nil {
		return nil
	}

	var conditionValues []filterRowCondition
	if !filterRows.FilterRowConditions.IsNull() && !filterRows.FilterRowConditions.IsUnknown() {
		diags := filterRows.FilterRowConditions.ElementsAs(ctx, &conditionValues, false)
		if diags.HasError() {
			return nil
		}
	}

	conditions := make([]filterParameters.FilterRowConditionInput, 0, len(conditionValues))
	for _, input := range conditionValues {
		condition := filterParameters.FilterRowConditionInput{
			Column:   input.Column.ValueString(),
			Operator: input.Operator.ValueString(),
			Argument: input.Argument.ValueString(),
		}
		conditions = append(conditions, condition)
	}
	return &filterParameters.FilterRowsInput{
		Condition:           filterRows.Condition.ValueString(),
		FilterRowConditions: conditions,
	}
}
