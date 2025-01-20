package filter

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	filterEntities "terraform-provider-trocco/internal/client/entities/job_definitions/filter"
	"terraform-provider-trocco/internal/client/parameters/job_definitions/filter"
)

type FilterStringTransform struct {
	ColumnName types.String `tfsdk:"column_name"`
	Type       types.String `tfsdk:"type"`
}

func NewFilterStringTransforms(filterStringTransforms []filterEntities.FilterStringTransform) []FilterStringTransform {
	if len(filterStringTransforms) == 0 {
		return nil
	}

	outputs := make([]FilterStringTransform, 0, len(filterStringTransforms))
	for _, input := range filterStringTransforms {
		filterStringTransform := FilterStringTransform{
			ColumnName: types.StringValue(input.ColumnName),
			Type:       types.StringValue(input.Type),
		}
		outputs = append(outputs, filterStringTransform)
	}
	return outputs
}

func (filterStringTransform FilterStringTransform) ToInput() filter.FilterStringTransformInput {
	return filter.FilterStringTransformInput{
		ColumnName: filterStringTransform.ColumnName.ValueString(),
		Type:       filterStringTransform.Type.ValueString(),
	}
}
