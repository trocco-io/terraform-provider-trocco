package filter

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	filterEntities "terraform-provider-trocco/internal/client/entities/job_definitions/filter"
)

type FilterMask struct {
	Name       types.String `tfsdk:"name"`
	MaskType   types.Int32  `tfsdk:"mask_type"`
	Length     types.Int64  `tfsdk:"length"`
	Pattern    types.String `tfsdk:"pattern"`
	StartIndex types.Int64  `tfsdk:"start_index"`
	EndIndex   types.Int64  `tfsdk:"end_index"`
}

func NewFilterMasks(filterMasks []filterEntities.FilterMask) []FilterMask {
	outputs := make([]FilterMask, 0, len(filterMasks))
	for _, input := range filterMasks {
		filterMask := FilterMask{
			Name:       types.StringValue(input.Name),
			MaskType:   types.Int32Value(input.MaskType),
			Length:     types.Int64PointerValue(input.Length),
			Pattern:    types.StringPointerValue(input.Pattern),
			StartIndex: types.Int64PointerValue(input.StartIndex),
			EndIndex:   types.Int64PointerValue(input.EndIndex),
		}
		outputs = append(outputs, filterMask)
	}
	return outputs
}
