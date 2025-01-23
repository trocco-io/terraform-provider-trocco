package filter

import (
	filterEntities "terraform-provider-trocco/internal/client/entities/job_definitions/filter"
	filter2 "terraform-provider-trocco/internal/client/parameter/job_definitions/filter"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterMask struct {
	Name       types.String `tfsdk:"name"`
	MaskType   types.String `tfsdk:"mask_type"`
	Length     types.Int64  `tfsdk:"length"`
	Pattern    types.String `tfsdk:"pattern"`
	StartIndex types.Int64  `tfsdk:"start_index"`
	EndIndex   types.Int64  `tfsdk:"end_index"`
}

func NewFilterMasks(filterMasks []filterEntities.FilterMask) []FilterMask {
	if len(filterMasks) == 0 {
		return nil
	}

	outputs := make([]FilterMask, 0, len(filterMasks))
	for _, input := range filterMasks {
		filterMask := FilterMask{
			Name:       types.StringValue(input.Name),
			MaskType:   types.StringValue(input.MaskType),
			Length:     types.Int64PointerValue(input.Length),
			Pattern:    types.StringPointerValue(input.Pattern),
			StartIndex: types.Int64PointerValue(input.StartIndex),
			EndIndex:   types.Int64PointerValue(input.EndIndex),
		}
		outputs = append(outputs, filterMask)
	}
	return outputs
}

func (filterMask FilterMask) ToInput() filter2.FilterMaskInput {
	input := filter2.FilterMaskInput{
		Name:       filterMask.Name.ValueString(),
		MaskType:   filterMask.MaskType.ValueString(),
		Length:     filterMask.Length.ValueInt64Pointer(),
		Pattern:    filterMask.Pattern.ValueStringPointer(),
		StartIndex: filterMask.StartIndex.ValueInt64Pointer(),
		EndIndex:   filterMask.EndIndex.ValueInt64Pointer(),
	}
	return input
}
