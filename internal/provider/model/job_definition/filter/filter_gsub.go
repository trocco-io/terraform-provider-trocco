package filter

import (
	filterEntities "terraform-provider-trocco/internal/client/entity/job_definition/filter"
	"terraform-provider-trocco/internal/client/parameter/job_definition/filter"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterGsub struct {
	ColumnName types.String `tfsdk:"column_name"`
	Pattern    types.String `tfsdk:"pattern"`
	To         types.String `tfsdk:"to"`
}

func NewFilterGsub(filterGsubs []filterEntities.FilterGsub) []FilterGsub {
	if len(filterGsubs) == 0 {
		return nil
	}

	outputs := make([]FilterGsub, 0, len(filterGsubs))
	for _, input := range filterGsubs {
		filterGsub := FilterGsub{
			ColumnName: types.StringValue(input.ColumnName),
			Pattern:    types.StringValue(input.Pattern),
			To:         types.StringValue(input.To),
		}
		outputs = append(outputs, filterGsub)
	}
	return outputs
}

func (filterGsub FilterGsub) ToInput() filter.FilterGsubInput {
	return filter.FilterGsubInput{
		ColumnName: filterGsub.ColumnName.ValueString(),
		Pattern:    filterGsub.Pattern.ValueString(),
		To:         filterGsub.To.ValueString(),
	}
}
