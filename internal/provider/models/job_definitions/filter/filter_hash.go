package filter

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions/filter"
)

type FilterHash struct {
	Name types.String `tfsdk:"name"`
}

func NewFilterHashes(filterHashes []filter.FilterHash) []FilterHash {
	outputs := make([]FilterHash, 0, len(filterHashes))
	for _, input := range filterHashes {
		filterHash := FilterHash{
			Name: types.StringValue(input.Name),
		}
		outputs = append(outputs, filterHash)
	}
	return outputs
}
