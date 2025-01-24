package filter

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/filter"
	filter2 "terraform-provider-trocco/internal/client/parameter/job_definitions/filter"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterHash struct {
	Name types.String `tfsdk:"name"`
}

func NewFilterHashes(filterHashes []filter.FilterHash) []FilterHash {
	if len(filterHashes) == 0 {
		return nil
	}
	outputs := make([]FilterHash, 0, len(filterHashes))
	for _, input := range filterHashes {
		filterHash := FilterHash{
			Name: types.StringValue(input.Name),
		}
		outputs = append(outputs, filterHash)
	}
	return outputs
}

func (filterHash FilterHash) ToInput() filter2.FilterHashInput {
	return filter2.FilterHashInput{
		Name: filterHash.Name.ValueString(),
	}
}
