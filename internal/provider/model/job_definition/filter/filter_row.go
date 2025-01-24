package filter

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/filter"
	filter2 "terraform-provider-trocco/internal/client/parameter/job_definition/filter"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterRows struct {
	Condition           types.String         `tfsdk:"condition"`
	FilterRowConditions []filterRowCondition `tfsdk:"filter_row_conditions"`
}

type filterRowCondition struct {
	Column   types.String `tfsdk:"column"`
	Operator types.String `tfsdk:"operator"`
	Argument types.String `tfsdk:"argument"`
}

func NewFilterRows(filterRows *filter.FilterRows) *FilterRows {
	if filterRows == nil {
		return nil
	}
	conditions := make([]filterRowCondition, 0, len(filterRows.FilterRowConditions))
	for _, input := range filterRows.FilterRowConditions {
		condition := filterRowCondition{
			Column:   types.StringValue(input.Column),
			Operator: types.StringValue(input.Operator),
			Argument: types.StringValue(input.Argument),
		}
		conditions = append(conditions, condition)
	}
	return &FilterRows{
		Condition:           types.StringValue(filterRows.Condition),
		FilterRowConditions: conditions,
	}
}

func (filterRows *FilterRows) ToInput() *filter2.FilterRowsInput {
	if filterRows == nil {
		return nil
	}

	conditions := make([]filter2.FilterRowConditionInput, 0, len(filterRows.FilterRowConditions))
	for _, input := range filterRows.FilterRowConditions {
		condition := filter2.FilterRowConditionInput{
			Column:   input.Column.ValueString(),
			Operator: input.Operator.ValueString(),
			Argument: input.Argument.ValueString(),
		}
		conditions = append(conditions, condition)
	}
	return &filter2.FilterRowsInput{
		Condition:           filterRows.Condition.ValueString(),
		FilterRowConditions: conditions,
	}
}
