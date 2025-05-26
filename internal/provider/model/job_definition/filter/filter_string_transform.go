package filter

import (
	filterEntities "terraform-provider-trocco/internal/client/entity/job_definition/filter"
	"terraform-provider-trocco/internal/client/parameter/job_definition/filter"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

func (s FilterStringTransform) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"column_name": types.StringType,
		"type":        types.StringType,
	}
}
